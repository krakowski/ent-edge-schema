// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/address"
	"entgo.io/bug/ent/company"
	"entgo.io/bug/ent/employeeaddress"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeAddressCreate is the builder for creating a EmployeeAddress entity.
type EmployeeAddressCreate struct {
	config
	mutation *EmployeeAddressMutation
	hooks    []Hook
}

// SetEmployeeID sets the "employee_id" field.
func (eac *EmployeeAddressCreate) SetEmployeeID(i int) *EmployeeAddressCreate {
	eac.mutation.SetEmployeeID(i)
	return eac
}

// SetAddressID sets the "address_id" field.
func (eac *EmployeeAddressCreate) SetAddressID(i int) *EmployeeAddressCreate {
	eac.mutation.SetAddressID(i)
	return eac
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (eac *EmployeeAddressCreate) SetCompanyID(id int) *EmployeeAddressCreate {
	eac.mutation.SetCompanyID(id)
	return eac
}

// SetCompany sets the "company" edge to the Company entity.
func (eac *EmployeeAddressCreate) SetCompany(c *Company) *EmployeeAddressCreate {
	return eac.SetCompanyID(c.ID)
}

// SetAddress sets the "address" edge to the Address entity.
func (eac *EmployeeAddressCreate) SetAddress(a *Address) *EmployeeAddressCreate {
	return eac.SetAddressID(a.ID)
}

// Mutation returns the EmployeeAddressMutation object of the builder.
func (eac *EmployeeAddressCreate) Mutation() *EmployeeAddressMutation {
	return eac.mutation
}

// Save creates the EmployeeAddress in the database.
func (eac *EmployeeAddressCreate) Save(ctx context.Context) (*EmployeeAddress, error) {
	return withHooks(ctx, eac.sqlSave, eac.mutation, eac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eac *EmployeeAddressCreate) SaveX(ctx context.Context) *EmployeeAddress {
	v, err := eac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eac *EmployeeAddressCreate) Exec(ctx context.Context) error {
	_, err := eac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eac *EmployeeAddressCreate) ExecX(ctx context.Context) {
	if err := eac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eac *EmployeeAddressCreate) check() error {
	if _, ok := eac.mutation.EmployeeID(); !ok {
		return &ValidationError{Name: "employee_id", err: errors.New(`ent: missing required field "EmployeeAddress.employee_id"`)}
	}
	if _, ok := eac.mutation.AddressID(); !ok {
		return &ValidationError{Name: "address_id", err: errors.New(`ent: missing required field "EmployeeAddress.address_id"`)}
	}
	if len(eac.mutation.CompanyIDs()) == 0 {
		return &ValidationError{Name: "company", err: errors.New(`ent: missing required edge "EmployeeAddress.company"`)}
	}
	if len(eac.mutation.AddressIDs()) == 0 {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required edge "EmployeeAddress.address"`)}
	}
	return nil
}

func (eac *EmployeeAddressCreate) sqlSave(ctx context.Context) (*EmployeeAddress, error) {
	if err := eac.check(); err != nil {
		return nil, err
	}
	_node, _spec := eac.createSpec()
	if err := sqlgraph.CreateNode(ctx, eac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	eac.mutation.id = &_node.ID
	eac.mutation.done = true
	return _node, nil
}

func (eac *EmployeeAddressCreate) createSpec() (*EmployeeAddress, *sqlgraph.CreateSpec) {
	var (
		_node = &EmployeeAddress{config: eac.config}
		_spec = sqlgraph.NewCreateSpec(employeeaddress.Table, sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt))
	)
	if nodes := eac.mutation.CompanyIDs(); len(nodes) > 0 {
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
		_node.EmployeeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := eac.mutation.AddressIDs(); len(nodes) > 0 {
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
		_node.AddressID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EmployeeAddressCreateBulk is the builder for creating many EmployeeAddress entities in bulk.
type EmployeeAddressCreateBulk struct {
	config
	err      error
	builders []*EmployeeAddressCreate
}

// Save creates the EmployeeAddress entities in the database.
func (eacb *EmployeeAddressCreateBulk) Save(ctx context.Context) ([]*EmployeeAddress, error) {
	if eacb.err != nil {
		return nil, eacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(eacb.builders))
	nodes := make([]*EmployeeAddress, len(eacb.builders))
	mutators := make([]Mutator, len(eacb.builders))
	for i := range eacb.builders {
		func(i int, root context.Context) {
			builder := eacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmployeeAddressMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, eacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eacb *EmployeeAddressCreateBulk) SaveX(ctx context.Context) []*EmployeeAddress {
	v, err := eacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eacb *EmployeeAddressCreateBulk) Exec(ctx context.Context) error {
	_, err := eacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eacb *EmployeeAddressCreateBulk) ExecX(ctx context.Context) {
	if err := eacb.Exec(ctx); err != nil {
		panic(err)
	}
}
