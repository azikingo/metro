package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("tg_id").Unique(),
		field.String("username").Nillable(),
		field.String("name").Nillable(),
		field.String("surname").Nillable(),
		field.String("email").Nillable(),
		field.String("phone").Nillable(),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
