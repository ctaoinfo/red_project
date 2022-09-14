package src

type Character struct {
	Name          string
	Class         string
	Level         int
	Health        int
	CurrentHealth int
	Inventory     []string
}

func (c *Character) Init(name string, class string, level int, health int, currentHealth int, inventory []string) {
	c.Name = name
	c.Class = class
	c.Level = level
	c.Health = health
	c.CurrentHealth = currentHealth
	c.Inventory = inventory
}
