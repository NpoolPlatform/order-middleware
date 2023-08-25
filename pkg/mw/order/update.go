package order

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateOrder(ctx context.Context, tx *ent.Tx, req *ordercrud.Req) error {
	orderstate, err := tx.OrderState.
		Query().
		Where(
			entorderstate.OrderID(*req.ID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if orderstate == nil {
		return fmt.Errorf("invalid order")
	}

	order, err := tx.Order.
		Query().
		Where(
			entorder.ID(*req.ID),
		).
		Only(ctx)
	if err != nil {
		return err
	}
	if order == nil {
		return fmt.Errorf("invalid order")
	}

	duration := orderstate.EndAt - orderstate.StartAt
	startAt := orderstate.StartAt
	if req.StartAt != nil && *req.StartAt > startAt {
		startAt = *req.StartAt
	}
	endAt := startAt + duration

	if orderstate.PaymentState != basetypes.PaymentState_PaymentStateWait.String() && order.OrderType == basetypes.OrderType_Normal.String() {
		if req.UserSetCanceled != nil && *req.UserSetCanceled {
			return fmt.Errorf("not wait payment")
		}
	}

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			OrderState:           req.OrderState,
			StartMode:            req.StartMode,
			StartAt:              &startAt,
			EndAt:                &endAt,
			LastBenefitAt:        req.LastBenefitAt,
			BenefitState:         req.BenefitState,
			UserSetPaid:          req.UserSetPaid,
			UserSetCanceled:      req.UserSetCanceled,
			PaymentTransactionID: req.PaymentTransactionID,
			PaymentFinishAmount:  req.PaymentFinishAmount,
			PaymentState:         req.PaymentState,
			OutOfGasHours:        req.OutOfGasHours,
			CompensateHours:      req.CompensateHours,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) UpdateOrder(ctx context.Context) (*npool.Order, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	handler := &updateHandler{
		Handler: h,
	}
	req := &ordercrud.Req{
		ID:                   h.ID,
		AppID:                h.AppID,
		OrderState:           h.OrderState,
		StartMode:            h.StartMode,
		StartAt:              h.StartAt,
		LastBenefitAt:        h.LastBenefitAt,
		BenefitState:         h.BenefitState,
		UserSetPaid:          h.UserSetPaid,
		UserSetCanceled:      h.UserSetCanceled,
		PaymentTransactionID: h.PaymentTransactionID,
		PaymentFinishAmount:  h.PaymentFinishAmount,
		PaymentState:         h.PaymentState,
		OutOfGasHours:        h.OutOfGasHours,
		CompensateHours:      h.CompensateHours,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateOrder(_ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOrder(ctx)
}

func (h *Handler) UpdateOrders(ctx context.Context) ([]*npool.Order, error) {
	handler := &updateHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.ID == nil {
				return fmt.Errorf("invalid id")
			}
			if err := handler.updateOrder(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &ordercrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
