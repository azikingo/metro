package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Station struct {
	ent.Schema
}

func (Station) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Optional(),
		field.Float32("latitude").Nillable().Optional(),
		field.Float32("longitude").Nillable().Optional(),
		field.Bool("is_end_station").Default(false),
		field.Time("created_at").Default(time.Now()),
	}
}

func (Station) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("station_schedules", Schedule.Type),
	}
}
