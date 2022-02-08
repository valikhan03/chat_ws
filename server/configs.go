package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"github.com/joho/godotenv"
)



func ReadMongoConfigs() map[string]string {
	godotenv.Load("mongo.env")
	uri := os.Getenv("MONGO_URI")
	db := os.Getenv("MONGO_DB")

	confs := map[string]string{"URI":uri, "DB":db}

	return confs
}

type PostgresConfigs struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	User    string `yaml:"user"`
	DBName  string `yaml:"dbname"`
	SSLMode string `yaml:"sslmode"`
}

func ReadPostgresConfigs() string {
	var confs PostgresConfigs

	file, err := os.Open("configs/postgres.yaml")
	if err != nil {
		log.Fatal(err)
	}

	byteConfigData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(byteConfigData, &confs)
	if err != nil {
		log.Fatal(err)
	}

	godotenv.Load("postgres.env")
	password := os.Getenv("POSTGRES_PASSWORD")

	conn_str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", confs.Host, confs.Port, confs.User,  confs.DBName, confs.SSLMode, password)

	return conn_str
}
