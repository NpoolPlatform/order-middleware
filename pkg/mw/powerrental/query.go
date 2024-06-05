package powerrental

import (
	"context"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	feeordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
	ordercouponmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order/coupon"
	paymentmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"
	ordercouponcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/coupon"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	paymentbalancecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment/balance"
	paymenttransfercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment/transfer"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entfeeorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorder"
	entfeeorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorderstate"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entordercoupon "github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	entpaymentbalance "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbalance"
	entpaymenttransfer "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymenttransfer"
	entpowerrental "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrental"
	entpowerrentalstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrentalstate"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount         *ent.OrderBaseSelect
	infos            []*npool.PowerRentalOrder
	orderCoupons     []*ordercouponmwpb.OrderCouponInfo
	paymentBalances  []*paymentmwpb.PaymentBalanceInfo
	paymentTransfers []*paymentmwpb.PaymentTransferInfo
	feeDurations     []*feeordermwpb.FeeDuration
	total            uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		if err := h.queryJoinPowerRental(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRental", "Error", err)
		}
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
		if err := h.queryJoinPowerRentalState(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRentalState", "Error", err)
		}
		h.queryJoinPaymentBase(s)
		h.queryJoinStockLock(s)
		if err := h.queryJoinOrderCoupon(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderCoupon", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) queryOrderCoupons(ctx context.Context, cli *ent.Client) error {
	orderIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.OrderID))
		}
		return
	}()

	stm, err := ordercouponcrud.SetQueryConds(
		cli.OrderCoupon.Query(),
		&ordercouponcrud.Conds{
			OrderIDs: &cruder.Cond{Op: cruder.IN, Val: orderIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entordercoupon.FieldOrderID,
		entordercoupon.FieldCouponID,
		entordercoupon.FieldCreatedAt,
	).Scan(ctx, &h.orderCoupons)
}

func (h *queryHandler) queryPaymentBalances(ctx context.Context, cli *ent.Client) error {
	paymentIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if _, err := uuid.Parse(info.PaymentID); err != nil {
				continue
			}
			uids = append(uids, uuid.MustParse(info.PaymentID))
		}
		return
	}()

	stm, err := paymentbalancecrud.SetQueryConds(
		cli.PaymentBalance.Query(),
		&paymentbalancecrud.Conds{
			PaymentIDs: &cruder.Cond{Op: cruder.IN, Val: paymentIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entpaymentbalance.FieldPaymentID,
		entpaymentbalance.FieldCoinTypeID,
		entpaymentbalance.FieldAmount,
		entpaymentbalance.FieldCoinUsdCurrency,
		entpaymentbalance.FieldLocalCoinUsdCurrency,
		entpaymentbalance.FieldLiveCoinUsdCurrency,
		entpaymentbalance.FieldCreatedAt,
	).Scan(ctx, &h.paymentBalances)
}

func (h *queryHandler) queryPaymentTransfers(ctx context.Context, cli *ent.Client) error {
	paymentIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if _, err := uuid.Parse(info.PaymentID); err != nil {
				continue
			}
			uids = append(uids, uuid.MustParse(info.PaymentID))
		}
		return
	}()

	stm, err := paymenttransfercrud.SetQueryConds(
		cli.PaymentTransfer.Query(),
		&paymenttransfercrud.Conds{
			PaymentIDs: &cruder.Cond{Op: cruder.IN, Val: paymentIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entpaymenttransfer.FieldPaymentID,
		entpaymenttransfer.FieldCoinTypeID,
		entpaymenttransfer.FieldAmount,
		entpaymenttransfer.FieldAccountID,
		entpaymenttransfer.FieldStartAmount,
		entpaymenttransfer.FieldCoinUsdCurrency,
		entpaymenttransfer.FieldLocalCoinUsdCurrency,
		entpaymenttransfer.FieldLiveCoinUsdCurrency,
		entpaymenttransfer.FieldFinishAmount,
		entpaymenttransfer.FieldCreatedAt,
	).Scan(ctx, &h.paymentTransfers)
}

func (h *queryHandler) queryFeeDurations(ctx context.Context, cli *ent.Client) error {
	cli = cli.Debug()
	orderIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.OrderID))
		}
		return
	}()
	stm, err := orderbasecrud.SetQueryConds(
		cli.OrderBase.Query(),
		&orderbasecrud.Conds{
			ParentOrderIDs: &cruder.Cond{Op: cruder.IN, Val: orderIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	return stm.GroupBy(
		entorderbase.FieldParentOrderID,
		entorderbase.FieldAppGoodID,
	).Aggregate(func(s *sql.Selector) string {
		t1 := sql.Table(entfeeorder.Table)
		s.Join(t1).
			On(
				s.C(entorderbase.FieldEntID),
				t1.C(entfeeorder.FieldOrderID),
			)
		return sql.As(
			sql.Sum(t1.C(entfeeorder.FieldDurationSeconds)),
			"total_duration_seconds",
		)
	}).Scan(ctx, &h.feeDurations)
}

func (h *queryHandler) queryOrdersPaymentGoodValueUSD(ctx context.Context, cli *ent.Client) error {
	cli = cli.Debug()
	orderIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.OrderID))
		}
		return
	}()
	stm, err := orderbasecrud.SetQueryConds(
		cli.OrderBase.Query(),
		&orderbasecrud.Conds{
			EntIDs: &cruder.Cond{Op: cruder.IN, Val: orderIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	goodValueUSDs := []struct {
		OrderID             string `json:"ent_id"`
		PaymentGoodValueUSD string `json:"payment_good_value_usd"`
	}{}
	if err := stm.GroupBy(
		entorderbase.FieldEntID,
	).Aggregate(func(s *sql.Selector) string {
		t0 := sql.Table(entorderstatebase.Table)
		t1 := sql.Table(entpowerrentalstate.Table)
		t2 := sql.Table(entfeeorderstate.Table)
		t3 := sql.Table(entfeeorderstate.Table)
		t4 := sql.Table(entfeeorderstate.Table)
		t5 := sql.Table(entpowerrental.Table)
		t6 := sql.Table(entfeeorder.Table)
		t7 := sql.Table(entfeeorder.Table)
		t8 := sql.Table(entfeeorder.Table)
		s.Join(t0).
			On(
				s.C(entorderbase.FieldEntID),
				t0.C(entorderstatebase.FieldOrderID),
			).
			OnP(
				sql.Or(
					sql.EQ(t0.C(entorderstatebase.FieldPaymentType), types.PaymentType_PayWithBalanceOnly.String()),
					sql.EQ(t0.C(entorderstatebase.FieldPaymentType), types.PaymentType_PayWithTransferOnly.String()),
					sql.EQ(t0.C(entorderstatebase.FieldPaymentType), types.PaymentType_PayWithTransferAndBalance.String()),
				),
			).
			LeftJoin(t1).
			On(
				s.C(entorderbase.FieldEntID),
				t1.C(entpowerrentalstate.FieldOrderID),
			).
			LeftJoin(t2).
			On(
				s.C(entorderbase.FieldEntID),
				t2.C(entfeeorderstate.FieldOrderID),
			).
			LeftJoin(t3).
			On(
				t1.C(entpowerrentalstate.FieldPaymentID),
				t3.C(entfeeorderstate.FieldPaymentID),
			).
			LeftJoin(t4).
			On(
				t2.C(entfeeorderstate.FieldPaymentID),
				t4.C(entfeeorderstate.FieldPaymentID),
			).
			LeftJoin(t5).
			On(
				t1.C(entpowerrentalstate.FieldOrderID),
				t5.C(entpowerrental.FieldOrderID),
			).
			LeftJoin(t6).
			On(
				t2.C(entfeeorderstate.FieldOrderID),
				t6.C(entfeeorder.FieldOrderID),
			).
			LeftJoin(t7).
			On(
				t3.C(entfeeorderstate.FieldOrderID),
				t7.C(entfeeorder.FieldOrderID),
			).
			LeftJoin(t8).
			On(
				t4.C(entfeeorderstate.FieldOrderID),
				t8.C(entfeeorder.FieldOrderID),
			)
		return sql.As(
			sql.Sum(
				"ifnull("+t5.C(entpowerrental.FieldGoodValueUsd)+", 0)+"+
					"ifnull("+t6.C(entfeeorder.FieldGoodValueUsd)+", 0)+"+
					"ifnull("+t7.C(entfeeorder.FieldGoodValueUsd)+", 0)+"+
					"ifnull("+t8.C(entfeeorder.FieldGoodValueUsd)+", 0)",
			),
			"payment_good_value_usd",
		)
	}).Scan(ctx, &goodValueUSDs); err != nil {
		return wlog.WrapError(err)
	}
	for _, info := range h.infos {
		for _, goodValueUSD := range goodValueUSDs {
			if info.OrderID == goodValueUSD.OrderID {
				info.PaymentGoodValueUSD = goodValueUSD.PaymentGoodValueUSD
				break
			}
		}
	}
	return nil
}

func (h *queryHandler) formalize() {
	orderCoupons := map[string][]*ordercouponmwpb.OrderCouponInfo{}
	paymentBalances := map[string][]*paymentmwpb.PaymentBalanceInfo{}
	paymentTransfers := map[string][]*paymentmwpb.PaymentTransferInfo{}
	feeDurations := map[string][]*feeordermwpb.FeeDuration{}

	for _, orderCoupon := range h.orderCoupons {
		orderCoupons[orderCoupon.OrderID] = append(orderCoupons[orderCoupon.OrderID], orderCoupon)
	}
	for _, paymentBalance := range h.paymentBalances {
		paymentBalances[paymentBalance.PaymentID] = append(paymentBalances[paymentBalance.PaymentID], paymentBalance)
		paymentBalance.Amount = func() string { amount, _ := decimal.NewFromString(paymentBalance.Amount); return amount.String() }()
		paymentBalance.CoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentBalance.CoinUSDCurrency)
			return amount.String()
		}()
		paymentBalance.LocalCoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentBalance.LocalCoinUSDCurrency)
			return amount.String()
		}()
		paymentBalance.LiveCoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentBalance.LiveCoinUSDCurrency)
			return amount.String()
		}()
	}
	for _, paymentTransfer := range h.paymentTransfers {
		paymentTransfers[paymentTransfer.PaymentID] = append(paymentTransfers[paymentTransfer.PaymentID], paymentTransfer)
		paymentTransfer.Amount = func() string { amount, _ := decimal.NewFromString(paymentTransfer.Amount); return amount.String() }()
		paymentTransfer.StartAmount = func() string { amount, _ := decimal.NewFromString(paymentTransfer.StartAmount); return amount.String() }()
		paymentTransfer.FinishAmount = func() string {
			amount, _ := decimal.NewFromString(paymentTransfer.FinishAmount)
			return amount.String()
		}()
		paymentTransfer.CoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentTransfer.CoinUSDCurrency)
			return amount.String()
		}()
		paymentTransfer.LocalCoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentTransfer.LocalCoinUSDCurrency)
			return amount.String()
		}()
		paymentTransfer.LiveCoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentTransfer.LiveCoinUSDCurrency)
			return amount.String()
		}()
	}
	for _, feeDuration := range h.feeDurations {
		feeDurations[feeDuration.ParentOrderID] = append(feeDurations[feeDuration.ParentOrderID], feeDuration)
	}

	for _, info := range h.infos {
		info.Units = func() string { amount, _ := decimal.NewFromString(info.Units); return amount.String() }()
		info.GoodValueUSD = func() string { amount, _ := decimal.NewFromString(info.GoodValueUSD); return amount.String() }()
		info.PaymentGoodValueUSD = func() string { amount, _ := decimal.NewFromString(info.PaymentGoodValueUSD); return amount.String() }()
		info.PaymentAmountUSD = func() string { amount, _ := decimal.NewFromString(info.PaymentAmountUSD); return amount.String() }()
		info.DiscountAmountUSD = func() string { amount, _ := decimal.NewFromString(info.DiscountAmountUSD); return amount.String() }()
		info.GoodType = goodtypes.GoodType(goodtypes.GoodType_value[info.GoodTypeStr])
		info.OrderType = types.OrderType(types.OrderType_value[info.OrderTypeStr])
		info.PaymentType = types.PaymentType(types.PaymentType_value[info.PaymentTypeStr])
		info.PaymentState = types.PaymentState(types.PaymentState_value[info.PaymentStateStr])
		info.OrderState = types.OrderState(types.OrderState_value[info.OrderStateStr])
		info.CancelState = types.OrderState(types.OrderState_value[info.CancelStateStr])
		info.RenewState = types.OrderRenewState(types.OrderRenewState_value[info.RenewStateStr])
		info.CreateMethod = types.OrderCreateMethod(types.OrderCreateMethod_value[info.CreateMethodStr])
		info.StartMode = types.OrderStartMode(types.OrderStartMode_value[info.StartModeStr])
		info.BenefitState = types.BenefitState(types.BenefitState_value[info.BenefitStateStr])
		info.InvestmentType = types.InvestmentType(types.InvestmentType_value[info.InvestmentTypeStr])
		info.Coupons = orderCoupons[info.OrderID]
		info.PaymentBalances = paymentBalances[info.PaymentID]
		info.PaymentTransfers = paymentTransfers[info.PaymentID]
		info.EndAt = info.StartAt + info.DurationSeconds + info.CompensateSeconds
		info.FeeDurations = feeDurations[info.OrderID]
	}
}

func (h *Handler) GetPowerRental(ctx context.Context) (*npool.PowerRentalOrder, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderBase(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentBalances(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentTransfers(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryFeeDurations(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryOrdersPaymentGoodValueUSD(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return handler.queryOrderCoupons(_ctx, cli)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetPowerRentals(ctx context.Context) ([]*npool.PowerRentalOrder, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entpowerrental.FieldCreatedAt))

		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentBalances(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentTransfers(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryFeeDurations(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryOrdersPaymentGoodValueUSD(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return handler.queryOrderCoupons(_ctx, cli)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
