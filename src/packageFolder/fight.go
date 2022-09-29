// CHUPIN Tao -- TARDELLI Andrea
// date
// YTrack Ynov Campus
// Red Project

package packageFolder // packageFolder
import "fmt"

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
