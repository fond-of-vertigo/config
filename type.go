package config

import (
	"errors"
	"strings"
)

type EnvVariable string

func (e *EnvVariable) UnmarshalJSON(data []byte) error {
	value := string(data)
	if !strings.HasPrefix(strings.ToUpper(value), "\"{ENV:") || !strings.HasSuffix(value, "}\"") {
		return errors.New("unknown EnvVariable tag: " + value)
	}
	variableName := value[6 : len(value)-2]
	*e = EnvVariable(MustGetEnv(variableName))
	return nil
}

func (e *EnvVariable) String() string {
	if e == nil {
		return ""
	}
	return string(*e)
}
