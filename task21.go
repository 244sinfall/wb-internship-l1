package main

import (
	"fmt"
)

// Введем общий интерфейс игрока
type Player interface {
	Attack()
	Walk(meters int)
	Introduce()
	Sleep(hours int)
}

// Создадим вполне конкретную структуру, соответствующую интерфейсу Player.
type Hero struct {
	name                  string
	yearsOnTheBattlefield int
}

func (h *Hero) Attack() {
	fmt.Printf("%s атакует врага голыми руками!\n", h.name)
}
func (h *Hero) Walk(meters int) {
	fmt.Printf("Стряхивая пот со лба, %s проходит %d метров...\n", h.name, meters)
}
func (h *Hero) Introduce() {
	fmt.Printf("Мое имя - %s. Я прибыл в Средиземье %d лет назад...\n", h.name, h.yearsOnTheBattlefield)
}
func (h *Hero) Sleep(hours int) {
	fmt.Printf("%s стелит лежак на полу и засыпает на %d часов...\n", h.name, hours)
}

// Предположим, что герой может стать волшебником . Это - другой интерфейс, схожий по смыслу действия,
// но различные в реализации. Реализуем их...
type Wizard interface {
	Fireball()
	Teleport(kilometers int)
	TellTheyWontPass()
	GoToAstral(years int)
}
type Magician struct {
	name string
}

func (m *Magician) Fireball() {
	fmt.Printf("%s атакует врага огненным шаром!\n", m.name)
}
func (m *Magician) Teleport(kilometers float64) {
	fmt.Printf("%s телепортируется на %f километров!\n", m.name, kilometers)
}
func (m *Magician) TellTheyWontPass() {
	fmt.Printf("%s говорит окружающим, что они <не пройдут>!\n", m.name)
}
func (m *Magician) GoToAstral(years float64) {
	fmt.Printf("%s уходит в астрал на %f лет!\n", m.name, years)
}

type WizardHeroAdapter struct {
	Magician
}

func (w *WizardHeroAdapter) Attack() {
	w.Fireball()
}
func (w *WizardHeroAdapter) Walk(meters int) {
	w.Teleport(float64(meters) / 1000.0)
}
func (w *WizardHeroAdapter) Introduce() {
	w.TellTheyWontPass()
}
func (w *WizardHeroAdapter) Sleep(hours int) {
	w.GoToAstral(float64(hours) / 8766.0) // Загуглил сколько часов в году)
}

// Теперь предположим, что у нас есть функция, которая позволяет игроку играть
func play(p Player) {
	p.Introduce()
	p.Walk(500)
	p.Attack()
	p.Sleep(10)
}

// Теперь запустим все это дело...
func task21() {
	hero1 := Hero{
		name:                  "Dima",
		yearsOnTheBattlefield: 5,
	}
	hero2 := Magician{name: "Gandalf"}
	play(&hero1) // С Димой все окей, поскольку он соответствует интерфейсу Player
	//play(&hero2) // С Гендальфом нет, потому что он не реализует методы интерфейса Player, но реализует те, что очень
	// похожи на них. Мы написали адаптер, который реализует методы Player используя методы Wizard
	realHero2 := WizardHeroAdapter{hero2}
	play(&realHero2) // С адаптером все работает!
}
