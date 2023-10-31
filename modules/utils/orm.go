package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"xorm.io/xorm"
)

type QueryCondition struct {
	Condition string
	Value     interface{}
}

//查询字段过滤
func MkConds(c *gin.Context, fields map[string]string) map[string]QueryCondition {
	conds := make(map[string]QueryCondition)

	fieldMap := make(map[string]string)
	for key, _ := range fields {
		fieldMap[key] = ""
	}
	err := c.ShouldBind(&fieldMap)
	if err != nil {
		return conds
	}
	for key, val := range fieldMap {
		if val != "" {
			conds[key] = QueryCondition{
				Condition: fields[key],
				Value:     val,
			}
		}
	}
	return conds
}

//TODO MkWhere 生成where查询条件  未完待续...
func MkWhere(session *xorm.Session, conds map[string]QueryCondition) *xorm.Session {
	for key, val := range conds {
		switch strings.ToLower(val.Condition) {
		case "eq":
			session.Where(key+" = ?", val.Value)
		case "neq":
			session.Where(key+" <> ?", val.Value)
		case "like":
			session.Where(key+" like ?", "%"+Strval(val.Value)+"%")
		case "in":
			session.Where(key+" in ?", val.Value)
		case "not in":
			session.Where(key+" not in ?", val.Value)
		case "gt":
			session.Where(key+" > ?", val.Value)
		case "egt":
			session.Where(key+" >= ?", val.Value)
		case "lt":
			session.Where(key+" < ?", val.Value)
		case "elt":
			session.Where(key+" <= ?", val.Value)
		default:
			continue
		}

	}
	return session
}

func KeyExists(key string, fields map[string]string) bool {
	for k, _ := range fields {
		if k == key {
			return true
		}
	}
	return false
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
