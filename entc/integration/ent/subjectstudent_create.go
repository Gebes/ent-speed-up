// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/student"
	"entgo.io/ent/entc/integration/ent/subject"
	"entgo.io/ent/entc/integration/ent/subjectstudent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SubjectStudentCreate is the builder for creating a SubjectStudent entity.
type SubjectStudentCreate struct {
	config
	mutation *SubjectStudentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNote sets the "note" field.
func (ssc *SubjectStudentCreate) SetNote(s string) *SubjectStudentCreate {
	ssc.mutation.SetNote(s)
	return ssc
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (ssc *SubjectStudentCreate) SetNillableNote(s *string) *SubjectStudentCreate {
	if s != nil {
		ssc.SetNote(*s)
	}
	return ssc
}

// SetSubjectID sets the "subject_id" field.
func (ssc *SubjectStudentCreate) SetSubjectID(u uuid.UUID) *SubjectStudentCreate {
	ssc.mutation.SetSubjectID(u)
	return ssc
}

// SetStudentID sets the "student_id" field.
func (ssc *SubjectStudentCreate) SetStudentID(u uuid.UUID) *SubjectStudentCreate {
	ssc.mutation.SetStudentID(u)
	return ssc
}

// SetID sets the "id" field.
func (ssc *SubjectStudentCreate) SetID(u uuid.UUID) *SubjectStudentCreate {
	ssc.mutation.SetID(u)
	return ssc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ssc *SubjectStudentCreate) SetNillableID(u *uuid.UUID) *SubjectStudentCreate {
	if u != nil {
		ssc.SetID(*u)
	}
	return ssc
}

// SetSubject sets the "subject" edge to the Subject entity.
func (ssc *SubjectStudentCreate) SetSubject(s *Subject) *SubjectStudentCreate {
	return ssc.SetSubjectID(s.ID)
}

// SetStudent sets the "student" edge to the Student entity.
func (ssc *SubjectStudentCreate) SetStudent(s *Student) *SubjectStudentCreate {
	return ssc.SetStudentID(s.ID)
}

// Mutation returns the SubjectStudentMutation object of the builder.
func (ssc *SubjectStudentCreate) Mutation() *SubjectStudentMutation {
	return ssc.mutation
}

// Save creates the SubjectStudent in the database.
func (ssc *SubjectStudentCreate) Save(ctx context.Context) (*SubjectStudent, error) {
	ssc.defaults()
	return withHooks(ctx, ssc.sqlSave, ssc.mutation, ssc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ssc *SubjectStudentCreate) SaveX(ctx context.Context) *SubjectStudent {
	v, err := ssc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ssc *SubjectStudentCreate) Exec(ctx context.Context) error {
	_, err := ssc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssc *SubjectStudentCreate) ExecX(ctx context.Context) {
	if err := ssc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ssc *SubjectStudentCreate) defaults() {
	if _, ok := ssc.mutation.ID(); !ok {
		v := subjectstudent.DefaultID()
		ssc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ssc *SubjectStudentCreate) check() error {
	if _, ok := ssc.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject_id", err: errors.New(`ent: missing required field "SubjectStudent.subject_id"`)}
	}
	if _, ok := ssc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student_id", err: errors.New(`ent: missing required field "SubjectStudent.student_id"`)}
	}
	if len(ssc.mutation.SubjectIDs()) == 0 {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required edge "SubjectStudent.subject"`)}
	}
	if len(ssc.mutation.StudentIDs()) == 0 {
		return &ValidationError{Name: "student", err: errors.New(`ent: missing required edge "SubjectStudent.student"`)}
	}
	return nil
}

func (ssc *SubjectStudentCreate) sqlSave(ctx context.Context) (*SubjectStudent, error) {
	if err := ssc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ssc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ssc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ssc.mutation.id = &_node.ID
	ssc.mutation.done = true
	return _node, nil
}

func (ssc *SubjectStudentCreate) createSpec() (*SubjectStudent, *sqlgraph.CreateSpec) {
	var (
		_node = &SubjectStudent{config: ssc.config}
		_spec = sqlgraph.NewCreateSpec(subjectstudent.Table, sqlgraph.NewFieldSpec(subjectstudent.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ssc.conflict
	if id, ok := ssc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ssc.mutation.Note(); ok {
		_spec.SetField(subjectstudent.FieldNote, field.TypeString, value)
		_node.Note = &value
	}
	if nodes := ssc.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subjectstudent.SubjectTable,
			Columns: []string{subjectstudent.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SubjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ssc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subjectstudent.StudentTable,
			Columns: []string{subjectstudent.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StudentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SubjectStudent.Create().
//		SetNote(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubjectStudentUpsert) {
//			SetNote(v+v).
//		}).
//		Exec(ctx)
func (ssc *SubjectStudentCreate) OnConflict(opts ...sql.ConflictOption) *SubjectStudentUpsertOne {
	ssc.conflict = opts
	return &SubjectStudentUpsertOne{
		create: ssc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SubjectStudent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ssc *SubjectStudentCreate) OnConflictColumns(columns ...string) *SubjectStudentUpsertOne {
	ssc.conflict = append(ssc.conflict, sql.ConflictColumns(columns...))
	return &SubjectStudentUpsertOne{
		create: ssc,
	}
}

type (
	// SubjectStudentUpsertOne is the builder for "upsert"-ing
	//  one SubjectStudent node.
	SubjectStudentUpsertOne struct {
		create *SubjectStudentCreate
	}

	// SubjectStudentUpsert is the "OnConflict" setter.
	SubjectStudentUpsert struct {
		*sql.UpdateSet
	}
)

// SetNote sets the "note" field.
func (u *SubjectStudentUpsert) SetNote(v string) *SubjectStudentUpsert {
	u.Set(subjectstudent.FieldNote, v)
	return u
}

// UpdateNote sets the "note" field to the value that was provided on create.
func (u *SubjectStudentUpsert) UpdateNote() *SubjectStudentUpsert {
	u.SetExcluded(subjectstudent.FieldNote)
	return u
}

// ClearNote clears the value of the "note" field.
func (u *SubjectStudentUpsert) ClearNote() *SubjectStudentUpsert {
	u.SetNull(subjectstudent.FieldNote)
	return u
}

// SetSubjectID sets the "subject_id" field.
func (u *SubjectStudentUpsert) SetSubjectID(v uuid.UUID) *SubjectStudentUpsert {
	u.Set(subjectstudent.FieldSubjectID, v)
	return u
}

// UpdateSubjectID sets the "subject_id" field to the value that was provided on create.
func (u *SubjectStudentUpsert) UpdateSubjectID() *SubjectStudentUpsert {
	u.SetExcluded(subjectstudent.FieldSubjectID)
	return u
}

// SetStudentID sets the "student_id" field.
func (u *SubjectStudentUpsert) SetStudentID(v uuid.UUID) *SubjectStudentUpsert {
	u.Set(subjectstudent.FieldStudentID, v)
	return u
}

// UpdateStudentID sets the "student_id" field to the value that was provided on create.
func (u *SubjectStudentUpsert) UpdateStudentID() *SubjectStudentUpsert {
	u.SetExcluded(subjectstudent.FieldStudentID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SubjectStudent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subjectstudent.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubjectStudentUpsertOne) UpdateNewValues() *SubjectStudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(subjectstudent.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SubjectStudent.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SubjectStudentUpsertOne) Ignore() *SubjectStudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubjectStudentUpsertOne) DoNothing() *SubjectStudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubjectStudentCreate.OnConflict
// documentation for more info.
func (u *SubjectStudentUpsertOne) Update(set func(*SubjectStudentUpsert)) *SubjectStudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubjectStudentUpsert{UpdateSet: update})
	}))
	return u
}

// SetNote sets the "note" field.
func (u *SubjectStudentUpsertOne) SetNote(v string) *SubjectStudentUpsertOne {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.SetNote(v)
	})
}

// UpdateNote sets the "note" field to the value that was provided on create.
func (u *SubjectStudentUpsertOne) UpdateNote() *SubjectStudentUpsertOne {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.UpdateNote()
	})
}

// ClearNote clears the value of the "note" field.
func (u *SubjectStudentUpsertOne) ClearNote() *SubjectStudentUpsertOne {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.ClearNote()
	})
}

// SetSubjectID sets the "subject_id" field.
func (u *SubjectStudentUpsertOne) SetSubjectID(v uuid.UUID) *SubjectStudentUpsertOne {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.SetSubjectID(v)
	})
}

// UpdateSubjectID sets the "subject_id" field to the value that was provided on create.
func (u *SubjectStudentUpsertOne) UpdateSubjectID() *SubjectStudentUpsertOne {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.UpdateSubjectID()
	})
}

// SetStudentID sets the "student_id" field.
func (u *SubjectStudentUpsertOne) SetStudentID(v uuid.UUID) *SubjectStudentUpsertOne {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.SetStudentID(v)
	})
}

// UpdateStudentID sets the "student_id" field to the value that was provided on create.
func (u *SubjectStudentUpsertOne) UpdateStudentID() *SubjectStudentUpsertOne {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.UpdateStudentID()
	})
}

// Exec executes the query.
func (u *SubjectStudentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubjectStudentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubjectStudentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SubjectStudentUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SubjectStudentUpsertOne.ID is not supported by MySQL driver. Use SubjectStudentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SubjectStudentUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SubjectStudentCreateBulk is the builder for creating many SubjectStudent entities in bulk.
type SubjectStudentCreateBulk struct {
	config
	err      error
	builders []*SubjectStudentCreate
	conflict []sql.ConflictOption
}

// Save creates the SubjectStudent entities in the database.
func (sscb *SubjectStudentCreateBulk) Save(ctx context.Context) ([]*SubjectStudent, error) {
	if sscb.err != nil {
		return nil, sscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sscb.builders))
	nodes := make([]*SubjectStudent, len(sscb.builders))
	mutators := make([]Mutator, len(sscb.builders))
	for i := range sscb.builders {
		func(i int, root context.Context) {
			builder := sscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubjectStudentMutation)
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
					_, err = mutators[i+1].Mutate(root, sscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sscb *SubjectStudentCreateBulk) SaveX(ctx context.Context) []*SubjectStudent {
	v, err := sscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sscb *SubjectStudentCreateBulk) Exec(ctx context.Context) error {
	_, err := sscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sscb *SubjectStudentCreateBulk) ExecX(ctx context.Context) {
	if err := sscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SubjectStudent.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubjectStudentUpsert) {
//			SetNote(v+v).
//		}).
//		Exec(ctx)
func (sscb *SubjectStudentCreateBulk) OnConflict(opts ...sql.ConflictOption) *SubjectStudentUpsertBulk {
	sscb.conflict = opts
	return &SubjectStudentUpsertBulk{
		create: sscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SubjectStudent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sscb *SubjectStudentCreateBulk) OnConflictColumns(columns ...string) *SubjectStudentUpsertBulk {
	sscb.conflict = append(sscb.conflict, sql.ConflictColumns(columns...))
	return &SubjectStudentUpsertBulk{
		create: sscb,
	}
}

// SubjectStudentUpsertBulk is the builder for "upsert"-ing
// a bulk of SubjectStudent nodes.
type SubjectStudentUpsertBulk struct {
	create *SubjectStudentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SubjectStudent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subjectstudent.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubjectStudentUpsertBulk) UpdateNewValues() *SubjectStudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(subjectstudent.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SubjectStudent.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SubjectStudentUpsertBulk) Ignore() *SubjectStudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubjectStudentUpsertBulk) DoNothing() *SubjectStudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubjectStudentCreateBulk.OnConflict
// documentation for more info.
func (u *SubjectStudentUpsertBulk) Update(set func(*SubjectStudentUpsert)) *SubjectStudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubjectStudentUpsert{UpdateSet: update})
	}))
	return u
}

// SetNote sets the "note" field.
func (u *SubjectStudentUpsertBulk) SetNote(v string) *SubjectStudentUpsertBulk {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.SetNote(v)
	})
}

// UpdateNote sets the "note" field to the value that was provided on create.
func (u *SubjectStudentUpsertBulk) UpdateNote() *SubjectStudentUpsertBulk {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.UpdateNote()
	})
}

// ClearNote clears the value of the "note" field.
func (u *SubjectStudentUpsertBulk) ClearNote() *SubjectStudentUpsertBulk {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.ClearNote()
	})
}

// SetSubjectID sets the "subject_id" field.
func (u *SubjectStudentUpsertBulk) SetSubjectID(v uuid.UUID) *SubjectStudentUpsertBulk {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.SetSubjectID(v)
	})
}

// UpdateSubjectID sets the "subject_id" field to the value that was provided on create.
func (u *SubjectStudentUpsertBulk) UpdateSubjectID() *SubjectStudentUpsertBulk {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.UpdateSubjectID()
	})
}

// SetStudentID sets the "student_id" field.
func (u *SubjectStudentUpsertBulk) SetStudentID(v uuid.UUID) *SubjectStudentUpsertBulk {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.SetStudentID(v)
	})
}

// UpdateStudentID sets the "student_id" field to the value that was provided on create.
func (u *SubjectStudentUpsertBulk) UpdateStudentID() *SubjectStudentUpsertBulk {
	return u.Update(func(s *SubjectStudentUpsert) {
		s.UpdateStudentID()
	})
}

// Exec executes the query.
func (u *SubjectStudentUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SubjectStudentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubjectStudentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubjectStudentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}