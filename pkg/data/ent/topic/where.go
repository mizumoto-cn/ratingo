// Code generated by ent, DO NOT EDIT.

package topic

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TopicName applies equality check predicate on the "topic_name" field. It's identical to TopicNameEQ.
func TopicName(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopicName), v))
	})
}

// TopicNameEQ applies the EQ predicate on the "topic_name" field.
func TopicNameEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopicName), v))
	})
}

// TopicNameNEQ applies the NEQ predicate on the "topic_name" field.
func TopicNameNEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTopicName), v))
	})
}

// TopicNameIn applies the In predicate on the "topic_name" field.
func TopicNameIn(vs ...string) predicate.Topic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTopicName), v...))
	})
}

// TopicNameNotIn applies the NotIn predicate on the "topic_name" field.
func TopicNameNotIn(vs ...string) predicate.Topic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTopicName), v...))
	})
}

// TopicNameGT applies the GT predicate on the "topic_name" field.
func TopicNameGT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTopicName), v))
	})
}

// TopicNameGTE applies the GTE predicate on the "topic_name" field.
func TopicNameGTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTopicName), v))
	})
}

// TopicNameLT applies the LT predicate on the "topic_name" field.
func TopicNameLT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTopicName), v))
	})
}

// TopicNameLTE applies the LTE predicate on the "topic_name" field.
func TopicNameLTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTopicName), v))
	})
}

// TopicNameContains applies the Contains predicate on the "topic_name" field.
func TopicNameContains(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTopicName), v))
	})
}

// TopicNameHasPrefix applies the HasPrefix predicate on the "topic_name" field.
func TopicNameHasPrefix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTopicName), v))
	})
}

// TopicNameHasSuffix applies the HasSuffix predicate on the "topic_name" field.
func TopicNameHasSuffix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTopicName), v))
	})
}

// TopicNameEqualFold applies the EqualFold predicate on the "topic_name" field.
func TopicNameEqualFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTopicName), v))
	})
}

// TopicNameContainsFold applies the ContainsFold predicate on the "topic_name" field.
func TopicNameContainsFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTopicName), v))
	})
}

// HasAnalysis applies the HasEdge predicate on the "analysis" edge.
func HasAnalysis() predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnalysisTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AnalysisTable, AnalysisColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnalysisWith applies the HasEdge predicate on the "analysis" edge with a given conditions (other predicates).
func HasAnalysisWith(preds ...predicate.Analysis) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnalysisInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AnalysisTable, AnalysisColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRatings applies the HasEdge predicate on the "ratings" edge.
func HasRatings() predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RatingsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RatingsTable, RatingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRatingsWith applies the HasEdge predicate on the "ratings" edge with a given conditions (other predicates).
func HasRatingsWith(preds ...predicate.Rating) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RatingsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RatingsTable, RatingsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Topic) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Topic) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Topic) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		p(s.Not())
	})
}
