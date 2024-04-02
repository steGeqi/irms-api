package project_common

type BusinessCode int
type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}
type Data struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
}

func (d *Data) BackData(code BusinessCode, msg string) *Data {
	d.Code = code
	d.Msg = msg
	return d
}

func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Msg = "Success"
	r.Data = data
	return r
}

func (r *Result) Fail(code BusinessCode, msg string, data any) *Result {
	r.Msg = msg
	r.Code = code
	r.Data = data
	return r
}
