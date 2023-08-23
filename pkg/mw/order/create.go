package order

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"

	orderbasetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
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

func (h *createHandler) validate() error {
	if h.PaymentAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("startamount is less than or equal to 0")
	}
	if h.PayWithBalanceAmount != nil && h.PayWithBalanceAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("amount is less than or equal to 0")
	}
	if h.PaymentAccountStartAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("coinusdcurrency is less than or equal to 0")
	}
	if h.PaymentCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("localcoinusdcurrency is less than or equal to 0")
	}
	if h.PaymentLiveUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
	}
	if h.PaymentLocalUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
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

	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateUserTransfer, *h.AppID, *h.GoodID, *h.UserID)
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
		orderState := orderbasetypes.OrderState_OrderStateWaitPayment

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
				UserSpecialReductionID: h.UserSpecialReductionID,
				StartAt:                h.StartAt,
				EndAt:                  h.EndAt,
				Type:                   h.Type,
				State:                  &orderState,
				CouponIDs:              &h.CouponIDs,
				InvestmentType:         h.InvestmentType,
			},
		).Save(ctx); err != nil {
			return err
		}

		id := uuid.New()
		if h.PaymentID == nil {
			h.PaymentID = &id
		}

		paymentState := orderbasetypes.PaymentState_PaymentStateWait

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

func (h *Handler) CreateOrders(ctx context.Context) ([]*npool.Order, uint32, error) {
	for _, req := range h.Reqs {
		if req.IsMainGood != nil && *req.IsMainGood {
			id := uuid.New()
			h.ParentOrderID = req.ID
			if req.ID == nil {
				h.ParentOrderID = &id
				req.ID = &id
			}
		}
	}

	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.IsMainGood != nil && !*req.IsMainGood {
				req.ParentOrderID = h.ParentOrderID
			}
			orderState := orderbasetypes.OrderState_OrderStateWaitPayment

			if _, err := ordercrud.CreateSet(
				tx.Order.Create(),
				&ordercrud.Req{
					ID:                     req.ID,
					GoodID:                 req.GoodID,
					AppID:                  req.AppID,
					UserID:                 req.UserID,
					ParentOrderID:          req.ParentOrderID,
					PayWithParent:          req.PayWithParent,
					Units:                  req.Units,
					PromotionID:            req.PromotionID,
					UserSpecialReductionID: req.UserSpecialReductionID,
					StartAt:                req.StartAt,
					EndAt:                  req.EndAt,
					Type:                   req.Type,
					State:                  &orderState,
					CouponIDs:              req.CouponIDs,
					InvestmentType:         req.InvestmentType,
				},
			).Save(ctx); err != nil {
				return err
			}

			id := uuid.New()
			if h.PaymentID == nil {
				h.PaymentID = &id
			}

			paymentState := orderbasetypes.PaymentState_PaymentStateWait

			if _, err := paymentcrud.CreateSet(
				tx.Payment.Create(),
				&paymentcrud.Req{
					ID:                   req.PaymentID,
					AppID:                req.AppID,
					UserID:               req.UserID,
					GoodID:               req.GoodID,
					OrderID:              req.ID,
					AccountID:            req.PaymentAccountID,
					StartAmount:          req.PaymentAccountStartAmount,
					Amount:               req.PaymentAmount,
					PayWithBalanceAmount: req.PayWithBalanceAmount,
					CoinUsdCurrency:      req.PaymentCoinUSDCurrency,
					LocalCoinUsdCurrency: req.PaymentLocalUSDCurrency,
					LiveCoinUsdCurrency:  req.PaymentLiveUSDCurrency,
					CoinInfoID:           req.PaymentCoinID,
					State:                &paymentState,
				},
			).Save(ctx); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	h.Conds = &ordercrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	return h.GetOrders(ctx)
}
