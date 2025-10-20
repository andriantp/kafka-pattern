package postgres

import (
	"context"
	"kafka-pattern/adapter/database/model"
	"kafka-pattern/logger"
	"log"

	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	ormlog "gorm.io/gorm/logger"
)

type Setting struct {
	Migrate bool
	Dsn     string
}

type postgresDb struct {
	db *gorm.DB
}

type DatabaseI interface {
	Db(ctx context.Context) interface{}
	WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context, dbt *gorm.DB) error) error
}

func NewPostgres(setting Setting) (DatabaseI, error) {
	newLogger := ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		ormlog.Config{
			SlowThreshold:             5 * time.Second, // Slow SQL threshold
			Colorful:                  true,            // Disable color
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			LogLevel:                  ormlog.Error,    // Log level
		},
	)

	db, err := gorm.Open(postgres.Open(setting.Dsn), &gorm.Config{
		Logger:  newLogger,
		NowFunc: func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return postgresDb{db: db}, err
	}

	if setting.Migrate {
		db.AutoMigrate(
			&model.Products{},
		)
	}

	return postgresDb{db: db}, nil
}

func (d postgresDb) Db(ctx context.Context) interface{} {
	tx := ctx.Value("txContext")
	if tx == nil {
		return d.db
	}
	return tx.(*gorm.DB)
}

func (d postgresDb) WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context, dbt *gorm.DB) error) error {
	tx := d.db.Begin()
	ctxWithTx := context.WithValue(ctx, tx, "txContext")
	errFn := fn(ctxWithTx, tx)
	if errFn != nil {
		errRlbck := tx.Rollback().Error
		if errRlbck != nil {
			logger.Level("error", "WithTransaction", "failed on rollback transaction:"+errRlbck.Error())
		}
		return errFn
	}

	errCmmt := tx.Commit().Error
	if errCmmt != nil {
		logger.Level("error", "WithTransaction", "failed on commit transaction:"+errCmmt.Error())
	}
	return errFn
}
