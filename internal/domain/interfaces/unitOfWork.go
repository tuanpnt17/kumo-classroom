package interfaces

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type IUnitOfWork interface {
	Begin(opts ...*sql.TxOptions) *gorm.DB
	SavePoint(name string) *gorm.DB
	RollbackTo(name string) *gorm.DB
	Rollback() *gorm.DB
	Commit() *gorm.DB

	Create(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	WithContext(ctx context.Context) (tx *gorm.DB)

	FindOne(ctx context.Context, dest interface{}, conds interface{}) error
	FindAll(ctx context.Context, dest interface{}, conds interface{}) error
	CountRows(ctx context.Context, conds interface{}) (int, error)

	InsertOne(ctx context.Context, src interface{}) error
	UpdateOne(ctx context.Context, old interface{}, new interface{}) error
	DeleteOne(ctx context.Context, dest interface{}, arg interface{}) error
}
