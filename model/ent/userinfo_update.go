// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flyskea/lightalk-user/model/ent/predicate"
	"github.com/flyskea/lightalk-user/model/ent/userinfo"
)

// UserInfoUpdate is the builder for updating UserInfo entities.
type UserInfoUpdate struct {
	config
	hooks    []Hook
	mutation *UserInfoMutation
}

// Where appends a list predicates to the UserInfoUpdate builder.
func (uiu *UserInfoUpdate) Where(ps ...predicate.UserInfo) *UserInfoUpdate {
	uiu.mutation.Where(ps...)
	return uiu
}

// SetUpdatedAt sets the "updated_at" field.
func (uiu *UserInfoUpdate) SetUpdatedAt(t time.Time) *UserInfoUpdate {
	uiu.mutation.SetUpdatedAt(t)
	return uiu
}

// SetDeletedAt sets the "deleted_at" field.
func (uiu *UserInfoUpdate) SetDeletedAt(t time.Time) *UserInfoUpdate {
	uiu.mutation.SetDeletedAt(t)
	return uiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uiu *UserInfoUpdate) SetNillableDeletedAt(t *time.Time) *UserInfoUpdate {
	if t != nil {
		uiu.SetDeletedAt(*t)
	}
	return uiu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uiu *UserInfoUpdate) ClearDeletedAt() *UserInfoUpdate {
	uiu.mutation.ClearDeletedAt()
	return uiu
}

// SetUserID sets the "user_id" field.
func (uiu *UserInfoUpdate) SetUserID(i int64) *UserInfoUpdate {
	uiu.mutation.ResetUserID()
	uiu.mutation.SetUserID(i)
	return uiu
}

// AddUserID adds i to the "user_id" field.
func (uiu *UserInfoUpdate) AddUserID(i int64) *UserInfoUpdate {
	uiu.mutation.AddUserID(i)
	return uiu
}

// SetSex sets the "sex" field.
func (uiu *UserInfoUpdate) SetSex(i int8) *UserInfoUpdate {
	uiu.mutation.ResetSex()
	uiu.mutation.SetSex(i)
	return uiu
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (uiu *UserInfoUpdate) SetNillableSex(i *int8) *UserInfoUpdate {
	if i != nil {
		uiu.SetSex(*i)
	}
	return uiu
}

// AddSex adds i to the "sex" field.
func (uiu *UserInfoUpdate) AddSex(i int8) *UserInfoUpdate {
	uiu.mutation.AddSex(i)
	return uiu
}

// SetAge sets the "age" field.
func (uiu *UserInfoUpdate) SetAge(i int8) *UserInfoUpdate {
	uiu.mutation.ResetAge()
	uiu.mutation.SetAge(i)
	return uiu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uiu *UserInfoUpdate) SetNillableAge(i *int8) *UserInfoUpdate {
	if i != nil {
		uiu.SetAge(*i)
	}
	return uiu
}

// AddAge adds i to the "age" field.
func (uiu *UserInfoUpdate) AddAge(i int8) *UserInfoUpdate {
	uiu.mutation.AddAge(i)
	return uiu
}

// ClearAge clears the value of the "age" field.
func (uiu *UserInfoUpdate) ClearAge() *UserInfoUpdate {
	uiu.mutation.ClearAge()
	return uiu
}

// SetBirthday sets the "birthday" field.
func (uiu *UserInfoUpdate) SetBirthday(s string) *UserInfoUpdate {
	uiu.mutation.SetBirthday(s)
	return uiu
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uiu *UserInfoUpdate) SetNillableBirthday(s *string) *UserInfoUpdate {
	if s != nil {
		uiu.SetBirthday(*s)
	}
	return uiu
}

// ClearBirthday clears the value of the "birthday" field.
func (uiu *UserInfoUpdate) ClearBirthday() *UserInfoUpdate {
	uiu.mutation.ClearBirthday()
	return uiu
}

// SetMotto sets the "motto" field.
func (uiu *UserInfoUpdate) SetMotto(s string) *UserInfoUpdate {
	uiu.mutation.SetMotto(s)
	return uiu
}

// SetNillableMotto sets the "motto" field if the given value is not nil.
func (uiu *UserInfoUpdate) SetNillableMotto(s *string) *UserInfoUpdate {
	if s != nil {
		uiu.SetMotto(*s)
	}
	return uiu
}

// ClearMotto clears the value of the "motto" field.
func (uiu *UserInfoUpdate) ClearMotto() *UserInfoUpdate {
	uiu.mutation.ClearMotto()
	return uiu
}

// Mutation returns the UserInfoMutation object of the builder.
func (uiu *UserInfoUpdate) Mutation() *UserInfoMutation {
	return uiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uiu *UserInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uiu.defaults()
	if len(uiu.hooks) == 0 {
		affected, err = uiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uiu.mutation = mutation
			affected, err = uiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uiu.hooks) - 1; i >= 0; i-- {
			if uiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uiu *UserInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := uiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uiu *UserInfoUpdate) Exec(ctx context.Context) error {
	_, err := uiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uiu *UserInfoUpdate) ExecX(ctx context.Context) {
	if err := uiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uiu *UserInfoUpdate) defaults() {
	if _, ok := uiu.mutation.UpdatedAt(); !ok {
		v := userinfo.UpdateDefaultUpdatedAt()
		uiu.mutation.SetUpdatedAt(v)
	}
}

func (uiu *UserInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userinfo.Table,
			Columns: userinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userinfo.FieldID,
			},
		},
	}
	if ps := uiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uiu.mutation.UpdatedAt(); ok {
		_spec.SetField(userinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uiu.mutation.DeletedAt(); ok {
		_spec.SetField(userinfo.FieldDeletedAt, field.TypeTime, value)
	}
	if uiu.mutation.DeletedAtCleared() {
		_spec.ClearField(userinfo.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uiu.mutation.UserID(); ok {
		_spec.SetField(userinfo.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := uiu.mutation.AddedUserID(); ok {
		_spec.AddField(userinfo.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := uiu.mutation.Sex(); ok {
		_spec.SetField(userinfo.FieldSex, field.TypeInt8, value)
	}
	if value, ok := uiu.mutation.AddedSex(); ok {
		_spec.AddField(userinfo.FieldSex, field.TypeInt8, value)
	}
	if value, ok := uiu.mutation.Age(); ok {
		_spec.SetField(userinfo.FieldAge, field.TypeInt8, value)
	}
	if value, ok := uiu.mutation.AddedAge(); ok {
		_spec.AddField(userinfo.FieldAge, field.TypeInt8, value)
	}
	if uiu.mutation.AgeCleared() {
		_spec.ClearField(userinfo.FieldAge, field.TypeInt8)
	}
	if value, ok := uiu.mutation.Birthday(); ok {
		_spec.SetField(userinfo.FieldBirthday, field.TypeString, value)
	}
	if uiu.mutation.BirthdayCleared() {
		_spec.ClearField(userinfo.FieldBirthday, field.TypeString)
	}
	if value, ok := uiu.mutation.Motto(); ok {
		_spec.SetField(userinfo.FieldMotto, field.TypeString, value)
	}
	if uiu.mutation.MottoCleared() {
		_spec.ClearField(userinfo.FieldMotto, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserInfoUpdateOne is the builder for updating a single UserInfo entity.
type UserInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserInfoMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uiuo *UserInfoUpdateOne) SetUpdatedAt(t time.Time) *UserInfoUpdateOne {
	uiuo.mutation.SetUpdatedAt(t)
	return uiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uiuo *UserInfoUpdateOne) SetDeletedAt(t time.Time) *UserInfoUpdateOne {
	uiuo.mutation.SetDeletedAt(t)
	return uiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uiuo *UserInfoUpdateOne) SetNillableDeletedAt(t *time.Time) *UserInfoUpdateOne {
	if t != nil {
		uiuo.SetDeletedAt(*t)
	}
	return uiuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uiuo *UserInfoUpdateOne) ClearDeletedAt() *UserInfoUpdateOne {
	uiuo.mutation.ClearDeletedAt()
	return uiuo
}

// SetUserID sets the "user_id" field.
func (uiuo *UserInfoUpdateOne) SetUserID(i int64) *UserInfoUpdateOne {
	uiuo.mutation.ResetUserID()
	uiuo.mutation.SetUserID(i)
	return uiuo
}

// AddUserID adds i to the "user_id" field.
func (uiuo *UserInfoUpdateOne) AddUserID(i int64) *UserInfoUpdateOne {
	uiuo.mutation.AddUserID(i)
	return uiuo
}

// SetSex sets the "sex" field.
func (uiuo *UserInfoUpdateOne) SetSex(i int8) *UserInfoUpdateOne {
	uiuo.mutation.ResetSex()
	uiuo.mutation.SetSex(i)
	return uiuo
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (uiuo *UserInfoUpdateOne) SetNillableSex(i *int8) *UserInfoUpdateOne {
	if i != nil {
		uiuo.SetSex(*i)
	}
	return uiuo
}

// AddSex adds i to the "sex" field.
func (uiuo *UserInfoUpdateOne) AddSex(i int8) *UserInfoUpdateOne {
	uiuo.mutation.AddSex(i)
	return uiuo
}

// SetAge sets the "age" field.
func (uiuo *UserInfoUpdateOne) SetAge(i int8) *UserInfoUpdateOne {
	uiuo.mutation.ResetAge()
	uiuo.mutation.SetAge(i)
	return uiuo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uiuo *UserInfoUpdateOne) SetNillableAge(i *int8) *UserInfoUpdateOne {
	if i != nil {
		uiuo.SetAge(*i)
	}
	return uiuo
}

// AddAge adds i to the "age" field.
func (uiuo *UserInfoUpdateOne) AddAge(i int8) *UserInfoUpdateOne {
	uiuo.mutation.AddAge(i)
	return uiuo
}

// ClearAge clears the value of the "age" field.
func (uiuo *UserInfoUpdateOne) ClearAge() *UserInfoUpdateOne {
	uiuo.mutation.ClearAge()
	return uiuo
}

// SetBirthday sets the "birthday" field.
func (uiuo *UserInfoUpdateOne) SetBirthday(s string) *UserInfoUpdateOne {
	uiuo.mutation.SetBirthday(s)
	return uiuo
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uiuo *UserInfoUpdateOne) SetNillableBirthday(s *string) *UserInfoUpdateOne {
	if s != nil {
		uiuo.SetBirthday(*s)
	}
	return uiuo
}

// ClearBirthday clears the value of the "birthday" field.
func (uiuo *UserInfoUpdateOne) ClearBirthday() *UserInfoUpdateOne {
	uiuo.mutation.ClearBirthday()
	return uiuo
}

// SetMotto sets the "motto" field.
func (uiuo *UserInfoUpdateOne) SetMotto(s string) *UserInfoUpdateOne {
	uiuo.mutation.SetMotto(s)
	return uiuo
}

// SetNillableMotto sets the "motto" field if the given value is not nil.
func (uiuo *UserInfoUpdateOne) SetNillableMotto(s *string) *UserInfoUpdateOne {
	if s != nil {
		uiuo.SetMotto(*s)
	}
	return uiuo
}

// ClearMotto clears the value of the "motto" field.
func (uiuo *UserInfoUpdateOne) ClearMotto() *UserInfoUpdateOne {
	uiuo.mutation.ClearMotto()
	return uiuo
}

// Mutation returns the UserInfoMutation object of the builder.
func (uiuo *UserInfoUpdateOne) Mutation() *UserInfoMutation {
	return uiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uiuo *UserInfoUpdateOne) Select(field string, fields ...string) *UserInfoUpdateOne {
	uiuo.fields = append([]string{field}, fields...)
	return uiuo
}

// Save executes the query and returns the updated UserInfo entity.
func (uiuo *UserInfoUpdateOne) Save(ctx context.Context) (*UserInfo, error) {
	var (
		err  error
		node *UserInfo
	)
	uiuo.defaults()
	if len(uiuo.hooks) == 0 {
		node, err = uiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uiuo.mutation = mutation
			node, err = uiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uiuo.hooks) - 1; i >= 0; i-- {
			if uiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uiuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uiuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (uiuo *UserInfoUpdateOne) SaveX(ctx context.Context) *UserInfo {
	node, err := uiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uiuo *UserInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := uiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uiuo *UserInfoUpdateOne) ExecX(ctx context.Context) {
	if err := uiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uiuo *UserInfoUpdateOne) defaults() {
	if _, ok := uiuo.mutation.UpdatedAt(); !ok {
		v := userinfo.UpdateDefaultUpdatedAt()
		uiuo.mutation.SetUpdatedAt(v)
	}
}

func (uiuo *UserInfoUpdateOne) sqlSave(ctx context.Context) (_node *UserInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userinfo.Table,
			Columns: userinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userinfo.FieldID,
			},
		},
	}
	id, ok := uiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userinfo.FieldID)
		for _, f := range fields {
			if !userinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uiuo.mutation.DeletedAt(); ok {
		_spec.SetField(userinfo.FieldDeletedAt, field.TypeTime, value)
	}
	if uiuo.mutation.DeletedAtCleared() {
		_spec.ClearField(userinfo.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uiuo.mutation.UserID(); ok {
		_spec.SetField(userinfo.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := uiuo.mutation.AddedUserID(); ok {
		_spec.AddField(userinfo.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := uiuo.mutation.Sex(); ok {
		_spec.SetField(userinfo.FieldSex, field.TypeInt8, value)
	}
	if value, ok := uiuo.mutation.AddedSex(); ok {
		_spec.AddField(userinfo.FieldSex, field.TypeInt8, value)
	}
	if value, ok := uiuo.mutation.Age(); ok {
		_spec.SetField(userinfo.FieldAge, field.TypeInt8, value)
	}
	if value, ok := uiuo.mutation.AddedAge(); ok {
		_spec.AddField(userinfo.FieldAge, field.TypeInt8, value)
	}
	if uiuo.mutation.AgeCleared() {
		_spec.ClearField(userinfo.FieldAge, field.TypeInt8)
	}
	if value, ok := uiuo.mutation.Birthday(); ok {
		_spec.SetField(userinfo.FieldBirthday, field.TypeString, value)
	}
	if uiuo.mutation.BirthdayCleared() {
		_spec.ClearField(userinfo.FieldBirthday, field.TypeString)
	}
	if value, ok := uiuo.mutation.Motto(); ok {
		_spec.SetField(userinfo.FieldMotto, field.TypeString, value)
	}
	if uiuo.mutation.MottoCleared() {
		_spec.ClearField(userinfo.FieldMotto, field.TypeString)
	}
	_node = &UserInfo{config: uiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
