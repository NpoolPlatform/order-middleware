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

// Compensate holds the schema definition for the Compensate entity.
type Compensate struct {
	ent.Schema
}

func (Compensate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Compensate.
func (Compensate) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		// Malfunction: ID is from good_manager.good_malfunction
		// Walfare: ID is from inspire_manager.npool_walfares
		// StarterDelay: ID is from good_manager.starter_delays
		field.
			UUID("compensate_from_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("compensate_type").
			Optional().
			Default(types.CompensateType_DefaultCompensateType.String()),
		field.
			Uint32("compensate_seconds").
			Optional().
			Default(0),
	}
}

// Edges of the Compensate.
func (Compensate) Edges() []ent.Edge {
	return nil
}

func (Compensate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
