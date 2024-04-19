package feeorder

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entfeeorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorder"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	// entordercoupon "github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderBaseSelect
}

func (h *baseQueryHandler) selectOrderBase(stm *ent.OrderBaseQuery) *ent.OrderBaseSelect {
	return stm.Select(entorderbase.FieldID)
}

func (h *baseQueryHandler) queryOrderBase(cli *ent.Client) error {
	if h.OrderID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.OrderBase.Query().Where(entorderbase.DeletedAt(0))
	if h.OrderID != nil {
		stm.Where(entorderbase.EntID(*h.OrderID))
	}
	h.stmSelect = h.selectOrderBase(stm)
	return nil
}

func (h *baseQueryHandler) queryOrderBases(cli *ent.Client) (*ent.OrderBaseSelect, error) {
	stm, err := orderbasecrud.SetQueryConds(cli.OrderBase.Query(), h.OrderBaseConds)
	if err != nil {
		return nil, err
	}
	return h.selectOrderBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldID),
			t.C(entorderbase.FieldID),
		)
	if h.OrderBaseConds.AppID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldAppID),
				h.OrderBaseConds.AppID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.UserID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldUserID),
				h.OrderBaseConds.UserID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.GoodID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldGoodID),
				h.OrderBaseConds.GoodID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.AppGoodID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldAppGoodID),
				h.OrderBaseConds.AppGoodID.Val.(uuid.UUID),
			),
		)
	}
	s.AppendSelect(
		t.C(entorderbase.FieldAppID),
		t.C(entorderbase.FieldUserID),
		t.C(entorderbase.FieldGoodID),
		t.C(entorderbase.FieldAppGoodID),
		t.C(entorderbase.FieldGoodType),
		t.C(entorderbase.FieldParentOrderID),
		t.C(entorderbase.FieldOrderType),
		t.C(entorderbase.FieldCreateMethod),
		t.C(entorderbase.FieldCreatedAt),
		t.C(entorderbase.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoinFeeOrder(s *sql.Selector) error { //nolint
	t := sql.Table(entfeeorder.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entfeeorder.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entfeeorder.FieldDeletedAt), 0),
		)
	s.AppendSelect(
		t.C(entfeeorder.FieldOrderID),
		t.C(entfeeorder.FieldGoodValueUsd),
		t.C(entfeeorder.FieldPaymentAmountUsd),
		t.C(entfeeorder.FieldDiscountAmountUsd),
		t.C(entfeeorder.FieldPromotionID),
		t.C(entfeeorder.FieldDurationSeconds),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinFeeOrder(s)
	})
}
