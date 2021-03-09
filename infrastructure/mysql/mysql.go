/**
 * @Title  mysql操作层
 * @Description mysql操作的封装
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:08
 */
package mysql

import (
	"github.com/alecthomas/log4go"
)

func Init(mysqlConfig MConfig) {
	initMysql(mysqlConfig)
}

type Query struct {
	Key   string
	value string
}

func GetDb() string {
	return "default"
}

/**
查找一个数据
*/
func MFindOne(query interface{}, dest interface{}) int64 {
	return getConn(GetDb()).Where(query).First(dest).RowsAffected
}

/**
查找数据列表
*/
func MFindList(query interface{}, dest interface{}, destList interface{}) int64 {
	return getConn(GetDb()).Where(query).Model(dest).Find(destList).RowsAffected
}

/**
插入数据
*/
func MAdd(dest interface{}) int64 {
	result := getConn(GetDb()).Create(dest)
	if result.Error != nil {
		_ = log4go.Error("mysql add error================", result.Error)
		return 0
	} else {
		return result.RowsAffected
	}
}

/**
更新数据
*/
func MSave(dest interface{}) int64 {
	result := getConn(GetDb()).Save(dest)
	if result.Error != nil {
		_ = log4go.Error("mysql save error================", result.Error)
		return 0
	} else {
		return result.RowsAffected
	}
}

/**
删除数据
*/
func MRemove(dest interface{}) int64 {
	result := getConn(GetDb()).Delete(dest)
	if result.Error != nil {
		_ = log4go.Error("mysql save error================", result.Error)
		return 0
	} else {
		return result.RowsAffected
	}
}
