package utils

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// SafelyConvertToInt 尝试将各种类型安全地转换为 int
// 支持 int, int32, int64, float32, float64, string 类型的转换
// 对于浮点数，要求必须是整数值（如 1.0），不接受带小数部分的值
// 对于字符串，尝试解析为整数或浮点数（浮点数必须是整数值）
func SafelyConvertToInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case float32:
		if float32(int(v)) == v {
			return int(v), nil
		}
		return 0, NewParamError("number", "must be an integer value")
	case float64:
		if float64(int(v)) == v {
			return int(v), nil
		}
		return 0, NewParamError("number", "must be an integer value")
	case string:
		var intValue int
		if _, err := fmt.Sscanf(v, "%d", &intValue); err == nil {
			return intValue, nil
		}
		var floatValue float64
		if _, err := fmt.Sscanf(v, "%f", &floatValue); err == nil {
			if math.Floor(floatValue) == floatValue {
				return int(floatValue), nil
			}
		}
		return 0, NewParamError("number", "must be a valid integer")
	default:
		return 0, NewParamError("number", fmt.Sprintf("unsupported type: %v", reflect.TypeOf(value)))
	}
}

// SafelyGetString 从 map 中安全提取 string 类型的参数值
// 支持 string, float64, float32, int, int32, int64 类型
// 对于数值类型自动转换为字符串表示，避免 JSON 反序列化导致的类型断言 panic
func SafelyGetString(key string, args map[string]interface{}) (string, error) {
	value, exists := args[key]
	if !exists {
		return "", NewParamError(key, "parameter is required")
	}
	switch v := value.(type) {
	case string:
		return v, nil
	case float64:
		return strconv.Itoa(int(v)), nil
	case float32:
		return strconv.Itoa(int(v)), nil
	case int:
		return strconv.Itoa(v), nil
	case int32:
		return strconv.Itoa(int(v)), nil
	case int64:
		return strconv.Itoa(int(v)), nil
	default:
		return "", NewParamError(key, fmt.Sprintf("unsupported type: %T", value))
	}
}

// ConvertArgumentsToMap 将 any 类型的参数安全转换为 map[string]interface{}
func ConvertArgumentsToMap(arguments any) (map[string]interface{}, error) {
	if arguments == nil {
		return make(map[string]interface{}), nil
	}

	if argMap, ok := arguments.(map[string]interface{}); ok {
		return argMap, nil
	}

	return nil, NewParamError("arguments", "arguments must be a map[string]interface{}")
}
