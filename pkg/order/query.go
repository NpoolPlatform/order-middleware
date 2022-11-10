package order

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
	order1 "github.com/NpoolPlatform/order-manager/pkg/db/ent/order"

	"github.com/NpoolPlatform/order-manager/pkg/db"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/payment"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	"github.com/google/uuid"
)

func GetOrder(ctx context.Context, id string) (info *npool.Order, err error) {
	infos := []*npool.Order{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Order.
			Query().
			Where(
				order1.ID(uuid.MustParse(id)),
			)

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("invalid order id")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	infos, err = expand(infos)
	if err != nil {
		return nil, err
	}

	return infos[0], nil
}

func GetOrders(ctx context.Context, appID, userID string, offset, limit int32) (infos []*npool.Order, total uint32, err error) {
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Order.
			Query().
			Where(
				order1.AppID(uuid.MustParse(appID)),
				order1.UserID(uuid.MustParse(userID)),
			)

		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		stm = stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	infos, err = expand(infos)
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func GetAppOrders(ctx context.Context, appID string, offset, limit int32) (infos []*npool.Order, total uint32, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Order.
			Query().
			Where(
				order1.AppID(uuid.MustParse(appID)),
			)

		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		stm = stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	infos, err = expand(infos)
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func join(stm *ent.OrderQuery) *ent.OrderSelect {
	return stm.
		Select(
			order1.FieldID,
			order1.FieldAppID,
			order1.FieldUserID,
			order1.FieldGoodID,
			order1.FieldUnits,
			order1.FieldType,
			order1.FieldState,
			order1.FieldParentOrderID,
			order1.FieldStartAt,
			order1.FieldEndAt,
			order1.FieldPayWithParent,
			order1.FieldFixAmountCouponID,
			order1.FieldDiscountCouponID,
			order1.FieldUserSpecialReductionID,
			order1.FieldCreatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(payment.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(order1.FieldID),
					t1.C(payment.FieldOrderID),
				).
				AppendSelect(
					sql.As(t1.C(payment.FieldCoinInfoID), "payment_coin_type_id"),
					sql.As(t1.C(payment.FieldCoinUsdCurrency), "payment_coin_usd_currency"),
					sql.As(t1.C(payment.FieldLiveCoinUsdCurrency), "payment_live_coin_usd_currency"),
					sql.As(t1.C(payment.FieldLocalCoinUsdCurrency), "payment_local_coin_usd_currency"),
					sql.As(t1.C(payment.FieldID), "payment_id"),
					sql.As(t1.C(payment.FieldAccountID), "payment_account_id"),
					sql.As(t1.C(payment.FieldAmount), "payment_amount"),
					sql.As(t1.C(payment.FieldState), "payment_state"),
					sql.As(t1.C(payment.FieldPayWithBalanceAmount), "pay_with_balance_amount"),
					sql.As(t1.C(payment.FieldUpdatedAt), "paid_at"),
					sql.As(t1.C(payment.FieldStartAmount), "payment_start_amount"),
					sql.As(t1.C(payment.FieldFinishAmount), "payment_finish_amount"),
					sql.As(t1.C(payment.FieldUserSetCanceled), "user_canceled"),
				)

			t2 := sql.Table(order1.Table)
			s.
				LeftJoin(t2).
				On(
					s.C(order1.FieldParentOrderID),
					t2.C(order1.FieldID),
				).
				AppendSelect(
					sql.As(t2.C(order1.FieldGoodID), "parent_order_good_id"),
				)
		})
}

func expand(infos []*npool.Order) ([]*npool.Order, error) { //nolint
	for _, info := range infos {
		info.OrderType = mgrpb.OrderType(mgrpb.OrderType_value[info.OrderTypeStr])
		info.OrderState = mgrpb.OrderState(mgrpb.OrderState_value[info.OrderStateStr])
	}
	return infos, nil
}
