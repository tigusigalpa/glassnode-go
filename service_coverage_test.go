package glassnode

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

var (
	contextType     = reflect.TypeOf((*context.Context)(nil)).Elem()
	metricQueryType = reflect.TypeOf((*MetricQuery)(nil))
	errorType       = reflect.TypeOf((*error)(nil)).Elem()
)

// TestCategoryServiceMethods exercises every exported convenience wrapper
// against a mock API. This verifies that wrappers remain wired to the shared
// transport while keeping endpoint-specific response handling centralized.
func TestCategoryServiceMethods(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[]`))
	})
	defer srv.Close()

	services := []interface{}{
		c.Addresses, c.Blockchain, c.Breakdowns, c.Bridges, c.DeFi,
		c.Derivatives, c.Distribution, c.Entities, c.ETH2, c.Fees,
		c.Global, c.Indicators, c.Institutions, c.Lightning, c.Macro,
		c.Market, c.Mempool, c.Mining, c.Options, c.Protocols, c.Signals,
		c.Supply, c.Transactions, c.Treasuries,
	}

	for _, service := range services {
		serviceValue := reflect.ValueOf(service)
		serviceType := serviceValue.Type()
		for i := 0; i < serviceValue.NumMethod(); i++ {
			method := serviceValue.Method(i)
			methodInfo := serviceType.Method(i)
			t.Run(serviceType.Elem().Name()+"/"+methodInfo.Name, func(t *testing.T) {
				args := coverageMethodArgs(t, method.Type())
				results := method.Call(args)
				last := results[len(results)-1]
				if !last.IsNil() {
					t.Fatalf("%s returned %v", methodInfo.Name, last.Interface())
				}
			})
		}
	}
}

func coverageMethodArgs(t *testing.T, methodType reflect.Type) []reflect.Value {
	t.Helper()
	args := make([]reflect.Value, methodType.NumIn())
	for i := range args {
		switch argType := methodType.In(i); {
		case argType == contextType:
			args[i] = reflect.ValueOf(context.Background())
		case argType == metricQueryType:
			args[i] = reflect.ValueOf(&MetricQuery{Asset: "BTC"})
		case argType.Kind() == reflect.String:
			args[i] = reflect.ValueOf("custom_metric")
		default:
			t.Fatalf("unsupported method argument %s", argType)
		}
	}
	if methodType.NumOut() == 0 || methodType.Out(methodType.NumOut()-1) != errorType {
		t.Fatalf("expected final error result in %s", methodType)
	}
	return args
}
