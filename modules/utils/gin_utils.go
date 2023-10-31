package utils

import (
	"bill/modules/constant"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// ExtractPaginationParams 从gin.context中取出分页查询基本参数
func ExtractPaginationParams(c *gin.Context) (int, int) {
	page := GetPageIndex(c)
	limit := GetPageLimit(c)
	return page, limit
}

// 获取页码
func GetPageIndex(c *gin.Context) int {
	str := c.Query("page")
	v, _ := strconv.ParseInt(str, 10, 0)
	p := int(v)
	if p <= 0 {
		p = 1
	}
	return int(v)
}

func GetPageLimit(c *gin.Context) int {
	str := c.Query("limit")
	v, _ := strconv.ParseInt(str, 10, 0)
	limit := int(v)
	if limit <= 0 {
		limit = 20
	}
	return limit
}

func GetUrlId(c *gin.Context) int {
	str := c.Param("id")
	v, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0
	}
	id := int(v)
	return id
}

func GetIntParam(c *gin.Context, key string) int {
	str := c.Param(key)
	v, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0
	}
	p := int(v)
	return p
}

// 获取排序信息
func GetPageSort(c *gin.Context) string {
	sort := c.DefaultQuery("sort", "")
	return sort
}

// 获取搜索关键词信息
func GetPageKey(c *gin.Context) string {
	key := c.DefaultQuery("key", "")
	return key
}

// 查询字段过滤
func ParseConds(c *gin.Context, fields []string) map[string]string {
	conds := make(map[string]string)

	for _, field := range fields {
		if value := c.Query(field); value != "" {
			conds[field] = value
		}
	}

	return conds
}

func SendErr(c *gin.Context, err error, data ...interface{}) {
	c.Abort()

	if len(data) > 0 && data[0] != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": constant.CodeError,
			"msg":  err.Error(),
			"data": data[0],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": constant.CodeError,
			"msg":  err.Error(),
		})
	}
}

func SendErrWithCode(c *gin.Context, code constant.ResponseCode, err error) {
	c.Abort()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err.Error(),
	})
}

func SendHttpErr(c *gin.Context, httpCode int, err error, data ...interface{}) {
	c.Abort()
	if len(data) > 0 && data[0] != nil {
		c.JSON(httpCode, gin.H{
			"msg":  err.Error(),
			"data": data[0],
		})
	} else {
		c.JSON(httpCode, gin.H{
			"msg": err.Error(),
		})
	}
}

// 返回一个带字符串模板的错误,后面跟参数
func SendTmplErr(c *gin.Context, err error, args ...interface{}) {
	c.Abort()
	//判断错误是否有对应的code
	c.JSON(http.StatusOK, gin.H{
		"code": constant.CodeError,
		"msg":  fmt.Sprintf(err.Error(), args),
	})
}

func SendSucc(c *gin.Context, data ...interface{}) {
	if len(data) > 0 && data[0] != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": constant.CodeSuccess,
			"data": data[0],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": constant.CodeSuccess,
			"data": "",
		})
	}
}

// 解决 xxx/:id 与xxx/create 这类路由冲突的问题,只需传入路由组和 xxx/create,以及
func MatchRouter(g *gin.RouterGroup, requestURI, relativePath string) bool {
	return strings.HasPrefix(requestURI, g.BasePath()+relativePath)
}
