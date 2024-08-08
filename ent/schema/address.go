package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Address struct {
	ent.Schema
}

func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("street"),
		field.String("house_number"),
		field.String("postal_code"),
		field.String("city"),
	}
}

func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("companies", Company.Type).
			Through("company_addresses", CompanyAddress.Type),
		edge.To("employees", Employee.Type).
			Through("employee_addresses", EmployeeAddress.Type),
	}
}
