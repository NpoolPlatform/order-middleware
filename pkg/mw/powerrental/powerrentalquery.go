package powerrental

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entordercoupon "github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	entpaymentbalance "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbalance"
	entpaymentbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbase"
	entpaymenttransfer "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymenttransfer"
	entpowerrental "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrental"
	entpowerrentalstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrentalstate"
)

type powerRentalQueryHandler struct {
	*Handler
	_ent powerRental
}

func (h *powerRentalQueryHandler) getPowerRentalEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.PowerRental.Query()
	if h.ID != nil {
		stm.Where(entpowerrental.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entpowerrental.EntID(*h.EntID))
	}
	if h.OrderID != nil {
		stm.Where(entpowerrental.OrderID(*h.OrderID))
	}
	if h._ent.entPowerRental, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return err
	}
	return nil
}

func (h *powerRentalQueryHandler) getPowerRentalState(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPowerRentalState, err = cli.
		PowerRentalState.
		Query().
		Where(
			entpowerrentalstate.OrderID(h._ent.entPowerRental.OrderID),
			entpowerrentalstate.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *powerRentalQueryHandler) getOrderBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderBase, err = cli.
		OrderBase.
		Query().
		Where(
			entorderbase.EntID(h._ent.entPowerRental.OrderID),
			entorderbase.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *powerRentalQueryHandler) getOrderStateBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderStateBase, err = cli.
		OrderStateBase.
		Query().
		Where(
			entorderstatebase.OrderID(h._ent.entPowerRental.OrderID),
			entorderstatebase.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *powerRentalQueryHandler) getPaymentBase(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.entPaymentBase, err = cli.
		PaymentBase.
		Query().
		Where(
			entpaymentbase.OrderID(h._ent.entPowerRental.OrderID),
			entpaymentbase.EntID(h._ent.entPowerRentalState.PaymentID),
			entpaymentbase.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
	}
	return err
}

func (h *powerRentalQueryHandler) getLedgerLock(ctx context.Context, cli *ent.Client) (err error) {
	// TODO: should get ID from payment balance lock firstly
	if h._ent.entLedgerLock, err = cli.
		OrderLock.
		Query().
		Where(
			entorderlock.OrderID(h._ent.entPowerRental.OrderID),
			entorderlock.LockType(types.OrderLockType_LockBalance.String()),
			entorderlock.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err); err != nil {
			return nil
		}
	}
	return err
}

func (h *powerRentalQueryHandler) getStockLock(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.entLedgerLock, err = cli.
		OrderLock.
		Query().
		Where(
			entorderlock.OrderID(h._ent.entPowerRental.OrderID),
			entorderlock.LockType(types.OrderLockType_LockStock.String()),
			entorderlock.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err); err != nil {
			return nil
		}
	}
	return err
}

func (h *powerRentalQueryHandler) getPaymentBalances(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPaymentBalances, err = cli.
		PaymentBalance.
		Query().
		Where(
			entpaymentbalance.PaymentID(h._ent.entPowerRentalState.PaymentID),
			entpaymentbalance.DeletedAt(0),
		).
		All(ctx)
	return err
}

func (h *powerRentalQueryHandler) getPaymentTransfers(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPaymentTransfers, err = cli.
		PaymentTransfer.
		Query().
		Where(
			entpaymenttransfer.PaymentID(h._ent.entPowerRentalState.PaymentID),
			entpaymenttransfer.DeletedAt(0),
		).
		All(ctx)
	return err
}

func (h *powerRentalQueryHandler) getOrderCoupons(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderCoupons, err = cli.
		OrderCoupon.
		Query().
		Where(
			entordercoupon.OrderID(h._ent.entPowerRental.OrderID),
			entordercoupon.DeletedAt(0),
		).
		All(ctx)
	return err
}

func (h *powerRentalQueryHandler) _getPowerRental(ctx context.Context, must bool) error {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return fmt.Errorf("invalid id")
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getPowerRentalEnt(_ctx, cli, must); err != nil {
			return err
		}
		if h._ent.entPowerRental == nil {
			return nil
		}
		if err := h.getPowerRentalState(_ctx, cli); err != nil {
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
		if err := h.getStockLock(_ctx, cli); err != nil {
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

func (h *powerRentalQueryHandler) getPowerRental(ctx context.Context) error {
	return h._getPowerRental(ctx, false)
}

func (h *powerRentalQueryHandler) requirePowerRental(ctx context.Context) error {
	return h._getPowerRental(ctx, true)
}