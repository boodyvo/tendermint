package world

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

type City struct {
	Name  string
	North string
	South string
	East  string
	West  string
}

// GetRandomDirection returns city name of random move or "" if it is isolated from other cities
func (c *City) GetRandomDirection() string {
	var possibleDirections []string
	if c.North != "" {
		possibleDirections = append(possibleDirections, c.North)
	}
	if c.East != "" {
		possibleDirections = append(possibleDirections, c.East)
	}
	if c.South != "" {
		possibleDirections = append(possibleDirections, c.South)
	}
	if c.West != "" {
		possibleDirections = append(possibleDirections, c.West)
	}
	if len(possibleDirections) == 0 {
		return ""
	}
	index := rand.Intn(len(possibleDirections))

	return possibleDirections[index]
}

type Map struct {
	WorldMap map[string]*City
}

func NewMap() *Map {
	rand.Seed(time.Now().Unix())
	return &Map{
		WorldMap: make(map[string]*City, 0),
	}
}

func (m *Map) GetCity(name string) *City {
	city, ok := m.WorldMap[name]
	if !ok {
		// City does not exist so we return nil
		return nil
	}

	return city
}

func (m *Map) GetCityWithCreation(name string) *City {
	city, ok := m.WorldMap[name]
	if !ok {
		// City does not exist so we need to create it
		city = &City{
			Name: name,
		}
		m.SetCity(city)
	}

	return city
}

// TODO(boodyvo): add error checking
func (m *Map) SetCity(city *City) {
	if city.Name != "" {
		m.WorldMap[city.Name] = city
	}
}

func (m *Map) RemoveCity(name string) {
	city, ok := m.WorldMap[name]
	if !ok {
		return
	}
	northCity := m.GetCity(city.North)
	if northCity != nil {
		northCity.South = ""
		m.SetCity(northCity)
	}

	westCity := m.GetCity(city.West)
	if westCity != nil {
		westCity.East = ""
		m.SetCity(westCity)
	}

	southCity := m.GetCity(city.South)
	if southCity != nil {
		southCity.North = ""
		m.SetCity(southCity)
	}

	eastCity := m.GetCity(city.East)
	if eastCity != nil {
		eastCity.West = ""
		m.SetCity(eastCity)
	}

	delete(m.WorldMap, name)
}

func (m *Map) GetRandomCity() *City {
	keys := reflect.ValueOf(m.WorldMap).MapKeys()
	if len(keys) == 0 {
		return nil
	}

	randKey := keys[rand.Intn(len(keys))].Interface().(string)

	return m.WorldMap[randKey]
}

func (m *Map) Print() {
	for _, city := range m.WorldMap {
		fmt.Print(city.Name)
		if city.North != "" {
			fmt.Printf(" north=%s", city.North)
		}
		if city.East != "" {
			fmt.Printf(" east=%s", city.East)
		}
		if city.South != "" {
			fmt.Printf(" south=%s", city.South)
		}
		if city.West != "" {
			fmt.Printf(" west=%s", city.West)
		}
		fmt.Print("\n")
	}
}

func ReadMapFromFile(filename string) (*Map, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	worldMap := NewMap()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		// we don't validate input file because of assumption that is is correct
		words := strings.Split(scanner.Text(), " ")
		name := words[0]
		city := worldMap.GetCityWithCreation(name)
		for i, directionSir := range words {
			// skip city name
			if i == 0 {
				continue
			}
			direction := strings.Split(directionSir, "=")
			edgeCity := worldMap.GetCityWithCreation(direction[1])
			switch direction[0] {
			case "north":
				city.North = edgeCity.Name
				edgeCity.South = city.Name
			case "west":
				city.West = edgeCity.Name
				edgeCity.East = city.Name
			case "south":
				city.South = edgeCity.Name
				edgeCity.North = city.Name
			case "east":
				city.East = edgeCity.Name
				edgeCity.West = city.Name
			default:
				// some unknown input
			}
			worldMap.SetCity(edgeCity)
		}
		worldMap.SetCity(city)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return worldMap, nil
}
