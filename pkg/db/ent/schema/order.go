package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("user_id", uuid.UUID{}),
		field.
			UUID("parent_order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("pay_with_parent").
			Optional().
			Default(false),
		field.
			Uint32("units").
			Optional().
			Default(0),
		field.
			Other("units_v1", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			UUID("promotion_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("discount_coupon_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_special_reduction_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
		field.
			UUID("fix_amount_coupon_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("type").
			Optional().
			Default(basetypes.OrderType_DefaultOrderType.String()),
		field.
			String("state").
			Optional().
			Default(basetypes.OrderState_DefaultOrderState.String()),
		field.
			JSON("coupon_ids", []uuid.UUID{}).
			Optional().
			Default(func() []uuid.UUID {
				return []uuid.UUID{}
			}),
		field.
			Uint32("last_benefit_at").
			Optional().
			Default(0),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
