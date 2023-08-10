package payment

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
)

type existHandler struct {
	*Handler
	stm *ent.PaymentQuery
}

func (h *existHandler) queryPayment(cli *ent.Client) {
	h.stm = cli.Payment.
		Query().
		Where(
			entpayment.ID(*h.ID),
			entpayment.DeletedAt(0),
		)
}

func (h *existHandler) queryPayments(cli *ent.Client) error {
	stm, err := paymentcrud.SetQueryConds(cli.Payment.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *Handler) ExistPayment(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryPayment(cli)
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistPaymentConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPayments(cli); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
