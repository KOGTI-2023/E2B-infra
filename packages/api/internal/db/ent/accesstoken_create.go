// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/accesstoken"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/user"
	"github.com/google/uuid"
)

// AccessTokenCreate is the builder for creating a AccessToken entity.
type AccessTokenCreate struct {
	config
	mutation *AccessTokenMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (atc *AccessTokenCreate) SetUserID(u uuid.UUID) *AccessTokenCreate {
	atc.mutation.SetUserID(u)
	return atc
}

// SetCreatedAt sets the "created_at" field.
func (atc *AccessTokenCreate) SetCreatedAt(t time.Time) *AccessTokenCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetID sets the "id" field.
func (atc *AccessTokenCreate) SetID(s string) *AccessTokenCreate {
	atc.mutation.SetID(s)
	return atc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (atc *AccessTokenCreate) AddUserIDs(ids ...uuid.UUID) *AccessTokenCreate {
	atc.mutation.AddUserIDs(ids...)
	return atc
}

// AddUsers adds the "users" edges to the User entity.
func (atc *AccessTokenCreate) AddUsers(u ...*User) *AccessTokenCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atc.AddUserIDs(ids...)
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atc *AccessTokenCreate) Mutation() *AccessTokenMutation {
	return atc.mutation
}

// Save creates the AccessToken in the database.
func (atc *AccessTokenCreate) Save(ctx context.Context) (*AccessToken, error) {
	return withHooks(ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AccessTokenCreate) SaveX(ctx context.Context) *AccessToken {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AccessTokenCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AccessTokenCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AccessTokenCreate) check() error {
	if _, ok := atc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AccessToken.user_id"`)}
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AccessToken.created_at"`)}
	}
	return nil
}

func (atc *AccessTokenCreate) sqlSave(ctx context.Context) (*AccessToken, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AccessToken.ID type: %T", _spec.ID.Value)
		}
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *AccessTokenCreate) createSpec() (*AccessToken, *sqlgraph.CreateSpec) {
	var (
		_node = &AccessToken{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(accesstoken.Table, sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeString))
	)
	_spec.Schema = atc.schemaConfig.AccessToken
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := atc.mutation.UserID(); ok {
		_spec.SetField(accesstoken.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.SetField(accesstoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := atc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accesstoken.UsersTable,
			Columns: []string{accesstoken.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = atc.schemaConfig.User
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccessTokenCreateBulk is the builder for creating many AccessToken entities in bulk.
type AccessTokenCreateBulk struct {
	config
	err      error
	builders []*AccessTokenCreate
}

// Save creates the AccessToken entities in the database.
func (atcb *AccessTokenCreateBulk) Save(ctx context.Context) ([]*AccessToken, error) {
	if atcb.err != nil {
		return nil, atcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AccessToken, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccessTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AccessTokenCreateBulk) SaveX(ctx context.Context) []*AccessToken {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AccessTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AccessTokenCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
