package appconfig

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	if h.ID == nil && h.EntID == nil && h.AppID == nil {
		return fmt.Errorf("invalid appconfigid")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update app_configs "
	if h.EnableSimulateOrder != nil {
		_sql += fmt.Sprintf("%venable_simulate_order = %v, ", set, *h.EnableSimulateOrder)
		set = ""
	}
	if h.SimulateOrderUnits != nil {
		_sql += fmt.Sprintf("%vsimulate_order_units = %v, ", set, *h.SimulateOrderUnits)
		set = ""
	}
	if h.SimulateOrderDurationSeconds != nil {
		_sql += fmt.Sprintf("%vsimulate_order_duration_seconds = %v, ", set, *h.SimulateOrderDurationSeconds)
		set = ""
	}
	if h.SimulateOrderCouponMode != nil {
		_sql += fmt.Sprintf("%vsimulate_order_coupon_mode = '%v', ", set, h.SimulateOrderCouponMode.String())
		set = ""
	}
	if h.SimulateOrderCouponProbability != nil {
		_sql += fmt.Sprintf(
			"%vsimulate_order_coupon_probability = '%v', ",
			set,
			*h.SimulateOrderCouponProbability,
		)
		set = ""
	}
	if h.SimulateOrderCashableProfitProbability != nil {
		_sql += fmt.Sprintf(
			"%vsimulate_order_cashable_profit_probability = %v, ",
			set,
			*h.SimulateOrderCashableProfitProbability,
		)
		set = ""
	}
	if h.MaxUnpaidOrders != nil {
		_sql += fmt.Sprintf("%vmax_unpaid_orders = %v, ", set, *h.MaxUnpaidOrders)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
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
	if h.AppID != nil {
		_sql += fmt.Sprintf("%v app_id = '%v'", whereAnd, *h.AppID)
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateAppConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("fail update appconfig: %v", err)
	}
	return nil
}

func (h *Handler) UpdateAppConfig(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.constructSQL(); err != nil {
		return err
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateAppConfig(ctx, tx)
	})
}
