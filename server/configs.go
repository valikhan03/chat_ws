package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
)



func ReadMongoConfigs() string {
	godotenv.Load("mongo.env")
	uri := os.Getenv("MONGO_URI")
	return uri
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
