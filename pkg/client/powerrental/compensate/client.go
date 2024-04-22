package compensate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	servicename "github.com/NpoolPlatform/order-middleware/pkg/servicename"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second, //nolint
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func CreateCompensate(ctx context.Context, in *npool.CompensateReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateCompensate(ctx, &npool.CreateCompensateRequest{
			Info: in,
		})
	})
	return err
}

func UpdateCompensate(ctx context.Context, in *npool.CompensateReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateCompensate(ctx, &npool.UpdateCompensateRequest{
			Info: in,
		})
	})
	return err
}

func GetCompensate(ctx context.Context, id string) (*npool.Compensate, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCompensate(ctx, &npool.GetCompensateRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Compensate), nil
}

func GetCompensates(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Compensate, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCompensates(ctx, &npool.GetCompensatesRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = resp.Total
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return _infos.([]*npool.Compensate), total, nil
}

func GetCompensateOnly(ctx context.Context, conds *npool.Conds) (*npool.Compensate, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCompensates(ctx, &npool.GetCompensatesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Compensate)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Compensate)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Compensate)[0], nil
}

func DeleteCompensate(ctx context.Context, id *uint32, entID *string) error {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.DeleteCompensate(ctx, &npool.DeleteCompensateRequest{
			Info: &npool.CompensateReq{
				ID:    id,
				EntID: entID,
			},
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Compensate), nil
}

func ExistCompensate(ctx context.Context, entID string) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistCompensate(ctx, &npool.ExistCompensateRequest{
			EntID: entID,
		})
		if err != nil {
			return false, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return exist.(bool), err
}

func ExistCompensateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistCompensateConds(ctx, &npool.ExistCompensateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return false, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return exist.(bool), err
}
