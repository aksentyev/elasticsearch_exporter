package exporter

type Config struct {
    URL             string
    Labels          map[string]string
    ExporterOptions map[string]string
    CacheTTL        int
}
