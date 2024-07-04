package db

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"

	// ent policy runtime
	_ "github.com/NpoolPlatform/order-middleware/pkg/db/ent/runtime"
	"github.com/google/uuid"
)

func client() (*ent.Client, error) {
	conn, err := mysql.GetConn()
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.MySQL, conn)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Init(hooks ...ent.Hook) error {
	cli, err := client()
	if err != nil {
		return wlog.WrapError(err)
	}
	cli.Use(hooks...)
	return cli.Schema.Create(context.Background())
}

func Client() (*ent.Client, error) {
	return client()
}

func txRun(ctx context.Context, tx *ent.Tx, fn func(ctx context.Context, tx *ent.Tx) error) error {
	runUuid := uuid.New() //nolint:stylecheck
	logger.Sugar().Infow("txRun start", "RunUuid", runUuid, "Tx", tx)
	succ := false
	defer func() {
		logger.Sugar().Infow("txRun end", "Success", succ, "RunUuid", runUuid)
		if !succ {
			err := tx.Rollback()
			if err != nil {
				logger.Sugar().Errorf("fail rollback: %v", err)
				return
			}
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}

	if err := tx.Commit(); err != nil {
		return wlog.Errorf("committing transaction: %v", err)
	}

	succ = true
	return nil
}

func WithTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	cli, err := Client()
	if err != nil {
		return wlog.WrapError(err)
	}
	tx, err := cli.Debug().Tx(ctx)
	if err != nil {
		return wlog.Errorf("fail get client transaction: %v", err)
	}
	return txRun(ctx, tx, fn)
}

func WithDebugTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	cli, err := Client()
	if err != nil {
		return wlog.WrapError(err)
	}
	tx, err := cli.Debug().Tx(ctx)
	if err != nil {
		return wlog.Errorf("fail get client transaction: %v", err)
	}
	return txRun(ctx, tx, fn)
}

func WithClient(ctx context.Context, fn func(ctx context.Context, cli *ent.Client) error) error {
	cli, err := Client()
	if err != nil {
		return wlog.Errorf("fail get db client: %v", err)
	}

	runUuid := uuid.New() //notlint:stylecheck
	logger.Sugar().Infow("ClientRun start", "RunUuid", runUuid, "Client", cli)
	defer logger.Sugar().Infow("ClientRun done", "RunUuid", runUuid)

	if err := fn(ctx, cli.Debug()); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}
