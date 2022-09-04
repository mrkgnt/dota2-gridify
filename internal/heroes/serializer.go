package heroes

type HeroesSerializer interface {
	Decode(input []byte) (*Heroes, error)
	Encode(input *Heroes) ([]byte, error)
}
