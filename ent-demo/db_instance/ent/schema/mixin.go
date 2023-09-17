package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

// -------------------------------------------------
// Mixin definition

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).
			Annotations(
				entsql.Annotation{
					Default: "CURRENT_TIMESTAMP",
				},
			),
		field.Time("update_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).UpdateDefault(time.Now).
			Annotations(
				entsql.Annotation{
					Default: "CURRENT_TIMESTAMP",
					Options: "ON UPDATE CURRENT_TIMESTAMP",
				},
			),
	}
}
