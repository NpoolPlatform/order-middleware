package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/order-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// PaymentBase holds the schema definition for the PaymentBase entity.
type PaymentBase struct {
	ent.Schema
}

func (PaymentBase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the PaymentBase.
func (PaymentBase) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		// When user change order payment method, we'll create a new payment and obselete the old one
		field.
			Bool("obseleted").
			Optional().
			Default(false),
	}
}

// Edges of the PaymentBase.
func (PaymentBase) Edges() []ent.Edge {
	return nil
}

func (PaymentBase) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
