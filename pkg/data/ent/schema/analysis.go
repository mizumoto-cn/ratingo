// Analysis holds the schema definition for the Analysis entity.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Analysis struct {
	ent.Schema
}

// Fields of the Analysis.
func (Analysis) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("topic_id"),
		field.Float("avg_rating"),
	}
}

// Edges of the Analysis.
func (Analysis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("underTopicOf", Topic.Type).Ref("analysis").Unique(),
	}
}
