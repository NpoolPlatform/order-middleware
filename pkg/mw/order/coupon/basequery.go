package ordercoupon

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	ordercouponcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/coupon"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entordercoupon "github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderCouponSelect
}

func (h *baseQueryHandler) selectOrderCoupon(stm *ent.OrderCouponQuery) *ent.OrderCouponSelect {
	return stm.Select(entordercoupon.FieldID)
}

func (h *baseQueryHandler) queryOrderCoupon(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.CouponID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.OrderCoupon.Query().Where(entordercoupon.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entordercoupon.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entordercoupon.EntID(*h.EntID))
	}
	if h.CouponID != nil {
		stm.Where(entordercoupon.CouponID(*h.CouponID))
	}
	h.stmSelect = h.selectOrderCoupon(stm)
	return nil
}

func (h *baseQueryHandler) queryOrderCoupons(cli *ent.Client) (*ent.OrderCouponSelect, error) {
	stm, err := ordercouponcrud.SetQueryConds(cli.OrderCoupon.Query(), h.OrderCouponConds)
	if err != nil {
		return nil, err
	}
	return h.selectOrderCoupon(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entordercoupon.Table)
	s.AppendSelect(
		t.C(entordercoupon.FieldID),
		t.C(entordercoupon.FieldEntID),
		t.C(entordercoupon.FieldOrderID),
		t.C(entordercoupon.FieldCouponID),
		t.C(entordercoupon.FieldCreatedAt),
		t.C(entordercoupon.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoinOrder(s *sql.Selector) error { //nolint
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entordercoupon.FieldOrderID),
			t.C(entorderbase.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entorderbase.FieldDeletedAt), 0),
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
	s.AppendSelect(
		t.C(entorderbase.FieldAppID),
		t.C(entorderbase.FieldUserID),
		t.C(entorderbase.FieldGoodID),
		t.C(entorderbase.FieldAppGoodID),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinOrder(s)
	})
}
