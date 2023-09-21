package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/rs/xid"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Mixin of the User.
func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
		//手动显示设置外键
		field.String("user_id").
			GoType(xid.ID{}).
			Optional(),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Field("user_id"). //手动显示设置外键
			Ref("pets").
			Unique(),
	}
}

func (Pet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"), //手动显示设置外键索引
	}
}
