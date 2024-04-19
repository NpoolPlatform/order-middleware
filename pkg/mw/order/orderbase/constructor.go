package orderbase

import (
	"fmt"
	"time"
)

//nolint:goconst,funlen
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into order_bases "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "good_id"
	_sql += comma + "app_good_id"
	if h.ParentOrderID != nil {
		_sql += comma + "parent_order_id"
	}
	_sql += comma + "order_type"
	_sql += comma + "payment_type"
	_sql += comma + "create_method"
	if h.Simulate != nil {
		_sql += comma + "simulate"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	if h.ParentOrderID != nil {
		_sql += fmt.Sprintf("%v'%v' as parent_order_id", comma, *h.ParentOrderID)
	}
	_sql += fmt.Sprintf("%v'%v' as order_type", comma, h.OrderType.String())
	_sql += fmt.Sprintf("%v'%v' as payment_type", comma, h.PaymentType.String())
	_sql += fmt.Sprintf("%v'%v' as create_method", comma, h.CreateMethod.String())
	if h.Simulate != nil {
		_sql += fmt.Sprintf("%v%v as simulate", comma, *h.Simulate)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	if h.ParentOrderID != nil {
		_sql += "where exists ("
		_sql += "select 1 from order_bases "
		_sql += fmt.Sprintf("where ent_id = '%v' ", *h.ParentOrderID)
		_sql += "limit 1)"
	}

	return _sql
}
