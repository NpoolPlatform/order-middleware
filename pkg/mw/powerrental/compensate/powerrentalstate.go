package compensate

import (
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type PowerRentalState interface {
	OrderID() uuid.UUID
}

type powerRentalState struct {
	entPowerRentalState *ent.PowerRentalState
}

func (f *powerRentalState) OrderID() uuid.UUID {
	return f.entPowerRentalState.OrderID
}

func (f *powerRentalState) CompensateSeconds() uint32 {
	return f.entPowerRentalState.CompensateSeconds
}
