package powerrental

import (
	"fmt"
	"time"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into power_rentals "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "order_id"
	comma = ", "
	if h.AppGoodStockID != nil {
		_sql += comma + "app_good_stock_id"
	}
	_sql += comma + "units"
	_sql += comma + "good_value_usd"
	_sql += comma + "payment_amount_usd"
	_sql += comma + "discount_amount_usd"
	if h.PromotionID != nil {
		_sql += comma + "promotion_id"
	}
	if h.InvestmentType != nil {
		_sql += comma + "investment_type"
	}
	_sql += comma + "duration_seconds"
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
	if h.AppGoodStockID != nil {
		_sql += fmt.Sprintf("%v'%v' as app_good_stock_id", comma, *h.AppGoodStockID)
	}
	_sql += fmt.Sprintf("%v'%v' as units", comma, *h.Units)
	_sql += fmt.Sprintf("%v'%v' as good_value_usd", comma, *h.GoodValueUSD)
	if h.PaymentAmountUSD != nil {
		_sql += fmt.Sprintf("%v'%v' as payment_amount_usd", comma, *h.PaymentAmountUSD)
	} else {
		_sql += fmt.Sprintf("%v'0' as payment_amount_usd", comma)
	}
	if h.DiscountAmountUSD != nil {
		_sql += fmt.Sprintf("%v'%v' as discount_amount_usd", comma, *h.DiscountAmountUSD)
	} else {
		_sql += fmt.Sprintf("%v'0' as discount_amount_usd", comma)
	}
	if h.PromotionID != nil {
		_sql += fmt.Sprintf("%v'%v' as promotion_id", comma, *h.PromotionID)
	}
	if h.InvestmentType != nil {
		_sql += fmt.Sprintf("%v'%v' as investment_type", comma, h.InvestmentType.String())
	}
	_sql += fmt.Sprintf("%v'%v' as duration_seconds", comma, *h.DurationSeconds)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp"

	return _sql
}
