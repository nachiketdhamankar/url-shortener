package json

import (
	"encoding/json"
	"github.com/pkg/errors"
	shortener2 "nachiket.me/url-shortener/shortener"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener2.Redirect, error) {
	redirect := &shortener2.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener2.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
