package lib

// imported from UpAndRunning2

// ******************
// * USER-RESPONSES *
// ******************

// Contains a success-bool and a message, which may be empty.
type BasicResponse struct {
	Success bool   `json:"requestSuccess"`
	Message string `json:"message"`
}

// Contains a success-bool and an array of BasicWebsites.
type WebsiteResponse struct {
	Success  bool           `json:"requestSuccess"`
	Websites []BasicWebsite `json:"websites"`
}

// Contains a success-bool and the Website's details.
type StatusResponse struct {
	Success               bool                `json:"requestSuccess"`
	WebsiteData           WebsiteData         `json:"websiteData"`
	Availability          WebsiteAvailability `json:"availability"`
	LastCheckResult       WebsiteCheckResult  `json:"lastCheckResult"`
	LastFailedCheckResult WebsiteCheckResult  `json:"lastFailedCheckResult"`
}

// Contains a success-bool and an array of WebsiteCheckResults.
type ResultsResponse struct {
	Success  bool                 `json:"requestSuccess"`
	Websites []WebsiteCheckResult `json:"results"`
}

// Contains the Website's basic data such as name, protocol, url and current status.
type BasicWebsite struct {
	Name         string `json:"name"`
	Protocol     string `json:"protocol"`
	Url          string `json:"url"`
	Status       string `json:"status"`
	ResponseTime string `json:"responseTime"`
}

// Contains the Website's basic data such as id, name and url.
type WebsiteData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Contains the Website's availability data like the amount of ups / downs, total checks and the average availability.
type WebsiteAvailability struct {
	Ups     int    `json:"ups"`
	Downs   int    `json:"downs"`
	Total   int    `json:"total"`
	Average string `json:"average"`
}

// Contains the Website's latest check result.
type WebsiteCheckResult struct {
	Status       string `json:"status"`
	ResponseTime string `json:"responseTime"`
	Time         string `json:"time"`
}

// *******************
// * ADMIN-RESPONSES *
// *******************

// Contains a success-bool and an array of AdminWebsites.
type DetailedWebsiteResponse struct {
	Success  bool              `json:"requestSuccess"`
	Websites []DetailedWebsite `json:"websites"`
}

// Contains a success-bool and a notification-object.
type WebsiteNotificationsResponse struct {
	Success       bool          `json:"requestSuccess"`
	Notifications Notifications `json:"notifications"`
}

// Contains the Website's data, which will be shown inside the admin-backend.
type DetailedWebsite struct {
	Id                   int                  `json:"id"`
	Name                 string               `json:"name"`
	Enabled              bool                 `json:"enabled"`
	Visible              bool                 `json:"visible"`
	Protocol             string               `json:"protocol"`
	Url                  string               `json:"url"`
	CheckMethod          string               `json:"checkMethod"`
	Status               string               `json:"status"`
	ResponseTime         string               `json:"responseTime"`
	Time                 string               `json:"time"`
	EnabledNotifications EnabledNotifications `json:"notifications"`
}

// Contains all saved notification settings of a website.
type Notifications struct {
	PushbulletKey string `json:"pushbulletKey"`
	Email         string `json:"email"`
}

// Contains whether a notification-type is enabled or not.
type EnabledNotifications struct {
	Pushbullet bool `json:"pushbullet"`
	Email      bool `json:"email"`
}
