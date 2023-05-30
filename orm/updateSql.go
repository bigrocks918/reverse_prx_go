package orm

import (
	"strings"
)

func updateSqlStr(sqlData *SqlData) string {
	str := sqlData.NameValue
	primary := sqlData.Key
	index := strings.LastIndex(str, ",")
	if index != -1 {
		str = str[:index]
	}

	index = strings.LastIndex(primary, ",")
	if index != -1 {
		primary = primary[:index]
	}
	primary = strings.Replace(primary, ",", " and ", -1)
	return "update " + sqlData.Table + " set " + str + " where " + primary
}

// --- struct to sql
func UpdateSqlStr(obj interface{}, params ...OpOption) string {
	op := &Op{sqlType: SQLTYPE_UPDATE}
	op.applyOpts(params)
	sqlData := &SqlData{}
	getTableName(obj, sqlData)
	parseStructSql(obj, sqlData, op)
	return updateSqlStr(sqlData)
}

func UpdateSql(obj interface{}, params ...OpOption) bool {
	str := UpdateSqlStr(obj, params...)
	return exec(str) == nil
}
