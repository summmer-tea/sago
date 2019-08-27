package db

import "github.com/go-xorm/xorm"

//提供封装的方式,保证调用风格统一
var MysqlDialect *DbDialect

type Session = xorm.Session
type DbDialect struct{}

func (m *DbDialect) Where(query interface{}, args ...interface{}) *xorm.Session {
	return MysqlEngine.Where(query, args)
}

func (m *DbDialect) Limit(limit int) *xorm.Session {
	return MysqlEngine.Limit(limit)
}

func (m *DbDialect) OrderBy(order string) *xorm.Session {
	return MysqlEngine.OrderBy(order)
}

func (m *DbDialect) Select(str string) *xorm.Session {
	return MysqlEngine.Select(str)
}

func (m *DbDialect) Count(bean ...interface{}) (int64, error) {
	return MysqlEngine.Count(bean)
}

func (m *DbDialect) Find(bean ...interface{}) error {
	return MysqlEngine.Find(bean)
}

func (m *DbDialect) FindBySql(sql string, items []interface{}) error {
	return MysqlEngine.SQL(sql).Find(&items)
}

func (m *DbDialect) GetById(id int, item interface{}) (bool, error) {
	return MysqlEngine.Id(id).Get(item)
}

func (m *DbDialect) UpdateById(id int, item interface{}) (int64, error) {
	return MysqlEngine.Id(id).Update(item)
}

func (m *DbDialect) Insert(bean ...interface{}) (int64, error) {
	return MysqlEngine.Insert(bean)
}
