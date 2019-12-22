# Assumptions

- Input file of cities is valid and we can read it without additional validation. 
It means that not only semantic correctness, 
but also that such map is not contradictory, for example:
```
Foo north=Bar
Bar north=Foo
``` 

- Every cities has only one particular city on each direction or has anything 
and it works for both cities, like if Foo has Bar on north 
that means Bar has Foo and only Foo on South.

- Aliens make their move simultaneously during the iteration.

- We need to output not only roads but also cities without roads, because in example 
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```
after destroying Bar city Bee still exists and we need to print it.

- During random creation of Aliens they can be placed random in the same place. 
On this 0 initial iteration they also will be fight.

# Usage


To build run from the folder project:
```bash
go mod vendor
go build .
```

To run after building run:
```bash
./tendermint -n=1 -input_path=tests/testdata/twocitiesmap.txt
```

To run tests:
```bash
cd world
go test
```

## Comand line paramenters

| Parameter        | Description                        | Default   |
|------------------|------------------------------------|-----------|
| _**n**_          | Number of randomly created aliens. | 0         |
| _**input_path**_ | Path to input file.                | input.txt |
