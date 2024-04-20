package powerrental

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	paymentbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	orderlock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/lock"
	orderstatebase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/statebase"
	paymentbase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment"
	paymentbalance1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/balance"
	paymentbalancelock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/balance/lock"
	paymenttransfer1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/transfer"
	powerrentalstate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental/state"

	"github.com/google/uuid"
)

type updateHandler struct {
	*powerRentalQueryHandler

	newPayment             bool
	newPaymentBalance      bool
	obseletePaymentBaseReq *paymentbasecrud.Req
	sqlObseletePaymentBase string

	sqlOrderStateBase     string
	sqlPowerRentalState   string
	sqlPaymentBase        string
	sqlOrderLocks         []string
	sqlPaymentBalanceLock string
	sqlPaymentBalances    []string
	sqlPaymentTransfers   []string

	updateNothing bool
}

func (h *updateHandler) constructOrderStateBaseSQL(ctx context.Context) (err error) {
	handler, _ := orderstatebase1.NewHandler(ctx)
	handler.Req = *h.OrderStateBaseReq
	handler.Req.StartMode = func() *types.OrderStartMode { e := types.OrderStartMode_OrderStartInstantly; return &e }()
	if h.sqlOrderStateBase, err = handler.ConstructUpdateSQL(); err == cruder.ErrUpdateNothing {
		return nil
	}
	return err
}

func (h *updateHandler) constructPowerRentalStateSQL(ctx context.Context) (err error) {
	handler, _ := powerrentalstate1.NewHandler(ctx)
	handler.Req = *h.PowerRentalStateReq
	if h.sqlPowerRentalState, err = handler.ConstructUpdateSQL(); err == cruder.ErrUpdateNothing {
		return nil
	}
	return err
}

func (h *updateHandler) constructOrderLockSQLs(ctx context.Context) {
	for _, req := range h.OrderLockReqs {
		handler, _ := orderlock1.NewHandler(ctx)
		handler.Req = *req
		h.sqlOrderLocks = append(h.sqlOrderLocks, handler.ConstructCreateSQL())
	}
}

func (h *updateHandler) constructPaymentBalanceLockSQL(ctx context.Context) {
	if !h.newPaymentBalance {
		return
	}
	handler, _ := paymentbalancelock1.NewHandler(ctx)
	handler.Req = *h.PaymentBalanceLockReq
	h.sqlPaymentBalanceLock = handler.ConstructCreateSQL()
}

func (h *updateHandler) constructPaymentBaseSQL(ctx context.Context) {
	if !h.newPayment {
		return
	}
	handler, _ := paymentbase1.NewHandler(ctx)
	handler.Req = *h.PaymentBaseReq
	h.sqlPaymentBase = handler.ConstructCreateSQL()
}

func (h *updateHandler) constructObseletePaymentBaseSQL(ctx context.Context) (err error) {
	if !h.newPayment {
		return
	}
	handler, _ := paymentbase1.NewHandler(ctx)
	handler.Req = *h.obseletePaymentBaseReq
	if h.sqlObseletePaymentBase, err = handler.ConstructUpdateSQL(); err == cruder.ErrUpdateNothing {
		return nil
	}
	return err
}

func (h *updateHandler) constructPaymentBalanceSQLs(ctx context.Context) {
	if !h.newPaymentBalance {
		return
	}
	for _, req := range h.PaymentBalanceReqs {
		handler, _ := paymentbalance1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentBalances = append(h.sqlPaymentBalances, handler.ConstructCreateSQL())
	}
}

func (h *updateHandler) constructPaymentTransferSQLs(ctx context.Context) error {
	for _, req := range h.PaymentTransferReqs {
		handler, _ := paymenttransfer1.NewHandler(ctx)
		handler.Req = *req
		if h.newPayment {
			h.sqlPaymentTransfers = append(h.sqlPaymentTransfers, handler.ConstructCreateSQL())
		} else {
			sql, err := handler.ConstructUpdateSQL()
			if err == cruder.ErrUpdateNothing {
				continue
			}
			if err != nil {
				return err
			}
			h.sqlPaymentTransfers = append(h.sqlPaymentTransfers, sql)
		}
	}
	return nil
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil {
		return err
	}
	if n == 1 {
		h.updateNothing = false
	}
	return nil
}

func (h *updateHandler) updateOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlOrderStateBase == "" {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlOrderStateBase)
}

func (h *updateHandler) updatePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	if h.sqlPowerRentalState == "" {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlPowerRentalState)
}

func (h *updateHandler) createOrderLocks(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlOrderLocks {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *updateHandler) createPaymentBalanceLock(ctx context.Context, tx *ent.Tx) error {
	if !h.newPaymentBalance {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlPaymentBalanceLock)
}

func (h *updateHandler) createPaymentBase(ctx context.Context, tx *ent.Tx) error {
	if !h.newPayment {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlPaymentBase)
}

func (h *updateHandler) updateObseletePaymentBase(ctx context.Context, tx *ent.Tx) error {
	if !h.newPayment {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlObseletePaymentBase)
}

func (h *updateHandler) createPaymentBalances(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentBalances {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *updateHandler) createOrUpdatePaymentTransfers(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentTransfers {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *updateHandler) formalizeOrderID() {
	if h.OrderID != nil {
		return
	}
	h.OrderID = func() *uuid.UUID { uid := h._ent.OrderID(); return &uid }()
	h.OrderBaseReq.EntID = h.OrderID
	h.OrderStateBaseReq.OrderID = h.OrderID
	h.PowerRentalStateReq.OrderID = h.OrderID
	h.PaymentBaseReq.OrderID = h.OrderID
}

func (h *updateHandler) formalizeOrderLocks() {
	for _, req := range h.OrderLockReqs {
		req.OrderID = h.OrderID
	}
}

func (h *updateHandler) formalizePaymentBalances() {
	if !h.newPaymentBalance {
		return
	}
	for _, req := range h.PaymentBalanceReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *updateHandler) formalizePaymentTransfers() {
	if !h.newPayment {
		return
	}
	for _, req := range h.PaymentTransferReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *updateHandler) formalizePaymentID() error {
	if h.PaymentBaseReq.EntID == nil || h._ent.PaymentID() == *h.PaymentBaseReq.EntID {
		return nil
	}

	h.newPayment = true
	ledgerLockID := func() *uuid.UUID {
		for _, req := range h.OrderLockReqs {
			if *req.LockType == types.OrderLockType_LockBalance {
				return req.EntID
			}
		}
		return nil
	}()
	h.newPaymentBalance = ledgerLockID != nil
	if h.newPaymentBalance && *ledgerLockID == h._ent.LedgerLockID() {
		return fmt.Errorf("invalid ledgerlock")
	}

	h.obseletePaymentBaseReq.EntID = func() *uuid.UUID { uid := h._ent.PaymentID(); return &uid }()
	h.PowerRentalStateReq.PaymentID = h.PaymentBaseReq.EntID
	return nil
}

func (h *updateHandler) formalizeEntIDs() {
	if h.PaymentBaseReq.EntID == nil {
		h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.PaymentBalanceLockReq.EntID == nil {
		h.PaymentBalanceLockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

func (h *Handler) UpdatePowerRentalWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &updateHandler{
		powerRentalQueryHandler: &powerRentalQueryHandler{
			Handler: h,
		},
		obseletePaymentBaseReq: &paymentbasecrud.Req{
			OrderID:       h.OrderID,
			ObseleteState: func() *types.PaymentObseleteState { e := types.PaymentObseleteState_PaymentObseleteWait; return &e }(),
		},
		updateNothing: true,
	}

	if err := handler.requirePowerRental(ctx); err != nil {
		return err
	}

	handler.formalizeOrderID()
	handler.formalizeOrderLocks()
	handler.formalizeEntIDs()
	if err := handler.formalizePaymentID(); err != nil {
		return err
	}
	handler.formalizePaymentBalances()
	handler.formalizePaymentTransfers()

	if err := handler.constructOrderStateBaseSQL(ctx); err != nil {
		return err
	}
	if err := handler.constructPowerRentalStateSQL(ctx); err != nil {
		return err
	}
	handler.constructOrderLockSQLs(ctx)
	handler.constructPaymentBalanceLockSQL(ctx)
	handler.constructPaymentBaseSQL(ctx)
	handler.constructPaymentBalanceSQLs(ctx)
	if err := handler.constructPaymentTransferSQLs(ctx); err != nil {
		return err
	}
	if err := handler.constructObseletePaymentBaseSQL(ctx); err != nil {
		return err
	}

	if err := handler.updateOrderStateBase(ctx, tx); err != nil {
		return err
	}
	if err := handler.updatePowerRentalState(ctx, tx); err != nil {
		return err
	}
	if err := handler.updateObseletePaymentBase(ctx, tx); err != nil {
		return err
	}
	if err := handler.createPaymentBase(ctx, tx); err != nil {
		return err
	}
	if err := handler.createOrderLocks(ctx, tx); err != nil {
		return err
	}
	if err := handler.createPaymentBalanceLock(ctx, tx); err != nil {
		return err
	}
	if err := handler.createPaymentBalances(ctx, tx); err != nil {
		return err
	}
	if err := handler.createOrUpdatePaymentTransfers(ctx, tx); err != nil {
		return err
	}
	if handler.updateNothing {
		return cruder.ErrUpdateNothing
	}
	return handler.FeeMultiHandler.UpdateFeeOrdersWithTx(ctx, tx)
}

func (h *Handler) UpdatePowerRental(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.UpdatePowerRentalWithTx(_ctx, tx)
	})
}