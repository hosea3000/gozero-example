package model

import (
	"time"
)

type Credential struct {
	ID

	// 认证信息
	ApiKey      string `gorm:"type:varchar(64);uniqueIndex;not null;comment:API Key" json:"-"`
	Secret      string `gorm:"type:varchar(255);not null;comment:加密后的Secret" json:"-"`
	DisplayName string `gorm:"type:varchar(255);not null;comment:凭证显示名称" json:"display_name"`

	// 权限控制
	Scopes      JSON `gorm:"type:json;not null;comment:权限范围" json:"scopes"`
	RateLimit   int  `gorm:"default:1000;comment:QPS限制" json:"rate_limit"`
	IPWhitelist JSON `gorm:"type:json;comment:IP白名单" json:"ip_whitelist"`

	// 状态管理
	IsActive  bool       `gorm:"default:true;comment:是否启用" json:"is_active"`
	ExpiresAt *time.Time `gorm:"index;comment:过期时间" json:"expires_at"`

	// 审计字段
	LastUsedAt *time.Time `gorm:"comment:最后使用时间" json:"last_used_at"`

	ModelTime
	ControlBy
}

func (Credential) TableName() string {
	return "credential"
}
