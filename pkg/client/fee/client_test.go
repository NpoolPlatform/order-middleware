package feeorder

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
	ordercouponmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order/coupon"
	paymentmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	powerrentalmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"
	powerrentalmwcli "github.com/NpoolPlatform/order-middleware/pkg/client/powerrental"
	testinit "github.com/NpoolPlatform/order-middleware/pkg/testinit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
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
	EntID:               uuid.NewString(),
	AppID:               uuid.NewString(),
	UserID:              uuid.NewString(),
	GoodID:              uuid.NewString(),
	GoodType:            goodtypes.GoodType_TechniqueServiceFee,
	AppGoodID:           uuid.NewString(),
	OrderID:             uuid.NewString(),
	ParentOrderID:       uuid.NewString(),
	ParentAppGoodID:     uuid.NewString(),
	ParentGoodType:      goodtypes.GoodType_PowerRental,
	OrderType:           types.OrderType_Normal,
	PaymentType:         types.PaymentType_PayWithBalanceOnly,
	CreateMethod:        types.OrderCreateMethod_OrderCreatedByPurchase,
	GoodValueUSD:        decimal.NewFromInt(120).String(),
	PaymentGoodValueUSD: decimal.NewFromInt(0).String(),
	PaymentAmountUSD:    decimal.NewFromInt(110).String(),
	DiscountAmountUSD:   decimal.NewFromInt(10).String(),
	PromotionID:         uuid.NewString(),
	DurationSeconds:     100000,
	LedgerLockID:        uuid.NewString(),
	PaymentID:           uuid.NewString(),
	Coupons: []*ordercouponmwpb.OrderCouponInfo{
		{
			CouponID: uuid.NewString(),
		},
	},
	PaymentBalances: []*paymentmwpb.PaymentBalanceInfo{
		{
			CoinTypeID:           uuid.NewString(),
			Amount:               decimal.NewFromInt(110).String(),
			CoinUSDCurrency:      decimal.NewFromInt(1).String(),
			LocalCoinUSDCurrency: decimal.NewFromInt(1).String(),
			LiveCoinUSDCurrency:  decimal.NewFromInt(1).String(),
		},
	},
	OrderState:   types.OrderState_OrderStateCreated,
	PaymentState: types.PaymentState_PaymentStateWait,
}

func setup(t *testing.T) func(*testing.T) {
	for _, paymentBalance := range ret.PaymentBalances {
		paymentBalance.PaymentID = ret.PaymentID
	}
	for _, orderCoupon := range ret.Coupons {
		orderCoupon.OrderID = ret.OrderID
	}

	ret.GoodTypeStr = ret.GoodType.String()
	ret.OrderTypeStr = ret.OrderType.String()
	ret.PaymentTypeStr = ret.PaymentType.String()
	ret.CreateMethodStr = ret.CreateMethod.String()
	ret.OrderStateStr = ret.OrderState.String()
	ret.PaymentStateStr = ret.PaymentState.String()
	ret.CancelStateStr = ret.CancelState.String()
	ret.ParentGoodTypeStr = ret.ParentGoodType.String()

	err := powerrentalmwcli.CreatePowerRentalOrder(context.Background(), &powerrentalmwpb.PowerRentalOrderReq{
		EntID:              func() *string { s := uuid.NewString(); return &s }(),
		AppID:              &ret.AppID,
		UserID:             &ret.UserID,
		GoodID:             func() *string { s := uuid.NewString(); return &s }(),
		GoodType:           &ret.ParentGoodType,
		AppGoodID:          &ret.ParentAppGoodID,
		OrderID:            &ret.ParentOrderID,
		OrderType:          func() *types.OrderType { e := types.OrderType_Offline; return &e }(),
		AppGoodStockID:     func() *string { s := uuid.NewString(); return &s }(),
		Units:              func() *string { s := decimal.NewFromInt(10).String(); return &s }(),
		PaymentType:        func() *types.PaymentType { e := types.PaymentType_PayWithOffline; return &e }(),
		CreateMethod:       func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(),
		GoodValueUSD:       func() *string { s := decimal.NewFromInt(10).String(); return &s }(),
		DurationSeconds:    func() *uint32 { u := uint32(123900); return &u }(),
		AppGoodStockLockID: func() *string { s := uuid.NewString(); return &s }(),
		StartMode:          func() *types.OrderStartMode { e := types.OrderStartMode_OrderStartInstantly; return &e }(),
		StartAt:            func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = powerrentalmwcli.DeletePowerRentalOrder(context.Background(), nil, nil, &ret.ParentOrderID)
	}
}

func createFeeOrder(t *testing.T) {
	err := CreateFeeOrder(context.Background(), &npool.FeeOrderReq{
		EntID:             &ret.EntID,
		AppID:             &ret.AppID,
		UserID:            &ret.UserID,
		GoodID:            &ret.GoodID,
		GoodType:          &ret.GoodType,
		AppGoodID:         &ret.AppGoodID,
		OrderID:           &ret.OrderID,
		OrderType:         &ret.OrderType,
		ParentOrderID:     &ret.ParentOrderID,
		PaymentType:       &ret.PaymentType,
		CreateMethod:      &ret.CreateMethod,
		GoodValueUSD:      &ret.GoodValueUSD,
		PaymentAmountUSD:  &ret.PaymentAmountUSD,
		DiscountAmountUSD: &ret.DiscountAmountUSD,
		PromotionID:       &ret.PromotionID,
		DurationSeconds:   &ret.DurationSeconds,
		LedgerLockID:      &ret.LedgerLockID,
		PaymentID:         &ret.PaymentID,
		CouponIDs: func() (_couponIDs []string) {
			for _, coupon := range ret.Coupons {
				_couponIDs = append(_couponIDs, coupon.CouponID)
			}
			return
		}(),
		PaymentBalances: func() (_reqs []*paymentmwpb.PaymentBalanceReq) {
			for _, req := range ret.PaymentBalances {
				_reqs = append(_reqs, &paymentmwpb.PaymentBalanceReq{
					CoinTypeID:           &req.CoinTypeID,
					Amount:               &req.Amount,
					LocalCoinUSDCurrency: &req.LocalCoinUSDCurrency,
					LiveCoinUSDCurrency:  &req.LiveCoinUSDCurrency,
				})
			}
			return
		}(),
	})
	if assert.Nil(t, err) {
		info, err := GetFeeOrder(context.Background(), ret.OrderID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			for _, paymentBalance := range ret.PaymentBalances {
				paymentBalance.CreatedAt = ret.CreatedAt
			}
			for _, orderCoupon := range ret.Coupons {
				orderCoupon.CreatedAt = ret.CreatedAt
			}
			assert.Equal(t, info, &ret)
		}
	}
}

func updateFeeOrder(t *testing.T) {
	ret.PaymentID = uuid.NewString()
	ret.LedgerLockID = uuid.NewString()
	for _, paymentBalance := range ret.PaymentBalances {
		paymentBalance.PaymentID = ret.PaymentID
	}

	err := UpdateFeeOrder(context.Background(), &npool.FeeOrderReq{
		ID:           &ret.ID,
		EntID:        &ret.EntID,
		OrderID:      &ret.OrderID,
		PaymentType:  &ret.PaymentType,
		LedgerLockID: &ret.LedgerLockID,
		PaymentID:    &ret.PaymentID,
		CouponIDs: func() (_couponIDs []string) {
			for _, coupon := range ret.Coupons {
				_couponIDs = append(_couponIDs, coupon.CouponID)
			}
			return
		}(),
		PaymentBalances: func() (_reqs []*paymentmwpb.PaymentBalanceReq) {
			for _, req := range ret.PaymentBalances {
				_reqs = append(_reqs, &paymentmwpb.PaymentBalanceReq{
					CoinTypeID:           &req.CoinTypeID,
					Amount:               &req.Amount,
					LocalCoinUSDCurrency: &req.LocalCoinUSDCurrency,
					LiveCoinUSDCurrency:  &req.LiveCoinUSDCurrency,
				})
			}
			return
		}(),
	})
	if assert.Nil(t, err) {
		info, err := GetFeeOrder(context.Background(), ret.OrderID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info.String(), ret.String())
		}
	}
}

func getFeeOrder(t *testing.T) {
	info, err := GetFeeOrder(context.Background(), ret.OrderID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func getFeeOrders(t *testing.T) {
	infos, _, err := GetFeeOrders(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0].String(), ret.String())
	}
}

func deleteFeeOrder(t *testing.T) {
	err := DeleteFeeOrder(context.Background(), &ret.ID, &ret.EntID, &ret.OrderID)
	assert.Nil(t, err)

	info, err := GetFeeOrder(context.Background(), ret.OrderID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setup(t)
	defer teardown(t)

	t.Run("createFeeOrder", createFeeOrder)
	t.Run("updateFeeOrder", updateFeeOrder)
	t.Run("getFeeOrder", getFeeOrder)
	t.Run("getFeeOrders", getFeeOrders)
	t.Run("deleteFeeOrder", deleteFeeOrder)
}
