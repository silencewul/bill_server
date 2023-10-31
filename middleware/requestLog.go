package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
)

func RequestLoggerMiddleware(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := ioutil.ReadAll(tee)
	c.Request.Body = ioutil.NopCloser(&buf)
	fmt.Printf("请求Body:%s\n", string(body))
	fmt.Printf("请求Header:%s\n", c.Request.Header)
	c.Next()
}
