package domain

import (
	"context"
	"github.com/lmxia/highway/pkg/errors"
	"github.com/lmxia/highway/pkg/util/structure"
	"gorm.io/gorm"

	"github.com/google/wire"
	"github.com/lmxia/highway/internal/app/dao/util"
	"github.com/lmxia/highway/internal/app/schema"
)

var DomainSet = wire.NewSet(wire.Struct(new(DomainRepo), "*"))

type DomainRepo struct {
	DB *gorm.DB
}

type SchemaDomain schema.Domain

func (a SchemaDomain) ToDomain() *Domain {
	item := new(Domain)
	structure.Copy(a, item)
	return item
}

func (a *DomainRepo) getQueryOption(opts ...schema.DomainQueryOptions) schema.DomainQueryOptions {
	var opt schema.DomainQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *DomainRepo) Query(ctx context.Context, params schema.DomainQueryParam, opts ...schema.DomainQueryOptions) (*schema.DomainQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetDomainDB(ctx, a.DB)
	if v := params.IDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}
	if v := params.Name; v != "" {
		db = db.Where("name=?", v)
	}
	if v := params.Status; v != 0 {
		db = db.Where("status=?", v)
	}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("name LIKE ?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Domains
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.DomainQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaRoles(),
	}

	return qr, nil
}

func (a *DomainRepo) Get(ctx context.Context, id uint64, opts ...schema.MenuQueryOptions) (*schema.Domain, error) {
	var item Domain
	ok, err := util.FindOne(ctx, GetDomainDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaDomain(), nil
}

func (a *DomainRepo) Create(ctx context.Context, item schema.Domain) error {
	eitem := SchemaDomain(item).ToDomain()
	result := GetDomainDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *DomainRepo) Update(ctx context.Context, id uint64, item schema.Domain) error {
	eitem := SchemaDomain(item).ToDomain()
	result := GetDomainDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *DomainRepo) Delete(ctx context.Context, id uint64) error {
	result := GetDomainDB(ctx, a.DB).Where("id=?", id).Delete(Domain{})
	return errors.WithStack(result.Error)
}

func (a *DomainRepo) UpdateStatus(ctx context.Context, id uint64, status int) error {
	result := GetDomainDB(ctx, a.DB).Where("id=?", id).Update("status", status)
	return errors.WithStack(result.Error)
}
