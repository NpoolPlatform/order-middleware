package outofgas

import (
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type outOfGas struct {
	entOutOfGas *ent.OutOfGas
}

func (f *outOfGas) OutOfGasID() uint32 {
	return f.entOutOfGas.ID
}

func (f *outOfGas) OrderID() uuid.UUID {
	return f.entOutOfGas.OrderID
}

func (f *outOfGas) StartAt() uint32 {
	return f.entOutOfGas.StartAt
}
