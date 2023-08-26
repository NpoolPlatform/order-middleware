package order

import (
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderReq struct {
	*ordercrud.Req
	OrderStateReq *orderstatecrud.Req
	PaymentReq    *paymentcrud.Req
}

func (h *Handler) ToOrderReq() *OrderReq {
	paymentID := uuid.New()

	req := &OrderReq{
		Req: &ordercrud.Req{
			ID:             h.ID,
			AppID:          h.AppID,
			UserID:         h.UserID,
			GoodID:         h.GoodID,
			AppGoodID:      h.AppGoodID,
			ParentOrderID:  h.ParentOrderID,
			Units:          h.Units,
			GoodValue:      h.GoodValue,
			PaymentAmount:  h.PaymentAmount,
			DiscountAmount: h.DiscountAmount,
			PromotionID:    h.PromotionID,
			DurationDays:   h.DurationDays,
			OrderType:      h.OrderType,
			InvestmentType: h.InvestmentType,
			CouponIDs:      &h.CouponIDs,
			PaymentType:    h.PaymentType,
		},
		OrderStateReq: &orderstatecrud.Req{
			OrderID:              h.ID,
			OrderState:           h.OrderState,
			StartMode:            h.StartMode,
			StartAt:              h.StartAt,
			EndAt:                h.EndAt,
			LastBenefitAt:        h.LastBenefitAt,
			BenefitState:         h.BenefitState,
			UserSetPaid:          h.UserSetPaid,
			UserSetCanceled:      h.UserSetCanceled,
			PaymentTransactionID: h.PaymentTransactionID,
			PaymentFinishAmount:  h.PaymentFinishAmount,
			OutOfGasHours:        h.OutOfGasHours,
			CompensateHours:      h.CompensateHours,
		},
	}

	if h.PaymentTransferAmount != nil && h.PaymentTransferAmount.Cmp(decimal.NewFromInt(0)) > 0 {
		req.Req.PaymentID = &paymentID
		req.PaymentReq = &paymentcrud.Req{
			ID:                   &paymentID,
			OrderID:              h.ID,
			AppID:                h.AppID,
			UserID:               h.UserID,
			GoodID:               h.GoodID,
			AccountID:            h.PaymentAccountID,
			CoinTypeID:           h.PaymentCoinTypeID,
			StartAmount:          h.PaymentStartAmount,
			TransferAmount:       h.PaymentTransferAmount,
			BalanceAmount:        h.PaymentBalanceAmount,
			CoinUSDCurrency:      h.PaymentCoinUSDCurrency,
			LocalCoinUSDCurrency: h.PaymentLocalCoinUSDCurrency,
			LiveCoinUSDCurrency:  h.PaymentLiveCoinUSDCurrency,
		}
	}
	return req
}

func (h *Handler) ToOrderReqs() []*OrderReq {
	reqs := []*OrderReq{}
	for _, req := range h.Reqs {
		paymentID := uuid.New()

		_req := &OrderReq{
			Req: &ordercrud.Req{
				ID:             req.ID,
				AppID:          req.AppID,
				UserID:         req.UserID,
				GoodID:         req.GoodID,
				AppGoodID:      req.AppGoodID,
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
			OrderStateReq: &orderstatecrud.Req{
				OrderID:              req.ID,
				OrderState:           req.OrderStateReq.OrderState,
				StartMode:            req.OrderStateReq.StartMode,
				StartAt:              req.OrderStateReq.StartAt,
				EndAt:                req.OrderStateReq.EndAt,
				LastBenefitAt:        req.OrderStateReq.LastBenefitAt,
				BenefitState:         req.OrderStateReq.BenefitState,
				UserSetPaid:          req.OrderStateReq.UserSetPaid,
				UserSetCanceled:      req.OrderStateReq.UserSetCanceled,
				PaymentTransactionID: req.OrderStateReq.PaymentTransactionID,
				PaymentFinishAmount:  req.OrderStateReq.PaymentFinishAmount,
				OutOfGasHours:        req.OrderStateReq.OutOfGasHours,
				CompensateHours:      req.OrderStateReq.CompensateHours,
			},
		}

		if h.PaymentTransferAmount != nil && h.PaymentTransferAmount.Cmp(decimal.NewFromInt(0)) > 0 {
			_req.Req.PaymentID = &paymentID
			_req.PaymentReq = &paymentcrud.Req{
				ID:                   &paymentID,
				OrderID:              req.ID,
				AppID:                req.PaymentReq.AppID,
				UserID:               req.PaymentReq.UserID,
				GoodID:               req.PaymentReq.GoodID,
				AccountID:            req.PaymentReq.AccountID,
				CoinTypeID:           req.PaymentReq.CoinTypeID,
				StartAmount:          req.PaymentReq.StartAmount,
				TransferAmount:       req.PaymentReq.TransferAmount,
				BalanceAmount:        req.PaymentReq.BalanceAmount,
				CoinUSDCurrency:      req.PaymentReq.CoinUSDCurrency,
				LocalCoinUSDCurrency: req.PaymentReq.LocalCoinUSDCurrency,
				LiveCoinUSDCurrency:  req.PaymentReq.LiveCoinUSDCurrency,
			}
		}
		reqs = append(reqs, _req)
	}
	return reqs
}
