package blfgo

var interceptorList []func(c *Controller) bool

func Intercept(c *Controller) bool {

	if len(interceptorList) == 0 {
		return true
	}

	for _, f := range interceptorList {
		if !f(c) {
			return false
		}
	}

	return true
}

func InsertInterceptor(f func(c *Controller) bool) {
	if len(interceptorList) == 0 {
		interceptorList = make([]func(c *Controller) bool, 0, 10)
	}
	interceptorList = append(interceptorList, f)
}
