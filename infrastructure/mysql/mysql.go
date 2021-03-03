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

/**
查找一个数据
*/
func MFindOne(db string, dest interface{}, query interface{}) {
	getConn(db).Where(query).First(dest)
}

/**
查找数据列表
*/
func MFind(db string, dest interface{}, query interface{}) {
	getConn(db).Where(query).Find(dest)
}

/**
插入数据
*/
func MAdd(db string, dest interface{}) int64 {
	result := getConn(db).Create(dest)
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
func MSave(db string, dest interface{}) int64 {
	result := getConn(db).Save(dest)
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
func MRemove(db string, dest interface{}) int64 {
	result := getConn(db).Delete(dest)
	if result.Error != nil {
		_ = log4go.Error("mysql save error================", result.Error)
		return 0
	} else {
		return result.RowsAffected
	}
}
