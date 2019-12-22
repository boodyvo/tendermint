package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/boodyvo/tendermint/world"
)

func TestReading(t *testing.T) {
	r := require.New(t)
	filename := "testdata/onecitymap.txt"
	worldMap, err := world.ReadMapFromFile(filename)
	r.Nil(err)

	r.Equal(
		1,
		len(worldMap.WorldMap),
		"Should be only one city",
	)
	r.NotNil(
		worldMap.GetCity("Foo"),
		"City Foo should exist",
	)
}