// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/bug/ent/address"
	"entgo.io/bug/ent/employee"
	"entgo.io/bug/ent/employeeaddress"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeAddressQuery is the builder for querying EmployeeAddress entities.
type EmployeeAddressQuery struct {
	config
	ctx          *QueryContext
	order        []employeeaddress.OrderOption
	inters       []Interceptor
	predicates   []predicate.EmployeeAddress
	withEmployee *EmployeeQuery
	withAddress  *AddressQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmployeeAddressQuery builder.
func (eaq *EmployeeAddressQuery) Where(ps ...predicate.EmployeeAddress) *EmployeeAddressQuery {
	eaq.predicates = append(eaq.predicates, ps...)
	return eaq
}

// Limit the number of records to be returned by this query.
func (eaq *EmployeeAddressQuery) Limit(limit int) *EmployeeAddressQuery {
	eaq.ctx.Limit = &limit
	return eaq
}

// Offset to start from.
func (eaq *EmployeeAddressQuery) Offset(offset int) *EmployeeAddressQuery {
	eaq.ctx.Offset = &offset
	return eaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eaq *EmployeeAddressQuery) Unique(unique bool) *EmployeeAddressQuery {
	eaq.ctx.Unique = &unique
	return eaq
}

// Order specifies how the records should be ordered.
func (eaq *EmployeeAddressQuery) Order(o ...employeeaddress.OrderOption) *EmployeeAddressQuery {
	eaq.order = append(eaq.order, o...)
	return eaq
}

// QueryEmployee chains the current query on the "employee" edge.
func (eaq *EmployeeAddressQuery) QueryEmployee() *EmployeeQuery {
	query := (&EmployeeClient{config: eaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(employeeaddress.Table, employeeaddress.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, employeeaddress.EmployeeTable, employeeaddress.EmployeeColumn),
		)
		fromU = sqlgraph.SetNeighbors(eaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAddress chains the current query on the "address" edge.
func (eaq *EmployeeAddressQuery) QueryAddress() *AddressQuery {
	query := (&AddressClient{config: eaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(employeeaddress.Table, employeeaddress.FieldID, selector),
			sqlgraph.To(address.Table, address.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, employeeaddress.AddressTable, employeeaddress.AddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(eaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EmployeeAddress entity from the query.
// Returns a *NotFoundError when no EmployeeAddress was found.
func (eaq *EmployeeAddressQuery) First(ctx context.Context) (*EmployeeAddress, error) {
	nodes, err := eaq.Limit(1).All(setContextOp(ctx, eaq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{employeeaddress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eaq *EmployeeAddressQuery) FirstX(ctx context.Context) *EmployeeAddress {
	node, err := eaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmployeeAddress ID from the query.
// Returns a *NotFoundError when no EmployeeAddress ID was found.
func (eaq *EmployeeAddressQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eaq.Limit(1).IDs(setContextOp(ctx, eaq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{employeeaddress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eaq *EmployeeAddressQuery) FirstIDX(ctx context.Context) int {
	id, err := eaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmployeeAddress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmployeeAddress entity is found.
// Returns a *NotFoundError when no EmployeeAddress entities are found.
func (eaq *EmployeeAddressQuery) Only(ctx context.Context) (*EmployeeAddress, error) {
	nodes, err := eaq.Limit(2).All(setContextOp(ctx, eaq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{employeeaddress.Label}
	default:
		return nil, &NotSingularError{employeeaddress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eaq *EmployeeAddressQuery) OnlyX(ctx context.Context) *EmployeeAddress {
	node, err := eaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmployeeAddress ID in the query.
// Returns a *NotSingularError when more than one EmployeeAddress ID is found.
// Returns a *NotFoundError when no entities are found.
func (eaq *EmployeeAddressQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eaq.Limit(2).IDs(setContextOp(ctx, eaq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{employeeaddress.Label}
	default:
		err = &NotSingularError{employeeaddress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eaq *EmployeeAddressQuery) OnlyIDX(ctx context.Context) int {
	id, err := eaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmployeeAddresses.
func (eaq *EmployeeAddressQuery) All(ctx context.Context) ([]*EmployeeAddress, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryAll)
	if err := eaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmployeeAddress, *EmployeeAddressQuery]()
	return withInterceptors[[]*EmployeeAddress](ctx, eaq, qr, eaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eaq *EmployeeAddressQuery) AllX(ctx context.Context) []*EmployeeAddress {
	nodes, err := eaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmployeeAddress IDs.
func (eaq *EmployeeAddressQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eaq.ctx.Unique == nil && eaq.path != nil {
		eaq.Unique(true)
	}
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryIDs)
	if err = eaq.Select(employeeaddress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eaq *EmployeeAddressQuery) IDsX(ctx context.Context) []int {
	ids, err := eaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eaq *EmployeeAddressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryCount)
	if err := eaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eaq, querierCount[*EmployeeAddressQuery](), eaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eaq *EmployeeAddressQuery) CountX(ctx context.Context) int {
	count, err := eaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eaq *EmployeeAddressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryExist)
	switch _, err := eaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eaq *EmployeeAddressQuery) ExistX(ctx context.Context) bool {
	exist, err := eaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmployeeAddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eaq *EmployeeAddressQuery) Clone() *EmployeeAddressQuery {
	if eaq == nil {
		return nil
	}
	return &EmployeeAddressQuery{
		config:       eaq.config,
		ctx:          eaq.ctx.Clone(),
		order:        append([]employeeaddress.OrderOption{}, eaq.order...),
		inters:       append([]Interceptor{}, eaq.inters...),
		predicates:   append([]predicate.EmployeeAddress{}, eaq.predicates...),
		withEmployee: eaq.withEmployee.Clone(),
		withAddress:  eaq.withAddress.Clone(),
		// clone intermediate query.
		sql:  eaq.sql.Clone(),
		path: eaq.path,
	}
}

// WithEmployee tells the query-builder to eager-load the nodes that are connected to
// the "employee" edge. The optional arguments are used to configure the query builder of the edge.
func (eaq *EmployeeAddressQuery) WithEmployee(opts ...func(*EmployeeQuery)) *EmployeeAddressQuery {
	query := (&EmployeeClient{config: eaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eaq.withEmployee = query
	return eaq
}

// WithAddress tells the query-builder to eager-load the nodes that are connected to
// the "address" edge. The optional arguments are used to configure the query builder of the edge.
func (eaq *EmployeeAddressQuery) WithAddress(opts ...func(*AddressQuery)) *EmployeeAddressQuery {
	query := (&AddressClient{config: eaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eaq.withAddress = query
	return eaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EmployeeID int `json:"employee_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmployeeAddress.Query().
//		GroupBy(employeeaddress.FieldEmployeeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eaq *EmployeeAddressQuery) GroupBy(field string, fields ...string) *EmployeeAddressGroupBy {
	eaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmployeeAddressGroupBy{build: eaq}
	grbuild.flds = &eaq.ctx.Fields
	grbuild.label = employeeaddress.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EmployeeID int `json:"employee_id,omitempty"`
//	}
//
//	client.EmployeeAddress.Query().
//		Select(employeeaddress.FieldEmployeeID).
//		Scan(ctx, &v)
func (eaq *EmployeeAddressQuery) Select(fields ...string) *EmployeeAddressSelect {
	eaq.ctx.Fields = append(eaq.ctx.Fields, fields...)
	sbuild := &EmployeeAddressSelect{EmployeeAddressQuery: eaq}
	sbuild.label = employeeaddress.Label
	sbuild.flds, sbuild.scan = &eaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmployeeAddressSelect configured with the given aggregations.
func (eaq *EmployeeAddressQuery) Aggregate(fns ...AggregateFunc) *EmployeeAddressSelect {
	return eaq.Select().Aggregate(fns...)
}

func (eaq *EmployeeAddressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eaq); err != nil {
				return err
			}
		}
	}
	for _, f := range eaq.ctx.Fields {
		if !employeeaddress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eaq.path != nil {
		prev, err := eaq.path(ctx)
		if err != nil {
			return err
		}
		eaq.sql = prev
	}
	return nil
}

func (eaq *EmployeeAddressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmployeeAddress, error) {
	var (
		nodes       = []*EmployeeAddress{}
		_spec       = eaq.querySpec()
		loadedTypes = [2]bool{
			eaq.withEmployee != nil,
			eaq.withAddress != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmployeeAddress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmployeeAddress{config: eaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eaq.withEmployee; query != nil {
		if err := eaq.loadEmployee(ctx, query, nodes, nil,
			func(n *EmployeeAddress, e *Employee) { n.Edges.Employee = e }); err != nil {
			return nil, err
		}
	}
	if query := eaq.withAddress; query != nil {
		if err := eaq.loadAddress(ctx, query, nodes, nil,
			func(n *EmployeeAddress, e *Address) { n.Edges.Address = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eaq *EmployeeAddressQuery) loadEmployee(ctx context.Context, query *EmployeeQuery, nodes []*EmployeeAddress, init func(*EmployeeAddress), assign func(*EmployeeAddress, *Employee)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EmployeeAddress)
	for i := range nodes {
		fk := nodes[i].EmployeeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(employee.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "employee_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eaq *EmployeeAddressQuery) loadAddress(ctx context.Context, query *AddressQuery, nodes []*EmployeeAddress, init func(*EmployeeAddress), assign func(*EmployeeAddress, *Address)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EmployeeAddress)
	for i := range nodes {
		fk := nodes[i].AddressID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(address.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "address_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eaq *EmployeeAddressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eaq.querySpec()
	_spec.Node.Columns = eaq.ctx.Fields
	if len(eaq.ctx.Fields) > 0 {
		_spec.Unique = eaq.ctx.Unique != nil && *eaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eaq.driver, _spec)
}

func (eaq *EmployeeAddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(employeeaddress.Table, employeeaddress.Columns, sqlgraph.NewFieldSpec(employeeaddress.FieldID, field.TypeInt))
	_spec.From = eaq.sql
	if unique := eaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eaq.path != nil {
		_spec.Unique = true
	}
	if fields := eaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employeeaddress.FieldID)
		for i := range fields {
			if fields[i] != employeeaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if eaq.withEmployee != nil {
			_spec.Node.AddColumnOnce(employeeaddress.FieldEmployeeID)
		}
		if eaq.withAddress != nil {
			_spec.Node.AddColumnOnce(employeeaddress.FieldAddressID)
		}
	}
	if ps := eaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eaq *EmployeeAddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eaq.driver.Dialect())
	t1 := builder.Table(employeeaddress.Table)
	columns := eaq.ctx.Fields
	if len(columns) == 0 {
		columns = employeeaddress.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eaq.sql != nil {
		selector = eaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eaq.ctx.Unique != nil && *eaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eaq.predicates {
		p(selector)
	}
	for _, p := range eaq.order {
		p(selector)
	}
	if offset := eaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmployeeAddressGroupBy is the group-by builder for EmployeeAddress entities.
type EmployeeAddressGroupBy struct {
	selector
	build *EmployeeAddressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eagb *EmployeeAddressGroupBy) Aggregate(fns ...AggregateFunc) *EmployeeAddressGroupBy {
	eagb.fns = append(eagb.fns, fns...)
	return eagb
}

// Scan applies the selector query and scans the result into the given value.
func (eagb *EmployeeAddressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eagb.build.ctx, ent.OpQueryGroupBy)
	if err := eagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeAddressQuery, *EmployeeAddressGroupBy](ctx, eagb.build, eagb, eagb.build.inters, v)
}

func (eagb *EmployeeAddressGroupBy) sqlScan(ctx context.Context, root *EmployeeAddressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eagb.fns))
	for _, fn := range eagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eagb.flds)+len(eagb.fns))
		for _, f := range *eagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmployeeAddressSelect is the builder for selecting fields of EmployeeAddress entities.
type EmployeeAddressSelect struct {
	*EmployeeAddressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eas *EmployeeAddressSelect) Aggregate(fns ...AggregateFunc) *EmployeeAddressSelect {
	eas.fns = append(eas.fns, fns...)
	return eas
}

// Scan applies the selector query and scans the result into the given value.
func (eas *EmployeeAddressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eas.ctx, ent.OpQuerySelect)
	if err := eas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeAddressQuery, *EmployeeAddressSelect](ctx, eas.EmployeeAddressQuery, eas, eas.inters, v)
}

func (eas *EmployeeAddressSelect) sqlScan(ctx context.Context, root *EmployeeAddressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eas.fns))
	for _, fn := range eas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
