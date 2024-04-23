package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type Schedule struct {
	ent.Schema
}

func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("station_id"),
		field.Int64("direction_id"),
		field.String("time"),
		field.Time("created_at").Default(time.Now()),
	}
}

func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("station", Station.Type).
			Field("station_id").
			Ref("station_schedules").
			Required().
			Unique(),
		edge.From("direction", Direction.Type).
			Field("direction_id").
			Ref("direction_schedules").
			Required().
			Unique(),
	}
}

func (Schedule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("station_id", "direction_id", "time").
			Unique().Annotations(entsql.DescColumns("station_id", "direction_id", "time")),
	}
}
