// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/thoohv5/template/internal/ent/phoneaccount"
	"github.com/thoohv5/template/internal/ent/predicate"
)

// PhoneAccountDelete is the builder for deleting a PhoneAccount entity.
type PhoneAccountDelete struct {
	config
	hooks      []Hook
	mutation   *PhoneAccountMutation
	predicates []predicate.PhoneAccount
}

// Where adds a new predicate to the delete builder.
func (pad *PhoneAccountDelete) Where(ps ...predicate.PhoneAccount) *PhoneAccountDelete {
	pad.predicates = append(pad.predicates, ps...)
	return pad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pad *PhoneAccountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pad.hooks) == 0 {
		affected, err = pad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhoneAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation user %T", m)
			}
			pad.mutation = mutation
			affected, err = pad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pad.hooks) - 1; i >= 0; i-- {
			mut = pad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pad *PhoneAccountDelete) ExecX(ctx context.Context) int {
	n, err := pad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pad *PhoneAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: phoneaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: phoneaccount.FieldID,
			},
		},
	}
	if ps := pad.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pad.driver, _spec)
}

// PhoneAccountDeleteOne is the builder for deleting a single PhoneAccount entity.
type PhoneAccountDeleteOne struct {
	pad *PhoneAccountDelete
}

// Exec executes the deletion query.
func (pado *PhoneAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := pado.pad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{phoneaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pado *PhoneAccountDeleteOne) ExecX(ctx context.Context) {
	pado.pad.ExecX(ctx)
}
