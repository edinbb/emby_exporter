package emby

type Status struct {
	ServerName      string
	Version         string
	OperatingSystem string

	HasUpdateAvailable bool
	HasPendingRestart  bool

	DeviceCount int

	MovieCount   int
	SeriesCount  int
	EpisodeCount int

	UserCount                int
	FailedLoginCount         map[string]int
	SessionCountByPlayMethod map[string]int
}

func (c *Client) Status() (*Status, error) {
	s := &Status{
		FailedLoginCount: make(map[string]int),
		SessionCountByPlayMethod: map[string]int{
			"DirectPlay":   0,
			"DirectStream": 0,
			"Transcode":    0,
		},
	}

	sys, err := c.GetSystemInfo()
	if err != nil {
		return nil, err
	}

	s.ServerName = sys.ServerName
	s.Version = sys.Version
	s.OperatingSystem = sys.OperatingSystem

	s.HasUpdateAvailable = sys.HasUpdateAvailable
	s.HasPendingRestart = sys.HasPendingRestart

	deviceList, err := c.GetDevices()
	if err != nil {
		return nil, err
	}

	s.DeviceCount = len(deviceList.Devices)

	library, err := c.GetLibraryItems()
	if err != nil {
		return nil, err
	}

	s.MovieCount = library.MovieCount
	s.SeriesCount = library.SeriesCount
	s.EpisodeCount = library.EpisodeCount

	userList, err := c.GetUsers()
	if err != nil {
		return nil, err
	}

	s.UserCount = userList.TotalCount
	for _, user := range userList.Users {
		s.FailedLoginCount[user.Name] = user.Policy.InvalidLoginAttemptCount
	}

	sessions, err := c.GetSessions()
	if err != nil {
		return nil, err
	}

	for _, session := range *sessions {
		if session.NowPlayingItem == nil {
			continue
		}

		s.SessionCountByPlayMethod[session.PlayState.PlayMethod]++
	}

	return s, nil
}
