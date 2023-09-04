package order

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) paymentState(req *ordercrud.Req) *types.PaymentState {
	if req.TransferAmount != nil && req.TransferAmount.Cmp(decimal.NewFromInt(0)) > 0 {
		return types.PaymentState_PaymentStateWait.Enum()
	}
	return types.PaymentState_PaymentStateNoPayment.Enum()
}

func (h *createHandler) createOrderState(ctx context.Context, tx *ent.Tx, req *orderstatecrud.Req) error {
	if *req.EndAt <= *req.StartAt {
		return fmt.Errorf("invalid startend")
	}

	if _, err := orderstatecrud.CreateSet(
		tx.OrderState.Create(),
		&orderstatecrud.Req{
			OrderID:      req.OrderID,
			OrderState:   req.OrderState,
			StartMode:    req.StartMode,
			StartAt:      req.StartAt,
			EndAt:        req.EndAt,
			PaymentState: req.PaymentState,
		},
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
	switch *req.PaymentType {
	case types.PaymentType_PayWithBalanceOnly:
		// Before we should already set transfer amount to 0
		fallthrough //nolint
	case types.PaymentType_PayWithTransferAndBalance:
		fallthrough //nolint
	case types.PaymentType_PayWithTransferOnly:
		// Before we should already set balance amount to 0
		if req.TransferAmount.Add(*req.BalanceAmount).Cmp(*req.PaymentAmount) != 0 {
			return fmt.Errorf("invalid paymentamount")
		}
	}

	if len(req.CouponIDs) > 0 {
		stm, err := ordercrud.SetQueryConds(tx.Order.Query(), &ordercrud.Conds{
			CouponIDs: &cruder.Cond{Op: cruder.IN, Val: req.CouponIDs},
		})
		if err != nil {
			return err
		}
		exist, err := stm.Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("coupons already used")
		}
	}

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

	req, err := h.ToOrderReq(ctx, true)
	if err != nil {
		return nil, err
	}
	h.PaymentState = handler.paymentState(req.Req)
	if *req.PaymentType == types.PaymentType_PayWithParentOrder {
		return nil, fmt.Errorf("invalid paymenttype")
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.createOrder(ctx, tx, req.Req); err != nil {
			return err
		}
		if err := handler.createOrderState(ctx, tx, req.OrderStateReq); err != nil {
			return err
		}
		if req.PaymentReq == nil {
			return nil
		}
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

func (h *createHandler) checkBatchParentOrder(ctx context.Context) error {
	var parentOrder uuid.UUID
	for _, req := range h.Reqs {
		if req.ParentOrderID == nil {
			continue
		}
		if parentOrder != uuid.Nil && *req.ParentOrderID != parentOrder {
			return fmt.Errorf("invalid parentorder")
		}
		parentOrder = *req.ParentOrderID
	}
	if parentOrder == uuid.Nil {
		return nil
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err := cli.
			Order.
			Query().
			Where(
				entorder.ID(parentOrder),
				entorder.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid parentorder")
		}
		return nil
	})
}

func (h *Handler) CreateOrders(ctx context.Context) ([]*npool.Order, error) {

	handler := &createHandler{
		Handler: h,
	}

	if err := handler.checkBatchParentOrder(ctx); err != nil {
		return nil, err
	}

	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			req.OrderStateReq.PaymentState = handler.paymentState(req.Req)
			id := uuid.New()
			if req.Req.ID == nil {
				req.Req.ID = &id
				req.OrderStateReq.OrderID = &id
			}
			if req.PaymentReq != nil {
				id1 := uuid.New()
				req.Req.PaymentID = &id1
			}
			if *req.PaymentType == types.PaymentType_PayWithParentOrder && req.ParentOrderID == nil {
				return fmt.Errorf("invalid parentorderid")
			}
			if err := handler.createOrder(ctx, tx, req.Req); err != nil {
				return err
			}
			if err := handler.createOrderState(ctx, tx, req.OrderStateReq); err != nil {
				return err
			}
			ids = append(ids, *req.Req.ID)

			if req.PaymentReq == nil {
				continue
			}
			req.PaymentReq.OrderID = req.Req.ID
			req.PaymentReq.ID = req.Req.PaymentID
			if err := handler.createPayment(ctx, tx, req.PaymentReq); err != nil {
				return err
			}
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
	return infos, nil
}
