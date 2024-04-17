package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PowerRental holds the schema definition for the PowerRental entity.
type PowerRental struct {
	ent.Schema
}

func (PowerRental) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the PowerRental.
//nolint:funlen
func (PowerRental) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("units", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("good_value_usd", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("payment_amount_usd", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("discount_amount_usd", decimal.Decimal{}).
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
			Uint32("duration").
			Optional().
			Default(0),
		field.
			String("investment_type").
			Optional().
			Default(types.InvestmentType_FullPayment.String()),
	}
}

// Edges of the PowerRental.
func (PowerRental) Edges() []ent.Edge {
	return nil
}

func (PowerRental) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
