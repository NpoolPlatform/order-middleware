package order

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
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

func (h *Handler) ToOrderReq() (*OrderReq, error) {
	req := &OrderReq{
		Req: &ordercrud.Req{
			ID:                   h.ID,
			AppID:                h.AppID,
			UserID:               h.UserID,
			GoodID:               h.GoodID,
			AppGoodID:            h.AppGoodID,
			ParentOrderID:        h.ParentOrderID,
			Units:                h.Units,
			GoodValue:            h.GoodValue,
			GoodValueUSD:         h.GoodValueUSD,
			PaymentAmount:        h.PaymentAmount,
			DiscountAmount:       h.DiscountAmount,
			PromotionID:          h.PromotionID,
			DurationDays:         h.DurationDays,
			OrderType:            h.OrderType,
			InvestmentType:       h.InvestmentType,
			CouponIDs:            &h.CouponIDs,
			PaymentType:          h.PaymentType,
			CoinTypeID:           h.CoinTypeID,
			PaymentCoinTypeID:    h.PaymentCoinTypeID,
			TransferAmount:       h.TransferAmount,
			BalanceAmount:        h.BalanceAmount,
			CoinUSDCurrency:      h.CoinUSDCurrency,
			LocalCoinUSDCurrency: h.LocalCoinUSDCurrency,
			LiveCoinUSDCurrency:  h.LiveCoinUSDCurrency,
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
			AdminSetCanceled:     h.AdminSetCanceled,
			PaymentTransactionID: h.PaymentTransactionID,
			PaymentFinishAmount:  h.PaymentFinishAmount,
			OutOfGasHours:        h.OutOfGasHours,
			CompensateHours:      h.CompensateHours,
		},
	}

	if req.PaymentType != nil {
		has, err := req.HasPayment()
		if err != nil {
			return req, err
		}
		if !has {
			return req, err
		}
	}

	paymentID := uuid.New()
	req.Req.PaymentID = &paymentID
	req.PaymentReq = &paymentcrud.Req{
		ID:          &paymentID,
		OrderID:     h.ID,
		AppID:       h.AppID,
		UserID:      h.UserID,
		GoodID:      h.GoodID,
		AccountID:   h.PaymentAccountID,
		StartAmount: h.PaymentStartAmount,
	}
	return req, nil
}

func (r *OrderReq) HasPayment() (bool, error) {
	switch *r.PaymentType {
	case basetypes.PaymentType_PayWithBalanceOnly:
		fallthrough //nolint
	case basetypes.PaymentType_PayWithTransferOnly:
		fallthrough //nolint
	case basetypes.PaymentType_PayWithTransferAndBalance:
		if r.PaymentAmount == nil {
			return false, fmt.Errorf("invalid paymentamount")
		}
		if r.PaymentAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return false, nil
		}

	case basetypes.PaymentType_PayWithParentOrder:
		fallthrough //nolint
	case basetypes.PaymentType_PayWithOffline:
		fallthrough //nolint
	case basetypes.PaymentType_PayWithNoPayment:
		return false, nil
	default:
		return false, fmt.Errorf("invalid paymenttype")
	}
	return true, nil
}
