package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	agencyv1 "com.gitlab/pobochiigo/bhole/proto/agency/v1"
	astronautv1 "com.gitlab/pobochiigo/bhole/proto/astronaut/v1"
	celestial_bodyv1 "com.gitlab/pobochiigo/bhole/proto/celestial_body/v1"
	docking_eventv1 "com.gitlab/pobochiigo/bhole/proto/docking_event/v1"
	eventv1 "com.gitlab/pobochiigo/bhole/proto/event/v1"
	expeditionv1 "com.gitlab/pobochiigo/bhole/proto/expedition/v1"
	landingv1 "com.gitlab/pobochiigo/bhole/proto/landing/v1"
	launchv1 "com.gitlab/pobochiigo/bhole/proto/launch/v1"
	launcherv1 "com.gitlab/pobochiigo/bhole/proto/launcher/v1"
	launcher_configurationv1 "com.gitlab/pobochiigo/bhole/proto/launcher_configuration/v1"
	locationv1 "com.gitlab/pobochiigo/bhole/proto/location/v1"
	padv1 "com.gitlab/pobochiigo/bhole/proto/pad/v1"
	payloadv1 "com.gitlab/pobochiigo/bhole/proto/payload/v1"
	programv1 "com.gitlab/pobochiigo/bhole/proto/program/v1"
	space_stationv1 "com.gitlab/pobochiigo/bhole/proto/space_station/v1"
	spacecraftv1 "com.gitlab/pobochiigo/bhole/proto/spacecraft/v1"
	spacewalkv1 "com.gitlab/pobochiigo/bhole/proto/spacewalk/v1"
	updatev1 "com.gitlab/pobochiigo/bhole/proto/update/v1"
)

type RESTClient struct {
	client  *http.Client
	baseURL string
}

func NewRESTClient(baseURL string, client *http.Client) *RESTClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &RESTClient{
		client:  client,
		baseURL: strings.TrimSuffix(baseURL, "/"),
	}
}
type AgencyDetailedJSON struct {
	Abbrev string `json:"abbrev"`
	Administrator *string `json:"administrator"`
	AttemptedLandings *int32 `json:"attempted_landings"`
	AttemptedLandingsPayload *int32 `json:"attempted_landings_payload"`
	AttemptedLandingsSpacecraft *int32 `json:"attempted_landings_spacecraft"`
	ConsecutiveSuccessfulLandings *int32 `json:"consecutive_successful_landings"`
	ConsecutiveSuccessfulLaunches *int32 `json:"consecutive_successful_launches"`
	Country []CountryJSON `json:"country"`
	Description *string `json:"description"`
	FailedLandings *int32 `json:"failed_landings"`
	FailedLandingsPayload *int32 `json:"failed_landings_payload"`
	FailedLandingsSpacecraft *int32 `json:"failed_landings_spacecraft"`
	FailedLaunches *int32 `json:"failed_launches"`
	Featured bool `json:"featured"`
	FoundingYear *int32 `json:"founding_year"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrl *string `json:"info_url"`
	Launchers string `json:"launchers"`
	Logo *ImageJSON `json:"logo"`
	Name string `json:"name"`
	Parent *string `json:"parent"`
	PendingLaunches *int32 `json:"pending_launches"`
	ResponseMode string `json:"response_mode"`
	SocialLogo *ImageJSON `json:"social_logo"`
	SocialMediaLinks []SocialMediaLinkJSON `json:"social_media_links"`
	Spacecraft string `json:"spacecraft"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	SuccessfulLandingsPayload *int32 `json:"successful_landings_payload"`
	SuccessfulLandingsSpacecraft *int32 `json:"successful_landings_spacecraft"`
	SuccessfulLaunches *int32 `json:"successful_launches"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	TypeVal *AgencyTypeJSON `json:"type"`
	Url string `json:"url"`
	WikiUrl *string `json:"wiki_url"`
}

type AgencyEndpointDetailedJSON struct {
	Abbrev string `json:"abbrev"`
	Administrator *string `json:"administrator"`
	AttemptedLandings *int32 `json:"attempted_landings"`
	AttemptedLandingsPayload *int32 `json:"attempted_landings_payload"`
	AttemptedLandingsSpacecraft *int32 `json:"attempted_landings_spacecraft"`
	ConsecutiveSuccessfulLandings *int32 `json:"consecutive_successful_landings"`
	ConsecutiveSuccessfulLaunches *int32 `json:"consecutive_successful_launches"`
	Country []CountryJSON `json:"country"`
	Description *string `json:"description"`
	FailedLandings *int32 `json:"failed_landings"`
	FailedLandingsPayload *int32 `json:"failed_landings_payload"`
	FailedLandingsSpacecraft *int32 `json:"failed_landings_spacecraft"`
	FailedLaunches *int32 `json:"failed_launches"`
	Featured bool `json:"featured"`
	FoundingYear *int32 `json:"founding_year"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrl *string `json:"info_url"`
	LauncherList []LauncherConfigDetailedSerializerNoManufacturerJSON `json:"launcher_list"`
	Launchers string `json:"launchers"`
	Logo *ImageJSON `json:"logo"`
	Name string `json:"name"`
	Parent *string `json:"parent"`
	PendingLaunches *int32 `json:"pending_launches"`
	ResponseMode string `json:"response_mode"`
	SocialLogo *ImageJSON `json:"social_logo"`
	SocialMediaLinks []SocialMediaLinkJSON `json:"social_media_links"`
	Spacecraft string `json:"spacecraft"`
	SpacecraftList []SpacecraftConfigDetailedJSON `json:"spacecraft_list"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	SuccessfulLandingsPayload *int32 `json:"successful_landings_payload"`
	SuccessfulLandingsSpacecraft *int32 `json:"successful_landings_spacecraft"`
	SuccessfulLaunches *int32 `json:"successful_launches"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	TypeVal *AgencyTypeJSON `json:"type"`
	Url string `json:"url"`
	WikiUrl *string `json:"wiki_url"`
}

type AgencyMiniJSON struct {
	Abbrev string `json:"abbrev"`
	Id int32 `json:"id"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	TypeVal *AgencyTypeJSON `json:"type"`
	Url string `json:"url"`
}

type AgencyNormalJSON struct {
	Abbrev string `json:"abbrev"`
	Administrator *string `json:"administrator"`
	Country []CountryJSON `json:"country"`
	Description *string `json:"description"`
	Featured bool `json:"featured"`
	FoundingYear *int32 `json:"founding_year"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Launchers string `json:"launchers"`
	Logo *ImageJSON `json:"logo"`
	Name string `json:"name"`
	Parent *string `json:"parent"`
	ResponseMode string `json:"response_mode"`
	SocialLogo *ImageJSON `json:"social_logo"`
	Spacecraft string `json:"spacecraft"`
	TypeVal *AgencyTypeJSON `json:"type"`
	Url string `json:"url"`
}

type AgencyTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type AstronautDetailedJSON struct {
	Age *int32 `json:"age"`
	Agency *AgencyMiniJSON `json:"agency"`
	Bio string `json:"bio"`
	DateOfBirth *string `json:"date_of_birth"`
	DateOfDeath *string `json:"date_of_death"`
	EvaTime string `json:"eva_time"`
	FirstFlight *string `json:"first_flight"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InSpace bool `json:"in_space"`
	LastFlight *string `json:"last_flight"`
	Name string `json:"name"`
	Nationality []CountryJSON `json:"nationality"`
	ResponseMode string `json:"response_mode"`
	SocialMediaLinks []SocialMediaLinkJSON `json:"social_media_links"`
	Status *AstronautStatusJSON `json:"status"`
	TimeInSpace *string `json:"time_in_space"`
	TypeVal *AstronautTypeJSON `json:"type"`
	Url string `json:"url"`
	Wiki *string `json:"wiki"`
}

type AstronautFlightJSON struct {
	Astronaut *AstronautDetailedJSON `json:"astronaut"`
	Id int32 `json:"id"`
	Role *AstronautRoleJSON `json:"role"`
}

type AstronautNormalJSON struct {
	Agency *AgencyMiniJSON `json:"agency"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Name string `json:"name"`
	Status *AstronautStatusJSON `json:"status"`
	Url string `json:"url"`
}

type AstronautRoleJSON struct {
	Id int32 `json:"id"`
	Priority int32 `json:"priority"`
	Role string `json:"role"`
}

type AstronautStatusJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type AstronautTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type CelestialBodyDetailedJSON struct {
	Atmosphere bool `json:"atmosphere"`
	Description *string `json:"description"`
	Diameter *float32 `json:"diameter"`
	FailedLandings int32 `json:"failed_landings"`
	FailedLaunches int32 `json:"failed_launches"`
	Gravity *float32 `json:"gravity"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	LengthOfDay *string `json:"length_of_day"`
	Mass *float32 `json:"mass"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	SuccessfulLandings int32 `json:"successful_landings"`
	SuccessfulLaunches int32 `json:"successful_launches"`
	TotalAttemptedLandings int32 `json:"total_attempted_landings"`
	TotalAttemptedLaunches int32 `json:"total_attempted_launches"`
	TypeVal *CelestialBodyTypeJSON `json:"type"`
	WikiUrl *string `json:"wiki_url"`
}

type CelestialBodyEndpointDetailedJSON struct {
	Atmosphere bool `json:"atmosphere"`
	Description *string `json:"description"`
	Diameter *float32 `json:"diameter"`
	FailedLandings int32 `json:"failed_landings"`
	FailedLaunches int32 `json:"failed_launches"`
	Gravity *float32 `json:"gravity"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	LengthOfDay *string `json:"length_of_day"`
	Locations []LocationSerializerNoCelestialBodyJSON `json:"locations"`
	Mass *float32 `json:"mass"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	SuccessfulLandings int32 `json:"successful_landings"`
	SuccessfulLaunches int32 `json:"successful_launches"`
	TotalAttemptedLandings int32 `json:"total_attempted_landings"`
	TotalAttemptedLaunches int32 `json:"total_attempted_launches"`
	TypeVal *CelestialBodyTypeJSON `json:"type"`
	WikiUrl *string `json:"wiki_url"`
}

type CelestialBodyMiniJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
}

type CelestialBodyNormalJSON struct {
	Atmosphere bool `json:"atmosphere"`
	Description *string `json:"description"`
	Diameter *float32 `json:"diameter"`
	Gravity *float32 `json:"gravity"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	LengthOfDay *string `json:"length_of_day"`
	Mass *float32 `json:"mass"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	TypeVal *CelestialBodyTypeJSON `json:"type"`
	WikiUrl *string `json:"wiki_url"`
}

type CelestialBodyTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type CountryJSON struct {
	Alpha2Code string `json:"alpha_2_code"`
	Alpha3Code string `json:"alpha_3_code"`
	Id int32 `json:"id"`
	Name string `json:"name"`
	NationalityName string `json:"nationality_name"`
	NationalityNameComposed string `json:"nationality_name_composed"`
}

type DockingEventDetailedSerializerForSpacestationJSON struct {
	Departure *string `json:"departure"`
	Docking string `json:"docking"`
	FlightVehicleChaser *SpacecraftFlightForDockingEventJSON `json:"flight_vehicle_chaser"`
	Id int32 `json:"id"`
	PayloadFlightChaser *PayloadFlightNormalJSON `json:"payload_flight_chaser"`
	SpaceStationChaser *SpaceStationNormalJSON `json:"space_station_chaser"`
	Url string `json:"url"`
}

type DockingEventEndpointDetailedJSON struct {
	Departure *string `json:"departure"`
	Docking string `json:"docking"`
	DockingLocation *DockingLocationJSON `json:"docking_location"`
	FlightVehicleChaser *SpacecraftFlightNormalJSON `json:"flight_vehicle_chaser"`
	FlightVehicleTarget *SpacecraftFlightMiniJSON `json:"flight_vehicle_target"`
	Id int32 `json:"id"`
	PayloadFlightChaser *PayloadFlightNormalJSON `json:"payload_flight_chaser"`
	PayloadFlightTarget *PayloadFlightMiniJSON `json:"payload_flight_target"`
	ResponseMode string `json:"response_mode"`
	SpaceStationChaser *SpaceStationNormalJSON `json:"space_station_chaser"`
	SpaceStationTarget *SpaceStationMiniJSON `json:"space_station_target"`
	Url string `json:"url"`
}

type DockingEventForChaserNormalJSON struct {
	Departure *string `json:"departure"`
	Docking string `json:"docking"`
	DockingLocation *DockingLocationJSON `json:"docking_location"`
	FlightVehicleTarget *SpacecraftFlightNormalJSON `json:"flight_vehicle_target"`
	Id int32 `json:"id"`
	PayloadFlightTarget *PayloadFlightNormalJSON `json:"payload_flight_target"`
	SpaceStationTarget *SpaceStationNormalJSON `json:"space_station_target"`
	Url string `json:"url"`
}

type DockingLocationJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Payload *PayloadMiniJSON `json:"payload"`
	Spacecraft *SpacecraftConfigNormalJSON `json:"spacecraft"`
	Spacestation *SpaceStationMiniJSON `json:"spacestation"`
}

type DockingLocationSerializerForSpacestationJSON struct {
	CurrentlyDocked *DockingEventDetailedSerializerForSpacestationJSON `json:"currently_docked"`
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type EventEndpointDetailedJSON struct {
	Agencies []AgencyMiniJSON `json:"agencies"`
	Astronauts []AstronautNormalJSON `json:"astronauts"`
	Date string `json:"date"`
	DatePrecision *NetPrecisionJSON `json:"date_precision"`
	Description string `json:"description"`
	Duration *string `json:"duration"`
	Expeditions []ExpeditionNormalJSON `json:"expeditions"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrls []InfoURLJSON `json:"info_urls"`
	LastUpdated string `json:"last_updated"`
	Launches []LaunchBasicJSON `json:"launches"`
	Location *string `json:"location"`
	Name string `json:"name"`
	Program []ProgramNormalJSON `json:"program"`
	ResponseMode string `json:"response_mode"`
	Slug string `json:"slug"`
	Spacestations []SpaceStationNormalJSON `json:"spacestations"`
	TypeVal *EventTypeJSON `json:"type"`
	Updates []UpdateJSON `json:"updates"`
	Url string `json:"url"`
	VidUrls []VidURLJSON `json:"vid_urls"`
	WebcastLive bool `json:"webcast_live"`
}

type EventNormalJSON struct {
	Date string `json:"date"`
	DatePrecision *NetPrecisionJSON `json:"date_precision"`
	Description string `json:"description"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrls []InfoURLJSON `json:"info_urls"`
	Location *string `json:"location"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	TypeVal *EventTypeJSON `json:"type"`
	Url string `json:"url"`
	VidUrls []VidURLJSON `json:"vid_urls"`
	WebcastLive bool `json:"webcast_live"`
}

type EventTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type ExpeditionDetailedJSON struct {
	Crew []AstronautFlightJSON `json:"crew"`
	End *string `json:"end"`
	Id int32 `json:"id"`
	MissionPatches []MissionPatchJSON `json:"mission_patches"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	Spacestation *SpaceStationDetailedJSON `json:"spacestation"`
	Spacewalks []SpacewalkListJSON `json:"spacewalks"`
	Start string `json:"start"`
	Url string `json:"url"`
}

type ExpeditionMiniJSON struct {
	End *string `json:"end"`
	Id int32 `json:"id"`
	Name string `json:"name"`
	Start string `json:"start"`
	Url string `json:"url"`
}

type ExpeditionNormalJSON struct {
	End *string `json:"end"`
	Id int32 `json:"id"`
	MissionPatches []MissionPatchJSON `json:"mission_patches"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	Spacestation *SpaceStationNormalJSON `json:"spacestation"`
	Spacewalks []SpacewalkListJSON `json:"spacewalks"`
	Start string `json:"start"`
	Url string `json:"url"`
}

type ExpeditionNormalSerializerForSpacewalkJSON struct {
	End *string `json:"end"`
	Id int32 `json:"id"`
	MissionPatches []MissionPatchJSON `json:"mission_patches"`
	Name string `json:"name"`
	Spacestation *SpaceStationNormalJSON `json:"spacestation"`
	Start string `json:"start"`
	Url string `json:"url"`
}

type FirstStageDetailedSerializerNoLandingJSON struct {
	Id int32 `json:"id"`
	Launcher *LauncherNormalJSON `json:"launcher"`
	LauncherFlightNumber *int32 `json:"launcher_flight_number"`
	PreviousFlight *LaunchNormalJSON `json:"previous_flight"`
	PreviousFlightDate *string `json:"previous_flight_date"`
	Reused *bool `json:"reused"`
	TurnAroundTime string `json:"turn_around_time"`
	TypeVal string `json:"type"`
}

type FirstStageNormalJSON struct {
	Id int32 `json:"id"`
	Landing *LandingJSON `json:"landing"`
	Launcher *LauncherNormalJSON `json:"launcher"`
	LauncherFlightNumber *int32 `json:"launcher_flight_number"`
	PreviousFlight *LaunchMiniJSON `json:"previous_flight"`
	PreviousFlightDate *string `json:"previous_flight_date"`
	Reused *bool `json:"reused"`
	TurnAroundTime string `json:"turn_around_time"`
	TypeVal string `json:"type"`
}

type ImageJSON struct {
	Credit *string `json:"credit"`
	Id int32 `json:"id"`
	ImageUrl string `json:"image_url"`
	License *ImageLicenseJSON `json:"license"`
	Name string `json:"name"`
	SingleUse bool `json:"single_use"`
	ThumbnailUrl string `json:"thumbnail_url"`
	Variants []ImageVariantJSON `json:"variants"`
}

type ImageLicenseJSON struct {
	Id int32 `json:"id"`
	Link *string `json:"link"`
	Name string `json:"name"`
	Priority int32 `json:"priority"`
}

type ImageVariantJSON struct {
	Id int32 `json:"id"`
	ImageUrl string `json:"image_url"`
	TypeVal *ImageVariantTypeJSON `json:"type"`
}

type ImageVariantTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type InfoURLJSON struct {
	Description *string `json:"description"`
	FeatureImage *string `json:"feature_image"`
	Language *LanguageJSON `json:"language"`
	Priority int32 `json:"priority"`
	Source *string `json:"source"`
	Title *string `json:"title"`
	TypeVal *InfoURLTypeJSON `json:"type"`
	Url string `json:"url"`
}

type InfoURLTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type LandingJSON struct {
	Attempt bool `json:"attempt"`
	Description string `json:"description"`
	DownrangeDistance *float32 `json:"downrange_distance"`
	Id int32 `json:"id"`
	LandingLocation *LandingLocationJSON `json:"landing_location"`
	Success *bool `json:"success"`
	TypeVal *LandingTypeJSON `json:"type"`
	Url string `json:"url"`
}

type LandingEndpointDetailedJSON struct {
	Attempt bool `json:"attempt"`
	Description string `json:"description"`
	DownrangeDistance *float32 `json:"downrange_distance"`
	Firststage *FirstStageDetailedSerializerNoLandingJSON `json:"firststage"`
	Id int32 `json:"id"`
	LandingLocation *LandingLocationJSON `json:"landing_location"`
	Payloadflight *PayloadFlightDetailedSerializerNoLandingJSON `json:"payloadflight"`
	ResponseMode string `json:"response_mode"`
	Spacecraftflight *SpacecraftFlightDetailedSerializerNoLandingJSON `json:"spacecraftflight"`
	Success *bool `json:"success"`
	TypeVal *LandingTypeJSON `json:"type"`
	Url string `json:"url"`
}

type LandingLocationJSON struct {
	Abbrev string `json:"abbrev"`
	Active bool `json:"active"`
	AttemptedLandings *int32 `json:"attempted_landings"`
	CelestialBody *CelestialBodyNormalJSON `json:"celestial_body"`
	Description *string `json:"description"`
	FailedLandings *int32 `json:"failed_landings"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Latitude *float32 `json:"latitude"`
	Location *LocationSerializerNoCelestialBodyJSON `json:"location"`
	Longitude *float32 `json:"longitude"`
	Name string `json:"name"`
	SuccessfulLandings *int32 `json:"successful_landings"`
}

type LandingTypeJSON struct {
	Abbrev string `json:"abbrev"`
	Description *string `json:"description"`
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type LanguageJSON struct {
	Code string `json:"code"`
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type LaunchBasicJSON struct {
	Id string `json:"id"`
	Image *ImageJSON `json:"image"`
	Infographic *string `json:"infographic"`
	LastUpdated string `json:"last_updated"`
	LaunchDesignator *string `json:"launch_designator"`
	Name string `json:"name"`
	Net string `json:"net"`
	NetPrecision *NetPrecisionJSON `json:"net_precision"`
	ResponseMode string `json:"response_mode"`
	Slug string `json:"slug"`
	Status *LaunchStatusJSON `json:"status"`
	Url string `json:"url"`
	WindowEnd string `json:"window_end"`
	WindowStart string `json:"window_start"`
}

type LaunchDetailedJSON struct {
	AgencyLaunchAttemptCount *int32 `json:"agency_launch_attempt_count"`
	AgencyLaunchAttemptCountYear *int32 `json:"agency_launch_attempt_count_year"`
	Failreason *string `json:"failreason"`
	FlightclubUrl *string `json:"flightclub_url"`
	Hashtag *string `json:"hashtag"`
	Id string `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrls []InfoURLJSON `json:"info_urls"`
	Infographic *string `json:"infographic"`
	LastUpdated string `json:"last_updated"`
	LaunchDesignator *string `json:"launch_designator"`
	LaunchServiceProvider *AgencyDetailedJSON `json:"launch_service_provider"`
	LocationLaunchAttemptCount *int32 `json:"location_launch_attempt_count"`
	LocationLaunchAttemptCountYear *int32 `json:"location_launch_attempt_count_year"`
	Mission *MissionJSON `json:"mission"`
	MissionPatches []MissionPatchJSON `json:"mission_patches"`
	Name string `json:"name"`
	Net string `json:"net"`
	NetPrecision *NetPrecisionJSON `json:"net_precision"`
	OrbitalLaunchAttemptCount *int32 `json:"orbital_launch_attempt_count"`
	OrbitalLaunchAttemptCountYear *int32 `json:"orbital_launch_attempt_count_year"`
	Pad *PadJSON `json:"pad"`
	PadLaunchAttemptCount *int32 `json:"pad_launch_attempt_count"`
	PadLaunchAttemptCountYear *int32 `json:"pad_launch_attempt_count_year"`
	PadTurnaround string `json:"pad_turnaround"`
	Probability *int32 `json:"probability"`
	Program []ProgramNormalJSON `json:"program"`
	ResponseMode string `json:"response_mode"`
	Rocket *RocketDetailedJSON `json:"rocket"`
	Slug string `json:"slug"`
	Status *LaunchStatusJSON `json:"status"`
	Timeline []TimelineEventJSON `json:"timeline"`
	Updates []UpdateJSON `json:"updates"`
	Url string `json:"url"`
	VidUrls []VidURLJSON `json:"vid_urls"`
	WeatherConcerns *string `json:"weather_concerns"`
	WebcastLive bool `json:"webcast_live"`
	WindowEnd string `json:"window_end"`
	WindowStart string `json:"window_start"`
}

type LaunchMiniJSON struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
}

type LaunchNormalJSON struct {
	AgencyLaunchAttemptCount *int32 `json:"agency_launch_attempt_count"`
	AgencyLaunchAttemptCountYear *int32 `json:"agency_launch_attempt_count_year"`
	Failreason *string `json:"failreason"`
	Hashtag *string `json:"hashtag"`
	Id string `json:"id"`
	Image *ImageJSON `json:"image"`
	Infographic *string `json:"infographic"`
	LastUpdated string `json:"last_updated"`
	LaunchDesignator *string `json:"launch_designator"`
	LaunchServiceProvider *AgencyMiniJSON `json:"launch_service_provider"`
	LocationLaunchAttemptCount *int32 `json:"location_launch_attempt_count"`
	LocationLaunchAttemptCountYear *int32 `json:"location_launch_attempt_count_year"`
	Mission *MissionJSON `json:"mission"`
	Name string `json:"name"`
	Net string `json:"net"`
	NetPrecision *NetPrecisionJSON `json:"net_precision"`
	OrbitalLaunchAttemptCount *int32 `json:"orbital_launch_attempt_count"`
	OrbitalLaunchAttemptCountYear *int32 `json:"orbital_launch_attempt_count_year"`
	Pad *PadJSON `json:"pad"`
	PadLaunchAttemptCount *int32 `json:"pad_launch_attempt_count"`
	PadLaunchAttemptCountYear *int32 `json:"pad_launch_attempt_count_year"`
	Probability *int32 `json:"probability"`
	Program []ProgramNormalJSON `json:"program"`
	ResponseMode string `json:"response_mode"`
	Rocket *RocketNormalJSON `json:"rocket"`
	Slug string `json:"slug"`
	Status *LaunchStatusJSON `json:"status"`
	Url string `json:"url"`
	WeatherConcerns *string `json:"weather_concerns"`
	WebcastLive bool `json:"webcast_live"`
	WindowEnd string `json:"window_end"`
	WindowStart string `json:"window_start"`
}

type LaunchStatusJSON struct {
	Abbrev string `json:"abbrev"`
	Description string `json:"description"`
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type LauncherConfigDetailedJSON struct {
	Active bool `json:"active"`
	Alias string `json:"alias"`
	Apogee *float32 `json:"apogee"`
	AttemptedLandings *int32 `json:"attempted_landings"`
	ConsecutiveSuccessfulLandings *int32 `json:"consecutive_successful_landings"`
	ConsecutiveSuccessfulLaunches *int32 `json:"consecutive_successful_launches"`
	Description string `json:"description"`
	Diameter *float32 `json:"diameter"`
	FailedLandings *int32 `json:"failed_landings"`
	FailedLaunches *int32 `json:"failed_launches"`
	Families []LauncherConfigFamilyDetailedJSON `json:"families"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	FullName string `json:"full_name"`
	GeoCapacity *float32 `json:"geo_capacity"`
	GtoCapacity *float32 `json:"gto_capacity"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrl *string `json:"info_url"`
	IsPlaceholder bool `json:"is_placeholder"`
	LaunchCost *int32 `json:"launch_cost"`
	LaunchMass *float32 `json:"launch_mass"`
	Length *float32 `json:"length"`
	LeoCapacity *float32 `json:"leo_capacity"`
	MaidenFlight *string `json:"maiden_flight"`
	Manufacturer *AgencyDetailedJSON `json:"manufacturer"`
	MaxStage *int32 `json:"max_stage"`
	MinStage *int32 `json:"min_stage"`
	Name string `json:"name"`
	PendingLaunches *int32 `json:"pending_launches"`
	Program []ProgramNormalJSON `json:"program"`
	ResponseMode string `json:"response_mode"`
	Reusable bool `json:"reusable"`
	SsoCapacity *float32 `json:"sso_capacity"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	SuccessfulLaunches *int32 `json:"successful_launches"`
	ToThrust *float32 `json:"to_thrust"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	Url string `json:"url"`
	Variant string `json:"variant"`
	WikiUrl *string `json:"wiki_url"`
}

type LauncherConfigDetailedSerializerNoManufacturerJSON struct {
	Active bool `json:"active"`
	Alias string `json:"alias"`
	Apogee *float32 `json:"apogee"`
	AttemptedLandings *int32 `json:"attempted_landings"`
	ConsecutiveSuccessfulLandings *int32 `json:"consecutive_successful_landings"`
	ConsecutiveSuccessfulLaunches *int32 `json:"consecutive_successful_launches"`
	Description string `json:"description"`
	Diameter *float32 `json:"diameter"`
	FailedLandings *int32 `json:"failed_landings"`
	FailedLaunches *int32 `json:"failed_launches"`
	Families []LauncherConfigFamilyDetailedJSON `json:"families"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	FullName string `json:"full_name"`
	GeoCapacity *float32 `json:"geo_capacity"`
	GtoCapacity *float32 `json:"gto_capacity"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrl *string `json:"info_url"`
	IsPlaceholder bool `json:"is_placeholder"`
	LaunchCost *int32 `json:"launch_cost"`
	LaunchMass *float32 `json:"launch_mass"`
	Length *float32 `json:"length"`
	LeoCapacity *float32 `json:"leo_capacity"`
	MaidenFlight *string `json:"maiden_flight"`
	MaxStage *int32 `json:"max_stage"`
	MinStage *int32 `json:"min_stage"`
	Name string `json:"name"`
	PendingLaunches *int32 `json:"pending_launches"`
	Program []ProgramNormalJSON `json:"program"`
	ResponseMode string `json:"response_mode"`
	Reusable bool `json:"reusable"`
	SsoCapacity *float32 `json:"sso_capacity"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	SuccessfulLaunches *int32 `json:"successful_launches"`
	ToThrust *float32 `json:"to_thrust"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	Url string `json:"url"`
	Variant string `json:"variant"`
	WikiUrl *string `json:"wiki_url"`
}

type LauncherConfigFamilyDetailedJSON struct {
	Active bool `json:"active"`
	AttemptedLandings *int32 `json:"attempted_landings"`
	ConsecutiveSuccessfulLandings *int32 `json:"consecutive_successful_landings"`
	ConsecutiveSuccessfulLaunches *int32 `json:"consecutive_successful_launches"`
	Description string `json:"description"`
	FailedLandings *int32 `json:"failed_landings"`
	FailedLaunches *int32 `json:"failed_launches"`
	Id int32 `json:"id"`
	MaidenFlight *string `json:"maiden_flight"`
	Manufacturer []AgencyDetailedJSON `json:"manufacturer"`
	Name string `json:"name"`
	Parent *LauncherConfigFamilyNormalJSON `json:"parent"`
	PendingLaunches *int32 `json:"pending_launches"`
	ResponseMode string `json:"response_mode"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	SuccessfulLaunches *int32 `json:"successful_launches"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
}

type LauncherConfigFamilyMiniJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
}

type LauncherConfigFamilyNormalJSON struct {
	Id int32 `json:"id"`
	Manufacturer []AgencyNormalJSON `json:"manufacturer"`
	Name string `json:"name"`
	Parent *LauncherConfigFamilyMiniJSON `json:"parent"`
	ResponseMode string `json:"response_mode"`
}

type LauncherConfigListJSON struct {
	Families []LauncherConfigFamilyMiniJSON `json:"families"`
	FullName string `json:"full_name"`
	Id int32 `json:"id"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	Url string `json:"url"`
	Variant string `json:"variant"`
}

type LauncherDetailedJSON struct {
	AttemptedLandings *int32 `json:"attempted_landings"`
	Details string `json:"details"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	FirstLaunchDate *string `json:"first_launch_date"`
	FlightProven bool `json:"flight_proven"`
	Flights *int32 `json:"flights"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	IsPlaceholder bool `json:"is_placeholder"`
	LastLaunchDate *string `json:"last_launch_date"`
	LauncherConfig *LauncherConfigListJSON `json:"launcher_config"`
	ResponseMode string `json:"response_mode"`
	SerialNumber *string `json:"serial_number"`
	Status *LauncherStatusJSON `json:"status"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	Url string `json:"url"`
}

type LauncherNormalJSON struct {
	AttemptedLandings *int32 `json:"attempted_landings"`
	Details string `json:"details"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	FirstLaunchDate *string `json:"first_launch_date"`
	FlightProven bool `json:"flight_proven"`
	Flights *int32 `json:"flights"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	IsPlaceholder bool `json:"is_placeholder"`
	LastLaunchDate *string `json:"last_launch_date"`
	ResponseMode string `json:"response_mode"`
	SerialNumber *string `json:"serial_number"`
	Status *LauncherStatusJSON `json:"status"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	Url string `json:"url"`
}

type LauncherStatusJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type LocationJSON struct {
	Active bool `json:"active"`
	CelestialBody *CelestialBodyDetailedJSON `json:"celestial_body"`
	Country *CountryJSON `json:"country"`
	Description *string `json:"description"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Latitude *float32 `json:"latitude"`
	Longitude *float32 `json:"longitude"`
	MapImage *string `json:"map_image"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	TimezoneName string `json:"timezone_name"`
	TotalLandingCount *int32 `json:"total_landing_count"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	Url string `json:"url"`
}

type LocationSerializerNoCelestialBodyJSON struct {
	Active bool `json:"active"`
	Country *CountryJSON `json:"country"`
	Description *string `json:"description"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Latitude *float32 `json:"latitude"`
	Longitude *float32 `json:"longitude"`
	MapImage *string `json:"map_image"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	TimezoneName string `json:"timezone_name"`
	TotalLandingCount *int32 `json:"total_landing_count"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	Url string `json:"url"`
}

type LocationSerializerWithPadsJSON struct {
	Active bool `json:"active"`
	CelestialBody *CelestialBodyDetailedJSON `json:"celestial_body"`
	Country *CountryJSON `json:"country"`
	Description *string `json:"description"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Latitude *float32 `json:"latitude"`
	Longitude *float32 `json:"longitude"`
	MapImage *string `json:"map_image"`
	Name string `json:"name"`
	Pads []PadSerializerNoLocationJSON `json:"pads"`
	ResponseMode string `json:"response_mode"`
	TimezoneName string `json:"timezone_name"`
	TotalLandingCount *int32 `json:"total_landing_count"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	Url string `json:"url"`
}

type MissionJSON struct {
	Agencies []AgencyDetailedJSON `json:"agencies"`
	Description string `json:"description"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrls []InfoURLJSON `json:"info_urls"`
	Name string `json:"name"`
	Orbit *OrbitJSON `json:"orbit"`
	TypeVal string `json:"type"`
	VidUrls []VidURLJSON `json:"vid_urls"`
}

type MissionPatchJSON struct {
	Agency *AgencyMiniJSON `json:"agency"`
	Id int32 `json:"id"`
	ImageUrl string `json:"image_url"`
	Name string `json:"name"`
	Priority int32 `json:"priority"`
	ResponseMode string `json:"response_mode"`
}

type NetPrecisionJSON struct {
	Abbrev string `json:"abbrev"`
	Description string `json:"description"`
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type OrbitJSON struct {
	Abbrev string `json:"abbrev"`
	CelestialBody *CelestialBodyMiniJSON `json:"celestial_body"`
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type PadJSON struct {
	Active bool `json:"active"`
	Agencies []AgencyNormalJSON `json:"agencies"`
	Country *CountryJSON `json:"country"`
	Description *string `json:"description"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrl *string `json:"info_url"`
	Latitude *float32 `json:"latitude"`
	Location *LocationJSON `json:"location"`
	Longitude *float32 `json:"longitude"`
	MapImage *string `json:"map_image"`
	MapUrl *string `json:"map_url"`
	Name string `json:"name"`
	OrbitalLaunchAttemptCount *int32 `json:"orbital_launch_attempt_count"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	Url string `json:"url"`
	WikiUrl *string `json:"wiki_url"`
}

type PadSerializerNoLocationJSON struct {
	Active bool `json:"active"`
	Agencies []AgencyMiniJSON `json:"agencies"`
	Country *CountryJSON `json:"country"`
	Description *string `json:"description"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrl *string `json:"info_url"`
	Latitude *float32 `json:"latitude"`
	Longitude *float32 `json:"longitude"`
	MapImage *string `json:"map_image"`
	MapUrl *string `json:"map_url"`
	Name string `json:"name"`
	OrbitalLaunchAttemptCount *int32 `json:"orbital_launch_attempt_count"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	Url string `json:"url"`
	WikiUrl *string `json:"wiki_url"`
}

type PayloadDetailedJSON struct {
	Cost *int32 `json:"cost"`
	Description string `json:"description"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoLink string `json:"info_link"`
	Manufacturer *AgencyDetailedJSON `json:"manufacturer"`
	Mass *float32 `json:"mass"`
	Name string `json:"name"`
	Operator *AgencyDetailedJSON `json:"operator"`
	Program []ProgramNormalJSON `json:"program"`
	ResponseMode string `json:"response_mode"`
	TypeVal *PayloadTypeJSON `json:"type"`
	WikiLink string `json:"wiki_link"`
}

type PayloadFlightDetailedSerializerNoLandingJSON struct {
	Amount int32 `json:"amount"`
	Destination *string `json:"destination"`
	DockingEvents []DockingEventForChaserNormalJSON `json:"docking_events"`
	Id int32 `json:"id"`
	Launch *LaunchNormalJSON `json:"launch"`
	Payload *PayloadDetailedJSON `json:"payload"`
	ResponseMode string `json:"response_mode"`
	Url string `json:"url"`
}

type PayloadFlightMiniJSON struct {
	Amount int32 `json:"amount"`
	Destination *string `json:"destination"`
	Id int32 `json:"id"`
	Landing *LandingJSON `json:"landing"`
	Launch *LaunchMiniJSON `json:"launch"`
	Payload *PayloadMiniJSON `json:"payload"`
	ResponseMode string `json:"response_mode"`
	Url string `json:"url"`
}

type PayloadFlightNormalJSON struct {
	Amount int32 `json:"amount"`
	Destination *string `json:"destination"`
	Id int32 `json:"id"`
	Landing *LandingJSON `json:"landing"`
	Launch *LaunchNormalJSON `json:"launch"`
	Payload *PayloadNormalJSON `json:"payload"`
	ResponseMode string `json:"response_mode"`
	Url string `json:"url"`
}

type PayloadFlightSerializerNoLaunchJSON struct {
	Amount int32 `json:"amount"`
	Destination *string `json:"destination"`
	DockingEvents []DockingEventForChaserNormalJSON `json:"docking_events"`
	Id int32 `json:"id"`
	Landing *LandingJSON `json:"landing"`
	Payload *PayloadDetailedJSON `json:"payload"`
	ResponseMode string `json:"response_mode"`
	Url string `json:"url"`
}

type PayloadMiniJSON struct {
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Manufacturer *AgencyMiniJSON `json:"manufacturer"`
	Name string `json:"name"`
	Operator *AgencyMiniJSON `json:"operator"`
	ResponseMode string `json:"response_mode"`
	TypeVal *PayloadTypeJSON `json:"type"`
}

type PayloadNormalJSON struct {
	Cost *int32 `json:"cost"`
	Description string `json:"description"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoLink string `json:"info_link"`
	Manufacturer *AgencyNormalJSON `json:"manufacturer"`
	Mass *float32 `json:"mass"`
	Name string `json:"name"`
	Operator *AgencyNormalJSON `json:"operator"`
	Program []ProgramMiniJSON `json:"program"`
	ResponseMode string `json:"response_mode"`
	TypeVal *PayloadTypeJSON `json:"type"`
	WikiLink string `json:"wiki_link"`
}

type PayloadTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type ProgramMiniJSON struct {
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrl *string `json:"info_url"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	Url string `json:"url"`
	WikiUrl *string `json:"wiki_url"`
}

type ProgramNormalJSON struct {
	Agencies []AgencyMiniJSON `json:"agencies"`
	Description *string `json:"description"`
	EndDate *string `json:"end_date"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InfoUrl *string `json:"info_url"`
	MissionPatches []MissionPatchJSON `json:"mission_patches"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	StartDate *string `json:"start_date"`
	TypeVal *ProgramTypeJSON `json:"type"`
	Url string `json:"url"`
	WikiUrl *string `json:"wiki_url"`
}

type ProgramTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type RocketDetailedJSON struct {
	Configuration *LauncherConfigDetailedJSON `json:"configuration"`
	Id int32 `json:"id"`
	LauncherStage []FirstStageNormalJSON `json:"launcher_stage"`
	Payloads []PayloadFlightSerializerNoLaunchJSON `json:"payloads"`
	SpacecraftStage []SpacecraftFlightDetailedSerializerNoLaunchJSON `json:"spacecraft_stage"`
}

type RocketNormalJSON struct {
	Configuration *LauncherConfigListJSON `json:"configuration"`
	Id int32 `json:"id"`
}

type SocialMediaJSON struct {
	Id int32 `json:"id"`
	Logo *ImageJSON `json:"logo"`
	Name string `json:"name"`
	Url *string `json:"url"`
}

type SocialMediaLinkJSON struct {
	Id int32 `json:"id"`
	SocialMedia *SocialMediaJSON `json:"social_media"`
	Url *string `json:"url"`
}

type SpaceStationDetailedJSON struct {
	Deorbited *string `json:"deorbited"`
	Description string `json:"description"`
	Founded string `json:"founded"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Name string `json:"name"`
	Orbit *string `json:"orbit"`
	Owners []AgencyNormalJSON `json:"owners"`
	Status *SpaceStationStatusJSON `json:"status"`
	TypeVal *SpaceStationTypeJSON `json:"type"`
	Url string `json:"url"`
}

type SpaceStationDetailedEndpointJSON struct {
	ActiveDockingEvents []DockingEventForChaserNormalJSON `json:"active_docking_events"`
	ActiveExpeditions []ExpeditionMiniJSON `json:"active_expeditions"`
	Deorbited *string `json:"deorbited"`
	Description string `json:"description"`
	DockedVehicles *int32 `json:"docked_vehicles"`
	DockingLocation []DockingLocationSerializerForSpacestationJSON `json:"docking_location"`
	Founded string `json:"founded"`
	Height *float32 `json:"height"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Mass *float32 `json:"mass"`
	Name string `json:"name"`
	OnboardCrew *int32 `json:"onboard_crew"`
	Orbit *string `json:"orbit"`
	Owners []AgencyNormalJSON `json:"owners"`
	ResponseMode string `json:"response_mode"`
	Status *SpaceStationStatusJSON `json:"status"`
	TypeVal *SpaceStationTypeJSON `json:"type"`
	Url string `json:"url"`
	Volume *int32 `json:"volume"`
	Width *float32 `json:"width"`
}

type SpaceStationMiniJSON struct {
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Name string `json:"name"`
	Url string `json:"url"`
}

type SpaceStationNormalJSON struct {
	Deorbited *string `json:"deorbited"`
	Description string `json:"description"`
	Founded string `json:"founded"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	Name string `json:"name"`
	Orbit *string `json:"orbit"`
	Status *SpaceStationStatusJSON `json:"status"`
	TypeVal *SpaceStationTypeJSON `json:"type"`
	Url string `json:"url"`
}

type SpaceStationStatusJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type SpaceStationTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type SpacecraftConfigDetailedJSON struct {
	Agency *AgencyNormalJSON `json:"agency"`
	AttemptedLandings *int32 `json:"attempted_landings"`
	Capability string `json:"capability"`
	CrewCapacity *int32 `json:"crew_capacity"`
	Details string `json:"details"`
	Diameter *float32 `json:"diameter"`
	FailedLandings *int32 `json:"failed_landings"`
	FailedLaunches *int32 `json:"failed_launches"`
	Family []SpacecraftConfigFamilyDetailedJSON `json:"family"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	FlightLife *string `json:"flight_life"`
	Height *float32 `json:"height"`
	History string `json:"history"`
	HumanRated bool `json:"human_rated"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InUse bool `json:"in_use"`
	InfoLink string `json:"info_link"`
	MaidenFlight *string `json:"maiden_flight"`
	Name string `json:"name"`
	PayloadCapacity *int32 `json:"payload_capacity"`
	PayloadReturnCapacity *int32 `json:"payload_return_capacity"`
	ResponseMode string `json:"response_mode"`
	SpacecraftFlown *int32 `json:"spacecraft_flown"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	SuccessfulLaunches *int32 `json:"successful_launches"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
	TypeVal *SpacecraftConfigTypeJSON `json:"type"`
	Url string `json:"url"`
	WikiLink string `json:"wiki_link"`
}

type SpacecraftConfigFamilyDetailedJSON struct {
	AttemptedLandings *int32 `json:"attempted_landings"`
	Description string `json:"description"`
	FailedLandings *int32 `json:"failed_landings"`
	FailedLaunches *int32 `json:"failed_launches"`
	Id int32 `json:"id"`
	MaidenFlight *string `json:"maiden_flight"`
	Manufacturer *AgencyNormalJSON `json:"manufacturer"`
	Name string `json:"name"`
	Parent *SpacecraftConfigFamilyNormalJSON `json:"parent"`
	ResponseMode string `json:"response_mode"`
	SpacecraftFlown *int32 `json:"spacecraft_flown"`
	SuccessfulLandings *int32 `json:"successful_landings"`
	SuccessfulLaunches *int32 `json:"successful_launches"`
	TotalLaunchCount *int32 `json:"total_launch_count"`
}

type SpacecraftConfigFamilyMiniJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
}

type SpacecraftConfigFamilyNormalJSON struct {
	Description string `json:"description"`
	Id int32 `json:"id"`
	MaidenFlight *string `json:"maiden_flight"`
	Manufacturer *AgencyMiniJSON `json:"manufacturer"`
	Name string `json:"name"`
	Parent *SpacecraftConfigFamilyMiniJSON `json:"parent"`
	ResponseMode string `json:"response_mode"`
}

type SpacecraftConfigNormalJSON struct {
	Agency *AgencyMiniJSON `json:"agency"`
	Family []SpacecraftConfigFamilyNormalJSON `json:"family"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InUse bool `json:"in_use"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	TypeVal *SpacecraftConfigTypeJSON `json:"type"`
	Url string `json:"url"`
}

type SpacecraftConfigTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type SpacecraftDetailedJSON struct {
	Description string `json:"description"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	FlightsCount *int32 `json:"flights_count"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InSpace bool `json:"in_space"`
	IsPlaceholder bool `json:"is_placeholder"`
	MissionEndsCount *int32 `json:"mission_ends_count"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	SerialNumber *string `json:"serial_number"`
	SpacecraftConfig *SpacecraftConfigDetailedJSON `json:"spacecraft_config"`
	Status *SpacecraftStatusJSON `json:"status"`
	TimeDocked *string `json:"time_docked"`
	TimeInSpace *string `json:"time_in_space"`
	Url string `json:"url"`
}

type SpacecraftEndpointDetailedJSON struct {
	Description string `json:"description"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	Flights []SpacecraftFlightNormalJSON `json:"flights"`
	FlightsCount *int32 `json:"flights_count"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InSpace bool `json:"in_space"`
	IsPlaceholder bool `json:"is_placeholder"`
	MissionEndsCount *int32 `json:"mission_ends_count"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	SerialNumber *string `json:"serial_number"`
	SpacecraftConfig *SpacecraftConfigDetailedJSON `json:"spacecraft_config"`
	Status *SpacecraftStatusJSON `json:"status"`
	TimeDocked *string `json:"time_docked"`
	TimeInSpace *string `json:"time_in_space"`
	Url string `json:"url"`
}

type SpacecraftFlightDetailedJSON struct {
	Destination *string `json:"destination"`
	DockingEvents []DockingEventForChaserNormalJSON `json:"docking_events"`
	Duration string `json:"duration"`
	Id int32 `json:"id"`
	Landing *LandingJSON `json:"landing"`
	LandingCrew []AstronautFlightJSON `json:"landing_crew"`
	Launch *LaunchNormalJSON `json:"launch"`
	LaunchCrew []AstronautFlightJSON `json:"launch_crew"`
	MissionEnd *string `json:"mission_end"`
	OnboardCrew []AstronautFlightJSON `json:"onboard_crew"`
	ResponseMode string `json:"response_mode"`
	Spacecraft *SpacecraftDetailedJSON `json:"spacecraft"`
	TurnAroundTime string `json:"turn_around_time"`
	Url string `json:"url"`
}

type SpacecraftFlightDetailedSerializerNoLandingJSON struct {
	Destination *string `json:"destination"`
	DockingEvents []DockingEventForChaserNormalJSON `json:"docking_events"`
	Duration string `json:"duration"`
	Id int32 `json:"id"`
	LandingCrew []AstronautFlightJSON `json:"landing_crew"`
	Launch *LaunchNormalJSON `json:"launch"`
	LaunchCrew []AstronautFlightJSON `json:"launch_crew"`
	MissionEnd *string `json:"mission_end"`
	OnboardCrew []AstronautFlightJSON `json:"onboard_crew"`
	ResponseMode string `json:"response_mode"`
	Spacecraft *SpacecraftDetailedJSON `json:"spacecraft"`
	TurnAroundTime string `json:"turn_around_time"`
	Url string `json:"url"`
}

type SpacecraftFlightDetailedSerializerNoLaunchJSON struct {
	Destination *string `json:"destination"`
	DockingEvents []DockingEventForChaserNormalJSON `json:"docking_events"`
	Duration string `json:"duration"`
	Id int32 `json:"id"`
	Landing *LandingJSON `json:"landing"`
	LandingCrew []AstronautFlightJSON `json:"landing_crew"`
	LaunchCrew []AstronautFlightJSON `json:"launch_crew"`
	MissionEnd *string `json:"mission_end"`
	OnboardCrew []AstronautFlightJSON `json:"onboard_crew"`
	ResponseMode string `json:"response_mode"`
	Spacecraft *SpacecraftDetailedJSON `json:"spacecraft"`
	TurnAroundTime string `json:"turn_around_time"`
	Url string `json:"url"`
}

type SpacecraftFlightForDockingEventJSON struct {
	Id int32 `json:"id"`
	Launch *LaunchNormalJSON `json:"launch"`
	Spacecraft *SpacecraftDetailedJSON `json:"spacecraft"`
	Url string `json:"url"`
}

type SpacecraftFlightMiniJSON struct {
	Destination *string `json:"destination"`
	Duration string `json:"duration"`
	Id int32 `json:"id"`
	Landing *LandingJSON `json:"landing"`
	Launch *LaunchMiniJSON `json:"launch"`
	MissionEnd *string `json:"mission_end"`
	Spacecraft *SpacecraftNormalJSON `json:"spacecraft"`
	TurnAroundTime string `json:"turn_around_time"`
	Url string `json:"url"`
}

type SpacecraftFlightNormalJSON struct {
	Destination *string `json:"destination"`
	Duration string `json:"duration"`
	Id int32 `json:"id"`
	Landing *LandingJSON `json:"landing"`
	Launch *LaunchNormalJSON `json:"launch"`
	MissionEnd *string `json:"mission_end"`
	ResponseMode string `json:"response_mode"`
	Spacecraft *SpacecraftNormalJSON `json:"spacecraft"`
	TurnAroundTime string `json:"turn_around_time"`
	Url string `json:"url"`
}

type SpacecraftNormalJSON struct {
	Description string `json:"description"`
	FastestTurnaround *string `json:"fastest_turnaround"`
	FlightsCount *int32 `json:"flights_count"`
	Id int32 `json:"id"`
	Image *ImageJSON `json:"image"`
	InSpace bool `json:"in_space"`
	IsPlaceholder bool `json:"is_placeholder"`
	MissionEndsCount *int32 `json:"mission_ends_count"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	SerialNumber *string `json:"serial_number"`
	SpacecraftConfig *SpacecraftConfigNormalJSON `json:"spacecraft_config"`
	Status *SpacecraftStatusJSON `json:"status"`
	TimeDocked *string `json:"time_docked"`
	TimeInSpace *string `json:"time_in_space"`
	Url string `json:"url"`
}

type SpacecraftStatusJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type SpacewalkEndpointDetailedJSON struct {
	Crew []AstronautFlightJSON `json:"crew"`
	Duration *string `json:"duration"`
	End *string `json:"end"`
	Event *EventNormalJSON `json:"event"`
	Expedition *ExpeditionNormalSerializerForSpacewalkJSON `json:"expedition"`
	Id int32 `json:"id"`
	Location *string `json:"location"`
	Name string `json:"name"`
	Program []ProgramNormalJSON `json:"program"`
	ResponseMode string `json:"response_mode"`
	SpacecraftFlight *SpacecraftFlightDetailedJSON `json:"spacecraft_flight"`
	Spacestation *SpaceStationNormalJSON `json:"spacestation"`
	Start *string `json:"start"`
	Url string `json:"url"`
}

type SpacewalkListJSON struct {
	Duration *string `json:"duration"`
	End *string `json:"end"`
	Id int32 `json:"id"`
	Location *string `json:"location"`
	Name string `json:"name"`
	ResponseMode string `json:"response_mode"`
	Start *string `json:"start"`
	Url string `json:"url"`
}

type TimelineEventJSON struct {
	RelativeTime *string `json:"relative_time"`
	TypeVal *TimelineEventTypeJSON `json:"type"`
}

type TimelineEventTypeJSON struct {
	Abbrev string `json:"abbrev"`
	Description string `json:"description"`
	Id int32 `json:"id"`
}

type UpdateJSON struct {
	Comment *string `json:"comment"`
	CreatedBy *string `json:"created_by"`
	CreatedOn string `json:"created_on"`
	Id int32 `json:"id"`
	InfoUrl *string `json:"info_url"`
	ProfileImage *string `json:"profile_image"`
}

type VidURLJSON struct {
	Description *string `json:"description"`
	EndTime *string `json:"end_time"`
	FeatureImage *string `json:"feature_image"`
	Language *LanguageJSON `json:"language"`
	Live bool `json:"live"`
	Priority int32 `json:"priority"`
	Publisher *string `json:"publisher"`
	Source *string `json:"source"`
	StartTime *string `json:"start_time"`
	Title *string `json:"title"`
	TypeVal *VidURLTypeJSON `json:"type"`
	Url string `json:"url"`
}

type VidURLTypeJSON struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type ListAgenciesResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []AgencyEndpointDetailedJSON `json:"results"`
}

type ListAstronautsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []AstronautDetailedJSON `json:"results"`
}

type ListCelestialBodiesResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []CelestialBodyEndpointDetailedJSON `json:"results"`
}

type ListDockingEventsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []DockingEventEndpointDetailedJSON `json:"results"`
}

type ListEventsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []EventEndpointDetailedJSON `json:"results"`
}

type ListExpeditionsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []ExpeditionDetailedJSON `json:"results"`
}

type ListLandingsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LandingEndpointDetailedJSON `json:"results"`
}

type ListLaunchersResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LauncherDetailedJSON `json:"results"`
}

type ListLauncherConfigurationsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LauncherConfigDetailedJSON `json:"results"`
}

type ListLaunchesResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LaunchDetailedJSON `json:"results"`
}

type ListLocationsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LocationSerializerWithPadsJSON `json:"results"`
}

type ListPadsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []PadJSON `json:"results"`
}

type ListPayloadsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []PayloadDetailedJSON `json:"results"`
}

type ListProgramsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []ProgramNormalJSON `json:"results"`
}

type ListSpaceStationsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []SpaceStationDetailedEndpointJSON `json:"results"`
}

type ListSpacecraftsResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []SpacecraftEndpointDetailedJSON `json:"results"`
}

type ListSpacewalksResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []SpacewalkEndpointDetailedJSON `json:"results"`
}

type ListUpdatesResponseJSON struct {
	Count int32 `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []UpdateJSON `json:"results"`
}

func (c *RESTClient) Do(req *http.Request) (*http.Response, error) {
	path := req.URL.Path
	reqContentType := req.Header.Get("Content-Type")

	switch path {
	case "/agency.v1.AgencyService/ListAgencies":
		var protoReq agencyv1.ListAgenciesRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/agencies/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListAgenciesResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &agencyv1.ListAgenciesResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapAgencyEndpointDetailedJSONToProto_agency(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/agency.v1.AgencyService/GetAgency":
		var protoReq agencyv1.GetAgencyRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/agencies/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp AgencyEndpointDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &agencyv1.GetAgencyResponse{
			Agency: mapAgencyEndpointDetailedJSONToProto_agency(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/astronaut.v1.AstronautService/ListAstronauts":
		var protoReq astronautv1.ListAstronautsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/astronauts/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListAstronautsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &astronautv1.ListAstronautsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapAstronautDetailedJSONToProto_astronaut(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/astronaut.v1.AstronautService/GetAstronaut":
		var protoReq astronautv1.GetAstronautRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/astronauts/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp AstronautDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &astronautv1.GetAstronautResponse{
			Astronaut: mapAstronautDetailedJSONToProto_astronaut(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/celestial_body.v1.CelestialBodyService/ListCelestialBodies":
		var protoReq celestial_bodyv1.ListCelestialBodiesRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/celestial_bodies/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListCelestialBodiesResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &celestial_bodyv1.ListCelestialBodiesResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapCelestialBodyEndpointDetailedJSONToProto_celestial_body(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/celestial_body.v1.CelestialBodyService/GetCelestialBody":
		var protoReq celestial_bodyv1.GetCelestialBodyRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/celestial_bodies/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp CelestialBodyEndpointDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &celestial_bodyv1.GetCelestialBodyResponse{
			CelestialBody: mapCelestialBodyEndpointDetailedJSONToProto_celestial_body(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/docking_event.v1.DockingEventService/ListDockingEvents":
		var protoReq docking_eventv1.ListDockingEventsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/docking_events/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListDockingEventsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &docking_eventv1.ListDockingEventsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapDockingEventEndpointDetailedJSONToProto_docking_event(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/docking_event.v1.DockingEventService/GetDockingEvent":
		var protoReq docking_eventv1.GetDockingEventRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/docking_events/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp DockingEventEndpointDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &docking_eventv1.GetDockingEventResponse{
			DockingEvent: mapDockingEventEndpointDetailedJSONToProto_docking_event(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/event.v1.EventService/ListEvents":
		var protoReq eventv1.ListEventsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/events/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListEventsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &eventv1.ListEventsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapEventEndpointDetailedJSONToProto_event(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/event.v1.EventService/GetEvent":
		var protoReq eventv1.GetEventRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/events/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp EventEndpointDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &eventv1.GetEventResponse{
			Event: mapEventEndpointDetailedJSONToProto_event(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/expedition.v1.ExpeditionService/ListExpeditions":
		var protoReq expeditionv1.ListExpeditionsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/expeditions/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListExpeditionsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &expeditionv1.ListExpeditionsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapExpeditionDetailedJSONToProto_expedition(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/expedition.v1.ExpeditionService/GetExpedition":
		var protoReq expeditionv1.GetExpeditionRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/expeditions/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ExpeditionDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &expeditionv1.GetExpeditionResponse{
			Expedition: mapExpeditionDetailedJSONToProto_expedition(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/landing.v1.LandingService/ListLandings":
		var protoReq landingv1.ListLandingsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/landings/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListLandingsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &landingv1.ListLandingsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapLandingEndpointDetailedJSONToProto_landing(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/landing.v1.LandingService/GetLanding":
		var protoReq landingv1.GetLandingRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/landings/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp LandingEndpointDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &landingv1.GetLandingResponse{
			Landing: mapLandingEndpointDetailedJSONToProto_landing(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/launch.v1.LaunchService/ListLaunches":
		var protoReq launchv1.ListLaunchesRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/launches/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListLaunchesResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &launchv1.ListLaunchesResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapLaunchDetailedJSONToProto_launch(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/launch.v1.LaunchService/GetLaunch":
		var protoReq launchv1.GetLaunchRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/launches/%v/?%s", c.baseURL, protoReq.Id, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp LaunchDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &launchv1.GetLaunchResponse{
			Launch: mapLaunchDetailedJSONToProto_launch(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/launcher.v1.LauncherService/ListLaunchers":
		var protoReq launcherv1.ListLaunchersRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/launchers/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListLaunchersResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &launcherv1.ListLaunchersResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapLauncherDetailedJSONToProto_launcher(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/launcher.v1.LauncherService/GetLauncher":
		var protoReq launcherv1.GetLauncherRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/launchers/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp LauncherDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &launcherv1.GetLauncherResponse{
			Launcher: mapLauncherDetailedJSONToProto_launcher(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/launcher_configuration.v1.LauncherConfigurationService/ListLauncherConfigurations":
		var protoReq launcher_configurationv1.ListLauncherConfigurationsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/launcher_configurations/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListLauncherConfigurationsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &launcher_configurationv1.ListLauncherConfigurationsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapLauncherConfigDetailedJSONToProto_launcher_configuration(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/launcher_configuration.v1.LauncherConfigurationService/GetLauncherConfiguration":
		var protoReq launcher_configurationv1.GetLauncherConfigurationRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/launcher_configurations/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp LauncherConfigDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &launcher_configurationv1.GetLauncherConfigurationResponse{
			LauncherConfiguration: mapLauncherConfigDetailedJSONToProto_launcher_configuration(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/location.v1.LocationService/ListLocations":
		var protoReq locationv1.ListLocationsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/locations/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListLocationsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &locationv1.ListLocationsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapLocationSerializerWithPadsJSONToProto_location(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/location.v1.LocationService/GetLocation":
		var protoReq locationv1.GetLocationRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/locations/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp LocationSerializerWithPadsJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &locationv1.GetLocationResponse{
			Location: mapLocationSerializerWithPadsJSONToProto_location(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/pad.v1.PadService/ListPads":
		var protoReq padv1.ListPadsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/pads/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListPadsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &padv1.ListPadsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapPadJSONToProto_pad(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/pad.v1.PadService/GetPad":
		var protoReq padv1.GetPadRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/pads/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp PadJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &padv1.GetPadResponse{
			Pad: mapPadJSONToProto_pad(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/payload.v1.PayloadService/ListPayloads":
		var protoReq payloadv1.ListPayloadsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/payloads/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListPayloadsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &payloadv1.ListPayloadsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapPayloadDetailedJSONToProto_payload(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/payload.v1.PayloadService/GetPayload":
		var protoReq payloadv1.GetPayloadRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/payloads/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp PayloadDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &payloadv1.GetPayloadResponse{
			Payload: mapPayloadDetailedJSONToProto_payload(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/program.v1.ProgramService/ListPrograms":
		var protoReq programv1.ListProgramsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/programs/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListProgramsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &programv1.ListProgramsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapProgramNormalJSONToProto_program(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/program.v1.ProgramService/GetProgram":
		var protoReq programv1.GetProgramRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/programs/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ProgramNormalJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &programv1.GetProgramResponse{
			Program: mapProgramNormalJSONToProto_program(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/space_station.v1.SpaceStationService/ListSpaceStations":
		var protoReq space_stationv1.ListSpaceStationsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/space_stations/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListSpaceStationsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &space_stationv1.ListSpaceStationsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapSpaceStationDetailedEndpointJSONToProto_space_station(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/space_station.v1.SpaceStationService/GetSpaceStation":
		var protoReq space_stationv1.GetSpaceStationRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/space_stations/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp SpaceStationDetailedEndpointJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &space_stationv1.GetSpaceStationResponse{
			SpaceStation: mapSpaceStationDetailedEndpointJSONToProto_space_station(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/spacecraft.v1.SpacecraftService/ListSpacecrafts":
		var protoReq spacecraftv1.ListSpacecraftsRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/spacecraft/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListSpacecraftsResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &spacecraftv1.ListSpacecraftsResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapSpacecraftEndpointDetailedJSONToProto_spacecraft(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/spacecraft.v1.SpacecraftService/GetSpacecraft":
		var protoReq spacecraftv1.GetSpacecraftRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/spacecraft/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp SpacecraftEndpointDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &spacecraftv1.GetSpacecraftResponse{
			Spacecraft: mapSpacecraftEndpointDetailedJSONToProto_spacecraft(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/spacewalk.v1.SpacewalkService/ListSpacewalks":
		var protoReq spacewalkv1.ListSpacewalksRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/spacewalks/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListSpacewalksResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &spacewalkv1.ListSpacewalksResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapSpacewalkEndpointDetailedJSONToProto_spacewalk(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/spacewalk.v1.SpacewalkService/GetSpacewalk":
		var protoReq spacewalkv1.GetSpacewalkRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/spacewalks/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp SpacewalkEndpointDetailedJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &spacewalkv1.GetSpacewalkResponse{
			Spacewalk: mapSpacewalkEndpointDetailedJSONToProto_spacewalk(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	case "/update.v1.UpdateService/ListUpdates":
		var protoReq updatev1.ListUpdatesRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Limit > 0 {
			q.Set("limit", strconv.Itoa(int(protoReq.Limit)))
		}
		if protoReq.Offset > 0 {
			q.Set("offset", strconv.Itoa(int(protoReq.Offset)))
		}
		if protoReq.Search != "" {
			q.Set("search", protoReq.Search)
		}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		restURL := fmt.Sprintf("%s/2.3.0/updates/?%s", c.baseURL, q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp ListUpdatesResponseJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &updatev1.ListUpdatesResponse{
			Count:    jsonResp.Count,
			Next:     jsonResp.Next,
			Previous: jsonResp.Previous,
		}
		for _, r := range jsonResp.Results {
			protoResp.Results = append(protoResp.Results, mapUpdateJSONToProto_update(&r))
		}

		return writeResponse(reqContentType, protoResp)

	case "/update.v1.UpdateService/GetUpdate":
		var protoReq updatev1.GetUpdateRequest
		if err := unmarshalRequest(req, &protoReq); err != nil {
			return nil, err
		}

		q := url.Values{}
		if protoReq.Mode != "" {
			q.Set("mode", protoReq.Mode)
		}

		// Map to REST URL path format with ID
		restURL := fmt.Sprintf("%s/2.3.0/updates/%v/?%s", c.baseURL, int(protoReq.Id), q.Encode())
		restResp, err := c.client.Get(restURL)
		if err != nil {
			return nil, err
		}
		defer restResp.Body.Close()

		if restResp.StatusCode != http.StatusOK {
			return makeErrorResponse(restResp)
		}

		var jsonResp UpdateJSON
		if err := json.NewDecoder(restResp.Body).Decode(&jsonResp); err != nil {
			return nil, err
		}

		protoResp := &updatev1.GetUpdateResponse{
			Update: mapUpdateJSONToProto_update(&jsonResp),
		}

		return writeResponse(reqContentType, protoResp)

	default:
		return nil, fmt.Errorf("unsupported connectrpc method path: %s", path)
	}
}

func unmarshalRequest(req *http.Request, msg proto.Message) error {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	contentType := req.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		return protojson.Unmarshal(bodyBytes, msg)
	}
	return proto.Unmarshal(bodyBytes, msg)
}

func writeResponse(reqContentType string, msg proto.Message) (*http.Response, error) {
	var body []byte
	var err error
	var contentType string

	if strings.Contains(reqContentType, "application/json") {
		body, err = protojson.Marshal(msg)
		contentType = "application/json"
	} else {
		body, err = proto.Marshal(msg)
		contentType = "application/proto"
	}

	if err != nil {
		return nil, err
	}

	resp := &http.Response{
		StatusCode:    http.StatusOK,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header),
	}
	resp.Header.Set("Content-Type", contentType)
	return resp, nil
}

func makeErrorResponse(resp *http.Response) (*http.Response, error) {
	bodyBytes, _ := io.ReadAll(resp.Body)
	return nil, fmt.Errorf("REST API returned status %d: %s", resp.StatusCode, string(bodyBytes))
}
func mapAgencyDetailedJSONToProto_agency(r *AgencyDetailedJSON) *agencyv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &agencyv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*agencyv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*agencyv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_agency(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_agency(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_agency(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_agency(r.SocialLogo),
		SocialMediaLinks: func() []*agencyv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*agencyv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_agency(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_agency(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyEndpointDetailedJSONToProto_agency(r *AgencyEndpointDetailedJSON) *agencyv1.Agency {
	if r == nil {
		return nil
	}
	l := &agencyv1.Agency{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*agencyv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*agencyv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_agency(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_agency(r.Image),
		InfoUrl: r.InfoUrl,
		LauncherList: func() []*agencyv1.LauncherConfigDetailedSerializerNoManufacturer {
			if r.LauncherList == nil {
				return nil
			}
			res := make([]*agencyv1.LauncherConfigDetailedSerializerNoManufacturer, len(r.LauncherList))
			for i, v := range r.LauncherList {
				res[i] = mapLauncherConfigDetailedSerializerNoManufacturerJSONToProto_agency(&v)
			}
			return res
		}(),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_agency(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_agency(r.SocialLogo),
		SocialMediaLinks: func() []*agencyv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*agencyv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_agency(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SpacecraftList: func() []*agencyv1.SpacecraftConfigDetailed {
			if r.SpacecraftList == nil {
				return nil
			}
			res := make([]*agencyv1.SpacecraftConfigDetailed, len(r.SpacecraftList))
			for i, v := range r.SpacecraftList {
				res[i] = mapSpacecraftConfigDetailedJSONToProto_agency(&v)
			}
			return res
		}(),
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_agency(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_agency(r *AgencyMiniJSON) *agencyv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &agencyv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_agency(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_agency(r *AgencyNormalJSON) *agencyv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &agencyv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*agencyv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*agencyv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_agency(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_agency(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_agency(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_agency(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_agency(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_agency(r *AgencyTypeJSON) *agencyv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &agencyv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_agency(r *CountryJSON) *agencyv1.Country {
	if r == nil {
		return nil
	}
	l := &agencyv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapImageJSONToProto_agency(r *ImageJSON) *agencyv1.Image {
	if r == nil {
		return nil
	}
	l := &agencyv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_agency(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*agencyv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*agencyv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_agency(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_agency(r *ImageLicenseJSON) *agencyv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &agencyv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_agency(r *ImageVariantJSON) *agencyv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &agencyv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_agency(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_agency(r *ImageVariantTypeJSON) *agencyv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &agencyv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigDetailedSerializerNoManufacturerJSONToProto_agency(r *LauncherConfigDetailedSerializerNoManufacturerJSON) *agencyv1.LauncherConfigDetailedSerializerNoManufacturer {
	if r == nil {
		return nil
	}
	l := &agencyv1.LauncherConfigDetailedSerializerNoManufacturer{
		Active: r.Active,
		Alias: r.Alias,
		Apogee: r.Apogee,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Families: func() []*agencyv1.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]*agencyv1.LauncherConfigFamilyDetailed, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyDetailedJSONToProto_agency(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FullName: r.FullName,
		GeoCapacity: r.GeoCapacity,
		GtoCapacity: r.GtoCapacity,
		Id: r.Id,
		Image: mapImageJSONToProto_agency(r.Image),
		InfoUrl: r.InfoUrl,
		IsPlaceholder: r.IsPlaceholder,
		LaunchCost: r.LaunchCost,
		LaunchMass: r.LaunchMass,
		Length: r.Length,
		LeoCapacity: r.LeoCapacity,
		MaidenFlight: r.MaidenFlight,
		MaxStage: r.MaxStage,
		MinStage: r.MinStage,
		Name: r.Name,
		PendingLaunches: r.PendingLaunches,
		Program: func() []*agencyv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*agencyv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_agency(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Reusable: r.Reusable,
		SsoCapacity: r.SsoCapacity,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		ToThrust: r.ToThrust,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		Variant: r.Variant,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapLauncherConfigFamilyDetailedJSONToProto_agency(r *LauncherConfigFamilyDetailedJSON) *agencyv1.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &agencyv1.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []*agencyv1.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*agencyv1.AgencyDetailed, len(r.Manufacturer))
			for i, v := range r.Manufacturer {
				res[i] = mapAgencyDetailedJSONToProto_agency(&v)
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapLauncherConfigFamilyNormalJSONToProto_agency(r.Parent),
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_agency(r *LauncherConfigFamilyMiniJSON) *agencyv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &agencyv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigFamilyNormalJSONToProto_agency(r *LauncherConfigFamilyNormalJSON) *agencyv1.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &agencyv1.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []*agencyv1.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*agencyv1.AgencyNormal, len(r.Manufacturer))
			for i, v := range r.Manufacturer {
				res[i] = mapAgencyNormalJSONToProto_agency(&v)
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapLauncherConfigFamilyMiniJSONToProto_agency(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapMissionPatchJSONToProto_agency(r *MissionPatchJSON) *agencyv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &agencyv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_agency(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapProgramNormalJSONToProto_agency(r *ProgramNormalJSON) *agencyv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &agencyv1.ProgramNormal{
		Agencies: func() []*agencyv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*agencyv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_agency(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_agency(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*agencyv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*agencyv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_agency(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_agency(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_agency(r *ProgramTypeJSON) *agencyv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &agencyv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSocialMediaJSONToProto_agency(r *SocialMediaJSON) *agencyv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &agencyv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_agency(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_agency(r *SocialMediaLinkJSON) *agencyv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &agencyv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_agency(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftConfigDetailedJSONToProto_agency(r *SpacecraftConfigDetailedJSON) *agencyv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	l := &agencyv1.SpacecraftConfigDetailed{
		Agency: mapAgencyNormalJSONToProto_agency(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*agencyv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*agencyv1.SpacecraftConfigFamilyDetailed, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyDetailedJSONToProto_agency(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FlightLife: r.FlightLife,
		Height: r.Height,
		History: r.History,
		HumanRated: r.HumanRated,
		Id: r.Id,
		Image: mapImageJSONToProto_agency(r.Image),
		InUse: r.InUse,
		InfoLink: r.InfoLink,
		MaidenFlight: r.MaidenFlight,
		Name: r.Name,
		PayloadCapacity: r.PayloadCapacity,
		PayloadReturnCapacity: r.PayloadReturnCapacity,
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapSpacecraftConfigTypeJSONToProto_agency(r.TypeVal),
		Url: r.Url,
		WikiLink: r.WikiLink,
	}
	return l
}

func mapSpacecraftConfigFamilyDetailedJSONToProto_agency(r *SpacecraftConfigFamilyDetailedJSON) *agencyv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &agencyv1.SpacecraftConfigFamilyDetailed{
		AttemptedLandings: r.AttemptedLandings,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyNormalJSONToProto_agency(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyNormalJSONToProto_agency(r.Parent),
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapSpacecraftConfigFamilyMiniJSONToProto_agency(r *SpacecraftConfigFamilyMiniJSON) *agencyv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &agencyv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigFamilyNormalJSONToProto_agency(r *SpacecraftConfigFamilyNormalJSON) *agencyv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &agencyv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyMiniJSONToProto_agency(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyMiniJSONToProto_agency(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigTypeJSONToProto_agency(r *SpacecraftConfigTypeJSON) *agencyv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	l := &agencyv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAgencyMiniJSONToProto_astronaut(r *AgencyMiniJSON) *astronautv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &astronautv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_astronaut(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_astronaut(r *AgencyTypeJSON) *astronautv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &astronautv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautDetailedJSONToProto_astronaut(r *AstronautDetailedJSON) *astronautv1.Astronaut {
	if r == nil {
		return nil
	}
	l := &astronautv1.Astronaut{
		Age: r.Age,
		Agency: mapAgencyMiniJSONToProto_astronaut(r.Agency),
		Bio: r.Bio,
		DateOfBirth: r.DateOfBirth,
		DateOfDeath: r.DateOfDeath,
		EvaTime: r.EvaTime,
		FirstFlight: r.FirstFlight,
		Id: r.Id,
		Image: mapImageJSONToProto_astronaut(r.Image),
		InSpace: r.InSpace,
		LastFlight: r.LastFlight,
		Name: r.Name,
		Nationality: func() []*astronautv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*astronautv1.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = mapCountryJSONToProto_astronaut(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*astronautv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*astronautv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_astronaut(&v)
			}
			return res
		}(),
		Status: mapAstronautStatusJSONToProto_astronaut(r.Status),
		TimeInSpace: r.TimeInSpace,
		Type: mapAstronautTypeJSONToProto_astronaut(r.TypeVal),
		Url: r.Url,
		Wiki: r.Wiki,
	}
	return l
}

func mapAstronautStatusJSONToProto_astronaut(r *AstronautStatusJSON) *astronautv1.AstronautStatus {
	if r == nil {
		return nil
	}
	l := &astronautv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautTypeJSONToProto_astronaut(r *AstronautTypeJSON) *astronautv1.AstronautType {
	if r == nil {
		return nil
	}
	l := &astronautv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_astronaut(r *CountryJSON) *astronautv1.Country {
	if r == nil {
		return nil
	}
	l := &astronautv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapImageJSONToProto_astronaut(r *ImageJSON) *astronautv1.Image {
	if r == nil {
		return nil
	}
	l := &astronautv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_astronaut(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*astronautv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*astronautv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_astronaut(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_astronaut(r *ImageLicenseJSON) *astronautv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &astronautv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_astronaut(r *ImageVariantJSON) *astronautv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &astronautv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_astronaut(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_astronaut(r *ImageVariantTypeJSON) *astronautv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &astronautv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSocialMediaJSONToProto_astronaut(r *SocialMediaJSON) *astronautv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &astronautv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_astronaut(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_astronaut(r *SocialMediaLinkJSON) *astronautv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &astronautv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_astronaut(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapCelestialBodyEndpointDetailedJSONToProto_celestial_body(r *CelestialBodyEndpointDetailedJSON) *celestial_bodyv1.CelestialBody {
	if r == nil {
		return nil
	}
	l := &celestial_bodyv1.CelestialBody{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_celestial_body(r.Image),
		LengthOfDay: r.LengthOfDay,
		Locations: func() []*celestial_bodyv1.LocationSerializerNoCelestialBody {
			if r.Locations == nil {
				return nil
			}
			res := make([]*celestial_bodyv1.LocationSerializerNoCelestialBody, len(r.Locations))
			for i, v := range r.Locations {
				res[i] = mapLocationSerializerNoCelestialBodyJSONToProto_celestial_body(&v)
			}
			return res
		}(),
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_celestial_body(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_celestial_body(r *CelestialBodyTypeJSON) *celestial_bodyv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &celestial_bodyv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_celestial_body(r *CountryJSON) *celestial_bodyv1.Country {
	if r == nil {
		return nil
	}
	l := &celestial_bodyv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapImageJSONToProto_celestial_body(r *ImageJSON) *celestial_bodyv1.Image {
	if r == nil {
		return nil
	}
	l := &celestial_bodyv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_celestial_body(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*celestial_bodyv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*celestial_bodyv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_celestial_body(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_celestial_body(r *ImageLicenseJSON) *celestial_bodyv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &celestial_bodyv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_celestial_body(r *ImageVariantJSON) *celestial_bodyv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &celestial_bodyv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_celestial_body(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_celestial_body(r *ImageVariantTypeJSON) *celestial_bodyv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &celestial_bodyv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLocationSerializerNoCelestialBodyJSONToProto_celestial_body(r *LocationSerializerNoCelestialBodyJSON) *celestial_bodyv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	l := &celestial_bodyv1.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapCountryJSONToProto_celestial_body(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_celestial_body(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapAgencyDetailedJSONToProto_docking_event(r *AgencyDetailedJSON) *docking_eventv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*docking_eventv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*docking_eventv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_docking_event(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_docking_event(r.SocialLogo),
		SocialMediaLinks: func() []*docking_eventv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*docking_eventv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_docking_event(r *AgencyMiniJSON) *docking_eventv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_docking_event(r *AgencyNormalJSON) *docking_eventv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*docking_eventv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*docking_eventv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_docking_event(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_docking_event(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_docking_event(r *AgencyTypeJSON) *docking_eventv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCelestialBodyDetailedJSONToProto_docking_event(r *CelestialBodyDetailedJSON) *docking_eventv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_docking_event(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyMiniJSONToProto_docking_event(r *CelestialBodyMiniJSON) *docking_eventv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapCelestialBodyNormalJSONToProto_docking_event(r *CelestialBodyNormalJSON) *docking_eventv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.CelestialBodyNormal{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapCelestialBodyTypeJSONToProto_docking_event(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_docking_event(r *CelestialBodyTypeJSON) *docking_eventv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_docking_event(r *CountryJSON) *docking_eventv1.Country {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapDockingEventEndpointDetailedJSONToProto_docking_event(r *DockingEventEndpointDetailedJSON) *docking_eventv1.DockingEvent {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.DockingEvent{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapDockingLocationJSONToProto_docking_event(r.DockingLocation),
		FlightVehicleChaser: mapSpacecraftFlightNormalJSONToProto_docking_event(r.FlightVehicleChaser),
		FlightVehicleTarget: mapSpacecraftFlightMiniJSONToProto_docking_event(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightChaser: mapPayloadFlightNormalJSONToProto_docking_event(r.PayloadFlightChaser),
		PayloadFlightTarget: mapPayloadFlightMiniJSONToProto_docking_event(r.PayloadFlightTarget),
		ResponseMode: r.ResponseMode,
		SpaceStationChaser: mapSpaceStationNormalJSONToProto_docking_event(r.SpaceStationChaser),
		SpaceStationTarget: mapSpaceStationMiniJSONToProto_docking_event(r.SpaceStationTarget),
		Url: r.Url,
	}
	return l
}

func mapDockingLocationJSONToProto_docking_event(r *DockingLocationJSON) *docking_eventv1.DockingLocation {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapPayloadMiniJSONToProto_docking_event(r.Payload),
		Spacecraft: mapSpacecraftConfigNormalJSONToProto_docking_event(r.Spacecraft),
		Spacestation: mapSpaceStationMiniJSONToProto_docking_event(r.Spacestation),
	}
	return l
}

func mapImageJSONToProto_docking_event(r *ImageJSON) *docking_eventv1.Image {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_docking_event(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*docking_eventv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*docking_eventv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_docking_event(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_docking_event(r *ImageLicenseJSON) *docking_eventv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_docking_event(r *ImageVariantJSON) *docking_eventv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_docking_event(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_docking_event(r *ImageVariantTypeJSON) *docking_eventv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapInfoURLJSONToProto_docking_event(r *InfoURLJSON) *docking_eventv1.InfoURL {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_docking_event(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		Type: mapInfoURLTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapInfoURLTypeJSONToProto_docking_event(r *InfoURLTypeJSON) *docking_eventv1.InfoURLType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLandingJSONToProto_docking_event(r *LandingJSON) *docking_eventv1.Landing {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Id: r.Id,
		LandingLocation: mapLandingLocationJSONToProto_docking_event(r.LandingLocation),
		Success: r.Success,
		Type: mapLandingTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapLandingLocationJSONToProto_docking_event(r *LandingLocationJSON) *docking_eventv1.LandingLocation {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.LandingLocation{
		Abbrev: r.Abbrev,
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		CelestialBody: mapCelestialBodyNormalJSONToProto_docking_event(r.CelestialBody),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		Latitude: r.Latitude,
		Location: mapLocationSerializerNoCelestialBodyJSONToProto_docking_event(r.Location),
		Longitude: r.Longitude,
		Name: r.Name,
		SuccessfulLandings: r.SuccessfulLandings,
	}
	return l
}

func mapLandingTypeJSONToProto_docking_event(r *LandingTypeJSON) *docking_eventv1.LandingType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLanguageJSONToProto_docking_event(r *LanguageJSON) *docking_eventv1.Language {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLaunchMiniJSONToProto_docking_event(r *LaunchMiniJSON) *docking_eventv1.LaunchMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.LaunchMini{
		Id: r.Id,
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapLaunchNormalJSONToProto_docking_event(r *LaunchNormalJSON) *docking_eventv1.LaunchNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.LaunchNormal{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapAgencyMiniJSONToProto_docking_event(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapMissionJSONToProto_docking_event(r.Mission),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapNetPrecisionJSONToProto_docking_event(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapPadJSONToProto_docking_event(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		Probability: r.Probability,
		Program: func() []*docking_eventv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*docking_eventv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_docking_event(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapRocketNormalJSONToProto_docking_event(r.Rocket),
		Slug: r.Slug,
		Status: mapLaunchStatusJSONToProto_docking_event(r.Status),
		Url: r.Url,
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
	return l
}

func mapLaunchStatusJSONToProto_docking_event(r *LaunchStatusJSON) *docking_eventv1.LaunchStatus {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_docking_event(r *LauncherConfigFamilyMiniJSON) *docking_eventv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigListJSONToProto_docking_event(r *LauncherConfigListJSON) *docking_eventv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.LauncherConfigList{
		Families: func() []*docking_eventv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*docking_eventv1.LauncherConfigFamilyMini, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyMiniJSONToProto_docking_event(&v)
			}
			return res
		}(),
		FullName: r.FullName,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		Variant: r.Variant,
	}
	return l
}

func mapLocationJSONToProto_docking_event(r *LocationJSON) *docking_eventv1.Location {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.Location{
		Active: r.Active,
		CelestialBody: mapCelestialBodyDetailedJSONToProto_docking_event(r.CelestialBody),
		Country: mapCountryJSONToProto_docking_event(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapLocationSerializerNoCelestialBodyJSONToProto_docking_event(r *LocationSerializerNoCelestialBodyJSON) *docking_eventv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapCountryJSONToProto_docking_event(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapMissionJSONToProto_docking_event(r *MissionJSON) *docking_eventv1.Mission {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.Mission{
		Agencies: func() []*docking_eventv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*docking_eventv1.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyDetailedJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		InfoUrls: func() []*docking_eventv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*docking_eventv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapOrbitJSONToProto_docking_event(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*docking_eventv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*docking_eventv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_docking_event(&v)
			}
			return res
		}(),
	}
	return l
}

func mapMissionPatchJSONToProto_docking_event(r *MissionPatchJSON) *docking_eventv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_docking_event(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapNetPrecisionJSONToProto_docking_event(r *NetPrecisionJSON) *docking_eventv1.NetPrecision {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapOrbitJSONToProto_docking_event(r *OrbitJSON) *docking_eventv1.Orbit {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapCelestialBodyMiniJSONToProto_docking_event(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapPadJSONToProto_docking_event(r *PadJSON) *docking_eventv1.Pad {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.Pad{
		Active: r.Active,
		Agencies: func() []*docking_eventv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*docking_eventv1.AgencyNormal, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyNormalJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Country: mapCountryJSONToProto_docking_event(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Location: mapLocationJSONToProto_docking_event(r.Location),
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapPayloadFlightMiniJSONToProto_docking_event(r *PayloadFlightMiniJSON) *docking_eventv1.PayloadFlightMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.PayloadFlightMini{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapLandingJSONToProto_docking_event(r.Landing),
		Launch: mapLaunchMiniJSONToProto_docking_event(r.Launch),
		Payload: mapPayloadMiniJSONToProto_docking_event(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
	return l
}

func mapPayloadFlightNormalJSONToProto_docking_event(r *PayloadFlightNormalJSON) *docking_eventv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.PayloadFlightNormal{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapLandingJSONToProto_docking_event(r.Landing),
		Launch: mapLaunchNormalJSONToProto_docking_event(r.Launch),
		Payload: mapPayloadNormalJSONToProto_docking_event(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
	return l
}

func mapPayloadMiniJSONToProto_docking_event(r *PayloadMiniJSON) *docking_eventv1.PayloadMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.PayloadMini{
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		Manufacturer: mapAgencyMiniJSONToProto_docking_event(r.Manufacturer),
		Name: r.Name,
		Operator: mapAgencyMiniJSONToProto_docking_event(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_docking_event(r.TypeVal),
	}
	return l
}

func mapPayloadNormalJSONToProto_docking_event(r *PayloadNormalJSON) *docking_eventv1.PayloadNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapAgencyNormalJSONToProto_docking_event(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapAgencyNormalJSONToProto_docking_event(r.Operator),
		Program: func() []*docking_eventv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*docking_eventv1.ProgramMini, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramMiniJSONToProto_docking_event(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_docking_event(r.TypeVal),
		WikiLink: r.WikiLink,
	}
	return l
}

func mapPayloadTypeJSONToProto_docking_event(r *PayloadTypeJSON) *docking_eventv1.PayloadType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapProgramMiniJSONToProto_docking_event(r *ProgramMiniJSON) *docking_eventv1.ProgramMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.ProgramMini{
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramNormalJSONToProto_docking_event(r *ProgramNormalJSON) *docking_eventv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.ProgramNormal{
		Agencies: func() []*docking_eventv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*docking_eventv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*docking_eventv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*docking_eventv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_docking_event(r *ProgramTypeJSON) *docking_eventv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapRocketNormalJSONToProto_docking_event(r *RocketNormalJSON) *docking_eventv1.RocketNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.RocketNormal{
		Configuration: mapLauncherConfigListJSONToProto_docking_event(r.Configuration),
		Id: r.Id,
	}
	return l
}

func mapSocialMediaJSONToProto_docking_event(r *SocialMediaJSON) *docking_eventv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_docking_event(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_docking_event(r *SocialMediaLinkJSON) *docking_eventv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_docking_event(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationMiniJSONToProto_docking_event(r *SpaceStationMiniJSON) *docking_eventv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpaceStationMini{
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSpaceStationNormalJSONToProto_docking_event(r *SpaceStationNormalJSON) *docking_eventv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpaceStationNormal{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Status: mapSpaceStationStatusJSONToProto_docking_event(r.Status),
		Type: mapSpaceStationTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationStatusJSONToProto_docking_event(r *SpaceStationStatusJSON) *docking_eventv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpaceStationTypeJSONToProto_docking_event(r *SpaceStationTypeJSON) *docking_eventv1.SpaceStationType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftConfigFamilyMiniJSONToProto_docking_event(r *SpacecraftConfigFamilyMiniJSON) *docking_eventv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigFamilyNormalJSONToProto_docking_event(r *SpacecraftConfigFamilyNormalJSON) *docking_eventv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyMiniJSONToProto_docking_event(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyMiniJSONToProto_docking_event(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigNormalJSONToProto_docking_event(r *SpacecraftConfigNormalJSON) *docking_eventv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpacecraftConfigNormal{
		Agency: mapAgencyMiniJSONToProto_docking_event(r.Agency),
		Family: func() []*docking_eventv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*docking_eventv1.SpacecraftConfigFamilyNormal, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyNormalJSONToProto_docking_event(&v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		InUse: r.InUse,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapSpacecraftConfigTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftConfigTypeJSONToProto_docking_event(r *SpacecraftConfigTypeJSON) *docking_eventv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftFlightMiniJSONToProto_docking_event(r *SpacecraftFlightMiniJSON) *docking_eventv1.SpacecraftFlightMini {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpacecraftFlightMini{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_docking_event(r.Landing),
		Launch: mapLaunchMiniJSONToProto_docking_event(r.Launch),
		MissionEnd: r.MissionEnd,
		Spacecraft: mapSpacecraftNormalJSONToProto_docking_event(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightNormalJSONToProto_docking_event(r *SpacecraftFlightNormalJSON) *docking_eventv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_docking_event(r.Landing),
		Launch: mapLaunchNormalJSONToProto_docking_event(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftNormalJSONToProto_docking_event(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftNormalJSONToProto_docking_event(r *SpacecraftNormalJSON) *docking_eventv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpacecraftNormal{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_docking_event(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigNormalJSONToProto_docking_event(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_docking_event(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftStatusJSONToProto_docking_event(r *SpacecraftStatusJSON) *docking_eventv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapVidURLJSONToProto_docking_event(r *VidURLJSON) *docking_eventv1.VidURL {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_docking_event(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		Type: mapVidURLTypeJSONToProto_docking_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapVidURLTypeJSONToProto_docking_event(r *VidURLTypeJSON) *docking_eventv1.VidURLType {
	if r == nil {
		return nil
	}
	l := &docking_eventv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAgencyMiniJSONToProto_event(r *AgencyMiniJSON) *eventv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &eventv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_event(r *AgencyTypeJSON) *eventv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &eventv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautNormalJSONToProto_event(r *AstronautNormalJSON) *eventv1.AstronautNormal {
	if r == nil {
		return nil
	}
	l := &eventv1.AstronautNormal{
		Agency: mapAgencyMiniJSONToProto_event(r.Agency),
		Id: r.Id,
		Image: mapImageJSONToProto_event(r.Image),
		Name: r.Name,
		Status: mapAstronautStatusJSONToProto_event(r.Status),
		Url: r.Url,
	}
	return l
}

func mapAstronautStatusJSONToProto_event(r *AstronautStatusJSON) *eventv1.AstronautStatus {
	if r == nil {
		return nil
	}
	l := &eventv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapEventEndpointDetailedJSONToProto_event(r *EventEndpointDetailedJSON) *eventv1.Event {
	if r == nil {
		return nil
	}
	l := &eventv1.Event{
		Agencies: func() []*eventv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*eventv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_event(&v)
			}
			return res
		}(),
		Astronauts: func() []*eventv1.AstronautNormal {
			if r.Astronauts == nil {
				return nil
			}
			res := make([]*eventv1.AstronautNormal, len(r.Astronauts))
			for i, v := range r.Astronauts {
				res[i] = mapAstronautNormalJSONToProto_event(&v)
			}
			return res
		}(),
		Date: r.Date,
		DatePrecision: mapNetPrecisionJSONToProto_event(r.DatePrecision),
		Description: r.Description,
		Duration: r.Duration,
		Expeditions: func() []*eventv1.ExpeditionNormal {
			if r.Expeditions == nil {
				return nil
			}
			res := make([]*eventv1.ExpeditionNormal, len(r.Expeditions))
			for i, v := range r.Expeditions {
				res[i] = mapExpeditionNormalJSONToProto_event(&v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapImageJSONToProto_event(r.Image),
		InfoUrls: func() []*eventv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*eventv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_event(&v)
			}
			return res
		}(),
		LastUpdated: r.LastUpdated,
		Launches: func() []*eventv1.LaunchBasic {
			if r.Launches == nil {
				return nil
			}
			res := make([]*eventv1.LaunchBasic, len(r.Launches))
			for i, v := range r.Launches {
				res[i] = mapLaunchBasicJSONToProto_event(&v)
			}
			return res
		}(),
		Location: r.Location,
		Name: r.Name,
		Program: func() []*eventv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*eventv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_event(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Slug: r.Slug,
		Spacestations: func() []*eventv1.SpaceStationNormal {
			if r.Spacestations == nil {
				return nil
			}
			res := make([]*eventv1.SpaceStationNormal, len(r.Spacestations))
			for i, v := range r.Spacestations {
				res[i] = mapSpaceStationNormalJSONToProto_event(&v)
			}
			return res
		}(),
		Type: mapEventTypeJSONToProto_event(r.TypeVal),
		Updates: func() []*eventv1.Update {
			if r.Updates == nil {
				return nil
			}
			res := make([]*eventv1.Update, len(r.Updates))
			for i, v := range r.Updates {
				res[i] = mapUpdateJSONToProto_event(&v)
			}
			return res
		}(),
		Url: r.Url,
		VidUrls: func() []*eventv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*eventv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_event(&v)
			}
			return res
		}(),
		WebcastLive: r.WebcastLive,
	}
	return l
}

func mapEventTypeJSONToProto_event(r *EventTypeJSON) *eventv1.EventType {
	if r == nil {
		return nil
	}
	l := &eventv1.EventType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapExpeditionNormalJSONToProto_event(r *ExpeditionNormalJSON) *eventv1.ExpeditionNormal {
	if r == nil {
		return nil
	}
	l := &eventv1.ExpeditionNormal{
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []*eventv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*eventv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_event(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Spacestation: mapSpaceStationNormalJSONToProto_event(r.Spacestation),
		Spacewalks: func() []*eventv1.SpacewalkList {
			if r.Spacewalks == nil {
				return nil
			}
			res := make([]*eventv1.SpacewalkList, len(r.Spacewalks))
			for i, v := range r.Spacewalks {
				res[i] = mapSpacewalkListJSONToProto_event(&v)
			}
			return res
		}(),
		Start: r.Start,
		Url: r.Url,
	}
	return l
}

func mapImageJSONToProto_event(r *ImageJSON) *eventv1.Image {
	if r == nil {
		return nil
	}
	l := &eventv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_event(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*eventv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*eventv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_event(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_event(r *ImageLicenseJSON) *eventv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &eventv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_event(r *ImageVariantJSON) *eventv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &eventv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_event(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_event(r *ImageVariantTypeJSON) *eventv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &eventv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapInfoURLJSONToProto_event(r *InfoURLJSON) *eventv1.InfoURL {
	if r == nil {
		return nil
	}
	l := &eventv1.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_event(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		Type: mapInfoURLTypeJSONToProto_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapInfoURLTypeJSONToProto_event(r *InfoURLTypeJSON) *eventv1.InfoURLType {
	if r == nil {
		return nil
	}
	l := &eventv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLanguageJSONToProto_event(r *LanguageJSON) *eventv1.Language {
	if r == nil {
		return nil
	}
	l := &eventv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLaunchBasicJSONToProto_event(r *LaunchBasicJSON) *eventv1.LaunchBasic {
	if r == nil {
		return nil
	}
	l := &eventv1.LaunchBasic{
		Id: r.Id,
		Image: mapImageJSONToProto_event(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapNetPrecisionJSONToProto_event(r.NetPrecision),
		ResponseMode: r.ResponseMode,
		Slug: r.Slug,
		Status: mapLaunchStatusJSONToProto_event(r.Status),
		Url: r.Url,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
	return l
}

func mapLaunchStatusJSONToProto_event(r *LaunchStatusJSON) *eventv1.LaunchStatus {
	if r == nil {
		return nil
	}
	l := &eventv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapMissionPatchJSONToProto_event(r *MissionPatchJSON) *eventv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &eventv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_event(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapNetPrecisionJSONToProto_event(r *NetPrecisionJSON) *eventv1.NetPrecision {
	if r == nil {
		return nil
	}
	l := &eventv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapProgramNormalJSONToProto_event(r *ProgramNormalJSON) *eventv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &eventv1.ProgramNormal{
		Agencies: func() []*eventv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*eventv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_event(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_event(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*eventv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*eventv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_event(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_event(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_event(r *ProgramTypeJSON) *eventv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &eventv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpaceStationNormalJSONToProto_event(r *SpaceStationNormalJSON) *eventv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	l := &eventv1.SpaceStationNormal{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapImageJSONToProto_event(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Status: mapSpaceStationStatusJSONToProto_event(r.Status),
		Type: mapSpaceStationTypeJSONToProto_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationStatusJSONToProto_event(r *SpaceStationStatusJSON) *eventv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	l := &eventv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpaceStationTypeJSONToProto_event(r *SpaceStationTypeJSON) *eventv1.SpaceStationType {
	if r == nil {
		return nil
	}
	l := &eventv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacewalkListJSONToProto_event(r *SpacewalkListJSON) *eventv1.SpacewalkList {
	if r == nil {
		return nil
	}
	l := &eventv1.SpacewalkList{
		Duration: r.Duration,
		End: r.End,
		Id: r.Id,
		Location: r.Location,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Start: r.Start,
		Url: r.Url,
	}
	return l
}

func mapUpdateJSONToProto_event(r *UpdateJSON) *eventv1.Update {
	if r == nil {
		return nil
	}
	l := &eventv1.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
	return l
}

func mapVidURLJSONToProto_event(r *VidURLJSON) *eventv1.VidURL {
	if r == nil {
		return nil
	}
	l := &eventv1.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_event(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		Type: mapVidURLTypeJSONToProto_event(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapVidURLTypeJSONToProto_event(r *VidURLTypeJSON) *eventv1.VidURLType {
	if r == nil {
		return nil
	}
	l := &eventv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAgencyMiniJSONToProto_expedition(r *AgencyMiniJSON) *expeditionv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &expeditionv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_expedition(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_expedition(r *AgencyNormalJSON) *expeditionv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &expeditionv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*expeditionv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*expeditionv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_expedition(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_expedition(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_expedition(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_expedition(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_expedition(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_expedition(r *AgencyTypeJSON) *expeditionv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &expeditionv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautDetailedJSONToProto_expedition(r *AstronautDetailedJSON) *expeditionv1.AstronautDetailed {
	if r == nil {
		return nil
	}
	l := &expeditionv1.AstronautDetailed{
		Age: r.Age,
		Agency: mapAgencyMiniJSONToProto_expedition(r.Agency),
		Bio: r.Bio,
		DateOfBirth: r.DateOfBirth,
		DateOfDeath: r.DateOfDeath,
		EvaTime: r.EvaTime,
		FirstFlight: r.FirstFlight,
		Id: r.Id,
		Image: mapImageJSONToProto_expedition(r.Image),
		InSpace: r.InSpace,
		LastFlight: r.LastFlight,
		Name: r.Name,
		Nationality: func() []*expeditionv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*expeditionv1.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = mapCountryJSONToProto_expedition(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*expeditionv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*expeditionv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_expedition(&v)
			}
			return res
		}(),
		Status: mapAstronautStatusJSONToProto_expedition(r.Status),
		TimeInSpace: r.TimeInSpace,
		Type: mapAstronautTypeJSONToProto_expedition(r.TypeVal),
		Url: r.Url,
		Wiki: r.Wiki,
	}
	return l
}

func mapAstronautFlightJSONToProto_expedition(r *AstronautFlightJSON) *expeditionv1.AstronautFlight {
	if r == nil {
		return nil
	}
	l := &expeditionv1.AstronautFlight{
		Astronaut: mapAstronautDetailedJSONToProto_expedition(r.Astronaut),
		Id: r.Id,
		Role: mapAstronautRoleJSONToProto_expedition(r.Role),
	}
	return l
}

func mapAstronautRoleJSONToProto_expedition(r *AstronautRoleJSON) *expeditionv1.AstronautRole {
	if r == nil {
		return nil
	}
	l := &expeditionv1.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
	return l
}

func mapAstronautStatusJSONToProto_expedition(r *AstronautStatusJSON) *expeditionv1.AstronautStatus {
	if r == nil {
		return nil
	}
	l := &expeditionv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautTypeJSONToProto_expedition(r *AstronautTypeJSON) *expeditionv1.AstronautType {
	if r == nil {
		return nil
	}
	l := &expeditionv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_expedition(r *CountryJSON) *expeditionv1.Country {
	if r == nil {
		return nil
	}
	l := &expeditionv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapExpeditionDetailedJSONToProto_expedition(r *ExpeditionDetailedJSON) *expeditionv1.Expedition {
	if r == nil {
		return nil
	}
	l := &expeditionv1.Expedition{
		Crew: func() []*expeditionv1.AstronautFlight {
			if r.Crew == nil {
				return nil
			}
			res := make([]*expeditionv1.AstronautFlight, len(r.Crew))
			for i, v := range r.Crew {
				res[i] = mapAstronautFlightJSONToProto_expedition(&v)
			}
			return res
		}(),
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []*expeditionv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*expeditionv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_expedition(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Spacestation: mapSpaceStationDetailedJSONToProto_expedition(r.Spacestation),
		Spacewalks: func() []*expeditionv1.SpacewalkList {
			if r.Spacewalks == nil {
				return nil
			}
			res := make([]*expeditionv1.SpacewalkList, len(r.Spacewalks))
			for i, v := range r.Spacewalks {
				res[i] = mapSpacewalkListJSONToProto_expedition(&v)
			}
			return res
		}(),
		Start: r.Start,
		Url: r.Url,
	}
	return l
}

func mapImageJSONToProto_expedition(r *ImageJSON) *expeditionv1.Image {
	if r == nil {
		return nil
	}
	l := &expeditionv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_expedition(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*expeditionv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*expeditionv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_expedition(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_expedition(r *ImageLicenseJSON) *expeditionv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &expeditionv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_expedition(r *ImageVariantJSON) *expeditionv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &expeditionv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_expedition(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_expedition(r *ImageVariantTypeJSON) *expeditionv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &expeditionv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapMissionPatchJSONToProto_expedition(r *MissionPatchJSON) *expeditionv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &expeditionv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_expedition(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSocialMediaJSONToProto_expedition(r *SocialMediaJSON) *expeditionv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &expeditionv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_expedition(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_expedition(r *SocialMediaLinkJSON) *expeditionv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &expeditionv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_expedition(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationDetailedJSONToProto_expedition(r *SpaceStationDetailedJSON) *expeditionv1.SpaceStationDetailed {
	if r == nil {
		return nil
	}
	l := &expeditionv1.SpaceStationDetailed{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapImageJSONToProto_expedition(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Owners: func() []*expeditionv1.AgencyNormal {
			if r.Owners == nil {
				return nil
			}
			res := make([]*expeditionv1.AgencyNormal, len(r.Owners))
			for i, v := range r.Owners {
				res[i] = mapAgencyNormalJSONToProto_expedition(&v)
			}
			return res
		}(),
		Status: mapSpaceStationStatusJSONToProto_expedition(r.Status),
		Type: mapSpaceStationTypeJSONToProto_expedition(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationStatusJSONToProto_expedition(r *SpaceStationStatusJSON) *expeditionv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	l := &expeditionv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpaceStationTypeJSONToProto_expedition(r *SpaceStationTypeJSON) *expeditionv1.SpaceStationType {
	if r == nil {
		return nil
	}
	l := &expeditionv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacewalkListJSONToProto_expedition(r *SpacewalkListJSON) *expeditionv1.SpacewalkList {
	if r == nil {
		return nil
	}
	l := &expeditionv1.SpacewalkList{
		Duration: r.Duration,
		End: r.End,
		Id: r.Id,
		Location: r.Location,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Start: r.Start,
		Url: r.Url,
	}
	return l
}

func mapAgencyDetailedJSONToProto_landing(r *AgencyDetailedJSON) *landingv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &landingv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*landingv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*landingv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_landing(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_landing(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_landing(r.SocialLogo),
		SocialMediaLinks: func() []*landingv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*landingv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_landing(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_landing(r *AgencyMiniJSON) *landingv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &landingv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_landing(r *AgencyNormalJSON) *landingv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*landingv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*landingv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_landing(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_landing(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_landing(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_landing(r *AgencyTypeJSON) *landingv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &landingv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautDetailedJSONToProto_landing(r *AstronautDetailedJSON) *landingv1.AstronautDetailed {
	if r == nil {
		return nil
	}
	l := &landingv1.AstronautDetailed{
		Age: r.Age,
		Agency: mapAgencyMiniJSONToProto_landing(r.Agency),
		Bio: r.Bio,
		DateOfBirth: r.DateOfBirth,
		DateOfDeath: r.DateOfDeath,
		EvaTime: r.EvaTime,
		FirstFlight: r.FirstFlight,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InSpace: r.InSpace,
		LastFlight: r.LastFlight,
		Name: r.Name,
		Nationality: func() []*landingv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*landingv1.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = mapCountryJSONToProto_landing(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*landingv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*landingv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_landing(&v)
			}
			return res
		}(),
		Status: mapAstronautStatusJSONToProto_landing(r.Status),
		TimeInSpace: r.TimeInSpace,
		Type: mapAstronautTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
		Wiki: r.Wiki,
	}
	return l
}

func mapAstronautFlightJSONToProto_landing(r *AstronautFlightJSON) *landingv1.AstronautFlight {
	if r == nil {
		return nil
	}
	l := &landingv1.AstronautFlight{
		Astronaut: mapAstronautDetailedJSONToProto_landing(r.Astronaut),
		Id: r.Id,
		Role: mapAstronautRoleJSONToProto_landing(r.Role),
	}
	return l
}

func mapAstronautRoleJSONToProto_landing(r *AstronautRoleJSON) *landingv1.AstronautRole {
	if r == nil {
		return nil
	}
	l := &landingv1.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
	return l
}

func mapAstronautStatusJSONToProto_landing(r *AstronautStatusJSON) *landingv1.AstronautStatus {
	if r == nil {
		return nil
	}
	l := &landingv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautTypeJSONToProto_landing(r *AstronautTypeJSON) *landingv1.AstronautType {
	if r == nil {
		return nil
	}
	l := &landingv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCelestialBodyDetailedJSONToProto_landing(r *CelestialBodyDetailedJSON) *landingv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	l := &landingv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_landing(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyMiniJSONToProto_landing(r *CelestialBodyMiniJSON) *landingv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	l := &landingv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapCelestialBodyNormalJSONToProto_landing(r *CelestialBodyNormalJSON) *landingv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.CelestialBodyNormal{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapCelestialBodyTypeJSONToProto_landing(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_landing(r *CelestialBodyTypeJSON) *landingv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &landingv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_landing(r *CountryJSON) *landingv1.Country {
	if r == nil {
		return nil
	}
	l := &landingv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapDockingEventForChaserNormalJSONToProto_landing(r *DockingEventForChaserNormalJSON) *landingv1.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.DockingEventForChaserNormal{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapDockingLocationJSONToProto_landing(r.DockingLocation),
		FlightVehicleTarget: mapSpacecraftFlightNormalJSONToProto_landing(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightTarget: mapPayloadFlightNormalJSONToProto_landing(r.PayloadFlightTarget),
		SpaceStationTarget: mapSpaceStationNormalJSONToProto_landing(r.SpaceStationTarget),
		Url: r.Url,
	}
	return l
}

func mapDockingLocationJSONToProto_landing(r *DockingLocationJSON) *landingv1.DockingLocation {
	if r == nil {
		return nil
	}
	l := &landingv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapPayloadMiniJSONToProto_landing(r.Payload),
		Spacecraft: mapSpacecraftConfigNormalJSONToProto_landing(r.Spacecraft),
		Spacestation: mapSpaceStationMiniJSONToProto_landing(r.Spacestation),
	}
	return l
}

func mapFirstStageDetailedSerializerNoLandingJSONToProto_landing(r *FirstStageDetailedSerializerNoLandingJSON) *landingv1.FirstStageDetailedSerializerNoLanding {
	if r == nil {
		return nil
	}
	l := &landingv1.FirstStageDetailedSerializerNoLanding{
		Id: r.Id,
		Launcher: mapLauncherNormalJSONToProto_landing(r.Launcher),
		LauncherFlightNumber: r.LauncherFlightNumber,
		PreviousFlight: mapLaunchNormalJSONToProto_landing(r.PreviousFlight),
		PreviousFlightDate: r.PreviousFlightDate,
		Reused: r.Reused,
		TurnAroundTime: r.TurnAroundTime,
		Type: r.TypeVal,
	}
	return l
}

func mapImageJSONToProto_landing(r *ImageJSON) *landingv1.Image {
	if r == nil {
		return nil
	}
	l := &landingv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_landing(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*landingv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*landingv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_landing(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_landing(r *ImageLicenseJSON) *landingv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &landingv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_landing(r *ImageVariantJSON) *landingv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &landingv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_landing(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_landing(r *ImageVariantTypeJSON) *landingv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &landingv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapInfoURLJSONToProto_landing(r *InfoURLJSON) *landingv1.InfoURL {
	if r == nil {
		return nil
	}
	l := &landingv1.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_landing(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		Type: mapInfoURLTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapInfoURLTypeJSONToProto_landing(r *InfoURLTypeJSON) *landingv1.InfoURLType {
	if r == nil {
		return nil
	}
	l := &landingv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLandingJSONToProto_landing(r *LandingJSON) *landingv1.LandingRecord {
	if r == nil {
		return nil
	}
	l := &landingv1.LandingRecord{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Id: r.Id,
		LandingLocation: mapLandingLocationJSONToProto_landing(r.LandingLocation),
		Success: r.Success,
		Type: mapLandingTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapLandingEndpointDetailedJSONToProto_landing(r *LandingEndpointDetailedJSON) *landingv1.Landing {
	if r == nil {
		return nil
	}
	l := &landingv1.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Firststage: mapFirstStageDetailedSerializerNoLandingJSONToProto_landing(r.Firststage),
		Id: r.Id,
		LandingLocation: mapLandingLocationJSONToProto_landing(r.LandingLocation),
		Payloadflight: mapPayloadFlightDetailedSerializerNoLandingJSONToProto_landing(r.Payloadflight),
		ResponseMode: r.ResponseMode,
		Spacecraftflight: mapSpacecraftFlightDetailedSerializerNoLandingJSONToProto_landing(r.Spacecraftflight),
		Success: r.Success,
		Type: mapLandingTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapLandingLocationJSONToProto_landing(r *LandingLocationJSON) *landingv1.LandingLocation {
	if r == nil {
		return nil
	}
	l := &landingv1.LandingLocation{
		Abbrev: r.Abbrev,
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		CelestialBody: mapCelestialBodyNormalJSONToProto_landing(r.CelestialBody),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		Latitude: r.Latitude,
		Location: mapLocationSerializerNoCelestialBodyJSONToProto_landing(r.Location),
		Longitude: r.Longitude,
		Name: r.Name,
		SuccessfulLandings: r.SuccessfulLandings,
	}
	return l
}

func mapLandingTypeJSONToProto_landing(r *LandingTypeJSON) *landingv1.LandingType {
	if r == nil {
		return nil
	}
	l := &landingv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLanguageJSONToProto_landing(r *LanguageJSON) *landingv1.Language {
	if r == nil {
		return nil
	}
	l := &landingv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLaunchNormalJSONToProto_landing(r *LaunchNormalJSON) *landingv1.LaunchNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.LaunchNormal{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapAgencyMiniJSONToProto_landing(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapMissionJSONToProto_landing(r.Mission),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapNetPrecisionJSONToProto_landing(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapPadJSONToProto_landing(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		Probability: r.Probability,
		Program: func() []*landingv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*landingv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_landing(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapRocketNormalJSONToProto_landing(r.Rocket),
		Slug: r.Slug,
		Status: mapLaunchStatusJSONToProto_landing(r.Status),
		Url: r.Url,
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
	return l
}

func mapLaunchStatusJSONToProto_landing(r *LaunchStatusJSON) *landingv1.LaunchStatus {
	if r == nil {
		return nil
	}
	l := &landingv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_landing(r *LauncherConfigFamilyMiniJSON) *landingv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &landingv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigListJSONToProto_landing(r *LauncherConfigListJSON) *landingv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	l := &landingv1.LauncherConfigList{
		Families: func() []*landingv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*landingv1.LauncherConfigFamilyMini, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyMiniJSONToProto_landing(&v)
			}
			return res
		}(),
		FullName: r.FullName,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		Variant: r.Variant,
	}
	return l
}

func mapLauncherNormalJSONToProto_landing(r *LauncherNormalJSON) *landingv1.LauncherNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.LauncherNormal{
		AttemptedLandings: r.AttemptedLandings,
		Details: r.Details,
		FastestTurnaround: r.FastestTurnaround,
		FirstLaunchDate: r.FirstLaunchDate,
		FlightProven: r.FlightProven,
		Flights: r.Flights,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		IsPlaceholder: r.IsPlaceholder,
		LastLaunchDate: r.LastLaunchDate,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		Status: mapLauncherStatusJSONToProto_landing(r.Status),
		SuccessfulLandings: r.SuccessfulLandings,
		Url: r.Url,
	}
	return l
}

func mapLauncherStatusJSONToProto_landing(r *LauncherStatusJSON) *landingv1.LauncherStatus {
	if r == nil {
		return nil
	}
	l := &landingv1.LauncherStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLocationJSONToProto_landing(r *LocationJSON) *landingv1.Location {
	if r == nil {
		return nil
	}
	l := &landingv1.Location{
		Active: r.Active,
		CelestialBody: mapCelestialBodyDetailedJSONToProto_landing(r.CelestialBody),
		Country: mapCountryJSONToProto_landing(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapLocationSerializerNoCelestialBodyJSONToProto_landing(r *LocationSerializerNoCelestialBodyJSON) *landingv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	l := &landingv1.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapCountryJSONToProto_landing(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapMissionJSONToProto_landing(r *MissionJSON) *landingv1.Mission {
	if r == nil {
		return nil
	}
	l := &landingv1.Mission{
		Agencies: func() []*landingv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*landingv1.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyDetailedJSONToProto_landing(&v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InfoUrls: func() []*landingv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*landingv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_landing(&v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapOrbitJSONToProto_landing(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*landingv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*landingv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_landing(&v)
			}
			return res
		}(),
	}
	return l
}

func mapMissionPatchJSONToProto_landing(r *MissionPatchJSON) *landingv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &landingv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_landing(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapNetPrecisionJSONToProto_landing(r *NetPrecisionJSON) *landingv1.NetPrecision {
	if r == nil {
		return nil
	}
	l := &landingv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapOrbitJSONToProto_landing(r *OrbitJSON) *landingv1.Orbit {
	if r == nil {
		return nil
	}
	l := &landingv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapCelestialBodyMiniJSONToProto_landing(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapPadJSONToProto_landing(r *PadJSON) *landingv1.Pad {
	if r == nil {
		return nil
	}
	l := &landingv1.Pad{
		Active: r.Active,
		Agencies: func() []*landingv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*landingv1.AgencyNormal, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyNormalJSONToProto_landing(&v)
			}
			return res
		}(),
		Country: mapCountryJSONToProto_landing(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Location: mapLocationJSONToProto_landing(r.Location),
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapPayloadDetailedJSONToProto_landing(r *PayloadDetailedJSON) *landingv1.PayloadDetailed {
	if r == nil {
		return nil
	}
	l := &landingv1.PayloadDetailed{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapAgencyDetailedJSONToProto_landing(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapAgencyDetailedJSONToProto_landing(r.Operator),
		Program: func() []*landingv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*landingv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_landing(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_landing(r.TypeVal),
		WikiLink: r.WikiLink,
	}
	return l
}

func mapPayloadFlightDetailedSerializerNoLandingJSONToProto_landing(r *PayloadFlightDetailedSerializerNoLandingJSON) *landingv1.PayloadFlightDetailedSerializerNoLanding {
	if r == nil {
		return nil
	}
	l := &landingv1.PayloadFlightDetailedSerializerNoLanding{
		Amount: r.Amount,
		Destination: r.Destination,
		DockingEvents: func() []*landingv1.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]*landingv1.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = mapDockingEventForChaserNormalJSONToProto_landing(&v)
			}
			return res
		}(),
		Id: r.Id,
		Launch: mapLaunchNormalJSONToProto_landing(r.Launch),
		Payload: mapPayloadDetailedJSONToProto_landing(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
	return l
}

func mapPayloadFlightNormalJSONToProto_landing(r *PayloadFlightNormalJSON) *landingv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.PayloadFlightNormal{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapLandingJSONToProto_landing(r.Landing),
		Launch: mapLaunchNormalJSONToProto_landing(r.Launch),
		Payload: mapPayloadNormalJSONToProto_landing(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
	return l
}

func mapPayloadMiniJSONToProto_landing(r *PayloadMiniJSON) *landingv1.PayloadMini {
	if r == nil {
		return nil
	}
	l := &landingv1.PayloadMini{
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		Manufacturer: mapAgencyMiniJSONToProto_landing(r.Manufacturer),
		Name: r.Name,
		Operator: mapAgencyMiniJSONToProto_landing(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_landing(r.TypeVal),
	}
	return l
}

func mapPayloadNormalJSONToProto_landing(r *PayloadNormalJSON) *landingv1.PayloadNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapAgencyNormalJSONToProto_landing(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapAgencyNormalJSONToProto_landing(r.Operator),
		Program: func() []*landingv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*landingv1.ProgramMini, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramMiniJSONToProto_landing(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_landing(r.TypeVal),
		WikiLink: r.WikiLink,
	}
	return l
}

func mapPayloadTypeJSONToProto_landing(r *PayloadTypeJSON) *landingv1.PayloadType {
	if r == nil {
		return nil
	}
	l := &landingv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapProgramMiniJSONToProto_landing(r *ProgramMiniJSON) *landingv1.ProgramMini {
	if r == nil {
		return nil
	}
	l := &landingv1.ProgramMini{
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramNormalJSONToProto_landing(r *ProgramNormalJSON) *landingv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.ProgramNormal{
		Agencies: func() []*landingv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*landingv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_landing(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*landingv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*landingv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_landing(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_landing(r *ProgramTypeJSON) *landingv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &landingv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapRocketNormalJSONToProto_landing(r *RocketNormalJSON) *landingv1.RocketNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.RocketNormal{
		Configuration: mapLauncherConfigListJSONToProto_landing(r.Configuration),
		Id: r.Id,
	}
	return l
}

func mapSocialMediaJSONToProto_landing(r *SocialMediaJSON) *landingv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &landingv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_landing(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_landing(r *SocialMediaLinkJSON) *landingv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &landingv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_landing(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationMiniJSONToProto_landing(r *SpaceStationMiniJSON) *landingv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	l := &landingv1.SpaceStationMini{
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSpaceStationNormalJSONToProto_landing(r *SpaceStationNormalJSON) *landingv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.SpaceStationNormal{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Status: mapSpaceStationStatusJSONToProto_landing(r.Status),
		Type: mapSpaceStationTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationStatusJSONToProto_landing(r *SpaceStationStatusJSON) *landingv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	l := &landingv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpaceStationTypeJSONToProto_landing(r *SpaceStationTypeJSON) *landingv1.SpaceStationType {
	if r == nil {
		return nil
	}
	l := &landingv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftConfigDetailedJSONToProto_landing(r *SpacecraftConfigDetailedJSON) *landingv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftConfigDetailed{
		Agency: mapAgencyNormalJSONToProto_landing(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*landingv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*landingv1.SpacecraftConfigFamilyDetailed, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyDetailedJSONToProto_landing(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FlightLife: r.FlightLife,
		Height: r.Height,
		History: r.History,
		HumanRated: r.HumanRated,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InUse: r.InUse,
		InfoLink: r.InfoLink,
		MaidenFlight: r.MaidenFlight,
		Name: r.Name,
		PayloadCapacity: r.PayloadCapacity,
		PayloadReturnCapacity: r.PayloadReturnCapacity,
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapSpacecraftConfigTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
		WikiLink: r.WikiLink,
	}
	return l
}

func mapSpacecraftConfigFamilyDetailedJSONToProto_landing(r *SpacecraftConfigFamilyDetailedJSON) *landingv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftConfigFamilyDetailed{
		AttemptedLandings: r.AttemptedLandings,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyNormalJSONToProto_landing(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyNormalJSONToProto_landing(r.Parent),
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapSpacecraftConfigFamilyMiniJSONToProto_landing(r *SpacecraftConfigFamilyMiniJSON) *landingv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigFamilyNormalJSONToProto_landing(r *SpacecraftConfigFamilyNormalJSON) *landingv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyMiniJSONToProto_landing(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyMiniJSONToProto_landing(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigNormalJSONToProto_landing(r *SpacecraftConfigNormalJSON) *landingv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftConfigNormal{
		Agency: mapAgencyMiniJSONToProto_landing(r.Agency),
		Family: func() []*landingv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*landingv1.SpacecraftConfigFamilyNormal, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyNormalJSONToProto_landing(&v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InUse: r.InUse,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapSpacecraftConfigTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftConfigTypeJSONToProto_landing(r *SpacecraftConfigTypeJSON) *landingv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftDetailedJSONToProto_landing(r *SpacecraftDetailedJSON) *landingv1.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftDetailed{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigDetailedJSONToProto_landing(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_landing(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightDetailedSerializerNoLandingJSONToProto_landing(r *SpacecraftFlightDetailedSerializerNoLandingJSON) *landingv1.SpacecraftFlightDetailedSerializerNoLanding {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftFlightDetailedSerializerNoLanding{
		Destination: r.Destination,
		DockingEvents: func() []*landingv1.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]*landingv1.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = mapDockingEventForChaserNormalJSONToProto_landing(&v)
			}
			return res
		}(),
		Duration: r.Duration,
		Id: r.Id,
		LandingCrew: func() []*landingv1.AstronautFlight {
			if r.LandingCrew == nil {
				return nil
			}
			res := make([]*landingv1.AstronautFlight, len(r.LandingCrew))
			for i, v := range r.LandingCrew {
				res[i] = mapAstronautFlightJSONToProto_landing(&v)
			}
			return res
		}(),
		Launch: mapLaunchNormalJSONToProto_landing(r.Launch),
		LaunchCrew: func() []*landingv1.AstronautFlight {
			if r.LaunchCrew == nil {
				return nil
			}
			res := make([]*landingv1.AstronautFlight, len(r.LaunchCrew))
			for i, v := range r.LaunchCrew {
				res[i] = mapAstronautFlightJSONToProto_landing(&v)
			}
			return res
		}(),
		MissionEnd: r.MissionEnd,
		OnboardCrew: func() []*landingv1.AstronautFlight {
			if r.OnboardCrew == nil {
				return nil
			}
			res := make([]*landingv1.AstronautFlight, len(r.OnboardCrew))
			for i, v := range r.OnboardCrew {
				res[i] = mapAstronautFlightJSONToProto_landing(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftDetailedJSONToProto_landing(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightNormalJSONToProto_landing(r *SpacecraftFlightNormalJSON) *landingv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_landing(r.Landing),
		Launch: mapLaunchNormalJSONToProto_landing(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftNormalJSONToProto_landing(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftNormalJSONToProto_landing(r *SpacecraftNormalJSON) *landingv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftNormal{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_landing(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigNormalJSONToProto_landing(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_landing(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftStatusJSONToProto_landing(r *SpacecraftStatusJSON) *landingv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	l := &landingv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapVidURLJSONToProto_landing(r *VidURLJSON) *landingv1.VidURL {
	if r == nil {
		return nil
	}
	l := &landingv1.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_landing(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		Type: mapVidURLTypeJSONToProto_landing(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapVidURLTypeJSONToProto_landing(r *VidURLTypeJSON) *landingv1.VidURLType {
	if r == nil {
		return nil
	}
	l := &landingv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapImageJSONToProto_launcher(r *ImageJSON) *launcherv1.Image {
	if r == nil {
		return nil
	}
	l := &launcherv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_launcher(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*launcherv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*launcherv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_launcher(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_launcher(r *ImageLicenseJSON) *launcherv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &launcherv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_launcher(r *ImageVariantJSON) *launcherv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &launcherv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_launcher(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_launcher(r *ImageVariantTypeJSON) *launcherv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &launcherv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_launcher(r *LauncherConfigFamilyMiniJSON) *launcherv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &launcherv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigListJSONToProto_launcher(r *LauncherConfigListJSON) *launcherv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	l := &launcherv1.LauncherConfigList{
		Families: func() []*launcherv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*launcherv1.LauncherConfigFamilyMini, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyMiniJSONToProto_launcher(&v)
			}
			return res
		}(),
		FullName: r.FullName,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		Variant: r.Variant,
	}
	return l
}

func mapLauncherDetailedJSONToProto_launcher(r *LauncherDetailedJSON) *launcherv1.Launcher {
	if r == nil {
		return nil
	}
	l := &launcherv1.Launcher{
		AttemptedLandings: r.AttemptedLandings,
		Details: r.Details,
		FastestTurnaround: r.FastestTurnaround,
		FirstLaunchDate: r.FirstLaunchDate,
		FlightProven: r.FlightProven,
		Flights: r.Flights,
		Id: r.Id,
		Image: mapImageJSONToProto_launcher(r.Image),
		IsPlaceholder: r.IsPlaceholder,
		LastLaunchDate: r.LastLaunchDate,
		LauncherConfig: mapLauncherConfigListJSONToProto_launcher(r.LauncherConfig),
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		Status: mapLauncherStatusJSONToProto_launcher(r.Status),
		SuccessfulLandings: r.SuccessfulLandings,
		Url: r.Url,
	}
	return l
}

func mapLauncherStatusJSONToProto_launcher(r *LauncherStatusJSON) *launcherv1.LauncherStatus {
	if r == nil {
		return nil
	}
	l := &launcherv1.LauncherStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAgencyDetailedJSONToProto_launcher_configuration(r *AgencyDetailedJSON) *launcher_configurationv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*launcher_configurationv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_launcher_configuration(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_launcher_configuration(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_launcher_configuration(r.SocialLogo),
		SocialMediaLinks: func() []*launcher_configurationv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_launcher_configuration(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_launcher_configuration(r *AgencyMiniJSON) *launcher_configurationv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_launcher_configuration(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_launcher_configuration(r *AgencyNormalJSON) *launcher_configurationv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*launcher_configurationv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_launcher_configuration(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_launcher_configuration(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_launcher_configuration(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_launcher_configuration(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_launcher_configuration(r *AgencyTypeJSON) *launcher_configurationv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_launcher_configuration(r *CountryJSON) *launcher_configurationv1.Country {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapImageJSONToProto_launcher_configuration(r *ImageJSON) *launcher_configurationv1.Image {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_launcher_configuration(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*launcher_configurationv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_launcher_configuration(r *ImageLicenseJSON) *launcher_configurationv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_launcher_configuration(r *ImageVariantJSON) *launcher_configurationv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_launcher_configuration(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_launcher_configuration(r *ImageVariantTypeJSON) *launcher_configurationv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigDetailedJSONToProto_launcher_configuration(r *LauncherConfigDetailedJSON) *launcher_configurationv1.LauncherConfiguration {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.LauncherConfiguration{
		Active: r.Active,
		Alias: r.Alias,
		Apogee: r.Apogee,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Families: func() []*launcher_configurationv1.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.LauncherConfigFamilyDetailed, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyDetailedJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FullName: r.FullName,
		GeoCapacity: r.GeoCapacity,
		GtoCapacity: r.GtoCapacity,
		Id: r.Id,
		Image: mapImageJSONToProto_launcher_configuration(r.Image),
		InfoUrl: r.InfoUrl,
		IsPlaceholder: r.IsPlaceholder,
		LaunchCost: r.LaunchCost,
		LaunchMass: r.LaunchMass,
		Length: r.Length,
		LeoCapacity: r.LeoCapacity,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyDetailedJSONToProto_launcher_configuration(r.Manufacturer),
		MaxStage: r.MaxStage,
		MinStage: r.MinStage,
		Name: r.Name,
		PendingLaunches: r.PendingLaunches,
		Program: func() []*launcher_configurationv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Reusable: r.Reusable,
		SsoCapacity: r.SsoCapacity,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		ToThrust: r.ToThrust,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		Variant: r.Variant,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapLauncherConfigFamilyDetailedJSONToProto_launcher_configuration(r *LauncherConfigFamilyDetailedJSON) *launcher_configurationv1.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []*launcher_configurationv1.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.AgencyDetailed, len(r.Manufacturer))
			for i, v := range r.Manufacturer {
				res[i] = mapAgencyDetailedJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapLauncherConfigFamilyNormalJSONToProto_launcher_configuration(r.Parent),
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_launcher_configuration(r *LauncherConfigFamilyMiniJSON) *launcher_configurationv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigFamilyNormalJSONToProto_launcher_configuration(r *LauncherConfigFamilyNormalJSON) *launcher_configurationv1.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []*launcher_configurationv1.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.AgencyNormal, len(r.Manufacturer))
			for i, v := range r.Manufacturer {
				res[i] = mapAgencyNormalJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapLauncherConfigFamilyMiniJSONToProto_launcher_configuration(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapMissionPatchJSONToProto_launcher_configuration(r *MissionPatchJSON) *launcher_configurationv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_launcher_configuration(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapProgramNormalJSONToProto_launcher_configuration(r *ProgramNormalJSON) *launcher_configurationv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.ProgramNormal{
		Agencies: func() []*launcher_configurationv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_launcher_configuration(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*launcher_configurationv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_launcher_configuration(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_launcher_configuration(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_launcher_configuration(r *ProgramTypeJSON) *launcher_configurationv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSocialMediaJSONToProto_launcher_configuration(r *SocialMediaJSON) *launcher_configurationv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_launcher_configuration(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_launcher_configuration(r *SocialMediaLinkJSON) *launcher_configurationv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &launcher_configurationv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_launcher_configuration(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapAgencyDetailedJSONToProto_launch(r *AgencyDetailedJSON) *launchv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*launchv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*launchv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_launch(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_launch(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_launch(r.SocialLogo),
		SocialMediaLinks: func() []*launchv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*launchv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_launch(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_launch(r *AgencyMiniJSON) *launchv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &launchv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_launch(r *AgencyNormalJSON) *launchv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*launchv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*launchv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_launch(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_launch(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_launch(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_launch(r *AgencyTypeJSON) *launchv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &launchv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautDetailedJSONToProto_launch(r *AstronautDetailedJSON) *launchv1.AstronautDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.AstronautDetailed{
		Age: r.Age,
		Agency: mapAgencyMiniJSONToProto_launch(r.Agency),
		Bio: r.Bio,
		DateOfBirth: r.DateOfBirth,
		DateOfDeath: r.DateOfDeath,
		EvaTime: r.EvaTime,
		FirstFlight: r.FirstFlight,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InSpace: r.InSpace,
		LastFlight: r.LastFlight,
		Name: r.Name,
		Nationality: func() []*launchv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*launchv1.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = mapCountryJSONToProto_launch(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*launchv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*launchv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_launch(&v)
			}
			return res
		}(),
		Status: mapAstronautStatusJSONToProto_launch(r.Status),
		TimeInSpace: r.TimeInSpace,
		Type: mapAstronautTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
		Wiki: r.Wiki,
	}
	return l
}

func mapAstronautFlightJSONToProto_launch(r *AstronautFlightJSON) *launchv1.AstronautFlight {
	if r == nil {
		return nil
	}
	l := &launchv1.AstronautFlight{
		Astronaut: mapAstronautDetailedJSONToProto_launch(r.Astronaut),
		Id: r.Id,
		Role: mapAstronautRoleJSONToProto_launch(r.Role),
	}
	return l
}

func mapAstronautRoleJSONToProto_launch(r *AstronautRoleJSON) *launchv1.AstronautRole {
	if r == nil {
		return nil
	}
	l := &launchv1.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
	return l
}

func mapAstronautStatusJSONToProto_launch(r *AstronautStatusJSON) *launchv1.AstronautStatus {
	if r == nil {
		return nil
	}
	l := &launchv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautTypeJSONToProto_launch(r *AstronautTypeJSON) *launchv1.AstronautType {
	if r == nil {
		return nil
	}
	l := &launchv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCelestialBodyDetailedJSONToProto_launch(r *CelestialBodyDetailedJSON) *launchv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_launch(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyMiniJSONToProto_launch(r *CelestialBodyMiniJSON) *launchv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	l := &launchv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapCelestialBodyNormalJSONToProto_launch(r *CelestialBodyNormalJSON) *launchv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.CelestialBodyNormal{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapCelestialBodyTypeJSONToProto_launch(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_launch(r *CelestialBodyTypeJSON) *launchv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &launchv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_launch(r *CountryJSON) *launchv1.Country {
	if r == nil {
		return nil
	}
	l := &launchv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapDockingEventForChaserNormalJSONToProto_launch(r *DockingEventForChaserNormalJSON) *launchv1.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.DockingEventForChaserNormal{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapDockingLocationJSONToProto_launch(r.DockingLocation),
		FlightVehicleTarget: mapSpacecraftFlightNormalJSONToProto_launch(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightTarget: mapPayloadFlightNormalJSONToProto_launch(r.PayloadFlightTarget),
		SpaceStationTarget: mapSpaceStationNormalJSONToProto_launch(r.SpaceStationTarget),
		Url: r.Url,
	}
	return l
}

func mapDockingLocationJSONToProto_launch(r *DockingLocationJSON) *launchv1.DockingLocation {
	if r == nil {
		return nil
	}
	l := &launchv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapPayloadMiniJSONToProto_launch(r.Payload),
		Spacecraft: mapSpacecraftConfigNormalJSONToProto_launch(r.Spacecraft),
		Spacestation: mapSpaceStationMiniJSONToProto_launch(r.Spacestation),
	}
	return l
}

func mapFirstStageNormalJSONToProto_launch(r *FirstStageNormalJSON) *launchv1.FirstStageNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.FirstStageNormal{
		Id: r.Id,
		Landing: mapLandingJSONToProto_launch(r.Landing),
		Launcher: mapLauncherNormalJSONToProto_launch(r.Launcher),
		LauncherFlightNumber: r.LauncherFlightNumber,
		PreviousFlight: mapLaunchMiniJSONToProto_launch(r.PreviousFlight),
		PreviousFlightDate: r.PreviousFlightDate,
		Reused: r.Reused,
		TurnAroundTime: r.TurnAroundTime,
		Type: r.TypeVal,
	}
	return l
}

func mapImageJSONToProto_launch(r *ImageJSON) *launchv1.Image {
	if r == nil {
		return nil
	}
	l := &launchv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_launch(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*launchv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*launchv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_launch(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_launch(r *ImageLicenseJSON) *launchv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &launchv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_launch(r *ImageVariantJSON) *launchv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &launchv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_launch(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_launch(r *ImageVariantTypeJSON) *launchv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &launchv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapInfoURLJSONToProto_launch(r *InfoURLJSON) *launchv1.InfoURL {
	if r == nil {
		return nil
	}
	l := &launchv1.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_launch(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		Type: mapInfoURLTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapInfoURLTypeJSONToProto_launch(r *InfoURLTypeJSON) *launchv1.InfoURLType {
	if r == nil {
		return nil
	}
	l := &launchv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLandingJSONToProto_launch(r *LandingJSON) *launchv1.Landing {
	if r == nil {
		return nil
	}
	l := &launchv1.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Id: r.Id,
		LandingLocation: mapLandingLocationJSONToProto_launch(r.LandingLocation),
		Success: r.Success,
		Type: mapLandingTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapLandingLocationJSONToProto_launch(r *LandingLocationJSON) *launchv1.LandingLocation {
	if r == nil {
		return nil
	}
	l := &launchv1.LandingLocation{
		Abbrev: r.Abbrev,
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		CelestialBody: mapCelestialBodyNormalJSONToProto_launch(r.CelestialBody),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		Latitude: r.Latitude,
		Location: mapLocationSerializerNoCelestialBodyJSONToProto_launch(r.Location),
		Longitude: r.Longitude,
		Name: r.Name,
		SuccessfulLandings: r.SuccessfulLandings,
	}
	return l
}

func mapLandingTypeJSONToProto_launch(r *LandingTypeJSON) *launchv1.LandingType {
	if r == nil {
		return nil
	}
	l := &launchv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLanguageJSONToProto_launch(r *LanguageJSON) *launchv1.Language {
	if r == nil {
		return nil
	}
	l := &launchv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLaunchDetailedJSONToProto_launch(r *LaunchDetailedJSON) *launchv1.Launch {
	if r == nil {
		return nil
	}
	l := &launchv1.Launch{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		FlightclubUrl: r.FlightclubUrl,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoUrls: func() []*launchv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*launchv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_launch(&v)
			}
			return res
		}(),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapAgencyDetailedJSONToProto_launch(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapMissionJSONToProto_launch(r.Mission),
		MissionPatches: func() []*launchv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*launchv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_launch(&v)
			}
			return res
		}(),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapNetPrecisionJSONToProto_launch(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapPadJSONToProto_launch(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		PadTurnaround: r.PadTurnaround,
		Probability: r.Probability,
		Program: func() []*launchv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapRocketDetailedJSONToProto_launch(r.Rocket),
		Slug: r.Slug,
		Status: mapLaunchStatusJSONToProto_launch(r.Status),
		Timeline: func() []*launchv1.TimelineEvent {
			if r.Timeline == nil {
				return nil
			}
			res := make([]*launchv1.TimelineEvent, len(r.Timeline))
			for i, v := range r.Timeline {
				res[i] = mapTimelineEventJSONToProto_launch(&v)
			}
			return res
		}(),
		Updates: func() []*launchv1.Update {
			if r.Updates == nil {
				return nil
			}
			res := make([]*launchv1.Update, len(r.Updates))
			for i, v := range r.Updates {
				res[i] = mapUpdateJSONToProto_launch(&v)
			}
			return res
		}(),
		Url: r.Url,
		VidUrls: func() []*launchv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*launchv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_launch(&v)
			}
			return res
		}(),
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
	return l
}

func mapLaunchMiniJSONToProto_launch(r *LaunchMiniJSON) *launchv1.LaunchMini {
	if r == nil {
		return nil
	}
	l := &launchv1.LaunchMini{
		Id: r.Id,
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapLaunchNormalJSONToProto_launch(r *LaunchNormalJSON) *launchv1.LaunchNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.LaunchNormal{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapAgencyMiniJSONToProto_launch(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapMissionJSONToProto_launch(r.Mission),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapNetPrecisionJSONToProto_launch(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapPadJSONToProto_launch(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		Probability: r.Probability,
		Program: func() []*launchv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapRocketNormalJSONToProto_launch(r.Rocket),
		Slug: r.Slug,
		Status: mapLaunchStatusJSONToProto_launch(r.Status),
		Url: r.Url,
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
	return l
}

func mapLaunchStatusJSONToProto_launch(r *LaunchStatusJSON) *launchv1.LaunchStatus {
	if r == nil {
		return nil
	}
	l := &launchv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigDetailedJSONToProto_launch(r *LauncherConfigDetailedJSON) *launchv1.LauncherConfigDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.LauncherConfigDetailed{
		Active: r.Active,
		Alias: r.Alias,
		Apogee: r.Apogee,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Families: func() []*launchv1.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]*launchv1.LauncherConfigFamilyDetailed, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyDetailedJSONToProto_launch(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FullName: r.FullName,
		GeoCapacity: r.GeoCapacity,
		GtoCapacity: r.GtoCapacity,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoUrl: r.InfoUrl,
		IsPlaceholder: r.IsPlaceholder,
		LaunchCost: r.LaunchCost,
		LaunchMass: r.LaunchMass,
		Length: r.Length,
		LeoCapacity: r.LeoCapacity,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyDetailedJSONToProto_launch(r.Manufacturer),
		MaxStage: r.MaxStage,
		MinStage: r.MinStage,
		Name: r.Name,
		PendingLaunches: r.PendingLaunches,
		Program: func() []*launchv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Reusable: r.Reusable,
		SsoCapacity: r.SsoCapacity,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		ToThrust: r.ToThrust,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		Variant: r.Variant,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapLauncherConfigFamilyDetailedJSONToProto_launch(r *LauncherConfigFamilyDetailedJSON) *launchv1.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []*launchv1.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*launchv1.AgencyDetailed, len(r.Manufacturer))
			for i, v := range r.Manufacturer {
				res[i] = mapAgencyDetailedJSONToProto_launch(&v)
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapLauncherConfigFamilyNormalJSONToProto_launch(r.Parent),
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_launch(r *LauncherConfigFamilyMiniJSON) *launchv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &launchv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigFamilyNormalJSONToProto_launch(r *LauncherConfigFamilyNormalJSON) *launchv1.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []*launchv1.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*launchv1.AgencyNormal, len(r.Manufacturer))
			for i, v := range r.Manufacturer {
				res[i] = mapAgencyNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapLauncherConfigFamilyMiniJSONToProto_launch(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigListJSONToProto_launch(r *LauncherConfigListJSON) *launchv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	l := &launchv1.LauncherConfigList{
		Families: func() []*launchv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*launchv1.LauncherConfigFamilyMini, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyMiniJSONToProto_launch(&v)
			}
			return res
		}(),
		FullName: r.FullName,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		Variant: r.Variant,
	}
	return l
}

func mapLauncherNormalJSONToProto_launch(r *LauncherNormalJSON) *launchv1.LauncherNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.LauncherNormal{
		AttemptedLandings: r.AttemptedLandings,
		Details: r.Details,
		FastestTurnaround: r.FastestTurnaround,
		FirstLaunchDate: r.FirstLaunchDate,
		FlightProven: r.FlightProven,
		Flights: r.Flights,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		IsPlaceholder: r.IsPlaceholder,
		LastLaunchDate: r.LastLaunchDate,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		Status: mapLauncherStatusJSONToProto_launch(r.Status),
		SuccessfulLandings: r.SuccessfulLandings,
		Url: r.Url,
	}
	return l
}

func mapLauncherStatusJSONToProto_launch(r *LauncherStatusJSON) *launchv1.LauncherStatus {
	if r == nil {
		return nil
	}
	l := &launchv1.LauncherStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLocationJSONToProto_launch(r *LocationJSON) *launchv1.Location {
	if r == nil {
		return nil
	}
	l := &launchv1.Location{
		Active: r.Active,
		CelestialBody: mapCelestialBodyDetailedJSONToProto_launch(r.CelestialBody),
		Country: mapCountryJSONToProto_launch(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapLocationSerializerNoCelestialBodyJSONToProto_launch(r *LocationSerializerNoCelestialBodyJSON) *launchv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	l := &launchv1.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapCountryJSONToProto_launch(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapMissionJSONToProto_launch(r *MissionJSON) *launchv1.Mission {
	if r == nil {
		return nil
	}
	l := &launchv1.Mission{
		Agencies: func() []*launchv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*launchv1.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyDetailedJSONToProto_launch(&v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoUrls: func() []*launchv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*launchv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_launch(&v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapOrbitJSONToProto_launch(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*launchv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*launchv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_launch(&v)
			}
			return res
		}(),
	}
	return l
}

func mapMissionPatchJSONToProto_launch(r *MissionPatchJSON) *launchv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &launchv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_launch(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapNetPrecisionJSONToProto_launch(r *NetPrecisionJSON) *launchv1.NetPrecision {
	if r == nil {
		return nil
	}
	l := &launchv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapOrbitJSONToProto_launch(r *OrbitJSON) *launchv1.Orbit {
	if r == nil {
		return nil
	}
	l := &launchv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapCelestialBodyMiniJSONToProto_launch(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapPadJSONToProto_launch(r *PadJSON) *launchv1.Pad {
	if r == nil {
		return nil
	}
	l := &launchv1.Pad{
		Active: r.Active,
		Agencies: func() []*launchv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*launchv1.AgencyNormal, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		Country: mapCountryJSONToProto_launch(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Location: mapLocationJSONToProto_launch(r.Location),
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapPayloadDetailedJSONToProto_launch(r *PayloadDetailedJSON) *launchv1.PayloadDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.PayloadDetailed{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapAgencyDetailedJSONToProto_launch(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapAgencyDetailedJSONToProto_launch(r.Operator),
		Program: func() []*launchv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_launch(r.TypeVal),
		WikiLink: r.WikiLink,
	}
	return l
}

func mapPayloadFlightNormalJSONToProto_launch(r *PayloadFlightNormalJSON) *launchv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.PayloadFlightNormal{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapLandingJSONToProto_launch(r.Landing),
		Launch: mapLaunchNormalJSONToProto_launch(r.Launch),
		Payload: mapPayloadNormalJSONToProto_launch(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
	return l
}

func mapPayloadFlightSerializerNoLaunchJSONToProto_launch(r *PayloadFlightSerializerNoLaunchJSON) *launchv1.PayloadFlightSerializerNoLaunch {
	if r == nil {
		return nil
	}
	l := &launchv1.PayloadFlightSerializerNoLaunch{
		Amount: r.Amount,
		Destination: r.Destination,
		DockingEvents: func() []*launchv1.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]*launchv1.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = mapDockingEventForChaserNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		Id: r.Id,
		Landing: mapLandingJSONToProto_launch(r.Landing),
		Payload: mapPayloadDetailedJSONToProto_launch(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
	return l
}

func mapPayloadMiniJSONToProto_launch(r *PayloadMiniJSON) *launchv1.PayloadMini {
	if r == nil {
		return nil
	}
	l := &launchv1.PayloadMini{
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		Manufacturer: mapAgencyMiniJSONToProto_launch(r.Manufacturer),
		Name: r.Name,
		Operator: mapAgencyMiniJSONToProto_launch(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_launch(r.TypeVal),
	}
	return l
}

func mapPayloadNormalJSONToProto_launch(r *PayloadNormalJSON) *launchv1.PayloadNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapAgencyNormalJSONToProto_launch(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapAgencyNormalJSONToProto_launch(r.Operator),
		Program: func() []*launchv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramMini, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramMiniJSONToProto_launch(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_launch(r.TypeVal),
		WikiLink: r.WikiLink,
	}
	return l
}

func mapPayloadTypeJSONToProto_launch(r *PayloadTypeJSON) *launchv1.PayloadType {
	if r == nil {
		return nil
	}
	l := &launchv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapProgramMiniJSONToProto_launch(r *ProgramMiniJSON) *launchv1.ProgramMini {
	if r == nil {
		return nil
	}
	l := &launchv1.ProgramMini{
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramNormalJSONToProto_launch(r *ProgramNormalJSON) *launchv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.ProgramNormal{
		Agencies: func() []*launchv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*launchv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_launch(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*launchv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*launchv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_launch(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_launch(r *ProgramTypeJSON) *launchv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &launchv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapRocketDetailedJSONToProto_launch(r *RocketDetailedJSON) *launchv1.RocketDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.RocketDetailed{
		Configuration: mapLauncherConfigDetailedJSONToProto_launch(r.Configuration),
		Id: r.Id,
		LauncherStage: func() []*launchv1.FirstStageNormal {
			if r.LauncherStage == nil {
				return nil
			}
			res := make([]*launchv1.FirstStageNormal, len(r.LauncherStage))
			for i, v := range r.LauncherStage {
				res[i] = mapFirstStageNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		Payloads: func() []*launchv1.PayloadFlightSerializerNoLaunch {
			if r.Payloads == nil {
				return nil
			}
			res := make([]*launchv1.PayloadFlightSerializerNoLaunch, len(r.Payloads))
			for i, v := range r.Payloads {
				res[i] = mapPayloadFlightSerializerNoLaunchJSONToProto_launch(&v)
			}
			return res
		}(),
		SpacecraftStage: func() []*launchv1.SpacecraftFlightDetailedSerializerNoLaunch {
			if r.SpacecraftStage == nil {
				return nil
			}
			res := make([]*launchv1.SpacecraftFlightDetailedSerializerNoLaunch, len(r.SpacecraftStage))
			for i, v := range r.SpacecraftStage {
				res[i] = mapSpacecraftFlightDetailedSerializerNoLaunchJSONToProto_launch(&v)
			}
			return res
		}(),
	}
	return l
}

func mapRocketNormalJSONToProto_launch(r *RocketNormalJSON) *launchv1.RocketNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.RocketNormal{
		Configuration: mapLauncherConfigListJSONToProto_launch(r.Configuration),
		Id: r.Id,
	}
	return l
}

func mapSocialMediaJSONToProto_launch(r *SocialMediaJSON) *launchv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &launchv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_launch(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_launch(r *SocialMediaLinkJSON) *launchv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &launchv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_launch(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationMiniJSONToProto_launch(r *SpaceStationMiniJSON) *launchv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	l := &launchv1.SpaceStationMini{
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSpaceStationNormalJSONToProto_launch(r *SpaceStationNormalJSON) *launchv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.SpaceStationNormal{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Status: mapSpaceStationStatusJSONToProto_launch(r.Status),
		Type: mapSpaceStationTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationStatusJSONToProto_launch(r *SpaceStationStatusJSON) *launchv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	l := &launchv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpaceStationTypeJSONToProto_launch(r *SpaceStationTypeJSON) *launchv1.SpaceStationType {
	if r == nil {
		return nil
	}
	l := &launchv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftConfigDetailedJSONToProto_launch(r *SpacecraftConfigDetailedJSON) *launchv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftConfigDetailed{
		Agency: mapAgencyNormalJSONToProto_launch(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*launchv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*launchv1.SpacecraftConfigFamilyDetailed, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyDetailedJSONToProto_launch(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FlightLife: r.FlightLife,
		Height: r.Height,
		History: r.History,
		HumanRated: r.HumanRated,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InUse: r.InUse,
		InfoLink: r.InfoLink,
		MaidenFlight: r.MaidenFlight,
		Name: r.Name,
		PayloadCapacity: r.PayloadCapacity,
		PayloadReturnCapacity: r.PayloadReturnCapacity,
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapSpacecraftConfigTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
		WikiLink: r.WikiLink,
	}
	return l
}

func mapSpacecraftConfigFamilyDetailedJSONToProto_launch(r *SpacecraftConfigFamilyDetailedJSON) *launchv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftConfigFamilyDetailed{
		AttemptedLandings: r.AttemptedLandings,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyNormalJSONToProto_launch(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyNormalJSONToProto_launch(r.Parent),
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapSpacecraftConfigFamilyMiniJSONToProto_launch(r *SpacecraftConfigFamilyMiniJSON) *launchv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigFamilyNormalJSONToProto_launch(r *SpacecraftConfigFamilyNormalJSON) *launchv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyMiniJSONToProto_launch(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyMiniJSONToProto_launch(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigNormalJSONToProto_launch(r *SpacecraftConfigNormalJSON) *launchv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftConfigNormal{
		Agency: mapAgencyMiniJSONToProto_launch(r.Agency),
		Family: func() []*launchv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*launchv1.SpacecraftConfigFamilyNormal, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InUse: r.InUse,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapSpacecraftConfigTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftConfigTypeJSONToProto_launch(r *SpacecraftConfigTypeJSON) *launchv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftDetailedJSONToProto_launch(r *SpacecraftDetailedJSON) *launchv1.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftDetailed{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigDetailedJSONToProto_launch(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_launch(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightDetailedSerializerNoLaunchJSONToProto_launch(r *SpacecraftFlightDetailedSerializerNoLaunchJSON) *launchv1.SpacecraftFlightDetailedSerializerNoLaunch {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftFlightDetailedSerializerNoLaunch{
		Destination: r.Destination,
		DockingEvents: func() []*launchv1.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]*launchv1.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = mapDockingEventForChaserNormalJSONToProto_launch(&v)
			}
			return res
		}(),
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_launch(r.Landing),
		LandingCrew: func() []*launchv1.AstronautFlight {
			if r.LandingCrew == nil {
				return nil
			}
			res := make([]*launchv1.AstronautFlight, len(r.LandingCrew))
			for i, v := range r.LandingCrew {
				res[i] = mapAstronautFlightJSONToProto_launch(&v)
			}
			return res
		}(),
		LaunchCrew: func() []*launchv1.AstronautFlight {
			if r.LaunchCrew == nil {
				return nil
			}
			res := make([]*launchv1.AstronautFlight, len(r.LaunchCrew))
			for i, v := range r.LaunchCrew {
				res[i] = mapAstronautFlightJSONToProto_launch(&v)
			}
			return res
		}(),
		MissionEnd: r.MissionEnd,
		OnboardCrew: func() []*launchv1.AstronautFlight {
			if r.OnboardCrew == nil {
				return nil
			}
			res := make([]*launchv1.AstronautFlight, len(r.OnboardCrew))
			for i, v := range r.OnboardCrew {
				res[i] = mapAstronautFlightJSONToProto_launch(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftDetailedJSONToProto_launch(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightNormalJSONToProto_launch(r *SpacecraftFlightNormalJSON) *launchv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_launch(r.Landing),
		Launch: mapLaunchNormalJSONToProto_launch(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftNormalJSONToProto_launch(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftNormalJSONToProto_launch(r *SpacecraftNormalJSON) *launchv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftNormal{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_launch(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigNormalJSONToProto_launch(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_launch(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftStatusJSONToProto_launch(r *SpacecraftStatusJSON) *launchv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	l := &launchv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapTimelineEventJSONToProto_launch(r *TimelineEventJSON) *launchv1.TimelineEvent {
	if r == nil {
		return nil
	}
	l := &launchv1.TimelineEvent{
		RelativeTime: r.RelativeTime,
		Type: mapTimelineEventTypeJSONToProto_launch(r.TypeVal),
	}
	return l
}

func mapTimelineEventTypeJSONToProto_launch(r *TimelineEventTypeJSON) *launchv1.TimelineEventType {
	if r == nil {
		return nil
	}
	l := &launchv1.TimelineEventType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
	}
	return l
}

func mapUpdateJSONToProto_launch(r *UpdateJSON) *launchv1.Update {
	if r == nil {
		return nil
	}
	l := &launchv1.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
	return l
}

func mapVidURLJSONToProto_launch(r *VidURLJSON) *launchv1.VidURL {
	if r == nil {
		return nil
	}
	l := &launchv1.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_launch(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		Type: mapVidURLTypeJSONToProto_launch(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapVidURLTypeJSONToProto_launch(r *VidURLTypeJSON) *launchv1.VidURLType {
	if r == nil {
		return nil
	}
	l := &launchv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAgencyMiniJSONToProto_location(r *AgencyMiniJSON) *locationv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &locationv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_location(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_location(r *AgencyTypeJSON) *locationv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &locationv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCelestialBodyDetailedJSONToProto_location(r *CelestialBodyDetailedJSON) *locationv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	l := &locationv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_location(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_location(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_location(r *CelestialBodyTypeJSON) *locationv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &locationv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_location(r *CountryJSON) *locationv1.Country {
	if r == nil {
		return nil
	}
	l := &locationv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapImageJSONToProto_location(r *ImageJSON) *locationv1.Image {
	if r == nil {
		return nil
	}
	l := &locationv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_location(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*locationv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*locationv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_location(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_location(r *ImageLicenseJSON) *locationv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &locationv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_location(r *ImageVariantJSON) *locationv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &locationv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_location(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_location(r *ImageVariantTypeJSON) *locationv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &locationv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLocationSerializerWithPadsJSONToProto_location(r *LocationSerializerWithPadsJSON) *locationv1.Location {
	if r == nil {
		return nil
	}
	l := &locationv1.Location{
		Active: r.Active,
		CelestialBody: mapCelestialBodyDetailedJSONToProto_location(r.CelestialBody),
		Country: mapCountryJSONToProto_location(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_location(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		Pads: func() []*locationv1.PadSerializerNoLocation {
			if r.Pads == nil {
				return nil
			}
			res := make([]*locationv1.PadSerializerNoLocation, len(r.Pads))
			for i, v := range r.Pads {
				res[i] = mapPadSerializerNoLocationJSONToProto_location(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapPadSerializerNoLocationJSONToProto_location(r *PadSerializerNoLocationJSON) *locationv1.PadSerializerNoLocation {
	if r == nil {
		return nil
	}
	l := &locationv1.PadSerializerNoLocation{
		Active: r.Active,
		Agencies: func() []*locationv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*locationv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_location(&v)
			}
			return res
		}(),
		Country: mapCountryJSONToProto_location(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapImageJSONToProto_location(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyNormalJSONToProto_pad(r *AgencyNormalJSON) *padv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &padv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*padv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*padv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_pad(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_pad(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_pad(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_pad(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_pad(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_pad(r *AgencyTypeJSON) *padv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &padv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCelestialBodyDetailedJSONToProto_pad(r *CelestialBodyDetailedJSON) *padv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	l := &padv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_pad(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_pad(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_pad(r *CelestialBodyTypeJSON) *padv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &padv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_pad(r *CountryJSON) *padv1.Country {
	if r == nil {
		return nil
	}
	l := &padv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapImageJSONToProto_pad(r *ImageJSON) *padv1.Image {
	if r == nil {
		return nil
	}
	l := &padv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_pad(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*padv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*padv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_pad(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_pad(r *ImageLicenseJSON) *padv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &padv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_pad(r *ImageVariantJSON) *padv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &padv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_pad(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_pad(r *ImageVariantTypeJSON) *padv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &padv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLocationJSONToProto_pad(r *LocationJSON) *padv1.Location {
	if r == nil {
		return nil
	}
	l := &padv1.Location{
		Active: r.Active,
		CelestialBody: mapCelestialBodyDetailedJSONToProto_pad(r.CelestialBody),
		Country: mapCountryJSONToProto_pad(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_pad(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapPadJSONToProto_pad(r *PadJSON) *padv1.Pad {
	if r == nil {
		return nil
	}
	l := &padv1.Pad{
		Active: r.Active,
		Agencies: func() []*padv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*padv1.AgencyNormal, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyNormalJSONToProto_pad(&v)
			}
			return res
		}(),
		Country: mapCountryJSONToProto_pad(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapImageJSONToProto_pad(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Location: mapLocationJSONToProto_pad(r.Location),
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyDetailedJSONToProto_payload(r *AgencyDetailedJSON) *payloadv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &payloadv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*payloadv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*payloadv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_payload(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_payload(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_payload(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_payload(r.SocialLogo),
		SocialMediaLinks: func() []*payloadv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*payloadv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_payload(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_payload(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_payload(r *AgencyMiniJSON) *payloadv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &payloadv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_payload(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_payload(r *AgencyTypeJSON) *payloadv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &payloadv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_payload(r *CountryJSON) *payloadv1.Country {
	if r == nil {
		return nil
	}
	l := &payloadv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapImageJSONToProto_payload(r *ImageJSON) *payloadv1.Image {
	if r == nil {
		return nil
	}
	l := &payloadv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_payload(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*payloadv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*payloadv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_payload(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_payload(r *ImageLicenseJSON) *payloadv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &payloadv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_payload(r *ImageVariantJSON) *payloadv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &payloadv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_payload(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_payload(r *ImageVariantTypeJSON) *payloadv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &payloadv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapMissionPatchJSONToProto_payload(r *MissionPatchJSON) *payloadv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &payloadv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_payload(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapPayloadDetailedJSONToProto_payload(r *PayloadDetailedJSON) *payloadv1.Payload {
	if r == nil {
		return nil
	}
	l := &payloadv1.Payload{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_payload(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapAgencyDetailedJSONToProto_payload(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapAgencyDetailedJSONToProto_payload(r.Operator),
		Program: func() []*payloadv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*payloadv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_payload(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_payload(r.TypeVal),
		WikiLink: r.WikiLink,
	}
	return l
}

func mapPayloadTypeJSONToProto_payload(r *PayloadTypeJSON) *payloadv1.PayloadType {
	if r == nil {
		return nil
	}
	l := &payloadv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapProgramNormalJSONToProto_payload(r *ProgramNormalJSON) *payloadv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &payloadv1.ProgramNormal{
		Agencies: func() []*payloadv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*payloadv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_payload(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_payload(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*payloadv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*payloadv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_payload(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_payload(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_payload(r *ProgramTypeJSON) *payloadv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &payloadv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSocialMediaJSONToProto_payload(r *SocialMediaJSON) *payloadv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &payloadv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_payload(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_payload(r *SocialMediaLinkJSON) *payloadv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &payloadv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_payload(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapAgencyMiniJSONToProto_program(r *AgencyMiniJSON) *programv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &programv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_program(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_program(r *AgencyTypeJSON) *programv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &programv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapImageJSONToProto_program(r *ImageJSON) *programv1.Image {
	if r == nil {
		return nil
	}
	l := &programv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_program(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*programv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*programv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_program(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_program(r *ImageLicenseJSON) *programv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &programv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_program(r *ImageVariantJSON) *programv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &programv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_program(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_program(r *ImageVariantTypeJSON) *programv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &programv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapMissionPatchJSONToProto_program(r *MissionPatchJSON) *programv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &programv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_program(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapProgramNormalJSONToProto_program(r *ProgramNormalJSON) *programv1.Program {
	if r == nil {
		return nil
	}
	l := &programv1.Program{
		Agencies: func() []*programv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*programv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_program(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_program(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*programv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*programv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_program(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_program(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_program(r *ProgramTypeJSON) *programv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &programv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAgencyDetailedJSONToProto_space_station(r *AgencyDetailedJSON) *space_stationv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &space_stationv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*space_stationv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*space_stationv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_space_station(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_space_station(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_space_station(r.SocialLogo),
		SocialMediaLinks: func() []*space_stationv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*space_stationv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_space_station(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_space_station(r *AgencyMiniJSON) *space_stationv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &space_stationv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_space_station(r *AgencyNormalJSON) *space_stationv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*space_stationv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*space_stationv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_space_station(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_space_station(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_space_station(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_space_station(r *AgencyTypeJSON) *space_stationv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCelestialBodyDetailedJSONToProto_space_station(r *CelestialBodyDetailedJSON) *space_stationv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	l := &space_stationv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_space_station(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyMiniJSONToProto_space_station(r *CelestialBodyMiniJSON) *space_stationv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	l := &space_stationv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapCelestialBodyNormalJSONToProto_space_station(r *CelestialBodyNormalJSON) *space_stationv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.CelestialBodyNormal{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapCelestialBodyTypeJSONToProto_space_station(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_space_station(r *CelestialBodyTypeJSON) *space_stationv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_space_station(r *CountryJSON) *space_stationv1.Country {
	if r == nil {
		return nil
	}
	l := &space_stationv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapDockingEventDetailedSerializerForSpacestationJSONToProto_space_station(r *DockingEventDetailedSerializerForSpacestationJSON) *space_stationv1.DockingEventDetailedSerializerForSpacestation {
	if r == nil {
		return nil
	}
	l := &space_stationv1.DockingEventDetailedSerializerForSpacestation{
		Departure: r.Departure,
		Docking: r.Docking,
		FlightVehicleChaser: mapSpacecraftFlightForDockingEventJSONToProto_space_station(r.FlightVehicleChaser),
		Id: r.Id,
		PayloadFlightChaser: mapPayloadFlightNormalJSONToProto_space_station(r.PayloadFlightChaser),
		SpaceStationChaser: mapSpaceStationNormalJSONToProto_space_station(r.SpaceStationChaser),
		Url: r.Url,
	}
	return l
}

func mapDockingEventForChaserNormalJSONToProto_space_station(r *DockingEventForChaserNormalJSON) *space_stationv1.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.DockingEventForChaserNormal{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapDockingLocationJSONToProto_space_station(r.DockingLocation),
		FlightVehicleTarget: mapSpacecraftFlightNormalJSONToProto_space_station(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightTarget: mapPayloadFlightNormalJSONToProto_space_station(r.PayloadFlightTarget),
		SpaceStationTarget: mapSpaceStationNormalJSONToProto_space_station(r.SpaceStationTarget),
		Url: r.Url,
	}
	return l
}

func mapDockingLocationJSONToProto_space_station(r *DockingLocationJSON) *space_stationv1.DockingLocation {
	if r == nil {
		return nil
	}
	l := &space_stationv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapPayloadMiniJSONToProto_space_station(r.Payload),
		Spacecraft: mapSpacecraftConfigNormalJSONToProto_space_station(r.Spacecraft),
		Spacestation: mapSpaceStationMiniJSONToProto_space_station(r.Spacestation),
	}
	return l
}

func mapDockingLocationSerializerForSpacestationJSONToProto_space_station(r *DockingLocationSerializerForSpacestationJSON) *space_stationv1.DockingLocationSerializerForSpacestation {
	if r == nil {
		return nil
	}
	l := &space_stationv1.DockingLocationSerializerForSpacestation{
		CurrentlyDocked: mapDockingEventDetailedSerializerForSpacestationJSONToProto_space_station(r.CurrentlyDocked),
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapExpeditionMiniJSONToProto_space_station(r *ExpeditionMiniJSON) *space_stationv1.ExpeditionMini {
	if r == nil {
		return nil
	}
	l := &space_stationv1.ExpeditionMini{
		End: r.End,
		Id: r.Id,
		Name: r.Name,
		Start: r.Start,
		Url: r.Url,
	}
	return l
}

func mapImageJSONToProto_space_station(r *ImageJSON) *space_stationv1.Image {
	if r == nil {
		return nil
	}
	l := &space_stationv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_space_station(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*space_stationv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*space_stationv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_space_station(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_space_station(r *ImageLicenseJSON) *space_stationv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &space_stationv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_space_station(r *ImageVariantJSON) *space_stationv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &space_stationv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_space_station(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_space_station(r *ImageVariantTypeJSON) *space_stationv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapInfoURLJSONToProto_space_station(r *InfoURLJSON) *space_stationv1.InfoURL {
	if r == nil {
		return nil
	}
	l := &space_stationv1.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_space_station(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		Type: mapInfoURLTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapInfoURLTypeJSONToProto_space_station(r *InfoURLTypeJSON) *space_stationv1.InfoURLType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLandingJSONToProto_space_station(r *LandingJSON) *space_stationv1.Landing {
	if r == nil {
		return nil
	}
	l := &space_stationv1.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Id: r.Id,
		LandingLocation: mapLandingLocationJSONToProto_space_station(r.LandingLocation),
		Success: r.Success,
		Type: mapLandingTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapLandingLocationJSONToProto_space_station(r *LandingLocationJSON) *space_stationv1.LandingLocation {
	if r == nil {
		return nil
	}
	l := &space_stationv1.LandingLocation{
		Abbrev: r.Abbrev,
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		CelestialBody: mapCelestialBodyNormalJSONToProto_space_station(r.CelestialBody),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Latitude: r.Latitude,
		Location: mapLocationSerializerNoCelestialBodyJSONToProto_space_station(r.Location),
		Longitude: r.Longitude,
		Name: r.Name,
		SuccessfulLandings: r.SuccessfulLandings,
	}
	return l
}

func mapLandingTypeJSONToProto_space_station(r *LandingTypeJSON) *space_stationv1.LandingType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLanguageJSONToProto_space_station(r *LanguageJSON) *space_stationv1.Language {
	if r == nil {
		return nil
	}
	l := &space_stationv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLaunchNormalJSONToProto_space_station(r *LaunchNormalJSON) *space_stationv1.LaunchNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.LaunchNormal{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapAgencyMiniJSONToProto_space_station(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapMissionJSONToProto_space_station(r.Mission),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapNetPrecisionJSONToProto_space_station(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapPadJSONToProto_space_station(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		Probability: r.Probability,
		Program: func() []*space_stationv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*space_stationv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_space_station(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapRocketNormalJSONToProto_space_station(r.Rocket),
		Slug: r.Slug,
		Status: mapLaunchStatusJSONToProto_space_station(r.Status),
		Url: r.Url,
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
	return l
}

func mapLaunchStatusJSONToProto_space_station(r *LaunchStatusJSON) *space_stationv1.LaunchStatus {
	if r == nil {
		return nil
	}
	l := &space_stationv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_space_station(r *LauncherConfigFamilyMiniJSON) *space_stationv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &space_stationv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigListJSONToProto_space_station(r *LauncherConfigListJSON) *space_stationv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	l := &space_stationv1.LauncherConfigList{
		Families: func() []*space_stationv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*space_stationv1.LauncherConfigFamilyMini, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyMiniJSONToProto_space_station(&v)
			}
			return res
		}(),
		FullName: r.FullName,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		Variant: r.Variant,
	}
	return l
}

func mapLocationJSONToProto_space_station(r *LocationJSON) *space_stationv1.Location {
	if r == nil {
		return nil
	}
	l := &space_stationv1.Location{
		Active: r.Active,
		CelestialBody: mapCelestialBodyDetailedJSONToProto_space_station(r.CelestialBody),
		Country: mapCountryJSONToProto_space_station(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapLocationSerializerNoCelestialBodyJSONToProto_space_station(r *LocationSerializerNoCelestialBodyJSON) *space_stationv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	l := &space_stationv1.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapCountryJSONToProto_space_station(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapMissionJSONToProto_space_station(r *MissionJSON) *space_stationv1.Mission {
	if r == nil {
		return nil
	}
	l := &space_stationv1.Mission{
		Agencies: func() []*space_stationv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*space_stationv1.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyDetailedJSONToProto_space_station(&v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InfoUrls: func() []*space_stationv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*space_stationv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_space_station(&v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapOrbitJSONToProto_space_station(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*space_stationv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*space_stationv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_space_station(&v)
			}
			return res
		}(),
	}
	return l
}

func mapMissionPatchJSONToProto_space_station(r *MissionPatchJSON) *space_stationv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &space_stationv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_space_station(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapNetPrecisionJSONToProto_space_station(r *NetPrecisionJSON) *space_stationv1.NetPrecision {
	if r == nil {
		return nil
	}
	l := &space_stationv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapOrbitJSONToProto_space_station(r *OrbitJSON) *space_stationv1.Orbit {
	if r == nil {
		return nil
	}
	l := &space_stationv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapCelestialBodyMiniJSONToProto_space_station(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapPadJSONToProto_space_station(r *PadJSON) *space_stationv1.Pad {
	if r == nil {
		return nil
	}
	l := &space_stationv1.Pad{
		Active: r.Active,
		Agencies: func() []*space_stationv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*space_stationv1.AgencyNormal, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyNormalJSONToProto_space_station(&v)
			}
			return res
		}(),
		Country: mapCountryJSONToProto_space_station(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Location: mapLocationJSONToProto_space_station(r.Location),
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapPayloadFlightNormalJSONToProto_space_station(r *PayloadFlightNormalJSON) *space_stationv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.PayloadFlightNormal{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapLandingJSONToProto_space_station(r.Landing),
		Launch: mapLaunchNormalJSONToProto_space_station(r.Launch),
		Payload: mapPayloadNormalJSONToProto_space_station(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
	return l
}

func mapPayloadMiniJSONToProto_space_station(r *PayloadMiniJSON) *space_stationv1.PayloadMini {
	if r == nil {
		return nil
	}
	l := &space_stationv1.PayloadMini{
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Manufacturer: mapAgencyMiniJSONToProto_space_station(r.Manufacturer),
		Name: r.Name,
		Operator: mapAgencyMiniJSONToProto_space_station(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_space_station(r.TypeVal),
	}
	return l
}

func mapPayloadNormalJSONToProto_space_station(r *PayloadNormalJSON) *space_stationv1.PayloadNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapAgencyNormalJSONToProto_space_station(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapAgencyNormalJSONToProto_space_station(r.Operator),
		Program: func() []*space_stationv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*space_stationv1.ProgramMini, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramMiniJSONToProto_space_station(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_space_station(r.TypeVal),
		WikiLink: r.WikiLink,
	}
	return l
}

func mapPayloadTypeJSONToProto_space_station(r *PayloadTypeJSON) *space_stationv1.PayloadType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapProgramMiniJSONToProto_space_station(r *ProgramMiniJSON) *space_stationv1.ProgramMini {
	if r == nil {
		return nil
	}
	l := &space_stationv1.ProgramMini{
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramNormalJSONToProto_space_station(r *ProgramNormalJSON) *space_stationv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.ProgramNormal{
		Agencies: func() []*space_stationv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*space_stationv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_space_station(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*space_stationv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*space_stationv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_space_station(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_space_station(r *ProgramTypeJSON) *space_stationv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapRocketNormalJSONToProto_space_station(r *RocketNormalJSON) *space_stationv1.RocketNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.RocketNormal{
		Configuration: mapLauncherConfigListJSONToProto_space_station(r.Configuration),
		Id: r.Id,
	}
	return l
}

func mapSocialMediaJSONToProto_space_station(r *SocialMediaJSON) *space_stationv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_space_station(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_space_station(r *SocialMediaLinkJSON) *space_stationv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_space_station(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationDetailedEndpointJSONToProto_space_station(r *SpaceStationDetailedEndpointJSON) *space_stationv1.SpaceStation {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpaceStation{
		ActiveDockingEvents: func() []*space_stationv1.DockingEventForChaserNormal {
			if r.ActiveDockingEvents == nil {
				return nil
			}
			res := make([]*space_stationv1.DockingEventForChaserNormal, len(r.ActiveDockingEvents))
			for i, v := range r.ActiveDockingEvents {
				res[i] = mapDockingEventForChaserNormalJSONToProto_space_station(&v)
			}
			return res
		}(),
		ActiveExpeditions: func() []*space_stationv1.ExpeditionMini {
			if r.ActiveExpeditions == nil {
				return nil
			}
			res := make([]*space_stationv1.ExpeditionMini, len(r.ActiveExpeditions))
			for i, v := range r.ActiveExpeditions {
				res[i] = mapExpeditionMiniJSONToProto_space_station(&v)
			}
			return res
		}(),
		Deorbited: r.Deorbited,
		Description: r.Description,
		DockedVehicles: r.DockedVehicles,
		DockingLocation: func() []*space_stationv1.DockingLocationSerializerForSpacestation {
			if r.DockingLocation == nil {
				return nil
			}
			res := make([]*space_stationv1.DockingLocationSerializerForSpacestation, len(r.DockingLocation))
			for i, v := range r.DockingLocation {
				res[i] = mapDockingLocationSerializerForSpacestationJSONToProto_space_station(&v)
			}
			return res
		}(),
		Founded: r.Founded,
		Height: r.Height,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Mass: r.Mass,
		Name: r.Name,
		OnboardCrew: r.OnboardCrew,
		Orbit: r.Orbit,
		Owners: func() []*space_stationv1.AgencyNormal {
			if r.Owners == nil {
				return nil
			}
			res := make([]*space_stationv1.AgencyNormal, len(r.Owners))
			for i, v := range r.Owners {
				res[i] = mapAgencyNormalJSONToProto_space_station(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Status: mapSpaceStationStatusJSONToProto_space_station(r.Status),
		Type: mapSpaceStationTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
		Volume: r.Volume,
		Width: r.Width,
	}
	return l
}

func mapSpaceStationMiniJSONToProto_space_station(r *SpaceStationMiniJSON) *space_stationv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpaceStationMini{
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSpaceStationNormalJSONToProto_space_station(r *SpaceStationNormalJSON) *space_stationv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpaceStationNormal{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Status: mapSpaceStationStatusJSONToProto_space_station(r.Status),
		Type: mapSpaceStationTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationStatusJSONToProto_space_station(r *SpaceStationStatusJSON) *space_stationv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpaceStationTypeJSONToProto_space_station(r *SpaceStationTypeJSON) *space_stationv1.SpaceStationType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftConfigDetailedJSONToProto_space_station(r *SpacecraftConfigDetailedJSON) *space_stationv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftConfigDetailed{
		Agency: mapAgencyNormalJSONToProto_space_station(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*space_stationv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*space_stationv1.SpacecraftConfigFamilyDetailed, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyDetailedJSONToProto_space_station(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FlightLife: r.FlightLife,
		Height: r.Height,
		History: r.History,
		HumanRated: r.HumanRated,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InUse: r.InUse,
		InfoLink: r.InfoLink,
		MaidenFlight: r.MaidenFlight,
		Name: r.Name,
		PayloadCapacity: r.PayloadCapacity,
		PayloadReturnCapacity: r.PayloadReturnCapacity,
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapSpacecraftConfigTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
		WikiLink: r.WikiLink,
	}
	return l
}

func mapSpacecraftConfigFamilyDetailedJSONToProto_space_station(r *SpacecraftConfigFamilyDetailedJSON) *space_stationv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftConfigFamilyDetailed{
		AttemptedLandings: r.AttemptedLandings,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyNormalJSONToProto_space_station(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyNormalJSONToProto_space_station(r.Parent),
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapSpacecraftConfigFamilyMiniJSONToProto_space_station(r *SpacecraftConfigFamilyMiniJSON) *space_stationv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigFamilyNormalJSONToProto_space_station(r *SpacecraftConfigFamilyNormalJSON) *space_stationv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyMiniJSONToProto_space_station(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyMiniJSONToProto_space_station(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigNormalJSONToProto_space_station(r *SpacecraftConfigNormalJSON) *space_stationv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftConfigNormal{
		Agency: mapAgencyMiniJSONToProto_space_station(r.Agency),
		Family: func() []*space_stationv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*space_stationv1.SpacecraftConfigFamilyNormal, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyNormalJSONToProto_space_station(&v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InUse: r.InUse,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapSpacecraftConfigTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftConfigTypeJSONToProto_space_station(r *SpacecraftConfigTypeJSON) *space_stationv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftDetailedJSONToProto_space_station(r *SpacecraftDetailedJSON) *space_stationv1.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftDetailed{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigDetailedJSONToProto_space_station(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_space_station(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightForDockingEventJSONToProto_space_station(r *SpacecraftFlightForDockingEventJSON) *space_stationv1.SpacecraftFlightForDockingEvent {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftFlightForDockingEvent{
		Id: r.Id,
		Launch: mapLaunchNormalJSONToProto_space_station(r.Launch),
		Spacecraft: mapSpacecraftDetailedJSONToProto_space_station(r.Spacecraft),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightNormalJSONToProto_space_station(r *SpacecraftFlightNormalJSON) *space_stationv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_space_station(r.Landing),
		Launch: mapLaunchNormalJSONToProto_space_station(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftNormalJSONToProto_space_station(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftNormalJSONToProto_space_station(r *SpacecraftNormalJSON) *space_stationv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftNormal{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_space_station(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigNormalJSONToProto_space_station(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_space_station(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftStatusJSONToProto_space_station(r *SpacecraftStatusJSON) *space_stationv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	l := &space_stationv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapVidURLJSONToProto_space_station(r *VidURLJSON) *space_stationv1.VidURL {
	if r == nil {
		return nil
	}
	l := &space_stationv1.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_space_station(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		Type: mapVidURLTypeJSONToProto_space_station(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapVidURLTypeJSONToProto_space_station(r *VidURLTypeJSON) *space_stationv1.VidURLType {
	if r == nil {
		return nil
	}
	l := &space_stationv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAgencyDetailedJSONToProto_spacecraft(r *AgencyDetailedJSON) *spacecraftv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*spacecraftv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*spacecraftv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_spacecraft(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_spacecraft(r.SocialLogo),
		SocialMediaLinks: func() []*spacecraftv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*spacecraftv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_spacecraft(r *AgencyMiniJSON) *spacecraftv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_spacecraft(r *AgencyNormalJSON) *spacecraftv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*spacecraftv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*spacecraftv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_spacecraft(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_spacecraft(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_spacecraft(r *AgencyTypeJSON) *spacecraftv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCelestialBodyDetailedJSONToProto_spacecraft(r *CelestialBodyDetailedJSON) *spacecraftv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_spacecraft(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyMiniJSONToProto_spacecraft(r *CelestialBodyMiniJSON) *spacecraftv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapCelestialBodyNormalJSONToProto_spacecraft(r *CelestialBodyNormalJSON) *spacecraftv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.CelestialBodyNormal{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapCelestialBodyTypeJSONToProto_spacecraft(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_spacecraft(r *CelestialBodyTypeJSON) *spacecraftv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_spacecraft(r *CountryJSON) *spacecraftv1.Country {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapImageJSONToProto_spacecraft(r *ImageJSON) *spacecraftv1.Image {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_spacecraft(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*spacecraftv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*spacecraftv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_spacecraft(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_spacecraft(r *ImageLicenseJSON) *spacecraftv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_spacecraft(r *ImageVariantJSON) *spacecraftv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_spacecraft(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_spacecraft(r *ImageVariantTypeJSON) *spacecraftv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapInfoURLJSONToProto_spacecraft(r *InfoURLJSON) *spacecraftv1.InfoURL {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_spacecraft(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		Type: mapInfoURLTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapInfoURLTypeJSONToProto_spacecraft(r *InfoURLTypeJSON) *spacecraftv1.InfoURLType {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLandingJSONToProto_spacecraft(r *LandingJSON) *spacecraftv1.Landing {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Id: r.Id,
		LandingLocation: mapLandingLocationJSONToProto_spacecraft(r.LandingLocation),
		Success: r.Success,
		Type: mapLandingTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapLandingLocationJSONToProto_spacecraft(r *LandingLocationJSON) *spacecraftv1.LandingLocation {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.LandingLocation{
		Abbrev: r.Abbrev,
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		CelestialBody: mapCelestialBodyNormalJSONToProto_spacecraft(r.CelestialBody),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		Latitude: r.Latitude,
		Location: mapLocationSerializerNoCelestialBodyJSONToProto_spacecraft(r.Location),
		Longitude: r.Longitude,
		Name: r.Name,
		SuccessfulLandings: r.SuccessfulLandings,
	}
	return l
}

func mapLandingTypeJSONToProto_spacecraft(r *LandingTypeJSON) *spacecraftv1.LandingType {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLanguageJSONToProto_spacecraft(r *LanguageJSON) *spacecraftv1.Language {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLaunchNormalJSONToProto_spacecraft(r *LaunchNormalJSON) *spacecraftv1.LaunchNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.LaunchNormal{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapAgencyMiniJSONToProto_spacecraft(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapMissionJSONToProto_spacecraft(r.Mission),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapNetPrecisionJSONToProto_spacecraft(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapPadJSONToProto_spacecraft(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		Probability: r.Probability,
		Program: func() []*spacecraftv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*spacecraftv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapRocketNormalJSONToProto_spacecraft(r.Rocket),
		Slug: r.Slug,
		Status: mapLaunchStatusJSONToProto_spacecraft(r.Status),
		Url: r.Url,
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
	return l
}

func mapLaunchStatusJSONToProto_spacecraft(r *LaunchStatusJSON) *spacecraftv1.LaunchStatus {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_spacecraft(r *LauncherConfigFamilyMiniJSON) *spacecraftv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigListJSONToProto_spacecraft(r *LauncherConfigListJSON) *spacecraftv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.LauncherConfigList{
		Families: func() []*spacecraftv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*spacecraftv1.LauncherConfigFamilyMini, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyMiniJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		FullName: r.FullName,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		Variant: r.Variant,
	}
	return l
}

func mapLocationJSONToProto_spacecraft(r *LocationJSON) *spacecraftv1.Location {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Location{
		Active: r.Active,
		CelestialBody: mapCelestialBodyDetailedJSONToProto_spacecraft(r.CelestialBody),
		Country: mapCountryJSONToProto_spacecraft(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapLocationSerializerNoCelestialBodyJSONToProto_spacecraft(r *LocationSerializerNoCelestialBodyJSON) *spacecraftv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapCountryJSONToProto_spacecraft(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapMissionJSONToProto_spacecraft(r *MissionJSON) *spacecraftv1.Mission {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Mission{
		Agencies: func() []*spacecraftv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacecraftv1.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyDetailedJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		InfoUrls: func() []*spacecraftv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*spacecraftv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapOrbitJSONToProto_spacecraft(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*spacecraftv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*spacecraftv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_spacecraft(&v)
			}
			return res
		}(),
	}
	return l
}

func mapMissionPatchJSONToProto_spacecraft(r *MissionPatchJSON) *spacecraftv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_spacecraft(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapNetPrecisionJSONToProto_spacecraft(r *NetPrecisionJSON) *spacecraftv1.NetPrecision {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapOrbitJSONToProto_spacecraft(r *OrbitJSON) *spacecraftv1.Orbit {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapCelestialBodyMiniJSONToProto_spacecraft(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapPadJSONToProto_spacecraft(r *PadJSON) *spacecraftv1.Pad {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Pad{
		Active: r.Active,
		Agencies: func() []*spacecraftv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacecraftv1.AgencyNormal, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyNormalJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Country: mapCountryJSONToProto_spacecraft(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Location: mapLocationJSONToProto_spacecraft(r.Location),
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramNormalJSONToProto_spacecraft(r *ProgramNormalJSON) *spacecraftv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.ProgramNormal{
		Agencies: func() []*spacecraftv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacecraftv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*spacecraftv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*spacecraftv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_spacecraft(r *ProgramTypeJSON) *spacecraftv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapRocketNormalJSONToProto_spacecraft(r *RocketNormalJSON) *spacecraftv1.RocketNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.RocketNormal{
		Configuration: mapLauncherConfigListJSONToProto_spacecraft(r.Configuration),
		Id: r.Id,
	}
	return l
}

func mapSocialMediaJSONToProto_spacecraft(r *SocialMediaJSON) *spacecraftv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_spacecraft(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_spacecraft(r *SocialMediaLinkJSON) *spacecraftv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_spacecraft(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftConfigDetailedJSONToProto_spacecraft(r *SpacecraftConfigDetailedJSON) *spacecraftv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftConfigDetailed{
		Agency: mapAgencyNormalJSONToProto_spacecraft(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*spacecraftv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*spacecraftv1.SpacecraftConfigFamilyDetailed, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyDetailedJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FlightLife: r.FlightLife,
		Height: r.Height,
		History: r.History,
		HumanRated: r.HumanRated,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		InUse: r.InUse,
		InfoLink: r.InfoLink,
		MaidenFlight: r.MaidenFlight,
		Name: r.Name,
		PayloadCapacity: r.PayloadCapacity,
		PayloadReturnCapacity: r.PayloadReturnCapacity,
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapSpacecraftConfigTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
		WikiLink: r.WikiLink,
	}
	return l
}

func mapSpacecraftConfigFamilyDetailedJSONToProto_spacecraft(r *SpacecraftConfigFamilyDetailedJSON) *spacecraftv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftConfigFamilyDetailed{
		AttemptedLandings: r.AttemptedLandings,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyNormalJSONToProto_spacecraft(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyNormalJSONToProto_spacecraft(r.Parent),
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapSpacecraftConfigFamilyMiniJSONToProto_spacecraft(r *SpacecraftConfigFamilyMiniJSON) *spacecraftv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigFamilyNormalJSONToProto_spacecraft(r *SpacecraftConfigFamilyNormalJSON) *spacecraftv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyMiniJSONToProto_spacecraft(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyMiniJSONToProto_spacecraft(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigNormalJSONToProto_spacecraft(r *SpacecraftConfigNormalJSON) *spacecraftv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftConfigNormal{
		Agency: mapAgencyMiniJSONToProto_spacecraft(r.Agency),
		Family: func() []*spacecraftv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*spacecraftv1.SpacecraftConfigFamilyNormal, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyNormalJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		InUse: r.InUse,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapSpacecraftConfigTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftConfigTypeJSONToProto_spacecraft(r *SpacecraftConfigTypeJSON) *spacecraftv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftEndpointDetailedJSONToProto_spacecraft(r *SpacecraftEndpointDetailedJSON) *spacecraftv1.Spacecraft {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.Spacecraft{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Flights: func() []*spacecraftv1.SpacecraftFlightNormal {
			if r.Flights == nil {
				return nil
			}
			res := make([]*spacecraftv1.SpacecraftFlightNormal, len(r.Flights))
			for i, v := range r.Flights {
				res[i] = mapSpacecraftFlightNormalJSONToProto_spacecraft(&v)
			}
			return res
		}(),
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigDetailedJSONToProto_spacecraft(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_spacecraft(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightNormalJSONToProto_spacecraft(r *SpacecraftFlightNormalJSON) *spacecraftv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_spacecraft(r.Landing),
		Launch: mapLaunchNormalJSONToProto_spacecraft(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftNormalJSONToProto_spacecraft(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftNormalJSONToProto_spacecraft(r *SpacecraftNormalJSON) *spacecraftv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftNormal{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_spacecraft(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigNormalJSONToProto_spacecraft(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_spacecraft(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftStatusJSONToProto_spacecraft(r *SpacecraftStatusJSON) *spacecraftv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapVidURLJSONToProto_spacecraft(r *VidURLJSON) *spacecraftv1.VidURL {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_spacecraft(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		Type: mapVidURLTypeJSONToProto_spacecraft(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapVidURLTypeJSONToProto_spacecraft(r *VidURLTypeJSON) *spacecraftv1.VidURLType {
	if r == nil {
		return nil
	}
	l := &spacecraftv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAgencyDetailedJSONToProto_spacewalk(r *AgencyDetailedJSON) *spacewalkv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*spacewalkv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*spacewalkv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_spacewalk(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_spacewalk(r.SocialLogo),
		SocialMediaLinks: func() []*spacewalkv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*spacewalkv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapAgencyTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapAgencyMiniJSONToProto_spacewalk(r *AgencyMiniJSON) *spacewalkv1.AgencyMini {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapAgencyTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyNormalJSONToProto_spacewalk(r *AgencyNormalJSON) *spacewalkv1.AgencyNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*spacewalkv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*spacewalkv1.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = mapCountryJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		Launchers: r.Launchers,
		Logo: mapImageJSONToProto_spacewalk(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapImageJSONToProto_spacewalk(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapAgencyTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapAgencyTypeJSONToProto_spacewalk(r *AgencyTypeJSON) *spacewalkv1.AgencyType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautDetailedJSONToProto_spacewalk(r *AstronautDetailedJSON) *spacewalkv1.AstronautDetailed {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AstronautDetailed{
		Age: r.Age,
		Agency: mapAgencyMiniJSONToProto_spacewalk(r.Agency),
		Bio: r.Bio,
		DateOfBirth: r.DateOfBirth,
		DateOfDeath: r.DateOfDeath,
		EvaTime: r.EvaTime,
		FirstFlight: r.FirstFlight,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InSpace: r.InSpace,
		LastFlight: r.LastFlight,
		Name: r.Name,
		Nationality: func() []*spacewalkv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*spacewalkv1.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = mapCountryJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*spacewalkv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*spacewalkv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = mapSocialMediaLinkJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Status: mapAstronautStatusJSONToProto_spacewalk(r.Status),
		TimeInSpace: r.TimeInSpace,
		Type: mapAstronautTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
		Wiki: r.Wiki,
	}
	return l
}

func mapAstronautFlightJSONToProto_spacewalk(r *AstronautFlightJSON) *spacewalkv1.AstronautFlight {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AstronautFlight{
		Astronaut: mapAstronautDetailedJSONToProto_spacewalk(r.Astronaut),
		Id: r.Id,
		Role: mapAstronautRoleJSONToProto_spacewalk(r.Role),
	}
	return l
}

func mapAstronautRoleJSONToProto_spacewalk(r *AstronautRoleJSON) *spacewalkv1.AstronautRole {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
	return l
}

func mapAstronautStatusJSONToProto_spacewalk(r *AstronautStatusJSON) *spacewalkv1.AstronautStatus {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapAstronautTypeJSONToProto_spacewalk(r *AstronautTypeJSON) *spacewalkv1.AstronautType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCelestialBodyDetailedJSONToProto_spacewalk(r *CelestialBodyDetailedJSON) *spacewalkv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		Type: mapCelestialBodyTypeJSONToProto_spacewalk(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyMiniJSONToProto_spacewalk(r *CelestialBodyMiniJSON) *spacewalkv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapCelestialBodyNormalJSONToProto_spacewalk(r *CelestialBodyNormalJSON) *spacewalkv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.CelestialBodyNormal{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapCelestialBodyTypeJSONToProto_spacewalk(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapCelestialBodyTypeJSONToProto_spacewalk(r *CelestialBodyTypeJSON) *spacewalkv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapCountryJSONToProto_spacewalk(r *CountryJSON) *spacewalkv1.Country {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
	return l
}

func mapDockingEventForChaserNormalJSONToProto_spacewalk(r *DockingEventForChaserNormalJSON) *spacewalkv1.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.DockingEventForChaserNormal{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapDockingLocationJSONToProto_spacewalk(r.DockingLocation),
		FlightVehicleTarget: mapSpacecraftFlightNormalJSONToProto_spacewalk(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightTarget: mapPayloadFlightNormalJSONToProto_spacewalk(r.PayloadFlightTarget),
		SpaceStationTarget: mapSpaceStationNormalJSONToProto_spacewalk(r.SpaceStationTarget),
		Url: r.Url,
	}
	return l
}

func mapDockingLocationJSONToProto_spacewalk(r *DockingLocationJSON) *spacewalkv1.DockingLocation {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapPayloadMiniJSONToProto_spacewalk(r.Payload),
		Spacecraft: mapSpacecraftConfigNormalJSONToProto_spacewalk(r.Spacecraft),
		Spacestation: mapSpaceStationMiniJSONToProto_spacewalk(r.Spacestation),
	}
	return l
}

func mapEventNormalJSONToProto_spacewalk(r *EventNormalJSON) *spacewalkv1.EventNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.EventNormal{
		Date: r.Date,
		DatePrecision: mapNetPrecisionJSONToProto_spacewalk(r.DatePrecision),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InfoUrls: func() []*spacewalkv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*spacewalkv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Location: r.Location,
		Name: r.Name,
		Slug: r.Slug,
		Type: mapEventTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
		VidUrls: func() []*spacewalkv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*spacewalkv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		WebcastLive: r.WebcastLive,
	}
	return l
}

func mapEventTypeJSONToProto_spacewalk(r *EventTypeJSON) *spacewalkv1.EventType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.EventType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapExpeditionNormalSerializerForSpacewalkJSONToProto_spacewalk(r *ExpeditionNormalSerializerForSpacewalkJSON) *spacewalkv1.ExpeditionNormalSerializerForSpacewalk {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.ExpeditionNormalSerializerForSpacewalk{
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []*spacewalkv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*spacewalkv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Name: r.Name,
		Spacestation: mapSpaceStationNormalJSONToProto_spacewalk(r.Spacestation),
		Start: r.Start,
		Url: r.Url,
	}
	return l
}

func mapImageJSONToProto_spacewalk(r *ImageJSON) *spacewalkv1.Image {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapImageLicenseJSONToProto_spacewalk(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*spacewalkv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*spacewalkv1.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = mapImageVariantJSONToProto_spacewalk(&v)
			}
			return res
		}(),
	}
	return l
}

func mapImageLicenseJSONToProto_spacewalk(r *ImageLicenseJSON) *spacewalkv1.ImageLicense {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
	return l
}

func mapImageVariantJSONToProto_spacewalk(r *ImageVariantJSON) *spacewalkv1.ImageVariant {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapImageVariantTypeJSONToProto_spacewalk(r.TypeVal),
	}
	return l
}

func mapImageVariantTypeJSONToProto_spacewalk(r *ImageVariantTypeJSON) *spacewalkv1.ImageVariantType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapInfoURLJSONToProto_spacewalk(r *InfoURLJSON) *spacewalkv1.InfoURL {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_spacewalk(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		Type: mapInfoURLTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapInfoURLTypeJSONToProto_spacewalk(r *InfoURLTypeJSON) *spacewalkv1.InfoURLType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLandingJSONToProto_spacewalk(r *LandingJSON) *spacewalkv1.Landing {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Id: r.Id,
		LandingLocation: mapLandingLocationJSONToProto_spacewalk(r.LandingLocation),
		Success: r.Success,
		Type: mapLandingTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapLandingLocationJSONToProto_spacewalk(r *LandingLocationJSON) *spacewalkv1.LandingLocation {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.LandingLocation{
		Abbrev: r.Abbrev,
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		CelestialBody: mapCelestialBodyNormalJSONToProto_spacewalk(r.CelestialBody),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		Latitude: r.Latitude,
		Location: mapLocationSerializerNoCelestialBodyJSONToProto_spacewalk(r.Location),
		Longitude: r.Longitude,
		Name: r.Name,
		SuccessfulLandings: r.SuccessfulLandings,
	}
	return l
}

func mapLandingTypeJSONToProto_spacewalk(r *LandingTypeJSON) *spacewalkv1.LandingType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLanguageJSONToProto_spacewalk(r *LanguageJSON) *spacewalkv1.Language {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLaunchNormalJSONToProto_spacewalk(r *LaunchNormalJSON) *spacewalkv1.LaunchNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.LaunchNormal{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapAgencyMiniJSONToProto_spacewalk(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapMissionJSONToProto_spacewalk(r.Mission),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapNetPrecisionJSONToProto_spacewalk(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapPadJSONToProto_spacewalk(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		Probability: r.Probability,
		Program: func() []*spacewalkv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*spacewalkv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapRocketNormalJSONToProto_spacewalk(r.Rocket),
		Slug: r.Slug,
		Status: mapLaunchStatusJSONToProto_spacewalk(r.Status),
		Url: r.Url,
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
	return l
}

func mapLaunchStatusJSONToProto_spacewalk(r *LaunchStatusJSON) *spacewalkv1.LaunchStatus {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapLauncherConfigFamilyMiniJSONToProto_spacewalk(r *LauncherConfigFamilyMiniJSON) *spacewalkv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapLauncherConfigListJSONToProto_spacewalk(r *LauncherConfigListJSON) *spacewalkv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.LauncherConfigList{
		Families: func() []*spacewalkv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*spacewalkv1.LauncherConfigFamilyMini, len(r.Families))
			for i, v := range r.Families {
				res[i] = mapLauncherConfigFamilyMiniJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		FullName: r.FullName,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		Variant: r.Variant,
	}
	return l
}

func mapLocationJSONToProto_spacewalk(r *LocationJSON) *spacewalkv1.Location {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Location{
		Active: r.Active,
		CelestialBody: mapCelestialBodyDetailedJSONToProto_spacewalk(r.CelestialBody),
		Country: mapCountryJSONToProto_spacewalk(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapLocationSerializerNoCelestialBodyJSONToProto_spacewalk(r *LocationSerializerNoCelestialBodyJSON) *spacewalkv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapCountryJSONToProto_spacewalk(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
	return l
}

func mapMissionJSONToProto_spacewalk(r *MissionJSON) *spacewalkv1.Mission {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Mission{
		Agencies: func() []*spacewalkv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacewalkv1.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyDetailedJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InfoUrls: func() []*spacewalkv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*spacewalkv1.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = mapInfoURLJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapOrbitJSONToProto_spacewalk(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*spacewalkv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*spacewalkv1.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = mapVidURLJSONToProto_spacewalk(&v)
			}
			return res
		}(),
	}
	return l
}

func mapMissionPatchJSONToProto_spacewalk(r *MissionPatchJSON) *spacewalkv1.MissionPatch {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.MissionPatch{
		Agency: mapAgencyMiniJSONToProto_spacewalk(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapNetPrecisionJSONToProto_spacewalk(r *NetPrecisionJSON) *spacewalkv1.NetPrecision {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapOrbitJSONToProto_spacewalk(r *OrbitJSON) *spacewalkv1.Orbit {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapCelestialBodyMiniJSONToProto_spacewalk(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapPadJSONToProto_spacewalk(r *PadJSON) *spacewalkv1.Pad {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Pad{
		Active: r.Active,
		Agencies: func() []*spacewalkv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacewalkv1.AgencyNormal, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyNormalJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Country: mapCountryJSONToProto_spacewalk(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Location: mapLocationJSONToProto_spacewalk(r.Location),
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapPayloadFlightNormalJSONToProto_spacewalk(r *PayloadFlightNormalJSON) *spacewalkv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.PayloadFlightNormal{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapLandingJSONToProto_spacewalk(r.Landing),
		Launch: mapLaunchNormalJSONToProto_spacewalk(r.Launch),
		Payload: mapPayloadNormalJSONToProto_spacewalk(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
	return l
}

func mapPayloadMiniJSONToProto_spacewalk(r *PayloadMiniJSON) *spacewalkv1.PayloadMini {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.PayloadMini{
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		Manufacturer: mapAgencyMiniJSONToProto_spacewalk(r.Manufacturer),
		Name: r.Name,
		Operator: mapAgencyMiniJSONToProto_spacewalk(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_spacewalk(r.TypeVal),
	}
	return l
}

func mapPayloadNormalJSONToProto_spacewalk(r *PayloadNormalJSON) *spacewalkv1.PayloadNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapAgencyNormalJSONToProto_spacewalk(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapAgencyNormalJSONToProto_spacewalk(r.Operator),
		Program: func() []*spacewalkv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*spacewalkv1.ProgramMini, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramMiniJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapPayloadTypeJSONToProto_spacewalk(r.TypeVal),
		WikiLink: r.WikiLink,
	}
	return l
}

func mapPayloadTypeJSONToProto_spacewalk(r *PayloadTypeJSON) *spacewalkv1.PayloadType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapProgramMiniJSONToProto_spacewalk(r *ProgramMiniJSON) *spacewalkv1.ProgramMini {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.ProgramMini{
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramNormalJSONToProto_spacewalk(r *ProgramNormalJSON) *spacewalkv1.ProgramNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.ProgramNormal{
		Agencies: func() []*spacewalkv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacewalkv1.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = mapAgencyMiniJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*spacewalkv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*spacewalkv1.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = mapMissionPatchJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapProgramTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
	return l
}

func mapProgramTypeJSONToProto_spacewalk(r *ProgramTypeJSON) *spacewalkv1.ProgramType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapRocketNormalJSONToProto_spacewalk(r *RocketNormalJSON) *spacewalkv1.RocketNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.RocketNormal{
		Configuration: mapLauncherConfigListJSONToProto_spacewalk(r.Configuration),
		Id: r.Id,
	}
	return l
}

func mapSocialMediaJSONToProto_spacewalk(r *SocialMediaJSON) *spacewalkv1.SocialMedia {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SocialMedia{
		Id: r.Id,
		Logo: mapImageJSONToProto_spacewalk(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSocialMediaLinkJSONToProto_spacewalk(r *SocialMediaLinkJSON) *spacewalkv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapSocialMediaJSONToProto_spacewalk(r.SocialMedia),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationMiniJSONToProto_spacewalk(r *SpaceStationMiniJSON) *spacewalkv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpaceStationMini{
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
	return l
}

func mapSpaceStationNormalJSONToProto_spacewalk(r *SpaceStationNormalJSON) *spacewalkv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpaceStationNormal{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Status: mapSpaceStationStatusJSONToProto_spacewalk(r.Status),
		Type: mapSpaceStationTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpaceStationStatusJSONToProto_spacewalk(r *SpaceStationStatusJSON) *spacewalkv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpaceStationTypeJSONToProto_spacewalk(r *SpaceStationTypeJSON) *spacewalkv1.SpaceStationType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftConfigDetailedJSONToProto_spacewalk(r *SpacecraftConfigDetailedJSON) *spacewalkv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftConfigDetailed{
		Agency: mapAgencyNormalJSONToProto_spacewalk(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*spacewalkv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*spacewalkv1.SpacecraftConfigFamilyDetailed, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyDetailedJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FlightLife: r.FlightLife,
		Height: r.Height,
		History: r.History,
		HumanRated: r.HumanRated,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InUse: r.InUse,
		InfoLink: r.InfoLink,
		MaidenFlight: r.MaidenFlight,
		Name: r.Name,
		PayloadCapacity: r.PayloadCapacity,
		PayloadReturnCapacity: r.PayloadReturnCapacity,
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapSpacecraftConfigTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
		WikiLink: r.WikiLink,
	}
	return l
}

func mapSpacecraftConfigFamilyDetailedJSONToProto_spacewalk(r *SpacecraftConfigFamilyDetailedJSON) *spacewalkv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftConfigFamilyDetailed{
		AttemptedLandings: r.AttemptedLandings,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyNormalJSONToProto_spacewalk(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyNormalJSONToProto_spacewalk(r.Parent),
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
	return l
}

func mapSpacecraftConfigFamilyMiniJSONToProto_spacewalk(r *SpacecraftConfigFamilyMiniJSON) *spacewalkv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigFamilyNormalJSONToProto_spacewalk(r *SpacecraftConfigFamilyNormalJSON) *spacewalkv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapAgencyMiniJSONToProto_spacewalk(r.Manufacturer),
		Name: r.Name,
		Parent: mapSpacecraftConfigFamilyMiniJSONToProto_spacewalk(r.Parent),
		ResponseMode: r.ResponseMode,
	}
	return l
}

func mapSpacecraftConfigNormalJSONToProto_spacewalk(r *SpacecraftConfigNormalJSON) *spacewalkv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftConfigNormal{
		Agency: mapAgencyMiniJSONToProto_spacewalk(r.Agency),
		Family: func() []*spacewalkv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*spacewalkv1.SpacecraftConfigFamilyNormal, len(r.Family))
			for i, v := range r.Family {
				res[i] = mapSpacecraftConfigFamilyNormalJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InUse: r.InUse,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapSpacecraftConfigTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapSpacecraftConfigTypeJSONToProto_spacewalk(r *SpacecraftConfigTypeJSON) *spacewalkv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacecraftDetailedJSONToProto_spacewalk(r *SpacecraftDetailedJSON) *spacewalkv1.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftDetailed{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigDetailedJSONToProto_spacewalk(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_spacewalk(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightDetailedJSONToProto_spacewalk(r *SpacecraftFlightDetailedJSON) *spacewalkv1.SpacecraftFlightDetailed {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftFlightDetailed{
		Destination: r.Destination,
		DockingEvents: func() []*spacewalkv1.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]*spacewalkv1.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = mapDockingEventForChaserNormalJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_spacewalk(r.Landing),
		LandingCrew: func() []*spacewalkv1.AstronautFlight {
			if r.LandingCrew == nil {
				return nil
			}
			res := make([]*spacewalkv1.AstronautFlight, len(r.LandingCrew))
			for i, v := range r.LandingCrew {
				res[i] = mapAstronautFlightJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Launch: mapLaunchNormalJSONToProto_spacewalk(r.Launch),
		LaunchCrew: func() []*spacewalkv1.AstronautFlight {
			if r.LaunchCrew == nil {
				return nil
			}
			res := make([]*spacewalkv1.AstronautFlight, len(r.LaunchCrew))
			for i, v := range r.LaunchCrew {
				res[i] = mapAstronautFlightJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		MissionEnd: r.MissionEnd,
		OnboardCrew: func() []*spacewalkv1.AstronautFlight {
			if r.OnboardCrew == nil {
				return nil
			}
			res := make([]*spacewalkv1.AstronautFlight, len(r.OnboardCrew))
			for i, v := range r.OnboardCrew {
				res[i] = mapAstronautFlightJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftDetailedJSONToProto_spacewalk(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftFlightNormalJSONToProto_spacewalk(r *SpacecraftFlightNormalJSON) *spacewalkv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapLandingJSONToProto_spacewalk(r.Landing),
		Launch: mapLaunchNormalJSONToProto_spacewalk(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapSpacecraftNormalJSONToProto_spacewalk(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftNormalJSONToProto_spacewalk(r *SpacecraftNormalJSON) *spacewalkv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftNormal{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapImageJSONToProto_spacewalk(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapSpacecraftConfigNormalJSONToProto_spacewalk(r.SpacecraftConfig),
		Status: mapSpacecraftStatusJSONToProto_spacewalk(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
	return l
}

func mapSpacecraftStatusJSONToProto_spacewalk(r *SpacecraftStatusJSON) *spacewalkv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapSpacewalkEndpointDetailedJSONToProto_spacewalk(r *SpacewalkEndpointDetailedJSON) *spacewalkv1.Spacewalk {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.Spacewalk{
		Crew: func() []*spacewalkv1.AstronautFlight {
			if r.Crew == nil {
				return nil
			}
			res := make([]*spacewalkv1.AstronautFlight, len(r.Crew))
			for i, v := range r.Crew {
				res[i] = mapAstronautFlightJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		Duration: r.Duration,
		End: r.End,
		Event: mapEventNormalJSONToProto_spacewalk(r.Event),
		Expedition: mapExpeditionNormalSerializerForSpacewalkJSONToProto_spacewalk(r.Expedition),
		Id: r.Id,
		Location: r.Location,
		Name: r.Name,
		Program: func() []*spacewalkv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*spacewalkv1.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = mapProgramNormalJSONToProto_spacewalk(&v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SpacecraftFlight: mapSpacecraftFlightDetailedJSONToProto_spacewalk(r.SpacecraftFlight),
		Spacestation: mapSpaceStationNormalJSONToProto_spacewalk(r.Spacestation),
		Start: r.Start,
		Url: r.Url,
	}
	return l
}

func mapVidURLJSONToProto_spacewalk(r *VidURLJSON) *spacewalkv1.VidURL {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapLanguageJSONToProto_spacewalk(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		Type: mapVidURLTypeJSONToProto_spacewalk(r.TypeVal),
		Url: r.Url,
	}
	return l
}

func mapVidURLTypeJSONToProto_spacewalk(r *VidURLTypeJSON) *spacewalkv1.VidURLType {
	if r == nil {
		return nil
	}
	l := &spacewalkv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
	return l
}

func mapUpdateJSONToProto_update(r *UpdateJSON) *updatev1.Update {
	if r == nil {
		return nil
	}
	l := &updatev1.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
	return l
}

