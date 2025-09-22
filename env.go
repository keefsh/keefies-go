package keefies

import (
	"fmt"
	"os"
)

type EnvError struct {
	arg     string
	message string
}

func (e *EnvError) Error() string {
	return fmt.Sprintf("%s: %s", e.arg, e.message)
}

func GetEnv(key string) (string, bool) {
	value, exists := os.LookupEnv(key)

	if !exists {
		return "", false
	}

	return value, true
}

func MustGetEnv(key string) string {
	value, ok := GetEnv(key)

	if !ok {
		panic(EnvError{arg: key, message: fmt.Sprintf("Expected to find %s", key)})
	}

	return value
}

type As[T any] = func(value string) (T, error)

func getEnvAs[T any](key string, mapper As[T]) (T, error) {
	value, ok := GetEnv(key)

	if !ok {
		var zero T
		return zero, &EnvError{arg: key, message: fmt.Sprintf("Expected to find %s", key)}
	}

	mapped, err := mapper(value)

	if err != nil {
		var zero T
		return zero, err
	}

	return mapped, nil
}
