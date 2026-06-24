package payload

type AgencyDetailed struct {
	Abbrev string
	Administrator *string
	AttemptedLandings *int32
	AttemptedLandingsPayload *int32
	AttemptedLandingsSpacecraft *int32
	ConsecutiveSuccessfulLandings *int32
	ConsecutiveSuccessfulLaunches *int32
	Country []Country
	Description *string
	FailedLandings *int32
	FailedLandingsPayload *int32
	FailedLandingsSpacecraft *int32
	FailedLaunches *int32
	Featured bool
	FoundingYear *int32
	Id int32
	Image *Image
	InfoUrl *string
	Launchers string
	Logo *Image
	Name string
	Parent *string
	PendingLaunches *int32
	ResponseMode string
	SocialLogo *Image
	SocialMediaLinks []SocialMediaLink
	Spacecraft string
	SuccessfulLandings *int32
	SuccessfulLandingsPayload *int32
	SuccessfulLandingsSpacecraft *int32
	SuccessfulLaunches *int32
	TotalLaunchCount *int32
	TypeVal *AgencyType
	Url string
	WikiUrl *string
}

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

type MissionPatch struct {
	Agency *AgencyMini
	Id int32
	ImageUrl string
	Name string
	Priority int32
	ResponseMode string
}

type Payload struct {
	Cost *int32
	Description string
	Id int32
	Image *Image
	InfoLink string
	Manufacturer *AgencyDetailed
	Mass *float32
	Name string
	Operator *AgencyDetailed
	Program []ProgramNormal
	ResponseMode string
	TypeVal *PayloadType
	WikiLink string
}

type PayloadType struct {
	Id int32
	Name string
}

type ProgramNormal struct {
	Agencies []AgencyMini
	Description *string
	EndDate *string
	Id int32
	Image *Image
	InfoUrl *string
	MissionPatches []MissionPatch
	Name string
	ResponseMode string
	StartDate *string
	TypeVal *ProgramType
	Url string
	WikiUrl *string
}

type ProgramType struct {
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

type ListPayloadsRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListPayloadsResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Payload
}

type GetPayloadRequest struct {
	ID   int32
	Mode string
}
