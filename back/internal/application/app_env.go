package application

import goenv "github.com/MarceloRezin/go-env/cmd/go-env"

type Env struct {
	DbUser string
	DbPass string
	DbName string
}

func LoadEnv() (Env, error) {
	envMap, err := goenv.Read(eNV_FILE)
	if err != nil {
		return Env{}, err
	}

	return Env{
		DbUser: envMap[dB_USER],
		DbPass: envMap[dB_PASS],
		DbName: envMap[dB_NAME],
	}, nil
}
