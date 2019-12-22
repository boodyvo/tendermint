package world

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// NewWorld creates new world with provided city map and number of aliens.
// It generates n random aliens and make fight step for aliens that were created in the same city.
func testWorld(worldMap *Map, aliens map[AlienName]*Alien) *worldImp {
	// Initialize distribution for aliens.
	distribution := make(map[string][]*Alien)
	for name, _ := range worldMap.WorldMap {
		distribution[name] = make([]*Alien, 0)
	}
	wi := &worldImp{
		Aliens:       aliens,
		WorldMap:     worldMap,
		Distribution: distribution,
	}
	wi.updateDistribution()
	wi.fightAliens()

	return wi
}

// Create 3 cities and 2 aliens in that way so the will fight after first iteration.
func TestAlienDestroyCity(t *testing.T) {
	r := require.New(t)
	cityFoo := &City{
		Name:  "Foo",
		North: "Bar",
	}
	cityBar := &City{
		Name:  "Bar",
		North: "Baz",
		South: "Foo",
	}
	cityBaz := &City{
		Name:  "Baz",
		South: "Bar",
	}
	testWorldMap3Cities := &Map{
		WorldMap: map[string]*City{
			cityFoo.Name: cityFoo,
			cityBar.Name: cityBar,
			cityBaz.Name: cityBaz,
		},
	}
	testAliens := map[AlienName]*Alien{
		1: {
			Name: 1,
			Location: cityFoo,
		},
		2: {
			Name: 2,
			Location: cityBaz,
		},
	}

	world := testWorld(testWorldMap3Cities, testAliens)
	world.MakeIteration()

	r.Zero(
		len(world.Aliens),
		"Aliens should be destroyed after fight",
	)
	r.Equal(
		2,
		len(world.WorldMap.WorldMap),
		"World should have only 2 cities after fight",
	)
	r.Nil(
		world.WorldMap.GetCity(cityBar.Name),
		"City Bar should be destroyed after fight",
	)

	for i := 1; i < 5; i++ {
		world.MakeIteration()

		r.Zero(
			len(world.Aliens),
			fmt.Sprintf("Aliens should be destroyed after iteration %d", i),
		)
		r.Equal(
			2,
			len(world.WorldMap.WorldMap),
			fmt.Sprintf("World should have only 2 cities after iteration %d", i),
		)
		r.Nil(
			world.WorldMap.GetCity(cityBar.Name),
			fmt.Sprintf("City Bar should be destroyd after second iteration %d", i),
		)
	}
}

// Create 2 cities and 2 aliens so aliens will never meet
func TestInfiniteLoop(t *testing.T) {
	r := require.New(t)
	cityFoo := &City{
		Name:  "Foo",
		North: "Bar",
	}
	cityBar := &City{
		Name:  "Bar",
		South: "Foo",
	}
	testWorldMap3Cities := &Map{
		WorldMap: map[string]*City{
			cityFoo.Name: cityFoo,
			cityBar.Name: cityBar,
		},
	}
	testAliens := map[AlienName]*Alien{
		1: {
			Name: 1,
			Location: cityFoo,
		},
		2: {
			Name: 2,
			Location: cityBar,
		},
	}

	world := testWorld(testWorldMap3Cities, testAliens)

	for i := 0; i < 100; i++ {
		world.MakeIteration()

		r.Equal(
			2,
			len(world.Aliens),
			fmt.Sprintf("Aliens should be alive after iteration %d", i),
		)
		r.Equal(
			2,
			len(world.WorldMap.WorldMap),
			fmt.Sprintf("All cities should exist after iteration %d", i),
		)
	}
}

// Isolated world
func TestIsolatedWorld(t *testing.T) {
	r := require.New(t)
	cityFoo := &City{
		Name:  "Foo",
	}
	cityBar := &City{
		Name:  "Bar",
	}
	cityBaz := &City{
		Name:  "Baz",
	}
	cityPon := &City{
		Name:  "Pon",
	}
	cityPin := &City{
		Name:  "Pin",
	}
	testWorldMap4Cities := &Map{
		WorldMap: map[string]*City{
			cityFoo.Name: cityFoo,
			cityBar.Name: cityBar,
			cityBaz.Name: cityBaz,
			cityPon.Name: cityPon,
			cityPin.Name: cityPin,
		},
	}
	testAliens := map[AlienName]*Alien{
		0: {
			Name: 0,
			Location: cityPin,
		},
		1: {
			Name: 1,
			Location: cityFoo,
		},
		2: {
			Name: 2,
			Location: cityBar,
		},
		3: {
			Name: 3,
			Location: cityBar,
		},
		4: {
			Name: 4,
			Location: cityBaz,
		},
		5: {
			Name: 5,
			Location: cityBaz,
		},
		6: {
			Name: 6,
			Location: cityBaz,
		},
	}

	world := testWorld(testWorldMap4Cities, testAliens)
	r.Equal(
		2,
		len(world.Aliens),
		fmt.Sprintf("Only 2 aliens should be alive after initiation"),
	)
	r.Equal(
		3,
		len(world.WorldMap.WorldMap),
		fmt.Sprintf("3 cities should exist after initiation"),
	)
	r.NotNil(
		world.WorldMap.GetCity(cityFoo.Name),
		fmt.Sprintf("City Foo should exist after initiation"),
	)
	r.NotNil(
		world.WorldMap.GetCity(cityPon.Name),
		fmt.Sprintf("City Pon should exist after initiation"),
	)
	r.NotNil(
		world.WorldMap.GetCity(cityPin.Name),
		fmt.Sprintf("City Pon should exist after initiation"),
	)

	for i := 0; i < 20; i++ {
		world.MakeIteration()

		r.Equal(
			2,
			len(world.Aliens),
			fmt.Sprintf("2 Aliens should be alive after iteration %d", i),
		)
		r.Equal(
			3,
			len(world.WorldMap.WorldMap),
			fmt.Sprintf("3 cities should exist after iteration %d", i),
		)
	}
}