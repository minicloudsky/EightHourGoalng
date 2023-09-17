package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)
import "entgo.io/ent/schema/field"

// DbInstance holds the schema definition for the DbInstance entity.
type DbInstance struct {
	ent.Schema
}

func (DbInstance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Annotations of the DBInstance.
func (DbInstance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "t_db_instance",
		},
	}
}

//type EmptyJson map[string]string{}

// Fields /***
// Fields of the DbInstance.
func (DbInstance) Fields() []ent.Field {
	return []ent.Field{
		field.String("instance_id").NotEmpty().Comment("实例id").Unique(),
		field.String("instance_name").NotEmpty().Comment("实例名").Default(""),
		field.String("host").NotEmpty().Comment("实例连接地址").Default(""),
		field.String("env").NotEmpty().Comment("环境"),
		field.String("instance_type").NotEmpty().Comment("实例类型").Default(""),
		field.String("engine").NotEmpty().Comment("引擎类型").Default(""),
		field.String("engine_version").NotEmpty().Comment("版本").Default(""),
		field.String("specification").NotEmpty().Comment("配置规格").Default(""),
		field.String("instance_status").NotEmpty().Comment("实例状态").Default(""),
		field.String("instance_create_time").NotEmpty().Comment("实例创建时间").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.JSON("annotations", map[string]string{}).Comment("实例注释; kv json").Optional(),
		field.JSON("labels", map[string]string{}).Comment("实例标签; kv json").Optional(),
	}
}

// Edges of the DbInstance.
func (DbInstance) Edges() []ent.Edge {
	return nil
}
