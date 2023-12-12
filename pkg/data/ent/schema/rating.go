// Rating holds the schema definition for the Rating entity.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Rating struct {
	ent.Schema
}

// Fields of the Rating.
func (Rating) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("user_id"),
		field.Int("topic_id"),
		field.Float("rating"),
		field.String("comment"),
	}
}

// Edges of the Rating.
func (Rating) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ratedBy", User.Type).Ref("ratings").Unique(),
		edge.From("underTopicOf", Topic.Type).Ref("ratings").Unique(),
	}
}
