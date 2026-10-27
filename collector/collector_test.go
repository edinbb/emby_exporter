package collector

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/williamclot/emby_exporter/emby"
)

func TestEmbyCollector(t *testing.T) {
	var tests = []struct {
		desc    string
		ss      *testStatusSource
		matches []*regexp.Regexp
	}{
		{
			desc: "empty",
			ss: &testStatusSource{
				s: &emby.Status{},
			},
			matches: []*regexp.Regexp{
				regexp.MustCompile(`emby_server_info{os="",server="",version=""} 1`),
				regexp.MustCompile(`emby_server_update_available{server=""} 0`),
				regexp.MustCompile(`emby_server_pending_restart{server=""} 0`),
			},
		},
		{
			desc: "full",
			ss: &testStatusSource{
				s: &emby.Status{
					ServerName:         "mars",
					Version:            "1.1.2",
					OperatingSystem:    "Linux",
					HasUpdateAvailable: true,
					HasPendingRestart:  false,
					DeviceCount:        2,
					MovieCount:         123,
					SeriesCount:        12,
					EpisodeCount:       362,
					UserCount:          2,
					FailedLoginCount: map[string]int{
						"bill": 0,
						"joe":  2,
					},
					SessionCountByPlayMethod: map[string]int{
						"DirectPlay":   0,
						"DirectStream": 1,
						"Transcode":    1,
					},
				},
			},

			matches: []*regexp.Regexp{
				regexp.MustCompile(`emby_server_info{os="Linux",server="mars",version="1\.1\.2"} 1`),
				regexp.MustCompile(`emby_server_update_available{server="mars"} 1`),
				regexp.MustCompile(`emby_server_pending_restart{server="mars"} 0`),
				regexp.MustCompile(`emby_device_count{server="mars"} 2`),
				regexp.MustCompile(`emby_movie_count{server="mars"} 123`),
				regexp.MustCompile(`emby_series_count{server="mars"} 12`),
				regexp.MustCompile(`emby_episode_count{server="mars"} 362`),
				regexp.MustCompile(`emby_user_count{server="mars"} 2`),
				regexp.MustCompile(`emby_failed_login{server="mars",user="bill"} 0`),
				regexp.MustCompile(`emby_failed_login{server="mars",user="joe"} 2`),
				regexp.MustCompile(`emby_session_count{play_method="DirectPlay",server="mars"} 0`),
				regexp.MustCompile(`emby_session_count{play_method="DirectStream",server="mars"} 1`),
				regexp.MustCompile(`emby_session_count{play_method="Transcode",server="mars"} 1`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			out := testCollector(t, New(tt.ss))
			t.Log(string(out))
			for _, m := range tt.matches {
				if !m.Match(out) {
					t.Fatalf("output failed to match regex (regexp: %v)", m)
				}
			}
		})
	}
}

func testCollector(t *testing.T, collector prometheus.Collector) []byte {
	t.Helper()

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(collector)

	srv := httptest.NewServer(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	defer srv.Close()

	c := &http.Client{Timeout: 1 * time.Second}
	resp, err := c.Get(srv.URL)
	if err != nil {
		t.Fatalf("failed to HTTP GET data from prometheus: %v", err)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read server response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected HTTP status %d: %s", resp.StatusCode, buf)
	}

	return buf
}

var _ StatusSource = &testStatusSource{}

type testStatusSource struct {
	s *emby.Status
}

func (ss *testStatusSource) Status() (*emby.Status, error) {
	return ss.s, nil
}
