package main

import (
	"fmt"
	"github.com/weiwenwang/big_data_sql/leftJoin"
	"github.com/weiwenwang/big_data_sql/noJoin"
)

func main() {
	//doLeftJoin()
	doNoJoin()
}

func doLeftJoin() {
	s := leftJoin.CreateLeftJoinSqlObject("event_db.event_wxafa45659f8ab8ff1", "xingye_ts_snap.user_info", "e", "u")
	s.Field([]string{"e.`events`", "LEFT(receive_time, 10) AS 'rt'", "count(DISTINCT e.user_id) AS '用户数'"})
	s.Where("e.app_id", "=", "wxafa45659f8ab8ff1")
	s.Where("e.`events`", "=", "login_events")
	s.Where("u.app_id", "=", "wxafa45659f8ab8ff1")
	s.Where("STR_TO_DATE(CONCAT(e.y, '-', e.m, '-', e.d), '%Y-%m-%d')", ">=", "2021-01-20")
	s.Where("STR_TO_DATE(CONCAT(e.y, '-', e.m, '-', e.d), '%Y-%m-%d')", "<=", "2021-01-20")
	s.Where("SUBSTR(u.register_time, 1, 10)", ">=", "2021-01-18")
	s.Where("SUBSTR(u.register_time, 1, 10)", "<=", "2021-01-20")
	s.On("e.user_id", "u.user_id")
	s.On("e.app_id", "u.app_id")
	s.GroupBy([]string{"e.`events`", "rt"})
	s.OrderBy([]string{"rt ASC", "用户数 DESC"})

	sql_str := s.GetSql()
	fmt.Println(sql_str)
}

func doNoJoin() {
	t2 := noJoin.CreateNoJoinSqlObject("xingye_ts_snap.user_info")
	t2.Field([]string{"DISTINCT user_id"})

	t2.Where("app_id", "=", "wxafa45659f8ab8ff1")
	t2.Where("SUBSTR(register_time, 1, 10)", ">=", "today-0")
	t2.Where("SUBSTR(register_time, 1, 10)", ">=", "today-0")

	t3 := noJoin.CreateNoJoinSqlObject("xingye_ts_snap.user_info")
	t3.Field([]string{"DISTINCT user_id"})
	t3.Where("app_id", "=", "wxafa45659f8ab8ff1")
	t3.Where("channel_id", "in", []interface{}{"share_10009", "share_10001"})

	t := noJoin.CreateNoJoinSqlObject("event_db.event_wxafa45659f8ab8ff1")
	t.Field([]string{"user_id", "data_time", "'event_0' AS `event_id`"})

	t.Where("events", "in", []interface{}{"31133"})
	t.Where("STR_TO_DATE(CONCAT(y, '-', m, '-', d), '%Y-%m-%d')", ">=", "2021-01-20")
	t.Where("STR_TO_DATE(CONCAT(y, '-', m, '-', d), '%Y-%m-%d')", "<=", "2021-01-20")
	t.Where("value2", "in", []interface{}{"3"})
	t.Where("user_id", "in", t2.GetSql())
	t.Where("user_id", "in", t3.GetSql())

	t4 := noJoin.CreateNoJoinSqlObject("(" + t.GetSql() + ")")
	t4.Field([]string{"user_id", "funnel_count_alpha(UNIX_TIMESTAMP(data_time), 300, `event_id`, 'event_0,event_1,event_2') AS max_ordered_match_length"})
	fmt.Println(t4.GetSql())
}
