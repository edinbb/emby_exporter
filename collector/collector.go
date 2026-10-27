package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/williamclot/emby_exporter/emby"
)

var _ StatusSource = &emby.Client{}

type StatusSource interface {
	Status() (*emby.Status, error)
}

type Collector struct {
	ServerInfo      *prometheus.Desc
	UpdateAvailable *prometheus.Desc
	PendingRestart  *prometheus.Desc

	DeviceCount      *prometheus.Desc
	MovieCount       *prometheus.Desc
	SeriesCount      *prometheus.Desc
	EpisodeCount     *prometheus.Desc
	UserCount        *prometheus.Desc
	FailedLoginCount *prometheus.Desc
	SessionCount     *prometheus.Desc

	ss StatusSource
}

func boolToFloat(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func New(ss StatusSource) prometheus.Collector {
	labels := []string{"server"}
	namespace := "emby"

	return &Collector{
		ServerInfo: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "server", "info"),
			"Metadata about a given Emby server.",
			[]string{"server", "version", "os"},
			nil,
		),

		UpdateAvailable: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "server", "update_available"),
			"Whether an Emby update is available (1 = yes, 0 = no).",
			labels,
			nil,
		),

		PendingRestart: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "server", "pending_restart"),
			"Whether Emby has a pending restart (1 = yes, 0 = no).",
			labels,
			nil,
		),

		DeviceCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "device", "count"),
			"Number of devices configured in Emby.",
			labels,
			nil,
		),

		MovieCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "movie", "count"),
			"Number of movies available in Emby.",
			labels,
			nil,
		),

		SeriesCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "series", "count"),
			"Number of tv shows available in Emby.",
			labels,
			nil,
		),

		EpisodeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "episode", "count"),
			"Number of tv show episodes available in Emby.",
			labels,
			nil,
		),

		UserCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "user", "count"),
			"Number of users configured in Emby.",
			labels,
			nil,
		),

		FailedLoginCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "failed", "login"),
			"Failed login counts per user in Emby.",
			[]string{"server", "user"},
			nil,
		),

		SessionCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "session", "count"),
			"Number of active playback sessions by play method.",
			[]string{"server", "play_method"},
			nil,
		),

		ss: ss,
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ds := []*prometheus.Desc{
		c.ServerInfo,
		c.UpdateAvailable,
		c.PendingRestart,
		c.DeviceCount,
		c.MovieCount,
		c.SeriesCount,
		c.EpisodeCount,
		c.UserCount,
		c.FailedLoginCount,
		c.SessionCount,
	}
	for _, d := range ds {
		ch <- d
	}
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	s, err := c.ss.Status()
	if err != nil {
		log.Printf("failed collecting emby metrics: %v", err)
		ch <- prometheus.NewInvalidMetric(c.ServerInfo, err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.ServerInfo,
		prometheus.GaugeValue,
		1,
		s.ServerName, s.Version, s.OperatingSystem,
	)

	ch <- prometheus.MustNewConstMetric(
		c.UpdateAvailable,
		prometheus.GaugeValue,
		boolToFloat(s.HasUpdateAvailable),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.PendingRestart,
		prometheus.GaugeValue,
		boolToFloat(s.HasPendingRestart),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.DeviceCount,
		prometheus.GaugeValue,
		float64(s.DeviceCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.MovieCount,
		prometheus.GaugeValue,
		float64(s.MovieCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.SeriesCount,
		prometheus.GaugeValue,
		float64(s.SeriesCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.EpisodeCount,
		prometheus.GaugeValue,
		float64(s.EpisodeCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.UserCount,
		prometheus.GaugeValue,
		float64(s.UserCount),
		s.ServerName,
	)

	for k, v := range s.FailedLoginCount {
		ch <- prometheus.MustNewConstMetric(
			c.FailedLoginCount,
			prometheus.GaugeValue,
			float64(v),
			s.ServerName, k,
		)
	}

	for method, n := range s.SessionCountByPlayMethod {
		ch <- prometheus.MustNewConstMetric(
			c.SessionCount,
			prometheus.GaugeValue,
			float64(n),
			s.ServerName, method,
		)
	}
}
