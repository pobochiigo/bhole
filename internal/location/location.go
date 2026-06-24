package location

type AgencyMini struct {
	Abbrev string
	Id int32
	Name string
	ResponseMode string
	TypeVal *AgencyType
	Url string
}

type AgencyType struct {
	Id int32
	Name string
}

type CelestialBodyDetailed struct {
	Atmosphere bool
	Description *string
	Diameter *float32
	FailedLandings int32
	FailedLaunches int32
	Gravity *float32
	Id int32
	Image *Image
	LengthOfDay *string
	Mass *float32
	Name string
	ResponseMode string
	SuccessfulLandings int32
	SuccessfulLaunches int32
	TotalAttemptedLandings int32
	TotalAttemptedLaunches int32
	TypeVal *CelestialBodyType
	WikiUrl *string
}

type CelestialBodyType struct {
	Id int32
	Name string
}

type Country struct {
	Alpha2Code string
	Alpha3Code string
	Id int32
	Name string
	NationalityName string
	NationalityNameComposed string
}

type Image struct {
	Credit *string
	Id int32
	ImageUrl string
	License *ImageLicense
	Name string
	SingleUse bool
	ThumbnailUrl string
	Variants []ImageVariant
}

type ImageLicense struct {
	Id int32
	Link *string
	Name string
	Priority int32
}

type ImageVariant struct {
	Id int32
	ImageUrl string
	TypeVal *ImageVariantType
}

type ImageVariantType struct {
	Id int32
	Name string
}

type Location struct {
	Active bool
	CelestialBody *CelestialBodyDetailed
	Country *Country
	Description *string
	Id int32
	Image *Image
	Latitude *float32
	Longitude *float32
	MapImage *string
	Name string
	Pads []PadSerializerNoLocation
	ResponseMode string
	TimezoneName string
	TotalLandingCount *int32
	TotalLaunchCount *int32
	Url string
}

type PadSerializerNoLocation struct {
	Active bool
	Agencies []AgencyMini
	Country *Country
	Description *string
	FastestTurnaround *string
	Id int32
	Image *Image
	InfoUrl *string
	Latitude *float32
	Longitude *float32
	MapImage *string
	MapUrl *string
	Name string
	OrbitalLaunchAttemptCount *int32
	TotalLaunchCount *int32
	Url string
	WikiUrl *string
}

type ListLocationsRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListLocationsResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Location
}

type GetLocationRequest struct {
	ID   int32
	Mode string
}
