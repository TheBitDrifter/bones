package bones

import "github.com/TheBitDrifter/warehouse"

type Position struct {
	X, Y float64
}

var PositionComponent = warehouse.FactoryNewComponent[Position]()
