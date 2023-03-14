package domain

import (
	"context"
	"github.com/lmxia/highway/internal/app/schema"
	"github.com/lmxia/highway/pkg/util/structure"
	"gorm.io/gorm"

	"github.com/lmxia/highway/internal/app/dao/util"
)

func GetDomainDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Domain))
}

// Domain sample for caskin.Domain interface
type Domain struct {
	util.Model
	DeletedAt    gorm.DeletedAt `gorm:"column:delete_at;index" json:"-"`
	Name         string         `gorm:"column:name;unique"     json:"name,omitempty"`
	Status       int            `gorm:"index;default:0;"` // 状态(1:启用 2:禁用)
	Memo         *string        `gorm:"size:1024;"`       // 备注
	MaintainerID uint64         `gorm:"index;default:0;"` // 管理员用户内码
}

func (a Domain) ToSchemaDomain() *schema.Domain {
	item := new(schema.Domain)
	structure.Copy(a, item)
	return item
}

type Domains []*Domain

func (a Domains) ToSchemaDomains() []*schema.Domain {
	list := make([]*schema.Domain, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDomain()
	}
	return list
}
