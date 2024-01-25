package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	"github.com/NpoolPlatform/order-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

func (Payment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("user_id", uuid.UUID{}),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			UUID("order_id", uuid.UUID{}),
		field.
			UUID("account_id", uuid.UUID{}),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional(),
		field.
			UUID("coin_info_id", uuid.UUID{}).
			Optional(),
		field.
			Other("start_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("multi_payment_coins").
			Optional().
			Default(false),
		field.
			JSON("payment_amounts", []ordermwpb.PaymentAmount{}).
			Optional().
			Default([]ordermwpb.PaymentAmount{}),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return nil
}

func (Payment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
