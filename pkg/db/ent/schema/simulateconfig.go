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

// SimulateConfig holds the schema definition for the SimulateConfig entity.
type SimulateConfig struct {
	ent.Schema
}

func (SimulateConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the SimulateConfig.
func (SimulateConfig) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			String("send_coupon_mode").
			Optional().
			Default(types.SendCouponMode_WithoutCoupon.String()),
		field.
			Other("send_coupon_probability", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("enabled_cashable_profit").
			Optional().
			Default(false),
		field.
			Other("cashable_profit_probability", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("enabled").
			Optional().
			Default(false),
	}
}

// Edges of the SimulateConfig.
func (SimulateConfig) Edges() []ent.Edge {
	return nil
}
