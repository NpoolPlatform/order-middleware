package payment

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
)

type queryHandler struct {
	*Handler
	stm   *ent.PaymentSelect
	infos []*npool.Payment
	total uint32
}

func (h *queryHandler) selectPayment(stm *ent.PaymentQuery) {
	h.stm = stm.Select(
		entpayment.FieldID,
		entpayment.FieldAppID,
		entpayment.FieldGoodID,
		entpayment.FieldUserID,
		entpayment.FieldOrderID,
		entpayment.FieldAccountID,
		entpayment.FieldStartAmount,
		entpayment.FieldAmount,
		entpayment.FieldPayWithBalanceAmount,
		entpayment.FieldFinishAmount,
		entpayment.FieldCoinUsdCurrency,
		entpayment.FieldLocalCoinUsdCurrency,
		entpayment.FieldLiveCoinUsdCurrency,
		entpayment.FieldCoinInfoID,
		entpayment.FieldState,
		entpayment.FieldChainTransactionID,
		entpayment.FieldUserSetPaid,
		entpayment.FieldUserSetCanceled,
		entpayment.FieldFakePayment,
		entpayment.FieldCreatedAt,
		entpayment.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryPayment(cli *ent.Client) {
	h.selectPayment(
		cli.Payment.
			Query().
			Where(
				entpayment.ID(*h.ID),
				entpayment.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryPayments(ctx context.Context, cli *ent.Client) error {
	stm, err := paymentcrud.SetQueryConds(cli.Payment.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectPayment(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.State = basetypes.PaymentState(basetypes.PaymentState_value[info.PaymentStateStr])
	}
}

func (h *Handler) GetPayment(ctx context.Context) (*npool.Payment, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryPayment(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetPayments(ctx context.Context) ([]*npool.Payment, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPayments(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Desc(entpayment.FieldCreatedAt))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
