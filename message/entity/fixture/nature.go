package fixture

import "github.com/PaluMacil/walkabout/message"

type Plant struct {
	message.Asset
	Harvestable bool
}

func (p *Plant) Update() {

}
