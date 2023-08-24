// Code generated by ent, DO NOT EDIT.

package payment

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// CoinInfoID applies equality check predicate on the "coin_info_id" field. It's identical to CoinInfoIDEQ.
func CoinInfoID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinInfoID), v))
	})
}

// StartAmount applies equality check predicate on the "start_amount" field. It's identical to StartAmountEQ.
func StartAmount(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAmount), v))
	})
}

// TransferAmount applies equality check predicate on the "transfer_amount" field. It's identical to TransferAmountEQ.
func TransferAmount(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransferAmount), v))
	})
}

// BalanceAmount applies equality check predicate on the "balance_amount" field. It's identical to BalanceAmountEQ.
func BalanceAmount(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBalanceAmount), v))
	})
}

// CoinUsdCurrency applies equality check predicate on the "coin_usd_currency" field. It's identical to CoinUsdCurrencyEQ.
func CoinUsdCurrency(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// LocalCoinUsdCurrency applies equality check predicate on the "local_coin_usd_currency" field. It's identical to LocalCoinUsdCurrencyEQ.
func LocalCoinUsdCurrency(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocalCoinUsdCurrency), v))
	})
}

// LiveCoinUsdCurrency applies equality check predicate on the "live_coin_usd_currency" field. It's identical to LiveCoinUsdCurrencyEQ.
func LiveCoinUsdCurrency(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLiveCoinUsdCurrency), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountID), v))
	})
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccountID), v...))
	})
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccountID), v...))
	})
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountID), v))
	})
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountID), v))
	})
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountID), v))
	})
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountID), v))
	})
}

// CoinInfoIDEQ applies the EQ predicate on the "coin_info_id" field.
func CoinInfoIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDNEQ applies the NEQ predicate on the "coin_info_id" field.
func CoinInfoIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDIn applies the In predicate on the "coin_info_id" field.
func CoinInfoIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinInfoID), v...))
	})
}

// CoinInfoIDNotIn applies the NotIn predicate on the "coin_info_id" field.
func CoinInfoIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinInfoID), v...))
	})
}

// CoinInfoIDGT applies the GT predicate on the "coin_info_id" field.
func CoinInfoIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDGTE applies the GTE predicate on the "coin_info_id" field.
func CoinInfoIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDLT applies the LT predicate on the "coin_info_id" field.
func CoinInfoIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDLTE applies the LTE predicate on the "coin_info_id" field.
func CoinInfoIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinInfoID), v))
	})
}

// StartAmountEQ applies the EQ predicate on the "start_amount" field.
func StartAmountEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAmount), v))
	})
}

// StartAmountNEQ applies the NEQ predicate on the "start_amount" field.
func StartAmountNEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAmount), v))
	})
}

// StartAmountIn applies the In predicate on the "start_amount" field.
func StartAmountIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartAmount), v...))
	})
}

// StartAmountNotIn applies the NotIn predicate on the "start_amount" field.
func StartAmountNotIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartAmount), v...))
	})
}

// StartAmountGT applies the GT predicate on the "start_amount" field.
func StartAmountGT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAmount), v))
	})
}

// StartAmountGTE applies the GTE predicate on the "start_amount" field.
func StartAmountGTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAmount), v))
	})
}

// StartAmountLT applies the LT predicate on the "start_amount" field.
func StartAmountLT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAmount), v))
	})
}

// StartAmountLTE applies the LTE predicate on the "start_amount" field.
func StartAmountLTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAmount), v))
	})
}

// StartAmountIsNil applies the IsNil predicate on the "start_amount" field.
func StartAmountIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartAmount)))
	})
}

// StartAmountNotNil applies the NotNil predicate on the "start_amount" field.
func StartAmountNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartAmount)))
	})
}

// TransferAmountEQ applies the EQ predicate on the "transfer_amount" field.
func TransferAmountEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransferAmount), v))
	})
}

// TransferAmountNEQ applies the NEQ predicate on the "transfer_amount" field.
func TransferAmountNEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransferAmount), v))
	})
}

// TransferAmountIn applies the In predicate on the "transfer_amount" field.
func TransferAmountIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTransferAmount), v...))
	})
}

// TransferAmountNotIn applies the NotIn predicate on the "transfer_amount" field.
func TransferAmountNotIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTransferAmount), v...))
	})
}

// TransferAmountGT applies the GT predicate on the "transfer_amount" field.
func TransferAmountGT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransferAmount), v))
	})
}

// TransferAmountGTE applies the GTE predicate on the "transfer_amount" field.
func TransferAmountGTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransferAmount), v))
	})
}

// TransferAmountLT applies the LT predicate on the "transfer_amount" field.
func TransferAmountLT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransferAmount), v))
	})
}

// TransferAmountLTE applies the LTE predicate on the "transfer_amount" field.
func TransferAmountLTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransferAmount), v))
	})
}

// TransferAmountIsNil applies the IsNil predicate on the "transfer_amount" field.
func TransferAmountIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTransferAmount)))
	})
}

// TransferAmountNotNil applies the NotNil predicate on the "transfer_amount" field.
func TransferAmountNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTransferAmount)))
	})
}

// BalanceAmountEQ applies the EQ predicate on the "balance_amount" field.
func BalanceAmountEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBalanceAmount), v))
	})
}

// BalanceAmountNEQ applies the NEQ predicate on the "balance_amount" field.
func BalanceAmountNEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBalanceAmount), v))
	})
}

// BalanceAmountIn applies the In predicate on the "balance_amount" field.
func BalanceAmountIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBalanceAmount), v...))
	})
}

// BalanceAmountNotIn applies the NotIn predicate on the "balance_amount" field.
func BalanceAmountNotIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBalanceAmount), v...))
	})
}

// BalanceAmountGT applies the GT predicate on the "balance_amount" field.
func BalanceAmountGT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBalanceAmount), v))
	})
}

// BalanceAmountGTE applies the GTE predicate on the "balance_amount" field.
func BalanceAmountGTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBalanceAmount), v))
	})
}

// BalanceAmountLT applies the LT predicate on the "balance_amount" field.
func BalanceAmountLT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBalanceAmount), v))
	})
}

// BalanceAmountLTE applies the LTE predicate on the "balance_amount" field.
func BalanceAmountLTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBalanceAmount), v))
	})
}

// BalanceAmountIsNil applies the IsNil predicate on the "balance_amount" field.
func BalanceAmountIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBalanceAmount)))
	})
}

// BalanceAmountNotNil applies the NotNil predicate on the "balance_amount" field.
func BalanceAmountNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBalanceAmount)))
	})
}

// CoinUsdCurrencyEQ applies the EQ predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyNEQ applies the NEQ predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyNEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyIn applies the In predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinUsdCurrency), v...))
	})
}

// CoinUsdCurrencyNotIn applies the NotIn predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyNotIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinUsdCurrency), v...))
	})
}

// CoinUsdCurrencyGT applies the GT predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyGT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyGTE applies the GTE predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyGTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyLT applies the LT predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyLT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyLTE applies the LTE predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyLTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyIsNil applies the IsNil predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinUsdCurrency)))
	})
}

// CoinUsdCurrencyNotNil applies the NotNil predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinUsdCurrency)))
	})
}

// LocalCoinUsdCurrencyEQ applies the EQ predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocalCoinUsdCurrency), v))
	})
}

// LocalCoinUsdCurrencyNEQ applies the NEQ predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyNEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocalCoinUsdCurrency), v))
	})
}

// LocalCoinUsdCurrencyIn applies the In predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLocalCoinUsdCurrency), v...))
	})
}

// LocalCoinUsdCurrencyNotIn applies the NotIn predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyNotIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLocalCoinUsdCurrency), v...))
	})
}

// LocalCoinUsdCurrencyGT applies the GT predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyGT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocalCoinUsdCurrency), v))
	})
}

// LocalCoinUsdCurrencyGTE applies the GTE predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyGTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocalCoinUsdCurrency), v))
	})
}

// LocalCoinUsdCurrencyLT applies the LT predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyLT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocalCoinUsdCurrency), v))
	})
}

// LocalCoinUsdCurrencyLTE applies the LTE predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyLTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocalCoinUsdCurrency), v))
	})
}

// LocalCoinUsdCurrencyIsNil applies the IsNil predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLocalCoinUsdCurrency)))
	})
}

// LocalCoinUsdCurrencyNotNil applies the NotNil predicate on the "local_coin_usd_currency" field.
func LocalCoinUsdCurrencyNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLocalCoinUsdCurrency)))
	})
}

// LiveCoinUsdCurrencyEQ applies the EQ predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLiveCoinUsdCurrency), v))
	})
}

// LiveCoinUsdCurrencyNEQ applies the NEQ predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyNEQ(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLiveCoinUsdCurrency), v))
	})
}

// LiveCoinUsdCurrencyIn applies the In predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLiveCoinUsdCurrency), v...))
	})
}

// LiveCoinUsdCurrencyNotIn applies the NotIn predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyNotIn(vs ...decimal.Decimal) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLiveCoinUsdCurrency), v...))
	})
}

// LiveCoinUsdCurrencyGT applies the GT predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyGT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLiveCoinUsdCurrency), v))
	})
}

// LiveCoinUsdCurrencyGTE applies the GTE predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyGTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLiveCoinUsdCurrency), v))
	})
}

// LiveCoinUsdCurrencyLT applies the LT predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyLT(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLiveCoinUsdCurrency), v))
	})
}

// LiveCoinUsdCurrencyLTE applies the LTE predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyLTE(v decimal.Decimal) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLiveCoinUsdCurrency), v))
	})
}

// LiveCoinUsdCurrencyIsNil applies the IsNil predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLiveCoinUsdCurrency)))
	})
}

// LiveCoinUsdCurrencyNotNil applies the NotNil predicate on the "live_coin_usd_currency" field.
func LiveCoinUsdCurrencyNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLiveCoinUsdCurrency)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		p(s.Not())
	})
}
