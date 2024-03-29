// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/schema"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/simulateconfig"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	compensateMixin := schema.Compensate{}.Mixin()
	compensate.Policy = privacy.NewPolicies(compensateMixin[0], schema.Compensate{})
	compensate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := compensate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	compensateMixinFields0 := compensateMixin[0].Fields()
	_ = compensateMixinFields0
	compensateMixinFields1 := compensateMixin[1].Fields()
	_ = compensateMixinFields1
	compensateFields := schema.Compensate{}.Fields()
	_ = compensateFields
	// compensateDescCreatedAt is the schema descriptor for created_at field.
	compensateDescCreatedAt := compensateMixinFields0[0].Descriptor()
	// compensate.DefaultCreatedAt holds the default value on creation for the created_at field.
	compensate.DefaultCreatedAt = compensateDescCreatedAt.Default.(func() uint32)
	// compensateDescUpdatedAt is the schema descriptor for updated_at field.
	compensateDescUpdatedAt := compensateMixinFields0[1].Descriptor()
	// compensate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	compensate.DefaultUpdatedAt = compensateDescUpdatedAt.Default.(func() uint32)
	// compensate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	compensate.UpdateDefaultUpdatedAt = compensateDescUpdatedAt.UpdateDefault.(func() uint32)
	// compensateDescDeletedAt is the schema descriptor for deleted_at field.
	compensateDescDeletedAt := compensateMixinFields0[2].Descriptor()
	// compensate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	compensate.DefaultDeletedAt = compensateDescDeletedAt.Default.(func() uint32)
	// compensateDescEntID is the schema descriptor for ent_id field.
	compensateDescEntID := compensateMixinFields1[1].Descriptor()
	// compensate.DefaultEntID holds the default value on creation for the ent_id field.
	compensate.DefaultEntID = compensateDescEntID.Default.(func() uuid.UUID)
	// compensateDescStartAt is the schema descriptor for start_at field.
	compensateDescStartAt := compensateFields[1].Descriptor()
	// compensate.DefaultStartAt holds the default value on creation for the start_at field.
	compensate.DefaultStartAt = compensateDescStartAt.Default.(uint32)
	// compensateDescEndAt is the schema descriptor for end_at field.
	compensateDescEndAt := compensateFields[2].Descriptor()
	// compensate.DefaultEndAt holds the default value on creation for the end_at field.
	compensate.DefaultEndAt = compensateDescEndAt.Default.(uint32)
	// compensateDescCompensateType is the schema descriptor for compensate_type field.
	compensateDescCompensateType := compensateFields[3].Descriptor()
	// compensate.DefaultCompensateType holds the default value on creation for the compensate_type field.
	compensate.DefaultCompensateType = compensateDescCompensateType.Default.(string)
	// compensateDescTitle is the schema descriptor for title field.
	compensateDescTitle := compensateFields[4].Descriptor()
	// compensate.DefaultTitle holds the default value on creation for the title field.
	compensate.DefaultTitle = compensateDescTitle.Default.(string)
	// compensateDescMessage is the schema descriptor for message field.
	compensateDescMessage := compensateFields[5].Descriptor()
	// compensate.DefaultMessage holds the default value on creation for the message field.
	compensate.DefaultMessage = compensateDescMessage.Default.(string)
	entorderMixin := schema.Order{}.Mixin()
	entorder.Policy = privacy.NewPolicies(entorderMixin[0], schema.Order{})
	entorder.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := entorder.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	entorderMixinFields0 := entorderMixin[0].Fields()
	_ = entorderMixinFields0
	entorderMixinFields1 := entorderMixin[1].Fields()
	_ = entorderMixinFields1
	entorderFields := schema.Order{}.Fields()
	_ = entorderFields
	// entorderDescCreatedAt is the schema descriptor for created_at field.
	entorderDescCreatedAt := entorderMixinFields0[0].Descriptor()
	// entorder.DefaultCreatedAt holds the default value on creation for the created_at field.
	entorder.DefaultCreatedAt = entorderDescCreatedAt.Default.(func() uint32)
	// entorderDescUpdatedAt is the schema descriptor for updated_at field.
	entorderDescUpdatedAt := entorderMixinFields0[1].Descriptor()
	// entorder.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	entorder.DefaultUpdatedAt = entorderDescUpdatedAt.Default.(func() uint32)
	// entorder.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	entorder.UpdateDefaultUpdatedAt = entorderDescUpdatedAt.UpdateDefault.(func() uint32)
	// entorderDescDeletedAt is the schema descriptor for deleted_at field.
	entorderDescDeletedAt := entorderMixinFields0[2].Descriptor()
	// entorder.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	entorder.DefaultDeletedAt = entorderDescDeletedAt.Default.(func() uint32)
	// entorderDescEntID is the schema descriptor for ent_id field.
	entorderDescEntID := entorderMixinFields1[1].Descriptor()
	// entorder.DefaultEntID holds the default value on creation for the ent_id field.
	entorder.DefaultEntID = entorderDescEntID.Default.(func() uuid.UUID)
	// entorderDescPaymentID is the schema descriptor for payment_id field.
	entorderDescPaymentID := entorderFields[4].Descriptor()
	// entorder.DefaultPaymentID holds the default value on creation for the payment_id field.
	entorder.DefaultPaymentID = entorderDescPaymentID.Default.(func() uuid.UUID)
	// entorderDescParentOrderID is the schema descriptor for parent_order_id field.
	entorderDescParentOrderID := entorderFields[5].Descriptor()
	// entorder.DefaultParentOrderID holds the default value on creation for the parent_order_id field.
	entorder.DefaultParentOrderID = entorderDescParentOrderID.Default.(func() uuid.UUID)
	// entorderDescUnitsV1 is the schema descriptor for units_v1 field.
	entorderDescUnitsV1 := entorderFields[6].Descriptor()
	// entorder.DefaultUnitsV1 holds the default value on creation for the units_v1 field.
	entorder.DefaultUnitsV1 = entorderDescUnitsV1.Default.(decimal.Decimal)
	// entorderDescGoodValue is the schema descriptor for good_value field.
	entorderDescGoodValue := entorderFields[7].Descriptor()
	// entorder.DefaultGoodValue holds the default value on creation for the good_value field.
	entorder.DefaultGoodValue = entorderDescGoodValue.Default.(decimal.Decimal)
	// entorderDescGoodValueUsd is the schema descriptor for good_value_usd field.
	entorderDescGoodValueUsd := entorderFields[8].Descriptor()
	// entorder.DefaultGoodValueUsd holds the default value on creation for the good_value_usd field.
	entorder.DefaultGoodValueUsd = entorderDescGoodValueUsd.Default.(decimal.Decimal)
	// entorderDescPaymentAmount is the schema descriptor for payment_amount field.
	entorderDescPaymentAmount := entorderFields[9].Descriptor()
	// entorder.DefaultPaymentAmount holds the default value on creation for the payment_amount field.
	entorder.DefaultPaymentAmount = entorderDescPaymentAmount.Default.(decimal.Decimal)
	// entorderDescDiscountAmount is the schema descriptor for discount_amount field.
	entorderDescDiscountAmount := entorderFields[10].Descriptor()
	// entorder.DefaultDiscountAmount holds the default value on creation for the discount_amount field.
	entorder.DefaultDiscountAmount = entorderDescDiscountAmount.Default.(decimal.Decimal)
	// entorderDescPromotionID is the schema descriptor for promotion_id field.
	entorderDescPromotionID := entorderFields[11].Descriptor()
	// entorder.DefaultPromotionID holds the default value on creation for the promotion_id field.
	entorder.DefaultPromotionID = entorderDescPromotionID.Default.(func() uuid.UUID)
	// entorderDescDuration is the schema descriptor for duration field.
	entorderDescDuration := entorderFields[12].Descriptor()
	// entorder.DefaultDuration holds the default value on creation for the duration field.
	entorder.DefaultDuration = entorderDescDuration.Default.(uint32)
	// entorderDescOrderType is the schema descriptor for order_type field.
	entorderDescOrderType := entorderFields[13].Descriptor()
	// entorder.DefaultOrderType holds the default value on creation for the order_type field.
	entorder.DefaultOrderType = entorderDescOrderType.Default.(string)
	// entorderDescInvestmentType is the schema descriptor for investment_type field.
	entorderDescInvestmentType := entorderFields[14].Descriptor()
	// entorder.DefaultInvestmentType holds the default value on creation for the investment_type field.
	entorder.DefaultInvestmentType = entorderDescInvestmentType.Default.(string)
	// entorderDescCouponIds is the schema descriptor for coupon_ids field.
	entorderDescCouponIds := entorderFields[15].Descriptor()
	// entorder.DefaultCouponIds holds the default value on creation for the coupon_ids field.
	entorder.DefaultCouponIds = entorderDescCouponIds.Default.(func() []uuid.UUID)
	// entorderDescPaymentType is the schema descriptor for payment_type field.
	entorderDescPaymentType := entorderFields[16].Descriptor()
	// entorder.DefaultPaymentType holds the default value on creation for the payment_type field.
	entorder.DefaultPaymentType = entorderDescPaymentType.Default.(string)
	// entorderDescPaymentCoinTypeID is the schema descriptor for payment_coin_type_id field.
	entorderDescPaymentCoinTypeID := entorderFields[18].Descriptor()
	// entorder.DefaultPaymentCoinTypeID holds the default value on creation for the payment_coin_type_id field.
	entorder.DefaultPaymentCoinTypeID = entorderDescPaymentCoinTypeID.Default.(func() uuid.UUID)
	// entorderDescTransferAmount is the schema descriptor for transfer_amount field.
	entorderDescTransferAmount := entorderFields[19].Descriptor()
	// entorder.DefaultTransferAmount holds the default value on creation for the transfer_amount field.
	entorder.DefaultTransferAmount = entorderDescTransferAmount.Default.(decimal.Decimal)
	// entorderDescBalanceAmount is the schema descriptor for balance_amount field.
	entorderDescBalanceAmount := entorderFields[20].Descriptor()
	// entorder.DefaultBalanceAmount holds the default value on creation for the balance_amount field.
	entorder.DefaultBalanceAmount = entorderDescBalanceAmount.Default.(decimal.Decimal)
	// entorderDescCoinUsdCurrency is the schema descriptor for coin_usd_currency field.
	entorderDescCoinUsdCurrency := entorderFields[21].Descriptor()
	// entorder.DefaultCoinUsdCurrency holds the default value on creation for the coin_usd_currency field.
	entorder.DefaultCoinUsdCurrency = entorderDescCoinUsdCurrency.Default.(decimal.Decimal)
	// entorderDescLocalCoinUsdCurrency is the schema descriptor for local_coin_usd_currency field.
	entorderDescLocalCoinUsdCurrency := entorderFields[22].Descriptor()
	// entorder.DefaultLocalCoinUsdCurrency holds the default value on creation for the local_coin_usd_currency field.
	entorder.DefaultLocalCoinUsdCurrency = entorderDescLocalCoinUsdCurrency.Default.(decimal.Decimal)
	// entorderDescLiveCoinUsdCurrency is the schema descriptor for live_coin_usd_currency field.
	entorderDescLiveCoinUsdCurrency := entorderFields[23].Descriptor()
	// entorder.DefaultLiveCoinUsdCurrency holds the default value on creation for the live_coin_usd_currency field.
	entorder.DefaultLiveCoinUsdCurrency = entorderDescLiveCoinUsdCurrency.Default.(decimal.Decimal)
	// entorderDescSimulate is the schema descriptor for simulate field.
	entorderDescSimulate := entorderFields[24].Descriptor()
	// entorder.DefaultSimulate holds the default value on creation for the simulate field.
	entorder.DefaultSimulate = entorderDescSimulate.Default.(bool)
	// entorderDescCreateMethod is the schema descriptor for create_method field.
	entorderDescCreateMethod := entorderFields[25].Descriptor()
	// entorder.DefaultCreateMethod holds the default value on creation for the create_method field.
	entorder.DefaultCreateMethod = entorderDescCreateMethod.Default.(string)
	// entorderDescMultiPaymentCoins is the schema descriptor for multi_payment_coins field.
	entorderDescMultiPaymentCoins := entorderFields[26].Descriptor()
	// entorder.DefaultMultiPaymentCoins holds the default value on creation for the multi_payment_coins field.
	entorder.DefaultMultiPaymentCoins = entorderDescMultiPaymentCoins.Default.(bool)
	// entorderDescPaymentAmounts is the schema descriptor for payment_amounts field.
	entorderDescPaymentAmounts := entorderFields[27].Descriptor()
	// entorder.DefaultPaymentAmounts holds the default value on creation for the payment_amounts field.
	entorder.DefaultPaymentAmounts = entorderDescPaymentAmounts.Default.([]order.PaymentAmount)
	orderlockMixin := schema.OrderLock{}.Mixin()
	orderlock.Policy = privacy.NewPolicies(orderlockMixin[0], schema.OrderLock{})
	orderlock.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := orderlock.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	orderlockMixinFields0 := orderlockMixin[0].Fields()
	_ = orderlockMixinFields0
	orderlockMixinFields1 := orderlockMixin[1].Fields()
	_ = orderlockMixinFields1
	orderlockFields := schema.OrderLock{}.Fields()
	_ = orderlockFields
	// orderlockDescCreatedAt is the schema descriptor for created_at field.
	orderlockDescCreatedAt := orderlockMixinFields0[0].Descriptor()
	// orderlock.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderlock.DefaultCreatedAt = orderlockDescCreatedAt.Default.(func() uint32)
	// orderlockDescUpdatedAt is the schema descriptor for updated_at field.
	orderlockDescUpdatedAt := orderlockMixinFields0[1].Descriptor()
	// orderlock.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderlock.DefaultUpdatedAt = orderlockDescUpdatedAt.Default.(func() uint32)
	// orderlock.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderlock.UpdateDefaultUpdatedAt = orderlockDescUpdatedAt.UpdateDefault.(func() uint32)
	// orderlockDescDeletedAt is the schema descriptor for deleted_at field.
	orderlockDescDeletedAt := orderlockMixinFields0[2].Descriptor()
	// orderlock.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	orderlock.DefaultDeletedAt = orderlockDescDeletedAt.Default.(func() uint32)
	// orderlockDescEntID is the schema descriptor for ent_id field.
	orderlockDescEntID := orderlockMixinFields1[1].Descriptor()
	// orderlock.DefaultEntID holds the default value on creation for the ent_id field.
	orderlock.DefaultEntID = orderlockDescEntID.Default.(func() uuid.UUID)
	// orderlockDescLockType is the schema descriptor for lock_type field.
	orderlockDescLockType := orderlockFields[3].Descriptor()
	// orderlock.DefaultLockType holds the default value on creation for the lock_type field.
	orderlock.DefaultLockType = orderlockDescLockType.Default.(string)
	orderstateMixin := schema.OrderState{}.Mixin()
	orderstate.Policy = privacy.NewPolicies(orderstateMixin[0], schema.OrderState{})
	orderstate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := orderstate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	orderstateMixinFields0 := orderstateMixin[0].Fields()
	_ = orderstateMixinFields0
	orderstateMixinFields1 := orderstateMixin[1].Fields()
	_ = orderstateMixinFields1
	orderstateFields := schema.OrderState{}.Fields()
	_ = orderstateFields
	// orderstateDescCreatedAt is the schema descriptor for created_at field.
	orderstateDescCreatedAt := orderstateMixinFields0[0].Descriptor()
	// orderstate.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderstate.DefaultCreatedAt = orderstateDescCreatedAt.Default.(func() uint32)
	// orderstateDescUpdatedAt is the schema descriptor for updated_at field.
	orderstateDescUpdatedAt := orderstateMixinFields0[1].Descriptor()
	// orderstate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderstate.DefaultUpdatedAt = orderstateDescUpdatedAt.Default.(func() uint32)
	// orderstate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderstate.UpdateDefaultUpdatedAt = orderstateDescUpdatedAt.UpdateDefault.(func() uint32)
	// orderstateDescDeletedAt is the schema descriptor for deleted_at field.
	orderstateDescDeletedAt := orderstateMixinFields0[2].Descriptor()
	// orderstate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	orderstate.DefaultDeletedAt = orderstateDescDeletedAt.Default.(func() uint32)
	// orderstateDescEntID is the schema descriptor for ent_id field.
	orderstateDescEntID := orderstateMixinFields1[1].Descriptor()
	// orderstate.DefaultEntID holds the default value on creation for the ent_id field.
	orderstate.DefaultEntID = orderstateDescEntID.Default.(func() uuid.UUID)
	// orderstateDescOrderState is the schema descriptor for order_state field.
	orderstateDescOrderState := orderstateFields[1].Descriptor()
	// orderstate.DefaultOrderState holds the default value on creation for the order_state field.
	orderstate.DefaultOrderState = orderstateDescOrderState.Default.(string)
	// orderstateDescCancelState is the schema descriptor for cancel_state field.
	orderstateDescCancelState := orderstateFields[2].Descriptor()
	// orderstate.DefaultCancelState holds the default value on creation for the cancel_state field.
	orderstate.DefaultCancelState = orderstateDescCancelState.Default.(string)
	// orderstateDescStartMode is the schema descriptor for start_mode field.
	orderstateDescStartMode := orderstateFields[3].Descriptor()
	// orderstate.DefaultStartMode holds the default value on creation for the start_mode field.
	orderstate.DefaultStartMode = orderstateDescStartMode.Default.(string)
	// orderstateDescStartAt is the schema descriptor for start_at field.
	orderstateDescStartAt := orderstateFields[4].Descriptor()
	// orderstate.DefaultStartAt holds the default value on creation for the start_at field.
	orderstate.DefaultStartAt = orderstateDescStartAt.Default.(uint32)
	// orderstateDescEndAt is the schema descriptor for end_at field.
	orderstateDescEndAt := orderstateFields[5].Descriptor()
	// orderstate.DefaultEndAt holds the default value on creation for the end_at field.
	orderstate.DefaultEndAt = orderstateDescEndAt.Default.(uint32)
	// orderstateDescPaidAt is the schema descriptor for paid_at field.
	orderstateDescPaidAt := orderstateFields[6].Descriptor()
	// orderstate.DefaultPaidAt holds the default value on creation for the paid_at field.
	orderstate.DefaultPaidAt = orderstateDescPaidAt.Default.(uint32)
	// orderstateDescLastBenefitAt is the schema descriptor for last_benefit_at field.
	orderstateDescLastBenefitAt := orderstateFields[7].Descriptor()
	// orderstate.DefaultLastBenefitAt holds the default value on creation for the last_benefit_at field.
	orderstate.DefaultLastBenefitAt = orderstateDescLastBenefitAt.Default.(uint32)
	// orderstateDescBenefitState is the schema descriptor for benefit_state field.
	orderstateDescBenefitState := orderstateFields[8].Descriptor()
	// orderstate.DefaultBenefitState holds the default value on creation for the benefit_state field.
	orderstate.DefaultBenefitState = orderstateDescBenefitState.Default.(string)
	// orderstateDescUserSetPaid is the schema descriptor for user_set_paid field.
	orderstateDescUserSetPaid := orderstateFields[9].Descriptor()
	// orderstate.DefaultUserSetPaid holds the default value on creation for the user_set_paid field.
	orderstate.DefaultUserSetPaid = orderstateDescUserSetPaid.Default.(bool)
	// orderstateDescUserSetCanceled is the schema descriptor for user_set_canceled field.
	orderstateDescUserSetCanceled := orderstateFields[10].Descriptor()
	// orderstate.DefaultUserSetCanceled holds the default value on creation for the user_set_canceled field.
	orderstate.DefaultUserSetCanceled = orderstateDescUserSetCanceled.Default.(bool)
	// orderstateDescAdminSetCanceled is the schema descriptor for admin_set_canceled field.
	orderstateDescAdminSetCanceled := orderstateFields[11].Descriptor()
	// orderstate.DefaultAdminSetCanceled holds the default value on creation for the admin_set_canceled field.
	orderstate.DefaultAdminSetCanceled = orderstateDescAdminSetCanceled.Default.(bool)
	// orderstateDescPaymentTransactionID is the schema descriptor for payment_transaction_id field.
	orderstateDescPaymentTransactionID := orderstateFields[12].Descriptor()
	// orderstate.DefaultPaymentTransactionID holds the default value on creation for the payment_transaction_id field.
	orderstate.DefaultPaymentTransactionID = orderstateDescPaymentTransactionID.Default.(string)
	// orderstateDescPaymentFinishAmount is the schema descriptor for payment_finish_amount field.
	orderstateDescPaymentFinishAmount := orderstateFields[13].Descriptor()
	// orderstate.DefaultPaymentFinishAmount holds the default value on creation for the payment_finish_amount field.
	orderstate.DefaultPaymentFinishAmount = orderstateDescPaymentFinishAmount.Default.(decimal.Decimal)
	// orderstateDescPaymentState is the schema descriptor for payment_state field.
	orderstateDescPaymentState := orderstateFields[14].Descriptor()
	// orderstate.DefaultPaymentState holds the default value on creation for the payment_state field.
	orderstate.DefaultPaymentState = orderstateDescPaymentState.Default.(string)
	// orderstateDescOutofgasHours is the schema descriptor for outofgas_hours field.
	orderstateDescOutofgasHours := orderstateFields[15].Descriptor()
	// orderstate.DefaultOutofgasHours holds the default value on creation for the outofgas_hours field.
	orderstate.DefaultOutofgasHours = orderstateDescOutofgasHours.Default.(uint32)
	// orderstateDescCompensateHours is the schema descriptor for compensate_hours field.
	orderstateDescCompensateHours := orderstateFields[16].Descriptor()
	// orderstate.DefaultCompensateHours holds the default value on creation for the compensate_hours field.
	orderstate.DefaultCompensateHours = orderstateDescCompensateHours.Default.(uint32)
	// orderstateDescRenewState is the schema descriptor for renew_state field.
	orderstateDescRenewState := orderstateFields[17].Descriptor()
	// orderstate.DefaultRenewState holds the default value on creation for the renew_state field.
	orderstate.DefaultRenewState = orderstateDescRenewState.Default.(string)
	// orderstateDescRenewNotifyAt is the schema descriptor for renew_notify_at field.
	orderstateDescRenewNotifyAt := orderstateFields[18].Descriptor()
	// orderstate.DefaultRenewNotifyAt holds the default value on creation for the renew_notify_at field.
	orderstate.DefaultRenewNotifyAt = orderstateDescRenewNotifyAt.Default.(uint32)
	outofgasMixin := schema.OutOfGas{}.Mixin()
	outofgas.Policy = privacy.NewPolicies(outofgasMixin[0], schema.OutOfGas{})
	outofgas.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := outofgas.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	outofgasMixinFields0 := outofgasMixin[0].Fields()
	_ = outofgasMixinFields0
	outofgasMixinFields1 := outofgasMixin[1].Fields()
	_ = outofgasMixinFields1
	outofgasFields := schema.OutOfGas{}.Fields()
	_ = outofgasFields
	// outofgasDescCreatedAt is the schema descriptor for created_at field.
	outofgasDescCreatedAt := outofgasMixinFields0[0].Descriptor()
	// outofgas.DefaultCreatedAt holds the default value on creation for the created_at field.
	outofgas.DefaultCreatedAt = outofgasDescCreatedAt.Default.(func() uint32)
	// outofgasDescUpdatedAt is the schema descriptor for updated_at field.
	outofgasDescUpdatedAt := outofgasMixinFields0[1].Descriptor()
	// outofgas.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	outofgas.DefaultUpdatedAt = outofgasDescUpdatedAt.Default.(func() uint32)
	// outofgas.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	outofgas.UpdateDefaultUpdatedAt = outofgasDescUpdatedAt.UpdateDefault.(func() uint32)
	// outofgasDescDeletedAt is the schema descriptor for deleted_at field.
	outofgasDescDeletedAt := outofgasMixinFields0[2].Descriptor()
	// outofgas.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	outofgas.DefaultDeletedAt = outofgasDescDeletedAt.Default.(func() uint32)
	// outofgasDescEntID is the schema descriptor for ent_id field.
	outofgasDescEntID := outofgasMixinFields1[1].Descriptor()
	// outofgas.DefaultEntID holds the default value on creation for the ent_id field.
	outofgas.DefaultEntID = outofgasDescEntID.Default.(func() uuid.UUID)
	// outofgasDescStartAt is the schema descriptor for start_at field.
	outofgasDescStartAt := outofgasFields[1].Descriptor()
	// outofgas.DefaultStartAt holds the default value on creation for the start_at field.
	outofgas.DefaultStartAt = outofgasDescStartAt.Default.(uint32)
	// outofgasDescEndAt is the schema descriptor for end_at field.
	outofgasDescEndAt := outofgasFields[2].Descriptor()
	// outofgas.DefaultEndAt holds the default value on creation for the end_at field.
	outofgas.DefaultEndAt = outofgasDescEndAt.Default.(uint32)
	paymentMixin := schema.Payment{}.Mixin()
	payment.Policy = privacy.NewPolicies(paymentMixin[0], schema.Payment{})
	payment.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := payment.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	paymentMixinFields0 := paymentMixin[0].Fields()
	_ = paymentMixinFields0
	paymentMixinFields1 := paymentMixin[1].Fields()
	_ = paymentMixinFields1
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescCreatedAt is the schema descriptor for created_at field.
	paymentDescCreatedAt := paymentMixinFields0[0].Descriptor()
	// payment.DefaultCreatedAt holds the default value on creation for the created_at field.
	payment.DefaultCreatedAt = paymentDescCreatedAt.Default.(func() uint32)
	// paymentDescUpdatedAt is the schema descriptor for updated_at field.
	paymentDescUpdatedAt := paymentMixinFields0[1].Descriptor()
	// payment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	payment.DefaultUpdatedAt = paymentDescUpdatedAt.Default.(func() uint32)
	// payment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	payment.UpdateDefaultUpdatedAt = paymentDescUpdatedAt.UpdateDefault.(func() uint32)
	// paymentDescDeletedAt is the schema descriptor for deleted_at field.
	paymentDescDeletedAt := paymentMixinFields0[2].Descriptor()
	// payment.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	payment.DefaultDeletedAt = paymentDescDeletedAt.Default.(func() uint32)
	// paymentDescEntID is the schema descriptor for ent_id field.
	paymentDescEntID := paymentMixinFields1[1].Descriptor()
	// payment.DefaultEntID holds the default value on creation for the ent_id field.
	payment.DefaultEntID = paymentDescEntID.Default.(func() uuid.UUID)
	// paymentDescStartAmount is the schema descriptor for start_amount field.
	paymentDescStartAmount := paymentFields[7].Descriptor()
	// payment.DefaultStartAmount holds the default value on creation for the start_amount field.
	payment.DefaultStartAmount = paymentDescStartAmount.Default.(decimal.Decimal)
	simulateconfigMixin := schema.SimulateConfig{}.Mixin()
	simulateconfig.Policy = privacy.NewPolicies(simulateconfigMixin[0], schema.SimulateConfig{})
	simulateconfig.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := simulateconfig.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	simulateconfigMixinFields0 := simulateconfigMixin[0].Fields()
	_ = simulateconfigMixinFields0
	simulateconfigMixinFields1 := simulateconfigMixin[1].Fields()
	_ = simulateconfigMixinFields1
	simulateconfigFields := schema.SimulateConfig{}.Fields()
	_ = simulateconfigFields
	// simulateconfigDescCreatedAt is the schema descriptor for created_at field.
	simulateconfigDescCreatedAt := simulateconfigMixinFields0[0].Descriptor()
	// simulateconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	simulateconfig.DefaultCreatedAt = simulateconfigDescCreatedAt.Default.(func() uint32)
	// simulateconfigDescUpdatedAt is the schema descriptor for updated_at field.
	simulateconfigDescUpdatedAt := simulateconfigMixinFields0[1].Descriptor()
	// simulateconfig.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	simulateconfig.DefaultUpdatedAt = simulateconfigDescUpdatedAt.Default.(func() uint32)
	// simulateconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	simulateconfig.UpdateDefaultUpdatedAt = simulateconfigDescUpdatedAt.UpdateDefault.(func() uint32)
	// simulateconfigDescDeletedAt is the schema descriptor for deleted_at field.
	simulateconfigDescDeletedAt := simulateconfigMixinFields0[2].Descriptor()
	// simulateconfig.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	simulateconfig.DefaultDeletedAt = simulateconfigDescDeletedAt.Default.(func() uint32)
	// simulateconfigDescEntID is the schema descriptor for ent_id field.
	simulateconfigDescEntID := simulateconfigMixinFields1[1].Descriptor()
	// simulateconfig.DefaultEntID holds the default value on creation for the ent_id field.
	simulateconfig.DefaultEntID = simulateconfigDescEntID.Default.(func() uuid.UUID)
	// simulateconfigDescSendCouponMode is the schema descriptor for send_coupon_mode field.
	simulateconfigDescSendCouponMode := simulateconfigFields[1].Descriptor()
	// simulateconfig.DefaultSendCouponMode holds the default value on creation for the send_coupon_mode field.
	simulateconfig.DefaultSendCouponMode = simulateconfigDescSendCouponMode.Default.(string)
	// simulateconfigDescSendCouponProbability is the schema descriptor for send_coupon_probability field.
	simulateconfigDescSendCouponProbability := simulateconfigFields[2].Descriptor()
	// simulateconfig.DefaultSendCouponProbability holds the default value on creation for the send_coupon_probability field.
	simulateconfig.DefaultSendCouponProbability = simulateconfigDescSendCouponProbability.Default.(decimal.Decimal)
	// simulateconfigDescCashableProfitProbability is the schema descriptor for cashable_profit_probability field.
	simulateconfigDescCashableProfitProbability := simulateconfigFields[3].Descriptor()
	// simulateconfig.DefaultCashableProfitProbability holds the default value on creation for the cashable_profit_probability field.
	simulateconfig.DefaultCashableProfitProbability = simulateconfigDescCashableProfitProbability.Default.(decimal.Decimal)
	// simulateconfigDescEnabled is the schema descriptor for enabled field.
	simulateconfigDescEnabled := simulateconfigFields[4].Descriptor()
	// simulateconfig.DefaultEnabled holds the default value on creation for the enabled field.
	simulateconfig.DefaultEnabled = simulateconfigDescEnabled.Default.(bool)
}

const (
	Version = "v0.11.3" // Version of ent codegen.
)
