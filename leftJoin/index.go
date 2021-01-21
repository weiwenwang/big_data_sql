package leftJoin

import (
	"admin-reactor/test/common"
	"strings"
)

type Table struct {
	Left_table_name        string
	Right_table_name       string
	Left_table_alias_name  string
	Right_table_alias_name string
	where                  [][]interface{}
	field                  []string
	on                     [][]string
	group_by               []string
	order_by               []string
	Limit                  int
}

func CreateLeftJoinSqlObject(left_table string, right_table string, left_table_alias string, right_table_alias string) (t Table) {
	t.Left_table_name = left_table
	t.Right_table_name = right_table
	t.Left_table_alias_name = left_table_alias
	t.Right_table_alias_name = right_table_alias

	t.where = make([][]interface{}, 0)
	t.field = make([]string, 0)
	t.on = make([][]string, 0)
	t.group_by = make([]string, 0)
	t.order_by = make([]string, 0)

	return t
}

func (t *Table) Where(a string, b string, c interface{}) {
	t.where = append(t.where, []interface{}{a, b, c})
}

func (t *Table) On(a string, b string) {
	t.on = append(t.on, []string{a, b})
}

func (t *Table) GroupBy(a []string) {
	t.group_by = append(t.group_by, a...)
}

func (t *Table) OrderBy(a []string) {
	t.order_by = append(t.order_by, a...)
}

func (t *Table) Field(obj []string) {
	t.field = append(t.field, obj...)
}

func (t Table) GetSql() (s string) {
	s = "select " + t.getField() + " from " + t.Left_table_name + " as " + t.Left_table_alias_name + " left join " + t.Right_table_name + " as " +
		t.Right_table_alias_name + " on " + t.getOn() + t.getWhere() + t.getGroupBy() + t.getOrder()
	return s
}

func (t *Table) getWhere() (str string) {
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
			}

		} else if strings.ToUpper(v[1].(string)) == ">" ||
			strings.ToUpper(v[1].(string)) == ">=" ||
			strings.ToUpper(v[1].(string)) == "<" ||
			strings.ToUpper(v[1].(string)) == "<=" ||
			strings.ToUpper(v[1].(string)) == "=" {
			if common.IsNum(v[2].(string)) {
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

func (t *Table) getOn() (on string) {
	for _, v := range t.on {
		on += " and " + v[0] + " = " + v[1]
	}
	on = strings.TrimLeft(on, " and")
	return on
}

func (t Table) getField() (str string) {
	for _, v := range t.field {
		str += "," + v
	}
	return strings.TrimLeft(str, ",")
}

func (t *Table) getOrder() (order_by string) {
	for k, v := range t.order_by {
		if k == 0 {
			order_by += "," + v
		} else {
			order_by += "," + v
		}
	}
	if strings.TrimLeft(order_by, ",", ) != "" {
		order_by = " order by " + strings.TrimLeft(order_by, ",", )
	} else {
		order_by = ""
	}
	return order_by
}

func (t *Table) getGroupBy() (str string) {
	for _, v := range t.group_by {
		str += "," + v
	}
	if strings.TrimLeft(str, ",", ) != "" {
		str = " group by " + strings.TrimLeft(str, ",", )
	} else {
		str = ""
	}
	return str
}
