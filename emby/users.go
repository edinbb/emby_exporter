package emby

import (
	"fmt"
	"net/http"
)

type UserList struct {
	TotalCount int    `json:"TotalRecordCount"`
	Users      []User `json:"Items"`
}

type User struct {
	Name                    string            `json:"Name"`
	ServerId                string            `json:"ServerId"`
	Prefix                  string            `json:"Prefix,omitempty"`
	ConnectUserName         string            `json:"ConnectUserName,omitempty"`
	ConnectLinkType         string            `json:"ConnectLinkType,omitempty"`
	DateCreated             string            `json:"DateCreated"`
	Id                      string            `json:"Id"`
	PrimaryImageTag         string            `json:"PrimaryImageTag,omitempty"`
	PrimaryImageAspectRatio float64           `json:"PrimaryImageAspectRatio,omitempty"`
	HasPassword             bool              `json:"HasPassword"`
	HasConfiguredPassword   bool              `json:"HasConfiguredPassword"`
	LastLoginDate           string            `json:"LastLoginDate,omitempty"`
	LastActivityDate        string            `json:"LastActivityDate,omitempty"`
	Configuration           UserConfiguration `json:"Configuration"`
	Policy                  UserPolicy        `json:"Policy"`
}

type UserConfiguration struct {
	PlayDefaultAudioTrack      bool     `json:"PlayDefaultAudioTrack"`
	AudioLanguagePreference    string   `json:"AudioLanguagePreference,omitempty"`
	SubtitleLanguagePreference string   `json:"SubtitleLanguagePreference"`
	DisplayMissingEpisodes     bool     `json:"DisplayMissingEpisodes"`
	SubtitleMode               string   `json:"SubtitleMode"`
	OrderedViews               []string `json:"OrderedViews"`
	LatestItemsExcludes        []string `json:"LatestItemsExcludes"`
	MyMediaExcludes            []string `json:"MyMediaExcludes"`
	HidePlayedInLatest         bool     `json:"HidePlayedInLatest"`
	HidePlayedInMoreLikeThis   bool     `json:"HidePlayedInMoreLikeThis"`
	HidePlayedInSuggestions    bool     `json:"HidePlayedInSuggestions"`
	RememberAudioSelections    bool     `json:"RememberAudioSelections"`
	RememberSubtitleSelections bool     `json:"RememberSubtitleSelections"`
	EnableNextEpisodeAutoPlay  bool     `json:"EnableNextEpisodeAutoPlay"`
	ResumeRewindSeconds        int      `json:"ResumeRewindSeconds"`
	IntroSkipMode              string   `json:"IntroSkipMode"`
	EnableLocalPassword        bool     `json:"EnableLocalPassword"`
}

type UserPolicy struct {
	IsAdministrator                  bool     `json:"IsAdministrator"`
	IsHidden                         bool     `json:"IsHidden"`
	IsHiddenRemotely                 bool     `json:"IsHiddenRemotely"`
	IsHiddenFromUnusedDevices        bool     `json:"IsHiddenFromUnusedDevices"`
	IsDisabled                       bool     `json:"IsDisabled"`
	LockedOutDate                    int64    `json:"LockedOutDate"`
	MaxParentalRating                *int     `json:"MaxParentalRating,omitempty"`
	AllowTagOrRating                 bool     `json:"AllowTagOrRating"`
	BlockedTags                      []string `json:"BlockedTags"`
	IsTagBlockingModeInclusive       bool     `json:"IsTagBlockingModeInclusive"`
	IncludeTags                      []string `json:"IncludeTags"`
	EnableUserPreferenceAccess       bool     `json:"EnableUserPreferenceAccess"`
	BlockUnratedItems                []string `json:"BlockUnratedItems"`
	EnableRemoteControlOfOtherUsers  bool     `json:"EnableRemoteControlOfOtherUsers"`
	EnableSharedDeviceControl        bool     `json:"EnableSharedDeviceControl"`
	EnableRemoteAccess               bool     `json:"EnableRemoteAccess"`
	EnableLiveTvManagement           bool     `json:"EnableLiveTvManagement"`
	EnableLiveTvAccess               bool     `json:"EnableLiveTvAccess"`
	EnableMediaPlayback              bool     `json:"EnableMediaPlayback"`
	EnableAudioPlaybackTranscoding   bool     `json:"EnableAudioPlaybackTranscoding"`
	EnableVideoPlaybackTranscoding   bool     `json:"EnableVideoPlaybackTranscoding"`
	EnableTranscodingQuality         bool     `json:"EnableTranscodingQuality"`
	AutoRemoteQuality                int      `json:"AutoRemoteQuality"`
	EnablePlaybackRemuxing           bool     `json:"EnablePlaybackRemuxing"`
	EnableContentDeletion            bool     `json:"EnableContentDeletion"`
	RestrictedFeatures               []string `json:"RestrictedFeatures"`
	EnableContentDeletionFromFolders []string `json:"EnableContentDeletionFromFolders"`
	EnableContentDownloading         bool     `json:"EnableContentDownloading"`
	EnableSubtitleDownloading        bool     `json:"EnableSubtitleDownloading"`
	EnableSubtitleManagement         bool     `json:"EnableSubtitleManagement"`
	EnableSyncTranscoding            bool     `json:"EnableSyncTranscoding"`
	EnableMediaConversion            bool     `json:"EnableMediaConversion"`
	EnabledChannels                  []string `json:"EnabledChannels"`
	EnableAllChannels                bool     `json:"EnableAllChannels"`
	EnabledFolders                   []string `json:"EnabledFolders"`
	EnableAllFolders                 bool     `json:"EnableAllFolders"`
	InvalidLoginAttemptCount         int      `json:"InvalidLoginAttemptCount"`
	EnablePublicSharing              bool     `json:"EnablePublicSharing"`
	RemoteClientBitrateLimit         int      `json:"RemoteClientBitrateLimit"`
	AuthenticationProviderId         string   `json:"AuthenticationProviderId"`
	ExcludedSubFolders               []string `json:"ExcludedSubFolders"`
	SimultaneousStreamLimit          int      `json:"SimultaneousStreamLimit"`
	EnabledDevices                   []string `json:"EnabledDevices"`
	EnableAllDevices                 bool     `json:"EnableAllDevices"`
	AllowCameraUpload                bool     `json:"AllowCameraUpload"`
	AllowSharingPersonalItems        bool     `json:"AllowSharingPersonalItems"`
}

func (c *Client) GetUsers() (*UserList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Users/Query", c.URL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	res := UserList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
