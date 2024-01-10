package copier

import (
	"GinTemplate/common/errs"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func Copy(toValue interface{}, fromValue interface{}) *errs.BError {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		zap.L().Error("copier error")
		return errs.CopierError
	}
	return nil
}
