package agency

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

type Agency struct {
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
	LauncherList []LauncherConfigDetailedSerializerNoManufacturer
	Launchers string
	Logo *Image
	Name string
	Parent *string
	PendingLaunches *int32
	ResponseMode string
	SocialLogo *Image
	SocialMediaLinks []SocialMediaLink
	Spacecraft string
	SpacecraftList []SpacecraftConfigDetailed
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

type LauncherConfigDetailedSerializerNoManufacturer struct {
	Active bool
	Alias string
	Apogee *float32
	AttemptedLandings *int32
	ConsecutiveSuccessfulLandings *int32
	ConsecutiveSuccessfulLaunches *int32
	Description string
	Diameter *float32
	FailedLandings *int32
	FailedLaunches *int32
	Families []LauncherConfigFamilyDetailed
	FastestTurnaround *string
	FullName string
	GeoCapacity *float32
	GtoCapacity *float32
	Id int32
	Image *Image
	InfoUrl *string
	IsPlaceholder bool
	LaunchCost *int32
	LaunchMass *float32
	Length *float32
	LeoCapacity *float32
	MaidenFlight *string
	MaxStage *int32
	MinStage *int32
	Name string
	PendingLaunches *int32
	Program []ProgramNormal
	ResponseMode string
	Reusable bool
	SsoCapacity *float32
	SuccessfulLandings *int32
	SuccessfulLaunches *int32
	ToThrust *float32
	TotalLaunchCount *int32
	Url string
	Variant string
	WikiUrl *string
}

type LauncherConfigFamilyDetailed struct {
	Active bool
	AttemptedLandings *int32
	ConsecutiveSuccessfulLandings *int32
	ConsecutiveSuccessfulLaunches *int32
	Description string
	FailedLandings *int32
	FailedLaunches *int32
	Id int32
	MaidenFlight *string
	Manufacturer []AgencyDetailed
	Name string
	Parent *LauncherConfigFamilyNormal
	PendingLaunches *int32
	ResponseMode string
	SuccessfulLandings *int32
	SuccessfulLaunches *int32
	TotalLaunchCount *int32
}

type LauncherConfigFamilyMini struct {
	Id int32
	Name string
	ResponseMode string
}

type LauncherConfigFamilyNormal struct {
	Id int32
	Manufacturer []AgencyNormal
	Name string
	Parent *LauncherConfigFamilyMini
	ResponseMode string
}

type MissionPatch struct {
	Agency *AgencyMini
	Id int32
	ImageUrl string
	Name string
	Priority int32
	ResponseMode string
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

type SpacecraftConfigDetailed struct {
	Agency *AgencyNormal
	AttemptedLandings *int32
	Capability string
	CrewCapacity *int32
	Details string
	Diameter *float32
	FailedLandings *int32
	FailedLaunches *int32
	Family []SpacecraftConfigFamilyDetailed
	FastestTurnaround *string
	FlightLife *string
	Height *float32
	History string
	HumanRated bool
	Id int32
	Image *Image
	InUse bool
	InfoLink string
	MaidenFlight *string
	Name string
	PayloadCapacity *int32
	PayloadReturnCapacity *int32
	ResponseMode string
	SpacecraftFlown *int32
	SuccessfulLandings *int32
	SuccessfulLaunches *int32
	TotalLaunchCount *int32
	TypeVal *SpacecraftConfigType
	Url string
	WikiLink string
}

type SpacecraftConfigFamilyDetailed struct {
	AttemptedLandings *int32
	Description string
	FailedLandings *int32
	FailedLaunches *int32
	Id int32
	MaidenFlight *string
	Manufacturer *AgencyNormal
	Name string
	Parent *SpacecraftConfigFamilyNormal
	ResponseMode string
	SpacecraftFlown *int32
	SuccessfulLandings *int32
	SuccessfulLaunches *int32
	TotalLaunchCount *int32
}

type SpacecraftConfigFamilyMini struct {
	Id int32
	Name string
	ResponseMode string
}

type SpacecraftConfigFamilyNormal struct {
	Description string
	Id int32
	MaidenFlight *string
	Manufacturer *AgencyMini
	Name string
	Parent *SpacecraftConfigFamilyMini
	ResponseMode string
}

type SpacecraftConfigType struct {
	Id int32
	Name string
}

type ListAgenciesRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListAgenciesResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Agency
}

type GetAgencyRequest struct {
	ID   int32
	Mode string
}
