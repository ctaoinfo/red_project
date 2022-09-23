// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

type Character struct { // Strcture du personnage
	Name          string         // Nom
	Class         string         // Classe
	Level         int            // Niveau
	Health        int            // Vie max
	CurrentHealth int            // Vie actuelle
	Money         int            // Argent
	Inventory     map[string]int // Inventaire Max 4 objet différent + une arme, nombre définit selon l'objet
	Skill         map[string]int // Compétence
}

func (c *Character) Init(name string, class string, level int, health int, currentHealth int, money int, inventory map[string]int, skill map[string]int) { // Fonction d'initialisation du personnage
	c.Name = name
	c.Class = class
	c.Level = level
	c.Health = health
	c.CurrentHealth = currentHealth
	c.Money = money
	c.Inventory = inventory
	c.Skill = skill
}

func (c *Character) CharCreation(name string, class int) {
	if class == 1 {
		c.Init(name, "Sorcier", 1, 120, 50, 150, make(map[string]int), make(map[string]int))
		c.Inventory["Potion"] = 3
		c.Skill["FireBall"] = 60
	} else if class == 2 {
		c.Init(name, "Archer", 1, 100, 40, 100, make(map[string]int), make(map[string]int))
		c.Inventory["Potion"] = 3
		c.Skill["FireArrow"] = 30
	} else if class == 3 {
		c.Init(name, "Tank", 1, 200, 70, 100, make(map[string]int), make(map[string]int))
		c.Inventory["Potion"] = 3
		c.Skill["Punch"] = 40
		c.Skill["Shield"] = 50
	} else if class == 999 {
		c.Init(name, "ADMIN", 100, 1000, 1000, 9999, make(map[string]int), make(map[string]int))
		c.Inventory["Potion"] = 99
		c.Inventory["AdvancedPotion"] = 99
		c.Inventory["ForcePotion"] = 99

	}
}

func (c *Character) InitCharCreation() {
	c.CharCreation(c.DisplayUserName(), c.DisplayUserClass())
}
