//nolint:dupl
package order

import (
	"context"
	"fmt"
	"time"

	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderlock"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID                   *uint32
	EntID                *uuid.UUID
	AppID                *uuid.UUID
	UserID               *uuid.UUID
	GoodID               *uuid.UUID
	AppGoodID            *uuid.UUID
	ParentOrderID        *uuid.UUID
	Units                *decimal.Decimal
	GoodValue            *decimal.Decimal
	GoodValueUSD         *decimal.Decimal
	PaymentAmount        *decimal.Decimal
	DiscountAmount       *decimal.Decimal
	PromotionID          *uuid.UUID
	Duration             *uint32
	OrderType            *types.OrderType
	InvestmentType       *types.InvestmentType
	CouponIDs            []uuid.UUID
	PaymentType          *types.PaymentType
	CoinTypeID           *uuid.UUID
	PaymentCoinTypeID    *uuid.UUID
	TransferAmount       *decimal.Decimal
	BalanceAmount        *decimal.Decimal
	CoinUSDCurrency      *decimal.Decimal
	LocalCoinUSDCurrency *decimal.Decimal
	LiveCoinUSDCurrency  *decimal.Decimal
	PaymentAccountID     *uuid.UUID
	PaymentStartAmount   *decimal.Decimal
	OrderState           *types.OrderState
	StartMode            *types.OrderStartMode
	StartAt              *uint32
	EndAt                *uint32
	LastBenefitAt        *uint32
	BenefitState         *types.BenefitState
	UserSetPaid          *bool
	UserSetCanceled      *bool
	AdminSetCanceled     *bool
	PaymentTransactionID *string
	PaymentFinishAmount  *decimal.Decimal
	PaymentState         *types.PaymentState
	OutOfGasHours        *uint32
	CompensateHours      *uint32
	AppGoodStockLockID   *uuid.UUID
	LedgerLockID         *uuid.UUID
	Rollback             *bool
	Simulate             *bool
	RenewState           *types.OrderRenewState
	RenewNotifyAt        *uint32
	CreateMethod         *types.OrderCreateMethod
	MultiPaymentCoins    *bool
	PaymentAmounts       []*npool.PaymentAmount
	Reqs                 []*OrderReq
	Conds                *ordercrud.Conds
	Offset               int32
	Limit                int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
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
		h.AppID = &_id
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
		h.UserID = &_id
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
		h.GoodID = &_id
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
		h.AppGoodID = &_id
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
		handler, err := NewHandler(
			ctx,
			WithEntID(id, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistOrder(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid parentorderid")
		}
		h.ParentOrderID = handler.EntID
		return nil
	}
}

func WithUnits(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid units")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("units is less than or equal to 0")
		}
		h.Units = &amount
		return nil
	}
}

func WithGoodValue(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid goodvalue")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("goodvalue is less than to 0")
		}
		h.GoodValue = &amount
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
			return fmt.Errorf("goodvalueusd is less than 0")
		}
		h.GoodValueUSD = &amount
		return nil
	}
}

func WithPaymentAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("paymentamount is less than 0")
		}
		h.PaymentAmount = &amount
		return nil
	}
}

func WithDiscountAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid discountamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("discountamount is less than 0")
		}
		h.DiscountAmount = &amount
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

func WithDuration(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return fmt.Errorf("invalid duration")
			}
			return nil
		}
		h.Duration = duration
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
		h.OrderType = orderType
		return nil
	}
}

func WithInvestmentType(_type *types.InvestmentType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid investmenttype")
			}
			return nil
		}
		switch *_type {
		case types.InvestmentType_FullPayment:
		case types.InvestmentType_UnionMining:
		default:
			return fmt.Errorf("invalid investmenttype")
		}
		h.InvestmentType = _type
		return nil
	}
}

func WithCouponIDs(ids []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if len(ids) == 0 {
			if must {
				return fmt.Errorf("invalid couponids")
			}
			return nil
		}
		_ids := []uuid.UUID{}
		idMap := map[uuid.UUID]struct{}{}
		for _, id := range ids {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			_ids = append(_ids, _id)
			idMap[_id] = struct{}{}
		}
		if len(_ids) != len(idMap) {
			return fmt.Errorf("invalid couponids")
		}
		h.CouponIDs = _ids
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
		case types.PaymentType_PayWithOffline:
		case types.PaymentType_PayWithNoPayment:
		default:
			return fmt.Errorf("invalid paymentType")
		}
		h.PaymentType = paymentType
		return nil
	}
}

func WithPaymentAccountID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid paymentaccountid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.PaymentAccountID = &_id
		return nil
	}
}

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithPaymentCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid paymentcointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.PaymentCoinTypeID = &_id
		return nil
	}
}

//nolint:dupl
func WithPaymentStartAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentstartamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("paymentstartamount is less than 0")
		}
		h.PaymentStartAmount = &amount
		return nil
	}
}

func WithTransferAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid transferamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("transferamount is less than 0")
		}
		h.TransferAmount = &amount
		return nil
	}
}

//nolint:dupl
func WithBalanceAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid balanceamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("balanceamount is less than 0")
		}
		h.BalanceAmount = &amount
		return nil
	}
}

//nolint:dupl
func WithCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid coinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("coinusdcurrency is less than to 0")
		}
		h.CoinUSDCurrency = &amount
		return nil
	}
}

//nolint:dupl
func WithLocalCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid localcoinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("localcoinusdcurrency is less than to 0")
		}
		h.LocalCoinUSDCurrency = &amount
		return nil
	}
}

//nolint:dupl
func WithLiveCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid livecoinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("livecoinusdcurrency is less than to 0")
		}
		h.LiveCoinUSDCurrency = &amount
		return nil
	}
}

//nolint:gocyclo
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
		h.OrderState = state
		return nil
	}
}

func WithStartMode(startMode *types.OrderStartMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startMode == nil {
			if must {
				return fmt.Errorf("invalid startmode")
			}
			return nil
		}
		switch *startMode {
		case types.OrderStartMode_OrderStartConfirmed:
		case types.OrderStartMode_OrderStartTBD:
		case types.OrderStartMode_OrderStartInstantly:
		case types.OrderStartMode_OrderStartNextDay:
		case types.OrderStartMode_OrderStartPreset:
		default:
			return fmt.Errorf("invalid startmode")
		}
		h.StartMode = startMode
		return nil
	}
}

func WithStartAt(startAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startAt == nil {
			if must {
				return fmt.Errorf("invalid startat")
			}
			return nil
		}
		now := uint32(time.Now().Unix())
		if *startAt < now {
			return fmt.Errorf("invalid startat")
		}
		h.StartAt = startAt
		return nil
	}
}

func WithEndAt(endAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if endAt == nil {
			if must {
				return fmt.Errorf("invalid endat")
			}
			return nil
		}
		h.EndAt = endAt
		return nil
	}
}

func WithLastBenefitAt(lastBenefitAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lastBenefitAt == nil {
			if must {
				return fmt.Errorf("invalid lastbenefitat")
			}
			return nil
		}
		h.LastBenefitAt = lastBenefitAt
		return nil
	}
}

func WithBenefitState(benefitState *types.BenefitState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if benefitState == nil {
			if must {
				return fmt.Errorf("invalid benefitstate")
			}
			return nil
		}
		switch *benefitState {
		case types.BenefitState_BenefitWait:
		case types.BenefitState_BenefitCalculated:
		case types.BenefitState_BenefitBookKept:
		default:
			return fmt.Errorf("invalid benefitstate")
		}
		h.BenefitState = benefitState
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
		h.UserSetPaid = value
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
		h.UserSetCanceled = value
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
		h.AdminSetCanceled = value
		return nil
	}
}

func WithPaymentTransactionID(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymenttransactionid")
			}
			return nil
		}
		h.PaymentTransactionID = value
		return nil
	}
}

func WithPaymentFinishAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentfinishamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("paymentfinishamount is less than 0")
		}
		h.PaymentFinishAmount = &amount
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
		h.PaymentState = state
		return nil
	}
}

func WithOutOfGasHours(outOfGasHours *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if outOfGasHours == nil {
			if must {
				return fmt.Errorf("invalid outofgashours")
			}
			return nil
		}
		h.OutOfGasHours = outOfGasHours
		return nil
	}
}

func WithCompensateHours(compensateHours *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if compensateHours == nil {
			if must {
				return fmt.Errorf("invalid compensatehours")
			}
			return nil
		}
		h.CompensateHours = compensateHours
		return nil
	}
}

func WithAppGoodStockLockID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appgoodstocklockid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppGoodStockLockID = &_id
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
		h.LedgerLockID = &_id
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

func WithSimulate(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid simulate")
			}
			return nil
		}
		h.Simulate = value
		return nil
	}
}

func WithRenewState(e *types.OrderRenewState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid renewstate")
			}
			return nil
		}
		switch *e {
		case types.OrderRenewState_OrderRenewWait:
		case types.OrderRenewState_OrderRenewCheck:
		case types.OrderRenewState_OrderRenewNotify:
		case types.OrderRenewState_OrderRenewExecute:
		case types.OrderRenewState_OrderRenewFail:
		default:
			return fmt.Errorf("invalid renewstate")
		}
		h.RenewState = e
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
		h.CreateMethod = e
		return nil
	}
}

func WithRenewNotifyAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RenewNotifyAt = n
		return nil
	}
}

func WithMultiPaymentCoins(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MultiPaymentCoins = b
		return nil
	}
}

func WithPaymentAmounts(amounts []*npool.PaymentAmount, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, amount := range amounts {
			_amount, err := decimal.NewFromString(amount.USDCurrency)
			if err != nil {
				return err
			}
			if _amount.Cmp(decimal.NewFromInt(0)) <= 0 {
				return fmt.Errorf("invalid coincurrency")
			}
			_amount, err = decimal.NewFromString(amount.Amount)
			if err != nil {
				return err
			}
			if _amount.Cmp(decimal.NewFromInt(0)) <= 0 {
				return fmt.Errorf("invalid paymentamount")
			}
		}
		h.PaymentAmounts = amounts
		return nil
	}
}

//nolint:funlen,gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &ordercrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op: conds.GetEntID().GetOp(), Val: id,
			}
		}
		if conds.IDs != nil {
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: conds.GetIDs().GetValue()}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: ids}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{Op: conds.GetGoodID().GetOp(), Val: id}
		}
		if conds.GoodIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetGoodIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.GoodIDs = &cruder.Cond{Op: conds.GetGoodIDs().GetOp(), Val: ids}
		}
		if conds.AppGoodID != nil {
			id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppGoodID = &cruder.Cond{Op: conds.GetAppGoodID().GetOp(), Val: id}
		}
		if conds.AppGoodIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetAppGoodIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.AppGoodIDs = &cruder.Cond{Op: conds.GetAppGoodIDs().GetOp(), Val: ids}
		}
		if conds.ParentOrderID != nil {
			id, err := uuid.Parse(conds.GetParentOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ParentOrderID = &cruder.Cond{Op: conds.GetParentOrderID().GetOp(), Val: id}
		}
		if conds.PaymentAmount != nil {
			amount, err := decimal.NewFromString(conds.GetPaymentAmount().GetValue())
			if err != nil {
				return err
			}
			h.Conds.PaymentAmount = &cruder.Cond{Op: conds.GetPaymentAmount().GetOp(), Val: amount}
		}
		if conds.OrderType != nil {
			switch conds.GetOrderType().GetValue() {
			case uint32(types.OrderType_Airdrop):
			case uint32(types.OrderType_Normal):
			case uint32(types.OrderType_Offline):
			default:
				return fmt.Errorf("invalid ordertype")
			}
			_type := conds.GetOrderType().GetValue()
			h.Conds.OrderType = &cruder.Cond{Op: conds.GetOrderType().GetOp(), Val: types.OrderType(_type)}
		}
		if conds.InvestmentType != nil {
			switch conds.GetInvestmentType().GetValue() {
			case uint32(types.InvestmentType_FullPayment):
			case uint32(types.InvestmentType_UnionMining):
			default:
				return fmt.Errorf("invalid investmenttype")
			}
			_type := conds.GetInvestmentType().GetValue()
			h.Conds.InvestmentType = &cruder.Cond{Op: conds.GetInvestmentType().GetOp(), Val: types.InvestmentType(_type)}
		}
		if conds.CouponIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetCouponIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.CouponIDs = &cruder.Cond{Op: conds.GetCouponIDs().GetOp(), Val: ids}
		}
		if conds.PaymentType != nil {
			switch conds.GetPaymentType().GetValue() {
			case uint32(types.PaymentType_PayWithBalanceOnly):
			case uint32(types.PaymentType_PayWithTransferOnly):
			case uint32(types.PaymentType_PayWithTransferAndBalance):
			case uint32(types.PaymentType_PayWithParentOrder):
			case uint32(types.PaymentType_PayWithOffline):
			case uint32(types.PaymentType_PayWithNoPayment):
			default:
				return fmt.Errorf("invalid paymenttype")
			}
			_type := conds.GetPaymentType().GetValue()
			h.Conds.PaymentType = &cruder.Cond{Op: conds.GetPaymentType().GetOp(), Val: types.PaymentType(_type)}
		}
		if conds.PaymentTypes != nil {
			_types := []types.PaymentType{}
			for _, _type := range conds.GetPaymentTypes().GetValue() {
				switch _type {
				case uint32(types.PaymentType_PayWithBalanceOnly):
				case uint32(types.PaymentType_PayWithTransferOnly):
				case uint32(types.PaymentType_PayWithTransferAndBalance):
				case uint32(types.PaymentType_PayWithParentOrder):
				case uint32(types.PaymentType_PayWithOffline):
				case uint32(types.PaymentType_PayWithNoPayment):
				default:
					return fmt.Errorf("invalid paymenttype")
				}
				_types = append(_types, types.PaymentType(_type))
			}
			h.Conds.PaymentTypes = &cruder.Cond{Op: conds.GetPaymentTypes().GetOp(), Val: _types}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{Op: conds.GetCoinTypeID().GetOp(), Val: id}
		}
		if conds.PaymentCoinTypeID != nil {
			id, err := uuid.Parse(conds.GetPaymentCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.PaymentCoinTypeID = &cruder.Cond{Op: conds.GetPaymentCoinTypeID().GetOp(), Val: id}
		}
		if conds.OrderState != nil {
			_state := conds.GetOrderState().GetValue()
			h.Conds.OrderState = &cruder.Cond{Op: conds.GetOrderState().GetOp(), Val: types.OrderState(_state)}
		}
		if conds.StartMode != nil {
			switch conds.GetStartMode().GetValue() {
			case uint32(types.OrderStartMode_OrderStartConfirmed):
			case uint32(types.OrderStartMode_OrderStartTBD):
			case uint32(types.OrderStartMode_OrderStartInstantly):
			case uint32(types.OrderStartMode_OrderStartNextDay):
			case uint32(types.OrderStartMode_OrderStartPreset):
			default:
				return fmt.Errorf("invalid startmode")
			}
			_state := conds.GetStartMode().GetValue()
			h.Conds.StartMode = &cruder.Cond{Op: conds.GetStartMode().GetOp(), Val: types.OrderStartMode(_state)}
		}
		if conds.LastBenefitAt != nil {
			h.Conds.LastBenefitAt = &cruder.Cond{Op: conds.GetLastBenefitAt().GetOp(), Val: conds.GetLastBenefitAt().GetValue()}
		}
		if conds.Simulate != nil {
			h.Conds.Simulate = &cruder.Cond{Op: conds.GetSimulate().GetOp(), Val: conds.GetSimulate().GetValue()}
		}

		if conds.BenefitState != nil {
			switch conds.GetBenefitState().GetValue() {
			case uint32(types.BenefitState_BenefitWait):
			case uint32(types.BenefitState_BenefitCalculated):
			case uint32(types.BenefitState_BenefitBookKept):
			default:
				return fmt.Errorf("invalid benefitstate")
			}
			_state := conds.GetBenefitState().GetValue()
			h.Conds.BenefitState = &cruder.Cond{
				Op:  conds.GetBenefitState().GetOp(),
				Val: types.BenefitState(_state),
			}
		}
		if conds.PaymentTransactionID != nil {
			h.Conds.PaymentTransactionID = &cruder.Cond{
				Op:  conds.GetPaymentTransactionID().GetOp(),
				Val: conds.GetPaymentTransactionID().GetValue(),
			}
		}
		if conds.CreatedAt != nil {
			h.Conds.CreatedAt = &cruder.Cond{
				Op:  conds.GetCreatedAt().GetOp(),
				Val: conds.GetCreatedAt().GetValue(),
			}
		}
		if conds.UpdatedAt != nil {
			h.Conds.UpdatedAt = &cruder.Cond{
				Op:  conds.GetUpdatedAt().GetOp(),
				Val: conds.GetUpdatedAt().GetValue(),
			}
		}
		if conds.AdminSetCanceled != nil {
			h.Conds.AdminSetCanceled = &cruder.Cond{
				Op:  conds.GetAdminSetCanceled().GetOp(),
				Val: conds.GetAdminSetCanceled().GetValue(),
			}
		}
		if conds.UserSetCanceled != nil {
			h.Conds.UserSetCanceled = &cruder.Cond{
				Op:  conds.GetUserSetCanceled().GetOp(),
				Val: conds.GetUserSetCanceled().GetValue(),
			}
		}

		if conds.PaymentState != nil {
			switch conds.GetPaymentState().GetValue() {
			case uint32(types.PaymentState_PaymentStateWait):
			case uint32(types.PaymentState_PaymentStateCanceled):
			case uint32(types.PaymentState_PaymentStateTimeout):
			case uint32(types.PaymentState_PaymentStateDone):
			case uint32(types.PaymentState_PaymentStateNoPayment):
			default:
				return fmt.Errorf("invalid paymentstate")
			}
			_state := conds.GetPaymentState().GetValue()
			h.Conds.PaymentState = &cruder.Cond{Op: conds.GetPaymentState().GetOp(), Val: types.PaymentState(_state)}
		}
		if conds.CouponID != nil {
			id, err := uuid.Parse(conds.GetCouponID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CouponID = &cruder.Cond{Op: conds.GetCouponID().GetOp(), Val: id}
		}
		if conds.OrderStates != nil {
			states := []string{}
			for _, state := range conds.GetOrderStates().GetValue() {
				_state := types.OrderState(state)
				states = append(states, _state.String())
			}
			if len(states) > 0 {
				h.Conds.OrderStates = &cruder.Cond{Op: conds.GetOrderStates().GetOp(), Val: states}
			}
		}
		if conds.ParentOrderIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetParentOrderIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.ParentOrderIDs = &cruder.Cond{Op: conds.GetParentOrderIDs().GetOp(), Val: ids}
		}
		if conds.RenewState != nil {
			h.Conds.RenewState = &cruder.Cond{
				Op:  conds.GetRenewState().GetOp(),
				Val: types.OrderRenewState(conds.GetRenewState().GetValue()),
			}
		}
		if conds.RenewNotifyAt != nil {
			h.Conds.RenewNotifyAt = &cruder.Cond{
				Op:  conds.GetRenewNotifyAt().GetOp(),
				Val: conds.GetRenewNotifyAt().GetValue(),
			}
		}
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

//nolint:funlen,gocyclo
func WithReqs(reqs []*npool.OrderReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*OrderReq{}
		for _, req := range reqs {
			_req := &OrderReq{
				Req: &ordercrud.Req{
					MultiPaymentCoins: req.MultiPaymentCoins,
					PaymentAmounts:    req.PaymentAmounts,
				},
				OrderStateReq: &orderstatecrud.Req{},
			}
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.UserID == nil {
					return fmt.Errorf("invalid userid")
				}
				if req.GoodID == nil {
					return fmt.Errorf("invalid goodid")
				}
				if req.AppGoodID == nil {
					return fmt.Errorf("invalid appgoodid")
				}
				if req.Units == nil {
					return fmt.Errorf("invalid units")
				}
				if req.GoodValueUSD == nil {
					return fmt.Errorf("invalid goodvalueusd")
				}
				if req.Duration == nil {
					return fmt.Errorf("invalid duration")
				}
				if req.OrderType == nil {
					return fmt.Errorf("invalid ordertype")
				}
				if req.InvestmentType == nil {
					return fmt.Errorf("invalid investmenttype")
				}
				if req.PaymentType == nil {
					return fmt.Errorf("invalid paymenttype")
				}
				if req.CoinUSDCurrency == nil {
					return fmt.Errorf("invalid coinusdcurrency")
				}
				if req.LiveCoinUSDCurrency == nil {
					return fmt.Errorf("invalid livecoinusdcurrency")
				}
				if req.StartAt == nil {
					return fmt.Errorf("invalid startat")
				}
				if req.EndAt == nil {
					return fmt.Errorf("invalid endat")
				}
			}
			if req.ID != nil {
				_req.ID = req.ID
			}
			if req.EntID != nil {
				id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &id
				_req.OrderStateReq.OrderID = &id
			}
			if req.AppID != nil {
				id, err := uuid.Parse(*req.AppID)
				if err != nil {
					return err
				}
				_req.AppID = &id
			}
			if req.UserID != nil {
				id, err := uuid.Parse(*req.UserID)
				if err != nil {
					return err
				}
				_req.UserID = &id
			}
			if req.GoodID != nil {
				id, err := uuid.Parse(*req.GoodID)
				if err != nil {
					return err
				}
				_req.GoodID = &id
			}
			if req.AppGoodID != nil {
				id, err := uuid.Parse(*req.AppGoodID)
				if err != nil {
					return err
				}
				_req.AppGoodID = &id
			}
			if req.ParentOrderID != nil {
				id, err := uuid.Parse(*req.ParentOrderID)
				if err != nil {
					return err
				}
				_req.ParentOrderID = &id
			}
			if req.Units != nil {
				amount, err := decimal.NewFromString(*req.Units)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
					return fmt.Errorf("units is less than or equal to 0")
				}
				_req.Units = &amount
			}
			if req.GoodValue != nil {
				amount, err := decimal.NewFromString(*req.GoodValue)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("goodvalue is less than 0")
				}
				_req.GoodValue = &amount
			}
			if req.GoodValueUSD != nil {
				amount, err := decimal.NewFromString(*req.GoodValueUSD)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("goodvalueusd is less than 0")
				}
				_req.GoodValueUSD = &amount
			}
			if req.PaymentAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("paymentamount is less than 0")
				}
				_req.PaymentAmount = &amount
			}
			if req.DiscountAmount != nil {
				amount, err := decimal.NewFromString(*req.DiscountAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("discountamount is less than 0")
				}
				_req.DiscountAmount = &amount
			}
			if req.PromotionID != nil {
				id, err := uuid.Parse(*req.PromotionID)
				if err != nil {
					return err
				}
				_req.PromotionID = &id
			}
			if req.Duration != nil {
				_req.Duration = req.Duration
			}
			if req.OrderType != nil {
				switch *req.OrderType {
				case types.OrderType_Airdrop:
				case types.OrderType_Offline:
				case types.OrderType_Normal:
				default:
					return fmt.Errorf("invalid ordertype")
				}
				_req.OrderType = req.OrderType
			}
			if req.CoinTypeID != nil {
				id, err := uuid.Parse(*req.CoinTypeID)
				if err != nil {
					return err
				}
				_req.CoinTypeID = &id
			}
			if req.PaymentCoinTypeID != nil {
				id, err := uuid.Parse(*req.PaymentCoinTypeID)
				if err != nil {
					return err
				}
				_req.PaymentCoinTypeID = &id
			}
			if req.TransferAmount != nil {
				amount, err := decimal.NewFromString(*req.TransferAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("transferamount is less than 0")
				}
				_req.TransferAmount = &amount
			}
			if req.BalanceAmount != nil {
				amount, err := decimal.NewFromString(*req.BalanceAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("balanceamount is less than 0")
				}
				_req.BalanceAmount = &amount
			}
			if req.CoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(*req.CoinUSDCurrency)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("coinusdcurrency is less than to 0")
				}
				_req.CoinUSDCurrency = &amount
			}
			if req.LocalCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(*req.LocalCoinUSDCurrency)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("localcoinusdcurrency is less than to 0")
				}
				_req.LocalCoinUSDCurrency = &amount
			}
			if req.LiveCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(*req.LiveCoinUSDCurrency)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("livecoinusdcurrency is less than to 0")
				}
				_req.LiveCoinUSDCurrency = &amount
			}
			if req.InvestmentType != nil {
				switch *req.InvestmentType {
				case types.InvestmentType_FullPayment:
				case types.InvestmentType_UnionMining:
				default:
					return fmt.Errorf("invalid investmenttype")
				}
				_req.InvestmentType = req.InvestmentType
			}
			if req.CouponIDs != nil {
				_ids := []uuid.UUID{}
				for _, id := range req.CouponIDs {
					_id, err := uuid.Parse(id)
					if err != nil {
						return err
					}
					_ids = append(_ids, _id)
				}
				_req.CouponIDs = _ids
			}
			if req.Simulate != nil {
				_req.Simulate = req.Simulate
			}
			if req.CreateMethod != nil {
				switch *req.CreateMethod {
				case types.OrderCreateMethod_OrderCreatedByPurchase:
				case types.OrderCreateMethod_OrderCreatedByAdmin:
				case types.OrderCreateMethod_OrderCreatedByRenew:
				default:
					return fmt.Errorf("invalid createmethod")
				}
				_req.CreateMethod = req.CreateMethod
			}
			if req.OrderState != nil {
				switch *req.OrderState {
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
				_req.OrderStateReq.OrderState = req.OrderState
			}
			if req.AppGoodStockLockID != nil {
				id, err := uuid.Parse(*req.AppGoodStockLockID)
				if err != nil {
					return err
				}
				_req.StockLockReq = &orderlockcrud.Req{}
				_req.StockLockReq.EntID = &id
				_req.StockLockReq.AppID = _req.AppID
				_req.StockLockReq.UserID = _req.UserID
				_req.StockLockReq.OrderID = _req.EntID
				_req.StockLockReq.LockType = types.OrderLockType_LockStock.Enum()
			}
			if req.StartMode != nil {
				switch *req.StartMode {
				case types.OrderStartMode_OrderStartConfirmed:
				case types.OrderStartMode_OrderStartTBD:
				case types.OrderStartMode_OrderStartInstantly:
				case types.OrderStartMode_OrderStartNextDay:
				case types.OrderStartMode_OrderStartPreset:
				default:
					return fmt.Errorf("invalid startmode")
				}
				_req.OrderStateReq.StartMode = req.StartMode
			}
			if req.StartAt != nil {
				now := uint32(time.Now().Unix())
				if *req.StartAt < now {
					return fmt.Errorf("invalid startat")
				}
				_req.OrderStateReq.StartAt = req.StartAt
			}
			if req.EndAt != nil {
				_req.OrderStateReq.EndAt = req.EndAt
			}
			if req.LastBenefitAt != nil {
				_req.OrderStateReq.LastBenefitAt = req.LastBenefitAt
			}
			if req.BenefitState != nil {
				switch *req.BenefitState {
				case types.BenefitState_BenefitWait:
				case types.BenefitState_BenefitCalculated:
				case types.BenefitState_BenefitBookKept:
				default:
					return fmt.Errorf("invalid benefitstate")
				}
				_req.OrderStateReq.BenefitState = req.BenefitState
			}
			if req.UserSetPaid != nil {
				_req.OrderStateReq.UserSetPaid = req.UserSetPaid
			}
			if req.UserSetCanceled != nil {
				_req.OrderStateReq.UserSetCanceled = req.UserSetCanceled
			}
			if req.AdminSetCanceled != nil {
				_req.OrderStateReq.AdminSetCanceled = req.AdminSetCanceled
			}
			if req.PaymentTransactionID != nil {
				_req.OrderStateReq.PaymentTransactionID = req.PaymentTransactionID
			}
			if req.PaymentFinishAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentFinishAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("paymentfinishamount is less than 0")
				}
				_req.OrderStateReq.PaymentFinishAmount = &amount
			}
			if req.PaymentState != nil {
				switch *req.PaymentState {
				case types.PaymentState_PaymentStateWait:
				case types.PaymentState_PaymentStateDone:
				case types.PaymentState_PaymentStateCanceled:
				case types.PaymentState_PaymentStateTimeout:
				case types.PaymentState_PaymentStateNoPayment:
				default:
					return fmt.Errorf("invalid paymentstate")
				}
				_req.OrderStateReq.PaymentState = req.PaymentState
			}
			if req.OutOfGasHours != nil {
				_req.OrderStateReq.OutOfGasHours = req.OutOfGasHours
			}
			if req.CompensateHours != nil {
				_req.OrderStateReq.CompensateHours = req.CompensateHours
			}

			if (req.BalanceAmount != nil && _req.BalanceAmount.Cmp(decimal.NewFromInt(0)) > 0) ||
				len(h.PaymentAmounts) > 0 { // In this case one ledger lock will relevant to multiple statements
				if req.LedgerLockID == nil {
					return fmt.Errorf("invalid ledgerlockid")
				}
				id, err := uuid.Parse(*req.LedgerLockID)
				if err != nil {
					return err
				}
				_req.BalanceLockReq = &orderlockcrud.Req{
					EntID:    &id,
					AppID:    _req.AppID,
					UserID:   _req.UserID,
					OrderID:  _req.EntID,
					LockType: types.OrderLockType_LockBalance.Enum(),
				}
			}

			if req.PaymentType != nil {
				_req.PaymentType = req.PaymentType
				if must {
					if err := _req.CheckOrderType(); err != nil {
						return err
					}
					has, err := _req.HasPayment()
					if err != nil {
						return err
					}
					if !has {
						_reqs = append(_reqs, _req)
						continue
					}
					if req.PaymentCoinTypeID == nil {
						return fmt.Errorf("invalid paymentcointypeid")
					}
				}
			}

			_req.PaymentReq = &paymentcrud.Req{
				OrderID: _req.EntID,
				AppID:   _req.AppID,
				GoodID:  _req.GoodID,
				UserID:  _req.UserID,
			}
			if req.PaymentAccountID != nil {
				id, err := uuid.Parse(*req.PaymentAccountID)
				if err != nil {
					return err
				}
				_req.PaymentReq.AccountID = &id
			}

			if req.PaymentStartAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentStartAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("invalid paymentstartamount")
				}
				_req.PaymentReq.StartAmount = &amount
			}

			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
