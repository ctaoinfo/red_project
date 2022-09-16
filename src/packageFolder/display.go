// CHUPIN Tao -- CORBEL Andrea
//
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
)

var Item int

func (c *Character) DisplayInfoPlayer() {

	fmt.Println("+-----------------------+")
	fmt.Println("|	Nom : ", c.Name, "	|")
	fmt.Println("+-----------------------+")
	fmt.Println("|	Niveau : ", c.Level, "	|")
	fmt.Println("+-----------------------+")
	fmt.Println("|	Classe : ", c.Class, "|")
	fmt.Println("+-----------------------+")
	fmt.Println("|	", c.CurrentHealth, "HP sur 100	|")
	fmt.Println("+-----------------------+")

}

func DisplayInventoryMenu() {
	fmt.Println("+-------------------------------+")
	fmt.Println("|	Acces à l'inventaire	|")
	fmt.Println("+-------------------------------+")
	fmt.Println("1. Afficher inventaire")
	fmt.Println("2. Utiliser un item")
	fmt.Println("3. Jeter un item")
	fmt.Println("4. Retour au menu")
	fmt.Println("---------------------------------")
}

func DisplayPrincipalMenu() { // Menu Principal
	fmt.Println("+-------------------------------+")
	fmt.Println("|	Affichage du menu	|")
	fmt.Println("+-------------------------------+")
	fmt.Println("1. Afficher information personnage")
	fmt.Println("2. Accés à l'inventaire")
	fmt.Println("3. Accoster le marchand")
	fmt.Println("4. return")
}

func DisplayMerchentFirstMeet() {
	fmt.Println("Bonjour, je suis le nouveau marchand du coin")
	fmt.Println("Je peux vous vendre une potion gratuite")
	fmt.Println("----------------------------")
	fmt.Println("Vous aurez la possibilité par la suite d'avoir accés à une boutique \n plus complète au fur et à mesure de votre aventure")
}

func DisplayMerchentMenu() { // Affichage menu du Marchand
	fmt.Println("+-------------------------------+")
	fmt.Println("|	Menu du marchand	|")
	fmt.Println("+-------------------------------+")
	fmt.Println("1. Acheter un objet")
	fmt.Println("2. Vendre un objet")
	fmt.Println("3. Retour au menu")
}

func DisplayListItemUtil() { // Affichage menu liste objet à vendre par le marchand
	fmt.Println("Liste Item Boutique")
	fmt.Println("1. Potion Soin Basic")

	fmt.Println("2. Potion Soin Avancé")
	fmt.Println("3. Potion Force")
	fmt.Println("4. Totem de Réanimation")
}

func DisplayListItemWeapon() {

}

func DisplayQuiSontIls() {
	fmt.Println("-")
	fmt.Println("-")
	fmt.Println("-")
}
