package order

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"

	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) paymentState() *ordertypes.PaymentState {
	if h.PaymentTransferAmount != nil && h.PaymentTransferAmount.Cmp(decimal.NewFromInt(0)) > 0 {
		return ordertypes.PaymentState_PaymentStateWait.Enum()
	}
	return ordertypes.PaymentState_PaymentStateNoPayment.Enum()
}

func (h *createHandler) createOrderState(ctx context.Context, tx *ent.Tx, req *orderstatecrud.Req) error {
	if _, err := orderstatecrud.CreateSet(
		tx.OrderState.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) createPayment(ctx context.Context, tx *ent.Tx, req *paymentcrud.Req) error {
	if _, err := paymentcrud.CreateSet(
		tx.Payment.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) createOrder(ctx context.Context, tx *ent.Tx, req *ordercrud.Req) error {
	if _, err := ordercrud.CreateSet(
		tx.Order.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CreateOrder(ctx context.Context) (*npool.Order, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	h.PaymentState = handler.paymentState()
	req := h.ToOrderReq()

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.createOrder(ctx, tx, req.Req); err != nil {
			return err
		}
		id := uuid.New()
		req.OrderStateReq.ID = &id
		if err := handler.createOrderState(ctx, tx, req.OrderStateReq); err != nil {
			return err
		}
		if req.PaymentReq == nil {
			return nil
		}
		req.PaymentReq.OrderID = req.Req.ID
		if err := handler.createPayment(ctx, tx, req.PaymentReq); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOrder(ctx)
}

func (h *Handler) CreateOrders(ctx context.Context) ([]*npool.Order, uint32, error) {
	handler := &createHandler{
		Handler: h,
	}
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.Req.ID == nil {
				req.Req.ID = &id
				req.OrderStateReq.OrderID = &id
			}
			if err := handler.createOrder(ctx, tx, req.Req); err != nil {
				return err
			}

			id = uuid.New()
			req.OrderStateReq.ID = &id
			if err := handler.createOrderState(ctx, tx, req.OrderStateReq); err != nil {
				return err
			}
			ids = append(ids, *req.Req.ID)

			if req.PaymentReq == nil {
				continue
			}
			req.PaymentReq.OrderID = req.Req.ID
			if err := handler.createPayment(ctx, tx, req.PaymentReq); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	h.Conds = &ordercrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	return h.GetOrders(ctx)
}
