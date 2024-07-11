package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	appconfigcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/app/config"
	entappconfig "github.com/NpoolPlatform/order-middleware/pkg/db/ent/appconfig"
)

func (h *Handler) ExistAppConfig(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil && h.AppID == nil {
		return false, wlog.Errorf("invalid id")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppConfig.Query().Where(entappconfig.DeletedAt(0))
		if h.EntID != nil {
			stm.Where(entappconfig.EntID(*h.EntID))
		}
		if h.AppID != nil {
			stm.Where(entappconfig.AppID(*h.AppID))
		}
		exist, err = stm.Exist(_ctx)
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistAppConfigConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appconfigcrud.SetQueryConds(cli.AppConfig.Query(), h.AppConfigConds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
