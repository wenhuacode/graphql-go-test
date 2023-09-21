package schema

import (
	"context"
	gen "ent-orm-test/ent"
	"ent-orm-test/ent/hook"
	"ent-orm-test/ent/intercept"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/rs/xid"
)

// BaseMixin to be shared will all different schemas.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the User.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(xid.ID{}).
			DefaultFunc(xid.New),
		field.Time("created_at").
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Default(func() time.Time { return time.Now() }).
			Comment("创建时间").
			StructTag(`gqlgen:"created_at"`),
		field.String("created_by").
			Optional().
			Comment("创建人").
			StructTag(`gqlgen:"created_by"`),
		field.Time("updated_at").
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Default(func() time.Time { return time.Now() }).
			Comment("更新时间").
			StructTag(`gqlgen:"updated_at"`),
		field.String("updated_by").
			Optional().
			Comment("更新人").
			StructTag(`gqlgen:"updated_by"`),
		field.Bool("is_deleted").
			Optional().
			Comment("是否删除").
			StructTag(`gqlgen:"is_deleted"`),
		field.Time("deleted_at").
			Optional().
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Comment("删除时间").
			StructTag(`gqlgen:"deleted_at"`),
		field.String("deleted_by").
			Optional().
			Comment("删除人").
			StructTag(`gqlgen:"deleted_by"`),
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

// Hooks of the SoftDeleteMixin.
func (d BaseMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(softDeleteHook(d), ent.OpDeleteOne|ent.OpDelete),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d BaseMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[6].Descriptor().Name),
	)
}
