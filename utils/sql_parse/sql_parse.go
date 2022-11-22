package sql_parse

import (
	"fmt"
	"github.com/wonderivan/logger"
	"strings"
	"sync"
)

type SqlParse struct {
	limit  int
	offset int
	where  []string
	table  string
	fields []string
	join   string
}

var (
	instance *SqlParse
	once     sync.Once
)

func NewSqlParse() *SqlParse {
	once.Do(func() {
		instance = new(SqlParse)
	})
	return instance
}

func (s *SqlParse) Table(table string) *SqlParse {
	s.table = table
	return s
}

func (s *SqlParse) Limit(limit int) *SqlParse {
	s.limit = limit
	return s
}

func (s *SqlParse) Offset(offset int) *SqlParse {
	s.offset = offset
	return s
}

func (s *SqlParse) Select(fields []string) *SqlParse {
	s.fields = fields
	return s
}

func (s *SqlParse) Join(join string) *SqlParse {
	s.join = join
	return s
}

func (s *SqlParse) Where(key string, value string) *SqlParse {
	where := key + value
	s.where = append(s.where, where)
	return s
}

func (s *SqlParse) clean() {
	fields := make([]string, 0)
	s.where = fields
	s.table = ""
	s.limit = 0
	s.offset = 0
	s.fields = fields
}

func (s *SqlParse) getFields() string {
	fields := "*"
	logger.Info("fields=%+v",s.fields)
	if len(s.fields) > 0 {
		fields = strings.Join(s.fields, ",")
	}
	return fields
}

func (s *SqlParse) getWhere() string {
	whereStr := ""
	if len(s.where) > 0 {
		where := strings.Join(s.where, " AND ")
		whereStr += fmt.Sprintf(" WHERE %s", where)
	}
	return whereStr
}

func (s *SqlParse) getLimit() string {
	limit := ""
	if s.limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d,%d", s.offset, s.limit)
	}
	return limit
}

func (s *SqlParse) getJoin() string {
	if s.join != "" {
		return s.join
	}
	return ""
}

func (s *SqlParse) WrapCharacter(v interface{}) string {
	switch v.(type) {
	case string:
		return fmt.Sprintf("'%s'", v.(string))
	case int:
		return fmt.Sprintf("%d", v.(int))
	case int64:
		return fmt.Sprintf("%d", v.(int64))
	}
	return ""
}

func (s *SqlParse) parseSql(sqlType string, value ...interface{}) string {
	sql := ""
	switch sqlType {
	case "GET":
		sql += fmt.Sprintf("SELECT %s FROM %s ", s.getFields(), s.table)
		sql += s.getJoin()
		sql += s.getWhere()
		sql += s.getLimit()
		return sql
	case "COUNT":
		sql += fmt.Sprintf("SELECT COUNT(*) AS count FROM %s", s.table)
		sql += s.getWhere()
		sql += s.getLimit()
		return sql
	case "INSERT":
		if len(value) > 0 {
			insert := value[0].(map[string]interface{})
			fields := []string{}
			values := []string{}
			for k, v := range insert {
				fields = append(fields, fmt.Sprintf("`%s`", k))
				values = append(values, s.WrapCharacter(v))
			}
			sql += fmt.Sprintf("INSERT INTO %s(%s)VALUES(%s)", s.table, strings.Join(fields, ","), strings.Join(values, ","))
			return sql
		}
	case "UPDATE":
		if len(value) > 0 {
			sql := fmt.Sprintf("UPDATE %s SET ", s.table)
			fields := []string{}
			for k, v := range value[0].(map[string]interface{}) {
				equal := fmt.Sprintf("`%s`", k) + "=" + s.WrapCharacter(v)
				fields = append(fields, equal)
			}
			sql += strings.Join(fields, ",")
			sql += s.getWhere()
			return sql
		}
	case "DELETE":
		where := s.getWhere()
		if where != "" {
			sql := fmt.Sprintf("DELETE FROM %s %s", s.table, where)
			return sql
		}
	}
	return ""
}

func (s *SqlParse) Get() string {
	sql := s.parseSql("GET")
	s.clean()
	return sql
}

func (s *SqlParse) Count() string {
	sql := s.parseSql("COUNT")
	s.clean()
	return sql
}

func (s *SqlParse) Delete() string {
	sql := s.parseSql("DELETE")
	s.clean()
	return sql
}

func (s *SqlParse) Update(updateData map[string]interface{}) string {
	sql := s.parseSql("UPDATE", updateData)
	s.clean()
	return sql
}

func (s *SqlParse) Insert(insert map[string]interface{}) string {
	sql := s.parseSql("INSERT", insert)
	s.clean()
	return sql
}
