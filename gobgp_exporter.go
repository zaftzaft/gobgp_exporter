package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	//"net/url"
	//"regexp"
	//"strconv"
	//"strings"
	//"sync"
	//"time"
	"os/exec"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	namespace = "gobgp"
)

var (
	tableAccepted = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "table_accepted"),
		"accepted BGP route.",
		nil, nil,
	)
)

type Exporter struct {
}

func NewExporter() (*Exporter, error) {
	return &Exporter{}, nil
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- tableAccepted
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	out, err := exec.Command("/usr/local/bin/gobgp", "-j", "neighbor").Output()
	if err != nil {
		fmt.Errorf("%v", err)
	}
	bytes := []byte(out)
	var neighbors []GobgpNeighbor

	err = json.Unmarshal(bytes, &neighbors)

	ch <- prometheus.MustNewConstMetric(
		tableAccepted, prometheus.GaugeValue, float64(neighbors[0].AfiSafis[0].State.Accepted),
	)

}

func init() {
	prometheus.MustRegister(version.NewCollector("gobgp_exporter"))
}

func main() {
	var (
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9510").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("gobgp_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	exporter, err := NewExporter()

	if err != nil {
		log.Fatalln(err)
	}

	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`gobgp_exporter`))
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}
