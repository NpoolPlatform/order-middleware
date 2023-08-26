package order

import (
	"context"
	"fmt"

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

func (h *createHandler) validate(req *paymentcrud.Req) error {
	if req.BalanceAmount != nil && req.BalanceAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("balanceamount is less than or equal to 0")
	}
	if req.StartAmount == nil {
		return fmt.Errorf("invalid startamount")
	}
	if req.StartAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("startamount is less than or equal to 0")
	}
	if req.CoinUSDCurrency == nil {
		return fmt.Errorf("invalid coinusdcurrency")
	}
	if req.CoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("coinusdcurrency is less than or equal to 0")
	}
	if req.LiveCoinUSDCurrency == nil {
		return fmt.Errorf("invalid livecoinusdcurrency")
	}
	if req.LiveCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
	}
	if req.LocalCoinUSDCurrency == nil {
		return fmt.Errorf("invalid localcoinusdcurrency")
	}
	if req.LocalCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
	}

	return nil
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

func (h *Handler) CreateOrders(ctx context.Context) ([]*npool.Order, uint32, error) {
	handler := &createHandler{
		Handler: h,
	}
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.OrderReq.ID == nil {
				req.OrderReq.ID = &id
				req.PaymentReq.OrderID = &id
				req.OrderStateReq.OrderID = &id
			}

			paymentState := ordertypes.PaymentState_PaymentStateNoPayment
			if req.PaymentReq.TransferAmount != nil && req.PaymentReq.TransferAmount.Cmp(decimal.NewFromInt(0)) > 0 {
				paymentState = ordertypes.PaymentState_PaymentStateWait
				id = uuid.New()
				req.PaymentReq.ID = &id
				req.OrderReq.PaymentID = &id
			}
			req.OrderStateReq.PaymentState = &paymentState

			if err := handler.createOrder(ctx, tx, req.OrderReq); err != nil {
				return err
			}

			if err := handler.createPayment(ctx, tx, req.PaymentReq); err != nil {
				return err
			}

			if err := handler.createOrderState(ctx, tx, req.OrderStateReq); err != nil {
				return err
			}
			ids = append(ids, *req.OrderReq.ID)
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
