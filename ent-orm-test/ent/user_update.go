// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent-orm-test/ent/pet"
	"ent-orm-test/ent/predicate"
	"ent-orm-test/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
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

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// AddAge adds i to the "age" field.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (uu *UserUpdate) SetSpouseID(id xid.ID) *UserUpdate {
	uu.mutation.SetSpouseID(id)
	return uu
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableSpouseID(id *xid.ID) *UserUpdate {
	if id != nil {
		uu = uu.SetSpouseID(*id)
	}
	return uu
}

// SetSpouse sets the "spouse" edge to the User entity.
func (uu *UserUpdate) SetSpouse(u *User) *UserUpdate {
	return uu.SetSpouseID(u.ID)
}

// SetPrevID sets the "prev" edge to the User entity by ID.
func (uu *UserUpdate) SetPrevID(id xid.ID) *UserUpdate {
	uu.mutation.SetPrevID(id)
	return uu
}

// SetNillablePrevID sets the "prev" edge to the User entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillablePrevID(id *xid.ID) *UserUpdate {
	if id != nil {
		uu = uu.SetPrevID(*id)
	}
	return uu
}

// SetPrev sets the "prev" edge to the User entity.
func (uu *UserUpdate) SetPrev(u *User) *UserUpdate {
	return uu.SetPrevID(u.ID)
}

// SetNextID sets the "next" edge to the User entity by ID.
func (uu *UserUpdate) SetNextID(id xid.ID) *UserUpdate {
	uu.mutation.SetNextID(id)
	return uu
}

// SetNillableNextID sets the "next" edge to the User entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableNextID(id *xid.ID) *UserUpdate {
	if id != nil {
		uu = uu.SetNextID(*id)
	}
	return uu
}

// SetNext sets the "next" edge to the User entity.
func (uu *UserUpdate) SetNext(u *User) *UserUpdate {
	return uu.SetNextID(u.ID)
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uu *UserUpdate) AddPetIDs(ids ...xid.ID) *UserUpdate {
	uu.mutation.AddPetIDs(ids...)
	return uu
}

// AddPets adds the "pets" edges to the Pet entity.
func (uu *UserUpdate) AddPets(p ...*Pet) *UserUpdate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPetIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearSpouse clears the "spouse" edge to the User entity.
func (uu *UserUpdate) ClearSpouse() *UserUpdate {
	uu.mutation.ClearSpouse()
	return uu
}

// ClearPrev clears the "prev" edge to the User entity.
func (uu *UserUpdate) ClearPrev() *UserUpdate {
	uu.mutation.ClearPrev()
	return uu
}

// ClearNext clears the "next" edge to the User entity.
func (uu *UserUpdate) ClearNext() *UserUpdate {
	uu.mutation.ClearNext()
	return uu
}

// ClearPets clears all "pets" edges to the Pet entity.
func (uu *UserUpdate) ClearPets() *UserUpdate {
	uu.mutation.ClearPets()
	return uu
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (uu *UserUpdate) RemovePetIDs(ids ...xid.ID) *UserUpdate {
	uu.mutation.RemovePetIDs(ids...)
	return uu
}

// RemovePets removes "pets" edges to Pet entities.
func (uu *UserUpdate) RemovePets(p ...*Pet) *UserUpdate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePetIDs(ids...)
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
	if v, ok := uu.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "User.age": %w`, err)}
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
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if uu.mutation.SpouseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PrevCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.PrevTable,
			Columns: []string{user.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.PrevTable,
			Columns: []string{user.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.NextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.NextTable,
			Columns: []string{user.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.NextTable,
			Columns: []string{user.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPetsIDs(); len(nodes) > 0 && !uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// AddAge adds i to the "age" field.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (uuo *UserUpdateOne) SetSpouseID(id xid.ID) *UserUpdateOne {
	uuo.mutation.SetSpouseID(id)
	return uuo
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSpouseID(id *xid.ID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSpouseID(*id)
	}
	return uuo
}

// SetSpouse sets the "spouse" edge to the User entity.
func (uuo *UserUpdateOne) SetSpouse(u *User) *UserUpdateOne {
	return uuo.SetSpouseID(u.ID)
}

// SetPrevID sets the "prev" edge to the User entity by ID.
func (uuo *UserUpdateOne) SetPrevID(id xid.ID) *UserUpdateOne {
	uuo.mutation.SetPrevID(id)
	return uuo
}

// SetNillablePrevID sets the "prev" edge to the User entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePrevID(id *xid.ID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetPrevID(*id)
	}
	return uuo
}

// SetPrev sets the "prev" edge to the User entity.
func (uuo *UserUpdateOne) SetPrev(u *User) *UserUpdateOne {
	return uuo.SetPrevID(u.ID)
}

// SetNextID sets the "next" edge to the User entity by ID.
func (uuo *UserUpdateOne) SetNextID(id xid.ID) *UserUpdateOne {
	uuo.mutation.SetNextID(id)
	return uuo
}

// SetNillableNextID sets the "next" edge to the User entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNextID(id *xid.ID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetNextID(*id)
	}
	return uuo
}

// SetNext sets the "next" edge to the User entity.
func (uuo *UserUpdateOne) SetNext(u *User) *UserUpdateOne {
	return uuo.SetNextID(u.ID)
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uuo *UserUpdateOne) AddPetIDs(ids ...xid.ID) *UserUpdateOne {
	uuo.mutation.AddPetIDs(ids...)
	return uuo
}

// AddPets adds the "pets" edges to the Pet entity.
func (uuo *UserUpdateOne) AddPets(p ...*Pet) *UserUpdateOne {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPetIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearSpouse clears the "spouse" edge to the User entity.
func (uuo *UserUpdateOne) ClearSpouse() *UserUpdateOne {
	uuo.mutation.ClearSpouse()
	return uuo
}

// ClearPrev clears the "prev" edge to the User entity.
func (uuo *UserUpdateOne) ClearPrev() *UserUpdateOne {
	uuo.mutation.ClearPrev()
	return uuo
}

// ClearNext clears the "next" edge to the User entity.
func (uuo *UserUpdateOne) ClearNext() *UserUpdateOne {
	uuo.mutation.ClearNext()
	return uuo
}

// ClearPets clears all "pets" edges to the Pet entity.
func (uuo *UserUpdateOne) ClearPets() *UserUpdateOne {
	uuo.mutation.ClearPets()
	return uuo
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (uuo *UserUpdateOne) RemovePetIDs(ids ...xid.ID) *UserUpdateOne {
	uuo.mutation.RemovePetIDs(ids...)
	return uuo
}

// RemovePets removes "pets" edges to Pet entities.
func (uuo *UserUpdateOne) RemovePets(p ...*Pet) *UserUpdateOne {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePetIDs(ids...)
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
	if v, ok := uuo.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "User.age": %w`, err)}
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
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if uuo.mutation.SpouseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PrevCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.PrevTable,
			Columns: []string{user.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.PrevTable,
			Columns: []string{user.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.NextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.NextTable,
			Columns: []string{user.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.NextTable,
			Columns: []string{user.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPetsIDs(); len(nodes) > 0 && !uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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
