package application

import (
	"errors"
	"fmt"

	golog "github.com/MarceloRezin/go-log/cmd/go-log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppContext struct {
	Log    *golog.Logger
	DB     *gorm.DB
	IsProd bool
}

func GetAppContext(env *Env) (AppContext, error) {
	appContext := AppContext{
		IsProd: env.IsProd,
	}

	var err error
	if env.IsProd {
		appContext.Log, err = golog.DefaultFile()
	} else {
		appContext.Log, err = golog.DefaultConsole()
	}

	if err != nil {
		return appContext, errors.New("it was not possible get logger")
	}

	appContext.Log.Info.Println("logger started")

	db, err := createDbConnection(env)
	if err != nil {
		return appContext, errors.New("it was not possible get db connection")
	}

	appContext.Log.Info.Println("database started")

	appContext.DB = db

	return appContext, nil
}

func createDbConnection(env *Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(144.22.223.201:1432)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DbUser, env.DbPass, env.DbName)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
