package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").
			Unique().
			Optional().
			Nillable().
			MinLen(1).
			MaxLen(255),
		field.Enum("provider").
			Values("credential").
			Default("credential"),
		field.String("access_token").
			Optional().
			Nillable(),
		field.String("refresh_token").
			Optional().
			Nillable(),
		field.Time("access_token_expires_at").
			Optional().
			Nillable(),
		field.Time("refresh_token_expires_at").
			Optional().
			Nillable(),
		field.String("scope").
			Optional().
			Nillable(),
		field.String("id_token").
			Optional().
			Nillable(),
		field.Text("password").
			Optional().
			Nillable(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("accounts").
			Unique().
			Required(),
	}
}
