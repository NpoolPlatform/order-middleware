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

func GetOrder(ctx context.Context, id string) (info *npool.Order, err error) {
	infos := []*npool.Order{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Order.
			Query().
			Where(
				order1.ID(uuid.MustParse(id)),
			)

		return join(stm).
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

	return post(infos[0]), nil
}

func GetOrders(ctx context.Context, appID, userID string, offset, limit int32) (infos []*npool.Order, total uint32, err error) {
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Order.
			Query().
			Where(
				order1.AppID(uuid.MustParse(appID)),
				order1.UserID(uuid.MustParse(userID)),
			)

		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		stm = stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	for i, info := range infos {
		infos[i] = post(info)
	}

	return infos, total, nil
}

func join(stm *ent.OrderQuery) *ent.OrderSelect {
	return stm.
		Select(
			order1.FieldID,
			order1.FieldAppID,
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
					sql.As(t1.C(payment.FieldUserSetCanceled), "user_canceled"),
				)

			t2 := sql.Table(order1.Table)
			s.
				LeftJoin(t2).
				On(
					s.C(order1.FieldParentOrderID),
					t2.C(order1.FieldID),
				).
				AppendSelect(
					sql.As(t2.C(order1.FieldGoodID), "parent_order_good_id"),
				)
		})
}

func Join(stm *ent.OrderQuery) *ent.OrderSelect {
	return join(stm)
}

func post(info *npool.Order) *npool.Order { //nolint
	info.PayWithParent = info.PayWithParentInt != 0

	switch info.OrderTypeStr {
	case constant.OrderTypeNormal:
		info.OrderType = ordermgrpb.OrderType_Normal
	case ordermgrpb.OrderType_Normal.String():
		info.OrderType = ordermgrpb.OrderType_Normal

	case constant.OrderTypeOffline:
		info.OrderType = ordermgrpb.OrderType_Offline
	case ordermgrpb.OrderType_Offline.String():
		info.OrderType = ordermgrpb.OrderType_Offline

	case constant.OrderTypeAirdrop:
		info.OrderType = ordermgrpb.OrderType_Airdrop
	case ordermgrpb.OrderType_Airdrop.String():
		info.OrderType = ordermgrpb.OrderType_Airdrop

	default:
		info.OrderType = ordermgrpb.OrderType_Normal
	}

	// TODO: state should from order state table
	switch info.PaymentState {
	case constant.PaymentStateTimeout:
		info.State = orderstatepb.EState_PaymentTimeout
	case orderstatepb.EState_PaymentTimeout.String():
		info.State = orderstatepb.EState_PaymentTimeout

	case constant.PaymentStateWait:
		info.State = orderstatepb.EState_WaitPayment
	case orderstatepb.EState_WaitPayment.String():
		info.State = orderstatepb.EState_WaitPayment

	case constant.PaymentStateDone:
		info.State = orderstatepb.EState_Paid
	case orderstatepb.EState_Paid.String():
		info.State = orderstatepb.EState_Paid

	case constant.PaymentStateCanceled:
		info.State = orderstatepb.EState_Canceled
	case orderstatepb.EState_Canceled.String():
		info.State = orderstatepb.EState_Canceled

	default:
		info.State = orderstatepb.EState_WaitPayment
	}

	if info.State == orderstatepb.EState_WaitPayment && info.UserCanceled != 0 {
		info.State = orderstatepb.EState_UserCanceled
	}

	info.PaymentState = info.State.String()

	// TODO: Should from order state table
	now := uint32(time.Now().Unix())
	if orderstatepb.EState_Paid == info.State {
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
	damount := func(samount string) string {
		return decimal.RequireFromString(samount).
			Div(decimal.NewFromInt(int64(accuracy))).
			String()
	}

	info.PaymentAmount = damount(info.PaymentAmount)
	info.PaymentCoinUSDCurrency = damount(info.PaymentCoinUSDCurrency)
	info.PaymentLiveUSDCurrency = damount(info.PaymentLiveUSDCurrency)
	info.PaymentLocalUSDCurrency = damount(info.PaymentLocalUSDCurrency)

	return info
}

func Post(info *npool.Order) *npool.Order {
	return post(info)
}
