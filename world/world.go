package world

import (
	"fmt"
	"strconv"
	"strings"
)

type World interface {
	MakeIteration()
	PrintCities()
	HasAliveAliens() bool
}

type worldImp struct {
	Aliens       map[AlienName]*Alien
	WorldMap     *Map
	// Distribution represents locations of aliens with key of city names.
	// We need distribution to make fast search through aliens in the same city.
	Distribution map[string][]*Alien
}

// NewWorld creates new world with provided city map and number of aliens.
// It generates n random aliens and make fight step for aliens that were created in the same city.
func NewWorld(worldMap *Map, n int) World {
	aliens := createRandomAliens(worldMap, n)

	// Initialize distribution for aliens.
	distribution := make(map[string][]*Alien)
	wi := &worldImp{
		Aliens:       aliens,
		WorldMap:     worldMap,
		Distribution: distribution,
	}
	wi.updateDistribution()
	wi.fightAliens()

	return wi
}

// MakeIteration is simple 1 iteration of process when alive aliens move randomly and after that fight
func (w *worldImp) MakeIteration() {
	w.moveAliens()
	w.updateDistribution()
	w.fightAliens()
}

func (w *worldImp) moveAliens() {
	for _, alien := range w.Aliens {
		alien.MakeRandomMove(w.WorldMap)
	}
}

func (w *worldImp) updateDistribution() {
	distribution := make(map[string][]*Alien)
	for name, _ := range w.WorldMap.WorldMap {
		distribution[name] = make([]*Alien, 0)
	}
	for _, alien := range w.Aliens {
		distribution[alien.Location.Name] = append(distribution[alien.Location.Name], alien)
	}

	w.Distribution = distribution
}

func (w *worldImp) fightAliens() {
	for cityName, aliens := range w.Distribution {
		// Destroy city and aliens if there are more than 1 alien in the city.
		if len(aliens) > 1 {
			names := make([]string, len(aliens))
			for i, alien := range aliens {
				names[i] = strconv.Itoa(int(alien.Name))
				delete(w.Aliens, alien.Name)
			}

			fmt.Printf("%s has been destroyed by aliens: %s\n", cityName, strings.Join(names[:], ", "))
			w.WorldMap.RemoveCity(cityName)
		}
	}
}

func (w *worldImp) PrintCities() {
	fmt.Println("---- World map ----")
	w.WorldMap.Print()
}

func (w *worldImp) HasAliveAliens() bool {
	return true
}