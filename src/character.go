package src

type Character struct {
	name          string
	class         string
	level         int
	health        int
	currentHealth int
	inventory     [10]string
}

func Init(name string, class string, level int, health int, currentHealth int) Character {

	p1 := Character{
		name:          name,
		class:         class,
		level:         level,
		health:        health,
		currentHealth: currentHealth,
		inventory:     [10]string{"potion", "potion", "potion"},
	}
	return p1

	// Possibilité d'un deuxième personnage

	// p2 := Character{
	// 	name: ""
	// 	class: "Berserker"
	// 	level: 1
	// 	health: 100
	// 	currentHealth: 40
	// 	inventory: [10]string{"potion", "potion", "potion"}
	// }
}
