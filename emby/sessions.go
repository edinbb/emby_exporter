package emby

import (
	"fmt"
	"net/http"
)

type SessionInfo struct {
	PlayState             PlayState        `json:"PlayState"`
	AdditionalUsers       []AdditionalUser `json:"AdditionalUsers"`
	RemoteEndPoint        string           `json:"RemoteEndPoint"`
	Protocol              string           `json:"Protocol"`
	PlayableMediaTypes    []string         `json:"PlayableMediaTypes"`
	PlaylistIndex         int              `json:"PlaylistIndex"`
	PlaylistLength        int              `json:"PlaylistLength"`
	Id                    string           `json:"Id"`
	ServerId              string           `json:"ServerId"`
	UserId                string           `json:"UserId,omitempty"`
	UserName              string           `json:"UserName,omitempty"`
	UserPrimaryImageTag   string           `json:"UserPrimaryImageTag,omitempty"`
	Client                string           `json:"Client"`
	LastActivityDate      string           `json:"LastActivityDate"`
	DeviceName            string           `json:"DeviceName"`
	NowPlayingItem        *NowPlayingItem  `json:"NowPlayingItem,omitempty"`
	InternalDeviceId      int              `json:"InternalDeviceId"`
	DeviceId              string           `json:"DeviceId"`
	ApplicationVersion    string           `json:"ApplicationVersion"`
	AppIconUrl            string           `json:"AppIconUrl,omitempty"`
	SupportedCommands     []string         `json:"SupportedCommands"`
	SupportsRemoteControl bool             `json:"SupportsRemoteControl"`
}

type AdditionalUser struct {
	UserId   string `json:"UserId"`
	UserName string `json:"UserName"`
}

type PlayState struct {
	PositionTicks       *int64  `json:"PositionTicks,omitempty"`
	CanSeek             bool    `json:"CanSeek"`
	IsPaused            bool    `json:"IsPaused"`
	IsMuted             bool    `json:"IsMuted"`
	VolumeLevel         *int    `json:"VolumeLevel,omitempty"`
	AudioStreamIndex    *int    `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex *int    `json:"SubtitleStreamIndex,omitempty"`
	MediaSourceId       string  `json:"MediaSourceId,omitempty"`
	PlayMethod          string  `json:"PlayMethod,omitempty"`
	RepeatMode          string  `json:"RepeatMode"`
	SleepTimerMode      string  `json:"SleepTimerMode"`
	SubtitleOffset      int     `json:"SubtitleOffset"`
	Shuffle             bool    `json:"Shuffle"`
	PlaybackRate        float64 `json:"PlaybackRate"`
}

type NowPlayingItem struct {
	Name                    string            `json:"Name"`
	OriginalTitle           string            `json:"OriginalTitle,omitempty"`
	ServerId                string            `json:"ServerId"`
	Id                      string            `json:"Id"`
	DateCreated             string            `json:"DateCreated"`
	DateModified            string            `json:"DateModified"`
	PresentationUniqueKey   string            `json:"PresentationUniqueKey,omitempty"`
	Container               string            `json:"Container,omitempty"`
	PremiereDate            *string           `json:"PremiereDate,omitempty"`
	ExternalUrls            []ExternalURL     `json:"ExternalUrls,omitempty"`
	Path                    string            `json:"Path,omitempty"`
	OfficialRating          string            `json:"OfficialRating,omitempty"`
	Overview                string            `json:"Overview,omitempty"`
	Taglines                []string          `json:"Taglines,omitempty"`
	Genres                  []string          `json:"Genres,omitempty"`
	CommunityRating         float64           `json:"CommunityRating,omitempty"`
	RunTimeTicks            int64             `json:"RunTimeTicks,omitempty"`
	Size                    int64             `json:"Size,omitempty"`
	FileName                string            `json:"FileName,omitempty"`
	Bitrate                 int64             `json:"Bitrate,omitempty"`
	ProductionYear          int               `json:"ProductionYear,omitempty"`
	ProviderIds             map[string]string `json:"ProviderIds,omitempty"`
	IsFolder                bool              `json:"IsFolder"`
	ParentId                string            `json:"ParentId,omitempty"`
	Type                    string            `json:"Type"`
	Studios                 []NameID          `json:"Studios,omitempty"`
	GenreItems              []NameID          `json:"GenreItems,omitempty"`
	PrimaryImageAspectRatio float64           `json:"PrimaryImageAspectRatio,omitempty"`
	MediaStreams            []MediaStream     `json:"MediaStreams,omitempty"`
	ImageTags               map[string]string `json:"ImageTags,omitempty"`
	BackdropImageTags       []string          `json:"BackdropImageTags,omitempty"`
	Chapters                []Chapter         `json:"Chapters,omitempty"`
	MediaType               string            `json:"MediaType,omitempty"`
	Width                   int               `json:"Width,omitempty"`
	Height                  int               `json:"Height,omitempty"`
	SeriesName              string            `json:"SeriesName,omitempty"`
	SeriesId                string            `json:"SeriesId,omitempty"`
	SeasonName              string            `json:"SeasonName,omitempty"`
	SeasonId                string            `json:"SeasonId,omitempty"`
	IndexNumber             *int              `json:"IndexNumber,omitempty"`
	ParentIndexNumber       *int              `json:"ParentIndexNumber,omitempty"`
	Album                   string            `json:"Album,omitempty"`
	AlbumArtist             string            `json:"AlbumArtist,omitempty"`
	Artists                 []string          `json:"Artists,omitempty"`
}

type ExternalURL struct {
	Name string `json:"Name"`
	Url  string `json:"Url"`
}

type NameID struct {
	Name string `json:"Name"`
	Id   int    `json:"Id"`
}

type MediaStream struct {
	Codec                           string  `json:"Codec"`
	Language                        string  `json:"Language,omitempty"`
	TimeBase                        string  `json:"TimeBase,omitempty"`
	Title                           string  `json:"Title,omitempty"`
	VideoRange                      string  `json:"VideoRange,omitempty"`
	DisplayTitle                    string  `json:"DisplayTitle"`
	DisplayLanguage                 string  `json:"DisplayLanguage,omitempty"`
	IsInterlaced                    bool    `json:"IsInterlaced"`
	ChannelLayout                   string  `json:"ChannelLayout,omitempty"`
	BitRate                         int     `json:"BitRate,omitempty"`
	BitDepth                        int     `json:"BitDepth,omitempty"`
	RefFrames                       int     `json:"RefFrames,omitempty"`
	Channels                        int     `json:"Channels,omitempty"`
	SampleRate                      int     `json:"SampleRate,omitempty"`
	IsDefault                       bool    `json:"IsDefault"`
	IsForced                        bool    `json:"IsForced"`
	IsHearingImpaired               bool    `json:"IsHearingImpaired"`
	Height                          int     `json:"Height,omitempty"`
	Width                           int     `json:"Width,omitempty"`
	AverageFrameRate                float64 `json:"AverageFrameRate,omitempty"`
	RealFrameRate                   float64 `json:"RealFrameRate,omitempty"`
	Profile                         string  `json:"Profile,omitempty"`
	Type                            string  `json:"Type"`
	AspectRatio                     string  `json:"AspectRatio,omitempty"`
	Index                           int     `json:"Index"`
	IsExternal                      bool    `json:"IsExternal"`
	IsTextSubtitleStream            bool    `json:"IsTextSubtitleStream"`
	SupportsExternalStream          bool    `json:"SupportsExternalStream"`
	Path                            string  `json:"Path,omitempty"`
	Protocol                        string  `json:"Protocol,omitempty"`
	PixelFormat                     string  `json:"PixelFormat,omitempty"`
	Level                           float64 `json:"Level,omitempty"`
	IsAnamorphic                    bool    `json:"IsAnamorphic,omitempty"`
	ExtendedVideoType               string  `json:"ExtendedVideoType,omitempty"`
	ExtendedVideoSubType            string  `json:"ExtendedVideoSubType,omitempty"`
	ExtendedVideoSubTypeDescription string  `json:"ExtendedVideoSubTypeDescription,omitempty"`
	AttachmentSize                  int     `json:"AttachmentSize"`
	SubtitleLocationType            string  `json:"SubtitleLocationType,omitempty"`
}

type Chapter struct {
	StartPositionTicks int64  `json:"StartPositionTicks"`
	Name               string `json:"Name"`
	MarkerType         string `json:"MarkerType"`
	ChapterIndex       int    `json:"ChapterIndex"`
}

func (c *Client) GetSessions() (*[]SessionInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Sessions", c.URL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	res := []SessionInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
