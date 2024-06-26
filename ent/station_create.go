// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"team-manager/ent/schedule"
	"team-manager/ent/station"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StationCreate is the builder for creating a Station entity.
type StationCreate struct {
	config
	mutation *StationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (sc *StationCreate) SetName(s string) *StationCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sc *StationCreate) SetNillableName(s *string) *StationCreate {
	if s != nil {
		sc.SetName(*s)
	}
	return sc
}

// SetLatitude sets the "latitude" field.
func (sc *StationCreate) SetLatitude(f float32) *StationCreate {
	sc.mutation.SetLatitude(f)
	return sc
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (sc *StationCreate) SetNillableLatitude(f *float32) *StationCreate {
	if f != nil {
		sc.SetLatitude(*f)
	}
	return sc
}

// SetLongitude sets the "longitude" field.
func (sc *StationCreate) SetLongitude(f float32) *StationCreate {
	sc.mutation.SetLongitude(f)
	return sc
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (sc *StationCreate) SetNillableLongitude(f *float32) *StationCreate {
	if f != nil {
		sc.SetLongitude(*f)
	}
	return sc
}

// SetIsEndStation sets the "is_end_station" field.
func (sc *StationCreate) SetIsEndStation(b bool) *StationCreate {
	sc.mutation.SetIsEndStation(b)
	return sc
}

// SetNillableIsEndStation sets the "is_end_station" field if the given value is not nil.
func (sc *StationCreate) SetNillableIsEndStation(b *bool) *StationCreate {
	if b != nil {
		sc.SetIsEndStation(*b)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StationCreate) SetCreatedAt(t time.Time) *StationCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StationCreate) SetNillableCreatedAt(t *time.Time) *StationCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StationCreate) SetID(i int64) *StationCreate {
	sc.mutation.SetID(i)
	return sc
}

// AddStationScheduleIDs adds the "station_schedules" edge to the Schedule entity by IDs.
func (sc *StationCreate) AddStationScheduleIDs(ids ...int64) *StationCreate {
	sc.mutation.AddStationScheduleIDs(ids...)
	return sc
}

// AddStationSchedules adds the "station_schedules" edges to the Schedule entity.
func (sc *StationCreate) AddStationSchedules(s ...*Schedule) *StationCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddStationScheduleIDs(ids...)
}

// Mutation returns the StationMutation object of the builder.
func (sc *StationCreate) Mutation() *StationMutation {
	return sc.mutation
}

// Save creates the Station in the database.
func (sc *StationCreate) Save(ctx context.Context) (*Station, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StationCreate) SaveX(ctx context.Context) *Station {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StationCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StationCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StationCreate) defaults() {
	if _, ok := sc.mutation.IsEndStation(); !ok {
		v := station.DefaultIsEndStation
		sc.mutation.SetIsEndStation(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := station.DefaultCreatedAt
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StationCreate) check() error {
	if _, ok := sc.mutation.IsEndStation(); !ok {
		return &ValidationError{Name: "is_end_station", err: errors.New(`ent: missing required field "Station.is_end_station"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Station.created_at"`)}
	}
	return nil
}

func (sc *StationCreate) sqlSave(ctx context.Context) (*Station, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StationCreate) createSpec() (*Station, *sqlgraph.CreateSpec) {
	var (
		_node = &Station{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(station.Table, sqlgraph.NewFieldSpec(station.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(station.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Latitude(); ok {
		_spec.SetField(station.FieldLatitude, field.TypeFloat32, value)
		_node.Latitude = &value
	}
	if value, ok := sc.mutation.Longitude(); ok {
		_spec.SetField(station.FieldLongitude, field.TypeFloat32, value)
		_node.Longitude = &value
	}
	if value, ok := sc.mutation.IsEndStation(); ok {
		_spec.SetField(station.FieldIsEndStation, field.TypeBool, value)
		_node.IsEndStation = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(station.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := sc.mutation.StationSchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   station.StationSchedulesTable,
			Columns: []string{station.StationSchedulesColumn},
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
//	client.Station.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StationUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sc *StationCreate) OnConflict(opts ...sql.ConflictOption) *StationUpsertOne {
	sc.conflict = opts
	return &StationUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Station.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *StationCreate) OnConflictColumns(columns ...string) *StationUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &StationUpsertOne{
		create: sc,
	}
}

type (
	// StationUpsertOne is the builder for "upsert"-ing
	//  one Station node.
	StationUpsertOne struct {
		create *StationCreate
	}

	// StationUpsert is the "OnConflict" setter.
	StationUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *StationUpsert) SetName(v string) *StationUpsert {
	u.Set(station.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StationUpsert) UpdateName() *StationUpsert {
	u.SetExcluded(station.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *StationUpsert) ClearName() *StationUpsert {
	u.SetNull(station.FieldName)
	return u
}

// SetLatitude sets the "latitude" field.
func (u *StationUpsert) SetLatitude(v float32) *StationUpsert {
	u.Set(station.FieldLatitude, v)
	return u
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *StationUpsert) UpdateLatitude() *StationUpsert {
	u.SetExcluded(station.FieldLatitude)
	return u
}

// AddLatitude adds v to the "latitude" field.
func (u *StationUpsert) AddLatitude(v float32) *StationUpsert {
	u.Add(station.FieldLatitude, v)
	return u
}

// ClearLatitude clears the value of the "latitude" field.
func (u *StationUpsert) ClearLatitude() *StationUpsert {
	u.SetNull(station.FieldLatitude)
	return u
}

// SetLongitude sets the "longitude" field.
func (u *StationUpsert) SetLongitude(v float32) *StationUpsert {
	u.Set(station.FieldLongitude, v)
	return u
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *StationUpsert) UpdateLongitude() *StationUpsert {
	u.SetExcluded(station.FieldLongitude)
	return u
}

// AddLongitude adds v to the "longitude" field.
func (u *StationUpsert) AddLongitude(v float32) *StationUpsert {
	u.Add(station.FieldLongitude, v)
	return u
}

// ClearLongitude clears the value of the "longitude" field.
func (u *StationUpsert) ClearLongitude() *StationUpsert {
	u.SetNull(station.FieldLongitude)
	return u
}

// SetIsEndStation sets the "is_end_station" field.
func (u *StationUpsert) SetIsEndStation(v bool) *StationUpsert {
	u.Set(station.FieldIsEndStation, v)
	return u
}

// UpdateIsEndStation sets the "is_end_station" field to the value that was provided on create.
func (u *StationUpsert) UpdateIsEndStation() *StationUpsert {
	u.SetExcluded(station.FieldIsEndStation)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *StationUpsert) SetCreatedAt(v time.Time) *StationUpsert {
	u.Set(station.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StationUpsert) UpdateCreatedAt() *StationUpsert {
	u.SetExcluded(station.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Station.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(station.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StationUpsertOne) UpdateNewValues() *StationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(station.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Station.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *StationUpsertOne) Ignore() *StationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StationUpsertOne) DoNothing() *StationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StationCreate.OnConflict
// documentation for more info.
func (u *StationUpsertOne) Update(set func(*StationUpsert)) *StationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StationUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StationUpsertOne) SetName(v string) *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StationUpsertOne) UpdateName() *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *StationUpsertOne) ClearName() *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.ClearName()
	})
}

// SetLatitude sets the "latitude" field.
func (u *StationUpsertOne) SetLatitude(v float32) *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.SetLatitude(v)
	})
}

// AddLatitude adds v to the "latitude" field.
func (u *StationUpsertOne) AddLatitude(v float32) *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.AddLatitude(v)
	})
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *StationUpsertOne) UpdateLatitude() *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.UpdateLatitude()
	})
}

// ClearLatitude clears the value of the "latitude" field.
func (u *StationUpsertOne) ClearLatitude() *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.ClearLatitude()
	})
}

// SetLongitude sets the "longitude" field.
func (u *StationUpsertOne) SetLongitude(v float32) *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.SetLongitude(v)
	})
}

// AddLongitude adds v to the "longitude" field.
func (u *StationUpsertOne) AddLongitude(v float32) *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.AddLongitude(v)
	})
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *StationUpsertOne) UpdateLongitude() *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.UpdateLongitude()
	})
}

// ClearLongitude clears the value of the "longitude" field.
func (u *StationUpsertOne) ClearLongitude() *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.ClearLongitude()
	})
}

// SetIsEndStation sets the "is_end_station" field.
func (u *StationUpsertOne) SetIsEndStation(v bool) *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.SetIsEndStation(v)
	})
}

// UpdateIsEndStation sets the "is_end_station" field to the value that was provided on create.
func (u *StationUpsertOne) UpdateIsEndStation() *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.UpdateIsEndStation()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *StationUpsertOne) SetCreatedAt(v time.Time) *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StationUpsertOne) UpdateCreatedAt() *StationUpsertOne {
	return u.Update(func(s *StationUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *StationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StationUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StationUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StationCreateBulk is the builder for creating many Station entities in bulk.
type StationCreateBulk struct {
	config
	err      error
	builders []*StationCreate
	conflict []sql.ConflictOption
}

// Save creates the Station entities in the database.
func (scb *StationCreateBulk) Save(ctx context.Context) ([]*Station, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Station, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StationMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StationCreateBulk) SaveX(ctx context.Context) []*Station {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StationCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StationCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Station.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StationUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (scb *StationCreateBulk) OnConflict(opts ...sql.ConflictOption) *StationUpsertBulk {
	scb.conflict = opts
	return &StationUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Station.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *StationCreateBulk) OnConflictColumns(columns ...string) *StationUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &StationUpsertBulk{
		create: scb,
	}
}

// StationUpsertBulk is the builder for "upsert"-ing
// a bulk of Station nodes.
type StationUpsertBulk struct {
	create *StationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Station.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(station.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StationUpsertBulk) UpdateNewValues() *StationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(station.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Station.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *StationUpsertBulk) Ignore() *StationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StationUpsertBulk) DoNothing() *StationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StationCreateBulk.OnConflict
// documentation for more info.
func (u *StationUpsertBulk) Update(set func(*StationUpsert)) *StationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StationUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StationUpsertBulk) SetName(v string) *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StationUpsertBulk) UpdateName() *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *StationUpsertBulk) ClearName() *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.ClearName()
	})
}

// SetLatitude sets the "latitude" field.
func (u *StationUpsertBulk) SetLatitude(v float32) *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.SetLatitude(v)
	})
}

// AddLatitude adds v to the "latitude" field.
func (u *StationUpsertBulk) AddLatitude(v float32) *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.AddLatitude(v)
	})
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *StationUpsertBulk) UpdateLatitude() *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.UpdateLatitude()
	})
}

// ClearLatitude clears the value of the "latitude" field.
func (u *StationUpsertBulk) ClearLatitude() *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.ClearLatitude()
	})
}

// SetLongitude sets the "longitude" field.
func (u *StationUpsertBulk) SetLongitude(v float32) *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.SetLongitude(v)
	})
}

// AddLongitude adds v to the "longitude" field.
func (u *StationUpsertBulk) AddLongitude(v float32) *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.AddLongitude(v)
	})
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *StationUpsertBulk) UpdateLongitude() *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.UpdateLongitude()
	})
}

// ClearLongitude clears the value of the "longitude" field.
func (u *StationUpsertBulk) ClearLongitude() *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.ClearLongitude()
	})
}

// SetIsEndStation sets the "is_end_station" field.
func (u *StationUpsertBulk) SetIsEndStation(v bool) *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.SetIsEndStation(v)
	})
}

// UpdateIsEndStation sets the "is_end_station" field to the value that was provided on create.
func (u *StationUpsertBulk) UpdateIsEndStation() *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.UpdateIsEndStation()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *StationUpsertBulk) SetCreatedAt(v time.Time) *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StationUpsertBulk) UpdateCreatedAt() *StationUpsertBulk {
	return u.Update(func(s *StationUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *StationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
