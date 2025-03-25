package main


type Move struct {
	Name string
	URL string
}

type OwnedPokemon struct {
	Name string
	// stats[0]
	HP int
	// stats[1]
	Attack int
	// stats[2]
	Defense int
	// stats[3]
	SpecialAttack int
	// stats[4]
	SpecialDefense int
	// stats[5]
	Speed int
	Moves []Move
}