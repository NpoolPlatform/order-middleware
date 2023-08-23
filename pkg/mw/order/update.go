package order

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateOrder(ctx context.Context, tx *ent.Tx, req *ordercrud.Req) error {
	payment, err := tx.Payment.
		Query().
		Where(
			entpayment.ID(*req.PaymentID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if payment.OrderID != *req.ID {
		return fmt.Errorf("invalid order")
	}

	order, err := tx.Order.
		Query().
		Where(
			entorder.ID(*req.ID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if order == nil {
		return fmt.Errorf("invalid order")
	}

	duration := order.EndAt - order.StartAt
	startAt := order.StartAt
	if req.StartAt != nil && *req.StartAt > startAt {
		startAt = *req.StartAt
	}
	endAt := startAt + duration

	if payment.StateV1 != basetypes.PaymentState_PaymentStateWait.String() && order.Type == basetypes.OrderType_Normal.String() {
		if req.PaymentUserSetCanceled != nil && *req.PaymentUserSetCanceled {
			return fmt.Errorf("not wait payment")
		}
	}

	if _, err := ordercrud.UpdateSet(
		order.Update(),
		&ordercrud.Req{
			State:         req.State,
			StartAt:       &startAt,
			EndAt:         &endAt,
			LastBenefitAt: req.LastBenefitAt,
		},
	).Save(ctx); err != nil {
		return err
	}

	if _, err := paymentcrud.UpdateSet(
		payment.Update(),
		&paymentcrud.Req{
			UserSetCanceled: req.PaymentUserSetCanceled,
			State:           req.PaymentState,
			FinishAmount:    req.PaymentFinishAmount,
			FakePayment:     req.PaymentFakePayment,
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
	if h.PaymentID == nil {
		return nil, fmt.Errorf("invalid payment")
	}
	handler := &updateHandler{
		Handler: h,
	}
	req := &ordercrud.Req{
		ID:                     handler.ID,
		AppID:                  handler.AppID,
		PaymentID:              handler.PaymentID,
		StartAt:                handler.StartAt,
		State:                  handler.State,
		LastBenefitAt:          handler.LastBenefitAt,
		PaymentUserSetCanceled: handler.PaymentUserSetCanceled,
		PaymentState:           handler.PaymentState,
		PaymentFinishAmount:    handler.PaymentFinishAmount,
		PaymentFakePayment:     handler.PaymentFakePayment,
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
			if req.PaymentID == nil {
				return fmt.Errorf("invalid payment")
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
