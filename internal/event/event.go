package event

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

type AstronautNormal struct {
	Agency *AgencyMini
	Id int32
	Image *Image
	Name string
	Status *AstronautStatus
	Url string
}

type AstronautStatus struct {
	Id int32
	Name string
}

type Event struct {
	Agencies []AgencyMini
	Astronauts []AstronautNormal
	Date string
	DatePrecision *NetPrecision
	Description string
	Duration *string
	Expeditions []ExpeditionNormal
	Id int32
	Image *Image
	InfoUrls []InfoURL
	LastUpdated string
	Launches []LaunchBasic
	Location *string
	Name string
	Program []ProgramNormal
	ResponseMode string
	Slug string
	Spacestations []SpaceStationNormal
	TypeVal *EventType
	Updates []Update
	Url string
	VidUrls []VidURL
	WebcastLive bool
}

type EventType struct {
	Id int32
	Name string
}

type ExpeditionNormal struct {
	End *string
	Id int32
	MissionPatches []MissionPatch
	Name string
	ResponseMode string
	Spacestation *SpaceStationNormal
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

type Language struct {
	Code string
	Id int32
	Name string
}

type LaunchBasic struct {
	Id string
	Image *Image
	Infographic *string
	LastUpdated string
	LaunchDesignator *string
	Name string
	Net string
	NetPrecision *NetPrecision
	ResponseMode string
	Slug string
	Status *LaunchStatus
	Url string
	WindowEnd string
	WindowStart string
}

type LaunchStatus struct {
	Abbrev string
	Description string
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

type NetPrecision struct {
	Abbrev string
	Description string
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

type SpaceStationNormal struct {
	Deorbited *string
	Description string
	Founded string
	Id int32
	Image *Image
	Name string
	Orbit *string
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

type Update struct {
	Comment *string
	CreatedBy *string
	CreatedOn string
	Id int32
	InfoUrl *string
	ProfileImage *string
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

type ListEventsRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListEventsResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Event
}

type GetEventRequest struct {
	ID   int32
	Mode string
}
