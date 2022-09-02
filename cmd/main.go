package main

import (
	"log"
	"os"

	todo "github.com/Nikita-Kuzhl/go-rest-api"
	"github.com/Nikita-Kuzhl/go-rest-api/package/handler"
	"github.com/Nikita-Kuzhl/go-rest-api/package/repository"
	"github.com/Nikita-Kuzhl/go-rest-api/package/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main(){
	if err:= initConfig(); err!=nil{
		log.Fatalf("error inittial config - %s",err.Error())
	}
	if err:= godotenv.Load();err!=nil{
		log.Fatalf("error env - %s",err.Error())
	}
	db,err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err!=nil {
		log.Fatalf("failed to init db - %s",err.Error())
	}

	repos := repository.NewRepository(db)
	services:=service.NewService(*repos)

	handlers:= handler.NewHandler(services)
	srv := new(todo.Server);
	if err:= srv.Run(viper.GetString("port"),handlers.InitRouter()); err != nil {
		log.Fatalf("error server - %s",err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}