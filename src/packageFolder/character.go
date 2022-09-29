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
	"fmt"     // Package pour l'affichage
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
	Skill         map[string]int // Compétences
	Equipement    *Equipement    // Compétences
}

func (c *Character) InitPlayer(name string, class string, level int, health int, currentHealth int, money int, inventory map[string]int, skill map[string]int, Equipement *Equipement) { // Fonction d'initialisation du personnage
	c.Name = name                   // Nom
	c.Class = class                 // Classe
	c.Level = level                 // Niveau
	c.Health = health               // Vie max
	c.CurrentHealth = currentHealth // Vie actuelle
	c.Money = money                 // Argent
	c.Inventory = inventory         // Inventaire
	c.Skill = skill                 // Compétences
	c.Equipement = Equipement       // Equipement
}

type Equipement struct { // Structure de l'équipement
	equiptete  string // Casque
	equiptorse string // Torse
	equippieds string // Pieds
}

type Mob struct { // Strcture du personnage
	Name          string // Nom
	Class         string // Classe
	Level         int    // Niveau
	Health        int    // Vie max
	CurrentHealth int    // Vie actuelle
}

func (m *Mob) InitMob(name string, class string, level int, health int, currentHealth int) { // Fonction d'initialisation du personnage
	m.Name = name                   // Nom
	m.Class = class                 // Classe
	m.Level = level                 // Niveau
	m.Health = health               // Vie max
	m.CurrentHealth = currentHealth // Vie actuelle
}

func (c *Character) CharCreation(name string, class int) { // Fonction de création du personnage
	if class == 1 { // Si la classe est 1
		c.InitPlayer(name, "Sorcier", 1, 120, 60, 150, make(map[string]int), make(map[string]int), &Equipement{}) // Initialisation du personnage Sorcier
		c.Inventory["Potion"] = 3                                                                                 // Ajout de 3 potions dans l'inventaire
	} else if class == 2 { // Si la classe est 2
		c.InitPlayer(name, "Archer", 1, 100, 50, 100, make(map[string]int), make(map[string]int), &Equipement{}) // Initialisation du personnage Archer
		c.Inventory["Potion"] = 3                                                                                // Ajout de 3 potions dans l'inventaire
	} else if class == 3 { // Si la classe est 3
		c.InitPlayer(name, "Tank", 1, 200, 100, 100, make(map[string]int), make(map[string]int), &Equipement{}) // Initialisation du personnage Tank
		c.Inventory["Potion"] = 3                                                                               // Ajout de 3 potions dans l'inventaire
	} else if class == 999 { // Si la classe est 999
		c.InitPlayer(name, "ADMIN", 100, 1000, 1000, 9999, make(map[string]int), make(map[string]int), &Equipement{}) // Initialisation du personnage ADMIN
		c.Inventory["Potion"] = 10                                                                                    // Ajout de 10 potions dans l'inventaire
		c.Inventory["AdvancedPotion"] = 10                                                                            // Ajout de 10 potions avancées dans l'inventaire
		c.Inventory["ForcePotion"] = 10                                                                               // Ajout de 10 potions de force dans l'inventaire
	}
}

func (c *Character) InitCharCreation() { // Fonction d'initialisation du personnage
	c.CharCreation(c.DisplayUserName(), c.DisplayUserClass()) // Création du personnage
}
func (c *Character) PlayerDead() { // Si le joueur est mort
	var health = c.Health / 2                        // On récupère la moitié de la vie du joueur
	if c.CurrentHealth == 0 || c.CurrentHealth < 0 { // Si la vie du joueur est égale à 0 ou inférieur à 0
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

// C'est une fonction qui crée un combat contre un gobelin.
func (m *Mob) FightMob(Round int, c *Character) { // Fonction de combat contre le monstre
	var choice int                                             // Choix du joueur
	m.InitMobCreation(Round)                                   // Initialisation du monstre
	DisplayFrame("Un "+m.Name+" sauvage apparait !", []string{ // Affichage de l'apparition du monstre
		" Que voulez-vous faire ? ",
		" 1. Attaquer ",
		" 2. Utiliser une potion ",
		" 3. Fuir ",
		" 4. Inventaire ",
		" 5. Statistiques "})
	for m.CurrentHealth > 0 && c.CurrentHealth > 0 { // Tant que le monstre et le joueur ont de la vie
		fmt.Scanln(&choice) // Choix du joueur
		if choice == 1 {    // Si le joueur choisi d'attaquer
			// c.AttackMob(m)
			DisplayFrame("Un "+m.Name+" sauvage apparait !", []string{ // Affichage de l'apparition du monstre
				" Que voulez-vous faire ? ",
				" 1. Attaquer ",
				" 2. Utiliser une potion ",
				" 3. Fuir ",
				" 4. Inventaire ",
				" 5. Statistiques "})
		} else if choice == 2 { // Si le joueur choisi d'utiliser une potion
			c.UseItemPerso(m)                                          // Utilisation d'une potion
			DisplayFrame("Un "+m.Name+" sauvage apparait !", []string{ // Affichage de l'apparition du monstre
				" Que voulez-vous faire ? ",
				" 1. Attaquer ",
				" 2. Utiliser une potion ",
				" 3. Fuir ",
				" 4. Inventaire ",
				" 5. Statistiques "})
		} else if choice == 3 { // Si le joueur choisi de fuir
			c.Fuir()                                                   // Fuite du joueur
			DisplayFrame("Un "+m.Name+" sauvage apparait !", []string{ // Affichage de l'apparition du monstre
				" Que voulez-vous faire ? ",
				" 1. Attaquer ",
				" 2. Utiliser une potion ",
				" 3. Fuir ",
				" 4. Inventaire ",
				" 5. Statistiques "})
		} else if choice == 4 { // Si le joueur choisi d'ouvrir l'inventaire
			c.DisplayInventory()                                       // Affichage de l'inventaire
			DisplayFrame("Un "+m.Name+" sauvage apparait !", []string{ // Affichage de l'apparition du monstre
				" Que voulez-vous faire ? ",
				" 1. Attaquer ",
				" 2. Utiliser une potion ",
				" 3. Fuir ",
				" 4. Inventaire ",
				" 5. Statistiques "})
		} else if choice == 5 { // Si le joueur choisi d'afficher ses statistiques
			c.StatsDisplay()                                           // Affichage des statistiques
			DisplayFrame("Un "+m.Name+" sauvage apparait !", []string{ // Affichage de l'apparition du monstre
				" Que voulez-vous faire ? ",
				" 1. Attaquer ",
				" 2. Utiliser une potion ",
				" 3. Fuir ",
				" 4. Inventaire ",
				" 5. Statistiques "})
		} else {
			DisplayErrorSwitch() // Affichage d'une erreur
		}
	}
}
