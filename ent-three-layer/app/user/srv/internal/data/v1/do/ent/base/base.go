package base

import (
	"context"
	gen "ent-three-layer/app/user/srv/internal/data/v1/do/ent"
	"ent-three-layer/app/user/srv/internal/data/v1/do/ent/hook"
	"ent-three-layer/app/user/srv/internal/data/v1/do/ent/intercept"
	"fmt"
	"github.com/oklog/ulid/v2"
	"log"
	"os"
	prand "pgregory.net/rand"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin to be shared will all different schemas.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the User.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Time("created_at").
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Default(func() time.Time { return time.Now() }).
			Comment("创建时间"),
		field.String("created_by").
			Optional().
			Comment("创建人"),
		field.Time("updated_at").
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Default(func() time.Time { return time.Now() }).
			Comment("更新时间"),
		field.String("updated_by").
			Optional().
			Comment("更新人"),
		field.Bool("is_deleted").
			Optional().
			Comment("是否删除"),
		field.Time("deleted_at").
			Optional().
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Comment("删除时间"),
		field.String("deleted_by").
			Optional().
			Comment("删除人"),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the SoftDeleteMixin. 跳过软删除
func (d BaseMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}
			d.P(q)
			return nil
		}),
	}
}

func softDeleteHook(d BaseMixin) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			// Skip soft-delete, means delete the entity permanently.
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return next.Mutate(ctx, m)
			}
			mx, ok := m.(interface {
				SetOp(ent.Op)
				Client() *gen.Client
				SetIsDeleted(bool)
				SetDeletedAt(time.Time)
				WhereP(...func(*sql.Selector))
			})
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			d.P(mx)
			mx.SetOp(ent.OpUpdate)
			mx.SetIsDeleted(true)
			mx.SetDeletedAt(time.Now())
			return mx.Client().Mutate(ctx, m)
		})
	}
}

func IDHook() ent.Hook {
	// 获取当前机器ID，这里假设你从环境变量中获取机器ID
	machineIDStr := os.Getenv("MACHINE_ID")
	if machineIDStr == "" {
		log.Fatalf("MACHINE_ID环境变量未设置")
		return nil
	}

	entropy := ulid.Monotonic(prand.New(), 777)
	id, _ := ulid.New(ulid.Timestamp(time.Now()), entropy)

	type IDSetter interface {
		SetID(string)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(IDSetter)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}
			is.SetID(id.String())
			return next.Mutate(ctx, m)
		})
	}
}

// Hooks of the SoftDeleteMixin.
func (d BaseMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(softDeleteHook(d), ent.OpDeleteOne|ent.OpDelete),
		hook.On(IDHook(), ent.OpCreate),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d BaseMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[6].Descriptor().Name),
	)
}
