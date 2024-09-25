package gorm

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type (
	transactionDB struct {
		db     *gorm.DB
		isDone bool
	}
	transactionCtx struct{}
)

func (gc *gormClient) GetDB(ctx context.Context) *gorm.DB {
	if tx := gc.getTx(ctx); tx != nil {
		return tx.db.WithContext(ctx)
	}

	return gc.db.WithContext(ctx)
}

func (gc *gormClient) Begin(ctx context.Context) (context.Context, func(), error) {
	if gc.getTx(ctx) != nil {
		return ctx, func() {}, errors.New("transaction already created")
	}

	db := gc.db.Begin()
	if db.Error != nil {
		return ctx, func() {}, fmt.Errorf("could not create transaction: %s", db.Error.Error())
	}

	tx := transactionDB{db: db}
	ctx = gc.setTx(ctx, &tx)

	txCallback := func() {
		if !tx.isDone {
			db.Rollback()
		}
	}

	return ctx, txCallback, nil
}

func (gc *gormClient) Commit(ctx context.Context) error {
	tx := gc.getTx(ctx)
	if tx == nil {
		return errors.New("no transaction in context")
	}

	if err := tx.db.Commit().Error; err != nil {
		return fmt.Errorf("could not commit transaction: %s", err.Error())
	}

	tx.isDone = true
	return nil
}

func (gc *gormClient) Rollback(ctx context.Context) error {
	tx := gc.getTx(ctx)
	if tx == nil {
		return errors.New("no transaction in context")
	}

	if err := tx.db.Rollback().Error; err != nil {
		return fmt.Errorf("could not rollback transaction: %s", err.Error())
	}

	tx.isDone = true
	return nil
}

func (gc *gormClient) getTx(ctx context.Context) *transactionDB {
	if tx, ok := ctx.Value(transactionCtx{}).(*transactionDB); ok {
		return tx
	}

	return nil
}

func (gc *gormClient) setTx(ctx context.Context, tx *transactionDB) context.Context {
	return context.WithValue(ctx, transactionCtx{}, tx)
}
