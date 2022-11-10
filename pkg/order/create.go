package order

import (
	"context"

	"go.opentelemetry.io/otel"

	scodes "go.opentelemetry.io/otel/codes"

	constant "github.com/NpoolPlatform/order-middleware/pkg/message/const"

	"github.com/NpoolPlatform/order-manager/pkg/db"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	orderpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"
	ordercrud "github.com/NpoolPlatform/order-manager/pkg/crud/order"

	paymentpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"
	paymentcrud "github.com/NpoolPlatform/order-manager/pkg/crud/payment"
)

func CreateOrder(ctx context.Context, in *npool.OrderReq) (info *npool.Order, err error) {
	var id string

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGood")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		orderState := orderpb.OrderState_WaitPayment

		c := tx.Order.Create()
		stm, err := ordercrud.CreateSet(c, &orderpb.OrderReq{
			ID:                     in.ID,
			GoodID:                 in.GoodID,
			AppID:                  in.AppID,
			UserID:                 in.UserID,
			ParentOrderID:          in.ParentOrderID,
			PayWithParent:          in.PayWithParent,
			Units:                  in.Units,
			PromotionID:            in.PromotionID,
			DiscountCouponID:       in.DiscountID,
			UserSpecialReductionID: in.SpecialOfferID,
			StartAt:                in.Start,
			EndAt:                  in.End,
			FixAmountCouponID:      in.FixAmountID,
			Type:                   in.OrderType,
			State:                  &orderState,
		})
		if err != nil {
			return err
		}

		ord, err := stm.Save(_ctx)
		if err != nil {
			return err
		}

		id = ord.ID.String()
		paymentState := paymentpb.PaymentState_Wait

		c1 := tx.Payment.Create()
		stm1, err := paymentcrud.CreateSet(c1, &paymentpb.PaymentReq{
			ID:                   in.PaymentID,
			AppID:                in.AppID,
			UserID:               in.UserID,
			GoodID:               in.GoodID,
			OrderID:              &id,
			AccountID:            in.PaymentAccountID,
			StartAmount:          in.PaymentAccountStartAmount,
			Amount:               in.PaymentAmount,
			PayWithBalanceAmount: in.PayWithBalanceAmount,
			CoinUsdCurrency:      in.PaymentCoinUSDCurrency,
			LocalCoinUsdCurrency: in.PaymentLocalUSDCurrency,
			LiveCoinUsdCurrency:  in.PaymentLiveUSDCurrency,
			CoinInfoID:           in.PaymentCoinID,
			State:                &paymentState,
		})
		if err != nil {
			return err
		}

		_, err = stm1.Save(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetOrder(ctx, id)
}
