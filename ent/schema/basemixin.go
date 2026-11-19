package schema

import (
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	ulid "github.com/oklog/ulid/v2"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(generateULID),
	}
}

func generateULID() string {
	return strings.ToLower(ulid.Make().String())
}
