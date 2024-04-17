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

// OrderStateBase holds the schema definition for the OrderStateBase entity.
type OrderStateBase struct {
	ent.Schema
}

func (OrderStateBase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OrderStateBase.
func (OrderStateBase) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("order_state").
			Optional().
			Default(types.OrderState_OrderStateCreated.String()),
		field.
			String("start_mode").
			Optional().
			Default(types.OrderStartMode_OrderStartConfirmed.String()),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("last_benefit_at").
			Optional().
			Default(0),
		field.
			String("benefit_state").
			Optional().
			Default(types.BenefitState_BenefitWait.String()),
	}
}

// Edges of the OrderStateBase.
func (OrderStateBase) Edges() []ent.Edge {
	return nil
}

func (OrderStateBase) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
