package world

type AlienName int

type Alien struct {
	Name     AlienName
	Location *City
}

func createRandomAliens(worldMap *Map, n int) map[AlienName]*Alien {
	aliens := make(map[AlienName]*Alien)

	for i := 0; i < n; i++ {
		randomCity := worldMap.GetRandomCity()
		alien := &Alien{
			Location: randomCity,
			Name:     AlienName(i),
		}
		aliens[alien.Name] = alien
	}

	return aliens
}

func (a *Alien) MakeRandomMove(worldMap *Map) {
	nextCityName := a.Location.GetRandomDirection()
	// If there is no possible moves for alien it is trapped and stays in the same city.
	if nextCityName == "" {
		return
	}
	a.Location = worldMap.GetCity(nextCityName)
}