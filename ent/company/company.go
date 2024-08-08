// Code generated by ent, DO NOT EDIT.

package company

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeAddress holds the string denoting the address edge name in mutations.
	EdgeAddress = "address"
	// EdgeCompanyAddresses holds the string denoting the company_addresses edge name in mutations.
	EdgeCompanyAddresses = "company_addresses"
	// Table holds the table name of the company in the database.
	Table = "companies"
	// AddressTable is the table that holds the address relation/edge. The primary key declared below.
	AddressTable = "company_addresses"
	// AddressInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressInverseTable = "addresses"
	// CompanyAddressesTable is the table that holds the company_addresses relation/edge.
	CompanyAddressesTable = "company_addresses"
	// CompanyAddressesInverseTable is the table name for the CompanyAddress entity.
	// It exists in this package in order to avoid circular dependency with the "companyaddress" package.
	CompanyAddressesInverseTable = "company_addresses"
	// CompanyAddressesColumn is the table column denoting the company_addresses relation/edge.
	CompanyAddressesColumn = "company_id"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// AddressPrimaryKey and AddressColumn2 are the table columns denoting the
	// primary key for the address relation (M2M).
	AddressPrimaryKey = []string{"address_id", "company_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Company queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAddressCount orders the results by address count.
func ByAddressCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAddressStep(), opts...)
	}
}

// ByAddress orders the results by address terms.
func ByAddress(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAddressStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompanyAddressesCount orders the results by company_addresses count.
func ByCompanyAddressesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCompanyAddressesStep(), opts...)
	}
}

// ByCompanyAddresses orders the results by company_addresses terms.
func ByCompanyAddresses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyAddressesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AddressTable, AddressPrimaryKey...),
	)
}
func newCompanyAddressesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyAddressesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CompanyAddressesTable, CompanyAddressesColumn),
	)
}
