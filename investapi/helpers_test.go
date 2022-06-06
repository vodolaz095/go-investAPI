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

func Test_MoneyValueFromFloat64(t *testing.T) {
	/*
		цена 100.02 для AAPL должна выставляться корректно.
		Проверьте, что она правильно преобразуется в формат MoneyValue,
		в котором дробная часть передается в миллиардных долях.
			100.02 = {units=100,nano=20000000}
	*/

	v := 100.02

	m := MoneyValueFromFloat64(v)
	t.Logf("%d %d ", m.Units, m.Nano)

	if m.Units != 100 || m.Nano != 20000000 {
		t.Fatal("цена не та")
	}
}
