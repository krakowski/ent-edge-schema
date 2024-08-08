// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/address"
	"entgo.io/bug/ent/company"
	"entgo.io/bug/ent/employeeaddress"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeAddressUpdate is the builder for updating EmployeeAddress entities.
type EmployeeAddressUpdate struct {
	config
	hooks    []Hook
	mutation *EmployeeAddressMutation
}

// Where appends a list predicates to the EmployeeAddressUpdate builder.
func (eau *EmployeeAddressUpdate) Where(ps ...predicate.EmployeeAddress) *EmployeeAddressUpdate {
	eau.mutation.Where(ps...)
	return eau
}

// SetEmployeeID sets the "employee_id" field.
func (eau *EmployeeAddressUpdate) SetEmployeeID(i int) *EmployeeAddressUpdate {
	eau.mutation.SetEmployeeID(i)
	return eau
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (eau *EmployeeAddressUpdate) SetNillableEmployeeID(i *int) *EmployeeAddressUpdate {
	if i != nil {
		eau.SetEmployeeID(*i)
	}
	return eau
}

// SetAddressID sets the "address_id" field.
func (eau *EmployeeAddressUpdate) SetAddressID(i int) *EmployeeAddressUpdate {
	eau.mutation.SetAddressID(i)
	return eau
}

// SetNillableAddressID sets the "address_id" field if the given value is not nil.
func (eau *EmployeeAddressUpdate) SetNillableAddressID(i *int) *EmployeeAddressUpdate {
	if i != nil {
		eau.SetAddressID(*i)
	}
	return eau
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (eau *EmployeeAddressUpdate) SetCompanyID(id int) *EmployeeAddressUpdate {
	eau.mutation.SetCompanyID(id)
	return eau
}

// SetCompany sets the "company" edge to the Company entity.
func (eau *EmployeeAddressUpdate) SetCompany(c *Company) *EmployeeAddressUpdate {
	return eau.SetCompanyID(c.ID)
}

// SetAddress sets the "address" edge to the Address entity.
func (eau *EmployeeAddressUpdate) SetAddress(a *Address) *EmployeeAddressUpdate {
	return eau.SetAddressID(a.ID)
}

// Mutation returns the EmployeeAddressMutation object of the builder.
func (eau *EmployeeAddressUpdate) Mutation() *EmployeeAddressMutation {
	return eau.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (eau *EmployeeAddressUpdate) ClearCompany() *EmployeeAddressUpdate {
	eau.mutation.ClearCompany()
	return eau
}

// ClearAddress clears the "address" edge to the Address entity.
func (eau *EmployeeAddressUpdate) ClearAddress() *EmployeeAddressUpdate {
	eau.mutation.ClearAddress()
	return eau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eau *EmployeeAddressUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eau.sqlSave, eau.mutation, eau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eau *EmployeeAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := eau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eau *EmployeeAddressUpdate) Exec(ctx context.Context) error {
	_, err := eau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eau *EmployeeAddressUpdate) ExecX(ctx context.Context) {
	if err := eau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eau *EmployeeAddressUpdate) check() error {
	if eau.mutation.CompanyCleared() && len(eau.mutation.CompanyIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EmployeeAddress.company"`)
	}
	if eau.mutation.AddressCleared() && len(eau.mutation.AddressIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EmployeeAddress.address"`)
	}
	return nil
}

func (eau *EmployeeAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(employeeaddress.Table, employeeaddress.Columns, sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt))
	if ps := eau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if eau.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   employeeaddress.CompanyTable,
			Columns: []string{employeeaddress.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eau.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   employeeaddress.CompanyTable,
			Columns: []string{employeeaddress.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eau.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   employeeaddress.AddressTable,
			Columns: []string{employeeaddress.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eau.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   employeeaddress.AddressTable,
			Columns: []string{employeeaddress.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employeeaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eau.mutation.done = true
	return n, nil
}

// EmployeeAddressUpdateOne is the builder for updating a single EmployeeAddress entity.
type EmployeeAddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmployeeAddressMutation
}

// SetEmployeeID sets the "employee_id" field.
func (eauo *EmployeeAddressUpdateOne) SetEmployeeID(i int) *EmployeeAddressUpdateOne {
	eauo.mutation.SetEmployeeID(i)
	return eauo
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (eauo *EmployeeAddressUpdateOne) SetNillableEmployeeID(i *int) *EmployeeAddressUpdateOne {
	if i != nil {
		eauo.SetEmployeeID(*i)
	}
	return eauo
}

// SetAddressID sets the "address_id" field.
func (eauo *EmployeeAddressUpdateOne) SetAddressID(i int) *EmployeeAddressUpdateOne {
	eauo.mutation.SetAddressID(i)
	return eauo
}

// SetNillableAddressID sets the "address_id" field if the given value is not nil.
func (eauo *EmployeeAddressUpdateOne) SetNillableAddressID(i *int) *EmployeeAddressUpdateOne {
	if i != nil {
		eauo.SetAddressID(*i)
	}
	return eauo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (eauo *EmployeeAddressUpdateOne) SetCompanyID(id int) *EmployeeAddressUpdateOne {
	eauo.mutation.SetCompanyID(id)
	return eauo
}

// SetCompany sets the "company" edge to the Company entity.
func (eauo *EmployeeAddressUpdateOne) SetCompany(c *Company) *EmployeeAddressUpdateOne {
	return eauo.SetCompanyID(c.ID)
}

// SetAddress sets the "address" edge to the Address entity.
func (eauo *EmployeeAddressUpdateOne) SetAddress(a *Address) *EmployeeAddressUpdateOne {
	return eauo.SetAddressID(a.ID)
}

// Mutation returns the EmployeeAddressMutation object of the builder.
func (eauo *EmployeeAddressUpdateOne) Mutation() *EmployeeAddressMutation {
	return eauo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (eauo *EmployeeAddressUpdateOne) ClearCompany() *EmployeeAddressUpdateOne {
	eauo.mutation.ClearCompany()
	return eauo
}

// ClearAddress clears the "address" edge to the Address entity.
func (eauo *EmployeeAddressUpdateOne) ClearAddress() *EmployeeAddressUpdateOne {
	eauo.mutation.ClearAddress()
	return eauo
}

// Where appends a list predicates to the EmployeeAddressUpdate builder.
func (eauo *EmployeeAddressUpdateOne) Where(ps ...predicate.EmployeeAddress) *EmployeeAddressUpdateOne {
	eauo.mutation.Where(ps...)
	return eauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eauo *EmployeeAddressUpdateOne) Select(field string, fields ...string) *EmployeeAddressUpdateOne {
	eauo.fields = append([]string{field}, fields...)
	return eauo
}

// Save executes the query and returns the updated EmployeeAddress entity.
func (eauo *EmployeeAddressUpdateOne) Save(ctx context.Context) (*EmployeeAddress, error) {
	return withHooks(ctx, eauo.sqlSave, eauo.mutation, eauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eauo *EmployeeAddressUpdateOne) SaveX(ctx context.Context) *EmployeeAddress {
	node, err := eauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eauo *EmployeeAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := eauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eauo *EmployeeAddressUpdateOne) ExecX(ctx context.Context) {
	if err := eauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eauo *EmployeeAddressUpdateOne) check() error {
	if eauo.mutation.CompanyCleared() && len(eauo.mutation.CompanyIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EmployeeAddress.company"`)
	}
	if eauo.mutation.AddressCleared() && len(eauo.mutation.AddressIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EmployeeAddress.address"`)
	}
	return nil
}

func (eauo *EmployeeAddressUpdateOne) sqlSave(ctx context.Context) (_node *EmployeeAddress, err error) {
	if err := eauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(employeeaddress.Table, employeeaddress.Columns, sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt))
	id, ok := eauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmployeeAddress.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employeeaddress.FieldID)
		for _, f := range fields {
			if !employeeaddress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != employeeaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if eauo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   employeeaddress.CompanyTable,
			Columns: []string{employeeaddress.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eauo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   employeeaddress.CompanyTable,
			Columns: []string{employeeaddress.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eauo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   employeeaddress.AddressTable,
			Columns: []string{employeeaddress.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eauo.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   employeeaddress.AddressTable,
			Columns: []string{employeeaddress.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EmployeeAddress{config: eauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employeeaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eauo.mutation.done = true
	return _node, nil
}
