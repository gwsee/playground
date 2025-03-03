// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gorm.io/playground/dal/model"
)

func newAccountInfo(db *gorm.DB, opts ...gen.DOOption) accountInfo {
	_accountInfo := accountInfo{}

	_accountInfo.accountInfoDo.UseDB(db, opts...)
	_accountInfo.accountInfoDo.UseModel(&model.AccountInfo{})

	tableName := _accountInfo.accountInfoDo.TableName()
	_accountInfo.ALL = field.NewAsterisk(tableName)
	_accountInfo.ID = field.NewInt64(tableName, "id")
	_accountInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_accountInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_accountInfo.DeletedAt = field.NewField(tableName, "deleted_at")
	_accountInfo.UserID = field.NewInt64(tableName, "user_id")
	_accountInfo.Number = field.NewString(tableName, "number")

	_accountInfo.fillFieldMap()

	return _accountInfo
}

type accountInfo struct {
	accountInfoDo accountInfoDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserID    field.Int64
	Number    field.String

	fieldMap map[string]field.Expr
}

func (a accountInfo) Table(newTableName string) *accountInfo {
	a.accountInfoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accountInfo) As(alias string) *accountInfo {
	a.accountInfoDo.DO = *(a.accountInfoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accountInfo) updateTableName(table string) *accountInfo {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.UserID = field.NewInt64(table, "user_id")
	a.Number = field.NewString(table, "number")

	a.fillFieldMap()

	return a
}

func (a *accountInfo) WithContext(ctx context.Context) *accountInfoDo {
	return a.accountInfoDo.WithContext(ctx)
}

func (a accountInfo) TableName() string { return a.accountInfoDo.TableName() }

func (a accountInfo) Alias() string { return a.accountInfoDo.Alias() }

func (a accountInfo) Columns(cols ...field.Expr) gen.Columns { return a.accountInfoDo.Columns(cols...) }

func (a *accountInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accountInfo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["number"] = a.Number
}

func (a accountInfo) clone(db *gorm.DB) accountInfo {
	a.accountInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accountInfo) replaceDB(db *gorm.DB) accountInfo {
	a.accountInfoDo.ReplaceDB(db)
	return a
}

type accountInfoDo struct{ gen.DO }

func (a accountInfoDo) Debug() *accountInfoDo {
	return a.withDO(a.DO.Debug())
}

func (a accountInfoDo) WithContext(ctx context.Context) *accountInfoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountInfoDo) ReadDB() *accountInfoDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountInfoDo) WriteDB() *accountInfoDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountInfoDo) Session(config *gorm.Session) *accountInfoDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountInfoDo) Clauses(conds ...clause.Expression) *accountInfoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountInfoDo) Returning(value interface{}, columns ...string) *accountInfoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountInfoDo) Not(conds ...gen.Condition) *accountInfoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountInfoDo) Or(conds ...gen.Condition) *accountInfoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountInfoDo) Select(conds ...field.Expr) *accountInfoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountInfoDo) Where(conds ...gen.Condition) *accountInfoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountInfoDo) Order(conds ...field.Expr) *accountInfoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountInfoDo) Distinct(cols ...field.Expr) *accountInfoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountInfoDo) Omit(cols ...field.Expr) *accountInfoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountInfoDo) Join(table schema.Tabler, on ...field.Expr) *accountInfoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *accountInfoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *accountInfoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountInfoDo) Group(cols ...field.Expr) *accountInfoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountInfoDo) Having(conds ...gen.Condition) *accountInfoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountInfoDo) Limit(limit int) *accountInfoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountInfoDo) Offset(offset int) *accountInfoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *accountInfoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountInfoDo) Unscoped() *accountInfoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountInfoDo) Create(values ...*model.AccountInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountInfoDo) CreateInBatches(values []*model.AccountInfo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountInfoDo) Save(values ...*model.AccountInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountInfoDo) First() (*model.AccountInfo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) Take() (*model.AccountInfo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) Last() (*model.AccountInfo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) Find() ([]*model.AccountInfo, error) {
	result, err := a.DO.Find()
	return result.([]*model.AccountInfo), err
}

func (a accountInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AccountInfo, err error) {
	buf := make([]*model.AccountInfo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountInfoDo) FindInBatches(result *[]*model.AccountInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountInfoDo) Attrs(attrs ...field.AssignExpr) *accountInfoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountInfoDo) Assign(attrs ...field.AssignExpr) *accountInfoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountInfoDo) Joins(fields ...field.RelationField) *accountInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountInfoDo) Preload(fields ...field.RelationField) *accountInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountInfoDo) FirstOrInit() (*model.AccountInfo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) FirstOrCreate() (*model.AccountInfo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) FindByPage(offset int, limit int) (result []*model.AccountInfo, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a accountInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountInfoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountInfoDo) Delete(models ...*model.AccountInfo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountInfoDo) withDO(do gen.Dao) *accountInfoDo {
	a.DO = *do.(*gen.DO)
	return a
}
