package exporter

import (
    "github.com/aksentyev/hubble/exportertools"

    "encoding/json"
    "io/ioutil"
    "net/http"

    "errors"
    "fmt"
)

type Collector struct {
    *Config
}

func NewCollector(config *Config) *Collector {
    return &Collector{config}
}
// Collecting metrics
func (c *Collector) Collect() ([]*exportertools.Metric, error) {
    var allStats NodeStatsResponse

    resp, err := http.Get(c.URL)
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    err = json.Unmarshal(body, &allStats)

    if err != nil {
        err = errors.New(fmt.Sprintf("Collecting stat for %v failed: %v", c.URL, err))
    }

    metrics := c.Transform(&allStats)

    return metrics, err
}
