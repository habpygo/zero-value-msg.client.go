package communications

import (
	"fmt"
	"sort"
	"time"

	"github.com/habpygo/zero-value-msg.client.go/mamutils"

	"github.com/iotaledger/giota"
)

type Transaction struct {
	Message   string
	Value     int64
	Timestamp time.Time
	Recipient string
}

type ApiTransactionsFinder interface {
	FindTransactions(giota.FindTransactionsRequest) ([]giota.Transaction, error)
}

// ReadTransactions reads all the historic transaction data from a particular address
func ReadTransactions(address string, f ApiTransactionsFinder) ([]Transaction, error) {
	iotaAddress, err := giota.ToAddress(address)
	if err != nil {
		return nil, err
	}

	req := giota.FindTransactionsRequest{
		Addresses: []giota.Address{iotaAddress},
	}

	foundTx, err := f.FindTransactions(req)
	//fmt.Println("foundTx is: ", foundTx)
	if err != nil {
		return nil, err
	}

	sort.Slice(foundTx, func(i, j int) bool {
		return !(foundTx[i].Timestamp.Unix() < foundTx[j].Timestamp.Unix())
	})

	transactions := make([]Transaction, len(foundTx))
	for i, t := range foundTx {
		//message, err := mamutils.FromMAMTrytes(t.SignatureMessageFragment)
		message, err := mamutils.FromMAMTrytes(t.SignatureMessageFragment)
		if err != nil {
			return nil, err
		}
		transactions[i] = Transaction{
			Message:   message,
			Value:     t.Value,
			Timestamp: t.Timestamp,
			Recipient: string(t.Address),
		}
	}

	fmt.Println("Total number of transactions found on tangle are: ", len(transactions))
	return transactions, nil
}

type ApiTransactionsReader interface {
	ReadTransactions([]giota.Trytes) ([]giota.Transaction, error)
}

func ReadTransaction(transactionID string, r ApiTransactionsReader) (Transaction, error) {
	tID, err := giota.ToTrytes(transactionID)
	if err != nil {
		return Transaction{}, err
	}

	txs, err := r.ReadTransactions([]giota.Trytes{tID})
	if len(txs) != 1 {
		return Transaction{}, fmt.Errorf("Requested 1 Transaction but got %d", len(txs))
	}
	if err != nil {
		return Transaction{}, err
	}

	tx := txs[0]
	message, err := mamutils.FromMAMTrytes(tx.SignatureMessageFragment)
	if err != nil {
		return Transaction{}, err
	}
	transaction := Transaction{
		Message:   message,
		Value:     tx.Value,
		Timestamp: tx.Timestamp,
		Recipient: string(tx.Address),
	}

	return transaction, nil
}

// func ReadTransactions(address string, f ApiTransactionsFinder) ([]Transaction, error) {
// 	iotaAddress, err := giota.ToAddress(address)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("iotaAddress is: ", iotaAddress)

// 	req := giota.FindTransactionsRequest{
// 		Addresses: []giota.Address{iotaAddress},
// 	}
// 	fmt.Println("req.Addresses is: ", req.Addresses)
// 	fmt.Println("req.Approvees is: ", req.Approvees)
// 	fmt.Println("req.Bundles is: ", req.Bundles)
// 	fmt.Println("req.Command is: ", req.Command)
// 	fmt.Println("req.Tags is: ", req.Tags)

// 	foundTx, err := f.FindTransactions(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	sort.Slice(foundTx, func(i, j int) bool {
// 		return !(foundTx[i].Timestamp.Unix() < foundTx[j].Timestamp.Unix())
// 	})

// 	transactions := make([]Transaction, len(foundTx))
// 	for i, t := range foundTx {
// 		message, err := mamutils.FromMAMTrytes(t.SignatureMessageFragment)
// 		if err != nil {
// 			return nil, err
// 		}
// 		transactions[i] = Transaction{
// 			Message:   message,
// 			Value:     t.Value,
// 			Timestamp: t.Timestamp,
// 			Recipient: string(t.Address),
// 		}
// 	}

// 	return transactions, nil
// }
