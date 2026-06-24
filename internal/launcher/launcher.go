package launcher

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

type Launcher struct {
	AttemptedLandings *int32
	Details string
	FastestTurnaround *string
	FirstLaunchDate *string
	FlightProven bool
	Flights *int32
	Id int32
	Image *Image
	IsPlaceholder bool
	LastLaunchDate *string
	LauncherConfig *LauncherConfigList
	ResponseMode string
	SerialNumber *string
	Status *LauncherStatus
	SuccessfulLandings *int32
	Url string
}

type LauncherStatus struct {
	Id int32
	Name string
}

type ListLaunchersRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListLaunchersResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Launcher
}

type GetLauncherRequest struct {
	ID   int32
	Mode string
}
