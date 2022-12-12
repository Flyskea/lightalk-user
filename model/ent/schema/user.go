package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int64("user_id").
			Unique().
			Comment("用户ID"),
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(20)",
			}).
			Comment("用户名"),
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}).
			Unique().
			Comment("用户邮箱"),
		field.String("password").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}).
			Comment("用户密码"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}
