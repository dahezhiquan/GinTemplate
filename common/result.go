package common

import "GinTemplate/common/errs"

type Result struct {
	Code errs.ErrorCode `json:"code"`
	Msg  string         `json:"msg"`
	Data any            `json:"data"`
}

func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}

func (r *Result) Fail(bErr *errs.BError) *Result {
	r.Code = bErr.Code
	r.Msg = bErr.Msg
	return r
}
