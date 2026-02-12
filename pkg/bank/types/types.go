package types

// Money — тип для представления денег в копейках/центах
type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	id       int
	pan      PAN
	Balance  Money
	currency Currency
	color    string
	name     string
	Active   bool
}
