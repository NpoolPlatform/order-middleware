package feeorder

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type MultiHandler struct {
	Handlers []*Handler
}

func (h *MultiHandler) AppendHandler(handler *Handler) {
	h.Handlers = append(h.Handlers, handler)
}

func (h *MultiHandler) GetHandlers() []*Handler {
	return h.Handlers
}

func (h *MultiHandler) validatePaymentOrder() (bool, error) {
	paymentOrders := 0
	payWithParents := 0
	payWithOthers := 0
	offlineOrder := false

	for _, handler := range h.Handlers {
		switch *handler.OrderBaseReq.OrderType {
		case types.OrderType_Offline:
			fallthrough //nolint
		case types.OrderType_Airdrop:
			offlineOrder = true
			continue
		}
		if offlineOrder {
			return false, wlog.Errorf("invalid ordertype")
		}
		if len(handler.PaymentTransferReqs) > 0 || len(handler.PaymentBalanceReqs) > 0 {
			paymentOrders += 1
		}
		if handler.OrderStateBaseReq.PaymentType == nil {
			continue
		}
		switch *handler.OrderStateBaseReq.PaymentType {
		case types.PaymentType_PayWithParentOrder:
			payWithParents += 1
		case types.PaymentType_PayWithOtherOrder:
			payWithOthers += 1
		}
	}
	switch paymentOrders {
	case 0:
		if !offlineOrder && payWithParents != len(h.Handlers) {
			return false, wlog.Errorf("invalid paywithparents")
		}
	case 1:
		if payWithOthers != len(h.Handlers)-1 {
			return false, wlog.Errorf("invalid paywithothers")
		}
	default:
		return false, wlog.Errorf("invalid paymentorder")
	}
	return paymentOrders > 0, nil
}

//nolint:gocyclo
func (h *MultiHandler) validatePaymentID() error {
	paymentIDs := map[uuid.UUID]struct{}{}

	for _, handler := range h.Handlers {
		if len(paymentIDs) > 0 && handler.PaymentBaseReq.EntID == nil {
			return wlog.Errorf("invalid paymentid")
		}
		if handler.PaymentBaseReq.EntID != nil {
			paymentIDs[*handler.PaymentBaseReq.EntID] = struct{}{}
		}
		for _, balance := range handler.PaymentBalanceReqs {
			if len(paymentIDs) > 0 && balance.PaymentID == nil {
				return wlog.Errorf("invalid paymentid")
			}
			if balance.PaymentID != nil {
				paymentIDs[*balance.PaymentID] = struct{}{}
			}
		}
		for _, transfer := range handler.PaymentTransferReqs {
			if len(paymentIDs) > 0 && transfer.PaymentID == nil {
				return wlog.Errorf("invalid paymentid")
			}
			if transfer.PaymentID != nil {
				paymentIDs[*transfer.PaymentID] = struct{}{}
			}
		}
	}
	if len(paymentIDs) > 1 {
		return wlog.Errorf("invalid paymentid")
	}
	if len(paymentIDs) == 1 {
		return nil
	}

	paymentID := uuid.New()
	for _, handler := range h.Handlers {
		handler.PaymentBaseReq.EntID = &paymentID
		for _, balance := range handler.PaymentBalanceReqs {
			balance.PaymentID = &paymentID
		}
		for _, transfer := range handler.PaymentTransferReqs {
			transfer.PaymentID = &paymentID
		}
	}

	return nil
}

func (h *MultiHandler) CreateFeeOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	shouldPay, err := h.validatePaymentOrder()
	if err != nil {
		return wlog.WrapError(err)
	}
	if shouldPay {
		if err := h.validatePaymentID(); err != nil {
			return wlog.WrapError(err)
		}
	}
	for _, handler := range h.Handlers {
		if err := handler.CreateFeeOrderWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) CreateFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateFeeOrdersWithTx(_ctx, tx)
	})
}

func (h *MultiHandler) UpdateFeeOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.UpdateFeeOrderWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) UpdateFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.UpdateFeeOrdersWithTx(_ctx, tx)
	})
}

func (h *MultiHandler) DeleteFeeOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.DeleteFeeOrderWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) DeleteFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteFeeOrdersWithTx(_ctx, tx)
	})
}
