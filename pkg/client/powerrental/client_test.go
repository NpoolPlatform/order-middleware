package powerrental

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
	feeordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
	ordercouponmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order/coupon"
	paymentmwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

var ret = npool.PowerRentalOrder{
	EntID:              uuid.NewString(),
	AppID:              uuid.NewString(),
	UserID:             uuid.NewString(),
	GoodID:             uuid.NewString(),
	GoodType:           goodtypes.GoodType_PowerRental,
	AppGoodID:          uuid.NewString(),
	OrderID:            uuid.NewString(),
	OrderType:          types.OrderType_Normal,
	AppGoodStockID:     uuid.NewString(),
	Units:              decimal.NewFromInt(10).String(),
	PaymentType:        types.PaymentType_PayWithBalanceOnly,
	CreateMethod:       types.OrderCreateMethod_OrderCreatedByAdmin,
	GoodValueUSD:       decimal.NewFromInt(120).String(),
	PaymentAmountUSD:   decimal.NewFromInt(110).String(),
	DiscountAmountUSD:  decimal.NewFromInt(10).String(),
	PromotionID:        uuid.NewString(),
	DurationSeconds:    100000,
	AppGoodStockLockID: uuid.NewString(),
	LedgerLockID:       uuid.NewString(),
	PaymentID:          uuid.NewString(),
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
	StartMode:    types.OrderStartMode_OrderStartInstantly,
	StartAt:      uint32(time.Now().Unix() + 100),
}

//nolint:unparam
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
	ret.StartModeStr = ret.StartMode.String()
	ret.RenewStateStr = types.OrderRenewState_OrderRenewWait.String()

	return func(*testing.T) {}
}

func createPowerRentalOrderWithFees(t *testing.T) {
	err := CreatePowerRentalOrderWithFees(context.Background(), &npool.PowerRentalOrderReq{
		EntID:              &ret.EntID,
		AppID:              &ret.AppID,
		UserID:             &ret.UserID,
		GoodID:             &ret.GoodID,
		GoodType:           &ret.GoodType,
		AppGoodID:          &ret.AppGoodID,
		OrderID:            &ret.OrderID,
		OrderType:          &ret.OrderType,
		AppGoodStockID:     &ret.AppGoodStockID,
		Units:              &ret.Units,
		PaymentType:        &ret.PaymentType,
		CreateMethod:       &ret.CreateMethod,
		GoodValueUSD:       &ret.GoodValueUSD,
		PaymentAmountUSD:   &ret.PaymentAmountUSD,
		DiscountAmountUSD:  &ret.DiscountAmountUSD,
		PromotionID:        &ret.PromotionID,
		DurationSeconds:    &ret.DurationSeconds,
		AppGoodStockLockID: &ret.AppGoodStockLockID,
		LedgerLockID:       &ret.LedgerLockID,
		PaymentID:          &ret.PaymentID,
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
		StartMode: &ret.StartMode,
		StartAt:   &ret.StartAt,
	},
		[]*feeordermwpb.FeeOrderReq{
			{
				EntID:           func() *string { s := uuid.NewString(); return &s }(),
				GoodID:          func() *string { s := uuid.NewString(); return &s }(),
				GoodType:        func() *goodtypes.GoodType { e := goodtypes.GoodType_TechniqueServiceFee; return &e }(),
				AppGoodID:       func() *string { s := uuid.NewString(); return &s }(),
				OrderID:         func() *string { s := uuid.NewString(); return &s }(),
				GoodValueUSD:    func() *string { s := decimal.NewFromInt(100).String(); return &s }(),
				DurationSeconds: func() *uint32 { u := uint32(150); return &u }(),
			},
		})
	if assert.Nil(t, err) {
		info, err := GetPowerRentalOrder(context.Background(), ret.OrderID)
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

func updatePowerRentalOrder(t *testing.T) {
	ret.PaymentID = uuid.NewString()
	ret.LedgerLockID = uuid.NewString()
	for _, paymentBalance := range ret.PaymentBalances {
		paymentBalance.PaymentID = ret.PaymentID
	}

	err := UpdatePowerRentalOrder(context.Background(), &npool.PowerRentalOrderReq{
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
		StartMode: &ret.StartMode,
		StartAt:   &ret.StartAt,
	})
	if assert.Nil(t, err) {
		info, err := GetPowerRentalOrder(context.Background(), ret.OrderID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info.String(), ret.String())
		}
	}
}

func getPowerRentalOrder(t *testing.T) {
	info, err := GetPowerRentalOrder(context.Background(), ret.OrderID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func getPowerRentalOrders(t *testing.T) {
	infos, _, err := GetPowerRentalOrders(context.Background(), &npool.Conds{
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
	}, 0, 1)
	if assert.Nil(t, err) && assert.Equal(t, len(infos), 1) {
		assert.Equal(t, infos[0].String(), ret.String())
	}
}

func deletePowerRentalOrder(t *testing.T) {
	err := DeletePowerRentalOrder(context.Background(), &ret.ID, &ret.EntID, &ret.OrderID)
	assert.Nil(t, err)

	info, err := GetPowerRentalOrder(context.Background(), ret.OrderID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestPowerRentalOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setup(t)
	defer teardown(t)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createPowerRentalOrderWithFees", createPowerRentalOrderWithFees)
	t.Run("updatePowerRentalOrder", updatePowerRentalOrder)
	t.Run("getPowerRentalOrder", getPowerRentalOrder)
	t.Run("getPowerRentalOrders", getPowerRentalOrders)
	t.Run("deletePowerRentalOrder", deletePowerRentalOrder)
}
