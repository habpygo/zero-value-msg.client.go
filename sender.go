package communications

import (
	"github.com/giota"
	"github.com/habpygo/mam.client.go/mamutils"
)

type ApiSender interface {
	SendToApi([]giota.Transfer) (giota.Bundle, error)
}

func Send(recipient string, value int64, message string, sender ApiSender) (string, error) {
	address, err := giota.ToAddress(recipient)
	if err != nil {
		return "", err
	}

	encodedMessage, err := mamutils.ToMAMTrytes(message)
	if err != nil {
		return "", err
	}

	trs := []giota.Transfer{
		giota.Transfer{
			Address: address,
			Value:   value,
			Message: encodedMessage,
			Tag:     "",
		},
	}

	mamBundle, sendErr := sender.SendToApi(trs)
	if sendErr != nil {
		return "", sendErr
	}

	return string(mamBundle[0].Hash()), nil
}

func SendBatch(recipient string, value int64, messages []string, sender ApiSender) (string, error) {
	address, err := giota.ToAddress(recipient)
	if err != nil {
		return "", err
	}

	trs := []giota.Transfer{}
	for _, message := range messages {

		encodedMessage, err := mamutils.ToMAMTrytes(message)
		if err != nil {
			return "", err
		}

		t := giota.Transfer{
			Address: address,
			Value:   value,
			Message: encodedMessage,
			Tag:     "",
		}

		trs = append(trs, t)
	}

	mamBundle, sendErr := sender.SendToApi(trs)
	if sendErr != nil {
		return "", sendErr
	}

	return string(mamBundle[0].Hash()), nil
}
