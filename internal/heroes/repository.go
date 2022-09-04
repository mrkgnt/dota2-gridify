package heroes

type HeroesRepository interface {
	GetHeroes() ([]byte, error)
}
