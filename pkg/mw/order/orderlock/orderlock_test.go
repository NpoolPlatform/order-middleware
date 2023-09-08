package orderlock

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"

	"github.com/google/uuid"
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

const secondsPerDay = 24 * 60 * 60
const seconds = 1

var (
	now   = uint32(time.Now().Unix())
	order = ordermwpb.Order{
		ID:                   uuid.NewString(),
		AppID:                uuid.NewString(),
		UserID:               uuid.NewString(),
		GoodID:               uuid.NewString(),
		AppGoodID:            uuid.NewString(),
		Units:                "100",
		GoodValue:            "1007",
		GoodValueUSD:         "1007",
		PaymentAmount:        "1121",
		DiscountAmount:       "10",
		PromotionID:          uuid.NewString(),
		DurationDays:         now + 5*secondsPerDay,
		OrderTypeStr:         ordertypes.OrderType_Normal.String(),
		OrderType:            ordertypes.OrderType_Normal,
		InvestmentType:       ordertypes.InvestmentType_FullPayment,
		InvestmentTypeStr:    ordertypes.InvestmentType_FullPayment.String(),
		PaymentTypeStr:       ordertypes.PaymentType_PayWithTransferAndBalance.String(),
		PaymentType:          ordertypes.PaymentType_PayWithTransferAndBalance,
		CoinTypeID:           uuid.NewString(),
		PaymentCoinTypeID:    uuid.NewString(),
		TransferAmount:       "1011",
		BalanceAmount:        "110",
		CoinUSDCurrency:      "1004",
		LocalCoinUSDCurrency: "1005",
		LiveCoinUSDCurrency:  "1006",

		PaymentAccountID:   uuid.NewString(),
		PaymentStartAmount: "1010",

		OrderStateStr:        ordertypes.OrderState_OrderStateWaitPayment.String(),
		OrderState:           ordertypes.OrderState_OrderStateWaitPayment,
		StartModeStr:         ordertypes.OrderStartMode_OrderStartConfirmed.String(),
		StartMode:            ordertypes.OrderStartMode_OrderStartConfirmed,
		StartAt:              now + 5*seconds,
		EndAt:                now + 5*secondsPerDay,
		LastBenefitAt:        0,
		BenefitStateStr:      ordertypes.BenefitState_BenefitWait.String(),
		BenefitState:         ordertypes.BenefitState_BenefitWait,
		UserSetPaid:          false,
		UserSetCanceled:      false,
		AdminSetCanceled:     false,
		PaymentTransactionID: "",
		PaymentFinishAmount:  "0",
		PaymentStateStr:      ordertypes.PaymentState_PaymentStateWait.String(),
		PaymentState:         ordertypes.PaymentState_PaymentStateWait,
		OutOfGasHours:        0,
		CompensateHours:      0,
		AppGoodStockLockID:   uuid.NewString(),
		LedgerLockID:         uuid.NewString(),
	}
	rets = []npool.OrderLock{
		{
			ID:          uuid.NewString(),
			OrderID:     order.ID,
			AppID:       order.AppID,
			UserID:      order.UserID,
			LockType:    ordertypes.OrderLockType_LockCommission,
			LockTypeStr: ordertypes.OrderLockType_LockCommission.String(),
		},
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := order1.NewHandler(
		context.Background(),
		order1.WithID(&order.ID, false),
		order1.WithAppID(&order.AppID, true),
		order1.WithUserID(&order.UserID, true),
		order1.WithGoodID(&order.GoodID, true),
		order1.WithAppGoodID(&order.AppGoodID, true),
		order1.WithUnits(&order.Units, true),
		order1.WithGoodValue(&order.GoodValue, true),
		order1.WithGoodValueUSD(&order.GoodValueUSD, true),
		order1.WithPaymentAmount(&order.PaymentAmount, true),
		order1.WithDiscountAmount(&order.DiscountAmount, false),
		order1.WithPromotionID(&order.PromotionID, false),
		order1.WithDurationDays(&order.DurationDays, true),
		order1.WithOrderType(&order.OrderType, true),
		order1.WithInvestmentType(&order.InvestmentType, true),
		order1.WithCouponIDs(order.CouponIDs, false),
		order1.WithPaymentType(&order.PaymentType, true),
		order1.WithCoinTypeID(&order.CoinTypeID, true),
		order1.WithPaymentCoinTypeID(&order.PaymentCoinTypeID, true),
		order1.WithTransferAmount(&order.TransferAmount, true),
		order1.WithBalanceAmount(&order.BalanceAmount, true),
		order1.WithCoinUSDCurrency(&order.CoinUSDCurrency, true),
		order1.WithLocalCoinUSDCurrency(&order.LocalCoinUSDCurrency, true),
		order1.WithLiveCoinUSDCurrency(&order.LiveCoinUSDCurrency, true),
		order1.WithPaymentAccountID(&order.PaymentAccountID, false),
		order1.WithPaymentStartAmount(&order.PaymentStartAmount, false),
		order1.WithStartMode(&order.StartMode, true),
		order1.WithStartAt(&order.StartAt, true),
		order1.WithEndAt(&order.EndAt, true),
		order1.WithPaymentState(&order.PaymentState, false),
		order1.WithAppGoodStockLockID(&order.AppGoodStockLockID, true),
		order1.WithLedgerLockID(&order.LedgerLockID, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateOrder(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteOrder(context.Background())
	}
}

func createOrderLocks(t *testing.T) {
	retsReq := []*npool.OrderLockReq{}
	for key := range rets {
		retReq := npool.OrderLockReq{
			ID:       &rets[key].ID,
			AppID:    &rets[key].AppID,
			UserID:   &rets[key].UserID,
			OrderID:  &rets[key].OrderID,
			LockType: &rets[key].LockType,
		}
		retsReq = append(retsReq, &retReq)
	}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(retsReq),
	)
	if assert.Nil(t, err) {
		infos, err := handler.CreateOrderLocks(context.Background())
		if assert.Nil(t, err) {
			rets[0].CreatedAt = infos[0].CreatedAt
			rets[0].UpdatedAt = infos[0].UpdatedAt
			assert.Equal(t, len(rets), len(infos))
		}
	}
}

func getOrderLock(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&rets[0].ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetOrderLock(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &rets[0], info)
		}
	}
}

func getOrderLocks(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID: &basetypes.StringVal{Op: cruder.EQ, Value: rets[0].ID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetOrderLocks(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &rets[0], infos[0])
			}
		}
	}
}

func deleteOrderLocks(t *testing.T) {
	retsReq := []*npool.OrderLockReq{}
	for key := range rets {
		retReq := npool.OrderLockReq{
			ID:       &rets[key].ID,
			AppID:    &rets[key].AppID,
			UserID:   &rets[key].UserID,
			OrderID:  &rets[key].OrderID,
			LockType: &rets[key].LockType,
		}
		retsReq = append(retsReq, &retReq)
	}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(retsReq),
	)
	if assert.Nil(t, err) {
		infos, err := handler.DeleteOrderLocks(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, len(infos), 0)
		}
	}
}

func TestOrderLock(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	time.Sleep(10 * time.Second)
	t.Run("createOrderLocks", createOrderLocks)
	t.Run("getOrderLock", getOrderLock)
	t.Run("getOrderLocks", getOrderLocks)
	t.Run("deleteOrderLocks", deleteOrderLocks)
}
