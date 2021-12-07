package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type mongoConfs struct {
	db_uri string `yaml:"uri"`
}

func ReadMongoConfigs() string {
	var confs mongoConfs

	file, err := os.Open("configs/mongo.yaml")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	byteConfigData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(byteConfigData, &confs)
	if err != nil {
		log.Fatal(err)
	}

	return confs.db_uri
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

	conn_str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", confs.Host, confs.Port, confs.User, "vvv15089", confs.DBName, confs.SSLMode)

	return conn_str
}
