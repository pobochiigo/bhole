package expedition

type AgencyMini struct {
	Abbrev string
	Id int32
	Name string
	ResponseMode string
	TypeVal *AgencyType
	Url string
}

type AgencyNormal struct {
	Abbrev string
	Administrator *string
	Country []Country
	Description *string
	Featured bool
	FoundingYear *int32
	Id int32
	Image *Image
	Launchers string
	Logo *Image
	Name string
	Parent *string
	ResponseMode string
	SocialLogo *Image
	Spacecraft string
	TypeVal *AgencyType
	Url string
}

type AgencyType struct {
	Id int32
	Name string
}

type AstronautDetailed struct {
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

type AstronautFlight struct {
	Astronaut *AstronautDetailed
	Id int32
	Role *AstronautRole
}

type AstronautRole struct {
	Id int32
	Priority int32
	Role string
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

type Expedition struct {
	Crew []AstronautFlight
	End *string
	Id int32
	MissionPatches []MissionPatch
	Name string
	ResponseMode string
	Spacestation *SpaceStationDetailed
	Spacewalks []SpacewalkList
	Start string
	Url string
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

type MissionPatch struct {
	Agency *AgencyMini
	Id int32
	ImageUrl string
	Name string
	Priority int32
	ResponseMode string
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

type SpaceStationDetailed struct {
	Deorbited *string
	Description string
	Founded string
	Id int32
	Image *Image
	Name string
	Orbit *string
	Owners []AgencyNormal
	Status *SpaceStationStatus
	TypeVal *SpaceStationType
	Url string
}

type SpaceStationStatus struct {
	Id int32
	Name string
}

type SpaceStationType struct {
	Id int32
	Name string
}

type SpacewalkList struct {
	Duration *string
	End *string
	Id int32
	Location *string
	Name string
	ResponseMode string
	Start *string
	Url string
}

type ListExpeditionsRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListExpeditionsResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Expedition
}

type GetExpeditionRequest struct {
	ID   int32
	Mode string
}
