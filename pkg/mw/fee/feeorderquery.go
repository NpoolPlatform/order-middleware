package feeorder

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entfeeorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorder"
	entfeeorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorderstate"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entordercoupon "github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	entpaymentbalance "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbalance"
	entpaymentbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbase"
	entpaymenttransfer "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymenttransfer"
)

type feeOrderQueryHandler struct {
	*Handler
	feeOrder
}

func (h *feeOrderQueryHandler) getFeeOrderEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.FeeOrder.Query()
	if h.ID != nil {
		stm.Where(entfeeorder.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entfeeorder.EntID(*h.EntID))
	}
	if h.OrderID != nil {
		stm.Where(entfeeorder.OrderID(*h.OrderID))
	}
	if h.entFeeOrder, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return err
	}
	return nil
}

func (h *feeOrderQueryHandler) getFeeOrderState(ctx context.Context, cli *ent.Client) (err error) {
	h.entFeeOrderState, err = cli.
		FeeOrderState.
		Query().
		Where(
			entfeeorderstate.OrderID(h.entFeeOrder.OrderID),
			entfeeorderstate.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *feeOrderQueryHandler) getOrderBase(ctx context.Context, cli *ent.Client) (err error) {
	h.entOrderBase, err = cli.
		OrderBase.
		Query().
		Where(
			entorderbase.EntID(h.entFeeOrder.OrderID),
			entorderbase.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *feeOrderQueryHandler) getOrderStateBase(ctx context.Context, cli *ent.Client) (err error) {
	h.entOrderStateBase, err = cli.
		OrderStateBase.
		Query().
		Where(
			entorderstatebase.OrderID(h.entFeeOrder.OrderID),
			entorderstatebase.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *feeOrderQueryHandler) getPaymentBase(ctx context.Context, cli *ent.Client) (err error) {
	h.entPaymentBase, err = cli.
		PaymentBase.
		Query().
		Where(
			entpaymentbase.OrderID(h.entFeeOrder.OrderID),
			entpaymentbase.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *feeOrderQueryHandler) getLedgerLock(ctx context.Context, cli *ent.Client) (err error) {
	h.entLedgerLock, err = cli.
		OrderLock.
		Query().
		Where(
			entorderlock.OrderID(h.entFeeOrder.OrderID),
			entorderlock.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *feeOrderQueryHandler) getPaymentBalances(ctx context.Context, cli *ent.Client) (err error) {
	h.entPaymentBalances, err = cli.
		PaymentBalance.
		Query().
		Where(
			entpaymentbalance.PaymentID(h.entPaymentBase.EntID),
			entpaymentbalance.DeletedAt(0),
		).
		All(ctx)
	return err
}

func (h *feeOrderQueryHandler) getPaymentTransfers(ctx context.Context, cli *ent.Client) (err error) {
	h.entPaymentTransfers, err = cli.
		PaymentTransfer.
		Query().
		Where(
			entpaymenttransfer.PaymentID(h.entPaymentBase.EntID),
			entpaymenttransfer.DeletedAt(0),
		).
		All(ctx)
	return err
}

func (h *feeOrderQueryHandler) getOrderCoupons(ctx context.Context, cli *ent.Client) (err error) {
	h.entOrderCoupons, err = cli.
		OrderCoupon.
		Query().
		Where(
			entordercoupon.OrderID(h.entFeeOrder.OrderID),
			entordercoupon.DeletedAt(0),
		).
		All(ctx)
	return err
}

func (h *feeOrderQueryHandler) _getFeeOrder(ctx context.Context, must bool) error {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return fmt.Errorf("invalid id")
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getFeeOrderEnt(_ctx, cli, must); err != nil {
			return err
		}
		if h.entFeeOrder == nil {
			return nil
		}
		if err := h.getFeeOrderState(_ctx, cli); err != nil {
			return err
		}
		if err := h.getOrderBase(_ctx, cli); err != nil {
			return err
		}
		if err := h.getOrderStateBase(_ctx, cli); err != nil {
			return err
		}
		if err := h.getPaymentBase(_ctx, cli); err != nil {
			return err
		}
		if err := h.getLedgerLock(_ctx, cli); err != nil {
			return err
		}
		if err := h.getPaymentBalances(_ctx, cli); err != nil {
			return err
		}
		if err := h.getPaymentTransfers(_ctx, cli); err != nil {
			return err
		}
		return h.getOrderCoupons(_ctx, cli)
	})
}

func (h *feeOrderQueryHandler) getFeeOrder(ctx context.Context) error {
	return nil
}

func (h *feeOrderQueryHandler) requireFeeOrder(ctx context.Context) error {
	return nil
}
