// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/thoohv5/template/internal/ent/miniprogramaccount"
	"github.com/thoohv5/template/internal/ent/predicate"
)

// MiniProgramAccountUpdate is the builder for updating MiniProgramAccount entities.
type MiniProgramAccountUpdate struct {
	config
	hooks      []Hook
	mutation   *MiniProgramAccountMutation
	predicates []predicate.MiniProgramAccount
}

// Where adds a new predicate for the builder.
func (mpau *MiniProgramAccountUpdate) Where(ps ...predicate.MiniProgramAccount) *MiniProgramAccountUpdate {
	mpau.predicates = append(mpau.predicates, ps...)
	return mpau
}

// SetUserIdentity sets the user_identity field.
func (mpau *MiniProgramAccountUpdate) SetUserIdentity(s string) *MiniProgramAccountUpdate {
	mpau.mutation.SetUserIdentity(s)
	return mpau
}

// SetNillableUserIdentity sets the user_identity field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableUserIdentity(s *string) *MiniProgramAccountUpdate {
	if s != nil {
		mpau.SetUserIdentity(*s)
	}
	return mpau
}

// ClearUserIdentity clears the value of user_identity.
func (mpau *MiniProgramAccountUpdate) ClearUserIdentity() *MiniProgramAccountUpdate {
	mpau.mutation.ClearUserIdentity()
	return mpau
}

// SetOpenID sets the open_id field.
func (mpau *MiniProgramAccountUpdate) SetOpenID(s string) *MiniProgramAccountUpdate {
	mpau.mutation.SetOpenID(s)
	return mpau
}

// SetNillableOpenID sets the open_id field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableOpenID(s *string) *MiniProgramAccountUpdate {
	if s != nil {
		mpau.SetOpenID(*s)
	}
	return mpau
}

// ClearOpenID clears the value of open_id.
func (mpau *MiniProgramAccountUpdate) ClearOpenID() *MiniProgramAccountUpdate {
	mpau.mutation.ClearOpenID()
	return mpau
}

// SetNickName sets the nick_name field.
func (mpau *MiniProgramAccountUpdate) SetNickName(s string) *MiniProgramAccountUpdate {
	mpau.mutation.SetNickName(s)
	return mpau
}

// SetNillableNickName sets the nick_name field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableNickName(s *string) *MiniProgramAccountUpdate {
	if s != nil {
		mpau.SetNickName(*s)
	}
	return mpau
}

// SetAvatarURL sets the avatar_url field.
func (mpau *MiniProgramAccountUpdate) SetAvatarURL(s string) *MiniProgramAccountUpdate {
	mpau.mutation.SetAvatarURL(s)
	return mpau
}

// SetNillableAvatarURL sets the avatar_url field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableAvatarURL(s *string) *MiniProgramAccountUpdate {
	if s != nil {
		mpau.SetAvatarURL(*s)
	}
	return mpau
}

// SetGender sets the gender field.
func (mpau *MiniProgramAccountUpdate) SetGender(i int32) *MiniProgramAccountUpdate {
	mpau.mutation.ResetGender()
	mpau.mutation.SetGender(i)
	return mpau
}

// SetNillableGender sets the gender field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableGender(i *int32) *MiniProgramAccountUpdate {
	if i != nil {
		mpau.SetGender(*i)
	}
	return mpau
}

// AddGender adds i to gender.
func (mpau *MiniProgramAccountUpdate) AddGender(i int32) *MiniProgramAccountUpdate {
	mpau.mutation.AddGender(i)
	return mpau
}

// SetCountry sets the country field.
func (mpau *MiniProgramAccountUpdate) SetCountry(s string) *MiniProgramAccountUpdate {
	mpau.mutation.SetCountry(s)
	return mpau
}

// SetNillableCountry sets the country field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableCountry(s *string) *MiniProgramAccountUpdate {
	if s != nil {
		mpau.SetCountry(*s)
	}
	return mpau
}

// SetProvince sets the province field.
func (mpau *MiniProgramAccountUpdate) SetProvince(s string) *MiniProgramAccountUpdate {
	mpau.mutation.SetProvince(s)
	return mpau
}

// SetNillableProvince sets the province field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableProvince(s *string) *MiniProgramAccountUpdate {
	if s != nil {
		mpau.SetProvince(*s)
	}
	return mpau
}

// SetCity sets the city field.
func (mpau *MiniProgramAccountUpdate) SetCity(s string) *MiniProgramAccountUpdate {
	mpau.mutation.SetCity(s)
	return mpau
}

// SetNillableCity sets the city field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableCity(s *string) *MiniProgramAccountUpdate {
	if s != nil {
		mpau.SetCity(*s)
	}
	return mpau
}

// SetLanguage sets the language field.
func (mpau *MiniProgramAccountUpdate) SetLanguage(s string) *MiniProgramAccountUpdate {
	mpau.mutation.SetLanguage(s)
	return mpau
}

// SetNillableLanguage sets the language field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableLanguage(s *string) *MiniProgramAccountUpdate {
	if s != nil {
		mpau.SetLanguage(*s)
	}
	return mpau
}

// SetCreatedAt sets the created_at field.
func (mpau *MiniProgramAccountUpdate) SetCreatedAt(t time.Time) *MiniProgramAccountUpdate {
	mpau.mutation.SetCreatedAt(t)
	return mpau
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableCreatedAt(t *time.Time) *MiniProgramAccountUpdate {
	if t != nil {
		mpau.SetCreatedAt(*t)
	}
	return mpau
}

// ClearCreatedAt clears the value of created_at.
func (mpau *MiniProgramAccountUpdate) ClearCreatedAt() *MiniProgramAccountUpdate {
	mpau.mutation.ClearCreatedAt()
	return mpau
}

// SetUpdatedAt sets the updated_at field.
func (mpau *MiniProgramAccountUpdate) SetUpdatedAt(t time.Time) *MiniProgramAccountUpdate {
	mpau.mutation.SetUpdatedAt(t)
	return mpau
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableUpdatedAt(t *time.Time) *MiniProgramAccountUpdate {
	if t != nil {
		mpau.SetUpdatedAt(*t)
	}
	return mpau
}

// ClearUpdatedAt clears the value of updated_at.
func (mpau *MiniProgramAccountUpdate) ClearUpdatedAt() *MiniProgramAccountUpdate {
	mpau.mutation.ClearUpdatedAt()
	return mpau
}

// SetDeletedAt sets the deleted_at field.
func (mpau *MiniProgramAccountUpdate) SetDeletedAt(t time.Time) *MiniProgramAccountUpdate {
	mpau.mutation.SetDeletedAt(t)
	return mpau
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (mpau *MiniProgramAccountUpdate) SetNillableDeletedAt(t *time.Time) *MiniProgramAccountUpdate {
	if t != nil {
		mpau.SetDeletedAt(*t)
	}
	return mpau
}

// ClearDeletedAt clears the value of deleted_at.
func (mpau *MiniProgramAccountUpdate) ClearDeletedAt() *MiniProgramAccountUpdate {
	mpau.mutation.ClearDeletedAt()
	return mpau
}

// Mutation returns the MiniProgramAccountMutation object of the builder.
func (mpau *MiniProgramAccountUpdate) Mutation() *MiniProgramAccountMutation {
	return mpau.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mpau *MiniProgramAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mpau.hooks) == 0 {
		affected, err = mpau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiniProgramAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation user %T", m)
			}
			mpau.mutation = mutation
			affected, err = mpau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mpau.hooks) - 1; i >= 0; i-- {
			mut = mpau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mpau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mpau *MiniProgramAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := mpau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mpau *MiniProgramAccountUpdate) Exec(ctx context.Context) error {
	_, err := mpau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpau *MiniProgramAccountUpdate) ExecX(ctx context.Context) {
	if err := mpau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mpau *MiniProgramAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   miniprogramaccount.Table,
			Columns: miniprogramaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: miniprogramaccount.FieldID,
			},
		},
	}
	if ps := mpau.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpau.mutation.UserIdentity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldUserIdentity,
		})
	}
	if mpau.mutation.UserIdentityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: miniprogramaccount.FieldUserIdentity,
		})
	}
	if value, ok := mpau.mutation.OpenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldOpenID,
		})
	}
	if mpau.mutation.OpenIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: miniprogramaccount.FieldOpenID,
		})
	}
	if value, ok := mpau.mutation.NickName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldNickName,
		})
	}
	if value, ok := mpau.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldAvatarURL,
		})
	}
	if value, ok := mpau.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: miniprogramaccount.FieldGender,
		})
	}
	if value, ok := mpau.mutation.AddedGender(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: miniprogramaccount.FieldGender,
		})
	}
	if value, ok := mpau.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldCountry,
		})
	}
	if value, ok := mpau.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldProvince,
		})
	}
	if value, ok := mpau.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldCity,
		})
	}
	if value, ok := mpau.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldLanguage,
		})
	}
	if value, ok := mpau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: miniprogramaccount.FieldCreatedAt,
		})
	}
	if mpau.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: miniprogramaccount.FieldCreatedAt,
		})
	}
	if value, ok := mpau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: miniprogramaccount.FieldUpdatedAt,
		})
	}
	if mpau.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: miniprogramaccount.FieldUpdatedAt,
		})
	}
	if value, ok := mpau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: miniprogramaccount.FieldDeletedAt,
		})
	}
	if mpau.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: miniprogramaccount.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mpau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{miniprogramaccount.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MiniProgramAccountUpdateOne is the builder for updating a single MiniProgramAccount entity.
type MiniProgramAccountUpdateOne struct {
	config
	hooks    []Hook
	mutation *MiniProgramAccountMutation
}

// SetUserIdentity sets the user_identity field.
func (mpauo *MiniProgramAccountUpdateOne) SetUserIdentity(s string) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetUserIdentity(s)
	return mpauo
}

// SetNillableUserIdentity sets the user_identity field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableUserIdentity(s *string) *MiniProgramAccountUpdateOne {
	if s != nil {
		mpauo.SetUserIdentity(*s)
	}
	return mpauo
}

// ClearUserIdentity clears the value of user_identity.
func (mpauo *MiniProgramAccountUpdateOne) ClearUserIdentity() *MiniProgramAccountUpdateOne {
	mpauo.mutation.ClearUserIdentity()
	return mpauo
}

// SetOpenID sets the open_id field.
func (mpauo *MiniProgramAccountUpdateOne) SetOpenID(s string) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetOpenID(s)
	return mpauo
}

// SetNillableOpenID sets the open_id field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableOpenID(s *string) *MiniProgramAccountUpdateOne {
	if s != nil {
		mpauo.SetOpenID(*s)
	}
	return mpauo
}

// ClearOpenID clears the value of open_id.
func (mpauo *MiniProgramAccountUpdateOne) ClearOpenID() *MiniProgramAccountUpdateOne {
	mpauo.mutation.ClearOpenID()
	return mpauo
}

// SetNickName sets the nick_name field.
func (mpauo *MiniProgramAccountUpdateOne) SetNickName(s string) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetNickName(s)
	return mpauo
}

// SetNillableNickName sets the nick_name field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableNickName(s *string) *MiniProgramAccountUpdateOne {
	if s != nil {
		mpauo.SetNickName(*s)
	}
	return mpauo
}

// SetAvatarURL sets the avatar_url field.
func (mpauo *MiniProgramAccountUpdateOne) SetAvatarURL(s string) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetAvatarURL(s)
	return mpauo
}

// SetNillableAvatarURL sets the avatar_url field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableAvatarURL(s *string) *MiniProgramAccountUpdateOne {
	if s != nil {
		mpauo.SetAvatarURL(*s)
	}
	return mpauo
}

// SetGender sets the gender field.
func (mpauo *MiniProgramAccountUpdateOne) SetGender(i int32) *MiniProgramAccountUpdateOne {
	mpauo.mutation.ResetGender()
	mpauo.mutation.SetGender(i)
	return mpauo
}

// SetNillableGender sets the gender field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableGender(i *int32) *MiniProgramAccountUpdateOne {
	if i != nil {
		mpauo.SetGender(*i)
	}
	return mpauo
}

// AddGender adds i to gender.
func (mpauo *MiniProgramAccountUpdateOne) AddGender(i int32) *MiniProgramAccountUpdateOne {
	mpauo.mutation.AddGender(i)
	return mpauo
}

// SetCountry sets the country field.
func (mpauo *MiniProgramAccountUpdateOne) SetCountry(s string) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetCountry(s)
	return mpauo
}

// SetNillableCountry sets the country field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableCountry(s *string) *MiniProgramAccountUpdateOne {
	if s != nil {
		mpauo.SetCountry(*s)
	}
	return mpauo
}

// SetProvince sets the province field.
func (mpauo *MiniProgramAccountUpdateOne) SetProvince(s string) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetProvince(s)
	return mpauo
}

// SetNillableProvince sets the province field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableProvince(s *string) *MiniProgramAccountUpdateOne {
	if s != nil {
		mpauo.SetProvince(*s)
	}
	return mpauo
}

// SetCity sets the city field.
func (mpauo *MiniProgramAccountUpdateOne) SetCity(s string) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetCity(s)
	return mpauo
}

// SetNillableCity sets the city field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableCity(s *string) *MiniProgramAccountUpdateOne {
	if s != nil {
		mpauo.SetCity(*s)
	}
	return mpauo
}

// SetLanguage sets the language field.
func (mpauo *MiniProgramAccountUpdateOne) SetLanguage(s string) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetLanguage(s)
	return mpauo
}

// SetNillableLanguage sets the language field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableLanguage(s *string) *MiniProgramAccountUpdateOne {
	if s != nil {
		mpauo.SetLanguage(*s)
	}
	return mpauo
}

// SetCreatedAt sets the created_at field.
func (mpauo *MiniProgramAccountUpdateOne) SetCreatedAt(t time.Time) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetCreatedAt(t)
	return mpauo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableCreatedAt(t *time.Time) *MiniProgramAccountUpdateOne {
	if t != nil {
		mpauo.SetCreatedAt(*t)
	}
	return mpauo
}

// ClearCreatedAt clears the value of created_at.
func (mpauo *MiniProgramAccountUpdateOne) ClearCreatedAt() *MiniProgramAccountUpdateOne {
	mpauo.mutation.ClearCreatedAt()
	return mpauo
}

// SetUpdatedAt sets the updated_at field.
func (mpauo *MiniProgramAccountUpdateOne) SetUpdatedAt(t time.Time) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetUpdatedAt(t)
	return mpauo
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableUpdatedAt(t *time.Time) *MiniProgramAccountUpdateOne {
	if t != nil {
		mpauo.SetUpdatedAt(*t)
	}
	return mpauo
}

// ClearUpdatedAt clears the value of updated_at.
func (mpauo *MiniProgramAccountUpdateOne) ClearUpdatedAt() *MiniProgramAccountUpdateOne {
	mpauo.mutation.ClearUpdatedAt()
	return mpauo
}

// SetDeletedAt sets the deleted_at field.
func (mpauo *MiniProgramAccountUpdateOne) SetDeletedAt(t time.Time) *MiniProgramAccountUpdateOne {
	mpauo.mutation.SetDeletedAt(t)
	return mpauo
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (mpauo *MiniProgramAccountUpdateOne) SetNillableDeletedAt(t *time.Time) *MiniProgramAccountUpdateOne {
	if t != nil {
		mpauo.SetDeletedAt(*t)
	}
	return mpauo
}

// ClearDeletedAt clears the value of deleted_at.
func (mpauo *MiniProgramAccountUpdateOne) ClearDeletedAt() *MiniProgramAccountUpdateOne {
	mpauo.mutation.ClearDeletedAt()
	return mpauo
}

// Mutation returns the MiniProgramAccountMutation object of the builder.
func (mpauo *MiniProgramAccountUpdateOne) Mutation() *MiniProgramAccountMutation {
	return mpauo.mutation
}

// Save executes the query and returns the updated entity.
func (mpauo *MiniProgramAccountUpdateOne) Save(ctx context.Context) (*MiniProgramAccount, error) {
	var (
		err  error
		node *MiniProgramAccount
	)
	if len(mpauo.hooks) == 0 {
		node, err = mpauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiniProgramAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation user %T", m)
			}
			mpauo.mutation = mutation
			node, err = mpauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mpauo.hooks) - 1; i >= 0; i-- {
			mut = mpauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mpauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mpauo *MiniProgramAccountUpdateOne) SaveX(ctx context.Context) *MiniProgramAccount {
	node, err := mpauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mpauo *MiniProgramAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := mpauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpauo *MiniProgramAccountUpdateOne) ExecX(ctx context.Context) {
	if err := mpauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mpauo *MiniProgramAccountUpdateOne) sqlSave(ctx context.Context) (_node *MiniProgramAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   miniprogramaccount.Table,
			Columns: miniprogramaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: miniprogramaccount.FieldID,
			},
		},
	}
	id, ok := mpauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing MiniProgramAccount.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := mpauo.mutation.UserIdentity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldUserIdentity,
		})
	}
	if mpauo.mutation.UserIdentityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: miniprogramaccount.FieldUserIdentity,
		})
	}
	if value, ok := mpauo.mutation.OpenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldOpenID,
		})
	}
	if mpauo.mutation.OpenIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: miniprogramaccount.FieldOpenID,
		})
	}
	if value, ok := mpauo.mutation.NickName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldNickName,
		})
	}
	if value, ok := mpauo.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldAvatarURL,
		})
	}
	if value, ok := mpauo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: miniprogramaccount.FieldGender,
		})
	}
	if value, ok := mpauo.mutation.AddedGender(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: miniprogramaccount.FieldGender,
		})
	}
	if value, ok := mpauo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldCountry,
		})
	}
	if value, ok := mpauo.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldProvince,
		})
	}
	if value, ok := mpauo.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldCity,
		})
	}
	if value, ok := mpauo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: miniprogramaccount.FieldLanguage,
		})
	}
	if value, ok := mpauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: miniprogramaccount.FieldCreatedAt,
		})
	}
	if mpauo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: miniprogramaccount.FieldCreatedAt,
		})
	}
	if value, ok := mpauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: miniprogramaccount.FieldUpdatedAt,
		})
	}
	if mpauo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: miniprogramaccount.FieldUpdatedAt,
		})
	}
	if value, ok := mpauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: miniprogramaccount.FieldDeletedAt,
		})
	}
	if mpauo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: miniprogramaccount.FieldDeletedAt,
		})
	}
	_node = &MiniProgramAccount{config: mpauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, mpauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{miniprogramaccount.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
