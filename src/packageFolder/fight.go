package packageFolder

import "strconv"

func (c *Character) PlayerDead() {
	var health = c.Health / 2
	if c.CurrentHealth == 0 || c.CurrentHealth < 0 {
		DisplayFrame("Mort", []string{
			"Vous avez péri.",
			"Vous avez réapparu avec " + strconv.Itoa(health) + " HP max."})
		c.CurrentHealth = c.Health / 2
	}
}

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

func (m *Character) AttackBasic() {

}

func (m *Character) AttackSpecial() {

}
