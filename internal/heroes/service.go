package heroes

type HeroesService interface {
	GetHeroes() ([]byte, error)
}
