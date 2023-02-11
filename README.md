# golang-cmgt-coin-miner

## Project

This project is a introduction assignment about blockchain. The program needs to interact with a public blockchain
hosted on school servers. This blockchain is accessible via a web service.

## Functional progamming

Functioneel programmeren is een manier van progammeren net zoals OOP een manier van programmeren is.
In Functional programming zijn er een aantal basis principes die belangrijk zijn.

- Pure functions

### Pure functions

Zijn functies die altijd hetzelfde resultaat terug gegeven. Dit wordt wel eens idempotent genoemd. Net zoals bij SAFE
methods in een API.
Een voorbeeld:

```go
func PureFunction(input string) string{
return input + "123"
}
```

### Function composition

Is het creëren van nieuwe functionaliteit door het combineren van meerdere pure functions.

voorbeeld:

```go

func GenerateName()string {
// Generates random name
name := "Bob"
return name
}

func GenerateLastName()string {
// Generates random name
name := "Marley"
return name
}

type Person struct{
name string
lastName string
}

func CreatePerson() Person{
return Person{name: GenerateName(), lastName: GenerateLastName()}
}

```

### Recursion

Recursion is een principe waarbij een stuk logica zich herhaald met verschillende inputs.
Dit is handig om verschillende outputs te genereren tegelijkertijd.

## Todo:

Do research about the following topics and describe implementation about those design patterns/ techniques in context of
this project.

- [x] Choose unit testing framework
- [x] Functional programming in golang
- [x] Pure functions
- [x] Function composition
- [x] Recursion in golang

Sources:

- [GopherCon 2020: Dylan Meeus - Functional Programming with Go](https://www.youtube.com/watch?v=wqs8n5Uk5OM&ab_channel=GopherAcademy)
