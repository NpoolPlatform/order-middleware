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

// OrderLock holds the schema definition for the OrderLock entity.
type OrderLock struct {
	ent.Schema
}

func (OrderLock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OrderLock.
func (OrderLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("user_id", uuid.UUID{}),
		field.
			UUID("order_id", uuid.UUID{}),
		field.
			String("lock_type").
			Optional().
			Default(types.OrderLockType_DefaultOrderLockType.String()),
	}
}

// Edges of the OrderLock.
func (OrderLock) Edges() []ent.Edge {
	return nil
}

func (OrderLock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
