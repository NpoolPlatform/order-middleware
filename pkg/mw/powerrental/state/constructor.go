package powerrentalstate

import (
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type constructorHandler struct {
	*Handler
}

func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into power_rental_states "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "order_id"
	comma = ", "
	_sql += comma + "cancel_state"
	_sql += comma + "end_at"
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
	_sql += fmt.Sprintf("%v'%v' as cancel_state", comma, h.CancelState.String())
	_sql += fmt.Sprintf("%v'%v' as end_at", comma, *h.EndAt)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from power_rental_states "
	_sql += fmt.Sprintf("where order_id = '%v' ", *h.OrderID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from order_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.OrderID)
	_sql += "limit 1)"

	return _sql
}

func (h *Handler) ConstructUpdateSQL() (string, error) {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return "", fmt.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update power_rental_states "
	if h.CancelState != nil {
		_sql += fmt.Sprintf("%vcancel_state = '%v', ", set, h.CancelState.String())
		set = ""
	}
	if h.EndAt != nil {
		_sql += fmt.Sprintf("%vend_at = %v, ", set, *h.EndAt)
		set = ""
	}
	if h.PaidAt != nil {
		_sql += fmt.Sprintf("%vpaid_at = %v, ", set, *h.PaidAt)
		set = ""
	}
	if h.UserSetPaid != nil {
		_sql += fmt.Sprintf("%vuser_set_paid = %v, ", set, *h.UserSetPaid)
		set = ""
	}
	if h.UserSetCanceled != nil {
		_sql += fmt.Sprintf("%vuser_set_canceled = %v, ", set, *h.UserSetCanceled)
		set = ""
	}
	if h.AdminSetCanceled != nil {
		_sql += fmt.Sprintf("%vadmin_set_canceled = %v, ", set, *h.AdminSetCanceled)
		set = ""
	}
	if h.PaymentState != nil {
		_sql += fmt.Sprintf("%vpayment_state = '%v', ", set, h.PaymentState.String())
		set = ""
	}
	if h.OutOfGasHours != nil {
		_sql += fmt.Sprintf("%vout_of_gas_hours = %v, ", set, *h.OutOfGasHours)
		set = ""
	}
	if h.CompensateHours != nil {
		_sql += fmt.Sprintf("%vcompensate_hours = %v, ", set, *h.CompensateHours)
		set = ""
	}
	if h.RenewState != nil {
		_sql += fmt.Sprintf("%vrenew_state = '%v', ", set, h.RenewState.String())
		set = ""
	}
	if h.RenewNotifyAt != nil {
		_sql += fmt.Sprintf("%vrenew_notify_at = %v, ", set, *h.RenewNotifyAt)
		set = ""
	}
	if set != "" {
		return "", cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	whereAnd := "where "
	if h.ID != nil {
		_sql += fmt.Sprintf("where id = %v ", *h.ID)
		whereAnd = "and"
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%v ent_id = '%v'", whereAnd, *h.EntID)
		whereAnd = "and"
	}
	if h.OrderID != nil {
		_sql += fmt.Sprintf("%v order_id = '%v'", whereAnd, *h.OrderID)
	}

	return _sql, nil
}
