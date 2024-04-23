// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"team-manager/ent/direction"
	"team-manager/ent/predicate"
	"team-manager/ent/schedule"
	"team-manager/ent/station"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleUpdate is the builder for updating Schedule entities.
type ScheduleUpdate struct {
	config
	hooks     []Hook
	mutation  *ScheduleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ScheduleUpdate builder.
func (su *ScheduleUpdate) Where(ps ...predicate.Schedule) *ScheduleUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetStationID sets the "station_id" field.
func (su *ScheduleUpdate) SetStationID(i int64) *ScheduleUpdate {
	su.mutation.SetStationID(i)
	return su
}

// SetNillableStationID sets the "station_id" field if the given value is not nil.
func (su *ScheduleUpdate) SetNillableStationID(i *int64) *ScheduleUpdate {
	if i != nil {
		su.SetStationID(*i)
	}
	return su
}

// SetDirectionID sets the "direction_id" field.
func (su *ScheduleUpdate) SetDirectionID(i int64) *ScheduleUpdate {
	su.mutation.SetDirectionID(i)
	return su
}

// SetNillableDirectionID sets the "direction_id" field if the given value is not nil.
func (su *ScheduleUpdate) SetNillableDirectionID(i *int64) *ScheduleUpdate {
	if i != nil {
		su.SetDirectionID(*i)
	}
	return su
}

// SetTime sets the "time" field.
func (su *ScheduleUpdate) SetTime(s string) *ScheduleUpdate {
	su.mutation.SetTime(s)
	return su
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (su *ScheduleUpdate) SetNillableTime(s *string) *ScheduleUpdate {
	if s != nil {
		su.SetTime(*s)
	}
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *ScheduleUpdate) SetCreatedAt(t time.Time) *ScheduleUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *ScheduleUpdate) SetNillableCreatedAt(t *time.Time) *ScheduleUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetStation sets the "station" edge to the Station entity.
func (su *ScheduleUpdate) SetStation(s *Station) *ScheduleUpdate {
	return su.SetStationID(s.ID)
}

// SetDirection sets the "direction" edge to the Direction entity.
func (su *ScheduleUpdate) SetDirection(d *Direction) *ScheduleUpdate {
	return su.SetDirectionID(d.ID)
}

// Mutation returns the ScheduleMutation object of the builder.
func (su *ScheduleUpdate) Mutation() *ScheduleMutation {
	return su.mutation
}

// ClearStation clears the "station" edge to the Station entity.
func (su *ScheduleUpdate) ClearStation() *ScheduleUpdate {
	su.mutation.ClearStation()
	return su
}

// ClearDirection clears the "direction" edge to the Direction entity.
func (su *ScheduleUpdate) ClearDirection() *ScheduleUpdate {
	su.mutation.ClearDirection()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScheduleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScheduleUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScheduleUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScheduleUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ScheduleUpdate) check() error {
	if _, ok := su.mutation.StationID(); su.mutation.StationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Schedule.station"`)
	}
	if _, ok := su.mutation.DirectionID(); su.mutation.DirectionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Schedule.direction"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *ScheduleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ScheduleUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *ScheduleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(schedule.Table, schedule.Columns, sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Time(); ok {
		_spec.SetField(schedule.FieldTime, field.TypeString, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(schedule.FieldCreatedAt, field.TypeTime, value)
	}
	if su.mutation.StationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.StationTable,
			Columns: []string{schedule.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(station.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.StationTable,
			Columns: []string{schedule.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(station.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.DirectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.DirectionTable,
			Columns: []string{schedule.DirectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(direction.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.DirectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.DirectionTable,
			Columns: []string{schedule.DirectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(direction.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ScheduleUpdateOne is the builder for updating a single Schedule entity.
type ScheduleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ScheduleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetStationID sets the "station_id" field.
func (suo *ScheduleUpdateOne) SetStationID(i int64) *ScheduleUpdateOne {
	suo.mutation.SetStationID(i)
	return suo
}

// SetNillableStationID sets the "station_id" field if the given value is not nil.
func (suo *ScheduleUpdateOne) SetNillableStationID(i *int64) *ScheduleUpdateOne {
	if i != nil {
		suo.SetStationID(*i)
	}
	return suo
}

// SetDirectionID sets the "direction_id" field.
func (suo *ScheduleUpdateOne) SetDirectionID(i int64) *ScheduleUpdateOne {
	suo.mutation.SetDirectionID(i)
	return suo
}

// SetNillableDirectionID sets the "direction_id" field if the given value is not nil.
func (suo *ScheduleUpdateOne) SetNillableDirectionID(i *int64) *ScheduleUpdateOne {
	if i != nil {
		suo.SetDirectionID(*i)
	}
	return suo
}

// SetTime sets the "time" field.
func (suo *ScheduleUpdateOne) SetTime(s string) *ScheduleUpdateOne {
	suo.mutation.SetTime(s)
	return suo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (suo *ScheduleUpdateOne) SetNillableTime(s *string) *ScheduleUpdateOne {
	if s != nil {
		suo.SetTime(*s)
	}
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *ScheduleUpdateOne) SetCreatedAt(t time.Time) *ScheduleUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *ScheduleUpdateOne) SetNillableCreatedAt(t *time.Time) *ScheduleUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetStation sets the "station" edge to the Station entity.
func (suo *ScheduleUpdateOne) SetStation(s *Station) *ScheduleUpdateOne {
	return suo.SetStationID(s.ID)
}

// SetDirection sets the "direction" edge to the Direction entity.
func (suo *ScheduleUpdateOne) SetDirection(d *Direction) *ScheduleUpdateOne {
	return suo.SetDirectionID(d.ID)
}

// Mutation returns the ScheduleMutation object of the builder.
func (suo *ScheduleUpdateOne) Mutation() *ScheduleMutation {
	return suo.mutation
}

// ClearStation clears the "station" edge to the Station entity.
func (suo *ScheduleUpdateOne) ClearStation() *ScheduleUpdateOne {
	suo.mutation.ClearStation()
	return suo
}

// ClearDirection clears the "direction" edge to the Direction entity.
func (suo *ScheduleUpdateOne) ClearDirection() *ScheduleUpdateOne {
	suo.mutation.ClearDirection()
	return suo
}

// Where appends a list predicates to the ScheduleUpdate builder.
func (suo *ScheduleUpdateOne) Where(ps ...predicate.Schedule) *ScheduleUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ScheduleUpdateOne) Select(field string, fields ...string) *ScheduleUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Schedule entity.
func (suo *ScheduleUpdateOne) Save(ctx context.Context) (*Schedule, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScheduleUpdateOne) SaveX(ctx context.Context) *Schedule {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScheduleUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScheduleUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ScheduleUpdateOne) check() error {
	if _, ok := suo.mutation.StationID(); suo.mutation.StationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Schedule.station"`)
	}
	if _, ok := suo.mutation.DirectionID(); suo.mutation.DirectionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Schedule.direction"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *ScheduleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ScheduleUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *ScheduleUpdateOne) sqlSave(ctx context.Context) (_node *Schedule, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(schedule.Table, schedule.Columns, sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Schedule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schedule.FieldID)
		for _, f := range fields {
			if !schedule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != schedule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Time(); ok {
		_spec.SetField(schedule.FieldTime, field.TypeString, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(schedule.FieldCreatedAt, field.TypeTime, value)
	}
	if suo.mutation.StationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.StationTable,
			Columns: []string{schedule.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(station.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.StationTable,
			Columns: []string{schedule.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(station.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.DirectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.DirectionTable,
			Columns: []string{schedule.DirectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(direction.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.DirectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.DirectionTable,
			Columns: []string{schedule.DirectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(direction.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Schedule{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}