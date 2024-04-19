package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// PowerRentalState holds the schema definition for the PowerRentalState entity.
type PowerRentalState struct {
	ent.Schema
}

func (PowerRentalState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the PowerRentalState.
func (PowerRentalState) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("cancel_state").
			Optional().
			Default(types.OrderState_DefaultOrderState.String()),
		field.
			Uint32("duration_seconds").
			Optional().
			Default(0),
		field.
			Uint32("paid_at").
			Optional().
			Default(0),
		field.
			Bool("user_set_paid").
			Optional().
			Default(false),
		field.
			Bool("user_set_canceled").
			Optional().
			Default(false),
		field.
			Bool("admin_set_canceled").
			Optional().
			Default(false),
		field.
			String("payment_state").
			Optional().
			Default(types.PaymentState_PaymentStateWait.String()),
		field.
			Uint32("outofgas_seconds").
			Optional().
			Default(0),
		field.
			Uint32("compensate_seconds").
			Optional().
			Default(0),
		field.
			String("renew_state").
			Optional().
			Default(types.OrderRenewState_OrderRenewWait.String()),
		field.
			Uint32("renew_notify_at").
			Optional().
			Default(0),
	}
}

// Edges of the PowerRentalState.
func (PowerRentalState) Edges() []ent.Edge {
	return nil
}

func (PowerRentalState) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
