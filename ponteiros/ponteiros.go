package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

func (c *Carteira) Sacar(valor Bitcoin) error {
	if valor > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= valor
	return nil

}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo

}

func main() {
	fmt.Println("Ponteiros")
}
