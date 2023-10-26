package utils

import (
	"errors"
	"net/http"
)

func ValidateRequest(r *http.Request) (configType string, name string, err error) {
	configType = r.URL.Query().Get("type")
	name = r.URL.Query().Get("name")

	if configType == "" || name == "" {
		return "", "", errors.New("missing config type or name in the request")
	}

	return configType, name, nil
}
