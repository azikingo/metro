// Code generated by ent, DO NOT EDIT.

package ent

import (
	"team-manager/ent/direction"
	"team-manager/ent/schedule"
	"team-manager/ent/schema"
	"team-manager/ent/station"
	"team-manager/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	directionFields := schema.Direction{}.Fields()
	_ = directionFields
	// directionDescCreatedAt is the schema descriptor for created_at field.
	directionDescCreatedAt := directionFields[3].Descriptor()
	// direction.DefaultCreatedAt holds the default value on creation for the created_at field.
	direction.DefaultCreatedAt = directionDescCreatedAt.Default.(time.Time)
	scheduleFields := schema.Schedule{}.Fields()
	_ = scheduleFields
	// scheduleDescCreatedAt is the schema descriptor for created_at field.
	scheduleDescCreatedAt := scheduleFields[4].Descriptor()
	// schedule.DefaultCreatedAt holds the default value on creation for the created_at field.
	schedule.DefaultCreatedAt = scheduleDescCreatedAt.Default.(time.Time)
	stationFields := schema.Station{}.Fields()
	_ = stationFields
	// stationDescIsEndStation is the schema descriptor for is_end_station field.
	stationDescIsEndStation := stationFields[4].Descriptor()
	// station.DefaultIsEndStation holds the default value on creation for the is_end_station field.
	station.DefaultIsEndStation = stationDescIsEndStation.Default.(bool)
	// stationDescCreatedAt is the schema descriptor for created_at field.
	stationDescCreatedAt := stationFields[5].Descriptor()
	// station.DefaultCreatedAt holds the default value on creation for the created_at field.
	station.DefaultCreatedAt = stationDescCreatedAt.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
}
