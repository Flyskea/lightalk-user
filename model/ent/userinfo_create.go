// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flyskea/lightalk-user-rpc/model/ent/userinfo"
)

// UserInfoCreate is the builder for creating a UserInfo entity.
type UserInfoCreate struct {
	config
	mutation *UserInfoMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uic *UserInfoCreate) SetCreatedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetCreatedAt(t)
	return uic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableCreatedAt(t *time.Time) *UserInfoCreate {
	if t != nil {
		uic.SetCreatedAt(*t)
	}
	return uic
}

// SetUpdatedAt sets the "updated_at" field.
func (uic *UserInfoCreate) SetUpdatedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetUpdatedAt(t)
	return uic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableUpdatedAt(t *time.Time) *UserInfoCreate {
	if t != nil {
		uic.SetUpdatedAt(*t)
	}
	return uic
}

// SetDeletedAt sets the "deleted_at" field.
func (uic *UserInfoCreate) SetDeletedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetDeletedAt(t)
	return uic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableDeletedAt(t *time.Time) *UserInfoCreate {
	if t != nil {
		uic.SetDeletedAt(*t)
	}
	return uic
}

// SetUserID sets the "user_id" field.
func (uic *UserInfoCreate) SetUserID(i int64) *UserInfoCreate {
	uic.mutation.SetUserID(i)
	return uic
}

// SetSex sets the "sex" field.
func (uic *UserInfoCreate) SetSex(i int8) *UserInfoCreate {
	uic.mutation.SetSex(i)
	return uic
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableSex(i *int8) *UserInfoCreate {
	if i != nil {
		uic.SetSex(*i)
	}
	return uic
}

// SetAge sets the "age" field.
func (uic *UserInfoCreate) SetAge(i int8) *UserInfoCreate {
	uic.mutation.SetAge(i)
	return uic
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableAge(i *int8) *UserInfoCreate {
	if i != nil {
		uic.SetAge(*i)
	}
	return uic
}

// SetBirthday sets the "birthday" field.
func (uic *UserInfoCreate) SetBirthday(s string) *UserInfoCreate {
	uic.mutation.SetBirthday(s)
	return uic
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableBirthday(s *string) *UserInfoCreate {
	if s != nil {
		uic.SetBirthday(*s)
	}
	return uic
}

// SetMotto sets the "motto" field.
func (uic *UserInfoCreate) SetMotto(s string) *UserInfoCreate {
	uic.mutation.SetMotto(s)
	return uic
}

// SetNillableMotto sets the "motto" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableMotto(s *string) *UserInfoCreate {
	if s != nil {
		uic.SetMotto(*s)
	}
	return uic
}

// SetID sets the "id" field.
func (uic *UserInfoCreate) SetID(i int) *UserInfoCreate {
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
			if node, err = uic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uic.hooks) - 1; i >= 0; i-- {
			if uic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserInfoMutation", v)
		}
		node = nv
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

// Exec executes the query.
func (uic *UserInfoCreate) Exec(ctx context.Context) error {
	_, err := uic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uic *UserInfoCreate) ExecX(ctx context.Context) {
	if err := uic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uic *UserInfoCreate) defaults() {
	if _, ok := uic.mutation.CreatedAt(); !ok {
		v := userinfo.DefaultCreatedAt()
		uic.mutation.SetCreatedAt(v)
	}
	if _, ok := uic.mutation.UpdatedAt(); !ok {
		v := userinfo.DefaultUpdatedAt()
		uic.mutation.SetUpdatedAt(v)
	}
	if _, ok := uic.mutation.Sex(); !ok {
		v := userinfo.DefaultSex
		uic.mutation.SetSex(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uic *UserInfoCreate) check() error {
	if _, ok := uic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserInfo.created_at"`)}
	}
	if _, ok := uic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserInfo.updated_at"`)}
	}
	if _, ok := uic.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserInfo.user_id"`)}
	}
	if _, ok := uic.mutation.Sex(); !ok {
		return &ValidationError{Name: "sex", err: errors.New(`ent: missing required field "UserInfo.sex"`)}
	}
	return nil
}

func (uic *UserInfoCreate) sqlSave(ctx context.Context) (*UserInfo, error) {
	_node, _spec := uic.createSpec()
	if err := sqlgraph.CreateNode(ctx, uic.driver, _spec); err != nil {
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

func (uic *UserInfoCreate) createSpec() (*UserInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &UserInfo{config: uic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userinfo.FieldID,
			},
		}
	)
	if id, ok := uic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uic.mutation.CreatedAt(); ok {
		_spec.SetField(userinfo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uic.mutation.UpdatedAt(); ok {
		_spec.SetField(userinfo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uic.mutation.DeletedAt(); ok {
		_spec.SetField(userinfo.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := uic.mutation.UserID(); ok {
		_spec.SetField(userinfo.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := uic.mutation.Sex(); ok {
		_spec.SetField(userinfo.FieldSex, field.TypeInt8, value)
		_node.Sex = value
	}
	if value, ok := uic.mutation.Age(); ok {
		_spec.SetField(userinfo.FieldAge, field.TypeInt8, value)
		_node.Age = value
	}
	if value, ok := uic.mutation.Birthday(); ok {
		_spec.SetField(userinfo.FieldBirthday, field.TypeString, value)
		_node.Birthday = &value
	}
	if value, ok := uic.mutation.Motto(); ok {
		_spec.SetField(userinfo.FieldMotto, field.TypeString, value)
		_node.Motto = value
	}
	return _node, _spec
}

// UserInfoCreateBulk is the builder for creating many UserInfo entities in bulk.
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
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) SaveX(ctx context.Context) []*UserInfo {
	v, err := uicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uicb *UserInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := uicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) ExecX(ctx context.Context) {
	if err := uicb.Exec(ctx); err != nil {
		panic(err)
	}
}
