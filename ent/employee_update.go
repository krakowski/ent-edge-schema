// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/address"
	"entgo.io/bug/ent/employee"
	"entgo.io/bug/ent/employeeaddress"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeUpdate is the builder for updating Employee entities.
type EmployeeUpdate struct {
	config
	hooks    []Hook
	mutation *EmployeeMutation
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (eu *EmployeeUpdate) Where(ps ...predicate.Employee) *EmployeeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetFirstname sets the "firstname" field.
func (eu *EmployeeUpdate) SetFirstname(s string) *EmployeeUpdate {
	eu.mutation.SetFirstname(s)
	return eu
}

// SetNillableFirstname sets the "firstname" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableFirstname(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetFirstname(*s)
	}
	return eu
}

// SetLastname sets the "lastname" field.
func (eu *EmployeeUpdate) SetLastname(s string) *EmployeeUpdate {
	eu.mutation.SetLastname(s)
	return eu
}

// SetNillableLastname sets the "lastname" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableLastname(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetLastname(*s)
	}
	return eu
}

// AddAddresIDs adds the "address" edge to the Address entity by IDs.
func (eu *EmployeeUpdate) AddAddresIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.AddAddresIDs(ids...)
	return eu
}

// AddAddress adds the "address" edges to the Address entity.
func (eu *EmployeeUpdate) AddAddress(a ...*Address) *EmployeeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return eu.AddAddresIDs(ids...)
}

// AddEmployeeAddressIDs adds the "employee_addresses" edge to the EmployeeAddress entity by IDs.
func (eu *EmployeeUpdate) AddEmployeeAddressIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.AddEmployeeAddressIDs(ids...)
	return eu
}

// AddEmployeeAddresses adds the "employee_addresses" edges to the EmployeeAddress entity.
func (eu *EmployeeUpdate) AddEmployeeAddresses(e ...*EmployeeAddress) *EmployeeUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eu.AddEmployeeAddressIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (eu *EmployeeUpdate) Mutation() *EmployeeMutation {
	return eu.mutation
}

// ClearAddress clears all "address" edges to the Address entity.
func (eu *EmployeeUpdate) ClearAddress() *EmployeeUpdate {
	eu.mutation.ClearAddress()
	return eu
}

// RemoveAddresIDs removes the "address" edge to Address entities by IDs.
func (eu *EmployeeUpdate) RemoveAddresIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.RemoveAddresIDs(ids...)
	return eu
}

// RemoveAddress removes "address" edges to Address entities.
func (eu *EmployeeUpdate) RemoveAddress(a ...*Address) *EmployeeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return eu.RemoveAddresIDs(ids...)
}

// ClearEmployeeAddresses clears all "employee_addresses" edges to the EmployeeAddress entity.
func (eu *EmployeeUpdate) ClearEmployeeAddresses() *EmployeeUpdate {
	eu.mutation.ClearEmployeeAddresses()
	return eu
}

// RemoveEmployeeAddressIDs removes the "employee_addresses" edge to EmployeeAddress entities by IDs.
func (eu *EmployeeUpdate) RemoveEmployeeAddressIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.RemoveEmployeeAddressIDs(ids...)
	return eu
}

// RemoveEmployeeAddresses removes "employee_addresses" edges to EmployeeAddress entities.
func (eu *EmployeeUpdate) RemoveEmployeeAddresses(e ...*EmployeeAddress) *EmployeeUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eu.RemoveEmployeeAddressIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmployeeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EmployeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Firstname(); ok {
		_spec.SetField(employee.FieldFirstname, field.TypeString, value)
	}
	if value, ok := eu.mutation.Lastname(); ok {
		_spec.SetField(employee.FieldLastname, field.TypeString, value)
	}
	if eu.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   employee.AddressTable,
			Columns: employee.AddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedAddressIDs(); len(nodes) > 0 && !eu.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   employee.AddressTable,
			Columns: employee.AddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   employee.AddressTable,
			Columns: employee.AddressPrimaryKey,
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
	if eu.mutation.EmployeeAddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   employee.EmployeeAddressesTable,
			Columns: []string{employee.EmployeeAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedEmployeeAddressesIDs(); len(nodes) > 0 && !eu.mutation.EmployeeAddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   employee.EmployeeAddressesTable,
			Columns: []string{employee.EmployeeAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.EmployeeAddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   employee.EmployeeAddressesTable,
			Columns: []string{employee.EmployeeAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmployeeUpdateOne is the builder for updating a single Employee entity.
type EmployeeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmployeeMutation
}

// SetFirstname sets the "firstname" field.
func (euo *EmployeeUpdateOne) SetFirstname(s string) *EmployeeUpdateOne {
	euo.mutation.SetFirstname(s)
	return euo
}

// SetNillableFirstname sets the "firstname" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableFirstname(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetFirstname(*s)
	}
	return euo
}

// SetLastname sets the "lastname" field.
func (euo *EmployeeUpdateOne) SetLastname(s string) *EmployeeUpdateOne {
	euo.mutation.SetLastname(s)
	return euo
}

// SetNillableLastname sets the "lastname" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableLastname(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetLastname(*s)
	}
	return euo
}

// AddAddresIDs adds the "address" edge to the Address entity by IDs.
func (euo *EmployeeUpdateOne) AddAddresIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.AddAddresIDs(ids...)
	return euo
}

// AddAddress adds the "address" edges to the Address entity.
func (euo *EmployeeUpdateOne) AddAddress(a ...*Address) *EmployeeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return euo.AddAddresIDs(ids...)
}

// AddEmployeeAddressIDs adds the "employee_addresses" edge to the EmployeeAddress entity by IDs.
func (euo *EmployeeUpdateOne) AddEmployeeAddressIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.AddEmployeeAddressIDs(ids...)
	return euo
}

// AddEmployeeAddresses adds the "employee_addresses" edges to the EmployeeAddress entity.
func (euo *EmployeeUpdateOne) AddEmployeeAddresses(e ...*EmployeeAddress) *EmployeeUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return euo.AddEmployeeAddressIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (euo *EmployeeUpdateOne) Mutation() *EmployeeMutation {
	return euo.mutation
}

// ClearAddress clears all "address" edges to the Address entity.
func (euo *EmployeeUpdateOne) ClearAddress() *EmployeeUpdateOne {
	euo.mutation.ClearAddress()
	return euo
}

// RemoveAddresIDs removes the "address" edge to Address entities by IDs.
func (euo *EmployeeUpdateOne) RemoveAddresIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.RemoveAddresIDs(ids...)
	return euo
}

// RemoveAddress removes "address" edges to Address entities.
func (euo *EmployeeUpdateOne) RemoveAddress(a ...*Address) *EmployeeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return euo.RemoveAddresIDs(ids...)
}

// ClearEmployeeAddresses clears all "employee_addresses" edges to the EmployeeAddress entity.
func (euo *EmployeeUpdateOne) ClearEmployeeAddresses() *EmployeeUpdateOne {
	euo.mutation.ClearEmployeeAddresses()
	return euo
}

// RemoveEmployeeAddressIDs removes the "employee_addresses" edge to EmployeeAddress entities by IDs.
func (euo *EmployeeUpdateOne) RemoveEmployeeAddressIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.RemoveEmployeeAddressIDs(ids...)
	return euo
}

// RemoveEmployeeAddresses removes "employee_addresses" edges to EmployeeAddress entities.
func (euo *EmployeeUpdateOne) RemoveEmployeeAddresses(e ...*EmployeeAddress) *EmployeeUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return euo.RemoveEmployeeAddressIDs(ids...)
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (euo *EmployeeUpdateOne) Where(ps ...predicate.Employee) *EmployeeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmployeeUpdateOne) Select(field string, fields ...string) *EmployeeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Employee entity.
func (euo *EmployeeUpdateOne) Save(ctx context.Context) (*Employee, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeUpdateOne) SaveX(ctx context.Context) *Employee {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmployeeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EmployeeUpdateOne) sqlSave(ctx context.Context) (_node *Employee, err error) {
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Employee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employee.FieldID)
		for _, f := range fields {
			if !employee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != employee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Firstname(); ok {
		_spec.SetField(employee.FieldFirstname, field.TypeString, value)
	}
	if value, ok := euo.mutation.Lastname(); ok {
		_spec.SetField(employee.FieldLastname, field.TypeString, value)
	}
	if euo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   employee.AddressTable,
			Columns: employee.AddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedAddressIDs(); len(nodes) > 0 && !euo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   employee.AddressTable,
			Columns: employee.AddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   employee.AddressTable,
			Columns: employee.AddressPrimaryKey,
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
	if euo.mutation.EmployeeAddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   employee.EmployeeAddressesTable,
			Columns: []string{employee.EmployeeAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedEmployeeAddressesIDs(); len(nodes) > 0 && !euo.mutation.EmployeeAddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   employee.EmployeeAddressesTable,
			Columns: []string{employee.EmployeeAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.EmployeeAddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   employee.EmployeeAddressesTable,
			Columns: []string{employee.EmployeeAddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Employee{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
