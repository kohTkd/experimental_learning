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
		field.String("email"),
		field.Int("account_type"),
		field.String("password"),
	}
}

// Edges of the Accounts.
func (Accounts) Edges() []ent.Edge {
	return nil
}
