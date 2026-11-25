package main

import "fmt"

type Character interface {
	GetAbilities() string
}

type Mario struct{}

func (m *Mario) GetAbilities() string {
	return "Mario"
}

type HeightUp struct {
	character Character
}

func NewHeightUp(ch Character) *HeightUp {
	return &HeightUp{character: ch}
}

func (h *HeightUp) GetAbilities() string {
	return h.character.GetAbilities() + " with HeightUp"

}

type GunPowerUp struct {
	character Character
}

func NewGunPowerUp(ch Character) *GunPowerUp {
	return &GunPowerUp{character: ch}
}

func (g *GunPowerUp) GetAbilities() string {
	return g.character.GetAbilities() + " with Gun"
}

type StarPowerUp struct {
	character Character // Wrapped component
}

func NewStarPowerUp(c Character) *StarPowerUp {
	return &StarPowerUp{character: c}
}

func (s *StarPowerUp) GetAbilities() string {
	return s.character.GetAbilities() + " with Star Power (Limited Time)"
}

func main() {
	var mario Character = &Mario{}
	fmt.Println("Basic Character:", mario.GetAbilities())

	// Decorate Mario with a HeightUp power-up
	mario = NewHeightUp(mario)
	fmt.Println("After HeightUp:", mario.GetAbilities())

	// Decorate Mario further with a GunPowerUp
	mario = NewGunPowerUp(mario)
	fmt.Println("After GunPowerUp:", mario.GetAbilities())

	// Finally, add a StarPowerUp decoration
	mario = NewStarPowerUp(mario)
	fmt.Println("After StarPowerUp:", mario.GetAbilities())

}
