package bug

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/bug/ent"
)

func TestEdgeSchemaTypes(t *testing.T) {
	company := ent.Company{}
	if reflect.TypeOf(company.Edges.Address).Kind() == reflect.Slice {
		t.Error("edge type is slice")
	}

	employee := ent.Employee{}
	if reflect.TypeOf(employee.Edges.Address).Kind() == reflect.Slice {
		t.Error("edge type is slice")
	}
}
