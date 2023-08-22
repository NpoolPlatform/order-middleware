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
	ID                        *uuid.UUID
	AppID                     *uuid.UUID
	GoodID                    *uuid.UUID
	UserID                    *uuid.UUID
	ParentOrderID             *uuid.UUID
	PayWithParent             *bool
	Units                     *decimal.Decimal
	PromotionID               *uuid.UUID
	UserSpecialReductionID    *uuid.UUID
	StartAt                   *uint32
	EndAt                     *uint32
	Type                      *basetypes.OrderType
	State                     *basetypes.OrderState
	CouponIDs                 []uuid.UUID
	LastBenefitAt             *uint32
	PaymentState              *basetypes.PaymentState
	InvestmentType            *basetypes.InvestmentType
	PaymentID                 *uuid.UUID
	PaymentAccountID          *uuid.UUID
	PaymentAccountStartAmount *decimal.Decimal
	PaymentAmount             *decimal.Decimal
	PayWithBalanceAmount      *decimal.Decimal
	PaymentCoinUSDCurrency    *decimal.Decimal
	PaymentLocalUSDCurrency   *decimal.Decimal
	PaymentLiveUSDCurrency    *decimal.Decimal
	PaymentCoinID             *uuid.UUID
	PaymentFinishAmount       *decimal.Decimal
	PaymentUserSetCanceled    *bool
	PaymentFakePayment        *bool
	Reqs                      []*ordercrud.Req
	Conds                     *ordercrud.Conds
	Offset                    int32
	Limit                     int32
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

func WithPayWithParent(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paywithparent")
			}
			return nil
		}
		h.PayWithParent = value
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

func WithUserSpecialReductionID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userspecialreductionid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserSpecialReductionID = &_id
		return nil
	}
}

func WithStartAt(start *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if start == nil {
			if must {
				return fmt.Errorf("invalid startat")
			}
			return nil
		}
		h.StartAt = start
		return nil
	}
}

func WithEndAt(end *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if end == nil {
			if must {
				return fmt.Errorf("invalid endat")
			}
			return nil
		}
		h.EndAt = end
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

func WithType(orderType *basetypes.OrderType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderType == nil {
			if must {
				return fmt.Errorf("invalid type")
			}
			return nil
		}
		switch *orderType {
		case basetypes.OrderType_Airdrop:
		case basetypes.OrderType_Offline:
		case basetypes.OrderType_Normal:
		default:
			return fmt.Errorf("invalid order type")
		}
		h.Type = orderType
		return nil
	}
}

func WithState(state *basetypes.OrderState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid state")
			}
			return nil
		}
		switch *state {
		case basetypes.OrderState_OrderStateCanceled:
		case basetypes.OrderState_OrderStateExpired:
		case basetypes.OrderState_OrderStatePaid:
		case basetypes.OrderState_OrderStateWaitPayment:
		case basetypes.OrderState_OrderStatePaymentTimeout:
		case basetypes.OrderState_OrderStateInService:
		default:
			return fmt.Errorf("invalid lockedby")
		}
		h.State = state
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
		case basetypes.PaymentState_PaymentStateTimeOut:
		case basetypes.PaymentState_PaymentStateDone:
		default:
			return fmt.Errorf("invalid paymentstate")
		}
		h.PaymentState = state
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
		h.PaymentID = &_id
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

func WithPaymentCoinID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid paymentcoinid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.PaymentCoinID = &_id
		return nil
	}
}

func WithPaymentAccountStartAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentsccountstartamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PaymentAccountStartAmount = &amount
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

func WithPayWithBalanceAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paywithbalanceamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PayWithBalanceAmount = &amount
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

func WithPaymentLocalUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentlocalusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PaymentLocalUSDCurrency = &amount
		return nil
	}
}

func WithPaymentLiveUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentliveusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PaymentLiveUSDCurrency = &amount
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

func WithPaymentUserSetCanceled(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentusersetcanceled")
			}
			return nil
		}
		h.PaymentUserSetCanceled = value
		return nil
	}
}

func WithPaymentFakePayment(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid paymentfakepayment")
			}
			return nil
		}
		h.PaymentFakePayment = value
		return nil
	}
}

//nolint:gocyclo
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
		if conds.UserSpecialReductionID != nil {
			id, err := uuid.Parse(conds.GetUserSpecialReductionID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserSpecialReductionID = &cruder.Cond{Op: conds.GetUserSpecialReductionID().GetOp(), Val: id}
		}
		if conds.State != nil {
			switch conds.GetState().GetValue() {
			case uint32(basetypes.OrderState_OrderStateCanceled):
			case uint32(basetypes.OrderState_OrderStateExpired):
			case uint32(basetypes.OrderState_OrderStateInService):
			case uint32(basetypes.OrderState_OrderStatePaid):
			case uint32(basetypes.OrderState_OrderStatePaymentTimeout):
			case uint32(basetypes.OrderState_OrderStateWaitPayment):
			default:
				return fmt.Errorf("invalid state")
			}
			_state := conds.GetState().GetValue()
			h.Conds.State = &cruder.Cond{Op: conds.GetState().GetOp(), Val: basetypes.PaymentState(_state)}
		}
		if conds.Type != nil {
			switch conds.GetType().GetValue() {
			case uint32(basetypes.OrderType_Airdrop):
			case uint32(basetypes.OrderType_Normal):
			case uint32(basetypes.OrderType_Offline):
			default:
				return fmt.Errorf("invalid order type")
			}
			_type := conds.GetType().GetValue()
			h.Conds.Type = &cruder.Cond{Op: conds.GetType().GetOp(), Val: basetypes.OrderType(_type)}
		}
		if conds.CouponID != nil {
			id, err := uuid.Parse(conds.GetCouponID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CouponID = &cruder.Cond{Op: conds.GetCouponID().GetOp(), Val: id}
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
		if conds.LastBenefitAt != nil {
			h.Conds.LastBenefitAt = &cruder.Cond{Op: conds.GetLastBenefitAt().GetOp(), Val: conds.GetLastBenefitAt().GetValue()}
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

//nolint:gocyclo
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
			if req.GoodID != nil {
				id, err := uuid.Parse(req.GetGoodID())
				if err != nil {
					return err
				}
				_req.GoodID = &id
			}
			if req.UserID != nil {
				id, err := uuid.Parse(req.GetUserID())
				if err != nil {
					return err
				}
				_req.UserID = &id
			}
			if req.PaymentID != nil {
				id, err := uuid.Parse(req.GetPaymentID())
				if err != nil {
					return err
				}
				_req.PaymentID = &id
			}
			if req.Start != nil {
				_req.StartAt = req.Start
			}
			if req.PaymentState != nil {
				switch req.GetPaymentState() {
				case basetypes.PaymentState_PaymentStateWait:
				case basetypes.PaymentState_PaymentStateDone:
				case basetypes.PaymentState_PaymentStateCanceled:
				case basetypes.PaymentState_PaymentStateTimeOut:
				default:
					return fmt.Errorf("invalid State")
				}
				_req.PaymentState = req.PaymentState
			}
			if req.LastBenefitAt != nil {
				_req.LastBenefitAt = req.LastBenefitAt
			}
			if req.Canceled != nil {
				_req.PaymentUserSetCanceled = req.Canceled
			}
			if req.PaymentFinishAmount != nil {
				amount, err := decimal.NewFromString(*req.PaymentFinishAmount)
				if err != nil {
					return err
				}
				_req.PaymentFinishAmount = &amount
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
