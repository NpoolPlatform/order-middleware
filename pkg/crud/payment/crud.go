package payment

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                   *uuid.UUID
	AppID                *uuid.UUID
	GoodID               *uuid.UUID
	UserID               *uuid.UUID
	OrderID              *uuid.UUID
	AccountID            *uuid.UUID
	StartAmount          *decimal.Decimal
	Amount               *decimal.Decimal
	PayWithBalanceAmount *decimal.Decimal
	FinishAmount         *decimal.Decimal
	CoinUsdCurrency      *decimal.Decimal
	LocalCoinUsdCurrency *decimal.Decimal
	LiveCoinUsdCurrency  *decimal.Decimal
	CoinInfoID           *uuid.UUID
	State                *basetypes.PaymentState
	ChainTransactionID   *string
	UserSetPaid          *bool
	UserSetCanceled      *bool
	FakePayment          *bool
	CreatedAt            *uint32
	DeletedAt            *uint32
}

//nolint:gocyclo
func CreateSet(c *ent.PaymentCreate, req *Req) *ent.PaymentCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.StartAmount != nil {
		c.SetStartAmount(*req.StartAmount)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.PayWithBalanceAmount != nil {
		c.SetPayWithBalanceAmount(*req.PayWithBalanceAmount)
	}
	if req.FinishAmount != nil {
		c.SetFinishAmount(*req.FinishAmount)
	}
	if req.CoinUsdCurrency != nil {
		c.SetCoinUsdCurrency(*req.CoinUsdCurrency)
	}
	if req.LocalCoinUsdCurrency != nil {
		c.SetLocalCoinUsdCurrency(*req.LocalCoinUsdCurrency)
	}
	if req.LiveCoinUsdCurrency != nil {
		c.SetLiveCoinUsdCurrency(*req.LiveCoinUsdCurrency)
	}
	if req.CoinInfoID != nil {
		c.SetCoinInfoID(*req.CoinInfoID)
	}
	if req.State != nil {
		c.SetState(req.State.String())
	}
	if req.ChainTransactionID != nil {
		c.SetChainTransactionID(*req.ChainTransactionID)
	}
	if req.UserSetPaid != nil {
		c.SetUserSetPaid(*req.UserSetPaid)
	}
	if req.UserSetCanceled != nil {
		c.SetUserSetCanceled(*req.UserSetCanceled)
	}
	if req.FakePayment != nil {
		c.SetFakePayment(*req.FakePayment)
	}
	if req.CreatedAt != nil {
		c.SetCreatedAt(*req.CreatedAt)
	}
	return c
}

func UpdateSet(u *ent.PaymentUpdateOne, req *Req) *ent.PaymentUpdateOne {
	if req.FinishAmount != nil {
		u.SetFinishAmount(*req.FinishAmount)
	}
	if req.State != nil {
		u.SetState(req.State.String())
	}
	if req.UserSetPaid != nil {
		u.SetUserSetPaid(*req.UserSetPaid)
	}
	if req.UserSetCanceled != nil {
		u.SetUserSetCanceled(*req.UserSetCanceled)
	}
	if req.FakePayment != nil {
		u.SetFakePayment(*req.FakePayment)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	GoodID     *cruder.Cond
	OrderID    *cruder.Cond
	AccountID  *cruder.Cond
	CoinInfoID *cruder.Cond
	State      *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.PaymentQuery, conds *Conds) (*ent.PaymentQuery, error) {
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
			q.Where(entpayment.ID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
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
				q.Where(entpayment.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid payment field")
			}
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entpayment.AppID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entpayment.UserID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entpayment.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entpayment.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entpayment.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.CoinInfoID != nil {
		id, ok := conds.CoinInfoID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coininfoid")
		}
		switch conds.CoinInfoID.Op {
		case cruder.EQ:
			q.Where(entpayment.CoinInfoID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(basetypes.PaymentState)
		if !ok {
			return nil, fmt.Errorf("invalid state")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entpayment.State(state.String()))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	q.Where(entpayment.DeletedAt(0))
	return q, nil
}
