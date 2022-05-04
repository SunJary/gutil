package gerrors

import "github.com/go-sql-driver/mysql"

// refer https://github.com/go-gorm/gorm/issues/5144

var (
	// ErrDuplicateEntryCode 命中唯一索引
	ErrDuplicateEntryCode = 1062
)

// MysqlErrCode 根据mysql错误信息返回错误代码
func MysqlErrCode(err error) int {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return 0
	}
	return int(mysqlErr.Number)
}

// todo 对同一个错误多次判断时，会有重复计算 MysqlErrCode ，可以优化
func IsDuplicateEntryErr(err error) bool {
	return MysqlErrCode(err) == ErrDuplicateEntryCode
}
