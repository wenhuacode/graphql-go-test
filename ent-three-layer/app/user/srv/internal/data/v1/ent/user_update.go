// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent-three-layer/app/user/srv/internal/data/v1/ent/predicate"
	"ent-three-layer/app/user/srv/internal/data/v1/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetCreatedBy sets the "created_by" field.
func (uu *UserUpdate) SetCreatedBy(s string) *UserUpdate {
	uu.mutation.SetCreatedBy(s)
	return uu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedBy(s *string) *UserUpdate {
	if s != nil {
		uu.SetCreatedBy(*s)
	}
	return uu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (uu *UserUpdate) ClearCreatedBy() *UserUpdate {
	uu.mutation.ClearCreatedBy()
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// SetUpdatedBy sets the "updated_by" field.
func (uu *UserUpdate) SetUpdatedBy(s string) *UserUpdate {
	uu.mutation.SetUpdatedBy(s)
	return uu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedBy(s *string) *UserUpdate {
	if s != nil {
		uu.SetUpdatedBy(*s)
	}
	return uu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uu *UserUpdate) ClearUpdatedBy() *UserUpdate {
	uu.mutation.ClearUpdatedBy()
	return uu
}

// SetIsDeleted sets the "is_deleted" field.
func (uu *UserUpdate) SetIsDeleted(b bool) *UserUpdate {
	uu.mutation.SetIsDeleted(b)
	return uu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsDeleted(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsDeleted(*b)
	}
	return uu
}

// ClearIsDeleted clears the value of the "is_deleted" field.
func (uu *UserUpdate) ClearIsDeleted() *UserUpdate {
	uu.mutation.ClearIsDeleted()
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// SetDeletedBy sets the "deleted_by" field.
func (uu *UserUpdate) SetDeletedBy(s string) *UserUpdate {
	uu.mutation.SetDeletedBy(s)
	return uu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedBy(s *string) *UserUpdate {
	if s != nil {
		uu.SetDeletedBy(*s)
	}
	return uu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (uu *UserUpdate) ClearDeletedBy() *UserUpdate {
	uu.mutation.ClearDeletedBy()
	return uu
}

// SetMobile sets the "mobile" field.
func (uu *UserUpdate) SetMobile(s string) *UserUpdate {
	uu.mutation.SetMobile(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickname(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickname(*s)
	}
	return uu
}

// ClearNickname clears the value of the "nickname" field.
func (uu *UserUpdate) ClearNickname() *UserUpdate {
	uu.mutation.ClearNickname()
	return uu
}

// SetBirthday sets the "birthday" field.
func (uu *UserUpdate) SetBirthday(t time.Time) *UserUpdate {
	uu.mutation.SetBirthday(t)
	return uu
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBirthday(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetBirthday(*t)
	}
	return uu
}

// ClearBirthday clears the value of the "birthday" field.
func (uu *UserUpdate) ClearBirthday() *UserUpdate {
	uu.mutation.ClearBirthday()
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(s string) *UserUpdate {
	uu.mutation.SetGender(s)
	return uu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGender(s *string) *UserUpdate {
	if s != nil {
		uu.SetGender(*s)
	}
	return uu
}

// SetRole sets the "Role" field.
func (uu *UserUpdate) SetRole(i int) *UserUpdate {
	uu.mutation.ResetRole()
	uu.mutation.SetRole(i)
	return uu
}

// SetNillableRole sets the "Role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(i *int) *UserUpdate {
	if i != nil {
		uu.SetRole(*i)
	}
	return uu
}

// AddRole adds i to the "Role" field.
func (uu *UserUpdate) AddRole(i int) *UserUpdate {
	uu.mutation.AddRole(i)
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeString, value)
	}
	if uu.mutation.CreatedByCleared() {
		_spec.ClearField(user.FieldCreatedBy, field.TypeString)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeString, value)
	}
	if uu.mutation.UpdatedByCleared() {
		_spec.ClearField(user.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := uu.mutation.IsDeleted(); ok {
		_spec.SetField(user.FieldIsDeleted, field.TypeBool, value)
	}
	if uu.mutation.IsDeletedCleared() {
		_spec.ClearField(user.FieldIsDeleted, field.TypeBool)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.DeletedBy(); ok {
		_spec.SetField(user.FieldDeletedBy, field.TypeString, value)
	}
	if uu.mutation.DeletedByCleared() {
		_spec.ClearField(user.FieldDeletedBy, field.TypeString)
	}
	if value, ok := uu.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uu.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uu.mutation.Birthday(); ok {
		_spec.SetField(user.FieldBirthday, field.TypeTime, value)
	}
	if uu.mutation.BirthdayCleared() {
		_spec.ClearField(user.FieldBirthday, field.TypeTime)
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeString, value)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetCreatedBy sets the "created_by" field.
func (uuo *UserUpdateOne) SetCreatedBy(s string) *UserUpdateOne {
	uuo.mutation.SetCreatedBy(s)
	return uuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedBy(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCreatedBy(*s)
	}
	return uuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (uuo *UserUpdateOne) ClearCreatedBy() *UserUpdateOne {
	uuo.mutation.ClearCreatedBy()
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// SetUpdatedBy sets the "updated_by" field.
func (uuo *UserUpdateOne) SetUpdatedBy(s string) *UserUpdateOne {
	uuo.mutation.SetUpdatedBy(s)
	return uuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedBy(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUpdatedBy(*s)
	}
	return uuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uuo *UserUpdateOne) ClearUpdatedBy() *UserUpdateOne {
	uuo.mutation.ClearUpdatedBy()
	return uuo
}

// SetIsDeleted sets the "is_deleted" field.
func (uuo *UserUpdateOne) SetIsDeleted(b bool) *UserUpdateOne {
	uuo.mutation.SetIsDeleted(b)
	return uuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsDeleted(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsDeleted(*b)
	}
	return uuo
}

// ClearIsDeleted clears the value of the "is_deleted" field.
func (uuo *UserUpdateOne) ClearIsDeleted() *UserUpdateOne {
	uuo.mutation.ClearIsDeleted()
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// SetDeletedBy sets the "deleted_by" field.
func (uuo *UserUpdateOne) SetDeletedBy(s string) *UserUpdateOne {
	uuo.mutation.SetDeletedBy(s)
	return uuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedBy(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDeletedBy(*s)
	}
	return uuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (uuo *UserUpdateOne) ClearDeletedBy() *UserUpdateOne {
	uuo.mutation.ClearDeletedBy()
	return uuo
}

// SetMobile sets the "mobile" field.
func (uuo *UserUpdateOne) SetMobile(s string) *UserUpdateOne {
	uuo.mutation.SetMobile(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickname(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickname(*s)
	}
	return uuo
}

// ClearNickname clears the value of the "nickname" field.
func (uuo *UserUpdateOne) ClearNickname() *UserUpdateOne {
	uuo.mutation.ClearNickname()
	return uuo
}

// SetBirthday sets the "birthday" field.
func (uuo *UserUpdateOne) SetBirthday(t time.Time) *UserUpdateOne {
	uuo.mutation.SetBirthday(t)
	return uuo
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBirthday(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetBirthday(*t)
	}
	return uuo
}

// ClearBirthday clears the value of the "birthday" field.
func (uuo *UserUpdateOne) ClearBirthday() *UserUpdateOne {
	uuo.mutation.ClearBirthday()
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(s string) *UserUpdateOne {
	uuo.mutation.SetGender(s)
	return uuo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGender(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetGender(*s)
	}
	return uuo
}

// SetRole sets the "Role" field.
func (uuo *UserUpdateOne) SetRole(i int) *UserUpdateOne {
	uuo.mutation.ResetRole()
	uuo.mutation.SetRole(i)
	return uuo
}

// SetNillableRole sets the "Role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetRole(*i)
	}
	return uuo
}

// AddRole adds i to the "Role" field.
func (uuo *UserUpdateOne) AddRole(i int) *UserUpdateOne {
	uuo.mutation.AddRole(i)
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeString, value)
	}
	if uuo.mutation.CreatedByCleared() {
		_spec.ClearField(user.FieldCreatedBy, field.TypeString)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeString, value)
	}
	if uuo.mutation.UpdatedByCleared() {
		_spec.ClearField(user.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := uuo.mutation.IsDeleted(); ok {
		_spec.SetField(user.FieldIsDeleted, field.TypeBool, value)
	}
	if uuo.mutation.IsDeletedCleared() {
		_spec.ClearField(user.FieldIsDeleted, field.TypeBool)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.DeletedBy(); ok {
		_spec.SetField(user.FieldDeletedBy, field.TypeString, value)
	}
	if uuo.mutation.DeletedByCleared() {
		_spec.ClearField(user.FieldDeletedBy, field.TypeString)
	}
	if value, ok := uuo.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uuo.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uuo.mutation.Birthday(); ok {
		_spec.SetField(user.FieldBirthday, field.TypeTime, value)
	}
	if uuo.mutation.BirthdayCleared() {
		_spec.ClearField(user.FieldBirthday, field.TypeTime)
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeInt, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
