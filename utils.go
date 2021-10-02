package mixed

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// copyContext create a new echo.Context with the same Request, Path, Params and Handler
// but with a Response that contain an in-memory ResponseWriter
func copyContext(c echo.Context) echo.Context {
	cc := c.Echo().AcquireContext()
	cc.SetRequest(c.Request())
	cc.SetPath(c.Path())
	cc.SetParamNames(c.ParamNames()...)
	cc.SetParamValues(c.ParamValues()...)
	cc.SetHandler(c.Handler())

	rw := tempResponseWriter{
		header:  make(http.Header),
		Content: []byte{},
	}
	resp := echo.NewResponse(&rw, c.Echo())
	cc.SetResponse(resp)
	return cc
}

// copyResponse copy c2 headers and content into c1
func copyResponse(c1, c2 echo.Context) {
	for k, v := range c2.Response().Header() {
		for _, vv := range v {
			c1.Response().Header().Set(k, vv)
		}
	}

	if c2.Response().Status > 0 {
		c1.Response().WriteHeader(c2.Response().Status)
	}

	c1.Response().Write(c2.Response().Writer.(*tempResponseWriter).Content)
}