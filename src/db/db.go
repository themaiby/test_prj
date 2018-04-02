package db

import (
	_ "github.com/denisenkom/go-mssqldb" // driver only
	"github.com/jmoiron/sqlx"
	log "github.com/kataras/golog"
	"github.com/spf13/viper"
)

var DB *sqlx.DB
var err error

var connectionString string

func Connect() {

	var host = viper.GetString("database.host")
	var port = viper.GetString("database.port")
	var instance = viper.GetString("database.instance")
	var user = viper.GetString("database.user")
	var password = viper.GetString("database.password")
	var db = viper.GetString("database.db")
	var connectionTimeout = "30"

	var portOrInstance string
	if instance != "" {
		// get setting from settings.json
		portOrInstance = "\\" + instance
	} else {
		portOrInstance = ":" + port
	}

	// make connection string by credentials
	connectionString = "odbc:server=" + host + portOrInstance + ";user id=" + user + ";password=" + password + " ;database=" + db + ";connection timeout=" + connectionTimeout

	// open DB connection
	DB, err = sqlx.Open("mssql", connectionString)

	// check errors. Exit program if no connect
	if err != nil {
		log.Error("[MSSQL] ", err)
		return
	} else {
		dberr := DB.Ping()
		if dberr != nil {
			log.Fatal("Error while DB connecting: ", dberr)
		} else {
			log.Info("[MSSQL] Connection success")
		}
	}
}

func IsConnected() bool {
	err := DB.Ping()
	if err != nil {
		log.Debug("DB NOT connected")
		return false
	}
	log.Debug("DB connected")
	return true
}

func Close() {
	DB.Close()
	log.Info("[MSSQL] DB connection closed")
}

func Reconnect() {
	DB, err = sqlx.Open("mssql", connectionString)
	if err != nil {
		log.Error("[MSSQL] Reconnect Error: ", err)
		return
	} else {
		log.Info("[MSSQL] Reconnection success")
	}
}
