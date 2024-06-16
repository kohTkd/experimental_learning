package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Accounts holds the schema definition for the Accounts entity.
type Accounts struct {
	ent.Schema
}

// Fields of the Accounts.
func (Accounts) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password"),
		field.Time("created_at"),
		field.Time("update_at"),
	}
}

// Edges of the Accounts.
func (Accounts) Edges() []ent.Edge {
	return nil
}
