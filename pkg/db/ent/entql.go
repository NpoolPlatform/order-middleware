// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 5)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   compensate.Table,
			Columns: compensate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: compensate.FieldID,
			},
		},
		Type: "Compensate",
		Fields: map[string]*sqlgraph.FieldSpec{
			compensate.FieldCreatedAt:      {Type: field.TypeUint32, Column: compensate.FieldCreatedAt},
			compensate.FieldUpdatedAt:      {Type: field.TypeUint32, Column: compensate.FieldUpdatedAt},
			compensate.FieldDeletedAt:      {Type: field.TypeUint32, Column: compensate.FieldDeletedAt},
			compensate.FieldOrderID:        {Type: field.TypeUUID, Column: compensate.FieldOrderID},
			compensate.FieldStartAt:        {Type: field.TypeUint32, Column: compensate.FieldStartAt},
			compensate.FieldEndAt:          {Type: field.TypeUint32, Column: compensate.FieldEndAt},
			compensate.FieldCompensateType: {Type: field.TypeString, Column: compensate.FieldCompensateType},
			compensate.FieldTitle:          {Type: field.TypeString, Column: compensate.FieldTitle},
			compensate.FieldMessage:        {Type: field.TypeString, Column: compensate.FieldMessage},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
		Type: "Order",
		Fields: map[string]*sqlgraph.FieldSpec{
			order.FieldCreatedAt:      {Type: field.TypeUint32, Column: order.FieldCreatedAt},
			order.FieldUpdatedAt:      {Type: field.TypeUint32, Column: order.FieldUpdatedAt},
			order.FieldDeletedAt:      {Type: field.TypeUint32, Column: order.FieldDeletedAt},
			order.FieldAppID:          {Type: field.TypeUUID, Column: order.FieldAppID},
			order.FieldUserID:         {Type: field.TypeUUID, Column: order.FieldUserID},
			order.FieldGoodID:         {Type: field.TypeUUID, Column: order.FieldGoodID},
			order.FieldPaymentID:      {Type: field.TypeUUID, Column: order.FieldPaymentID},
			order.FieldParentOrderID:  {Type: field.TypeUUID, Column: order.FieldParentOrderID},
			order.FieldUnitsV1:        {Type: field.TypeOther, Column: order.FieldUnitsV1},
			order.FieldGoodValue:      {Type: field.TypeOther, Column: order.FieldGoodValue},
			order.FieldPaymentAmount:  {Type: field.TypeOther, Column: order.FieldPaymentAmount},
			order.FieldDiscountAmount: {Type: field.TypeOther, Column: order.FieldDiscountAmount},
			order.FieldPromotionID:    {Type: field.TypeUUID, Column: order.FieldPromotionID},
			order.FieldDurationDays:   {Type: field.TypeUint32, Column: order.FieldDurationDays},
			order.FieldOrderType:      {Type: field.TypeString, Column: order.FieldOrderType},
			order.FieldInvestmentType: {Type: field.TypeString, Column: order.FieldInvestmentType},
			order.FieldCouponIds:      {Type: field.TypeJSON, Column: order.FieldCouponIds},
			order.FieldPaymentType:    {Type: field.TypeString, Column: order.FieldPaymentType},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   orderstate.Table,
			Columns: orderstate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderstate.FieldID,
			},
		},
		Type: "OrderState",
		Fields: map[string]*sqlgraph.FieldSpec{
			orderstate.FieldCreatedAt:            {Type: field.TypeUint32, Column: orderstate.FieldCreatedAt},
			orderstate.FieldUpdatedAt:            {Type: field.TypeUint32, Column: orderstate.FieldUpdatedAt},
			orderstate.FieldDeletedAt:            {Type: field.TypeUint32, Column: orderstate.FieldDeletedAt},
			orderstate.FieldOrderID:              {Type: field.TypeUUID, Column: orderstate.FieldOrderID},
			orderstate.FieldOrderState:           {Type: field.TypeString, Column: orderstate.FieldOrderState},
			orderstate.FieldStartMode:            {Type: field.TypeString, Column: orderstate.FieldStartMode},
			orderstate.FieldStartAt:              {Type: field.TypeUint32, Column: orderstate.FieldStartAt},
			orderstate.FieldEndAt:                {Type: field.TypeUint32, Column: orderstate.FieldEndAt},
			orderstate.FieldLastBenefitAt:        {Type: field.TypeUint32, Column: orderstate.FieldLastBenefitAt},
			orderstate.FieldBenefitState:         {Type: field.TypeString, Column: orderstate.FieldBenefitState},
			orderstate.FieldUserSetPaid:          {Type: field.TypeBool, Column: orderstate.FieldUserSetPaid},
			orderstate.FieldUserSetCancelled:     {Type: field.TypeBool, Column: orderstate.FieldUserSetCancelled},
			orderstate.FieldPaymentTransactionID: {Type: field.TypeString, Column: orderstate.FieldPaymentTransactionID},
			orderstate.FieldPaymentFinishAmount:  {Type: field.TypeOther, Column: orderstate.FieldPaymentFinishAmount},
			orderstate.FieldPaymentState:         {Type: field.TypeString, Column: orderstate.FieldPaymentState},
			orderstate.FieldOutofgasHours:        {Type: field.TypeUint32, Column: orderstate.FieldOutofgasHours},
			orderstate.FieldCompensateHours:      {Type: field.TypeUint32, Column: orderstate.FieldCompensateHours},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   outofgas.Table,
			Columns: outofgas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		},
		Type: "OutOfGas",
		Fields: map[string]*sqlgraph.FieldSpec{
			outofgas.FieldCreatedAt: {Type: field.TypeUint32, Column: outofgas.FieldCreatedAt},
			outofgas.FieldUpdatedAt: {Type: field.TypeUint32, Column: outofgas.FieldUpdatedAt},
			outofgas.FieldDeletedAt: {Type: field.TypeUint32, Column: outofgas.FieldDeletedAt},
			outofgas.FieldOrderID:   {Type: field.TypeUUID, Column: outofgas.FieldOrderID},
			outofgas.FieldStartAt:   {Type: field.TypeUint32, Column: outofgas.FieldStartAt},
			outofgas.FieldEndAt:     {Type: field.TypeUint32, Column: outofgas.FieldEndAt},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: payment.FieldID,
			},
		},
		Type: "Payment",
		Fields: map[string]*sqlgraph.FieldSpec{
			payment.FieldCreatedAt:            {Type: field.TypeUint32, Column: payment.FieldCreatedAt},
			payment.FieldUpdatedAt:            {Type: field.TypeUint32, Column: payment.FieldUpdatedAt},
			payment.FieldDeletedAt:            {Type: field.TypeUint32, Column: payment.FieldDeletedAt},
			payment.FieldAppID:                {Type: field.TypeUUID, Column: payment.FieldAppID},
			payment.FieldUserID:               {Type: field.TypeUUID, Column: payment.FieldUserID},
			payment.FieldGoodID:               {Type: field.TypeUUID, Column: payment.FieldGoodID},
			payment.FieldOrderID:              {Type: field.TypeUUID, Column: payment.FieldOrderID},
			payment.FieldAccountID:            {Type: field.TypeUUID, Column: payment.FieldAccountID},
			payment.FieldCoinInfoID:           {Type: field.TypeUUID, Column: payment.FieldCoinInfoID},
			payment.FieldStartAmount:          {Type: field.TypeOther, Column: payment.FieldStartAmount},
			payment.FieldTransferAmount:       {Type: field.TypeOther, Column: payment.FieldTransferAmount},
			payment.FieldBalanceAmount:        {Type: field.TypeOther, Column: payment.FieldBalanceAmount},
			payment.FieldCoinUsdCurrency:      {Type: field.TypeOther, Column: payment.FieldCoinUsdCurrency},
			payment.FieldLocalCoinUsdCurrency: {Type: field.TypeOther, Column: payment.FieldLocalCoinUsdCurrency},
			payment.FieldLiveCoinUsdCurrency:  {Type: field.TypeOther, Column: payment.FieldLiveCoinUsdCurrency},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *CompensateQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CompensateQuery builder.
func (cq *CompensateQuery) Filter() *CompensateFilter {
	return &CompensateFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CompensateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CompensateMutation builder.
func (m *CompensateMutation) Filter() *CompensateFilter {
	return &CompensateFilter{config: m.config, predicateAdder: m}
}

// CompensateFilter provides a generic filtering capability at runtime for CompensateQuery.
type CompensateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CompensateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CompensateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(compensate.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CompensateFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CompensateFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CompensateFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldDeletedAt))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *CompensateFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(compensate.FieldOrderID))
}

// WhereStartAt applies the entql uint32 predicate on the start_at field.
func (f *CompensateFilter) WhereStartAt(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldStartAt))
}

// WhereEndAt applies the entql uint32 predicate on the end_at field.
func (f *CompensateFilter) WhereEndAt(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldEndAt))
}

// WhereCompensateType applies the entql string predicate on the compensate_type field.
func (f *CompensateFilter) WhereCompensateType(p entql.StringP) {
	f.Where(p.Field(compensate.FieldCompensateType))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *CompensateFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(compensate.FieldTitle))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *CompensateFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(compensate.FieldMessage))
}

// addPredicate implements the predicateAdder interface.
func (oq *OrderQuery) addPredicate(pred func(s *sql.Selector)) {
	oq.predicates = append(oq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OrderQuery builder.
func (oq *OrderQuery) Filter() *OrderFilter {
	return &OrderFilter{config: oq.config, predicateAdder: oq}
}

// addPredicate implements the predicateAdder interface.
func (m *OrderMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OrderMutation builder.
func (m *OrderMutation) Filter() *OrderFilter {
	return &OrderFilter{config: m.config, predicateAdder: m}
}

// OrderFilter provides a generic filtering capability at runtime for OrderQuery.
type OrderFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OrderFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *OrderFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(order.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *OrderFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(order.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *OrderFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(order.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *OrderFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(order.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *OrderFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(order.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *OrderFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(order.FieldUserID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *OrderFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(order.FieldGoodID))
}

// WherePaymentID applies the entql [16]byte predicate on the payment_id field.
func (f *OrderFilter) WherePaymentID(p entql.ValueP) {
	f.Where(p.Field(order.FieldPaymentID))
}

// WhereParentOrderID applies the entql [16]byte predicate on the parent_order_id field.
func (f *OrderFilter) WhereParentOrderID(p entql.ValueP) {
	f.Where(p.Field(order.FieldParentOrderID))
}

// WhereUnitsV1 applies the entql other predicate on the units_v1 field.
func (f *OrderFilter) WhereUnitsV1(p entql.OtherP) {
	f.Where(p.Field(order.FieldUnitsV1))
}

// WhereGoodValue applies the entql other predicate on the good_value field.
func (f *OrderFilter) WhereGoodValue(p entql.OtherP) {
	f.Where(p.Field(order.FieldGoodValue))
}

// WherePaymentAmount applies the entql other predicate on the payment_amount field.
func (f *OrderFilter) WherePaymentAmount(p entql.OtherP) {
	f.Where(p.Field(order.FieldPaymentAmount))
}

// WhereDiscountAmount applies the entql other predicate on the discount_amount field.
func (f *OrderFilter) WhereDiscountAmount(p entql.OtherP) {
	f.Where(p.Field(order.FieldDiscountAmount))
}

// WherePromotionID applies the entql [16]byte predicate on the promotion_id field.
func (f *OrderFilter) WherePromotionID(p entql.ValueP) {
	f.Where(p.Field(order.FieldPromotionID))
}

// WhereDurationDays applies the entql uint32 predicate on the duration_days field.
func (f *OrderFilter) WhereDurationDays(p entql.Uint32P) {
	f.Where(p.Field(order.FieldDurationDays))
}

// WhereOrderType applies the entql string predicate on the order_type field.
func (f *OrderFilter) WhereOrderType(p entql.StringP) {
	f.Where(p.Field(order.FieldOrderType))
}

// WhereInvestmentType applies the entql string predicate on the investment_type field.
func (f *OrderFilter) WhereInvestmentType(p entql.StringP) {
	f.Where(p.Field(order.FieldInvestmentType))
}

// WhereCouponIds applies the entql json.RawMessage predicate on the coupon_ids field.
func (f *OrderFilter) WhereCouponIds(p entql.BytesP) {
	f.Where(p.Field(order.FieldCouponIds))
}

// WherePaymentType applies the entql string predicate on the payment_type field.
func (f *OrderFilter) WherePaymentType(p entql.StringP) {
	f.Where(p.Field(order.FieldPaymentType))
}

// addPredicate implements the predicateAdder interface.
func (osq *OrderStateQuery) addPredicate(pred func(s *sql.Selector)) {
	osq.predicates = append(osq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OrderStateQuery builder.
func (osq *OrderStateQuery) Filter() *OrderStateFilter {
	return &OrderStateFilter{config: osq.config, predicateAdder: osq}
}

// addPredicate implements the predicateAdder interface.
func (m *OrderStateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OrderStateMutation builder.
func (m *OrderStateMutation) Filter() *OrderStateFilter {
	return &OrderStateFilter{config: m.config, predicateAdder: m}
}

// OrderStateFilter provides a generic filtering capability at runtime for OrderStateQuery.
type OrderStateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OrderStateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *OrderStateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(orderstate.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *OrderStateFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(orderstate.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *OrderStateFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(orderstate.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *OrderStateFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(orderstate.FieldDeletedAt))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *OrderStateFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(orderstate.FieldOrderID))
}

// WhereOrderState applies the entql string predicate on the order_state field.
func (f *OrderStateFilter) WhereOrderState(p entql.StringP) {
	f.Where(p.Field(orderstate.FieldOrderState))
}

// WhereStartMode applies the entql string predicate on the start_mode field.
func (f *OrderStateFilter) WhereStartMode(p entql.StringP) {
	f.Where(p.Field(orderstate.FieldStartMode))
}

// WhereStartAt applies the entql uint32 predicate on the start_at field.
func (f *OrderStateFilter) WhereStartAt(p entql.Uint32P) {
	f.Where(p.Field(orderstate.FieldStartAt))
}

// WhereEndAt applies the entql uint32 predicate on the end_at field.
func (f *OrderStateFilter) WhereEndAt(p entql.Uint32P) {
	f.Where(p.Field(orderstate.FieldEndAt))
}

// WhereLastBenefitAt applies the entql uint32 predicate on the last_benefit_at field.
func (f *OrderStateFilter) WhereLastBenefitAt(p entql.Uint32P) {
	f.Where(p.Field(orderstate.FieldLastBenefitAt))
}

// WhereBenefitState applies the entql string predicate on the benefit_state field.
func (f *OrderStateFilter) WhereBenefitState(p entql.StringP) {
	f.Where(p.Field(orderstate.FieldBenefitState))
}

// WhereUserSetPaid applies the entql bool predicate on the user_set_paid field.
func (f *OrderStateFilter) WhereUserSetPaid(p entql.BoolP) {
	f.Where(p.Field(orderstate.FieldUserSetPaid))
}

// WhereUserSetCancelled applies the entql bool predicate on the user_set_cancelled field.
func (f *OrderStateFilter) WhereUserSetCancelled(p entql.BoolP) {
	f.Where(p.Field(orderstate.FieldUserSetCancelled))
}

// WherePaymentTransactionID applies the entql string predicate on the payment_transaction_id field.
func (f *OrderStateFilter) WherePaymentTransactionID(p entql.StringP) {
	f.Where(p.Field(orderstate.FieldPaymentTransactionID))
}

// WherePaymentFinishAmount applies the entql other predicate on the payment_finish_amount field.
func (f *OrderStateFilter) WherePaymentFinishAmount(p entql.OtherP) {
	f.Where(p.Field(orderstate.FieldPaymentFinishAmount))
}

// WherePaymentState applies the entql string predicate on the payment_state field.
func (f *OrderStateFilter) WherePaymentState(p entql.StringP) {
	f.Where(p.Field(orderstate.FieldPaymentState))
}

// WhereOutofgasHours applies the entql uint32 predicate on the outofgas_hours field.
func (f *OrderStateFilter) WhereOutofgasHours(p entql.Uint32P) {
	f.Where(p.Field(orderstate.FieldOutofgasHours))
}

// WhereCompensateHours applies the entql uint32 predicate on the compensate_hours field.
func (f *OrderStateFilter) WhereCompensateHours(p entql.Uint32P) {
	f.Where(p.Field(orderstate.FieldCompensateHours))
}

// addPredicate implements the predicateAdder interface.
func (oogq *OutOfGasQuery) addPredicate(pred func(s *sql.Selector)) {
	oogq.predicates = append(oogq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OutOfGasQuery builder.
func (oogq *OutOfGasQuery) Filter() *OutOfGasFilter {
	return &OutOfGasFilter{config: oogq.config, predicateAdder: oogq}
}

// addPredicate implements the predicateAdder interface.
func (m *OutOfGasMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OutOfGasMutation builder.
func (m *OutOfGasMutation) Filter() *OutOfGasFilter {
	return &OutOfGasFilter{config: m.config, predicateAdder: m}
}

// OutOfGasFilter provides a generic filtering capability at runtime for OutOfGasQuery.
type OutOfGasFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OutOfGasFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *OutOfGasFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(outofgas.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *OutOfGasFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *OutOfGasFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *OutOfGasFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldDeletedAt))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *OutOfGasFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(outofgas.FieldOrderID))
}

// WhereStartAt applies the entql uint32 predicate on the start_at field.
func (f *OutOfGasFilter) WhereStartAt(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldStartAt))
}

// WhereEndAt applies the entql uint32 predicate on the end_at field.
func (f *OutOfGasFilter) WhereEndAt(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldEndAt))
}

// addPredicate implements the predicateAdder interface.
func (pq *PaymentQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PaymentQuery builder.
func (pq *PaymentQuery) Filter() *PaymentFilter {
	return &PaymentFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PaymentMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PaymentMutation builder.
func (m *PaymentMutation) Filter() *PaymentFilter {
	return &PaymentFilter{config: m.config, predicateAdder: m}
}

// PaymentFilter provides a generic filtering capability at runtime for PaymentQuery.
type PaymentFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PaymentFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PaymentFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PaymentFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PaymentFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PaymentFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *PaymentFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *PaymentFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldUserID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *PaymentFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldGoodID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *PaymentFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldOrderID))
}

// WhereAccountID applies the entql [16]byte predicate on the account_id field.
func (f *PaymentFilter) WhereAccountID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldAccountID))
}

// WhereCoinInfoID applies the entql [16]byte predicate on the coin_info_id field.
func (f *PaymentFilter) WhereCoinInfoID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldCoinInfoID))
}

// WhereStartAmount applies the entql other predicate on the start_amount field.
func (f *PaymentFilter) WhereStartAmount(p entql.OtherP) {
	f.Where(p.Field(payment.FieldStartAmount))
}

// WhereTransferAmount applies the entql other predicate on the transfer_amount field.
func (f *PaymentFilter) WhereTransferAmount(p entql.OtherP) {
	f.Where(p.Field(payment.FieldTransferAmount))
}

// WhereBalanceAmount applies the entql other predicate on the balance_amount field.
func (f *PaymentFilter) WhereBalanceAmount(p entql.OtherP) {
	f.Where(p.Field(payment.FieldBalanceAmount))
}

// WhereCoinUsdCurrency applies the entql other predicate on the coin_usd_currency field.
func (f *PaymentFilter) WhereCoinUsdCurrency(p entql.OtherP) {
	f.Where(p.Field(payment.FieldCoinUsdCurrency))
}

// WhereLocalCoinUsdCurrency applies the entql other predicate on the local_coin_usd_currency field.
func (f *PaymentFilter) WhereLocalCoinUsdCurrency(p entql.OtherP) {
	f.Where(p.Field(payment.FieldLocalCoinUsdCurrency))
}

// WhereLiveCoinUsdCurrency applies the entql other predicate on the live_coin_usd_currency field.
func (f *PaymentFilter) WhereLiveCoinUsdCurrency(p entql.OtherP) {
	f.Where(p.Field(payment.FieldLiveCoinUsdCurrency))
}
