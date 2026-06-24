package spacecraft

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

type CelestialBodyMini struct {
	Id int32
	Name string
	ResponseMode string
}

type CelestialBodyNormal struct {
	Atmosphere bool
	Description *string
	Diameter *float32
	Gravity *float32
	Id int32
	Image *Image
	LengthOfDay *string
	Mass *float32
	Name string
	ResponseMode string
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

type InfoURL struct {
	Description *string
	FeatureImage *string
	Language *Language
	Priority int32
	Source *string
	Title *string
	TypeVal *InfoURLType
	Url string
}

type InfoURLType struct {
	Id int32
	Name string
}

type Landing struct {
	Attempt bool
	Description string
	DownrangeDistance *float32
	Id int32
	LandingLocation *LandingLocation
	Success *bool
	TypeVal *LandingType
	Url string
}

type LandingLocation struct {
	Abbrev string
	Active bool
	AttemptedLandings *int32
	CelestialBody *CelestialBodyNormal
	Description *string
	FailedLandings *int32
	Id int32
	Image *Image
	Latitude *float32
	Location *LocationSerializerNoCelestialBody
	Longitude *float32
	Name string
	SuccessfulLandings *int32
}

type LandingType struct {
	Abbrev string
	Description *string
	Id int32
	Name string
}

type Language struct {
	Code string
	Id int32
	Name string
}

type LaunchNormal struct {
	AgencyLaunchAttemptCount *int32
	AgencyLaunchAttemptCountYear *int32
	Failreason *string
	Hashtag *string
	Id string
	Image *Image
	Infographic *string
	LastUpdated string
	LaunchDesignator *string
	LaunchServiceProvider *AgencyMini
	LocationLaunchAttemptCount *int32
	LocationLaunchAttemptCountYear *int32
	Mission *Mission
	Name string
	Net string
	NetPrecision *NetPrecision
	OrbitalLaunchAttemptCount *int32
	OrbitalLaunchAttemptCountYear *int32
	Pad *Pad
	PadLaunchAttemptCount *int32
	PadLaunchAttemptCountYear *int32
	Probability *int32
	Program []ProgramNormal
	ResponseMode string
	Rocket *RocketNormal
	Slug string
	Status *LaunchStatus
	Url string
	WeatherConcerns *string
	WebcastLive bool
	WindowEnd string
	WindowStart string
}

type LaunchStatus struct {
	Abbrev string
	Description string
	Id int32
	Name string
}

type LauncherConfigFamilyMini struct {
	Id int32
	Name string
	ResponseMode string
}

type LauncherConfigList struct {
	Families []LauncherConfigFamilyMini
	FullName string
	Id int32
	Name string
	ResponseMode string
	Url string
	Variant string
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
	ResponseMode string
	TimezoneName string
	TotalLandingCount *int32
	TotalLaunchCount *int32
	Url string
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

type Mission struct {
	Agencies []AgencyDetailed
	Description string
	Id int32
	Image *Image
	InfoUrls []InfoURL
	Name string
	Orbit *Orbit
	TypeVal string
	VidUrls []VidURL
}

type MissionPatch struct {
	Agency *AgencyMini
	Id int32
	ImageUrl string
	Name string
	Priority int32
	ResponseMode string
}

type NetPrecision struct {
	Abbrev string
	Description string
	Id int32
	Name string
}

type Orbit struct {
	Abbrev string
	CelestialBody *CelestialBodyMini
	Id int32
	Name string
}

type Pad struct {
	Active bool
	Agencies []AgencyNormal
	Country *Country
	Description *string
	FastestTurnaround *string
	Id int32
	Image *Image
	InfoUrl *string
	Latitude *float32
	Location *Location
	Longitude *float32
	MapImage *string
	MapUrl *string
	Name string
	OrbitalLaunchAttemptCount *int32
	TotalLaunchCount *int32
	Url string
	WikiUrl *string
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

type RocketNormal struct {
	Configuration *LauncherConfigList
	Id int32
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

type SpacecraftConfigNormal struct {
	Agency *AgencyMini
	Family []SpacecraftConfigFamilyNormal
	Id int32
	Image *Image
	InUse bool
	Name string
	ResponseMode string
	TypeVal *SpacecraftConfigType
	Url string
}

type SpacecraftConfigType struct {
	Id int32
	Name string
}

type Spacecraft struct {
	Description string
	FastestTurnaround *string
	Flights []SpacecraftFlightNormal
	FlightsCount *int32
	Id int32
	Image *Image
	InSpace bool
	IsPlaceholder bool
	MissionEndsCount *int32
	Name string
	ResponseMode string
	SerialNumber *string
	SpacecraftConfig *SpacecraftConfigDetailed
	Status *SpacecraftStatus
	TimeDocked *string
	TimeInSpace *string
	Url string
}

type SpacecraftFlightNormal struct {
	Destination *string
	Duration string
	Id int32
	Landing *Landing
	Launch *LaunchNormal
	MissionEnd *string
	ResponseMode string
	Spacecraft *SpacecraftNormal
	TurnAroundTime string
	Url string
}

type SpacecraftNormal struct {
	Description string
	FastestTurnaround *string
	FlightsCount *int32
	Id int32
	Image *Image
	InSpace bool
	IsPlaceholder bool
	MissionEndsCount *int32
	Name string
	ResponseMode string
	SerialNumber *string
	SpacecraftConfig *SpacecraftConfigNormal
	Status *SpacecraftStatus
	TimeDocked *string
	TimeInSpace *string
	Url string
}

type SpacecraftStatus struct {
	Id int32
	Name string
}

type VidURL struct {
	Description *string
	EndTime *string
	FeatureImage *string
	Language *Language
	Live bool
	Priority int32
	Publisher *string
	Source *string
	StartTime *string
	Title *string
	TypeVal *VidURLType
	Url string
}

type VidURLType struct {
	Id int32
	Name string
}

type ListSpacecraftsRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListSpacecraftsResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Spacecraft
}

type GetSpacecraftRequest struct {
	ID   int32
	Mode string
}
