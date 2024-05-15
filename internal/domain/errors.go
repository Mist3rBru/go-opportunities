package domain

import "fmt"

func RequiredParamError(param string, typ string) error {
	return fmt.Errorf("param `%s(type:%s)` is required", param, typ)
}

func InvalidParamError(param string) error {
	return fmt.Errorf("param `%s` is invalid", param)
}

func NotFoundError(item string) error {
	return fmt.Errorf("%s not found", item)
}
