// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
)

var Item int

func (c *Character) DisplayInfo() {

	fmt.Println("-------------------------")
	fmt.Println("|	Nom : ", c.Name, "	|")
	fmt.Println("-------------------------")
	fmt.Println("|	Niveau : ", c.Level, "	|")
	fmt.Println("-------------------------")
	fmt.Println("|	Classe : ", c.Class, "	|")
	fmt.Println("-------------------------")
	fmt.Println("|	", c.CurrentHealth, "HP sur 100	|")
	fmt.Println("-------------------------")
	// fmt.Println("Appuyer enter pour sortir")

}

func (c *Character) AccessInventory() {

	var choice int
	fmt.Println("---------------------------------")
	fmt.Println("|	Acces à l'inventaire	|")
	fmt.Println("---------------------------------")
	fmt.Println("1. Afficher inventaire")
	fmt.Println("2. Utiliser un item")
	fmt.Println("3. Jeter un item")
	fmt.Println("4. Retour au menu")
	fmt.Println("---------------------------------")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		for _, itemPresent := range c.Inventory {
			fmt.Println("\t")
			fmt.Println(itemPresent)
		}
		// fmt.Println(Count(c.Inventory))
	case 2:
		c.UseItem()

	case 3:
		c.throwItem()

	default:
		return
	}
}

func (c *Character) Menu() {
	var choice int

	fmt.Println("---------------------------------")
	fmt.Println("|	Affichage du menu	|")
	fmt.Println("---------------------------------")
	fmt.Println("1. Afficher information personnage")
	fmt.Println("2. Accés à l'inventaire")
	fmt.Println("3. Accoster le marchand")
	fmt.Println("4. return")

	fmt.Scanln(&choice)
	fmt.Print("\n")

	switch choice {
	case 1:
		c.DisplayInfo()
	case 2:
		c.AccessInventory()
	case 3:
		c.Marchand()
	case 4:
		return
	}

}

func (c Character) Marchand() {
	var choice int
	var firstMeet bool = true

	if firstMeet {

		fmt.Println("Bonjour, je suis le nouveau marchand du coin")
		fmt.Println("Je peux vous vendre une potion gratuite")
		fmt.Println("----------------------------")
		fmt.Println("Vous aurez la possibilité par la suite d'avoir accés à une boutique plus complète au fur et à mesure de votre aventure")

		firstMeet = false
	}

	// Menu Marchand
	fmt.Println("---------------------------------")
	fmt.Println("|	Menu du marchand	|")
	fmt.Println("---------------------------------")
	fmt.Println("1. Acheter un objet")
	fmt.Println("2. Vendre un objet")
	fmt.Println("3. Retour au menu")

	fmt.Scanln(&choice)

	switch choice {
	case 1:

	case 2:

	case 3:

	}

}
