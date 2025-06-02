package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/yaml.v2"
)

// var pool *pgxpool.Pool
var DB *pgxpool.Pool

func failOnError(err error, msg string) { //Делаем более читаемую и компактную обработку ошибок.
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// функция подключения к базе данных
func Connect() error {
	const configPath = "config.yml"
	type Cfg struct {
		POSTGRES_HOST           string `yaml:"postgres_host"`
		POSTGRES_PORT           int    `yaml:"postgres_port"`
		POSTGRES_DB             string `yaml:"postgres_db"`
		POSTGRES_USER           string `yaml:"postgres_user"`
		POSTGRES_PASS           string `yaml:"postgres_pass"`
		POSTGRES_SSL            string `yaml:"postgres_ssl"`
		POSTGRES_POOL_MAX_CONNS int    `yaml:"postgres_pool_max_conns"`
	}
	var AppConfig *Cfg
	f, err := os.Open(configPath)
	failOnError(err, "Can't open config.\n")
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)
	failOnError(err, "Can't decode config.\n")

	postgres_host := AppConfig.POSTGRES_HOST
	postgres_port := AppConfig.POSTGRES_PORT
	postgres_db := AppConfig.POSTGRES_DB
	postgres_user := AppConfig.POSTGRES_USER
	postgres_pass := AppConfig.POSTGRES_PASS
	postgres_ssl := AppConfig.POSTGRES_SSL
	postgres_pool_max_conns := AppConfig.POSTGRES_POOL_MAX_CONNS

	//Инициализация БД
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&pool_max_conns=%d",
		postgres_user, postgres_pass, postgres_host, postgres_port, postgres_db, postgres_ssl, postgres_pool_max_conns)

	pool, err := pgxpool.New(context.Background(), dbURL)
	failOnError(err, "Unable to connection to database: %v.\n")
	defer pool.Close()
	log.Print("Connected to database!\n")

	DB = pool
	log.Println("Успешно подключились к базе данных")
	return nil
}
