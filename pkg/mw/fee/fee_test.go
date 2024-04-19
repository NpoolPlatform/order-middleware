package feeorder

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	// basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
	ordercouponmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order/coupon"
	paymentmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	orderbase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/orderbase"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/order-middleware/pkg/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.FeeOrder{
	EntID:             uuid.NewString(),
	AppID:             uuid.NewString(),
	UserID:            uuid.NewString(),
	GoodID:            uuid.NewString(),
	GoodType:          goodtypes.GoodType_TechniqueServiceFee,
	AppGoodID:         uuid.NewString(),
	OrderID:           uuid.NewString(),
	ParentOrderID:     uuid.NewString(),
	OrderType:         types.OrderType_Normal,
	PaymentType:       types.PaymentType_PayWithBalanceOnly,
	CreateMethod:      types.OrderCreateMethod_OrderCreatedByAdmin,
	GoodValueUSD:      decimal.NewFromInt(120).String(),
	PaymentAmountUSD:  decimal.NewFromInt(110).String(),
	DiscountAmountUSD: decimal.NewFromInt(10).String(),
	PromotionID:       uuid.NewString(),
	DurationSeconds:   100000,
	LedgerLockID:      uuid.NewString(),
	PaymentID:         uuid.NewString(),
	Coupons: []*ordercouponmwpb.OrderCouponInfo{
		{
			CouponID: uuid.NewString(),
		},
	},
	PaymentBalances: []*paymentmwpb.PaymentBalanceInfo{
		{
			CoinTypeID:           uuid.NewString(),
			Amount:               decimal.NewFromInt(110).String(),
			LocalCoinUSDCurrency: decimal.NewFromInt(1).String(),
			LiveCoinUSDCurrency:  decimal.NewFromInt(1).String(),
		},
	},
}

//nolint:funlen
func setup(t *testing.T) func(*testing.T) {
	h1, err := orderbase1.NewHandler(
		context.Background(),
		orderbase1.WithEntID(&ret.ParentOrderID, false),
		orderbase1.WithAppID(&ret.AppID, true),
		orderbase1.WithUserID(&ret.UserID, true),
		orderbase1.WithGoodID(func() *string { s := uuid.NewString(); return &s }(), true),
		orderbase1.WithGoodType(&ret.GoodType, true),
		orderbase1.WithAppGoodID(func() *string { s := uuid.NewString(); return &s }(), true),
		orderbase1.WithOrderType(func() *types.OrderType { e := types.OrderType_Offline; return &e }(), true),
		orderbase1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateOrderBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteOrderBase(context.Background())
	}
}

func createFeeOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithOrderID(&ret.OrderID, true),
		WithParentOrderID(&ret.ParentOrderID, true),
		WithOrderType(&ret.OrderType, true),
		WithPaymentType(&ret.PaymentType, true),
		WithCreateMethod(&ret.CreateMethod, true),
		WithGoodValueUSD(&ret.GoodValueUSD, true),
		WithPaymentAmountUSD(&ret.PaymentAmountUSD, true),
		WithDiscountAmountUSD(&ret.DiscountAmountUSD, true),
		WithPromotionID(&ret.PromotionID, true),
		WithDurationSeconds(&ret.DurationSeconds, true),
		WithLedgerLockID(&ret.LedgerLockID, true),
		WithPaymentID(&ret.PaymentID, true),
		WithCouponIDs(func() (_couponIDs []string) {
			for _, coupon := range ret.Coupons {
				_couponIDs = append(_couponIDs, coupon.CouponID)
			}
			return
		}(), true),
		WithPaymentBalances(func() (_reqs []*paymentmwpb.PaymentBalanceReq) {
			for _, req := range ret.PaymentBalances {
				_reqs = append(_reqs, &paymentmwpb.PaymentBalanceReq{
					CoinTypeID:           &req.CoinTypeID,
					Amount:               &req.Amount,
					LocalCoinUSDCurrency: &req.LocalCoinUSDCurrency,
					LiveCoinUSDCurrency:  &req.LiveCoinUSDCurrency,
				})
			}
			return
		}(), true),
		WithPaymentTransfers([]*paymentmwpb.PaymentTransferReq{}, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateFeeOrder(context.Background())
		if assert.Nil(t, err) {
			/*
				info, err := handler.GetFeeOrder(context.Background())
				if assert.Nil(t, err) {
					ret.CreatedAt = info.CreatedAt
					ret.UpdatedAt = info.UpdatedAt
					ret.ID = info.ID
					assert.Equal(t, &ret, info)
				}
			*/
		}
	}
}

/*
func updateFeeOrder(t *testing.T) {
	ret.FeeOrderSeconds = 180
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, false),
		WithFeeOrderSeconds(&ret.FeeOrderSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateFeeOrder(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetFeeOrder(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getFeeOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetFeeOrder(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getFeeOrders(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetFeeOrders(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteFeeOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteFeeOrder(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetFeeOrder(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}
*/

func TestFeeOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createFeeOrder", createFeeOrder)
	// t.Run("updateFeeOrder", updateFeeOrder)
	// t.Run("getFeeOrder", getFeeOrder)
	// t.Run("getFeeOrders", getFeeOrders)
	// t.Run("deleteFeeOrder", deleteFeeOrder)
}
