package utils4g

var (
	dbInstance = &dbUtil{
		Mongo: newDbMgoUtil(),
		Mysql: newDbMysqlUtil(),
	}
)

func newDbUtil() *dbUtil {
	return dbInstance
}
