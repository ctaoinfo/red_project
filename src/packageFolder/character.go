// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

type Character struct { // Strcture du personnage
	Name          string   // Nom
	Class         string   // Classe
	Level         int      // Niveau
	Health        int      // Vie max
	CurrentHealth int      // Vie actuelle
	Inventory     []string // Inventaire Max 4 objet différent + une arme, nombre définit selon l'objet
}

func (c *Character) Init(name string, class string, level int, health int, currentHealth int, inventory []string) { // Fonction d'initialisation du personnage
	c.Name = name
	c.Class = class
	c.Level = level
	c.Health = health
	c.CurrentHealth = currentHealth
	c.Inventory = inventory
}
