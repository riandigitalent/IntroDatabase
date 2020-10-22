package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/riandigitalent/IntroDatabase/sql-generic/config"
	"github.com/riandigitalent/IntroDatabase/sql-generic/database"
	"github.com/spf13/viper"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := connect(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}
	//coba insert
	//database.InsertCustomer(database.Customer{
	//	FirstName:    "Rian",
	//	LastName:     "Yunandar",
	//	NpwpId: "id-1",
	//	Age:          10,
	//	CustomerType: "Sultan",
	//	Street:       "tiuhbalak",
	//	City:         "Waykanan",
	//	State:        "Indo",
	//	ZipCode:      "35122",
	//	PhoneNumber:  "0812a3232384",
	//},db)

	//coba tarik data constumer
	//database.GetCustomers(db)

	//coba hapus databy id
	//database.DeleteCustomer(1,db)

	//update umur
	database.UpdateCustomer(28, 2, db)
}

//fungsi tarik data yml
func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

//fungsi konek ke db mysql
func connect(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Config))
	if err != nil {
		return nil, err
	}

	log.Println("db successfully connected")
	return db, nil
}
