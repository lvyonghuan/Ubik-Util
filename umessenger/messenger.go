package umessenger

import (
	"bytes"
	"errors"
	"log"
	"net/http"

	"github.com/lvyonghuan/Ubik-Util/uerr"
	"github.com/lvyonghuan/Ubik-Util/ujson"
)

type UMessenger struct {
	addr string //address for delivery
	uuid string //UUID of the follower that sent the message. If is a leader, it is empty.
}

// NewUMessenger creates a new UMessenger instance with the given address.
func NewUMessenger(deliveryAddr, uuid string) *UMessenger {
	return &UMessenger{
		addr: deliveryAddr,
		uuid: uuid,
	}
}

// PostMessage sends a message asynchronously to the specified address without returning a response.
// It will use a goroutine to send the message.
func (messenger *UMessenger) PostMessage(envelope UEnvelope) {
	go func() {
		req, err := messenger.newRequest(envelope)
		if err != nil {
			log.Println("UMessenger: failed to make request: " + err.Error())
		}

		_, err = messenger.call(req)
		if err != nil {
			log.Println("UMessenger: failed to send request: " + err.Error())
			return
		}
	}()
}

func (messenger *UMessenger) newRequest(envelope UEnvelope) (*http.Request, error) {
	envelopeBytes, err := ujson.Marshal(envelope)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(envelopeBytes)

	req, err := http.NewRequest(envelope.methods, messenger.addr+"?UUID="+envelope.uuid, bodyReader)
	if err != nil {
		return nil, uerr.NewError(errors.New("failed to create new request: " + err.Error()))
	}

	return req, nil
}

func (messenger *UMessenger) call(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, uerr.NewError(errors.New("failed to send request: " + err.Error()))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, uerr.NewError(errors.New("received non-200 response: " + resp.Status))
	}

	return resp, nil
}
