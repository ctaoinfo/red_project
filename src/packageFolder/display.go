// CHUPIN Tao -- TARDELLI Andrea
// date
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
	"strconv"

	"github.com/rivo/uniseg"
)

func DisplayFrame(title string, text []string) {
	windowSizeLenght := 0
	for _, word := range text {
		if len(word) > windowSizeLenght {
			windowSizeLenght = 0
			for i := 0; i <= uniseg.GraphemeClusterCount(word); i++ {
				windowSizeLenght++
			}
		}
	}
	windowSizeLenght += 5
	fmt.Printf("╒")
	for i := 0; i < windowSizeLenght; i++ {
		fmt.Printf("═")
	}

	fmt.Println("╕") // Angle haut droit

	fmt.Printf("│")
	fmt.Print(title)
	for i := 0; i < windowSizeLenght-uniseg.GraphemeClusterCount(title); i++ {
		fmt.Printf(" ")
	}
	fmt.Println("│")

	fmt.Printf("╞")
	for i := 0; i < windowSizeLenght; i++ {
		fmt.Printf("═")
	}
	fmt.Println("╡")
	for _, word := range text {
		fmt.Printf("│")
		fmt.Print(word)
		for i := 0; i < windowSizeLenght-uniseg.GraphemeClusterCount(word); i++ {
			fmt.Printf(" ")
		}
		fmt.Println("│")
	}
	fmt.Printf("└")
	for i := 0; i < windowSizeLenght; i++ {
		fmt.Printf("─")
	}
	fmt.Println("┘")
}

func DisplayDescName() {
	DisplayFrame(" Personnalisation du pseudo", []string{
		"Votre pseudo doit comporter une majuscule en première lettre",
		"Suivi par des minuscules",
		"Tous les numéros ou caractères spéciaux sont interdit",
		"Le pseudo ne peut dépasser 15 caractères",
		"",
		"Veuillez entrer votre pseudo : "})
}

func DisplayDescClass() {
	DisplayFrame(" Personnalisation de la classe", []string{
		"Veuillez choisir votre classe :",
		"- 1. Sorcier",
		"- 2. Archer",
		"- 3. Tank"})
}

func DisplaySortAchete() {
	DisplayFrame("Nouveau Sort", []string{
		"Vous avez obtenu un nouveau sort !"})
}

// J'adore la bite mon gars
func DisplaySortDetenu() {
	DisplayFrame("Sort détenu", []string{"Vous avez déjà le sort."})
}

func (c *Character) DisplayUserName() string {
	var name string

	DisplayDescName()

	fmt.Scanln(&name)

	var nameMajAccepted bool
	var nameMinAccepted bool
	// var newName string
	// if len(name) >= 15 {
	// 	DisplayErrorNameLen()
	// }
	if len(name) > 0 {
		if rune(name[0]) >= 65 && rune(name[0]) <= 90 {
			nameMajAccepted = true
			for i := 1; i < len(name); i++ {
				if rune(name[i]) >= 97 && rune(name[i]) <= 122 {
					nameMinAccepted = true
					continue
				} else {
					nameMinAccepted = false
				}
			}
		} else {
			nameMajAccepted = false
		}
	}

	if !nameMajAccepted {
		DisplayErrorNameMaj()
		c.DisplayUserName()
	} else if !nameMinAccepted {
		DisplayErrorNameMin()
		c.DisplayUserName()
	} else {
		return name
	}
	return ""
}

func (c *Character) DisplayUserClass() int {
	var class int

	DisplayDescClass()

	fmt.Scanln(&class)

	if class == 0 {
		DisplayErrorClass()
		c.DisplayUserClass()
	} else if class == 999 {
		return class
	} else if class < 1 || class > 3 {
		DisplayErrorClass()
		c.DisplayUserClass()
	}
	return class
}

func DisplayError() {
	DisplayFrame("Erreur", []string{
		"Erreur Système,",
		"Recommencez"})
}

func DisplayErrorClass() {
	DisplayFrame("Erreur dans la classe", []string{
		"Votre classe doit être valide"})
}

func DisplayErrorNameMaj() {
	DisplayFrame("Nom pas accepté", []string{
		"Votre pseudo doit commencer par une majuscule"})
}

func DisplayErrorNameMin() {
	DisplayFrame("Erreur pseudo", []string{
		"Mise à pars la première lettre le reste doit être en minuscules",
		"Et ne doit pas comporter de caractères spéciaux"})
}

func DisplayErrorNameLen() {
	DisplayFrame("Erreur pseudo", []string{
		"Le pseudo ne peut comporter plus de 15 caractères"})
}

func (c *Character) DisplayInfoPlayer() { // Affichage information joueur
	DisplayFrame(" Info Joueur", []string{
		" Nom : " + c.Name,
		" Niveau : " + strconv.Itoa(c.Level),
		" Classe : " + c.Class,
		" " + strconv.Itoa(c.CurrentHealth) + " HP sur " + strconv.Itoa(c.Health) + " HP."})
}

func DisplayPrincipalMenu() { // Menu Principal
	DisplayFrame(" Menu Principal", []string{
		"- 1. Afficher information personnage",
		"- 2. Accés à l'inventaire",
		"- 3. Accoster le marchand",
		"- 4. Aller au combat",
		"- 5. Qui sont-ils ?",
		"- 6. Quitter le jeu"})
}

func DisplayInventoryMenu() { // Affichage menu de l'inventaire
	DisplayFrame(" Menu Inventaire", []string{
		"- 1. Afficher inventaire",
		"- 2. Afficher équipement",
		"- 3. Afficher mon argent",
		"- 4. Utiliser un item",
		"- 5. Jeter un item",
		"- 6. Retour au menu"})
}

func DisplayStuffInventory() {
	DisplayFrame(" Menu Equipement", []string{
		"- 1. Afficher l'équipement détenu",
		"- 2. Equiper un équipement",
		"- 3. Déséquiper un équipement",
		"- 4. Retour au menu"})
}

func DisplayMerchentFirstMeet() { // Affichage Première visite du marchand
	DisplayFrame(" Discour Marchand", []string{
		"Bonjour, je suis le nouveau marchand du coin",
		"Je vous ai donné une potion gratuite aller voir votre inventaire",
		"Vous aurez la possibilité par la suite d'avoir",
		"accés à une boutique plus complète au fur et à ",
		"mesure de votre aventure."})
}

func DisplayMerchentMenu() { // Affichage menu du Marchand
	DisplayFrame(" Menu Marchand", []string{
		"- 1. Acheter un objet",
		"- 2. Acheter un sort",
		"- 3. Achat équipement",
		"- 4. Vendre un objet",
		"- 5. Afficher inventaire",
		"- 6. Retour au menu"})
}

func DisplayListItemUtil() { // Affichage menu liste objet (objet utilitaire) à vendre par le marchand
	DisplayFrame(" Boutique Item", []string{
		"- 1. Potion Soin Basic",
		"- 2. Potion Soin Avancé",
		"- 3. Potion Force"})
}

func DisplayListStuff() {
	DisplayFrame(" Boutique Equipement", []string{
		"- 1. Plume de Corbeau - 3€",
		"- 2. Fourrure de Loup - 5€",
		"- 3. Peau de Troll - 8€",
		"- 4. Cuir de Sanglier - 11€",
		"- 5. Fil d'araigné - 15€",
		"- 6. Retour au menu"})
}

func DisplayMenuFight() {
	DisplayFrame(" Menu Combat", []string{
		"- 1. Info Ennemi",
		"- 2. Attaque",
		"- 3. Item",
		"- 4. Retour Menu"})
}

func (mob *Character) DisplayInfoEnemy() {
	DisplayFrame(" Info Ennemi", []string{
		" Nom : " + mob.Name,
		" Niveau : " + strconv.Itoa(mob.Level),
		" " + strconv.Itoa(mob.CurrentHealth) + " HP sur 100",
		"Puissance caché"})
}

func (c *Character) DisplayListSpell() {
	if c.Class == "Sorcier" {
		DisplayFrame(
			" Boutique Skill", []string{
				"- 1. Boule de feu",
				"- 2. Eclair",
				"- 3. Anneau soin",
				"- 4. Retour menu"})
	} else if c.Class == "Archer" {
		DisplayFrame(" Boutique Skill", []string{
			"- 1. Flèche",
			"- 2. Rafale flèche",
			"- 3. Escive",
			"- 4. Retour menu"})
	} else if c.Class == "Tank" {
		DisplayFrame(" Boutique Skill", []string{
			"- 1. Coup de poing",
			"- 2. Coup de tête",
			"- 3. Bouclier",
			"- 4. Retour menu"})
	} else if c.Class == "ADMIN" {
		DisplayFrame(" Boutique Skill", []string{
			"- 1. Destruction",
			"- 4. Retour menu"})
	}
}

func DisplayQuiSontIls() { // Quête Bonus Noms 2 artistes cachés
	DisplayFrame(" Qui sont-ils ?", []string{
		"- ABBA",
		"- Steven Spielberg"})
}

func (c *Character) DisplayPlayerDead() {
	DisplayFrame(" La Mort", []string{
		"%s à péri", c.Name})
}

func DisplayErrorSwitch() {
	DisplayFrame(" Erreur", []string{
		"Vous avez pas entrer une valeur invalide"})
}

func (c *Character) DisplayMenuSkillBook() {
	if c.Class == "Sorcier" {
		DisplayFrame(" Quel sort voulez-vous utiliser ?", []string{
			"- 1. Boule de feu",
			"- 2. Eclair",
			"- 3. Anneau de soin",
			"- 4. Retour au menu"})
	} else if c.Class == "Archer" {
		DisplayFrame(" Quel sort voulez-vous utiliser ?", []string{
			"- 1. Tire flèche",
			"- 2. Rafale",
			"- 3. Esquive",
			"- 4. Retour au menu"})
	} else if c.Class == "Tank" {
		DisplayFrame(" Quel sort voulez-vous utiliser ?", []string{
			"- 1. Coup de poing",
			"- 2. Coup de tête",
			"- 3. Bouclier",
			"- 4. Retour au menu"})
	} else if c.Class == "ADMIN" {
		DisplayFrame(" Quel sort voulez-vous utiliser ?", []string{
			"- 1. Destruction",
			"- 4. Retour au menu"})
	}
}

func (c *Character) DisplayInventory() {

	var txt string
	var listTxt []string
	for i, it := range c.Inventory {
		if i == "Potion" {
			txt = "Potion Basic"
		} else if i == "AdvancedPotion" {
			txt = "Potion Avancé"
		} else if i == "PoisonPotion" {
			txt = "Potion Poison"
		} else if i == "ForcePotion" {
			txt = "Potion Force"
		} else if i == "PlumeCorbeau" {
			txt = "Plume de Corbeau"
		} else if i == "FourrureLoup" {
			txt = "Fourrure de Loup"
		} else if i == "PeauTroll" {
			txt = "Peau de Troll"
		} else if i == "CuirSanglier" {
			txt = "Cuir de Sanglier"
		} else if i == "FildAraignee" {
			txt = "Fil d'Araignée"
		}
		listTxt = append(listTxt, "- "+txt+" - "+strconv.Itoa(it)+" disponible")

	}
	DisplayFrame(" Inventaire", listTxt)
}

func (c *Character) DisplayMoney() {
	DisplayFrame("Compte en banque", []string{
		"Vous avez " + strconv.Itoa(c.Money) + "€"})
}

func (c *Character) DisplayFullInventory() {
	DisplayFrame("Inventaire complet", []string{
		"Vous avez trop d'item dans votre inventaire.",
		"Vous pouvez en jetez dans le menu inventaire."})
}

func DisplayFireBall() {

	DisplayFrame("Lance Boulle de feu", []string{
		"Vous avez infligé x dégats"})
}

func DisplayLightning() {
	DisplayFrame("Lance Eclair", []string{
		"Vous avez infligé x dégats"})
}

func DisplayHealing() {
	DisplayFrame("Lance Soin", []string{
		"Vous avez infligé x dégats"})
}

func DisplayArrow() {
	DisplayFrame("Lance Flêche", []string{
		"Vous avez infligé x dégats"})
}

func DisplayGustArrow() {
	DisplayFrame("Lance Rafale", []string{
		"Vous avez infligé x dégats"})
}

func DisplayEscape() {
	DisplayFrame("Lance Esquive", []string{
		"Vous avez esquiver la prochaine attaque"})
}

func DisplayPunch() {
	DisplayFrame("Lance Coup de Poing", []string{
		"Vous avez infligé x dégats"})
}

func DisplayHeadButt() {
	DisplayFrame("Lance Coup de Tête", []string{
		"Vous avez infligé x dégats"})
}

func DisplayShield() {
	DisplayFrame("Lance Bouclier", []string{
		"Vous avez absorber x dégats"})
}

func DisplayDestruction() {
	DisplayFrame("Lance Destruction", []string{
		"Vous avez détruit votre adversaire"})
}

func DisplaySellItem() {
	DisplayFrame("Quelle item jeter ?", []string{
		"- 1. Potion Basic",
		"- 2. Potion avancée",
		"- 3. Potion de force",
		"- 4. Potion de poison",
		"- 5. Retour"})
}

func DisplaySellMenu() {
	DisplayFrame("Que voulez vous jeter ?", []string{
		"- 1. Un item",
		// "- 2. Un équipement",
		"- 5. Afficher l'inventaire",
		"- 6. Menu Principal"})
}
