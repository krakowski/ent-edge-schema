// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/address"
	"entgo.io/bug/ent/company"
	"entgo.io/bug/ent/companyaddress"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// CompanyAddress is the model entity for the CompanyAddress schema.
type CompanyAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CompanyID holds the value of the "company_id" field.
	CompanyID int `json:"company_id,omitempty"`
	// AddressID holds the value of the "address_id" field.
	AddressID int `json:"address_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyAddressQuery when eager-loading is set.
	Edges        CompanyAddressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CompanyAddressEdges holds the relations/edges for other nodes in the graph.
type CompanyAddressEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Address holds the value of the address edge.
	Address *Address `json:"address,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompanyAddressEdges) CompanyOrErr() (*Company, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// AddressOrErr returns the Address value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompanyAddressEdges) AddressOrErr() (*Address, error) {
	if e.Address != nil {
		return e.Address, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: address.Label}
	}
	return nil, &NotLoadedError{edge: "address"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CompanyAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case companyaddress.FieldID, companyaddress.FieldCompanyID, companyaddress.FieldAddressID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CompanyAddress fields.
func (ca *CompanyAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case companyaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ca.ID = int(value.Int64)
		case companyaddress.FieldCompanyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field company_id", values[i])
			} else if value.Valid {
				ca.CompanyID = int(value.Int64)
			}
		case companyaddress.FieldAddressID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field address_id", values[i])
			} else if value.Valid {
				ca.AddressID = int(value.Int64)
			}
		default:
			ca.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CompanyAddress.
// This includes values selected through modifiers, order, etc.
func (ca *CompanyAddress) Value(name string) (ent.Value, error) {
	return ca.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the CompanyAddress entity.
func (ca *CompanyAddress) QueryCompany() *CompanyQuery {
	return NewCompanyAddressClient(ca.config).QueryCompany(ca)
}

// QueryAddress queries the "address" edge of the CompanyAddress entity.
func (ca *CompanyAddress) QueryAddress() *AddressQuery {
	return NewCompanyAddressClient(ca.config).QueryAddress(ca)
}

// Update returns a builder for updating this CompanyAddress.
// Note that you need to call CompanyAddress.Unwrap() before calling this method if this CompanyAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (ca *CompanyAddress) Update() *CompanyAddressUpdateOne {
	return NewCompanyAddressClient(ca.config).UpdateOne(ca)
}

// Unwrap unwraps the CompanyAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ca *CompanyAddress) Unwrap() *CompanyAddress {
	_tx, ok := ca.config.driver.(*txDriver)
	if !ok {
		panic("ent: CompanyAddress is not a transactional entity")
	}
	ca.config.driver = _tx.drv
	return ca
}

// String implements the fmt.Stringer.
func (ca *CompanyAddress) String() string {
	var builder strings.Builder
	builder.WriteString("CompanyAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ca.ID))
	builder.WriteString("company_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.CompanyID))
	builder.WriteString(", ")
	builder.WriteString("address_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.AddressID))
	builder.WriteByte(')')
	return builder.String()
}

// CompanyAddresses is a parsable slice of CompanyAddress.
type CompanyAddresses []*CompanyAddress
