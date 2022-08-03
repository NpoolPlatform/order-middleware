package order

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"

	constant "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"

	orderstatepb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order/state"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	"github.com/google/uuid"
)

func CreateOrder(ctx context.Context, in *npool.OrderReq) (info *npool.Order, err error) {
	var id string

	// TODO: move to manager when refactor
	// TODO: refactor amount to string
	const accuracy = 1000000000000
	amount := func(samount string) uint64 {
		damount := decimal.RequireFromString(samount)
		return uint64(damount.Mul(decimal.NewFromInt(accuracy)).IntPart())
	}
	state := func(st orderstatepb.EState) payment.State {
		switch st {
		case orderstatepb.EState_WaitPayment:
			return payment.State(constant.PaymentStateWait)
		case orderstatepb.EState_Paid:
			return payment.State(constant.PaymentStateDone)
		case orderstatepb.EState_PaymentTimeout:
			return payment.State(constant.PaymentStateTimeout)
		case orderstatepb.EState_Canceled:
			return payment.State(constant.PaymentStateCanceled)
		default:
			return payment.State(constant.PaymentStateWait)
		}
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.
			Order.
			Create().
			SetAppID(uuid.MustParse(in.GetAppID())).
			SetUserID(uuid.MustParse(in.GetUserID())).
			SetGoodID(uuid.MustParse(in.GetGoodID())).
			SetUnits(in.GetUnits()).
			SetStart(in.GetStart()).
			SetEnd(in.GetEnd()).
			SetOrderType(in.GetOrderType().String())

		if in.FixAmountID != nil {
			stm = stm.SetCouponID(uuid.MustParse(in.GetFixAmountID()))
		}
		if in.DiscountID != nil {
			stm = stm.SetDiscountCouponID(uuid.MustParse(in.GetDiscountID()))
		}
		if in.SpecialOfferID != nil {
			stm = stm.SetUserSpecialReductionID(uuid.MustParse(in.GetSpecialOfferID()))
		}
		if in.PromotionID != nil {
			stm = stm.SetPromotionID(uuid.MustParse(in.GetPromotionID()))
		}

		ord, err := stm.Save(ctx)
		if err != nil {
			return err
		}

		id = ord.ID.String()

		stm1 := tx.
			Payment.
			Create().
			SetAppID(uuid.MustParse(in.GetAppID())).
			SetUserID(uuid.MustParse(in.GetUserID())).
			SetGoodID(uuid.MustParse(in.GetGoodID())).
			SetOrderID(ord.ID).
			SetCoinInfoID(uuid.MustParse(in.GetPaymentCoinID())).
			SetAccountID(uuid.MustParse(in.GetPaymentAccountID())).
			SetStartAmount(amount(in.GetPaymentAccountStartAmount())).
			SetFinishAmount(0).
			SetAmount(amount(in.GetPaymentAmount())).
			SetCoinUsdCurrency(amount(in.GetPaymentCoinUSDCurrency())).
			SetLocalCoinUsdCurrency(amount(in.GetPaymentLocalUSDCurrency())).
			SetLiveCoinUsdCurrency(amount(in.GetPaymentLiveUSDCurrency())).
			SetState(state(orderstatepb.EState_WaitPayment))
		if in.PayWithBalanceAmount != nil {
			stm1 = stm1.SetPayWithBalanceAmount(decimal.RequireFromString(in.GetPayWithBalanceAmount()))
		}
		_, err = stm1.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return GetOrder(ctx, id)
}
