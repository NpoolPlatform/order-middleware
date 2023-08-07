// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/schema"
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
	// compensateDescStart is the schema descriptor for start field.
	compensateDescStart := compensateFields[2].Descriptor()
	// compensate.DefaultStart holds the default value on creation for the start field.
	compensate.DefaultStart = compensateDescStart.Default.(uint32)
	// compensateDescEnd is the schema descriptor for end field.
	compensateDescEnd := compensateFields[3].Descriptor()
	// compensate.DefaultEnd holds the default value on creation for the end field.
	compensate.DefaultEnd = compensateDescEnd.Default.(uint32)
	// compensateDescMessage is the schema descriptor for message field.
	compensateDescMessage := compensateFields[4].Descriptor()
	// compensate.DefaultMessage holds the default value on creation for the message field.
	compensate.DefaultMessage = compensateDescMessage.Default.(string)
	// compensateDescID is the schema descriptor for id field.
	compensateDescID := compensateFields[0].Descriptor()
	// compensate.DefaultID holds the default value on creation for the id field.
	compensate.DefaultID = compensateDescID.Default.(func() uuid.UUID)
	orderMixin := schema.Order{}.Mixin()
	order.Policy = privacy.NewPolicies(orderMixin[0], schema.Order{})
	order.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := order.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderMixinFields0[0].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() uint32)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderMixinFields0[1].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() uint32)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() uint32)
	// orderDescDeletedAt is the schema descriptor for deleted_at field.
	orderDescDeletedAt := orderMixinFields0[2].Descriptor()
	// order.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	order.DefaultDeletedAt = orderDescDeletedAt.Default.(func() uint32)
	// orderDescParentOrderID is the schema descriptor for parent_order_id field.
	orderDescParentOrderID := orderFields[4].Descriptor()
	// order.DefaultParentOrderID holds the default value on creation for the parent_order_id field.
	order.DefaultParentOrderID = orderDescParentOrderID.Default.(func() uuid.UUID)
	// orderDescPayWithParent is the schema descriptor for pay_with_parent field.
	orderDescPayWithParent := orderFields[5].Descriptor()
	// order.DefaultPayWithParent holds the default value on creation for the pay_with_parent field.
	order.DefaultPayWithParent = orderDescPayWithParent.Default.(bool)
	// orderDescUnits is the schema descriptor for units field.
	orderDescUnits := orderFields[6].Descriptor()
	// order.DefaultUnits holds the default value on creation for the units field.
	order.DefaultUnits = orderDescUnits.Default.(uint32)
	// orderDescUnitsV1 is the schema descriptor for units_v1 field.
	orderDescUnitsV1 := orderFields[7].Descriptor()
	// order.DefaultUnitsV1 holds the default value on creation for the units_v1 field.
	order.DefaultUnitsV1 = orderDescUnitsV1.Default.(decimal.Decimal)
	// orderDescPromotionID is the schema descriptor for promotion_id field.
	orderDescPromotionID := orderFields[8].Descriptor()
	// order.DefaultPromotionID holds the default value on creation for the promotion_id field.
	order.DefaultPromotionID = orderDescPromotionID.Default.(func() uuid.UUID)
	// orderDescDiscountCouponID is the schema descriptor for discount_coupon_id field.
	orderDescDiscountCouponID := orderFields[9].Descriptor()
	// order.DefaultDiscountCouponID holds the default value on creation for the discount_coupon_id field.
	order.DefaultDiscountCouponID = orderDescDiscountCouponID.Default.(func() uuid.UUID)
	// orderDescUserSpecialReductionID is the schema descriptor for user_special_reduction_id field.
	orderDescUserSpecialReductionID := orderFields[10].Descriptor()
	// order.DefaultUserSpecialReductionID holds the default value on creation for the user_special_reduction_id field.
	order.DefaultUserSpecialReductionID = orderDescUserSpecialReductionID.Default.(func() uuid.UUID)
	// orderDescStartAt is the schema descriptor for start_at field.
	orderDescStartAt := orderFields[11].Descriptor()
	// order.DefaultStartAt holds the default value on creation for the start_at field.
	order.DefaultStartAt = orderDescStartAt.Default.(uint32)
	// orderDescEndAt is the schema descriptor for end_at field.
	orderDescEndAt := orderFields[12].Descriptor()
	// order.DefaultEndAt holds the default value on creation for the end_at field.
	order.DefaultEndAt = orderDescEndAt.Default.(uint32)
	// orderDescFixAmountCouponID is the schema descriptor for fix_amount_coupon_id field.
	orderDescFixAmountCouponID := orderFields[13].Descriptor()
	// order.DefaultFixAmountCouponID holds the default value on creation for the fix_amount_coupon_id field.
	order.DefaultFixAmountCouponID = orderDescFixAmountCouponID.Default.(func() uuid.UUID)
	// orderDescType is the schema descriptor for type field.
	orderDescType := orderFields[14].Descriptor()
	// order.DefaultType holds the default value on creation for the type field.
	order.DefaultType = orderDescType.Default.(string)
	// orderDescState is the schema descriptor for state field.
	orderDescState := orderFields[15].Descriptor()
	// order.DefaultState holds the default value on creation for the state field.
	order.DefaultState = orderDescState.Default.(string)
	// orderDescCouponIds is the schema descriptor for coupon_ids field.
	orderDescCouponIds := orderFields[16].Descriptor()
	// order.DefaultCouponIds holds the default value on creation for the coupon_ids field.
	order.DefaultCouponIds = orderDescCouponIds.Default.(func() []uuid.UUID)
	// orderDescLastBenefitAt is the schema descriptor for last_benefit_at field.
	orderDescLastBenefitAt := orderFields[17].Descriptor()
	// order.DefaultLastBenefitAt holds the default value on creation for the last_benefit_at field.
	order.DefaultLastBenefitAt = orderDescLastBenefitAt.Default.(uint32)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderFields[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() uuid.UUID)
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
	// outofgasDescStart is the schema descriptor for start field.
	outofgasDescStart := outofgasFields[2].Descriptor()
	// outofgas.DefaultStart holds the default value on creation for the start field.
	outofgas.DefaultStart = outofgasDescStart.Default.(uint32)
	// outofgasDescEnd is the schema descriptor for end field.
	outofgasDescEnd := outofgasFields[3].Descriptor()
	// outofgas.DefaultEnd holds the default value on creation for the end field.
	outofgas.DefaultEnd = outofgasDescEnd.Default.(uint32)
	// outofgasDescID is the schema descriptor for id field.
	outofgasDescID := outofgasFields[0].Descriptor()
	// outofgas.DefaultID holds the default value on creation for the id field.
	outofgas.DefaultID = outofgasDescID.Default.(func() uuid.UUID)
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
	// paymentDescStartAmount is the schema descriptor for start_amount field.
	paymentDescStartAmount := paymentFields[6].Descriptor()
	// payment.DefaultStartAmount holds the default value on creation for the start_amount field.
	payment.DefaultStartAmount = paymentDescStartAmount.Default.(decimal.Decimal)
	// paymentDescAmount is the schema descriptor for amount field.
	paymentDescAmount := paymentFields[7].Descriptor()
	// payment.DefaultAmount holds the default value on creation for the amount field.
	payment.DefaultAmount = paymentDescAmount.Default.(decimal.Decimal)
	// paymentDescPayWithBalanceAmount is the schema descriptor for pay_with_balance_amount field.
	paymentDescPayWithBalanceAmount := paymentFields[8].Descriptor()
	// payment.DefaultPayWithBalanceAmount holds the default value on creation for the pay_with_balance_amount field.
	payment.DefaultPayWithBalanceAmount = paymentDescPayWithBalanceAmount.Default.(decimal.Decimal)
	// paymentDescFinishAmount is the schema descriptor for finish_amount field.
	paymentDescFinishAmount := paymentFields[9].Descriptor()
	// payment.DefaultFinishAmount holds the default value on creation for the finish_amount field.
	payment.DefaultFinishAmount = paymentDescFinishAmount.Default.(decimal.Decimal)
	// paymentDescCoinUsdCurrency is the schema descriptor for coin_usd_currency field.
	paymentDescCoinUsdCurrency := paymentFields[10].Descriptor()
	// payment.DefaultCoinUsdCurrency holds the default value on creation for the coin_usd_currency field.
	payment.DefaultCoinUsdCurrency = paymentDescCoinUsdCurrency.Default.(decimal.Decimal)
	// paymentDescLocalCoinUsdCurrency is the schema descriptor for local_coin_usd_currency field.
	paymentDescLocalCoinUsdCurrency := paymentFields[11].Descriptor()
	// payment.DefaultLocalCoinUsdCurrency holds the default value on creation for the local_coin_usd_currency field.
	payment.DefaultLocalCoinUsdCurrency = paymentDescLocalCoinUsdCurrency.Default.(decimal.Decimal)
	// paymentDescLiveCoinUsdCurrency is the schema descriptor for live_coin_usd_currency field.
	paymentDescLiveCoinUsdCurrency := paymentFields[12].Descriptor()
	// payment.DefaultLiveCoinUsdCurrency holds the default value on creation for the live_coin_usd_currency field.
	payment.DefaultLiveCoinUsdCurrency = paymentDescLiveCoinUsdCurrency.Default.(decimal.Decimal)
	// paymentDescState is the schema descriptor for state field.
	paymentDescState := paymentFields[14].Descriptor()
	// payment.DefaultState holds the default value on creation for the state field.
	payment.DefaultState = paymentDescState.Default.(string)
	// paymentDescChainTransactionID is the schema descriptor for chain_transaction_id field.
	paymentDescChainTransactionID := paymentFields[15].Descriptor()
	// payment.DefaultChainTransactionID holds the default value on creation for the chain_transaction_id field.
	payment.DefaultChainTransactionID = paymentDescChainTransactionID.Default.(string)
	// paymentDescUserSetPaid is the schema descriptor for user_set_paid field.
	paymentDescUserSetPaid := paymentFields[16].Descriptor()
	// payment.DefaultUserSetPaid holds the default value on creation for the user_set_paid field.
	payment.DefaultUserSetPaid = paymentDescUserSetPaid.Default.(bool)
	// paymentDescUserSetCanceled is the schema descriptor for user_set_canceled field.
	paymentDescUserSetCanceled := paymentFields[17].Descriptor()
	// payment.DefaultUserSetCanceled holds the default value on creation for the user_set_canceled field.
	payment.DefaultUserSetCanceled = paymentDescUserSetCanceled.Default.(bool)
	// paymentDescFakePayment is the schema descriptor for fake_payment field.
	paymentDescFakePayment := paymentFields[18].Descriptor()
	// payment.DefaultFakePayment holds the default value on creation for the fake_payment field.
	payment.DefaultFakePayment = paymentDescFakePayment.Default.(bool)
	// paymentDescID is the schema descriptor for id field.
	paymentDescID := paymentFields[0].Descriptor()
	// payment.DefaultID holds the default value on creation for the id field.
	payment.DefaultID = paymentDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.3" // Version of ent codegen.
)
