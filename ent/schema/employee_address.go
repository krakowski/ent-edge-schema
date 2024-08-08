package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type EmployeeAddress struct {
	ent.Schema
}

func (EmployeeAddress) Fields() []ent.Field {
	return []ent.Field{
		field.Int("employee_id"),
		field.Int("address_id"),
	}
}

func (EmployeeAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Unique().
			Required().
			Field("employee_id"),
		edge.To("address", Address.Type).
			Unique().
			Required().
			Field("address_id"),
	}
}

func (EmployeeAddress) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("employee_id").
			Unique(),
	}
}
