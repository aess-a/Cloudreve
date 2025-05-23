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
	"github.com/cloudreve/Cloudreve/v4/ent/davaccount"
	"github.com/cloudreve/Cloudreve/v4/ent/user"
	"github.com/cloudreve/Cloudreve/v4/inventory/types"
	"github.com/cloudreve/Cloudreve/v4/pkg/boolset"
)

// DavAccountCreate is the builder for creating a DavAccount entity.
type DavAccountCreate struct {
	config
	mutation *DavAccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dac *DavAccountCreate) SetCreatedAt(t time.Time) *DavAccountCreate {
	dac.mutation.SetCreatedAt(t)
	return dac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dac *DavAccountCreate) SetNillableCreatedAt(t *time.Time) *DavAccountCreate {
	if t != nil {
		dac.SetCreatedAt(*t)
	}
	return dac
}

// SetUpdatedAt sets the "updated_at" field.
func (dac *DavAccountCreate) SetUpdatedAt(t time.Time) *DavAccountCreate {
	dac.mutation.SetUpdatedAt(t)
	return dac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dac *DavAccountCreate) SetNillableUpdatedAt(t *time.Time) *DavAccountCreate {
	if t != nil {
		dac.SetUpdatedAt(*t)
	}
	return dac
}

// SetDeletedAt sets the "deleted_at" field.
func (dac *DavAccountCreate) SetDeletedAt(t time.Time) *DavAccountCreate {
	dac.mutation.SetDeletedAt(t)
	return dac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dac *DavAccountCreate) SetNillableDeletedAt(t *time.Time) *DavAccountCreate {
	if t != nil {
		dac.SetDeletedAt(*t)
	}
	return dac
}

// SetName sets the "name" field.
func (dac *DavAccountCreate) SetName(s string) *DavAccountCreate {
	dac.mutation.SetName(s)
	return dac
}

// SetURI sets the "uri" field.
func (dac *DavAccountCreate) SetURI(s string) *DavAccountCreate {
	dac.mutation.SetURI(s)
	return dac
}

// SetPassword sets the "password" field.
func (dac *DavAccountCreate) SetPassword(s string) *DavAccountCreate {
	dac.mutation.SetPassword(s)
	return dac
}

// SetOptions sets the "options" field.
func (dac *DavAccountCreate) SetOptions(bs *boolset.BooleanSet) *DavAccountCreate {
	dac.mutation.SetOptions(bs)
	return dac
}

// SetProps sets the "props" field.
func (dac *DavAccountCreate) SetProps(tap *types.DavAccountProps) *DavAccountCreate {
	dac.mutation.SetProps(tap)
	return dac
}

// SetOwnerID sets the "owner_id" field.
func (dac *DavAccountCreate) SetOwnerID(i int) *DavAccountCreate {
	dac.mutation.SetOwnerID(i)
	return dac
}

// SetOwner sets the "owner" edge to the User entity.
func (dac *DavAccountCreate) SetOwner(u *User) *DavAccountCreate {
	return dac.SetOwnerID(u.ID)
}

// Mutation returns the DavAccountMutation object of the builder.
func (dac *DavAccountCreate) Mutation() *DavAccountMutation {
	return dac.mutation
}

// Save creates the DavAccount in the database.
func (dac *DavAccountCreate) Save(ctx context.Context) (*DavAccount, error) {
	if err := dac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dac.sqlSave, dac.mutation, dac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dac *DavAccountCreate) SaveX(ctx context.Context) *DavAccount {
	v, err := dac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dac *DavAccountCreate) Exec(ctx context.Context) error {
	_, err := dac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dac *DavAccountCreate) ExecX(ctx context.Context) {
	if err := dac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dac *DavAccountCreate) defaults() error {
	if _, ok := dac.mutation.CreatedAt(); !ok {
		if davaccount.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized davaccount.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := davaccount.DefaultCreatedAt()
		dac.mutation.SetCreatedAt(v)
	}
	if _, ok := dac.mutation.UpdatedAt(); !ok {
		if davaccount.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized davaccount.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := davaccount.DefaultUpdatedAt()
		dac.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dac *DavAccountCreate) check() error {
	if _, ok := dac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DavAccount.created_at"`)}
	}
	if _, ok := dac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DavAccount.updated_at"`)}
	}
	if _, ok := dac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DavAccount.name"`)}
	}
	if _, ok := dac.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "DavAccount.uri"`)}
	}
	if _, ok := dac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "DavAccount.password"`)}
	}
	if _, ok := dac.mutation.Options(); !ok {
		return &ValidationError{Name: "options", err: errors.New(`ent: missing required field "DavAccount.options"`)}
	}
	if _, ok := dac.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "DavAccount.owner_id"`)}
	}
	if _, ok := dac.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "DavAccount.owner"`)}
	}
	return nil
}

func (dac *DavAccountCreate) sqlSave(ctx context.Context) (*DavAccount, error) {
	if err := dac.check(); err != nil {
		return nil, err
	}
	_node, _spec := dac.createSpec()
	if err := sqlgraph.CreateNode(ctx, dac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dac.mutation.id = &_node.ID
	dac.mutation.done = true
	return _node, nil
}

func (dac *DavAccountCreate) createSpec() (*DavAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &DavAccount{config: dac.config}
		_spec = sqlgraph.NewCreateSpec(davaccount.Table, sqlgraph.NewFieldSpec(davaccount.FieldID, field.TypeInt))
	)

	if id, ok := dac.mutation.ID(); ok {
		_node.ID = id
		id64 := int64(id)
		_spec.ID.Value = id64
	}

	_spec.OnConflict = dac.conflict
	if value, ok := dac.mutation.CreatedAt(); ok {
		_spec.SetField(davaccount.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dac.mutation.UpdatedAt(); ok {
		_spec.SetField(davaccount.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dac.mutation.DeletedAt(); ok {
		_spec.SetField(davaccount.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := dac.mutation.Name(); ok {
		_spec.SetField(davaccount.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dac.mutation.URI(); ok {
		_spec.SetField(davaccount.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if value, ok := dac.mutation.Password(); ok {
		_spec.SetField(davaccount.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := dac.mutation.Options(); ok {
		_spec.SetField(davaccount.FieldOptions, field.TypeBytes, value)
		_node.Options = value
	}
	if value, ok := dac.mutation.Props(); ok {
		_spec.SetField(davaccount.FieldProps, field.TypeJSON, value)
		_node.Props = value
	}
	if nodes := dac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   davaccount.OwnerTable,
			Columns: []string{davaccount.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DavAccount.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DavAccountUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dac *DavAccountCreate) OnConflict(opts ...sql.ConflictOption) *DavAccountUpsertOne {
	dac.conflict = opts
	return &DavAccountUpsertOne{
		create: dac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DavAccount.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dac *DavAccountCreate) OnConflictColumns(columns ...string) *DavAccountUpsertOne {
	dac.conflict = append(dac.conflict, sql.ConflictColumns(columns...))
	return &DavAccountUpsertOne{
		create: dac,
	}
}

type (
	// DavAccountUpsertOne is the builder for "upsert"-ing
	//  one DavAccount node.
	DavAccountUpsertOne struct {
		create *DavAccountCreate
	}

	// DavAccountUpsert is the "OnConflict" setter.
	DavAccountUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *DavAccountUpsert) SetUpdatedAt(v time.Time) *DavAccountUpsert {
	u.Set(davaccount.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DavAccountUpsert) UpdateUpdatedAt() *DavAccountUpsert {
	u.SetExcluded(davaccount.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DavAccountUpsert) SetDeletedAt(v time.Time) *DavAccountUpsert {
	u.Set(davaccount.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DavAccountUpsert) UpdateDeletedAt() *DavAccountUpsert {
	u.SetExcluded(davaccount.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DavAccountUpsert) ClearDeletedAt() *DavAccountUpsert {
	u.SetNull(davaccount.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *DavAccountUpsert) SetName(v string) *DavAccountUpsert {
	u.Set(davaccount.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DavAccountUpsert) UpdateName() *DavAccountUpsert {
	u.SetExcluded(davaccount.FieldName)
	return u
}

// SetURI sets the "uri" field.
func (u *DavAccountUpsert) SetURI(v string) *DavAccountUpsert {
	u.Set(davaccount.FieldURI, v)
	return u
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *DavAccountUpsert) UpdateURI() *DavAccountUpsert {
	u.SetExcluded(davaccount.FieldURI)
	return u
}

// SetPassword sets the "password" field.
func (u *DavAccountUpsert) SetPassword(v string) *DavAccountUpsert {
	u.Set(davaccount.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *DavAccountUpsert) UpdatePassword() *DavAccountUpsert {
	u.SetExcluded(davaccount.FieldPassword)
	return u
}

// SetOptions sets the "options" field.
func (u *DavAccountUpsert) SetOptions(v *boolset.BooleanSet) *DavAccountUpsert {
	u.Set(davaccount.FieldOptions, v)
	return u
}

// UpdateOptions sets the "options" field to the value that was provided on create.
func (u *DavAccountUpsert) UpdateOptions() *DavAccountUpsert {
	u.SetExcluded(davaccount.FieldOptions)
	return u
}

// SetProps sets the "props" field.
func (u *DavAccountUpsert) SetProps(v *types.DavAccountProps) *DavAccountUpsert {
	u.Set(davaccount.FieldProps, v)
	return u
}

// UpdateProps sets the "props" field to the value that was provided on create.
func (u *DavAccountUpsert) UpdateProps() *DavAccountUpsert {
	u.SetExcluded(davaccount.FieldProps)
	return u
}

// ClearProps clears the value of the "props" field.
func (u *DavAccountUpsert) ClearProps() *DavAccountUpsert {
	u.SetNull(davaccount.FieldProps)
	return u
}

// SetOwnerID sets the "owner_id" field.
func (u *DavAccountUpsert) SetOwnerID(v int) *DavAccountUpsert {
	u.Set(davaccount.FieldOwnerID, v)
	return u
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *DavAccountUpsert) UpdateOwnerID() *DavAccountUpsert {
	u.SetExcluded(davaccount.FieldOwnerID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DavAccount.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DavAccountUpsertOne) UpdateNewValues() *DavAccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(davaccount.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DavAccount.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DavAccountUpsertOne) Ignore() *DavAccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DavAccountUpsertOne) DoNothing() *DavAccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DavAccountCreate.OnConflict
// documentation for more info.
func (u *DavAccountUpsertOne) Update(set func(*DavAccountUpsert)) *DavAccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DavAccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DavAccountUpsertOne) SetUpdatedAt(v time.Time) *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DavAccountUpsertOne) UpdateUpdatedAt() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DavAccountUpsertOne) SetDeletedAt(v time.Time) *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DavAccountUpsertOne) UpdateDeletedAt() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DavAccountUpsertOne) ClearDeletedAt() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *DavAccountUpsertOne) SetName(v string) *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DavAccountUpsertOne) UpdateName() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateName()
	})
}

// SetURI sets the "uri" field.
func (u *DavAccountUpsertOne) SetURI(v string) *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *DavAccountUpsertOne) UpdateURI() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateURI()
	})
}

// SetPassword sets the "password" field.
func (u *DavAccountUpsertOne) SetPassword(v string) *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *DavAccountUpsertOne) UpdatePassword() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdatePassword()
	})
}

// SetOptions sets the "options" field.
func (u *DavAccountUpsertOne) SetOptions(v *boolset.BooleanSet) *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetOptions(v)
	})
}

// UpdateOptions sets the "options" field to the value that was provided on create.
func (u *DavAccountUpsertOne) UpdateOptions() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateOptions()
	})
}

// SetProps sets the "props" field.
func (u *DavAccountUpsertOne) SetProps(v *types.DavAccountProps) *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetProps(v)
	})
}

// UpdateProps sets the "props" field to the value that was provided on create.
func (u *DavAccountUpsertOne) UpdateProps() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateProps()
	})
}

// ClearProps clears the value of the "props" field.
func (u *DavAccountUpsertOne) ClearProps() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.ClearProps()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *DavAccountUpsertOne) SetOwnerID(v int) *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *DavAccountUpsertOne) UpdateOwnerID() *DavAccountUpsertOne {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateOwnerID()
	})
}

// Exec executes the query.
func (u *DavAccountUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DavAccountCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DavAccountUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DavAccountUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DavAccountUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (m *DavAccountCreate) SetRawID(t int) *DavAccountCreate {
	m.mutation.SetRawID(t)
	return m
}

// DavAccountCreateBulk is the builder for creating many DavAccount entities in bulk.
type DavAccountCreateBulk struct {
	config
	err      error
	builders []*DavAccountCreate
	conflict []sql.ConflictOption
}

// Save creates the DavAccount entities in the database.
func (dacb *DavAccountCreateBulk) Save(ctx context.Context) ([]*DavAccount, error) {
	if dacb.err != nil {
		return nil, dacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dacb.builders))
	nodes := make([]*DavAccount, len(dacb.builders))
	mutators := make([]Mutator, len(dacb.builders))
	for i := range dacb.builders {
		func(i int, root context.Context) {
			builder := dacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DavAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, dacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, dacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dacb *DavAccountCreateBulk) SaveX(ctx context.Context) []*DavAccount {
	v, err := dacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dacb *DavAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := dacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dacb *DavAccountCreateBulk) ExecX(ctx context.Context) {
	if err := dacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DavAccount.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DavAccountUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dacb *DavAccountCreateBulk) OnConflict(opts ...sql.ConflictOption) *DavAccountUpsertBulk {
	dacb.conflict = opts
	return &DavAccountUpsertBulk{
		create: dacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DavAccount.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dacb *DavAccountCreateBulk) OnConflictColumns(columns ...string) *DavAccountUpsertBulk {
	dacb.conflict = append(dacb.conflict, sql.ConflictColumns(columns...))
	return &DavAccountUpsertBulk{
		create: dacb,
	}
}

// DavAccountUpsertBulk is the builder for "upsert"-ing
// a bulk of DavAccount nodes.
type DavAccountUpsertBulk struct {
	create *DavAccountCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DavAccount.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DavAccountUpsertBulk) UpdateNewValues() *DavAccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(davaccount.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DavAccount.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DavAccountUpsertBulk) Ignore() *DavAccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DavAccountUpsertBulk) DoNothing() *DavAccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DavAccountCreateBulk.OnConflict
// documentation for more info.
func (u *DavAccountUpsertBulk) Update(set func(*DavAccountUpsert)) *DavAccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DavAccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DavAccountUpsertBulk) SetUpdatedAt(v time.Time) *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DavAccountUpsertBulk) UpdateUpdatedAt() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DavAccountUpsertBulk) SetDeletedAt(v time.Time) *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DavAccountUpsertBulk) UpdateDeletedAt() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DavAccountUpsertBulk) ClearDeletedAt() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *DavAccountUpsertBulk) SetName(v string) *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DavAccountUpsertBulk) UpdateName() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateName()
	})
}

// SetURI sets the "uri" field.
func (u *DavAccountUpsertBulk) SetURI(v string) *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *DavAccountUpsertBulk) UpdateURI() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateURI()
	})
}

// SetPassword sets the "password" field.
func (u *DavAccountUpsertBulk) SetPassword(v string) *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *DavAccountUpsertBulk) UpdatePassword() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdatePassword()
	})
}

// SetOptions sets the "options" field.
func (u *DavAccountUpsertBulk) SetOptions(v *boolset.BooleanSet) *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetOptions(v)
	})
}

// UpdateOptions sets the "options" field to the value that was provided on create.
func (u *DavAccountUpsertBulk) UpdateOptions() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateOptions()
	})
}

// SetProps sets the "props" field.
func (u *DavAccountUpsertBulk) SetProps(v *types.DavAccountProps) *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetProps(v)
	})
}

// UpdateProps sets the "props" field to the value that was provided on create.
func (u *DavAccountUpsertBulk) UpdateProps() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateProps()
	})
}

// ClearProps clears the value of the "props" field.
func (u *DavAccountUpsertBulk) ClearProps() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.ClearProps()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *DavAccountUpsertBulk) SetOwnerID(v int) *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *DavAccountUpsertBulk) UpdateOwnerID() *DavAccountUpsertBulk {
	return u.Update(func(s *DavAccountUpsert) {
		s.UpdateOwnerID()
	})
}

// Exec executes the query.
func (u *DavAccountUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DavAccountCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DavAccountCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DavAccountUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
