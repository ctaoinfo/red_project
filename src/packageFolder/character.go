//
// @property {string} Name - Le nom de l'objet.
// @property {string} Class - La classe du personnage.
// @property {int} Level - Le niveau du personnage.
// @property {int} Health - La santé maximale du personnage.
// @property {int} CurrentHealth - La santé actuelle du personnage.
// @property {int} Money - La somme d'argent dont dispose le joueur.
// @property Inventory - Une carte des articles et leurs quantités.
// @property Skill - Une carte des compétences, où la clé est le nom de la compétence et la valeur est
// le niveau de compétence.
// @property Equipement - C'est une structure qui contient l'équipement du joueur.
// CHUPIN Tao -- TARDELLI Andrea
// date
// YTrack Ynov Campus
// Red Project

package packageFolder // Package packageFolder

import ( // Importation des packages

	"strconv" // Package pour la conversion de type
)

type Character struct { // Strcture du personnage
	Name          string         // Nom
	Class         string         // Classe
	Level         int            // Niveau
	Health        int            // Vie max
	CurrentHealth int            // Vie actuelle
	Money         int            // Argent
	Inventory     map[string]int // Inventaire
	Skills        []string       // Compétences
	Equipement    *Equipement    // Équipement
}

func (c *Character) InitPlayer(name string, class string, level int, health int, currentHealth int, money int, inventory map[string]int, skills []string, Equipement *Equipement) { // Fonction d'initialisation du personnage
	c.Name = name                   // Nom
	c.Class = class                 // Classe
	c.Level = level                 // Niveau
	c.Health = health               // Vie max
	c.CurrentHealth = currentHealth // Vie actuelle
	c.Money = money                 // Argent
	c.Inventory = inventory         // Inventaire
	c.Skills = skills               // Compétences
	c.Equipement = Equipement       // Equipement
}

type Equipement struct { // Structure de l'équipement
	equiptete  string // Casque
	equiptorse string // Torse
	equippieds string // Pieds
}

func (c *Character) CharCreation(name string, class int) { // Fonction de création du personnage
	if class == 1 { // Si la classe est 1
		c.InitPlayer(name, "Sorcier", 1, 120, 60, 150, make(map[string]int), make([]string, 0), &Equipement{}) // Initialisation du personnage Sorcier
		c.Inventory["Potion"] = 3                                                                              // Ajout de 3 potions dans l'inventaire
	} else if class == 2 { // Si la classe est 2
		c.InitPlayer(name, "Archer", 1, 100, 50, 100, make(map[string]int), make([]string, 0), &Equipement{}) // Initialisation du personnage Archer
		c.Inventory["Potion"] = 3                                                                             // Ajout de 3 potions dans l'inventaire
	} else if class == 3 { // Si la classe est 3
		c.InitPlayer(name, "Tank", 1, 200, 100, 100, make(map[string]int), make([]string, 0), &Equipement{}) // Initialisation du personnage Tank
		c.Inventory["Potion"] = 3                                                                            // Ajout de 3 potions dans l'inventaire
	} else if class == 999 { // Si la classe est 999
		c.InitPlayer(name, "ADMIN", 100, 1000, 1000, 9999, make(map[string]int), make([]string, 0), &Equipement{}) // Initialisation du personnage ADMIN
		c.Inventory["Potion"] = 10                                                                                 // Ajout de 10 potions dans l'inventaire
		c.Inventory["AdvancedPotion"] = 10                                                                         // Ajout de 10 potions avancées dans l'inventaire
		c.Inventory["ForcePotion"] = 10                                                                            // Ajout de 10 potions de force dans l'inventaire
	}
}

func (c *Character) InitCharCreation() { // Fonction d'initialisation du personnage
	c.CharCreation(c.DisplayUserName(), c.DisplayUserClass()) // Création du personnage
}
func (c *Character) PlayerDead() { // Si le joueur est mort
	var health = c.Health / 2 // On récupère la moitié de la vie du joueur
	if c.CurrentHealth <= 0 { // Si la vie du joueur est égale à 0 ou inférieur à 0
		DisplayFrame("Mort", []string{ // On affiche un message de mort
			"Vous avez péri.",
			"Vous avez réapparu avec " + strconv.Itoa(health) + " HP max."})
		c.CurrentHealth = c.Health / 2 // On réinitialise la vie du joueur à la moitié de sa vie max
	}
}

func (c *Character) EquipArmor(armor string, maxHPplus int) { // Fonction d'équipement d'armure
	if c.Inventory[armor] >= 1 { // Si l'armure est dans l'inventaire
		c.Health += maxHPplus                // Ajout de la vie max
		c.Inventory[armor]--                 // Suppression de l'armure dans l'inventaire
		c.Equipement.equippieds = armor      // Ajout de l'armure dans l'équipement
		c.Equipement.equiptete = armor       // Ajout de l'armure dans l'équipement
		c.Equipement.equiptorse = armor      // Ajout de l'armure dans l'équipement
		DisplayFrame("Equipement", []string{ // Affichage de l'équipement
			" Vous avez équipé " + armor,
			"Votre vie maximale est maintenant de " + strconv.Itoa(c.Health)})
	} else { // Si l'armure n'est pas dans l'inventaire
		DisplayFrame("Equipement", []string{ // Affichage de l'équipement
			" Vous n'avez pas d'" + armor + " dans votre inventaire"})
	}
}

func (c *Character) EquipPied() { // Fonction d'équipement de chaussure
	c.EquipArmor("Bottes de cuir", 10) // Equipement de chaussure
}

func (c *Character) EquipTete() { // Fonction d'équipement de casque
	c.EquipArmor("Casque de cuir", 30) // Equipement de casque
}

func (c *Character) EquipTorse() { // Fonction d'équipement de plastron
	c.EquipArmor("Tunique de cuir", 50) // Equipement de plastron
}

func (m *Mob) InitMobCreation(Round int) { // Fonction d'initialisation du monstre
	m.InitMob("Goblin", "Goblin", Round, 100, 100) // Initialisation du monstre
}

func (c *Character) UpdateMoney(value int) {
	c.Money += value
	if c.Money <= 0 {
		c.Money = 0
	}
}

func (c *Character) AddItem(AddItem string) { // fonction ajouter un item
	verif := false                  // variable verif
	for item := range c.Inventory { // boucle pour vérifier si l'item est dans l'inventaire
		if AddItem == item { // si l'item est dans l'inventaire
			c.Inventory[AddItem]++ // on ajoute 1 à l'item
			verif = true           // verif = true
			break                  // on sort de la boucle
		} else { // si l'item n'est pas dans l'inventaire
			verif = false // verif = false
		}
	}
	if !verif { // si verif = false
		c.Inventory[AddItem] = 1 // on ajoute l'item
	}

}

func (c *Character) RemoveItem(item string, count int) {
	if c.Inventory[item]-count >= 0 {
		c.Inventory[item] -= count
	}
	if c.Inventory[item] == 0 {
		delete(c.Inventory, item)
	}
}

func (c *Character) VerifFullInventoryItems() bool { // Vérifie si l'inventaire est plein
	return len(c.Inventory) < 10 // Retourne vrai si l'inventaire est plein
}

func (c *Character) VerifFullInventoryItemsAmount(item string) bool { // Vérifie si l'inventaire est plein

	return c.Inventory[item] < 10 // Retourne vrai si l'inventaire est plein

}

func (c *Character) SpellBook(SkillAdd string) {
	if !c.HasSkill(SkillAdd) {
		c.Skills = append(c.Skills, SkillAdd)
	}
}

func (c *Character) HasSkill(skill string) bool {
	for _, c_skill := range c.Skills {
		if skill == c_skill {
			return false
		}
	}
	return true
}
