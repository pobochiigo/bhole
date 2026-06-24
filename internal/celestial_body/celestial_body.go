package celestial_body

type CelestialBody struct {
	Atmosphere bool
	Description *string
	Diameter *float32
	FailedLandings int32
	FailedLaunches int32
	Gravity *float32
	Id int32
	Image *Image
	LengthOfDay *string
	Locations []LocationSerializerNoCelestialBody
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

type LocationSerializerNoCelestialBody struct {
	Active bool
	Country *Country
	Description *string
	Id int32
	Image *Image
	Latitude *float32
	Longitude *float32
	MapImage *string
	Name string
	ResponseMode string
	TimezoneName string
	TotalLandingCount *int32
	TotalLaunchCount *int32
	Url string
}

type ListCelestialBodiesRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListCelestialBodiesResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []CelestialBody
}

type GetCelestialBodyRequest struct {
	ID   int32
	Mode string
}
