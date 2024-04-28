package tools

import (
	"testing"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func TestParasToml(t *testing.T) {
	print(Conf.MySQLServer.Database)
}
