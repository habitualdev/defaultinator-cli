package utils

import (
	"errors"
	"github.com/mitchellh/go-homedir"
	"os"
	"strings"
)

func SaveCreds(apiKey string) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	file, err := os.Create(home + "/.defaultinator.ini")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("apiKey:" + apiKey)
	if err != nil {
		return err
	}
	return nil
}

func GetCreds() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	file, err := os.ReadFile(home + "/.defaultinator.ini")
	if err != nil {
		return "", err
	}
	apiKey := strings.Split(string(file), ":")

	if len(apiKey) < 2 {
		return "", errors.New("No API Key found")
	}

	return apiKey[1], nil
}
