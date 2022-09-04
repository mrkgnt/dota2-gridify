package json

import (
	"encoding/json"

	"github.com/mrkgnt/dota2-gridify/internal/heroes"
	"github.com/pkg/errors"
)

type Heroes struct{}

func (h *Heroes) Decode(input []byte) (*heroes.Heroes, error) {
	heroes := &heroes.Heroes{}

	if err := json.Unmarshal(input, heroes); err != nil {
		return nil, errors.Wrap(err, "serializer.json.Heroes.Decode")
	}

	return heroes, nil
}

func (h *Heroes) Encode(input *heroes.Heroes) ([]byte, error) {
	msg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.json.Heroes.Encode")
	}

	return msg, nil
}
