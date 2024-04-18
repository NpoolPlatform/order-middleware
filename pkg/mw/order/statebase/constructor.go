package statebase

import (
	"fmt"
	"time"
)

type constructorHandler struct {
	*Handler
}

func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into order_state_bases "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "order_id"
	comma = ", "
	_sql += comma + "order_state"
	_sql += comma + "start_mode"
	if h.StartAt != nil {
		_sql += comma + "start_at"
	}
	_sql += comma + "last_benefit_at"
	_sql += comma + "benefit_state"
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
	_sql += fmt.Sprintf("%v'%v' as order_id", comma, *h.OrderID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as order_state", comma, h.OrderState.String())
	_sql += fmt.Sprintf("%v'%v' as start_mode", comma, h.StartMode.String())
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%v%v as start_at", comma, *h.StartAt)
	}
	_sql += fmt.Sprintf("%v%v as last_benefit_at", comma, *h.LastBenefitAt)
	_sql += fmt.Sprintf("%v'%v' as benefit_state", comma, h.BenefitState.String())
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from order_state_bases "
	_sql += fmt.Sprintf("where order_id = '%v' ", *h.OrderID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from order_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.OrderID)
	_sql += "limit 1)"

	return _sql
}
