package orderstm

import (
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type order struct {
	entOrderBase      *ent.OrderBase
	entOrderStateBase *ent.OrderStateBase
}

func (o *order) OrderState() types.OrderState {
	return types.OrderState(types.OrderState_value[o.entOrderStateBase.OrderState])
}

func (o *order) OrderType() types.OrderType {
	return types.OrderType(types.OrderType_value[o.entOrderBase.OrderType])
}

func (o *order) PaymentType() types.PaymentType {
	return types.PaymentType(types.PaymentType_value[o.entOrderStateBase.PaymentType])
}
