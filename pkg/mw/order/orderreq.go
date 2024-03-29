package order

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderlock"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderReq struct {
	*ordercrud.Req
	OrderStateReq  *orderstatecrud.Req
	PaymentReq     *paymentcrud.Req
	StockLockReq   *orderlockcrud.Req
	BalanceLockReq *orderlockcrud.Req
}

//nolint:funlen,gocyclo
func (h *Handler) ToOrderReq(ctx context.Context, newOrder bool) (*OrderReq, error) {
	if !newOrder {
		info, err := h.GetOrder(ctx)
		if err != nil {
			return nil, err
		}
		if info == nil {
			return nil, fmt.Errorf("invalid order")
		}
		if h.EntID == nil {
			id := uuid.MustParse(info.EntID)
			h.EntID = &id
		}
		if h.StartAt == nil {
			h.StartAt = &info.StartAt
		}
		if h.EndAt == nil {
			h.EndAt = &info.EndAt
		}
	}

	if h.StartAt == nil || h.EndAt == nil {
		return nil, fmt.Errorf("invalid duration")
	}

	if *h.EndAt <= *h.StartAt {
		return nil, fmt.Errorf("invalid order")
	}

	req := &OrderReq{
		Req: &ordercrud.Req{
			ID:                   h.ID,
			EntID:                h.EntID,
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
			Duration:             h.Duration,
			OrderType:            h.OrderType,
			InvestmentType:       h.InvestmentType,
			CouponIDs:            h.CouponIDs,
			PaymentType:          h.PaymentType,
			CoinTypeID:           h.CoinTypeID,
			PaymentCoinTypeID:    h.PaymentCoinTypeID,
			TransferAmount:       h.TransferAmount,
			BalanceAmount:        h.BalanceAmount,
			CoinUSDCurrency:      h.CoinUSDCurrency,
			LocalCoinUSDCurrency: h.LocalCoinUSDCurrency,
			LiveCoinUSDCurrency:  h.LiveCoinUSDCurrency,
			Simulate:             h.Simulate,
			CreateMethod:         h.CreateMethod,
			MultiPaymentCoins:    h.MultiPaymentCoins,
			PaymentAmounts:       h.PaymentAmounts,
		},
		OrderStateReq: &orderstatecrud.Req{
			OrderID:              h.EntID,
			OrderState:           h.OrderState,
			StartMode:            h.StartMode,
			PaymentState:         h.PaymentState,
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
			RenewState:           h.RenewState,
			RenewNotifyAt:        h.RenewNotifyAt,
		},
	}

	if h.AppGoodStockLockID != nil {
		req.StockLockReq = &orderlockcrud.Req{
			EntID:    h.AppGoodStockLockID,
			AppID:    h.AppID,
			UserID:   h.UserID,
			OrderID:  h.EntID,
			LockType: basetypes.OrderLockType_LockStock.Enum(),
		}
	}

	if (h.BalanceAmount != nil && h.BalanceAmount.Cmp(decimal.NewFromInt(0)) > 0) ||
		len(h.PaymentAmounts) > 0 { // In this case one ledger lock will relevant to multiple statements
		req.BalanceLockReq = &orderlockcrud.Req{
			EntID:    h.LedgerLockID,
			AppID:    h.AppID,
			UserID:   h.UserID,
			OrderID:  h.EntID,
			LockType: basetypes.OrderLockType_LockBalance.Enum(),
		}
	}

	if newOrder && req.PaymentType != nil {
		if err := req.CheckOrderType(); err != nil {
			return nil, err
		}
		has, err := req.HasPayment()
		if err != nil {
			return nil, err
		}
		if !has {
			return req, nil
		}
		if req.PaymentCoinTypeID == nil {
			return nil, fmt.Errorf("invalid paymentcointypeid")
		}
	}

	paymentID := uuid.New()
	req.Req.PaymentID = &paymentID
	req.PaymentReq = &paymentcrud.Req{
		EntID:       &paymentID,
		OrderID:     h.EntID,
		AppID:       h.AppID,
		UserID:      h.UserID,
		GoodID:      h.GoodID,
		AccountID:   h.PaymentAccountID,
		StartAmount: h.PaymentStartAmount,
	}
	return req, nil
}

//nolint:gocyclo
func (r *OrderReq) CheckOrderType() error {
	switch *r.OrderType {
	case basetypes.OrderType_Normal:
		switch *r.PaymentType {
		case basetypes.PaymentType_PayWithTransferOnly:
			fallthrough //nolint
		case basetypes.PaymentType_PayWithTransferAndBalance:
			fallthrough //nolint
		case basetypes.PaymentType_PayWithBalanceOnly:
			fallthrough //nolint
		case basetypes.PaymentType_PayWithParentOrder:
			if r.ParentOrderID != nil {
				return nil
			}
			if r.Simulate != nil && *r.Simulate {
				return fmt.Errorf("invalid paymenttype")
			}
		case basetypes.PaymentType_PayWithNoPayment:
			if r.Simulate == nil || !*r.Simulate {
				return fmt.Errorf("invalid paymenttype")
			}
		default:
			return fmt.Errorf("invalid paymenttype")
		}
	case basetypes.OrderType_Offline:
		if *r.PaymentType != basetypes.PaymentType_PayWithOffline {
			return fmt.Errorf("invalid paymenttype")
		}
	case basetypes.OrderType_Airdrop:
		if *r.PaymentType != basetypes.PaymentType_PayWithNoPayment {
			return fmt.Errorf("invalid paymenttype")
		}
	default:
		return fmt.Errorf("invalid ordertype")
	}
	return nil
}

func (r *OrderReq) HasPayment() (bool, error) {
	zeroAmount := decimal.NewFromInt(0)
	if r.PaymentAmount == nil {
		r.PaymentAmount = &zeroAmount
	}
	if r.BalanceAmount == nil {
		r.BalanceAmount = &zeroAmount
	}
	if r.TransferAmount == nil {
		r.TransferAmount = &zeroAmount
	}
	switch *r.PaymentType {
	case basetypes.PaymentType_PayWithBalanceOnly:
		r.TransferAmount = &zeroAmount
		if r.TransferAmount.Add(*r.BalanceAmount).Cmp(*r.PaymentAmount) != 0 {
			return false, fmt.Errorf("invalid paymentAmount")
		}
		return false, nil
	case basetypes.PaymentType_PayWithTransferOnly:
		r.BalanceAmount = &zeroAmount
		fallthrough
	case basetypes.PaymentType_PayWithTransferAndBalance:
		if r.TransferAmount.Add(*r.BalanceAmount).Cmp(*r.PaymentAmount) != 0 {
			return false, fmt.Errorf("invalid paymentAmount")
		}
	case basetypes.PaymentType_PayWithOffline:
		r.BalanceAmount = &zeroAmount
		r.TransferAmount = &zeroAmount
		return false, nil
	case basetypes.PaymentType_PayWithParentOrder:
		fallthrough //nolint
	case basetypes.PaymentType_PayWithNoPayment:
		r.PaymentAmount = &zeroAmount
		r.BalanceAmount = &zeroAmount
		r.TransferAmount = &zeroAmount
		return false, nil
	default:
		return false, fmt.Errorf("invalid paymenttype")
	}
	return true, nil
}
