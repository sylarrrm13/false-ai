package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}

	tTime := time.Time(*t)
	if tTime.IsZero() {
		return []byte("null"), nil

	}
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (lt LocalTime) Value() (driver.Value, error) {
	if time.Time(lt).IsZero() {
		return nil, nil
	}
	// 将 LocalTime 转换为 "2006-01-02 15:04:05" 格式的字符串
	return time.Time(lt).Format("2006-01-02 15:04:05"), nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	// 使用 time.Parse 解析时间字符串
	formats := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		// 添加其他格式字符串...
	}
	// 遍历所有的格式字符串，尝试解析输入时间字符串
	for _, layout := range formats {
		tmp, err := time.Parse(layout, str)
		if err == nil {
			*t = LocalTime(tmp)
			return nil
		}
	}
	// 如果所有的格式字符串都不能解析输入时间字符串，返回一个错误
	return fmt.Errorf("invalid time format: %s", str)
}
