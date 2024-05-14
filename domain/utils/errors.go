package domain_utils

import "fmt"

func RequiredParamError(param string, typ string) error {
	return fmt.Errorf("param: %s(type:%s) is required", param, typ)
}
