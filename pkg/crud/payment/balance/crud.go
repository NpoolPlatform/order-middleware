package paymentbalance

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpaymentbalance "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbalance"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                *uuid.UUID
	PaymentID            *uuid.UUID
	CoinTypeID           *uuid.UUID
	Amount               *decimal.Decimal
	CoinUSDCurrency      *decimal.Decimal
	LocalCoinUSDCurrency *decimal.Decimal
	LiveCoinUSDCurrency  *decimal.Decimal
	DeletedAt            *uint32
}

func CreateSet(c *ent.PaymentBalanceCreate, req *Req) *ent.PaymentBalanceCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.PaymentID != nil {
		c.SetPaymentID(*req.PaymentID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
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
	return c
}

func UpdateSet(u *ent.PaymentBalanceUpdateOne, req *Req) *ent.PaymentBalanceUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	PaymentID  *cruder.Cond
	PaymentIDs *cruder.Cond
	CoinTypeID *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.PaymentBalanceQuery, conds *Conds) (*ent.PaymentBalanceQuery, error) {
	q.Where(entpaymentbalance.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entpaymentbalance.ID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		if len(ids) > 0 {
			switch conds.IDs.Op {
			case cruder.IN:
				q.Where(entpaymentbalance.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid payment field")
			}
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entpaymentbalance.EntID(id))
		case cruder.NEQ:
			q.Where(entpaymentbalance.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		if len(ids) > 0 {
			switch conds.EntIDs.Op {
			case cruder.IN:
				q.Where(entpaymentbalance.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid payment field")
			}
		}
	}
	if conds.PaymentID != nil {
		id, ok := conds.PaymentID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.PaymentID.Op {
		case cruder.EQ:
			q.Where(entpaymentbalance.PaymentID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.PaymentIDs != nil {
		ids, ok := conds.PaymentIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid paymentids")
		}
		if len(ids) > 0 {
			switch conds.PaymentIDs.Op {
			case cruder.IN:
				q.Where(entpaymentbalance.PaymentIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid payment field")
			}
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entpaymentbalance.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	return q, nil
}
