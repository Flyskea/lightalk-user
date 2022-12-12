package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type UserInfo struct {
	ent.Schema
}

// Fields of the User.
func (UserInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int64("user_id").
			Unique().
			Comment("用户ID"),
		field.Int8("sex").
			Default(0).
			Comment("用户性别"),
		field.Int8("age").
			Optional().
			Comment("用户年龄"),
		field.String("birthday").
			SchemaType(map[string]string{
				dialect.MySQL: "date",
			}).
			Optional().
			Nillable().
			Comment("用户生日"),
		field.String("motto").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}).
			Optional().
			Comment("用户座右铭"),
	}
}

// Edges of the User.
func (UserInfo) Edges() []ent.Edge {
	return nil
}

func (UserInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (UserInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_info"},
	}
}
