package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Employee struct {
	ent.Schema
}

func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstname"),
		field.String("lastname"),
	}
}

func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("address", Address.Type).
			Ref("employees").
			Through("employee_addresses", EmployeeAddress.Type),
	}
}
