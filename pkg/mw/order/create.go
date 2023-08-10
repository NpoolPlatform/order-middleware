package order

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *createHandler) validate() error {
	if h.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if h.UserID == nil {
		return fmt.Errorf("invalid userid")
	}
	if h.GoodID == nil {
		return fmt.Errorf("invalid goodid")
	}
	if h.PaymentCoinID == nil {
		return fmt.Errorf("invalid paymentcoinid")
	}
	if h.PaymentAccountID == nil {
		return fmt.Errorf("invalid paymentaccountid")
	}
	if h.PaymentAmount == nil {
		return fmt.Errorf("invalid startamount")
	}
	if h.PaymentAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("startamount is less than or equal to 0")
	}
	if h.PayWithBalanceAmount != nil && h.PayWithBalanceAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("amount is less than or equal to 0")
	}

	if h.PaymentAccountStartAmount == nil {
		return fmt.Errorf("invalid coinusdcurrency")
	}
	if h.PaymentAccountStartAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("coinusdcurrency is less than or equal to 0")
	}
	if h.PaymentCoinUSDCurrency == nil {
		return fmt.Errorf("invalid localcoinusdcurrency")
	}
	if h.PaymentCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("localcoinusdcurrency is less than or equal to 0")
	}
	if h.PaymentLiveUSDCurrency == nil {
		return fmt.Errorf("invalid livecoinusdcurrency")
	}
	if h.PaymentLiveUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
	}
	if h.PaymentLocalUSDCurrency == nil {
		return fmt.Errorf("invalid livecoinusdcurrency")
	}
	if h.PaymentLocalUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
	}
	if h.Type == nil {
		return fmt.Errorf("invalid ordertype")
	}
	switch *h.Type {
	case basetypes.OrderType_OrderTypeNormal:
	case basetypes.OrderType_OrderTypeOffline:
	case basetypes.OrderType_OrderTypeAirdrop:
	default:
		return fmt.Errorf("invalid ordertype")
	}
	return nil
}

func (h *Handler) CreateOrder(ctx context.Context) (*npool.Order, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateUserTransfer, *h.ParentOrderID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		orderState := basetypes.OrderState_OrderStateWaitPayment

		if _, err := ordercrud.CreateSet(
			tx.Order.Create(),
			&ordercrud.Req{
				ID:                     h.ID,
				GoodID:                 h.GoodID,
				AppID:                  h.AppID,
				UserID:                 h.UserID,
				ParentOrderID:          h.ParentOrderID,
				PayWithParent:          h.PayWithParent,
				Units:                  h.Units,
				PromotionID:            h.PromotionID,
				DiscountCouponID:       h.DiscountCouponID,
				UserSpecialReductionID: h.UserSpecialReductionID,
				StartAt:                h.StartAt,
				EndAt:                  h.EndAt,
				FixAmountCouponID:      h.FixAmountCouponID,
				Type:                   basetypes.OrderType_DefaultOrderType.Enum(),
				State:                  &orderState,
				CouponIDs:              &h.CouponIDs,
			},
		).Save(ctx); err != nil {
			return err
		}

		paymentState := basetypes.PaymentState_PaymentStateWait

		if _, err := paymentcrud.CreateSet(
			tx.Payment.Create(),
			&paymentcrud.Req{
				ID:                   h.PaymentID,
				AppID:                h.AppID,
				UserID:               h.UserID,
				GoodID:               h.GoodID,
				OrderID:              h.ID,
				AccountID:            h.PaymentAccountID,
				StartAmount:          h.PaymentAccountStartAmount,
				Amount:               h.PaymentAmount,
				PayWithBalanceAmount: h.PayWithBalanceAmount,
				CoinUsdCurrency:      h.PaymentCoinUSDCurrency,
				LocalCoinUsdCurrency: h.PaymentLocalUSDCurrency,
				LiveCoinUsdCurrency:  h.PaymentLiveUSDCurrency,
				CoinInfoID:           h.PaymentCoinID,
				State:                &paymentState,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOrder(ctx)
}