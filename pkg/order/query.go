package order

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	order1 "github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/order"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"

	constant "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"

	ordermgrpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order/order"
	orderstatepb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order/state"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	"github.com/google/uuid"
)

// nolint
func GetOrder(ctx context.Context, id string) (info *npool.Order, err error) {
	infos := []*npool.Order{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return cli.
			Order.
			Query().
			Select(
				order1.FieldID,
				order1.FieldUserID,
				order1.FieldGoodID,
				order1.FieldUnits,
				order1.FieldOrderType,
				order1.FieldParentOrderID,
				order1.FieldCouponID,
				order1.FieldDiscountCouponID,
				order1.FieldUserSpecialReductionID,
				order1.FieldCreateAt,
				order1.FieldPayWithParent,
				order1.FieldStart,
				order1.FieldEnd,
			).
			Where(
				order1.ID(uuid.MustParse(id)),
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(payment.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(order1.FieldID),
						t1.C(payment.FieldOrderID),
					).
					AppendSelect(
						sql.As(t1.C(payment.FieldID), "payment_id"),
						sql.As(t1.C(payment.FieldCoinInfoID), "payment_coin_type_id"),
						sql.As(t1.C(payment.FieldCoinUsdCurrency), "payment_coin_usd_currency"),
						sql.As(t1.C(payment.FieldLiveCoinUsdCurrency), "payment_live_coin_usd_currency"),
						sql.As(t1.C(payment.FieldLocalCoinUsdCurrency), "payment_local_coin_usd_currency"),
						sql.As(t1.C(payment.FieldAccountID), "payment_account_id"),
						sql.As(t1.C(payment.FieldAmount), "payment_amount"),
						sql.As(t1.C(payment.FieldState), "payment_state"),
						sql.As(t1.C(payment.FieldPayWithBalanceAmount), "pay_with_balance_amount"),
						sql.As(t1.C(payment.FieldUpdateAt), "paid_at"),
					)
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	info = infos[0]
	info.PayWithParent = info.PayWithParentInt != 0

	switch info.OrderType {
	case constant.OrderTypeNormal:
		info.Type = ordermgrpb.OrderType_Normal
	case constant.OrderTypeOffline:
		info.Type = ordermgrpb.OrderType_Offline
	case constant.OrderTypeAirdrop:
		info.Type = ordermgrpb.OrderType_Airdrop
	case ordermgrpb.OrderType_Normal.String():
		info.Type = ordermgrpb.OrderType_Normal
	case ordermgrpb.OrderType_Offline.String():
		info.Type = ordermgrpb.OrderType_Offline
	case ordermgrpb.OrderType_Airdrop.String():
		info.Type = ordermgrpb.OrderType_Airdrop
	default:
		info.Type = ordermgrpb.OrderType_Normal
	}

	switch info.PaymentState {
	case constant.PaymentStateDone:
		info.State = orderstatepb.EState_Paid
	case constant.PaymentStateTimeout:
		info.State = orderstatepb.EState_PaymentTimeout
	case constant.PaymentStateCanceled:
		info.State = orderstatepb.EState_Canceled
	case constant.PaymentStateWait:
		info.State = orderstatepb.EState_WaitPayment
	case orderstatepb.EState_WaitPayment.String():
		info.State = orderstatepb.EState_WaitPayment
	case orderstatepb.EState_Paid.String():
		info.State = orderstatepb.EState_Paid
	case orderstatepb.EState_PaymentTimeout.String():
		info.State = orderstatepb.EState_PaymentTimeout
	case orderstatepb.EState_Canceled.String():
		info.State = orderstatepb.EState_Canceled
	case orderstatepb.EState_InService.String():
		info.State = orderstatepb.EState_InService
	case orderstatepb.EState_Expired.String():
		info.State = orderstatepb.EState_Expired
	default:
		info.State = orderstatepb.EState_WaitPayment
	}

	now := uint32(time.Now().Unix())
	switch info.State {
	case orderstatepb.EState_Paid:
		if info.Start >= now {
			info.State = orderstatepb.EState_InService
		}
		if now > info.End {
			info.State = orderstatepb.EState_Expired
		}
	}

	invalidID := uuid.UUID{}.String()
	if info.PaymentID == invalidID {
		info.State = orderstatepb.EState_WaitPayment
		if now > info.CreatedAt+constant.TimeoutSeconds {
			info.State = orderstatepb.EState_PaymentTimeout
		}
	}

	const accuracy = 1000000000000
	damount := func(amount uint64) decimal.Decimal {
		return decimal.NewFromInt(int64(amount)).
			Div(decimal.NewFromInt(int64(accuracy)))
	}

	info.PaymentCoinUSDCurrency = damount(info.PaymentCoinUSDCurrencyUint).String()
	info.PaymentLiveCoinUSDCurrency = damount(info.PaymentLiveUSDCurrencyUint).String()
	info.PaymentLocalCoinUSDCurrency = damount(info.PaymentLocalUSDCurrencyUint).String()
	info.PaymentAmount = damount(info.PaymentAmountUint).String()

	return info, nil
}
