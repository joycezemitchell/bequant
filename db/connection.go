package DB

import (
	"allyapps.com/bequant/models"
	"os"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"encoding/json"	
)

var Dbs *sql.DB

func Init()  {

	configuration := getConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configuration.Host, configuration.Port, configuration.User, configuration.Password, configuration.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Successfully connected!")
	

	Dbs = db
	
}

func getConfig() model.Configuration{
	jsonFile, _ := os.Open(getConfigurationPath() + "/"   + "config.json")
	defer jsonFile.Close()
	var configuration model.Configuration
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &configuration)
	return configuration
}

func getConfigurationPath() string{
	wd, _ := os.Getwd()
	cofigurationPath := fmt.Sprintf("%s", wd)

	return cofigurationPath
}
