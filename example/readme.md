```
SELECT e.`events`, LEFT(receive_time, 10) AS "rt"
   	, count(DISTINCT e.user_id) AS "用户数"
   FROM event_db.event_wxafa45659f8ab8ff1 e
   	LEFT JOIN xingye_ts_snap.user_info u
   	ON e.user_id = u.user_id
   		AND e.app_id = u.app_id
   WHERE e.app_id = 'wxafa45659f8ab8ff1'
   	AND e.`events` = 'login_events'
   	AND u.app_id = 'wxafa45659f8ab8ff1'
   	AND STR_TO_DATE(CONCAT(e.y, '-', e.m, '-', e.d), '%Y-%m-%d') >= '2021-01-20'
   	AND STR_TO_DATE(CONCAT(e.y, '-', e.m, '-', e.d), '%Y-%m-%d') <= '2021-01-20'
   	AND SUBSTR(u.register_time, 1, 10) >= '2021-01-18'
   	AND SUBSTR(u.register_time, 1, 10) <= '2021-01-20'
   GROUP BY e.`events`, rt
   ORDER BY rt ASC, 用户数 DESC
```

```
SELECT user_id
   	, funnel_count_alpha(UNIX_TIMESTAMP(data_time), 300, `event_id`, 'event_0,event_1,event_2') AS max_ordered_match_length
   FROM (
   	SELECT user_id, data_time, 'event_0' AS `event_id`
   	FROM event_db.event_wxafa45659f8ab8ff1
   	WHERE events IN ('31133')
   		AND STR_TO_DATE(CONCAT(y, '-', m, '-', d), '%Y-%m-%d') >= '2021-01-20'
   		AND STR_TO_DATE(CONCAT(y, '-', m, '-', d), '%Y-%m-%d') <= '2021-01-20'
   		AND value2 IN ('3')
   		AND user_id IN (
   			SELECT DISTINCT user_id
   			FROM xingye_ts_snap.user_info
   			WHERE pp_id = 'wxafa45659f8ab8ff1'
   				AND SUBSTR(register_time, 1, 10) >= 'today-0'
   				AND SUBSTR(register_time, 1, 10) >= 'today-0'
   		)
   		AND user_id IN (
   			SELECT DISTINCT user_id
   			FROM xingye_ts_snap.user_info
   			WHERE pp_id = 'wxafa45659f8ab8ff1'
   				AND channel_id IN ('share_10009', 'share_10001')
   		)
   )
```