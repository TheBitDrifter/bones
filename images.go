package bones

import "github.com/TheBitDrifter/warehouse"

type (
	ImageLocationsSmall [5]warehouse.CacheLocation
	ImageLocationsMed   [10]warehouse.CacheLocation
	ImageLocationsLarge [20]warehouse.CacheLocation
	ImageLocationsXL    [40]warehouse.CacheLocation
)

var (
	ImageLocationsSmallComponent = warehouse.FactoryNewComponent[ImageLocationsSmall]()
	ImageLocationsMedComponent   = warehouse.FactoryNewComponent[ImageLocationsMed]()
	ImageLocationsLargeComponent = warehouse.FactoryNewComponent[ImageLocationsLarge]()
	ImageLocationsXLComponent    = warehouse.FactoryNewComponent[ImageLocationsXL]()
)
