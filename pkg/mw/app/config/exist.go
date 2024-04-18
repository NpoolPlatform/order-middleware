package appconfig

import (
	"context"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	appconfigcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/app/config"
	entappconfig "github.com/NpoolPlatform/order-middleware/pkg/db/ent/appconfig"
)

func (h *Handler) ExistAppConfig(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			AppConfig.
			Query().
			Where(
				entappconfig.EntID(*h.EntID),
				entappconfig.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistAppConfigConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appconfigcrud.SetQueryConds(cli.AppConfig.Query(), h.AppConfigConds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
