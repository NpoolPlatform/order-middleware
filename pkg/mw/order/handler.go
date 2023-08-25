package order

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID                          *uuid.UUID
	AppID                       *uuid.UUID
	UserID                      *uuid.UUID
	GoodID                      *uuid.UUID
	AppGoodID                   *uuid.UUID
	ParentOrderID               *uuid.UUID
	Units                       *decimal.Decimal
	GoodValue                   *decimal.Decimal
	PaymentAmount               *decimal.Decimal
	DiscountAmount              *decimal.Decimal
	PromotionID                 *uuid.UUID
	DurationDays                *uint32
	OrderType                   *basetypes.OrderType
	InvestmentType              *basetypes.InvestmentType
	CouponIDs                   []uuid.UUID
	PaymentType                 *basetypes.PaymentType
	PaymentAccountID            *uuid.UUID
	PaymentCoinTypeID           *uuid.UUID
	PaymentStartAmount          *decimal.Decimal
	PaymentTransferAmount       *decimal.Decimal
	PaymentBalanceAmount        *decimal.Decimal
	PaymentCoinUSDCurrency      *decimal.Decimal
	PaymentLocalCoinUSDCurrency *decimal.Decimal
	PaymentLiveCoinUSDCurrency  *decimal.Decimal
	OrderState                  *basetypes.OrderState
	StartMode                   *basetypes.OrderStartMode
	StartAt                     *uint32
	EndAt                       *uint32
	LastBenefitAt               *uint32
	BenefitState                *basetypes.BenefitState
	UserSetPaid                 *bool
	UserSetCanceled             *bool
	PaymentTransactionID        *string
	PaymentFinishAmount         *decimal.Decimal
	PaymentState                *basetypes.PaymentState
	OutOfGasHours               *uint32
	CompensateHours             *uint32
	Reqs                        []*ordercrud.Req
	Conds                       *ordercrud.Conds
	Offset                      int32
	Limit                       int32
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

func WithID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ParentOrderID = &_id
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
		h.GoodValue = &amount
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

func WithDurationDays(durationDays *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if durationDays == nil {
			if must {
				return fmt.Errorf("invalid durationdays")
			}
			return nil
		}
		h.DurationDays = durationDays
		return nil
	}
}

func WithOrderType(orderType *basetypes.OrderType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderType == nil {
			if must {
				return fmt.Errorf("invalid ordertype")
			}
			return nil
		}
		switch *orderType {
		case basetypes.OrderType_Airdrop:
		case basetypes.OrderType_Offline:
		case basetypes.OrderType_Normal:
		default:
			return fmt.Errorf("invalid ordertype")
		}
		h.OrderType = orderType
		return nil
	}
}

func WithInvestmentType(_type *basetypes.InvestmentType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid investmenttype")
			}
			return nil
		}
		switch *_type {
		case basetypes.InvestmentType_FullPayment:
		case basetypes.InvestmentType_UnionMining:
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
		for _, id := range ids {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			_ids = append(_ids, _id)
		}
		h.CouponIDs = _ids
		return nil
	}
}

func WithPaymentType(paymentType *basetypes.PaymentType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if paymentType == nil {
			if must {
				return fmt.Errorf("invalid paymenttype")
			}
			return nil
		}
		switch *paymentType {
		case basetypes.PaymentType_PayWithBalanceOnly:
		case basetypes.PaymentType_PayWithTransferOnly:
		case basetypes.PaymentType_PayWithTransferAndBalance:
		case basetypes.PaymentType_PayWithParentOrder:
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
		h.PaymentStartAmount = &amount
		return nil
	}
}

func WithPaymentTransferAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymenttransferamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PaymentTransferAmount = &amount
		return nil
	}
}

func WithPaymentBalanceAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentbalanceamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PaymentBalanceAmount = &amount
		return nil
	}
}

func WithPaymentCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentcoinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PaymentCoinUSDCurrency = &amount
		return nil
	}
}

func WithPaymentLocalCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentlocalcoinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PaymentLocalCoinUSDCurrency = &amount
		return nil
	}
}

func WithPaymentLiveCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentlivecoinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PaymentLiveCoinUSDCurrency = &amount
		return nil
	}
}

func WithOrderState(state *basetypes.OrderState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid orderstate")
			}
			return nil
		}
		switch *state {
		case basetypes.OrderState_OrderStateWaitPayment:
		case basetypes.OrderState_OrderStateCheckPayment:
		case basetypes.OrderState_OrderStatePaid:
		case basetypes.OrderState_OrderStatePaymentTimeout:
		case basetypes.OrderState_OrderStateCanceled:
		case basetypes.OrderState_OrderStateInService:
		case basetypes.OrderState_OrderStateExpired:
		default:
			return fmt.Errorf("invalid orderstate")
		}
		h.OrderState = state
		return nil
	}
}

func WithStartMode(startMode *basetypes.OrderStartMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startMode == nil {
			if must {
				return fmt.Errorf("invalid startmode")
			}
			return nil
		}
		switch *startMode {
		case basetypes.OrderStartMode_OrderStartConfirmed:
		case basetypes.OrderStartMode_OrderStartTBD:
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

func WithBenefitState(benefitState *basetypes.BenefitState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if benefitState == nil {
			if must {
				return fmt.Errorf("invalid benefitstate")
			}
			return nil
		}
		switch *benefitState {
		case basetypes.BenefitState_BenefitWait:
		case basetypes.BenefitState_BenefitCalculated:
		case basetypes.BenefitState_BenefitDone:
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
		h.PaymentFinishAmount = &amount
		return nil
	}
}

func WithPaymentState(state *basetypes.PaymentState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid paymentstate")
			}
			return nil
		}
		switch *state {
		case basetypes.PaymentState_PaymentStateWait:
		case basetypes.PaymentState_PaymentStateCanceled:
		case basetypes.PaymentState_PaymentStateTimeout:
		case basetypes.PaymentState_PaymentStateDone:
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

//nolint:funlen,gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &ordercrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
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
		if conds.AppGoodID != nil {
			id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppGoodID = &cruder.Cond{Op: conds.GetAppGoodID().GetOp(), Val: id}
		}
		if conds.ParentOrderID != nil {
			id, err := uuid.Parse(conds.GetParentOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ParentOrderID = &cruder.Cond{Op: conds.GetParentOrderID().GetOp(), Val: id}
		}
		if conds.OrderType != nil {
			switch conds.GetOrderType().GetValue() {
			case uint32(basetypes.OrderType_Airdrop):
			case uint32(basetypes.OrderType_Normal):
			case uint32(basetypes.OrderType_Offline):
			default:
				return fmt.Errorf("invalid ordertype")
			}
			_type := conds.GetOrderType().GetValue()
			h.Conds.OrderType = &cruder.Cond{Op: conds.GetOrderType().GetOp(), Val: basetypes.OrderType(_type)}
		}
		if conds.InvestmentType != nil {
			switch conds.GetInvestmentType().GetValue() {
			case uint32(basetypes.InvestmentType_FullPayment):
			case uint32(basetypes.InvestmentType_UnionMining):
			default:
				return fmt.Errorf("invalid ordertype")
			}
			_type := conds.GetInvestmentType().GetValue()
			h.Conds.InvestmentType = &cruder.Cond{Op: conds.GetInvestmentType().GetOp(), Val: basetypes.InvestmentType(_type)}
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
			case uint32(basetypes.PaymentState_PaymentStateWait):
			case uint32(basetypes.PaymentState_PaymentStateCanceled):
			case uint32(basetypes.PaymentState_PaymentStateTimeout):
			case uint32(basetypes.PaymentState_PaymentStateDone):
			default:
				return fmt.Errorf("invalid paymenttype")
			}
			_type := conds.GetPaymentType().GetValue()
			h.Conds.PaymentType = &cruder.Cond{Op: conds.GetPaymentType().GetOp(), Val: basetypes.PaymentType(_type)}
		}

		if conds.PaymentCoinTypeID != nil {
			id, err := uuid.Parse(conds.GetPaymentCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.PaymentCoinTypeID = &cruder.Cond{Op: conds.GetPaymentCoinTypeID().GetOp(), Val: id}
		}
		if conds.OrderState != nil {
			switch conds.GetOrderState().GetValue() {
			case uint32(basetypes.OrderState_OrderStateWaitPayment):
			case uint32(basetypes.OrderState_OrderStateCheckPayment):
			case uint32(basetypes.OrderState_OrderStatePaid):
			case uint32(basetypes.OrderState_OrderStatePaymentTimeout):
			case uint32(basetypes.OrderState_OrderStateCanceled):
			case uint32(basetypes.OrderState_OrderStateInService):
			case uint32(basetypes.OrderState_OrderStateExpired):
			default:
				return fmt.Errorf("invalid orderstate")
			}
			_state := conds.GetOrderState().GetValue()
			h.Conds.OrderState = &cruder.Cond{Op: conds.GetOrderState().GetOp(), Val: basetypes.OrderState(_state)}
		}
		if conds.StartMode != nil {
			switch conds.GetStartMode().GetValue() {
			case uint32(basetypes.OrderStartMode_OrderStartConfirmed):
			case uint32(basetypes.OrderStartMode_OrderStartTBD):
			default:
				return fmt.Errorf("invalid orderstate")
			}
			_state := conds.GetStartMode().GetValue()
			h.Conds.StartMode = &cruder.Cond{Op: conds.GetStartMode().GetOp(), Val: basetypes.OrderStartMode(_state)}
		}
		if conds.LastBenefitAt != nil {
			h.Conds.LastBenefitAt = &cruder.Cond{Op: conds.GetLastBenefitAt().GetOp(), Val: conds.GetLastBenefitAt().GetValue()}
		}

		if conds.BenefitState != nil {
			switch conds.GetOrderState().GetValue() {
			case uint32(basetypes.BenefitState_BenefitWait):
			case uint32(basetypes.BenefitState_BenefitCalculated):
			case uint32(basetypes.BenefitState_BenefitDone):
			default:
				return fmt.Errorf("invalid benefitstate")
			}
			_state := conds.GetBenefitState().GetValue()
			h.Conds.BenefitState = &cruder.Cond{Op: conds.GetBenefitState().GetOp(), Val: basetypes.BenefitState(_state)}
		}
		if conds.PaymentTransactionID != nil {
			h.Conds.PaymentTransactionID = &cruder.Cond{Op: conds.GetPaymentTransactionID().GetOp(), Val: conds.GetPaymentTransactionID().GetValue()}
		}

		if conds.PaymentState != nil {
			switch conds.GetOrderState().GetValue() {
			case uint32(basetypes.PaymentState_PaymentStateWait):
			case uint32(basetypes.PaymentState_PaymentStateCanceled):
			case uint32(basetypes.PaymentState_PaymentStateTimeout):
			case uint32(basetypes.PaymentState_PaymentStateDone):
			default:
				return fmt.Errorf("invalid paymentstate")
			}
			_state := conds.GetPaymentState().GetValue()
			h.Conds.PaymentState = &cruder.Cond{Op: conds.GetPaymentState().GetOp(), Val: basetypes.PaymentState(_state)}
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
func WithReqs(reqs []*npool.OrderReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*ordercrud.Req{}
		for _, req := range reqs {
			_req := &ordercrud.Req{}
			if req.ID != nil {
				id, err := uuid.Parse(req.GetID())
				if err != nil {
					return err
				}
				_req.ID = &id
			}
			if req.AppID != nil {
				id, err := uuid.Parse(req.GetAppID())
				if err != nil {
					return err
				}
				_req.AppID = &id
			}
			if req.UserID != nil {
				id, err := uuid.Parse(req.GetUserID())
				if err != nil {
					return err
				}
				_req.UserID = &id
			}
			if req.GoodID != nil {
				id, err := uuid.Parse(req.GetGoodID())
				if err != nil {
					return err
				}
				_req.GoodID = &id
			}
			if req.AppGoodID != nil {
				id, err := uuid.Parse(req.GetAppGoodID())
				if err != nil {
					return err
				}
				_req.AppGoodID = &id
			}
			if req.ParentOrderID != nil {
				id, err := uuid.Parse(req.GetParentOrderID())
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
				_req.Units = &amount
			}
			if req.GoodValue != nil {
				amount, err := decimal.NewFromString(*req.GoodValue)
				if err != nil {
					return err
				}
				_req.GoodValue = &amount
			}
			if req.PaymentAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentAmount)
				if err != nil {
					return err
				}
				_req.PaymentAmount = &amount
			}
			if req.DiscountAmount != nil {
				amount, err := decimal.NewFromString(*req.DiscountAmount)
				if err != nil {
					return err
				}
				_req.DiscountAmount = &amount
			}
			if req.PromotionID != nil {
				id, err := uuid.Parse(req.GetPromotionID())
				if err != nil {
					return err
				}
				_req.PromotionID = &id
			}
			if req.DurationDays != nil {
				_req.DurationDays = req.DurationDays
			}
			if req.OrderType != nil {
				switch req.GetOrderType() {
				case basetypes.OrderType_Airdrop:
				case basetypes.OrderType_Offline:
				case basetypes.OrderType_Normal:
				default:
					return fmt.Errorf("invalid ordertype")
				}
				_req.OrderType = req.OrderType
			}
			if req.InvestmentType != nil {
				switch req.GetInvestmentType() {
				case basetypes.InvestmentType_FullPayment:
				case basetypes.InvestmentType_UnionMining:
				default:
					return fmt.Errorf("invalid investmenttype")
				}
				_req.InvestmentType = req.InvestmentType
			}
			if req.CouponIDs != nil {
				_ids := []uuid.UUID{}
				for _, id := range req.GetCouponIDs() {
					_id, err := uuid.Parse(id)
					if err != nil {
						return err
					}
					_ids = append(_ids, _id)
				}
				_req.CouponIDs = &_ids
			}
			if req.PaymentType != nil {
				switch req.GetPaymentState() {
				case basetypes.PaymentState_PaymentStateWait:
				case basetypes.PaymentState_PaymentStateCanceled:
				case basetypes.PaymentState_PaymentStateTimeout:
				case basetypes.PaymentState_PaymentStateDone:
				default:
					return fmt.Errorf("invalid State")
				}
				_req.PaymentType = req.PaymentType
			}
			if req.PaymentAccountID != nil {
				id, err := uuid.Parse(req.GetPaymentAccountID())
				if err != nil {
					return err
				}
				_req.PaymentAccountID = &id
			}
			if req.PaymentCoinTypeID != nil {
				id, err := uuid.Parse(req.GetPaymentCoinTypeID())
				if err != nil {
					return err
				}
				_req.PaymentCoinTypeID = &id
			}
			if req.PaymentStartAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentStartAmount)
				if err != nil {
					return err
				}
				_req.PaymentStartAmount = &amount
			}
			if req.PaymentTransferAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentTransferAmount)
				if err != nil {
					return err
				}
				_req.PaymentTransferAmount = &amount
			}
			if req.PaymentBalanceAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentBalanceAmount)
				if err != nil {
					return err
				}
				_req.PaymentBalanceAmount = &amount
			}
			if req.PaymentCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(*req.PaymentCoinUSDCurrency)
				if err != nil {
					return err
				}
				_req.PaymentCoinUSDCurrency = &amount
			}
			if req.PaymentLocalCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(*req.PaymentLocalCoinUSDCurrency)
				if err != nil {
					return err
				}
				_req.PaymentLocalCoinUSDCurrency = &amount
			}
			if req.PaymentLiveCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(*req.PaymentLiveCoinUSDCurrency)
				if err != nil {
					return err
				}
				_req.PaymentLiveCoinUSDCurrency = &amount
			}
			if req.OrderState != nil {
				switch req.GetOrderState() {
				case basetypes.OrderState_OrderStateWaitPayment:
				case basetypes.OrderState_OrderStateCheckPayment:
				case basetypes.OrderState_OrderStatePaid:
				case basetypes.OrderState_OrderStatePaymentTimeout:
				case basetypes.OrderState_OrderStateCanceled:
				case basetypes.OrderState_OrderStateInService:
				case basetypes.OrderState_OrderStateExpired:
				default:
					return fmt.Errorf("invalid orderstate")
				}
				_req.OrderState = req.OrderState
			}
			if req.StartMode != nil {
				switch req.GetStartMode() {
				case basetypes.OrderStartMode_OrderStartConfirmed:
				case basetypes.OrderStartMode_OrderStartTBD:
				default:
					return fmt.Errorf("invalid startmode")
				}
				_req.StartMode = req.StartMode
			}
			if req.StartAt != nil {
				_req.StartAt = req.StartAt
			}
			if req.EndAt != nil {
				_req.EndAt = req.EndAt
			}
			if req.LastBenefitAt != nil {
				_req.LastBenefitAt = req.LastBenefitAt
			}
			if req.BenefitState != nil {
				switch req.GetBenefitState() {
				case basetypes.BenefitState_BenefitWait:
				case basetypes.BenefitState_BenefitCalculated:
				case basetypes.BenefitState_BenefitDone:
				default:
					return fmt.Errorf("invalid benefitstate")
				}
				_req.BenefitState = req.BenefitState
			}
			if req.UserSetPaid != nil {
				_req.UserSetPaid = req.UserSetPaid
			}
			if req.UserSetCanceled != nil {
				_req.UserSetCanceled = req.UserSetCanceled
			}
			if req.PaymentTransactionID != nil {
				_req.PaymentTransactionID = req.PaymentTransactionID
			}
			if req.PaymentFinishAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentFinishAmount)
				if err != nil {
					return err
				}
				_req.PaymentFinishAmount = &amount
			}
			if req.PaymentState != nil {
				switch req.GetPaymentState() {
				case basetypes.PaymentState_PaymentStateWait:
				case basetypes.PaymentState_PaymentStateDone:
				case basetypes.PaymentState_PaymentStateCanceled:
				case basetypes.PaymentState_PaymentStateTimeout:
				default:
					return fmt.Errorf("invalid paymentstate")
				}
				_req.PaymentState = req.PaymentState
			}
			if req.OutOfGasHours != nil {
				_req.OutOfGasHours = req.OutOfGasHours
			}
			if req.CompensateHours != nil {
				_req.CompensateHours = req.CompensateHours
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
