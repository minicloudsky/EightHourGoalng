package schema

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
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

/**
field.Time("created_at").
    SchemaType(map[string]string{
        dialect.MySQL:    "datetime",
        dialect.Postgres: "date",
    })
*/
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).
			Annotations(
				entsql.Annotation{
					Table:       "",
					Charset:     "",
					Collation:   "",
					Default:     "CURRENT_TIMESTAMP",
					Options:     "",
					Size:        0,
					Incremental: nil,
					OnDelete:    "",
					Check:       "",
					Checks:      nil,
				},
			),
		field.Time("update_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).UpdateDefault(time.Now).
			Annotations(
				entsql.Annotation{
					Table:       "",
					Charset:     "",
					Collation:   "",
					Default:     "CURRENT_TIMESTAMP",
					Options:     "ON UPDATE CURRENT_TIMESTAMP",
					Size:        0,
					Incremental: nil,
					OnDelete:    "",
					Check:       "",
					Checks:      nil,
				},
			),
	}
}

// DetailsMixin implements the ent.Mixin for sharing
// entity details fields with package schemas.
type DetailsMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (DetailsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			NotEmpty(),
	}
}

// -------------------------------------------------
// Schema definition

// Pet schema mixed-in the DetailsMixin fields and therefore
// has 3 fields: `age`, `name` and `weight`.
type Pet struct {
	ent.Schema
}

func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DetailsMixin{},
	}
}

func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.Float("weight"),
	}
}
