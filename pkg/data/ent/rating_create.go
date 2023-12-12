// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/rating"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/topic"
	"github.com/mizumoto-cn/ratingo/pkg/data/ent/user"
)

// RatingCreate is the builder for creating a Rating entity.
type RatingCreate struct {
	config
	mutation *RatingMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (rc *RatingCreate) SetUserID(i int) *RatingCreate {
	rc.mutation.SetUserID(i)
	return rc
}

// SetTopicID sets the "topic_id" field.
func (rc *RatingCreate) SetTopicID(i int) *RatingCreate {
	rc.mutation.SetTopicID(i)
	return rc
}

// SetRating sets the "rating" field.
func (rc *RatingCreate) SetRating(f float64) *RatingCreate {
	rc.mutation.SetRating(f)
	return rc
}

// SetComment sets the "comment" field.
func (rc *RatingCreate) SetComment(s string) *RatingCreate {
	rc.mutation.SetComment(s)
	return rc
}

// SetID sets the "id" field.
func (rc *RatingCreate) SetID(i int) *RatingCreate {
	rc.mutation.SetID(i)
	return rc
}

// SetRatedByID sets the "ratedBy" edge to the User entity by ID.
func (rc *RatingCreate) SetRatedByID(id int) *RatingCreate {
	rc.mutation.SetRatedByID(id)
	return rc
}

// SetNillableRatedByID sets the "ratedBy" edge to the User entity by ID if the given value is not nil.
func (rc *RatingCreate) SetNillableRatedByID(id *int) *RatingCreate {
	if id != nil {
		rc = rc.SetRatedByID(*id)
	}
	return rc
}

// SetRatedBy sets the "ratedBy" edge to the User entity.
func (rc *RatingCreate) SetRatedBy(u *User) *RatingCreate {
	return rc.SetRatedByID(u.ID)
}

// SetUnderTopicOfID sets the "underTopicOf" edge to the Topic entity by ID.
func (rc *RatingCreate) SetUnderTopicOfID(id int) *RatingCreate {
	rc.mutation.SetUnderTopicOfID(id)
	return rc
}

// SetNillableUnderTopicOfID sets the "underTopicOf" edge to the Topic entity by ID if the given value is not nil.
func (rc *RatingCreate) SetNillableUnderTopicOfID(id *int) *RatingCreate {
	if id != nil {
		rc = rc.SetUnderTopicOfID(*id)
	}
	return rc
}

// SetUnderTopicOf sets the "underTopicOf" edge to the Topic entity.
func (rc *RatingCreate) SetUnderTopicOf(t *Topic) *RatingCreate {
	return rc.SetUnderTopicOfID(t.ID)
}

// Mutation returns the RatingMutation object of the builder.
func (rc *RatingCreate) Mutation() *RatingMutation {
	return rc.mutation
}

// Save creates the Rating in the database.
func (rc *RatingCreate) Save(ctx context.Context) (*Rating, error) {
	var (
		err  error
		node *Rating
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RatingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Rating)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RatingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RatingCreate) SaveX(ctx context.Context) *Rating {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RatingCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RatingCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RatingCreate) check() error {
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Rating.user_id"`)}
	}
	if _, ok := rc.mutation.TopicID(); !ok {
		return &ValidationError{Name: "topic_id", err: errors.New(`ent: missing required field "Rating.topic_id"`)}
	}
	if _, ok := rc.mutation.Rating(); !ok {
		return &ValidationError{Name: "rating", err: errors.New(`ent: missing required field "Rating.rating"`)}
	}
	if _, ok := rc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "Rating.comment"`)}
	}
	return nil
}

func (rc *RatingCreate) sqlSave(ctx context.Context) (*Rating, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (rc *RatingCreate) createSpec() (*Rating, *sqlgraph.CreateSpec) {
	var (
		_node = &Rating{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rating.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rating.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.UserID(); ok {
		_spec.SetField(rating.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := rc.mutation.TopicID(); ok {
		_spec.SetField(rating.FieldTopicID, field.TypeInt, value)
		_node.TopicID = value
	}
	if value, ok := rc.mutation.Rating(); ok {
		_spec.SetField(rating.FieldRating, field.TypeFloat64, value)
		_node.Rating = value
	}
	if value, ok := rc.mutation.Comment(); ok {
		_spec.SetField(rating.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if nodes := rc.mutation.RatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rating.RatedByTable,
			Columns: []string{rating.RatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_ratings = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UnderTopicOfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rating.UnderTopicOfTable,
			Columns: []string{rating.UnderTopicOfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.topic_ratings = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RatingCreateBulk is the builder for creating many Rating entities in bulk.
type RatingCreateBulk struct {
	config
	builders []*RatingCreate
}

// Save creates the Rating entities in the database.
func (rcb *RatingCreateBulk) Save(ctx context.Context) ([]*Rating, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rating, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RatingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RatingCreateBulk) SaveX(ctx context.Context) []*Rating {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RatingCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RatingCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
