package order

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderlock"
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
	parentOrderID uuid.UUID
	createParent  bool
}

func (h *createHandler) paymentState(req *ordercrud.Req) *types.PaymentState {
	if req.PaymentType != nil {
		switch *req.PaymentType {
		case types.PaymentType_PayWithNoPayment:
			return types.PaymentState_PaymentStateNoPayment.Enum()
		case types.PaymentType_PayWithOffline:
			return types.PaymentState_PaymentStateWait.Enum()
		}
	}
	if (req.TransferAmount != nil && req.TransferAmount.Cmp(decimal.NewFromInt(0)) > 0) ||
		(req.BalanceAmount != nil && req.BalanceAmount.Cmp(decimal.NewFromInt(0)) > 0) {
		return types.PaymentState_PaymentStateWait.Enum()
	}
	return types.PaymentState_PaymentStateNoPayment.Enum()
}

func (h *createHandler) createOrderLocks(ctx context.Context, tx *ent.Tx, stockLockReq, balanceLockReq *orderlockcrud.Req) error {
	reqs := []*ent.OrderLockCreate{}
	if stockLockReq != nil {
		if stockLockReq.EntID == nil {
			return fmt.Errorf("invalid stocklockid")
		}
		reqs = append(reqs, orderlockcrud.CreateSet(tx.OrderLock.Create(), stockLockReq))
	}
	if balanceLockReq != nil {
		if balanceLockReq.EntID == nil {
			return fmt.Errorf("invalid balancelockid")
		}
		reqs = append(reqs, orderlockcrud.CreateSet(tx.OrderLock.Create(), balanceLockReq))
	}
	if len(reqs) == 0 {
		return nil
	}
	if _, err := tx.OrderLock.CreateBulk(reqs...).Save(ctx); err != nil {
		return err
	}

	return nil
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
	if h.EntID == nil {
		h.EntID = &id
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

	start := time.Now()
	callID := uuid.New()

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		req.OrderStateReq.PaymentState = handler.paymentState(req.Req)
		if err := handler.createOrder(ctx, tx, req.Req); err != nil {
			return err
		}
		if err := handler.createOrderState(ctx, tx, req.OrderStateReq); err != nil {
			return err
		}
		if req.BalanceLockReq != nil {
			req.BalanceLockReq.OrderID = req.Req.EntID
		}
		if req.StockLockReq != nil {
			req.StockLockReq.OrderID = req.Req.EntID
		}
		if err := handler.createOrderLocks(ctx, tx, req.StockLockReq, req.BalanceLockReq); err != nil {
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
	logger.Sugar().Infow(
		"CreateOrder",
		"EntID", *h.EntID,
		"CallID", callID,
		"Elapsed", time.Since(start),
		"Error", err,
	)
	if err != nil {
		return nil, err
	}

	info, err := h.GetOrder(ctx)
	if err != nil {
		logger.Sugar().Warnw(
			"CreateOrders",
			"CallID", callID,
			"EntID", *h.EntID,
			"Error", err,
		)
	}
	return info, nil
}

func (h *createHandler) checkBatchParentOrder(ctx context.Context) error {
	for _, req := range h.Reqs {
		if req.ParentOrderID == nil {
			h.parentOrderID = *req.EntID
			h.createParent = true
			continue
		}
		if h.parentOrderID != uuid.Nil && *req.ParentOrderID != h.parentOrderID {
			return fmt.Errorf("invalid parentorder")
		}
	}
	if h.parentOrderID == uuid.Nil {
		return nil
	}
	if h.createParent {
		return nil
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err := cli.
			Order.
			Query().
			Where(
				entorder.EntID(h.parentOrderID),
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

//nolint:gocyclo
func (h *Handler) CreateOrders(ctx context.Context) ([]*npool.Order, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.checkBatchParentOrder(ctx); err != nil {
		return nil, err
	}

	start := time.Now()
	ids := []uuid.UUID{}
	callID := uuid.New()

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			req.OrderStateReq.PaymentState = handler.paymentState(req.Req)
			if req.Req.EntID == nil {
				id := uuid.New()
				req.Req.EntID = &id
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
			if req.BalanceLockReq != nil {
				req.BalanceLockReq.OrderID = req.Req.EntID
			}
			if req.StockLockReq != nil {
				req.StockLockReq.OrderID = req.Req.EntID
			}
			if err := handler.createOrderLocks(ctx, tx, req.StockLockReq, req.BalanceLockReq); err != nil {
				return err
			}
			ids = append(ids, *req.Req.EntID)

			if req.PaymentReq == nil {
				continue
			}
			req.PaymentReq.OrderID = req.Req.EntID
			req.PaymentReq.EntID = req.Req.PaymentID
			if err := handler.createPayment(ctx, tx, req.PaymentReq); err != nil {
				return err
			}
		}
		return nil
	})

	logger.Sugar().Infow(
		"CreateOrders",
		"ID", handler.parentOrderID,
		"CallID", callID,
		"CreateParent", handler.createParent,
		"Elapsed", time.Since(start),
		"Error", err,
	)
	if err != nil {
		return nil, err
	}

	h.Conds = &ordercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetOrders(ctx)
	if err != nil {
		logger.Sugar().Warnw(
			"CreateOrders",
			"CallID", callID,
			"IDs", ids,
			"Error", err,
		)
	}
	return infos, nil
}
