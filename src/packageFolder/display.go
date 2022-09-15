package packageFolder

import "fmt"

func (c Character) DisplayInfo() {
	fmt.Println(c.Name)
}

func (c Character) AccessInventory() {
	fmt.Println(c.Inventory)
}
