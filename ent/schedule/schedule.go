// Code generated by ent, DO NOT EDIT.

package schedule

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the schedule type in the database.
	Label = "schedule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStationID holds the string denoting the station_id field in the database.
	FieldStationID = "station_id"
	// FieldDirectionID holds the string denoting the direction_id field in the database.
	FieldDirectionID = "direction_id"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeStation holds the string denoting the station edge name in mutations.
	EdgeStation = "station"
	// EdgeDirection holds the string denoting the direction edge name in mutations.
	EdgeDirection = "direction"
	// Table holds the table name of the schedule in the database.
	Table = "schedules"
	// StationTable is the table that holds the station relation/edge.
	StationTable = "schedules"
	// StationInverseTable is the table name for the Station entity.
	// It exists in this package in order to avoid circular dependency with the "station" package.
	StationInverseTable = "stations"
	// StationColumn is the table column denoting the station relation/edge.
	StationColumn = "station_id"
	// DirectionTable is the table that holds the direction relation/edge.
	DirectionTable = "schedules"
	// DirectionInverseTable is the table name for the Direction entity.
	// It exists in this package in order to avoid circular dependency with the "direction" package.
	DirectionInverseTable = "directions"
	// DirectionColumn is the table column denoting the direction relation/edge.
	DirectionColumn = "direction_id"
)

// Columns holds all SQL columns for schedule fields.
var Columns = []string{
	FieldID,
	FieldStationID,
	FieldDirectionID,
	FieldTime,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Schedule queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStationID orders the results by the station_id field.
func ByStationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStationID, opts...).ToFunc()
}

// ByDirectionID orders the results by the direction_id field.
func ByDirectionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDirectionID, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByStationField orders the results by station field.
func ByStationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStationStep(), sql.OrderByField(field, opts...))
	}
}

// ByDirectionField orders the results by direction field.
func ByDirectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDirectionStep(), sql.OrderByField(field, opts...))
	}
}
func newStationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, StationTable, StationColumn),
	)
}
func newDirectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DirectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DirectionTable, DirectionColumn),
	)
}