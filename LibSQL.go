/*
Author: Pranav KV
Mail: pranavkvnmabiar@gmail.com
*/
package golib_v1

import (
	"fmt"
	"time"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

var OrmDB *gorm.DB
var SqlDB *sql.DB

func dsn() string {
	username := GetString("database.username")
	pwd := GetString("database.password")
	host := GetString("database.host")
	port := GetString("database.port")
	database := GetString("database.name")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, pwd, host, port, database)

}

func InitMySQL() {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn()}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent),
		PrepareStmt:            true,
		SkipDefaultTransaction: true})
	if err != nil {
		Logger.Error("Unable to open database connection: ", err)
	}

	sqldb, _ := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqldb.SetMaxIdleConns(GetInt("database.max.idle"))
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqldb.SetMaxOpenConns(GetInt("database.max.open"))
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqldb.SetConnMaxLifetime(time.Minute * 30)

	db.Use(prometheus.New(prometheus.Config{
		DBName:          GetString("database.name"),  // use `DBName` as metrics label
		RefreshInterval: 15,                          // Refresh metrics interval (default 15 seconds)
		PushAddr:        "prometheus pusher address", // push metrics if `PushAddr` configured
		// StartServer:     true,                        // start http server to expose metrics
		HTTPServerPort: 8080, // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				Prefix:        "gorm_status_",
				Interval:      100,
				VariableNames: []string{"Threads_running"},
			},
		}, // user defined metrics
	}))

	OrmDB = db
	SqlDB = sqldb

	Logger.Info("database connection success!!")
}
