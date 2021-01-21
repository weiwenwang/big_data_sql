package noJoin

import (
	"fmt"
	"strconv"
	"strings"
)

type Table struct {
	Db_Table_Name string
	where         [][]interface{}
	field         []string
	group_by      []string
	order_by      []string
	Limit         int
}

func CreateNoJoinSqlObject(table string) (t Table) {
	t.Db_Table_Name = table
	t.where = make([][]interface{}, 0)
	t.field = make([]string, 0)
	t.group_by = make([]string, 0)
	t.order_by = make([]string, 0)
	return t
}


func (t *Table) Where(a string, b string, c interface{}) {
	t.where = append(t.where, []interface{}{a, b, c})
}

func (t *Table) Field(obj []string) {
	t.field = append(t.field, obj...)
}

func (t *Table) GetSql() (s string) {
	s = "select " + t.GetField() + " from " + t.Db_Table_Name + t.GetWhere()
	return s
}

func (t *Table) GetField() (str string) {
	for _, v := range t.field {
		str += "," + v
	}
	str = strings.TrimLeft(str, ",")
	return str
}

func (t *Table) GetWhere() (str string) {
	for _, v := range t.where {
		if strings.ToUpper(v[1].(string)) == "IN" {
			switch v[2].(type) {
			case string:
				str += " and " + v[0].(string) + " in " + "(" + v[2].(string) + ")"
			case []interface{}:
				one := ""
				for _, vv := range v[2].([]interface{}) {
					one += ",'" + vv.(string) + "'"
				}
				str += " and " + v[0].(string) + " in " + "(" + one[1:] + ")"
			case bool:
				fmt.Printf("%v is bool", v)
			}

		} else if strings.ToUpper(v[1].(string)) == ">" ||
			strings.ToUpper(v[1].(string)) == ">=" ||
			strings.ToUpper(v[1].(string)) == "<" ||
			strings.ToUpper(v[1].(string)) == "<=" ||
			strings.ToUpper(v[1].(string)) == "=" {
			if IsNum(v[2].(string)) {
				str += " and " + v[0].(string) + " " + v[1].(string) + " " + v[2].(string)
			} else {
				str += " and " + v[0].(string) + " " + v[1].(string) + " '" + v[2].(string) + "'"
			}
		}
	}
	if str == "" {
		return str
	} else {
		str = " where " + strings.TrimLeft(str, " and")
		return str
	}
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
