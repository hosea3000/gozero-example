package model

import (
	"database/sql/driver"
	"encoding/json"
)

// JSON 自定义JSON类型
type JSON []string

func (j *JSON) Scan(value interface{}) error {
	// 实现数据库扫描逻辑
	return json.Unmarshal(value.([]byte), j)
}

func (j *JSON) Value() (driver.Value, error) {
	// 实现数据库存储逻辑
	return json.Marshal(j)
}
