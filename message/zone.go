package message

type ZoneReport struct {
	Header
	Layers []Layer
}

type Layer struct {
	ID    uint8
	Asset Asset
}

type Asset struct {
	ID          uint
	X           uint
	Y           uint
	StateLength uint
	State       interface{}
}
