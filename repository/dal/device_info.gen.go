// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"bytetrack_server/repository/model"
)

func newDeviceInfo(db *gorm.DB, opts ...gen.DOOption) deviceInfo {
	_deviceInfo := deviceInfo{}

	_deviceInfo.deviceInfoDo.UseDB(db, opts...)
	_deviceInfo.deviceInfoDo.UseModel(&model.DeviceInfo{})

	tableName := _deviceInfo.deviceInfoDo.TableName()
	_deviceInfo.ALL = field.NewAsterisk(tableName)
	_deviceInfo.ID = field.NewInt32(tableName, "id")
	_deviceInfo.Name = field.NewString(tableName, "name")
	_deviceInfo.IP = field.NewString(tableName, "ip")
	_deviceInfo.Describe = field.NewString(tableName, "describe")

	_deviceInfo.fillFieldMap()

	return _deviceInfo
}

type deviceInfo struct {
	deviceInfoDo

	ALL      field.Asterisk
	ID       field.Int32
	Name     field.String
	IP       field.String
	Describe field.String

	fieldMap map[string]field.Expr
}

func (d deviceInfo) Table(newTableName string) *deviceInfo {
	d.deviceInfoDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deviceInfo) As(alias string) *deviceInfo {
	d.deviceInfoDo.DO = *(d.deviceInfoDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deviceInfo) updateTableName(table string) *deviceInfo {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt32(table, "id")
	d.Name = field.NewString(table, "name")
	d.IP = field.NewString(table, "ip")
	d.Describe = field.NewString(table, "describe")

	d.fillFieldMap()

	return d
}

func (d *deviceInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deviceInfo) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 4)
	d.fieldMap["id"] = d.ID
	d.fieldMap["name"] = d.Name
	d.fieldMap["ip"] = d.IP
	d.fieldMap["describe"] = d.Describe
}

func (d deviceInfo) clone(db *gorm.DB) deviceInfo {
	d.deviceInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deviceInfo) replaceDB(db *gorm.DB) deviceInfo {
	d.deviceInfoDo.ReplaceDB(db)
	return d
}

type deviceInfoDo struct{ gen.DO }

func (d deviceInfoDo) Debug() *deviceInfoDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceInfoDo) WithContext(ctx context.Context) *deviceInfoDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceInfoDo) ReadDB() *deviceInfoDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceInfoDo) WriteDB() *deviceInfoDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceInfoDo) Session(config *gorm.Session) *deviceInfoDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceInfoDo) Clauses(conds ...clause.Expression) *deviceInfoDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceInfoDo) Returning(value interface{}, columns ...string) *deviceInfoDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceInfoDo) Not(conds ...gen.Condition) *deviceInfoDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceInfoDo) Or(conds ...gen.Condition) *deviceInfoDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceInfoDo) Select(conds ...field.Expr) *deviceInfoDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceInfoDo) Where(conds ...gen.Condition) *deviceInfoDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceInfoDo) Order(conds ...field.Expr) *deviceInfoDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceInfoDo) Distinct(cols ...field.Expr) *deviceInfoDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceInfoDo) Omit(cols ...field.Expr) *deviceInfoDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceInfoDo) Join(table schema.Tabler, on ...field.Expr) *deviceInfoDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *deviceInfoDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *deviceInfoDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceInfoDo) Group(cols ...field.Expr) *deviceInfoDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceInfoDo) Having(conds ...gen.Condition) *deviceInfoDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceInfoDo) Limit(limit int) *deviceInfoDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceInfoDo) Offset(offset int) *deviceInfoDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *deviceInfoDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceInfoDo) Unscoped() *deviceInfoDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceInfoDo) Create(values ...*model.DeviceInfo) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceInfoDo) CreateInBatches(values []*model.DeviceInfo, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceInfoDo) Save(values ...*model.DeviceInfo) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceInfoDo) First() (*model.DeviceInfo, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceInfo), nil
	}
}

func (d deviceInfoDo) Take() (*model.DeviceInfo, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceInfo), nil
	}
}

func (d deviceInfoDo) Last() (*model.DeviceInfo, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceInfo), nil
	}
}

func (d deviceInfoDo) Find() ([]*model.DeviceInfo, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeviceInfo), err
}

func (d deviceInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeviceInfo, err error) {
	buf := make([]*model.DeviceInfo, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceInfoDo) FindInBatches(result *[]*model.DeviceInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceInfoDo) Attrs(attrs ...field.AssignExpr) *deviceInfoDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceInfoDo) Assign(attrs ...field.AssignExpr) *deviceInfoDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceInfoDo) Joins(fields ...field.RelationField) *deviceInfoDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceInfoDo) Preload(fields ...field.RelationField) *deviceInfoDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceInfoDo) FirstOrInit() (*model.DeviceInfo, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceInfo), nil
	}
}

func (d deviceInfoDo) FirstOrCreate() (*model.DeviceInfo, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceInfo), nil
	}
}

func (d deviceInfoDo) FindByPage(offset int, limit int) (result []*model.DeviceInfo, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deviceInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceInfoDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceInfoDo) Delete(models ...*model.DeviceInfo) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceInfoDo) withDO(do gen.Dao) *deviceInfoDo {
	d.DO = *do.(*gen.DO)
	return d
}
