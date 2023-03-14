package schema

import (
	"time"
)

// Domain Domain对象
type Domain struct {
	ID        uint64    `json:"id,string"`                             // 唯一标识
	Name      string    `json:"name" binding:"required"`               // domain名称
	Memo      string    `json:"memo"`                                  // 备注
	Status    int       `json:"status" binding:"required,max=2,min=1"` // 状态(1:启用 2:禁用)
	CreatedAt time.Time `json:"created_at"`                            // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                            // 更新时间
}

// DomainQueryParam 查询条件
type DomainQueryParam struct {
	PaginationParam
	IDs        []uint64 `form:"-"`          // 唯一标识列表
	Name       string   `form:"-"`          // domain名称
	QueryValue string   `form:"queryValue"` // 模糊查询
	Status     int      `form:"status"`     // 状态(1:启用 2:禁用)
}

// DomainQueryOptions 查询可选参数项
type DomainQueryOptions struct {
	OrderFields  []*OrderField // 排序字段
	SelectFields []string      // 查询字段
}

// DomainQueryResult 查询结果
type DomainQueryResult struct {
	Data       Domains
	PageResult *PaginationResult
}

// Domains Domain对象列表
type Domains []*Domain

// ToNames 获取domain名称列表
func (a Domains) ToNames() []string {
	names := make([]string, len(a))
	for i, item := range a {
		names[i] = item.Name
	}
	return names
}

// ToMap 转换为键值存储
func (a Domains) ToMap() map[uint64]*Domain {
	m := make(map[uint64]*Domain)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}
