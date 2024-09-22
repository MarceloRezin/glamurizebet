package application

import (
	"errors"

	goenv "github.com/MarceloRezin/go-env/cmd/go-env"
)

type Env struct {
	DbUser string
	DbPass string
	DbName string
	IsProd bool
}

func LoadEnv() (Env, error) {
	envMap, err := goenv.Read(eNV_FILE)
	if err != nil {
		return Env{}, err
	}

	et := envMap[envType]
	if et == "" {
		return Env{}, errors.New("env don't setted")
	}

	return Env{
		DbUser: envMap[dB_USER],
		DbPass: envMap[dB_PASS],
		DbName: envMap[dB_NAME],
		IsProd: et == prod,
	}, nil
}
