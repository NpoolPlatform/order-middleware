package paymenttransfer

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpaymenttransfer "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymenttransfer"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                *uuid.UUID
	OrderID              *uuid.UUID
	CoinTypeID           *uuid.UUID
	AccountID            *uuid.UUID
	Amount               *decimal.Decimal
	StartAmount          *decimal.Decimal
	TransactionID        *string
	FinishAmount         *decimal.Decimal
	CoinUSDCurrency      *decimal.Decimal
	LocalCoinUSDCurrency *decimal.Decimal
	LiveCoinUSDCurrency  *decimal.Decimal
	DeletedAt            *uint32
}

func CreateSet(c *ent.PaymentTransferCreate, req *Req) *ent.PaymentTransferCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.StartAmount != nil {
		c.SetStartAmount(*req.StartAmount)
	}
	if req.TransactionID != nil {
		c.SetTransactionID(*req.TransactionID)
	}
	if req.FinishAmount != nil {
		c.SetFinishAmount(*req.FinishAmount)
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

func UpdateSet(u *ent.PaymentTransferUpdateOne, req *Req) *ent.PaymentTransferUpdateOne {
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
	OrderID    *cruder.Cond
	CoinTypeID *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.PaymentTransferQuery, conds *Conds) (*ent.PaymentTransferQuery, error) {
	q.Where(entpaymenttransfer.DeletedAt(0))
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
			q.Where(entpaymenttransfer.ID(id))
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
				q.Where(entpaymenttransfer.IDIn(ids...))
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
			q.Where(entpaymenttransfer.EntID(id))
		case cruder.NEQ:
			q.Where(entpaymenttransfer.EntIDNEQ(id))
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
				q.Where(entpaymenttransfer.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid payment field")
			}
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entpaymenttransfer.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entpaymenttransfer.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	return q, nil
}
