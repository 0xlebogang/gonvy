package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Verification holds the schema definition for the Verification entity.
type Verification struct {
	ent.Schema
}

func (Verification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}

// Fields of the Verification.
func (Verification) Fields() []ent.Field {
	return []ent.Field{
		field.String("identifier").
			NotEmpty(),
		field.String("value").
			NotEmpty(),
		field.Time("expires_at"),
	}
}

// Indexes of the Verification.
func (Verification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("identifier"),
	}
}
