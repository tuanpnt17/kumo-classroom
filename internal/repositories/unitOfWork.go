package repositories

import (
	"context"
	"database/sql"
	"github.com/tuanpnt17/kumo-classroom-be/global"
	"github.com/tuanpnt17/kumo-classroom-be/internal/domain/interfaces"
	"gorm.io/gorm"
)

type unitOfWork struct {
	// DB *gorm.DB
}

func NewUnitOfWork() interfaces.IUnitOfWork {
	return &unitOfWork{}
}
func (u unitOfWork) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return global.AppDB.Begin(opts...)
}

func (u unitOfWork) SavePoint(name string) *gorm.DB {
	return global.AppDB.SavePoint(name)
}

func (u unitOfWork) RollbackTo(name string) *gorm.DB {
	return global.AppDB.RollbackTo(name)
}

func (u unitOfWork) Rollback() *gorm.DB {
	return global.AppDB.Rollback()
}

func (u unitOfWork) Commit() *gorm.DB {
	return global.AppDB.Commit()
}

func (u unitOfWork) Create(value interface{}) (tx *gorm.DB) {
	return global.AppDB.Create(value)
}

func (u unitOfWork) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return global.AppDB.First(dest, conds...)
}

func (u unitOfWork) Select(query interface{}, args ...interface{}) (tx *gorm.DB) {
	return global.AppDB.Select(query, args...)
}

func (u unitOfWork) Where(query interface{}, args ...interface{}) (tx *gorm.DB) {
	return global.AppDB.Where(query, args...)
}

func (u unitOfWork) WithContext(ctx context.Context) (tx *gorm.DB) {
	return global.AppDB.WithContext(ctx)
}

func (u unitOfWork) FindOne(ctx context.Context, dest interface{}, conds interface{}) error {
	result := global.AppDB.WithContext(ctx).Where(conds).First(dest)
	return result.Error
}

func (u unitOfWork) FindAll(ctx context.Context, dest interface{}, conds interface{}) error {
	result := global.AppDB.WithContext(ctx).Where(conds).Find(dest)
	return result.Error
}

func (u unitOfWork) CountRows(ctx context.Context, conds interface{}) (int, error) {
	var count int64
	result := global.AppDB.WithContext(ctx).Model(conds).Where(conds).Count(&count)
	return int(count), result.Error
}

func (u unitOfWork) InsertOne(ctx context.Context, src interface{}) error {
	result := global.AppDB.WithContext(ctx).Create(src)
	return result.Error
}

func (u unitOfWork) UpdateOne(ctx context.Context, old interface{}, new interface{}) error {
	result := global.AppDB.WithContext(ctx).Model(old).Updates(new)
	return result.Error
}

func (u unitOfWork) DeleteOne(ctx context.Context, dest interface{}, arg interface{}) error {
	result := global.AppDB.WithContext(ctx).Delete(arg)
	return result.Error
}
