package feeorder

import (
	"context"
	"fmt"

	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
	paymentmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	feeordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/fee"
	feeorderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/fee/state"
	ordercouponcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/coupon"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/lock"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	orderstatebasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/statebase"
	paymentbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	paymentbalancecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment/balance"
	paymenttransfercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment/transfer"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID       *uint32
	Rollback *bool
	feeordercrud.Req
	OrderBaseReq        *orderbasecrud.Req
	OrderStateBaseReq   *orderstatebasecrud.Req
	FeeOrderStateReq    *feeorderstatecrud.Req
	PaymentBaseReq      *paymentbasecrud.Req
	PaymentBalanceReqs  []*paymentbalancecrud.Req
	PaymentTransferReqs []*paymenttransfercrud.Req
	LedgerLockReq       *orderlockcrud.Req
	OrderCouponReqs     []*ordercouponcrud.Req
	FeeOrderConds       *feeordercrud.Conds
	OrderBaseConds      *orderbasecrud.Conds
	OrderStateBaseConds *orderstatebasecrud.Conds
	FeeOrderStateConds  *feeorderstatecrud.Conds
	Offset              int32
	Limit               int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		OrderBaseReq:      &orderbasecrud.Req{},
		OrderStateBaseReq: &orderstatebasecrud.Req{},
		FeeOrderStateReq:  &feeorderstatecrud.Req{},
		LedgerLockReq: &orderlockcrud.Req{
			LockType: func() *types.OrderLockType { e := types.OrderLockType_LockBalance; return &e }(),
		},
		PaymentBaseReq:      &paymentbasecrud.Req{},
		FeeOrderConds:       &feeordercrud.Conds{},
		OrderBaseConds:      &orderbasecrud.Conds{},
		OrderStateBaseConds: &orderstatebasecrud.Conds{},
		FeeOrderStateConds:  &feeorderstatecrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderBaseReq.AppID = &_id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderBaseReq.UserID = &_id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderBaseReq.GoodID = &_id
		return nil
	}
}

func WithGoodType(e *goodtypes.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case goodtypes.GoodType_TechniqueServiceFee:
		case goodtypes.GoodType_ElectricityFee:
		default:
			return fmt.Errorf("invalid goodtype")
		}
		h.OrderBaseReq.GoodType = e
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderBaseReq.AppGoodID = &_id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderID = &_id
		h.OrderBaseReq.EntID = &_id
		h.OrderStateBaseReq.OrderID = &_id
		h.FeeOrderStateReq.OrderID = &_id
		h.LedgerLockReq.OrderID = &_id
		h.PaymentBaseReq.OrderID = &_id
		return nil
	}
}

func WithParentOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid parentorderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderBaseReq.ParentOrderID = &_id
		return nil
	}
}

func WithOrderType(orderType *types.OrderType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderType == nil {
			if must {
				return fmt.Errorf("invalid ordertype")
			}
			return nil
		}
		switch *orderType {
		case types.OrderType_Airdrop:
		case types.OrderType_Offline:
		case types.OrderType_Normal:
		default:
			return fmt.Errorf("invalid ordertype")
		}
		h.OrderBaseReq.OrderType = orderType
		return nil
	}
}

func WithPaymentType(paymentType *types.PaymentType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if paymentType == nil {
			if must {
				return fmt.Errorf("invalid paymenttype")
			}
			return nil
		}
		switch *paymentType {
		case types.PaymentType_PayWithBalanceOnly:
		case types.PaymentType_PayWithTransferOnly:
		case types.PaymentType_PayWithTransferAndBalance:
		case types.PaymentType_PayWithParentOrder:
		case types.PaymentType_PayWithContract:
		case types.PaymentType_PayWithOffline:
		case types.PaymentType_PayWithNoPayment:
		default:
			return fmt.Errorf("invalid paymentType")
		}
		h.OrderBaseReq.PaymentType = paymentType
		return nil
	}
}

func WithGoodValueUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid goodvalueusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid goodvalueusd")
		}
		h.GoodValueUSD = &amount
		return nil
	}
}

func WithPaymentAmountUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentamountusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid paymentamountusd")
		}
		h.PaymentAmountUSD = &amount
		return nil
	}
}

func WithDiscountAmountUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid discountamountusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid discountamountusd")
		}
		h.DiscountAmountUSD = &amount
		return nil
	}
}

func WithPromotionID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid promotionid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.PromotionID = &_id
		return nil
	}
}

func WithDurationSeconds(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return fmt.Errorf("invalid durationseconds")
			}
			return nil
		}
		if *duration <= 0 {
			return fmt.Errorf("invalid durationseconds")
		}
		h.DurationSeconds = duration
		return nil
	}
}

func WithOrderState(state *types.OrderState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid orderstate")
			}
			return nil
		}
		switch *state {
		case types.OrderState_OrderStateCreated:
		case types.OrderState_OrderStateWaitPayment:
		case types.OrderState_OrderStatePaymentTransferReceived:
		case types.OrderState_OrderStatePaymentTransferBookKeeping:
		case types.OrderState_OrderStatePaymentSpendBalance:
		case types.OrderState_OrderStateTransferGoodStockLocked:
		case types.OrderState_OrderStateAddCommission:
		case types.OrderState_OrderStateAchievementBookKeeping:
		case types.OrderState_OrderStateUpdatePaidChilds:
		case types.OrderState_OrderStateChildPaidByParent:
		case types.OrderState_OrderStatePaymentUnlockAccount:
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateTransferGoodStockWaitStart:
		case types.OrderState_OrderStateUpdateInServiceChilds:
		case types.OrderState_OrderStateChildInServiceByParent:
		case types.OrderState_OrderStateInService:
		case types.OrderState_OrderStatePaymentTimeout:
		case types.OrderState_OrderStatePreCancel:
		case types.OrderState_OrderStatePreExpired:
		case types.OrderState_OrderStateRestoreExpiredStock:
		case types.OrderState_OrderStateUpdateExpiredChilds:
		case types.OrderState_OrderStateChildExpiredByParent:
		case types.OrderState_OrderStateRestoreCanceledStock:
		case types.OrderState_OrderStateCancelAchievement:
		case types.OrderState_OrderStateDeductLockedCommission:
		case types.OrderState_OrderStateReturnCanceledBalance:
		case types.OrderState_OrderStateUpdateCanceledChilds:
		case types.OrderState_OrderStateChildCanceledByParent:
		case types.OrderState_OrderStateCanceledTransferBookKeeping:
		case types.OrderState_OrderStateCancelUnlockPaymentAccount:
		case types.OrderState_OrderStateCanceled:
		case types.OrderState_OrderStateExpired:
		default:
			return fmt.Errorf("invalid orderstate")
		}
		h.OrderStateBaseReq.OrderState = state
		return nil
	}
}

func WithCreateMethod(e *types.OrderCreateMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid createmethod")
			}
			return nil
		}
		switch *e {
		case types.OrderCreateMethod_OrderCreatedByPurchase:
		case types.OrderCreateMethod_OrderCreatedByAdmin:
		case types.OrderCreateMethod_OrderCreatedByRenew:
		default:
			return fmt.Errorf("invalid createmethod")
		}
		h.OrderBaseReq.CreateMethod = e
		return nil
	}
}

func WithUserSetPaid(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid usersetpaid")
			}
			return nil
		}
		h.FeeOrderStateReq.UserSetPaid = value
		return nil
	}
}

func WithUserSetCanceled(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid usersetcanceled")
			}
			return nil
		}
		h.FeeOrderStateReq.UserSetCanceled = value
		return nil
	}
}

func WithAdminSetCanceled(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid adminsetcanceled")
			}
			return nil
		}
		h.FeeOrderStateReq.AdminSetCanceled = value
		return nil
	}
}

func WithPaymentState(state *types.PaymentState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid paymentstate")
			}
			return nil
		}
		switch *state {
		case types.PaymentState_PaymentStateWait:
		case types.PaymentState_PaymentStateCanceled:
		case types.PaymentState_PaymentStateTimeout:
		case types.PaymentState_PaymentStateDone:
		case types.PaymentState_PaymentStateNoPayment:
		default:
			return fmt.Errorf("invalid paymentstate")
		}
		h.FeeOrderStateReq.PaymentState = state
		return nil
	}
}

func WithRollback(rollback *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if rollback == nil {
			if must {
				return fmt.Errorf("invalid rollback")
			}
			return nil
		}
		h.Rollback = rollback
		return nil
	}
}

func WithLedgerLockID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid ledgerlockid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.LedgerLockReq.EntID = &_id
		return nil
	}
}

func WithPaymentID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid paymentid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.PaymentBaseReq.EntID = &_id
		return nil
	}
}

func WithCouponIDs(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, s := range ss {
			id, err := uuid.Parse(s)
			if err != nil {
				return err
			}
			// Fill order id later
			h.OrderCouponReqs = append(h.OrderCouponReqs, &ordercouponcrud.Req{
				CouponID: &id,
			})
		}
		return nil
	}
}

func WithPaymentBalances(bs []*paymentmwpb.PaymentBalanceReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, b := range bs {
			req := &paymentbalancecrud.Req{}

			id, err := uuid.Parse(b.GetCoinTypeID())
			if err != nil {
				return err
			}
			req.CoinTypeID = &id

			if b.LocalCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(b.GetLocalCoinUSDCurrency())
				if err != nil {
					return err
				}
				req.LocalCoinUSDCurrency = &amount
			}

			amount1, err := decimal.NewFromString(b.GetAmount())
			if err != nil {
				return err
			}
			req.Amount = &amount1

			amount2, err := decimal.NewFromString(b.GetLiveCoinUSDCurrency())
			if err != nil {
				return err
			}
			req.LiveCoinUSDCurrency = &amount2

			h.PaymentBalanceReqs = append(h.PaymentBalanceReqs, req)
		}
		return nil
	}
}

func WithPaymentTransfers(bs []*paymentmwpb.PaymentTransferReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, b := range bs {
			req := &paymenttransfercrud.Req{}

			id1, err := uuid.Parse(b.GetCoinTypeID())
			if err != nil {
				return err
			}
			req.CoinTypeID = &id1

			if b.LocalCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(*b.LocalCoinUSDCurrency)
				if err != nil {
					return err
				}
				req.LocalCoinUSDCurrency = &amount
			}

			id2, err := uuid.Parse(b.GetAccountID())
			if err != nil {
				return err
			}
			req.AccountID = &id2

			amount1, err := decimal.NewFromString(b.GetAmount())
			if err != nil {
				return err
			}
			req.Amount = &amount1

			amount2, err := decimal.NewFromString(b.GetStartAmount())
			if err != nil {
				return err
			}
			req.StartAmount = &amount2

			amount3, err := decimal.NewFromString(*b.LiveCoinUSDCurrency)
			if err != nil {
				return err
			}
			req.LiveCoinUSDCurrency = &amount3

			amount4, err := decimal.NewFromString(b.GetFinishAmount())
			if err != nil {
				return err
			}
			req.FinishAmount = &amount4

			h.PaymentTransferReqs = append(h.PaymentTransferReqs, req)
		}
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}