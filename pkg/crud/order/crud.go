package order

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                   *uuid.UUID
	AppID                *uuid.UUID
	UserID               *uuid.UUID
	GoodID               *uuid.UUID
	AppGoodID            *uuid.UUID
	PaymentID            *uuid.UUID
	ParentOrderID        *uuid.UUID
	Units                *decimal.Decimal
	GoodValue            *decimal.Decimal
	GoodValueUSD         *decimal.Decimal
	PaymentAmount        *decimal.Decimal
	DiscountAmount       *decimal.Decimal
	PromotionID          *uuid.UUID
	DurationDays         *uint32
	OrderType            *basetypes.OrderType
	InvestmentType       *basetypes.InvestmentType
	CouponIDs            *[]uuid.UUID
	PaymentType          *basetypes.PaymentType
	PaymentCoinTypeID    *uuid.UUID
	CoinTypeID           *uuid.UUID
	TransferAmount       *decimal.Decimal
	BalanceAmount        *decimal.Decimal
	CoinUSDCurrency      *decimal.Decimal
	LocalCoinUSDCurrency *decimal.Decimal
	LiveCoinUSDCurrency  *decimal.Decimal
	CreatedAt            *uint32
	DeletedAt            *uint32
}

//nolint:funlen,gocyclo
func CreateSet(c *ent.OrderCreate, req *Req) *ent.OrderCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.PaymentID != nil {
		c.SetPaymentID(*req.PaymentID)
	}
	if req.ParentOrderID != nil {
		c.SetParentOrderID(*req.ParentOrderID)
	}
	if req.Units != nil {
		c.SetUnitsV1(*req.Units)
	}
	if req.GoodValue != nil {
		c.SetGoodValue(*req.GoodValue)
	}
	if req.GoodValueUSD != nil {
		c.SetGoodValueUsd(*req.GoodValueUSD)
	}
	if req.PaymentAmount != nil {
		c.SetPaymentAmount(*req.PaymentAmount)
	}
	if req.DiscountAmount != nil {
		c.SetDiscountAmount(*req.DiscountAmount)
	}
	if req.PromotionID != nil {
		c.SetPromotionID(*req.PromotionID)
	}
	if req.DurationDays != nil {
		c.SetDurationDays(*req.DurationDays)
	}
	if req.OrderType != nil {
		c.SetOrderType(req.OrderType.String())
	}
	if req.InvestmentType != nil {
		c.SetInvestmentType(req.InvestmentType.String())
	}
	if req.CouponIDs != nil {
		c.SetCouponIds(*req.CouponIDs)
	}
	if req.PaymentType != nil {
		c.SetPaymentType(req.PaymentType.String())
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.PaymentCoinTypeID != nil {
		c.SetPaymentCoinTypeID(*req.PaymentCoinTypeID)
	}
	if req.TransferAmount != nil {
		c.SetTransferAmount(*req.TransferAmount)
	}
	if req.BalanceAmount != nil {
		c.SetBalanceAmount(*req.BalanceAmount)
	}
	if req.CoinUSDCurrency != nil {
		c.SetCoinUsdCurrency(*req.CoinUSDCurrency)
	}
	if req.LocalCoinUSDCurrency != nil {
		c.SetLocalCoinUsdCurrency(*req.LocalCoinUSDCurrency)
	}
	if req.LiveCoinUSDCurrency != nil {
		c.SetLiveCoinUsdCurrency(*req.LiveCoinUSDCurrency)
	}
	if req.CreatedAt != nil {
		c.SetCreatedAt(*req.CreatedAt)
	}

	return c
}

func UpdateSet(u *ent.OrderUpdateOne, req *Req) *ent.OrderUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	orderstatecrud.Conds
	ID                *cruder.Cond
	AppID             *cruder.Cond
	UserID            *cruder.Cond
	GoodID            *cruder.Cond
	AppGoodID         *cruder.Cond
	ParentOrderID     *cruder.Cond
	PaymentAmount     *cruder.Cond
	OrderType         *cruder.Cond
	InvestmentType    *cruder.Cond
	PaymentType       *cruder.Cond
	CoinTypeID        *cruder.Cond
	PaymentCoinTypeID *cruder.Cond
	IDs               *cruder.Cond
	CouponID          *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.OrderQuery, conds *Conds) (*ent.OrderQuery, error) {
	q.Where(entorder.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entorder.ID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entorder.AppID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entorder.UserID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entorder.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entorder.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.ParentOrderID != nil {
		id, ok := conds.ParentOrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid parentorderid")
		}
		switch conds.ParentOrderID.Op {
		case cruder.EQ:
			q.Where(entorder.ParentOrderID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.PaymentAmount != nil {
		paymentAmount, ok := conds.PaymentAmount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid paymentamount")
		}
		switch conds.PaymentAmount.Op {
		case cruder.EQ:
			q.Where(entorder.PaymentAmount(paymentAmount))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.OrderType != nil {
		ordertype, ok := conds.OrderType.Val.(basetypes.OrderType)
		if !ok {
			return nil, fmt.Errorf("invalid ordertype")
		}
		switch conds.OrderType.Op {
		case cruder.EQ:
			q.Where(entorder.OrderType(ordertype.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.InvestmentType != nil {
		investmenttype, ok := conds.InvestmentType.Val.(basetypes.InvestmentType)
		if !ok {
			return nil, fmt.Errorf("invalid investmenttype")
		}
		switch conds.InvestmentType.Op {
		case cruder.EQ:
			q.Where(entorder.InvestmentType(investmenttype.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.PaymentType != nil {
		paymenttype, ok := conds.PaymentType.Val.(basetypes.PaymentType)
		if !ok {
			return nil, fmt.Errorf("invalid paymenttype")
		}
		switch conds.PaymentType.Op {
		case cruder.EQ:
			q.Where(entorder.PaymentType(paymenttype.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entorder.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.PaymentCoinTypeID != nil {
		id, ok := conds.PaymentCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid paymentcointypeid")
		}
		switch conds.PaymentCoinTypeID.Op {
		case cruder.EQ:
			q.Where(entorder.PaymentCoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		if len(ids) > 0 {
			switch conds.IDs.Op {
			case cruder.IN:
				q.Where(entorder.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid order field")
			}
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.CouponID.Op {
		case cruder.LIKE:
			q.Where(func(selector *sql.Selector) {
				selector.Where(sqljson.ValueContains(entorder.FieldCouponIds, id))
			})
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	return q, nil
}
