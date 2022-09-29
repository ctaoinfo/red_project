// CHUPIN Tao -- TARDELLI Andrea
// date
// YTrack Ynov Campus
// Red Project

package packageFolder

import (
	"fmt"
	"strconv"

	"github.com/rivo/uniseg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func DisplayFrame(title string, text []string) { // Affichage cadre
	windowSizeLenght := 0 // Taille de la fenêtre

	for _, word := range text { // Parcours du tableau de texte
		if len(word) > windowSizeLenght { // Si la taille du mot est plus grande que la taille de la fenêtre
			windowSizeLenght = 0                                      // On remet la taille de la fenêtre à 0
			for i := 0; i <= uniseg.GraphemeClusterCount(word); i++ { // On parcours le mot
				windowSizeLenght++ // On incrémente la taille de la fenêtre
			}
		}
		if len(title) > len(word) && len(title) > windowSizeLenght { // Si la taille du titre est plus grande que la taille du mot et que la taille du titre est plus grande que la taille de la fenêtre
			windowSizeLenght = 0                                       // On remet la taille de la fenêtre à 0
			for i := 0; i <= uniseg.GraphemeClusterCount(title); i++ { // On parcours le titre
				windowSizeLenght++ // On incrémente la taille de la fenêtre
			}
		}
	}
	windowSizeLenght += 5                   // On ajoute 5 à la taille de la fenêtre
	fmt.Printf("╒")                         // On affiche le coin supérieur gauche
	for i := 0; i < windowSizeLenght; i++ { // On parcours la taille de la fenêtre
		fmt.Printf("═") // On affiche un trait horizontal
	}

	fmt.Println("╕") // On affiche le coin supérieur droit

	fmt.Printf("│")                                                            // On affiche un trait vertical
	fmt.Print(title)                                                           // On affiche le titre
	for i := 0; i < windowSizeLenght-uniseg.GraphemeClusterCount(title); i++ { // On parcours la taille de la fenêtre moins la taille du titre
		fmt.Printf(" ") // On affiche un espace
	}
	fmt.Println("│") // On affiche un trait vertical

	fmt.Printf("╞")                         // On affiche le coin central gauche
	for i := 0; i < windowSizeLenght; i++ { // On parcours la taille de la fenêtre
		fmt.Printf("═") // On affiche un trait horizontal
	}
	fmt.Println("╡")            // On affiche le coin central droit
	for _, word := range text { // On parcours le tableau de texte
		fmt.Printf("│")                                                           // On affiche un trait vertical
		fmt.Print(word)                                                           // On affiche le mot
		for i := 0; i < windowSizeLenght-uniseg.GraphemeClusterCount(word); i++ { // On parcours la taille de la fenêtre moins la taille du mot
			fmt.Printf(" ") // On affiche un espace
		}
		fmt.Println("│") // On affiche un trait vertical
	}
	fmt.Printf("└")                         // On affiche le coin inférieur gauche
	for i := 0; i < windowSizeLenght; i++ { // On parcours la taille de la fenêtre
		fmt.Printf("─") // On affiche un trait horizontal
	}
	fmt.Println("┘") // On affiche le coin inférieur droit
}

func DisplayDescName() { // Affichage description nom
	DisplayFrame(" Personnalisation du pseudo", []string{ // Affichage cadre personnalisation nom
		"Votre pseudo doit comporter une majuscule en première lettre",
		"Suivi par des minuscules",
		"Tous les numéros ou caractères spéciaux sont interdit",
		"Le pseudo ne peut dépasser 15 caractères",
		"",
		"Veuillez entrer votre pseudo : "})
}

func DisplayDescClass() { // Affichage description classe
	DisplayFrame(" Personnalisation de la classe", []string{ // Affichage cadre personnalisation classe
		"Veuillez choisir votre classe :",
		"- 1. Sorcier",
		"- 2. Archer",
		"- 3. Tank"})
}

func DisplaySortAchete() { // Affichage sort acheté
	DisplayFrame("Nouveau Sort", []string{ // Affichage cadre nouveau sort
		"Vous avez obtenu un nouveau sort !"})
}

func DisplaySortDetenu() { // Affichage sort détenu
	DisplayFrame("Sort détenu", []string{"Vous avez déjà le sort."}) // Affichage cadre sort détenu
}

func CheckName(name *string) bool { // Vérification nom
	for _, letter := range *name { // Parcours du nom
		if !(letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z') { // Si la lettre n'est pas comprise entre a et z ou A et Z
			return false // On retourne faux
		}
	}
	caser := cases.Title(language.English)

	*name = caser.String(*name) // On met la première lettre en majuscule et le reste en minuscule
	return true                 // On retourne vrai
}

func (c *Character) DisplayUserName() string { // Affichage nom
	var name string          // Nom
	DisplayImageAsciiDebut() // Affichage image ascii début

	DisplayDescName() // Affichage description nom

	fmt.Scanf("%s\n", &name) // On récupère le nom

	if len(name) > 0 { // Si la taille du nom est supérieur à 0
		if len(name) >= 15 { // Si la taille du nom est supérieur ou égale à 15
			DisplayErrorNameLen() // On affiche l'erreur de taille du nom
			c.DisplayUserName()   // On relance la fonction
		} else if CheckName(&name) { // Sinon si le nom est correct
			return name // On retourne le nom
		} else {
			DisplayErrorName()  // On affiche l'erreur de nom
			c.DisplayUserName() // On relance la fonction
		}
	} else if len(name) == 0 { // Sinon si la taille du nom est égale à 0
		DisplayErrorNameLen() // On affiche l'erreur de taille du nom
		c.DisplayUserName()   // On relance la fonction
	}
	return "" // On retourne une chaine vide
}

func DisplayErrorName() { // Affichage erreur nom
	DisplayFrame("Erreur", []string{"Votre pseudo n'ai pas valide."}) // Affichage cadre erreur nom
}

func (c *Character) DisplayUserClass() int { // Affichage classe
	var class int // Classe

	DisplayDescClass() // Affichage description classe

	fmt.Scanln(&class) // On récupère la classe

	if class == 0 { // Si la classe est égale à 0
		DisplayErrorClass()  // On affiche l'erreur de classe
		c.DisplayUserClass() // On relance la fonction
	} else if class == 999 { // Sinon si la classe est égale à 999
		return class // On retourne la classe
	} else if class < 1 || class > 3 { // Sinon si la classe est inférieur à 1 ou supérieur à 3
		DisplayErrorClass()  // On affiche l'erreur de classe
		c.DisplayUserClass() // On relance la fonction
	}
	return class // On retourne la classe
}

func DisplayError() { // Affichage erreur
	DisplayFrame("Erreur", []string{ // Affichage cadre erreur
		"Erreur Système,",
		"Recommencez"})
}

func (c *Character) DisplayNoMoney() {
	DisplayFrame("Pas d'argent", []string{
		"Vous n'avais pas asser d'argent"})
}

func DisplayErrorClass() { // Affichage erreur classe
	DisplayFrame("Erreur dans la classe", []string{ // Affichage cadre erreur classe
		"Votre classe doit être valide"})
}

func DisplayErrorNameLen() { // Affichage erreur nom
	DisplayFrame("Erreur pseudo", []string{ // Affichage cadre erreur nom
		"Le pseudo ne peut comporter plus de 15 caractères"})
}

func (c *Character) DisplayInfoPlayer() { // Affichage information joueur
	DisplayFrame(" Info Joueur", []string{ // Affichage cadre information joueur
		" Nom : " + c.Name,                   // Nom
		" Niveau : " + strconv.Itoa(c.Level), // Niveau
		" Classe : " + c.Class,               // Classe
		" " + strconv.Itoa(c.CurrentHealth) + " HP sur " + strconv.Itoa(c.Health) + " HP."}) // Points de vie
}

func DisplayPrincipalMenu() { // Menu Principal
	DisplayFrame(" Menu Principal", []string{ // Affichage cadre menu principal
		"- 1. Afficher information personnage",
		"- 2. Accés à l'inventaire",
		"- 3. Accoster le marchand",
		"- 4. Accoster le forgeron",
		"- 5. Aller au combat",
		"- 6. Qui sont-ils ?",
		"- 7. Quitter le jeu"})
}

func DisplayInventoryMenu() { // Affichage menu de l'inventaire
	DisplayFrame(" Menu Inventaire", []string{ // Affichage cadre menu inventaire
		"- 1. Afficher inventaire",
		"- 2. Afficher équipement",
		"- 3. Afficher mon argent",
		"- 4. Utiliser un item",
		"- 5. Voir sort",
		"- 6. Retour au menu"})
}

func DisplayMenuStuffInventory() { // Affichage menu équipement inventaire
	DisplayFrame(" Menu Equipement", []string{ // Affichage cadre menu équipement
		"- 1. Afficher l'équipement détenu",
		"- 2. Equiper un équipement",
		"- 3. Déséquiper un équipement",
		"- 4. Retour au menu"})
}

func DisplayMerchentFirstMeet() { // Affichage Première visite du marchand
	DisplayFrame(" Discour Marchand", []string{ // Affichage cadre première visite du marchand
		"Bonjour, je suis le nouveau marchand du coin",
		"Je vous ai donné une potion gratuite aller voir votre inventaire",
		"Vous aurez la possibilité par la suite d'avoir",
		"accés à une boutique plus complète au fur et à ",
		"mesure de votre aventure."})
}

func DisplayMerchentMenu() { // Affichage menu du Marchand
	DisplayFrame(" Menu Marchand", []string{ // Affichage cadre menu marchand
		"- 1. Acheter un objet",
		"- 2. Acheter un sort",
		"- 3. Achat équipement",
		"- 4. Afficher inventaire",
		"- 5. Retour au menu"})
}

func DisplayListItemUtil() { // Affichage menu liste objet (objet utilitaire) à vendre par le marchand
	DisplayFrame(" Boutique Item", []string{ // Affichage cadre liste objet à vendre par le marchand
		"- 1. Potion Soin Basic",
		"- 2. Potion Soin Avancé",
		"- 3. Potion Force"})
}

func DisplayListStuff() { // Affichage menu liste équipement à vendre par le marchand
	DisplayFrame(" Boutique Equipement", []string{ // Affichage cadre liste équipement à vendre par le marchand
		"- 1. Plume de Corbeau - 36€",
		"- 2. Fourrure de Loup - 45€",
		"- 3. Peau de Troll - 32€",
		"- 4. Cuir de Sanglier -28€",
		"- 5. Fil d'araigné - 34€",
		"- 6. Retour au menu"})
}

func DisplayMenuFight() { // Affichage menu combat
	DisplayFrame(" Menu Combat", []string{ // Affichage cadre menu combat
		"- 1. Info Ennemi",
		"- 2. Attaque",
		"- 3. Item",
		"- 4. Retour Menu"})
}

func (mob *Character) DisplayInfoEnemy() { // Affichage information ennemi
	DisplayFrame(" Info Ennemi", []string{ // Affichage cadre information ennemi
		" Nom : " + mob.Name,
		" Niveau : " + strconv.Itoa(mob.Level),
		" " + strconv.Itoa(mob.CurrentHealth) + " HP sur 100",
		"Puissance caché"})
}

func (c *Character) DisplayListBoutiqueSpell() { // Affichage liste sort à vendre par le marchand
	if c.Class == "Sorcier" { // Si la classe est sorcier
		DisplayFrame(" Boutique Skill", []string{ // Affichage cadre liste sort à vendre par le marchand
			"- 1. Boule de feu",
			"- 2. Eclair",
			"- 3. Anneau soin",
			"- 4. Retour menu"})
	} else if c.Class == "Archer" { // Si la classe est archer
		DisplayFrame(" Boutique Skill", []string{ // Affichage cadre liste sort à vendre par le marchand
			"- 1. Flèche",
			"- 2. Rafale flèche",
			"- 3. Escive",
			"- 4. Retour menu"})
	} else if c.Class == "Tank" { // Si la classe est tank
		DisplayFrame(" Boutique Skill", []string{ // Affichage cadre liste sort à vendre par le marchand
			"- 1. Coup de poing",
			"- 2. Coup de tête",
			"- 3. Bouclier",
			"- 4. Retour menu"})
	} else if c.Class == "ADMIN" { // Si la classe est ADMIN
		DisplayFrame(" Boutique Skill", []string{ // Affichage cadre liste sort à vendre par le marchand
			"- 1. Destruction",
			"- 4. Retour menu"})
	}
}

func DisplayQuiSontIls() { // Quête Bonus Noms 2 artistes cachés
	DisplayFrame(" Qui sont-ils ?", []string{ // Affichage cadre quête bonus
		"- ABBA",
		"- Steven Spielberg"})
}

func (c *Character) DisplayPlayerDead() { // Affichage message mort du joueur
	DisplayFrame(" La Mort", []string{ // Affichage cadre message mort du joueur
		"%s à péri", c.Name})
}

func DisplayErrorSwitch() { // Affichage message erreur switch
	DisplayFrame(" Erreur", []string{ // Affichage cadre message erreur
		"Vous avez pas entrer une valeur invalide"})
}

func (c *Character) DisplayMenuSkillBook() { // Affichage menu livre de sort

	if c.Class == "Sorcier" { // Si la classe est sorcier
		var txt string               // Déclaration variable txt
		var listTxt []string         // Déclaration variable listTxt
		for _, i := range c.Skills { // Boucle pour afficher les sorts du joueur
			if i == "Boule de feu" { // Si le sort est boule de feu
				txt = "1 - Utiliser Boule de feu" // Affichage du sort
			} else if i == "Eclair" { // Si le sort est eclair
				txt = "2 - Utiliser Éclair" // Affichage du sort
			} else if i == "Anneau de soin" { // Si le sort est anneau de soin
				txt = "3 - Utiliser Anneau de soin" // Affichage du sort
			}
			listTxt = append(listTxt, txt) // Ajout du sort dans la liste
		}
		DisplayFrame(" Liste sort Sorcier", listTxt) // Affichage cadre liste sort
	} else if c.Class == "Archer" { // Si la classe est archer
		var txt string               // Déclaration variable txt
		var listTxt []string         // Déclaration variable listTxt
		for _, i := range c.Skills { // Boucle pour afficher les sorts du joueur
			if i == "Tire flèche" { // Si le sort est tire flèche
				txt = "1 - Utiliser Tire flèche" // Affichage du sort
			} else if i == "Rafale" { // Si le sort est rafale
				txt = "2 - Utiliser Rafale" // Affichage du sort
			} else if i == "Esquive" { // Si le sort est esquive
				txt = "3 - Utiliser Esquive" // Affichage du sort
			}
			listTxt = append(listTxt, txt) // Ajout du sort dans la liste
		}
		DisplayFrame(" Liste sort Archer", listTxt) // Affichage cadre liste sort
	} else if c.Class == "Tank" { // Si la classe est tank
		var txt string               // Déclaration variable txt
		var listTxt []string         // Déclaration variable listTxt
		for _, i := range c.Skills { // Boucle pour afficher les sorts du joueur
			if i == "Coup de poing" { // Si le sort est coup de poing
				txt = "1 - Coup de poing" // Affichage du sort
			} else if i == "Coup de tête" { // Si le sort est coup de tête
				txt = "2 - Utiliser Coup de tête" // Affichage du sort
			} else if i == "Bouclier" { // Si le sort est bouclier
				txt = "3 - Utiliser Bouclier" // Affichage du sort
			}
			listTxt = append(listTxt, txt) // Ajout du sort dans la liste
		}
		DisplayFrame(" Liste sort Tank", listTxt) // Affichage cadre liste sort
	} else if c.Class == "ADMIN" { // Si la classe est ADMIN
		DisplayFrame(" Liste sort ADMIN", []string{ // Affichage cadre liste sort
			"- 1. Destruction",
			"- 4. Retour au menu"})
	}
}

func (c *Character) DisplayInventory() { // Affichage inventaire

	var txt string                   // Déclaration variable txt
	var listTxt []string             // Déclaration variable listTxt
	for i, it := range c.Inventory { // Boucle pour afficher les objets de l'inventaire
		if i == "Potion" { // Si l'objet est potion
			txt = "Potion Basic" // Affichage de l'objet Potion Basic
		} else if i == "AdvancedPotion" { // Si l'objet est potion avancée
			txt = "Potion Avancé" // Affichage de l'objet Potion Avancé
		} else if i == "PoisonPotion" { // Si l'objet est potion poison
			txt = "Potion Poison" // Affichage de l'objet Potion Poison
		} else if i == "ForcePotion" { // Si l'objet est potion force
			txt = "Potion Force" // Affichage de l'objet Potion Force
		} else if i == "PlumeCorbeau" { // Si l'objet est plume de corbeau
			txt = "Plume de Corbeau" // Affichage de l'objet Plume de Corbeau
		} else if i == "FourrureLoup" { // Si l'objet est fourrure de loup
			txt = "Fourrure de Loup" // Affichage de l'objet Fourrure de Loup
		} else if i == "PeauTroll" { // Si l'objet est peau de troll
			txt = "Peau de Troll" // Affichage de l'objet Peau de Troll
		} else if i == "CuirSanglier" { // Si l'objet est cuir de sanglier
			txt = "Cuir de Sanglier" // Affichage de l'objet Cuir de Sanglier
		} else if i == "FildAraignee" { // Si l'objet est fil d'araignée
			txt = "Fil d'Araignée" // Affichage de l'objet Fil d'Araignée
		}
		listTxt = append(listTxt, "- "+txt+" - "+strconv.Itoa(it)+" disponible") // Ajout de l'objet dans la liste

		DisplayFrame(" Inventaire", listTxt) // Affichage cadre inventaire
	}
}
func (c *Character) DisplayMoney() { // Affichage argent
	DisplayFrame("Compte en banque", []string{ // Affichage cadre argent
		"Vous avez " + strconv.Itoa(c.Money) + "€"})
}

func (c *Character) DisplayFullInventory() { // Affichage inventaire complet
	DisplayFrame("Inventaire complet", []string{ // Affichage cadre inventaire complet
		"Vous avez trop d'item dans votre inventaire.",
		"Vous pouvez en jetez dans le menu inventaire."})
}

func DisplayFireBall() { // Affichage boule de feu

	DisplayFrame("Lance Boule de feu", []string{ // Affichage cadre boule de feu
		"Vous avez infligé x dégats"})
}

func DisplayLightning() { // Affichage éclair
	DisplayFrame("Lance Eclair", []string{ // Affichage cadre éclair
		"Vous avez infligé x dégats"})
}

func DisplayHealing() {
	DisplayFrame("Lance Soin", []string{ // Affichage cadre soin
		"Vous avez soigné x points de vie"})
}

func DisplayArrow() { // Affichage flèche
	DisplayFrame("Lance Flêche", []string{ // Affichage cadre flèche
		"Vous avez infligé x dégats"})
}

func DisplayGustArrow() { // Affichage rafale
	DisplayFrame("Lance Rafale", []string{ // Affichage cadre rafale
		"Vous avez infligé x dégats"})
}

func DisplayEscape() { // Affichage esquive
	DisplayFrame("Lance Esquive", []string{ // Affichage cadre esquive
		"Vous avez esquiver la prochaine attaque"})
}

func DisplayPunch() { // Affichage coup de poing
	DisplayFrame("Lance Coup de Poing", []string{ // Affichage cadre coup de poing
		"Vous avez infligé x dégats"})
}

func DisplayHeadButt() { // Affichage coup de tête
	DisplayFrame("Lance Coup de Tête", []string{ // Affichage cadre coup de tête
		"Vous avez infligé x dégats"})
}

func DisplayShield() { // Affichage bouclier
	DisplayFrame("Lance Bouclier", []string{ // Affichage cadre bouclier
		"Vous avez absorber x dégats"})
}

func DisplayDestruction() { // Affichage destruction
	DisplayFrame("Lance Destruction", []string{ // Affichage cadre destruction
		"Vous avez détruit votre adversaire"})
}

func DisplaySellItem() { // Affichage vente
	DisplayFrame("Quelle item jeter ?", []string{ // Affichage cadre vente
		"- 1. Potion Basic",
		"- 2. Potion avancée",
		"- 3. Potion de force",
		"- 4. Potion de poison",
		"- 5. Retour"})
}

func DisplaySellMenu() { // Affichage menu vente
	DisplayFrame("Que voulez vous jeter ?", []string{ // Affichage cadre menu vente
		"- 1. Un item",
		// "- 2. Un équipement",
		"- 5. Afficher l'inventaire",
		"- 6. Menu Principal"})
}

func (m *Mob) DisplayGoblinInfo() { // Affichage info goblin
	DisplayFrame("Goblin", []string{ // Affichage cadre info goblin
		"Vie : " + strconv.Itoa(m.CurrentHealth) + "/" + strconv.Itoa(m.Health)})
}

func (c *Character) DisplaySortObtenuBouleDeFeu() { // Affichage sort obtenu boule de feu
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu boule de feu
		" Vous avez obtenu le sort : Boule de Feu"})
}

func (c *Character) DisplaySortObtenuEclair() { // Affichage sort obtenu éclair
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu éclair
		" Vous avez obtenu le sort : Eclair"})
}

func (c *Character) DisplaySortObtenuAnneauSoin() { // Affichage sort obtenu anneau de soin
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu anneau de soin
		" Vous avez obtenu le sort : Soin"})
}

func (c *Character) DisplaySortObtenuFleche() { // Affichage sort obtenu flèche
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu flèche
		" Vous avez obtenu le sort : Flêche"})
}

func (c *Character) DisplaySortObtenuRafale() { // Affichage sort obtenu rafale
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu rafale
		" Vous avez obtenu le sort : Rafale"})
}

func (c *Character) DisplaySortObtenuEsquive() { // Affichage sort obtenu esquive
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu esquive
		" Vous avez obtenu le sort : Esquive"})
}

func (c *Character) DisplaySortObtenuCoupDePoing() { // Affichage sort obtenu coup de poing
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu coup de poing
		" Vous avez obtenu le sort : Coup de Poing"})
}

func (c *Character) DisplaySortObtenuCoupDeTete() { // Affichage sort obtenu coup de tête
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu coup de tête
		" Vous avez obtenu le sort : Coup de Tête"})
}

func (c *Character) DisplaySortObtenuBouclier() { // Affichage sort obtenu bouclier
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu bouclier
		" Vous avez obtenu le sort : Bouclier"})
}

func (c *Character) DisplaySortObtenuDestruction() { // Affichage sort obtenu destruction
	DisplayFrame(" Sort Obtenu", []string{ // Affichage cadre sort obtenu destruction
		" Vous avez obtenu le sort : Destruction"})
}

func (c *Character) DisplayMenuForgeron() { // Affichage menu forgeron
	DisplayFrame("Menu Forgeron", []string{ // Affichage cadre menu forgeron
		"- 1. Acheter un équipement",
		"- 2. Afficher l'inventaire",
		"- 3. Menu Principal"})
}

func DisplayListEquipment() { // Affichage liste équipement
	DisplayFrame("Liste des équipements", []string{ // Affichage cadre liste équipement
		"- 1. Casque || Plume de corbeau : 1, Cuir de sanglier : 1",
		"- 2. Plastron || Fourrure de loup : 2, Peau de troll : 1",
		"- 3. Botte || Fourrure de loup : 1, Cuir de sanglier : 1",
		"- 4. Retour"})
}

func DisplayFuir() { // Affichage fuir
	DisplayFrame("Fuir", []string{ // Affichage cadre fuir
		"Vous avez fuit le combat",
		"Vous avez perdu 10 de votre argent",
		"Vous avez réapparu avec la moitié de votre vie max"})
}

func DisplayImageAsciiDebut() { // Affichage image ascii début
	DisplayFrame("", []string{ // Affichage cadre image ascii début
		" _     _   ",
		"| |   (_)  ",
		"| |__  _  ___ _ ____   _____ _ __  _   _  ___ ",
		"| '_ \\| |/ _ \\ '_ \\ \\ / / _ \\ '_ \\| | | |/ _ \\",
		"| |_) | |  __/ | | \\ V /  __/ | | | |_| |  __/",
		"|_.__/|_|\\___|_| |_|\\_/ \\___|_| |_|\\__,_|\\___|"})
}

func DisplayImageAsciiFin() { // Affichage image ascii fin
	DisplayFrame("", []string{ // Affichage cadre image ascii fin
		"_______ _____ __   _      ______  _     _      _____ _______ _     _",
		"|______   |   | \\  |      |     \\ |     |        |   |______ |     |",
		"|       __|__ |  \\_|      |_____/ |_____|      __|   |______ |_____|"})
}

func (c *Character) StatsDisplay() { // Affichage stats
	DisplayFrame("Statistiques de "+c.Name, []string{ // Affichage cadre stats
		" Classe : " + c.Class,               // Affichage classe
		" Niveau : " + strconv.Itoa(c.Level), // Affichage niveau
		" Vie : " + strconv.Itoa(c.CurrentHealth) + "/" + strconv.Itoa(c.Health), // Affichage vie
		" Argent : " + strconv.Itoa(c.Money)})                                    // Affichage argent
}
