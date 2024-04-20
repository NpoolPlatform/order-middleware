package compensate

import (
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type powerRentalStates struct {
	entPowerRentalStates []*ent.PowerRentalState
}

func (f *powerRentalStates) PowerRentalStateID() uint32 {
	return f.entPowerRentalStates[0].ID
}

func (f *powerRentalStates) OrderID() uuid.UUID {
	return f.entPowerRentalStates[0].OrderID
}

func (f *powerRentalStates) CompensateSeconds() uint32 {
	return f.entPowerRentalStates[0].CompensateSeconds
}

func (f *powerRentalStates) OrderIDWithIndex(index int) uuid.UUID {
	return f.entPowerRentalStates[index].OrderID
}

func (f *powerRentalStates) CompensateSecondsWithIndex(index int) uint32 {
	return f.entPowerRentalStates[index].CompensateSeconds
}

func (f *powerRentalStates) Exhausted() bool {
	return len(f.entPowerRentalStates) == 0
}
