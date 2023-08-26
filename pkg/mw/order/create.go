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

func (h *createHandler) paymentState() ordertypes.PaymentState {
	if h.PaymentTransferAmount != nil && h.PaymentTransferAmount.Cmp(decimal.NewFromInt(0)) > 0 {
		return ordertypes.PaymentState_PaymentStateWait
	}
	return ordertypes.PaymentState_PaymentStateNoPayment
}

func (h *createHandler) initCreateReq(req OrderReq) *OrderReq {
	id := uuid.New()
	if req.Req.ID == nil {
		req.Req.ID = &id
		req.OrderStateReq.OrderID = &id
	}

	paymentState := h.paymentState()
	if req.OrderStateReq.PaymentState != nil {
		paymentState = *h.PaymentState
	}
	req.OrderStateReq.PaymentState = &paymentState
	return &req
}

func (h *createHandler) createOrderState(ctx context.Context, tx *ent.Tx, req *orderstatecrud.Req) error {
	orderState := ordertypes.OrderState_OrderStateWaitPayment
	id := uuid.New()
	req.ID = &id
	req.OrderState = &orderState
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
	req := h.ToOrderReq()

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		_req := handler.initCreateReq(*req)
		h.ID = _req.Req.ID

		if err := handler.createOrder(ctx, tx, _req.Req); err != nil {
			return err
		}
		if err := handler.createOrderState(ctx, tx, _req.OrderStateReq); err != nil {
			return err
		}
		if _req.PaymentReq == nil {
			return nil
		}
		_req.PaymentReq.OrderID = _req.Req.ID
		if err := handler.createPayment(ctx, tx, _req.PaymentReq); err != nil {
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
	reqs := h.ToOrderReqs()
	handler := &createHandler{
		Handler: h,
	}
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range reqs {
			_req := handler.initCreateReq(*req)

			if err := handler.createOrder(ctx, tx, _req.Req); err != nil {
				return err
			}
			if err := handler.createOrderState(ctx, tx, _req.OrderStateReq); err != nil {
				return err
			}
			ids = append(ids, *_req.Req.ID)

			if _req.PaymentReq == nil {
				continue
			}
			_req.PaymentReq.OrderID = _req.Req.ID
			if err := handler.createPayment(ctx, tx, _req.PaymentReq); err != nil {
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
