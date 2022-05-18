package investapi

import (
	"testing"
)

func TestFloat64ToQuotation(t *testing.T) {
	quot := Float64ToQuotation(0.82)
	t.Logf("Unit %v", quot.Units)
	t.Logf("Nano %v", quot.Nano)
}

func TestQuotationToFloat(t *testing.T) {
	q := new(Quotation)
	q.Nano = 820000000
	q.Units = 0
	t.Logf("Quotation %s", q.String())
	price := q.ToFloat64()
	t.Logf("Price %.5f", price)
}
