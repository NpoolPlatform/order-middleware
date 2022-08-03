package order

import (
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
)

func trace(span trace1.Span, in *npool.OrderReq, index int) trace1.Span {
	// TODO: add order tracer
	return span
}

func Trace(span trace1.Span, in *npool.OrderReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	// TODO: add order conds tracer
	return span
}

func TraceMany(span trace1.Span, infos []*npool.OrderReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
