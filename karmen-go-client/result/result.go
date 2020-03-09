package result

//Result - result struct
type Result struct {
	status bool //false = fail, true = pass
}

//Pass - passing result
func (r *Result) Pass() {
	r.status = true
}

//Fail - failing result
func (r *Result) Fail() {
	r.status = false
}

//GetResult -
func (r *Result) GetResult() bool {
	return r.status
}
