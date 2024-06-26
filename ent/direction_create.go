// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"team-manager/ent/direction"
	"team-manager/ent/schedule"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DirectionCreate is the builder for creating a Direction entity.
type DirectionCreate struct {
	config
	mutation *DirectionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStartStationID sets the "start_station_id" field.
func (dc *DirectionCreate) SetStartStationID(i int64) *DirectionCreate {
	dc.mutation.SetStartStationID(i)
	return dc
}

// SetEndStationID sets the "end_station_id" field.
func (dc *DirectionCreate) SetEndStationID(i int64) *DirectionCreate {
	dc.mutation.SetEndStationID(i)
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DirectionCreate) SetCreatedAt(t time.Time) *DirectionCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DirectionCreate) SetNillableCreatedAt(t *time.Time) *DirectionCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DirectionCreate) SetID(i int64) *DirectionCreate {
	dc.mutation.SetID(i)
	return dc
}

// AddDirectionScheduleIDs adds the "direction_schedules" edge to the Schedule entity by IDs.
func (dc *DirectionCreate) AddDirectionScheduleIDs(ids ...int64) *DirectionCreate {
	dc.mutation.AddDirectionScheduleIDs(ids...)
	return dc
}

// AddDirectionSchedules adds the "direction_schedules" edges to the Schedule entity.
func (dc *DirectionCreate) AddDirectionSchedules(s ...*Schedule) *DirectionCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddDirectionScheduleIDs(ids...)
}

// Mutation returns the DirectionMutation object of the builder.
func (dc *DirectionCreate) Mutation() *DirectionMutation {
	return dc.mutation
}

// Save creates the Direction in the database.
func (dc *DirectionCreate) Save(ctx context.Context) (*Direction, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DirectionCreate) SaveX(ctx context.Context) *Direction {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DirectionCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DirectionCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DirectionCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := direction.DefaultCreatedAt
		dc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DirectionCreate) check() error {
	if _, ok := dc.mutation.StartStationID(); !ok {
		return &ValidationError{Name: "start_station_id", err: errors.New(`ent: missing required field "Direction.start_station_id"`)}
	}
	if _, ok := dc.mutation.EndStationID(); !ok {
		return &ValidationError{Name: "end_station_id", err: errors.New(`ent: missing required field "Direction.end_station_id"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Direction.created_at"`)}
	}
	return nil
}

func (dc *DirectionCreate) sqlSave(ctx context.Context) (*Direction, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DirectionCreate) createSpec() (*Direction, *sqlgraph.CreateSpec) {
	var (
		_node = &Direction{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(direction.Table, sqlgraph.NewFieldSpec(direction.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.StartStationID(); ok {
		_spec.SetField(direction.FieldStartStationID, field.TypeInt64, value)
		_node.StartStationID = value
	}
	if value, ok := dc.mutation.EndStationID(); ok {
		_spec.SetField(direction.FieldEndStationID, field.TypeInt64, value)
		_node.EndStationID = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(direction.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := dc.mutation.DirectionSchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   direction.DirectionSchedulesTable,
			Columns: []string{direction.DirectionSchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Direction.Create().
//		SetStartStationID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DirectionUpsert) {
//			SetStartStationID(v+v).
//		}).
//		Exec(ctx)
func (dc *DirectionCreate) OnConflict(opts ...sql.ConflictOption) *DirectionUpsertOne {
	dc.conflict = opts
	return &DirectionUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Direction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DirectionCreate) OnConflictColumns(columns ...string) *DirectionUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DirectionUpsertOne{
		create: dc,
	}
}

type (
	// DirectionUpsertOne is the builder for "upsert"-ing
	//  one Direction node.
	DirectionUpsertOne struct {
		create *DirectionCreate
	}

	// DirectionUpsert is the "OnConflict" setter.
	DirectionUpsert struct {
		*sql.UpdateSet
	}
)

// SetStartStationID sets the "start_station_id" field.
func (u *DirectionUpsert) SetStartStationID(v int64) *DirectionUpsert {
	u.Set(direction.FieldStartStationID, v)
	return u
}

// UpdateStartStationID sets the "start_station_id" field to the value that was provided on create.
func (u *DirectionUpsert) UpdateStartStationID() *DirectionUpsert {
	u.SetExcluded(direction.FieldStartStationID)
	return u
}

// AddStartStationID adds v to the "start_station_id" field.
func (u *DirectionUpsert) AddStartStationID(v int64) *DirectionUpsert {
	u.Add(direction.FieldStartStationID, v)
	return u
}

// SetEndStationID sets the "end_station_id" field.
func (u *DirectionUpsert) SetEndStationID(v int64) *DirectionUpsert {
	u.Set(direction.FieldEndStationID, v)
	return u
}

// UpdateEndStationID sets the "end_station_id" field to the value that was provided on create.
func (u *DirectionUpsert) UpdateEndStationID() *DirectionUpsert {
	u.SetExcluded(direction.FieldEndStationID)
	return u
}

// AddEndStationID adds v to the "end_station_id" field.
func (u *DirectionUpsert) AddEndStationID(v int64) *DirectionUpsert {
	u.Add(direction.FieldEndStationID, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DirectionUpsert) SetCreatedAt(v time.Time) *DirectionUpsert {
	u.Set(direction.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DirectionUpsert) UpdateCreatedAt() *DirectionUpsert {
	u.SetExcluded(direction.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Direction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(direction.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DirectionUpsertOne) UpdateNewValues() *DirectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(direction.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Direction.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DirectionUpsertOne) Ignore() *DirectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DirectionUpsertOne) DoNothing() *DirectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DirectionCreate.OnConflict
// documentation for more info.
func (u *DirectionUpsertOne) Update(set func(*DirectionUpsert)) *DirectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DirectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetStartStationID sets the "start_station_id" field.
func (u *DirectionUpsertOne) SetStartStationID(v int64) *DirectionUpsertOne {
	return u.Update(func(s *DirectionUpsert) {
		s.SetStartStationID(v)
	})
}

// AddStartStationID adds v to the "start_station_id" field.
func (u *DirectionUpsertOne) AddStartStationID(v int64) *DirectionUpsertOne {
	return u.Update(func(s *DirectionUpsert) {
		s.AddStartStationID(v)
	})
}

// UpdateStartStationID sets the "start_station_id" field to the value that was provided on create.
func (u *DirectionUpsertOne) UpdateStartStationID() *DirectionUpsertOne {
	return u.Update(func(s *DirectionUpsert) {
		s.UpdateStartStationID()
	})
}

// SetEndStationID sets the "end_station_id" field.
func (u *DirectionUpsertOne) SetEndStationID(v int64) *DirectionUpsertOne {
	return u.Update(func(s *DirectionUpsert) {
		s.SetEndStationID(v)
	})
}

// AddEndStationID adds v to the "end_station_id" field.
func (u *DirectionUpsertOne) AddEndStationID(v int64) *DirectionUpsertOne {
	return u.Update(func(s *DirectionUpsert) {
		s.AddEndStationID(v)
	})
}

// UpdateEndStationID sets the "end_station_id" field to the value that was provided on create.
func (u *DirectionUpsertOne) UpdateEndStationID() *DirectionUpsertOne {
	return u.Update(func(s *DirectionUpsert) {
		s.UpdateEndStationID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DirectionUpsertOne) SetCreatedAt(v time.Time) *DirectionUpsertOne {
	return u.Update(func(s *DirectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DirectionUpsertOne) UpdateCreatedAt() *DirectionUpsertOne {
	return u.Update(func(s *DirectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *DirectionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DirectionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DirectionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DirectionUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DirectionUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DirectionCreateBulk is the builder for creating many Direction entities in bulk.
type DirectionCreateBulk struct {
	config
	err      error
	builders []*DirectionCreate
	conflict []sql.ConflictOption
}

// Save creates the Direction entities in the database.
func (dcb *DirectionCreateBulk) Save(ctx context.Context) ([]*Direction, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Direction, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DirectionMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DirectionCreateBulk) SaveX(ctx context.Context) []*Direction {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DirectionCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DirectionCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Direction.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DirectionUpsert) {
//			SetStartStationID(v+v).
//		}).
//		Exec(ctx)
func (dcb *DirectionCreateBulk) OnConflict(opts ...sql.ConflictOption) *DirectionUpsertBulk {
	dcb.conflict = opts
	return &DirectionUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Direction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DirectionCreateBulk) OnConflictColumns(columns ...string) *DirectionUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DirectionUpsertBulk{
		create: dcb,
	}
}

// DirectionUpsertBulk is the builder for "upsert"-ing
// a bulk of Direction nodes.
type DirectionUpsertBulk struct {
	create *DirectionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Direction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(direction.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DirectionUpsertBulk) UpdateNewValues() *DirectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(direction.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Direction.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DirectionUpsertBulk) Ignore() *DirectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DirectionUpsertBulk) DoNothing() *DirectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DirectionCreateBulk.OnConflict
// documentation for more info.
func (u *DirectionUpsertBulk) Update(set func(*DirectionUpsert)) *DirectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DirectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetStartStationID sets the "start_station_id" field.
func (u *DirectionUpsertBulk) SetStartStationID(v int64) *DirectionUpsertBulk {
	return u.Update(func(s *DirectionUpsert) {
		s.SetStartStationID(v)
	})
}

// AddStartStationID adds v to the "start_station_id" field.
func (u *DirectionUpsertBulk) AddStartStationID(v int64) *DirectionUpsertBulk {
	return u.Update(func(s *DirectionUpsert) {
		s.AddStartStationID(v)
	})
}

// UpdateStartStationID sets the "start_station_id" field to the value that was provided on create.
func (u *DirectionUpsertBulk) UpdateStartStationID() *DirectionUpsertBulk {
	return u.Update(func(s *DirectionUpsert) {
		s.UpdateStartStationID()
	})
}

// SetEndStationID sets the "end_station_id" field.
func (u *DirectionUpsertBulk) SetEndStationID(v int64) *DirectionUpsertBulk {
	return u.Update(func(s *DirectionUpsert) {
		s.SetEndStationID(v)
	})
}

// AddEndStationID adds v to the "end_station_id" field.
func (u *DirectionUpsertBulk) AddEndStationID(v int64) *DirectionUpsertBulk {
	return u.Update(func(s *DirectionUpsert) {
		s.AddEndStationID(v)
	})
}

// UpdateEndStationID sets the "end_station_id" field to the value that was provided on create.
func (u *DirectionUpsertBulk) UpdateEndStationID() *DirectionUpsertBulk {
	return u.Update(func(s *DirectionUpsert) {
		s.UpdateEndStationID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DirectionUpsertBulk) SetCreatedAt(v time.Time) *DirectionUpsertBulk {
	return u.Update(func(s *DirectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DirectionUpsertBulk) UpdateCreatedAt() *DirectionUpsertBulk {
	return u.Update(func(s *DirectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *DirectionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DirectionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DirectionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DirectionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
