// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/thoohv5/template/internal/ent/userinfo"
)

// UserInfoCreate is the builder for creating a UserInfo entity.
type UserInfoCreate struct {
	config
	mutation *UserInfoMutation
	hooks    []Hook
}

// SetUserIdentity sets the user_identity field.
func (uic *UserInfoCreate) SetUserIdentity(s string) *UserInfoCreate {
	uic.mutation.SetUserIdentity(s)
	return uic
}

// SetNillableUserIdentity sets the user_identity field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableUserIdentity(s *string) *UserInfoCreate {
	if s != nil {
		uic.SetUserIdentity(*s)
	}
	return uic
}

// SetChannel sets the channel field.
func (uic *UserInfoCreate) SetChannel(i int32) *UserInfoCreate {
	uic.mutation.SetChannel(i)
	return uic
}

// SetNillableChannel sets the channel field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableChannel(i *int32) *UserInfoCreate {
	if i != nil {
		uic.SetChannel(*i)
	}
	return uic
}

// SetForm sets the form field.
func (uic *UserInfoCreate) SetForm(i int32) *UserInfoCreate {
	uic.mutation.SetForm(i)
	return uic
}

// SetNillableForm sets the form field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableForm(i *int32) *UserInfoCreate {
	if i != nil {
		uic.SetForm(*i)
	}
	return uic
}

// SetCreatedAt sets the created_at field.
func (uic *UserInfoCreate) SetCreatedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetCreatedAt(t)
	return uic
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableCreatedAt(t *time.Time) *UserInfoCreate {
	if t != nil {
		uic.SetCreatedAt(*t)
	}
	return uic
}

// SetUpdatedAt sets the updated_at field.
func (uic *UserInfoCreate) SetUpdatedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetUpdatedAt(t)
	return uic
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableUpdatedAt(t *time.Time) *UserInfoCreate {
	if t != nil {
		uic.SetUpdatedAt(*t)
	}
	return uic
}

// SetDeletedAt sets the deleted_at field.
func (uic *UserInfoCreate) SetDeletedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetDeletedAt(t)
	return uic
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableDeletedAt(t *time.Time) *UserInfoCreate {
	if t != nil {
		uic.SetDeletedAt(*t)
	}
	return uic
}

// SetID sets the id field.
func (uic *UserInfoCreate) SetID(i int64) *UserInfoCreate {
	uic.mutation.SetID(i)
	return uic
}

// Mutation returns the UserInfoMutation object of the builder.
func (uic *UserInfoCreate) Mutation() *UserInfoMutation {
	return uic.mutation
}

// Save creates the UserInfo in the database.
func (uic *UserInfoCreate) Save(ctx context.Context) (*UserInfo, error) {
	var (
		err  error
		node *UserInfo
	)
	uic.defaults()
	if len(uic.hooks) == 0 {
		if err = uic.check(); err != nil {
			return nil, err
		}
		node, err = uic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uic.check(); err != nil {
				return nil, err
			}
			uic.mutation = mutation
			node, err = uic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uic.hooks) - 1; i >= 0; i-- {
			mut = uic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uic *UserInfoCreate) SaveX(ctx context.Context) *UserInfo {
	v, err := uic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uic *UserInfoCreate) defaults() {
	if _, ok := uic.mutation.Channel(); !ok {
		v := userinfo.DefaultChannel
		uic.mutation.SetChannel(v)
	}
	if _, ok := uic.mutation.Form(); !ok {
		v := userinfo.DefaultForm
		uic.mutation.SetForm(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uic *UserInfoCreate) check() error {
	if _, ok := uic.mutation.Channel(); !ok {
		return &ValidationError{Name: "channel", err: errors.New("ent: missing required field \"channel\"")}
	}
	if _, ok := uic.mutation.Form(); !ok {
		return &ValidationError{Name: "form", err: errors.New("ent: missing required field \"form\"")}
	}
	return nil
}

func (uic *UserInfoCreate) sqlSave(ctx context.Context) (*UserInfo, error) {
	_node, _spec := uic.createSpec()
	if err := sqlgraph.CreateNode(ctx, uic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (uic *UserInfoCreate) createSpec() (*UserInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &UserInfo{config: uic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userinfo.FieldID,
			},
		}
	)
	if id, ok := uic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uic.mutation.UserIdentity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userinfo.FieldUserIdentity,
		})
		_node.UserIdentity = value
	}
	if value, ok := uic.mutation.Channel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userinfo.FieldChannel,
		})
		_node.Channel = value
	}
	if value, ok := uic.mutation.Form(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userinfo.FieldForm,
		})
		_node.Form = value
	}
	if value, ok := uic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userinfo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userinfo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uic.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userinfo.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	return _node, _spec
}

// UserInfoCreateBulk is the builder for creating a bulk of UserInfo entities.
type UserInfoCreateBulk struct {
	config
	builders []*UserInfoCreate
}

// Save creates the UserInfo entities in the database.
func (uicb *UserInfoCreateBulk) Save(ctx context.Context) ([]*UserInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uicb.builders))
	nodes := make([]*UserInfo, len(uicb.builders))
	mutators := make([]Mutator, len(uicb.builders))
	for i := range uicb.builders {
		func(i int, root context.Context) {
			builder := uicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, uicb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uicb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (uicb *UserInfoCreateBulk) SaveX(ctx context.Context) []*UserInfo {
	v, err := uicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
