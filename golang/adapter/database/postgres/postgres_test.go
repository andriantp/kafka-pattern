package postgres_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"kafka-pattern/adapter/database/model"
	"kafka-pattern/adapter/database/postgres"
	"testing"

	"gorm.io/gorm"
)

func Test_Connect(t *testing.T) {
	setting := postgres.Setting{
		Migrate: true,
		Dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "54321",
			"inventory", "postgres", "postgres"),
	}

	_, err := postgres.NewPostgres(setting)
	if err != nil {
		t.Fatalf("NewPostgres:%v", err)
	}
}

func Test_ReadALL(t *testing.T) {
	ctx := context.Background()
	setting := postgres.Setting{
		Migrate: true,
		Dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "54321",
			"inventory", "postgres", "postgres"),
	}

	pg, err := postgres.NewPostgres(setting)
	if err != nil {
		t.Fatalf("NewPostgres:%v", err)
	}

	data := new([]model.Products)
	db := pg.Db(ctx).(*gorm.DB)
	result := db.Table(model.Products{}.TableName()).Scan(data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		t.Fatalf("record not found")
	}

	if len(*data) == 0 {
		t.Fatalf("no data")
	}

	for i, v := range *data {
		js, _ := json.MarshalIndent(v, "", "  ")
		t.Logf("Data [%d]: %s", i+1, string(js))
	}
}

func Test_FindProduct(t *testing.T) {
	ctx := context.Background()
	setting := postgres.Setting{
		Migrate: true,
		Dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "54321",
			"inventory", "postgres", "postgres"),
	}

	pg, err := postgres.NewPostgres(setting)
	if err != nil {
		t.Fatalf("NewPostgres:%v", err)
	}

	name := "Apple"
	t.Logf("Find by name : %s", name)
	data := new(model.Products)
	db := pg.Db(ctx).(*gorm.DB)
	result := db.Table(model.Products{}.TableName()).Where("name = ?", name).Find(data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		t.Fatalf("record not found")
	}

	js, _ := json.MarshalIndent(data, "", "  ")
	t.Logf("Data : %s", string(js))

}

func Test_Insert(t *testing.T) {
	ctx := context.Background()
	setting := postgres.Setting{
		Migrate: true,
		Dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "54321",
			"inventory", "postgres", "postgres"),
	}

	pg, err := postgres.NewPostgres(setting)
	if err != nil {
		t.Fatalf("NewPostgres:%v", err)
	}

	data := model.Products{
		Name:  "Manggo",
		Price: 0.25,
		Stock: 25,
	}
	errTx := pg.WithTransaction(ctx, func(ctxWithTx context.Context, db *gorm.DB) error {
		if err := data.Create(db); err != nil {
			return fmt.Errorf("failed at Create:%w", err)
		}
		return nil
	})
	if errTx != nil {
		t.Fatalf("WithTransaction:%v", errTx)
	}
}

func Test_Update(t *testing.T) {
	ctx := context.Background()
	setting := postgres.Setting{
		Migrate: true,
		Dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "54321",
			"inventory", "postgres", "postgres"),
	}

	pg, err := postgres.NewPostgres(setting)
	if err != nil {
		t.Fatalf("NewPostgres:%v", err)
	}

	name := "Manggo"
	t.Logf("Find by name : %s", name)
	dataBefore := new(model.Products)
	db := pg.Db(ctx).(*gorm.DB)
	result := db.Table(model.Products{}.TableName()).Where("name = ?", name).Find(dataBefore)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		t.Fatalf("record not found")
	}
	js, _ := json.MarshalIndent(dataBefore, "", "  ")
	t.Logf("Before update : %s", string(js))

	newStock := uint(50)
	t.Logf("Change stock, from [%d] to [%d]", dataBefore.Stock, newStock)
	dataBefore.Stock = newStock
	errTx := pg.WithTransaction(ctx, func(ctxWithTx context.Context, db *gorm.DB) error {
		if err := dataBefore.Update(db); err != nil {
			return fmt.Errorf("failed at Update:%w", err)
		}
		return nil
	})
	if errTx != nil {
		t.Fatalf("WithTransaction:%v", errTx)
	}

	dataAfter := new(model.Products)
	result = db.Table(model.Products{}.TableName()).Where("name = ?", name).Find(dataAfter)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		t.Fatalf("record not found")
	}
	js, _ = json.MarshalIndent(dataAfter, "", "  ")
	t.Logf("After update : %s", string(js))
}

func Test_SoftDelete(t *testing.T) {
	ctx := context.Background()
	setting := postgres.Setting{
		Migrate: true,
		Dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "54321",
			"inventory", "postgres", "postgres"),
	}

	pg, err := postgres.NewPostgres(setting)
	if err != nil {
		t.Fatalf("NewPostgres:%v", err)
	}

	data := new(model.Products)
	data.ID = 7
	t.Logf("Soft Delete by ID : %v", data.ID)
	errTx := pg.WithTransaction(ctx, func(ctxWithTx context.Context, db *gorm.DB) error {
		if err := data.Delete(db); err != nil {
			return fmt.Errorf("failed at Delete:%w", err)
		}
		return nil
	})
	if errTx != nil {
		t.Fatalf("WithTransaction:%v", errTx)
	}
}

func Test_Remove(t *testing.T) {
	ctx := context.Background()
	setting := postgres.Setting{
		Migrate: true,
		Dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", "localhost", "54321",
			"inventory", "postgres", "postgres"),
	}

	pg, err := postgres.NewPostgres(setting)
	if err != nil {
		t.Fatalf("NewPostgres:%v", err)
	}

	data := new(model.Products)
	data.ID = 7
	t.Logf("Remove by ID : %v", data.ID)
	errTx := pg.WithTransaction(ctx, func(ctxWithTx context.Context, db *gorm.DB) error {
		if err := data.Remove(db); err != nil {
			return fmt.Errorf("failed at Remove:%w", err)
		}
		return nil
	})
	if errTx != nil {
		t.Fatalf("WithTransaction:%v", errTx)
	}
}
