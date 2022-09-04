package msgpack

import (
	"github.com/vmihailenco/msgpack"

	"github.com/mrkgnt/dota2-gridify/internal/heroes"
	"github.com/pkg/errors"
)

type Heroes struct{}

func (h *Heroes) Decode(input []byte) (*heroes.Heroes, error) {
	heroes := &heroes.Heroes{}

	if err := msgpack.Unmarshal(input, heroes); err != nil {
		return nil, errors.Wrap(err, "serializer.msgpack.Heroes.Decode")
	}

	return heroes, nil
}

func (h *Heroes) Encode(input *heroes.Heroes) ([]byte, error) {
	msg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.msgpack.Heroes.Encode")
	}

	return msg, nil
}
