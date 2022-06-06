package investapi

import (
	"fmt"
	"math"
	"strings"
)

const nano = 1000000000

// ToFloat64 возвращает значения Quotation как число типа float64
func (x *Quotation) ToFloat64() (output float64) {
	return float64(x.Units) + float64(x.Nano)/nano
}

// MoneyValueFromFloat64 создаёт MoneyValue c пустой валютой из числа типа float64
func MoneyValueFromFloat64(a float64) *MoneyValue {
	i, f := math.Modf(a)
	return &MoneyValue{
		Units:    int64(i),
		Nano:     int32(math.Round(f * nano)),
		Currency: "",
	}
}

// Float64ToQuotation преобразует float64 в ссылку на тип Quotation
func Float64ToQuotation(input float64) (output *Quotation) {
	integerPart, floatPart := math.Modf(input)
	output = new(Quotation)
	output.Units = int64(integerPart)
	output.Nano = int32(math.Round(nano * floatPart))
	return
}

// ToStringWithoutCurrency выводит строковое представление цены без указания валюты
func (x *MoneyValue) ToStringWithoutCurrency() string {
	return fmt.Sprintf("%.2f", x.ToFloat64())
}

// ToFloat64 выводит представление цены в виде дробного числа float64
func (x *MoneyValue) ToFloat64() float64 {
	return float64(x.Units) + float64(x.Nano)/nano
}

// ToStringWithCurrency  выводит строковое представление цены с указанием валюты
func (x *MoneyValue) ToStringWithCurrency() string {
	return fmt.Sprintf("%.2f %s", x.ToFloat64(), strings.ToUpper(x.Currency))
}
