package helper

import "github.com/joho/godotenv"

func GetEnv(envPath string) (map[string]string, error) {
	env, err := godotenv.Read(envPath)
	if err != nil {
		return env, err
	}

	return env, nil
}
