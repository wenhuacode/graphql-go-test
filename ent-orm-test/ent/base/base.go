package base

import (
	"context"
	"ent-orm-test/ent/hook"
	"ent-orm-test/ent/intercept"
	"fmt"
	"github.com/rs/xid"
	"time"

	gen "ent-orm-test/ent"
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
		field.String("id").
			GoType(xid.ID{}).
			DefaultFunc(xid.New),
		field.Time("created_at").
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Default(time.Now).
			Comment("创建时间"),
		field.String("created_by").
			Optional().
			Comment("创建人"),
		field.Time("updated_at").
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Optional().
			UpdateDefault(time.Now).
			Comment("更新时间"),
		field.String("updated_by").
			Optional().
			Comment("更新人"),
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
			userID := "test12"
			mx, ok := m.(interface {
				SetOp(ent.Op)
				Client() *gen.Client
				SetDeletedAt(time.Time)
				SetDeletedBy(string)
				WhereP(...func(*sql.Selector))
			})
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			d.P(mx)
			mx.SetOp(ent.OpUpdate)
			mx.SetDeletedAt(time.Now())
			mx.SetDeletedBy(userID)
			return mx.Client().Mutate(ctx, m)
		})
	}
}

// A AuditHook is an example for audit-log hook. 安全删除
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger wraps the methods that are shared between all mutations of
	// schemas that embed the AuditLog mixin. The variable "exists" is true, if
	// the field already exists in the mutation (e.g. was set by a different hook).
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedBy(string)
		CreatedBy() (id string, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedBy(string)
		UpdatedBy() (id string, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		//获取用户id
		userID := "test12"

		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())
			if _, exists := ml.CreatedBy(); !exists {
				ml.SetCreatedBy(userID)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())
			if _, exists := ml.UpdatedBy(); !exists {
				ml.SetUpdatedBy(userID)
			}
		}
		return next.Mutate(ctx, m)
	})
}

// Hooks of the SoftDeleteMixin.
func (d BaseMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
		hook.On(softDeleteHook(d), ent.OpDeleteOne|ent.OpDelete),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d BaseMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[6].Descriptor().Name),
	)
}
