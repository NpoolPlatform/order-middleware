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

// OrderBase holds the schema definition for the OrderBase entity.
type OrderBase struct {
	ent.Schema
}

func (OrderBase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OrderBase.
//nolint:funlen
func (OrderBase) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("good_type").
			Optional().
			Default(""),
		field.
			UUID("parent_order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("order_type").
			Optional().
			Default(types.OrderType_Normal.String()),
		field.
			String("payment_type").
			Optional().
			Default(types.PaymentType_PayWithBalanceOnly.String()),
		field.
			String("create_method").
			Optional().
			Default(types.OrderCreateMethod_OrderCreatedByPurchase.String()),
		field.
			Bool("simulate").
			Optional().
			Default(false),
	}
}

// Edges of the OrderBase.
func (OrderBase) Edges() []ent.Edge {
	return nil
}

func (OrderBase) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "app_good_id"),
	}
}
