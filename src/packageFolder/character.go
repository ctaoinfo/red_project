// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

type Character struct { // Strcture du personnage
	Name               string         // Nom
	Class              string         // Classe
	Level              int            // Niveau
	Health             int            // Vie max
	CurrentHealth      int            // Vie actuelle
	Money              int            // Argent
	Inventory          map[string]int // Inventaire Max 4 objet différent + une arme, nombre définit selon l'objet
	EquipmentInventory []string       // Inventaire d'équipement
	Skill              map[string]int // Compétence
	Equipment          Equipment      // Equipements
}

type Equipment struct {
	Helmet     map[string]int // Casque
	Chestplate map[string]int // Plastron
	Boots      map[string]int // Chaussure
}

func (c *Character) InitPlayer(name string, class string, level int, health int, currentHealth int, money int, inventory map[string]int, skill map[string]int) { // Fonction d'initialisation du personnage
	c.Name = name
	c.Class = class
	c.Level = level
	c.Health = health
	c.CurrentHealth = currentHealth
	c.Money = money
	c.Inventory = inventory
	c.Skill = skill
}

func (mob *Character) InitMob(name string, class string, level int, health int, currentHealth int, money int, inventory map[string]int, skill map[string]int) { // Fonction d'initialisation du personnage
	mob.Name = name
	mob.Class = class
	mob.Level = level
	mob.Health = health
	mob.CurrentHealth = currentHealth
	mob.Money = money
	mob.Inventory = inventory
	mob.Skill = skill
}

func (c *Equipment) InitEquipment(helmet map[string]int, chestplate map[string]int, boots map[string]int) {
	c.Helmet = helmet
	c.Chestplate = chestplate
	c.Boots = boots
}

func (c *Character) CharCreation(name string, class int) {
	if class == 1 {
		c.InitPlayer(name, "Sorcier", 1, 120, 60, 150, make(map[string]int), make(map[string]int))
		c.Inventory["Potion"] = 3
		c.Skill["FireBall"] = 60
	} else if class == 2 {
		c.InitPlayer(name, "Archer", 1, 100, 50, 100, make(map[string]int), make(map[string]int))
		c.Inventory["Potion"] = 3
		c.Skill["FireArrow"] = 30
	} else if class == 3 {
		c.InitPlayer(name, "Tank", 1, 200, 100, 100, make(map[string]int), make(map[string]int))
		c.Inventory["Potion"] = 3
		c.Skill["Punch"] = 40
		c.Skill["Shield"] = 50
	} else if class == 999 {
		c.InitPlayer(name, "ADMIN", 100, 1000, 1000, 9999, make(map[string]int), make(map[string]int))
		c.Inventory["Potion"] = 99
		c.Inventory["AdvancedPotion"] = 99
		c.Inventory["ForcePotion"] = 99

	}
}

func (mob *Character) MobCreation(Round int) {
	mob.InitMob("name", "ADMIN", 100, 1000, 1000, 9999, make(map[string]int), make(map[string]int))
}

func (c *Equipment) EquipementCreation() {
	c.InitEquipment(make(map[string]int), make(map[string]int), make(map[string]int))
}

func (c *Character) InitCharCreation() {
	c.CharCreation(c.DisplayUserName(), c.DisplayUserClass())
}
