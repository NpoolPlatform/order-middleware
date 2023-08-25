package order

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"

	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate(req *ordercrud.Req) error {
	if req.PaymentAmount == nil {
		return fmt.Errorf("invalid paymentamount")
	}
	if req.PaymentAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("startamount is less than or equal to 0")
	}
	if req.PaymentBalanceAmount != nil && req.PaymentBalanceAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("amount is less than or equal to 0")
	}
	if req.PaymentStartAmount == nil {
		return fmt.Errorf("invalid paymentstartamount")
	}
	if req.PaymentStartAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("coinusdcurrency is less than or equal to 0")
	}
	if req.PaymentCoinUSDCurrency == nil {
		return fmt.Errorf("invalid paymentcoinusdcurrency")
	}
	if req.PaymentCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("localcoinusdcurrency is less than or equal to 0")
	}
	if req.PaymentLiveCoinUSDCurrency == nil {
		return fmt.Errorf("invalid paymentlivecoinusdcurrency")
	}
	if req.PaymentLiveCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
	}
	if req.PaymentLocalCoinUSDCurrency == nil {
		return fmt.Errorf("invalid paymentlocalcoinusdcurrency")
	}
	if req.PaymentLocalCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
	}

	return nil
}

func (h *createHandler) createOrder(ctx context.Context, tx *ent.Tx, req *ordercrud.Req) error {
	orderState := ordertypes.OrderState_OrderStateWaitPayment
	paymentState := ordertypes.PaymentState_PaymentStateNoPayment
	if req.PaymentTransferAmount != nil && req.PaymentTransferAmount.Cmp(decimal.NewFromInt(0)) > 0 {
		if err := h.validate(req); err != nil {
			return err
		}
		id := uuid.New()
		req.PaymentID = &id
		paymentState = ordertypes.PaymentState_PaymentStateWait
		if _, err := paymentcrud.CreateSet(
			tx.Payment.Create(),
			&paymentcrud.Req{
				ID:                   req.PaymentID,
				AppID:                req.AppID,
				UserID:               req.UserID,
				GoodID:               req.GoodID,
				OrderID:              req.ID,
				AccountID:            req.PaymentAccountID,
				CoinTypeID:           req.PaymentCoinTypeID,
				StartAmount:          req.PaymentStartAmount,
				TransferAmount:       req.PaymentTransferAmount,
				BalanceAmount:        req.PaymentBalanceAmount,
				CoinUSDCurrency:      req.PaymentCoinUSDCurrency,
				LocalCoinUSDCurrency: req.PaymentLocalCoinUSDCurrency,
				LiveCoinUSDCurrency:  req.PaymentLiveCoinUSDCurrency,
			},
		).Save(ctx); err != nil {
			return err
		}
	}
	if _, err := ordercrud.CreateSet(
		tx.Order.Create(),
		&ordercrud.Req{
			ID:             req.ID,
			AppID:          req.AppID,
			UserID:         req.UserID,
			GoodID:         req.GoodID,
			AppGoodID:      req.AppGoodID,
			PaymentID:      req.PaymentID,
			ParentOrderID:  req.ParentOrderID,
			Units:          req.Units,
			GoodValue:      req.GoodValue,
			PaymentAmount:  req.PaymentAmount,
			DiscountAmount: req.DiscountAmount,
			PromotionID:    req.PromotionID,
			DurationDays:   req.DurationDays,
			OrderType:      req.OrderType,
			InvestmentType: req.InvestmentType,
			CouponIDs:      req.CouponIDs,
			PaymentType:    req.PaymentType,
		},
	).Save(ctx); err != nil {
		return err
	}

	id := uuid.New()
	if _, err := orderstatecrud.CreateSet(
		tx.OrderState.Create(),
		&orderstatecrud.Req{
			ID:                   &id,
			OrderID:              req.ID,
			OrderState:           &orderState,
			StartMode:            req.StartMode,
			StartAt:              req.StartAt,
			EndAt:                req.EndAt,
			LastBenefitAt:        req.LastBenefitAt,
			BenefitState:         req.BenefitState,
			UserSetPaid:          req.UserSetPaid,
			UserSetCanceled:      req.UserSetCanceled,
			PaymentTransactionID: req.PaymentTransactionID,
			PaymentFinishAmount:  req.PaymentFinishAmount,
			PaymentState:         &paymentState,
			OutOfGasHours:        req.OutOfGasHours,
			CompensateHours:      req.CompensateHours,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateOrder(ctx context.Context) (*npool.Order, error) {
	handler := &createHandler{
		Handler: h,
	}

	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateUserTransfer, *h.AppID, *h.GoodID, *h.UserID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	req := &ordercrud.Req{
		ID:                          h.ID,
		AppID:                       h.AppID,
		UserID:                      h.UserID,
		GoodID:                      h.GoodID,
		AppGoodID:                   h.AppGoodID,
		ParentOrderID:               h.ParentOrderID,
		Units:                       h.Units,
		GoodValue:                   h.GoodValue,
		PaymentAmount:               h.PaymentAmount,
		DiscountAmount:              h.DiscountAmount,
		PromotionID:                 h.PromotionID,
		DurationDays:                h.DurationDays,
		OrderType:                   h.OrderType,
		InvestmentType:              h.InvestmentType,
		CouponIDs:                   &h.CouponIDs,
		PaymentType:                 h.PaymentType,
		PaymentAccountID:            h.PaymentAccountID,
		PaymentCoinTypeID:           h.PaymentCoinTypeID,
		PaymentStartAmount:          h.PaymentStartAmount,
		PaymentTransferAmount:       h.PaymentTransferAmount,
		PaymentBalanceAmount:        h.PaymentBalanceAmount,
		PaymentCoinUSDCurrency:      h.PaymentCoinUSDCurrency,
		PaymentLocalCoinUSDCurrency: h.PaymentLocalCoinUSDCurrency,
		PaymentLiveCoinUSDCurrency:  h.PaymentLiveCoinUSDCurrency,
		StartMode:                   h.StartMode,
		StartAt:                     h.StartAt,
		EndAt:                       h.EndAt,
		LastBenefitAt:               h.LastBenefitAt,
		BenefitState:                h.BenefitState,
		UserSetPaid:                 h.UserSetPaid,
		UserSetCanceled:             h.UserSetCanceled,
		PaymentTransactionID:        h.PaymentTransactionID,
		PaymentFinishAmount:         h.PaymentFinishAmount,
		OutOfGasHours:               h.OutOfGasHours,
		CompensateHours:             h.CompensateHours,
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		id := uuid.New()
		if req.ID == nil {
			req.ID = &id
		}
		if err := handler.createOrder(ctx, tx, req); err != nil {
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
	handler := &createHandler{
		Handler: h,
	}
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID == nil {
				req.ID = &id
			}
			if err := handler.createOrder(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *req.ID)
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
