package order

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	ordermgrpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order/order"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	"github.com/google/uuid"
)

func validate(info *npool.OrderReq) error { //nolint
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	if info.GoodID == nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID)
		return status.Error(codes.InvalidArgument, "GoodID is empty")
	}

	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GetGoodID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if info.PaymentCoinID == nil {
		logger.Sugar().Errorw("validate", "PaymentCoinID", info.PaymentCoinID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetPaymentCoinID()); err != nil {
		logger.Sugar().Errorw("validate", "PaymentCoinID", info.GetPaymentCoinID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentCoinID is invalid: %v", err))
	}

	if info.PaymentAccountID == nil {
		logger.Sugar().Errorw("validate", "PaymentAccountID", info.PaymentAccountID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetPaymentAccountID()); err != nil {
		logger.Sugar().Errorw("validate", "PaymentAccountID", info.GetPaymentAccountID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentAccountID is invalid: %v", err))
	}

	amount, err := decimal.NewFromString(info.GetPaymentAmount())
	if err != nil {
		logger.Sugar().Errorw("validate", "PaymentAmount", info.GetPaymentAmount(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentAmount is invalid: %v", err))
	}
	if amount.Cmp(decimal.NewFromInt(0)) < 0 {
		logger.Sugar().Errorw("validate", "PaymentAmount", info.GetPaymentAmount(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "PaymentAmount is less than 0")
	}

	amount, err = decimal.NewFromString(info.GetPayWithBalanceAmount())
	if err != nil {
		logger.Sugar().Errorw("validate", "PayWithBalanceAmount", info.GetPayWithBalanceAmount(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PayWithBalanceAmount is invalid: %v", err))
	}
	if amount.Cmp(decimal.NewFromInt(0)) < 0 {
		logger.Sugar().Errorw("validate", "PayWithBalanceAmount", info.GetPayWithBalanceAmount(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "PayWithBalanceAmount is less than 0")
	}

	amount, err = decimal.NewFromString(info.GetPaymentAccountStartAmount())
	if err != nil {
		logger.Sugar().Errorw("validate", "PaymentAccountStartAmount", info.GetPaymentAccountStartAmount(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentAccountStartAmount is invalid: %v", err))
	}
	if amount.Cmp(decimal.NewFromInt(0)) < 0 {
		logger.Sugar().Errorw("validate", "PaymentAccountStartAmount", info.GetPaymentAccountStartAmount(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "PaymentAccountStartAmount is less than 0")
	}

	currency, err := decimal.NewFromString(info.GetPaymentCoinUSDCurrency())
	if err != nil {
		logger.Sugar().Errorw("validate", "PaymentCoinUSDCurrency", info.GetPaymentCoinUSDCurrency(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentCoinUSDCurrency is invalid: %v", err))
	}
	if currency.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "PaymentCoinUSDCurrency", info.GetPaymentCoinUSDCurrency(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "PaymentCoinUSDCurrency is less than 0")
	}

	currency, err = decimal.NewFromString(info.GetPaymentLiveUSDCurrency())
	if err != nil {
		logger.Sugar().Errorw("validate", "PaymentLiveUSDCurrency", info.GetPaymentLiveUSDCurrency(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentLiveUSDCurrency is invalid: %v", err))
	}
	if currency.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "PaymentLiveUSDCurrency", info.GetPaymentLiveUSDCurrency(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "PaymentLiveUSDCurrency is less than 0")
	}

	currency, err = decimal.NewFromString(info.GetPaymentLocalUSDCurrency())
	if err != nil {
		logger.Sugar().Errorw("validate", "PaymentLocalUSDCurrency", info.GetPaymentLocalUSDCurrency(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentLocalUSDCurrency is invalid: %v", err))
	}
	if currency.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "PaymentLocalUSDCurrency", info.GetPaymentLocalUSDCurrency(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "PaymentLocalUSDCurrency is less than 0")
	}

	if info.OrderType == nil {
		logger.Sugar().Errorw("validate", "OrderType", info.OrderType)
		return status.Error(codes.InvalidArgument, "OrderType is empty")
	}

	switch info.GetOrderType() {
	case ordermgrpb.OrderType_Normal:
	case ordermgrpb.OrderType_Offline:
	case ordermgrpb.OrderType_Airdrop:
	default:
		logger.Sugar().Errorw("validate", "OrderType", info.GetOrderType())
		return status.Error(codes.InvalidArgument, "OrderType is invalid")
	}

	if info.FixAmountID != nil {
		if _, err := uuid.Parse(info.GetFixAmountID()); err != nil {
			logger.Sugar().Errorw("validate", "FixAmountID", info.GetFixAmountID(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("FixAmountID is invalid: %v", err))
		}
	}

	if info.DiscountID != nil {
		if _, err := uuid.Parse(info.GetDiscountID()); err != nil {
			logger.Sugar().Errorw("validate", "DiscountID", info.GetDiscountID(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("DiscountID is invalid: %v", err))
		}
	}

	if info.SpecialOfferID != nil {
		if _, err := uuid.Parse(info.GetSpecialOfferID()); err != nil {
			logger.Sugar().Errorw("validate", "SpecialOfferID", info.GetSpecialOfferID(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("SpecialOfferID is invalid: %v", err))
		}
	}

	if info.PromotionID != nil {
		if _, err := uuid.Parse(info.GetPromotionID()); err != nil {
			logger.Sugar().Errorw("validate", "PromotionID", info.GetPromotionID(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("PromotionID is invalid: %v", err))
		}
	}

	return nil
}
