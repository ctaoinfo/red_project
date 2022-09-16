package packageFolder

func RoundFight() {
	var Round int

	switch Round {
	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	case 6:

	case 7:

	case 8:

	case 9:

	case 101: // Level secret 1
	case 102: // Level secret 2
	case 103: // Level secret 3
	case 104: // Level secret 4
	}
}

// func (c Character) PlayerDead() {
// 	if c.currentHealth == 0 {
// 		if /* Totem présent dans l'inventaire */ {
// 			/* Enlever le totem de l'inventaire  */
// 			c.CurrentHealth = 50
// 		}else {
// 			fmt.Println("Vous avez malheureusement perdu tout vos HP")
// 		}
// 	}
// }

type Mob struct { // Strcture du personnage
	Name          string // Nom
	Rank          string // Rang
	Level         int    // Niveau
	Health        int    // Vie max
	CurrentHealth int    // Vie actuelle
}

func (m *Mob) InitMob(name string, rank string, level int, health int, currentHealth int) { // Fonction d'initialisation des Mobs
	m.Name = name
	m.Rank = rank
	m.Level = level
	m.Health = health
	m.CurrentHealth = currentHealth
}

func (m *Mob) AttackBasic() {

}

func (m *Mob) AttackSpecial() {

}
