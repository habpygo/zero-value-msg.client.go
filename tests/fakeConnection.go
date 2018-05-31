package communications

import (
	"errors"
	"time"

	"github.com/iotaledger/giota"
)

//NewConnection establishes a connection with the given provider and the seed
func NewFakeConnection(provider, seed string) (*FakeConnection, error) {

	return &FakeConnection{
		capsules: make(map[giota.Address][]giota.Transaction),
	}, nil
}

type FakeConnection struct {
	capsules map[giota.Address][]giota.Transaction
}

func (c *FakeConnection) SendToApi(trs []giota.Transfer) (giota.Bundle, error) {
	for _, t := range trs {
		_, found := c.capsules[t.Address]
		if !found {
			c.capsules[t.Address] = []giota.Transaction{}
		}
		c.capsules[t.Address] = append(c.capsules[t.Address], giota.Transaction{
			Address: t.Address,
			Value:   t.Value,
			SignatureMessageFragment: t.Message,
			Timestamp:                time.Now(),
		})
	}

	return giota.Bundle{giota.Transaction{}}, nil
}

func (c *FakeConnection) FindTransactions(req giota.FindTransactionsRequest) ([]giota.Transaction, error) {

	res := []giota.Transaction{}
	for _, a := range req.Addresses {
		ts, found := c.capsules[a]

		if !found {
			continue
		}

		for _, t := range ts {
			res = append(res, t)
		}
	}

	return res, nil
}

func (c *FakeConnection) ReadTransactions(tIDs []giota.Trytes) ([]giota.Transaction, error) {
	return nil, errors.New("Not implemented yet")
}
