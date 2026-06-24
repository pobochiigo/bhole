package astronaut

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

type Astronaut struct {
	Age *int32
	Agency *AgencyMini
	Bio string
	DateOfBirth *string
	DateOfDeath *string
	EvaTime string
	FirstFlight *string
	Id int32
	Image *Image
	InSpace bool
	LastFlight *string
	Name string
	Nationality []Country
	ResponseMode string
	SocialMediaLinks []SocialMediaLink
	Status *AstronautStatus
	TimeInSpace *string
	TypeVal *AstronautType
	Url string
	Wiki *string
}

type AstronautStatus struct {
	Id int32
	Name string
}

type AstronautType struct {
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

type SocialMedia struct {
	Id int32
	Logo *Image
	Name string
	Url *string
}

type SocialMediaLink struct {
	Id int32
	SocialMedia *SocialMedia
	Url *string
}

type ListAstronautsRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListAstronautsResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Astronaut
}

type GetAstronautRequest struct {
	ID   int32
	Mode string
}
