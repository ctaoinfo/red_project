package packageFolder

// import (
// 	"fmt"
// 	"math/rand"
// 	"strings"
// 	"time"
// )

// func (c *Character) spiritTurn() {
// 	DisplayFrame("Tour de l'esprit de la forêt", []string{
// 		"Que voulez-vous faire ?",
// 		"1. - Attaquer",
// 		"2. - Utiliser un objet",
// 		"3. - Retour",
// 	})

// 	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
// 	m = strings.Replace(m, "\r\n", "", -1)
// 	switch m {
// 	case "1":
// 		c.spiritAttack()
// 	case "2":
// 		ClearLog()
// 		c.AccessInventory()
// 		c.useInventory(true)
// 	default:
// 		c.spiritTurn()
// 	}
// }

// func (c Character) spiritAttack() {
// 	ClearLog()
// 	fmt.Println(BIWhite + ">> Combat d'entraînement <<" + Reset)
// 	fmt.Println("1 >> Tempête verte")
// 	fmt.Println("2 >> Boule de feu")
// 	fmt.Println("3 >> Retour")
// 	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
// 	m = strings.Replace(m, "\r\n", "", -1)
// 	switch m {
// 	case "1":
// 		c.useStorm()
// 	case "2":
// 		c.useFireball("Esprit de la forêt")
// 	case "3":
// 		c.spiritTurn()
// 	default:
// 		c.spiritAttack()
// 	}
// }

// func (c Character) useStorm() {
// 	rand.Seed(time.Now().UnixNano())
// 	stormDmg := rand.Intn(11)
// 	if c.skill["Tempête verte"] >= 1 {
// 		fmt.Println(c.nom, "utilise", BIGreen+"Tempête verte"+Reset, "et inflige", stormDmg, "dégats à", M1.name)
// 		M1.HP -= stormDmg
// 		time.Sleep(400 * time.Millisecond)
// 		if M1.HP > 0 {
// 			stormDmg = rand.Intn(11)
// 			fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
// 			time.Sleep(2 * time.Second)
// 			fmt.Println("La", BIGreen+"Tempête verte"+Reset, "surgit à nouveau et fait", stormDmg, "dégats à", M1.name)
// 			M1.HP -= stormDmg
// 			if M1.HP > 0 {
// 				fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
// 			}
// 		}

// 	} else {
// 		fmt.Println(">> Impossible <<")
// 		c.spiritAttack()
// 	}
// }
////////////////////////////////////////////////
