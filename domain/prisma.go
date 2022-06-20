/**
domain 类似于 DDD 中领域服务, 这里主要存储领域模型、模型对应的方法、定义与 delivery、repositpry 层交互接口
*/
package domain

import (
	"context"
	"time"
)

type Prisma struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Schema     string    `json:"schema"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
	IsDel      int8      `json:"-"`
}

type PrismaUseCase interface {
	Create(ctx context.Context, p *Prisma) error
	Update(ctx context.Context, p *Prisma) error
	Fetch(ctx context.Context) ([]Prisma, error)
	Delete(ctx context.Context, id int64) error
}

type PirsmaRepository interface {
	Create(ctx context.Context, p *Prisma) error
	Update(ctx context.Context, p *Prisma) error
	Fetch(ctx context.Context) ([]Prisma, error)
	Delete(ctx context.Context, id int64) error
}
