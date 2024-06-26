// Code generated by ent, DO NOT EDIT.

package ent

import (
	"team-manager/ent/direction"
	"team-manager/ent/predicate"
	"team-manager/ent/schedule"
	"team-manager/ent/station"
	"team-manager/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 4)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   direction.Table,
			Columns: direction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: direction.FieldID,
			},
		},
		Type: "Direction",
		Fields: map[string]*sqlgraph.FieldSpec{
			direction.FieldStartStationID: {Type: field.TypeInt64, Column: direction.FieldStartStationID},
			direction.FieldEndStationID:   {Type: field.TypeInt64, Column: direction.FieldEndStationID},
			direction.FieldCreatedAt:      {Type: field.TypeTime, Column: direction.FieldCreatedAt},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   schedule.Table,
			Columns: schedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: schedule.FieldID,
			},
		},
		Type: "Schedule",
		Fields: map[string]*sqlgraph.FieldSpec{
			schedule.FieldStationID:   {Type: field.TypeInt64, Column: schedule.FieldStationID},
			schedule.FieldDirectionID: {Type: field.TypeInt64, Column: schedule.FieldDirectionID},
			schedule.FieldTime:        {Type: field.TypeString, Column: schedule.FieldTime},
			schedule.FieldCreatedAt:   {Type: field.TypeTime, Column: schedule.FieldCreatedAt},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   station.Table,
			Columns: station.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: station.FieldID,
			},
		},
		Type: "Station",
		Fields: map[string]*sqlgraph.FieldSpec{
			station.FieldName:         {Type: field.TypeString, Column: station.FieldName},
			station.FieldLatitude:     {Type: field.TypeFloat32, Column: station.FieldLatitude},
			station.FieldLongitude:    {Type: field.TypeFloat32, Column: station.FieldLongitude},
			station.FieldIsEndStation: {Type: field.TypeBool, Column: station.FieldIsEndStation},
			station.FieldCreatedAt:    {Type: field.TypeTime, Column: station.FieldCreatedAt},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldTgID:      {Type: field.TypeInt64, Column: user.FieldTgID},
			user.FieldUsername:  {Type: field.TypeString, Column: user.FieldUsername},
			user.FieldName:      {Type: field.TypeString, Column: user.FieldName},
			user.FieldSurname:   {Type: field.TypeString, Column: user.FieldSurname},
			user.FieldEmail:     {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldPhone:     {Type: field.TypeString, Column: user.FieldPhone},
			user.FieldCreatedAt: {Type: field.TypeTime, Column: user.FieldCreatedAt},
		},
	}
	graph.MustAddE(
		"direction_schedules",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   direction.DirectionSchedulesTable,
			Columns: []string{direction.DirectionSchedulesColumn},
			Bidi:    false,
		},
		"Direction",
		"Schedule",
	)
	graph.MustAddE(
		"station",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.StationTable,
			Columns: []string{schedule.StationColumn},
			Bidi:    false,
		},
		"Schedule",
		"Station",
	)
	graph.MustAddE(
		"direction",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.DirectionTable,
			Columns: []string{schedule.DirectionColumn},
			Bidi:    false,
		},
		"Schedule",
		"Direction",
	)
	graph.MustAddE(
		"station_schedules",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   station.StationSchedulesTable,
			Columns: []string{station.StationSchedulesColumn},
			Bidi:    false,
		},
		"Station",
		"Schedule",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (dq *DirectionQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DirectionQuery builder.
func (dq *DirectionQuery) Filter() *DirectionFilter {
	return &DirectionFilter{config: dq.config, predicateAdder: dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DirectionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DirectionMutation builder.
func (m *DirectionMutation) Filter() *DirectionFilter {
	return &DirectionFilter{config: m.config, predicateAdder: m}
}

// DirectionFilter provides a generic filtering capability at runtime for DirectionQuery.
type DirectionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *DirectionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *DirectionFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(direction.FieldID))
}

// WhereStartStationID applies the entql int64 predicate on the start_station_id field.
func (f *DirectionFilter) WhereStartStationID(p entql.Int64P) {
	f.Where(p.Field(direction.FieldStartStationID))
}

// WhereEndStationID applies the entql int64 predicate on the end_station_id field.
func (f *DirectionFilter) WhereEndStationID(p entql.Int64P) {
	f.Where(p.Field(direction.FieldEndStationID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *DirectionFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(direction.FieldCreatedAt))
}

// WhereHasDirectionSchedules applies a predicate to check if query has an edge direction_schedules.
func (f *DirectionFilter) WhereHasDirectionSchedules() {
	f.Where(entql.HasEdge("direction_schedules"))
}

// WhereHasDirectionSchedulesWith applies a predicate to check if query has an edge direction_schedules with a given conditions (other predicates).
func (f *DirectionFilter) WhereHasDirectionSchedulesWith(preds ...predicate.Schedule) {
	f.Where(entql.HasEdgeWith("direction_schedules", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (sq *ScheduleQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ScheduleQuery builder.
func (sq *ScheduleQuery) Filter() *ScheduleFilter {
	return &ScheduleFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *ScheduleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ScheduleMutation builder.
func (m *ScheduleMutation) Filter() *ScheduleFilter {
	return &ScheduleFilter{config: m.config, predicateAdder: m}
}

// ScheduleFilter provides a generic filtering capability at runtime for ScheduleQuery.
type ScheduleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ScheduleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *ScheduleFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(schedule.FieldID))
}

// WhereStationID applies the entql int64 predicate on the station_id field.
func (f *ScheduleFilter) WhereStationID(p entql.Int64P) {
	f.Where(p.Field(schedule.FieldStationID))
}

// WhereDirectionID applies the entql int64 predicate on the direction_id field.
func (f *ScheduleFilter) WhereDirectionID(p entql.Int64P) {
	f.Where(p.Field(schedule.FieldDirectionID))
}

// WhereTime applies the entql string predicate on the time field.
func (f *ScheduleFilter) WhereTime(p entql.StringP) {
	f.Where(p.Field(schedule.FieldTime))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *ScheduleFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(schedule.FieldCreatedAt))
}

// WhereHasStation applies a predicate to check if query has an edge station.
func (f *ScheduleFilter) WhereHasStation() {
	f.Where(entql.HasEdge("station"))
}

// WhereHasStationWith applies a predicate to check if query has an edge station with a given conditions (other predicates).
func (f *ScheduleFilter) WhereHasStationWith(preds ...predicate.Station) {
	f.Where(entql.HasEdgeWith("station", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasDirection applies a predicate to check if query has an edge direction.
func (f *ScheduleFilter) WhereHasDirection() {
	f.Where(entql.HasEdge("direction"))
}

// WhereHasDirectionWith applies a predicate to check if query has an edge direction with a given conditions (other predicates).
func (f *ScheduleFilter) WhereHasDirectionWith(preds ...predicate.Direction) {
	f.Where(entql.HasEdgeWith("direction", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (sq *StationQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the StationQuery builder.
func (sq *StationQuery) Filter() *StationFilter {
	return &StationFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *StationMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the StationMutation builder.
func (m *StationMutation) Filter() *StationFilter {
	return &StationFilter{config: m.config, predicateAdder: m}
}

// StationFilter provides a generic filtering capability at runtime for StationQuery.
type StationFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *StationFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *StationFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(station.FieldID))
}

// WhereName applies the entql string predicate on the name field.
func (f *StationFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(station.FieldName))
}

// WhereLatitude applies the entql float32 predicate on the latitude field.
func (f *StationFilter) WhereLatitude(p entql.Float32P) {
	f.Where(p.Field(station.FieldLatitude))
}

// WhereLongitude applies the entql float32 predicate on the longitude field.
func (f *StationFilter) WhereLongitude(p entql.Float32P) {
	f.Where(p.Field(station.FieldLongitude))
}

// WhereIsEndStation applies the entql bool predicate on the is_end_station field.
func (f *StationFilter) WhereIsEndStation(p entql.BoolP) {
	f.Where(p.Field(station.FieldIsEndStation))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *StationFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(station.FieldCreatedAt))
}

// WhereHasStationSchedules applies a predicate to check if query has an edge station_schedules.
func (f *StationFilter) WhereHasStationSchedules() {
	f.Where(entql.HasEdge("station_schedules"))
}

// WhereHasStationSchedulesWith applies a predicate to check if query has an edge station_schedules with a given conditions (other predicates).
func (f *StationFilter) WhereHasStationSchedulesWith(preds ...predicate.Schedule) {
	f.Where(entql.HasEdgeWith("station_schedules", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *UserFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(user.FieldID))
}

// WhereTgID applies the entql int64 predicate on the tg_id field.
func (f *UserFilter) WhereTgID(p entql.Int64P) {
	f.Where(p.Field(user.FieldTgID))
}

// WhereUsername applies the entql string predicate on the username field.
func (f *UserFilter) WhereUsername(p entql.StringP) {
	f.Where(p.Field(user.FieldUsername))
}

// WhereName applies the entql string predicate on the name field.
func (f *UserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WhereSurname applies the entql string predicate on the surname field.
func (f *UserFilter) WhereSurname(p entql.StringP) {
	f.Where(p.Field(user.FieldSurname))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WherePhone applies the entql string predicate on the phone field.
func (f *UserFilter) WherePhone(p entql.StringP) {
	f.Where(p.Field(user.FieldPhone))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *UserFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreatedAt))
}
