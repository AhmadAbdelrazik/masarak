package httputils

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func ReadIDParam(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
}

func ReadString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

func ReadCSV(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)

	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}
