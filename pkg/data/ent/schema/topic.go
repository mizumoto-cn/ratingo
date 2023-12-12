// Topic holds the schema definition for the Topic entity.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("topic_name"),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("analysis", Analysis.Type).Unique(),
		edge.To("ratings", Rating.Type),
	}
}
