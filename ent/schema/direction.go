package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Direction struct {
	ent.Schema
}

func (Direction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("start_station_id"),
		field.Int64("end_station_id"),
		field.Time("created_at").Default(time.Now()),
	}
}

func (Direction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("direction_schedules", Schedule.Type),
	}
}
