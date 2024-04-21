package outofgas

import (
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type powerRentalState struct {
	entPowerRentalState *ent.PowerRentalState
}

func (f *powerRentalState) PowerRentalStateID() uint32 {
	return f.entPowerRentalState.ID
}

func (f *powerRentalState) OrderID() uuid.UUID {
	return f.entPowerRentalState.OrderID
}

func (f *powerRentalState) OutOfGasSeconds() uint32 {
	return f.entPowerRentalState.OutofgasSeconds
}
