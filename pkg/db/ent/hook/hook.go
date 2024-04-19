// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

// The AppConfigFunc type is an adapter to allow the use of ordinary
// function as AppConfig mutator.
type AppConfigFunc func(context.Context, *ent.AppConfigMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppConfigFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppConfigMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppConfigMutation", m)
	}
	return f(ctx, mv)
}

// The CompensateFunc type is an adapter to allow the use of ordinary
// function as Compensate mutator.
type CompensateFunc func(context.Context, *ent.CompensateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompensateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CompensateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompensateMutation", m)
	}
	return f(ctx, mv)
}

// The FeeOrderFunc type is an adapter to allow the use of ordinary
// function as FeeOrder mutator.
type FeeOrderFunc func(context.Context, *ent.FeeOrderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeeOrderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FeeOrderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeeOrderMutation", m)
	}
	return f(ctx, mv)
}

// The FeeOrderStateFunc type is an adapter to allow the use of ordinary
// function as FeeOrderState mutator.
type FeeOrderStateFunc func(context.Context, *ent.FeeOrderStateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeeOrderStateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FeeOrderStateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeeOrderStateMutation", m)
	}
	return f(ctx, mv)
}

// The OrderFunc type is an adapter to allow the use of ordinary
// function as Order mutator.
type OrderFunc func(context.Context, *ent.OrderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderMutation", m)
	}
	return f(ctx, mv)
}

// The OrderBaseFunc type is an adapter to allow the use of ordinary
// function as OrderBase mutator.
type OrderBaseFunc func(context.Context, *ent.OrderBaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderBaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderBaseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderBaseMutation", m)
	}
	return f(ctx, mv)
}

// The OrderCouponFunc type is an adapter to allow the use of ordinary
// function as OrderCoupon mutator.
type OrderCouponFunc func(context.Context, *ent.OrderCouponMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderCouponFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderCouponMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderCouponMutation", m)
	}
	return f(ctx, mv)
}

// The OrderLockFunc type is an adapter to allow the use of ordinary
// function as OrderLock mutator.
type OrderLockFunc func(context.Context, *ent.OrderLockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderLockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderLockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderLockMutation", m)
	}
	return f(ctx, mv)
}

// The OrderStateFunc type is an adapter to allow the use of ordinary
// function as OrderState mutator.
type OrderStateFunc func(context.Context, *ent.OrderStateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderStateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderStateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderStateMutation", m)
	}
	return f(ctx, mv)
}

// The OrderStateBaseFunc type is an adapter to allow the use of ordinary
// function as OrderStateBase mutator.
type OrderStateBaseFunc func(context.Context, *ent.OrderStateBaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderStateBaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderStateBaseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderStateBaseMutation", m)
	}
	return f(ctx, mv)
}

// The OutOfGasFunc type is an adapter to allow the use of ordinary
// function as OutOfGas mutator.
type OutOfGasFunc func(context.Context, *ent.OutOfGasMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OutOfGasFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OutOfGasMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OutOfGasMutation", m)
	}
	return f(ctx, mv)
}

// The PaymentFunc type is an adapter to allow the use of ordinary
// function as Payment mutator.
type PaymentFunc func(context.Context, *ent.PaymentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaymentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentMutation", m)
	}
	return f(ctx, mv)
}

// The PaymentBalanceFunc type is an adapter to allow the use of ordinary
// function as PaymentBalance mutator.
type PaymentBalanceFunc func(context.Context, *ent.PaymentBalanceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentBalanceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaymentBalanceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentBalanceMutation", m)
	}
	return f(ctx, mv)
}

// The PaymentContractFunc type is an adapter to allow the use of ordinary
// function as PaymentContract mutator.
type PaymentContractFunc func(context.Context, *ent.PaymentContractMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentContractFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaymentContractMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentContractMutation", m)
	}
	return f(ctx, mv)
}

// The PaymentTransferFunc type is an adapter to allow the use of ordinary
// function as PaymentTransfer mutator.
type PaymentTransferFunc func(context.Context, *ent.PaymentTransferMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentTransferFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaymentTransferMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentTransferMutation", m)
	}
	return f(ctx, mv)
}

// The PowerRentalFunc type is an adapter to allow the use of ordinary
// function as PowerRental mutator.
type PowerRentalFunc func(context.Context, *ent.PowerRentalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PowerRentalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PowerRentalMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PowerRentalMutation", m)
	}
	return f(ctx, mv)
}

// The PowerRentalStateFunc type is an adapter to allow the use of ordinary
// function as PowerRentalState mutator.
type PowerRentalStateFunc func(context.Context, *ent.PowerRentalStateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PowerRentalStateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PowerRentalStateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PowerRentalStateMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
