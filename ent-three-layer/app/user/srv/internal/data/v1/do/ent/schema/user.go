package schema

import (
	"ent-three-layer/app/user/srv/internal/data/v1/do/ent/base"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		base.BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("mobile").
			Unique().
			Comment("手机号"),
		field.String("password").
			NotEmpty().
			Comment("密码"),
		field.String("nickname").
			Optional().
			Comment("昵称"),
		field.Time("birthday").
			Optional().
			Comment("生日"),
		field.String("gender").
			Default("male").
			Comment("性别"),
		field.Int("Role").
			Default(1).
			Comment("角色"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
