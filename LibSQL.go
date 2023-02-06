/*
Author: Pranav KV
Mail: pranavkvnambiar@gmail.com
*/
package golib_v1

import (
	"context"
	"fmt"
	"log"
	"time"

	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

var OrmDB *gorm.DB
var SqlDB *sql.DB
var MongoDB *mongo.Database

func dsn() string {

	username := GetString("database.username")
	pwd := GetString("database.password")
	host := GetString("database.host")
	port := GetString("database.port")
	database := GetString("database.name")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, pwd, host, port, database)

}

func postGresDsn() string {

	username := GetString("database.username")
	pwd := GetString("database.password")
	port := GetString("database.port")
	host := GetString("database.host")
	database := GetString("database.name")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, pwd, database, port)
}

func mongoDbDsn() string {

	//mongodb://localhost:27017
	port := GetString("mongodb.port")
	host := GetString("mongodb.host")
	return fmt.Sprintf("mongodb://%s:%s", host, port)
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

func InitPostGreSql() {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: postGresDsn()}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		Logger.Error("Unable to open database connection: ", err)
	}

	postgresdb, _ := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	postgresdb.SetMaxIdleConns(GetInt("database.max.idle"))
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	postgresdb.SetMaxOpenConns(GetInt("database.max.open"))
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	postgresdb.SetConnMaxLifetime(time.Minute * 30)

	db.Use(prometheus.New(prometheus.Config{
		DBName:          GetString("database.name"),  // use `DBName` as metrics label
		RefreshInterval: 15,                          // Refresh metrics interval (default 15 seconds)
		PushAddr:        "prometheus pusher address", // push metrics if `PushAddr` configured
		// StartServer:     true,                        // start http server to expose metrics
		HTTPServerPort: 8080, // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.Postgres{
				Prefix:        "gorm_status_",
				Interval:      100,
				VariableNames: []string{"Threads_running"},
			},
		}, // user defined metrics
	}))

	OrmDB = db
	//PostgresDB = postgresdb

	Logger.Info("postgres database connection success!!")
}

func InitMongoDB() {

	clientOptions := options.Client().ApplyURI(mongoDbDsn())

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	MongoDB = client.Database(GetString("mongodb.name"))

}
