package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type CompanyAddress struct {
	ent.Schema
}

func (CompanyAddress) Fields() []ent.Field {
	return []ent.Field{
		field.Int("company_id"),
		field.Int("address_id"),
	}
}

func (CompanyAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Unique().
			Required().
			Field("company_id"),
		edge.To("address", Address.Type).
			Unique().
			Required().
			Field("address_id"),
	}
}

func (CompanyAddress) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("company_id").
			Unique(),
	}
}
