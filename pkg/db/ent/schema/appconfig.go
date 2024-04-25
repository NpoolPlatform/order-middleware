package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppConfig holds the schema definition for the AppConfig entity.
type AppConfig struct {
	ent.Schema
}

func (AppConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppConfig.
func (AppConfig) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Bool("enable_simulate_order").
			Optional().
			Default(false),
		field.
			Other("simulate_order_units", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.NewFromInt(0)),
		field.
			String("simulate_order_coupon_mode").
			Optional().
			Default(types.SimulateOrderCouponMode_WithoutCoupon.String()),
		field.
			Other("simulate_order_coupon_probability", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.NewFromInt(0)),
		field.
			Other("simulate_order_cashable_profit_probability", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.NewFromInt(0)),
		field.
			Uint32("max_unpaid_orders").
			Optional().
			Default(5), //nolint
	}
}

// Edges of the AppConfig.
func (AppConfig) Edges() []ent.Edge {
	return nil
}
