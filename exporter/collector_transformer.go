package exporter

import (
    "github.com/aksentyev/hubble/exportertools"
    "regexp"
    "fmt"
)

/*
Transform json data to exportertools.Metric list
*/
func (c *Collector) Transform(allStats *NodeStatsResponse) (metrics []*exportertools.Metric) {
    for _, stats := range allStats.Nodes {
        // GC Stats
        for _, gcstats := range stats.JVM.GC.Collectors {
            metrics = append(metrics, c.ConvertToMetric("jvm_gc_collection_seconds_count",
                                                        float64(gcstats.CollectionCount),
                                                        "COUNTER",
                                                        nil))

            metrics = append(metrics, c.ConvertToMetric("jvm_gc_collection_seconds_sum",
                                                        float64(gcstats.CollectionTime / 1000),
                                                        "COUNTER",
                                                        nil))
        }

        // Breaker stats
        for _, bstats := range stats.Breakers {
            metrics = append(metrics, c.ConvertToMetric("breakers_estimated_size_bytes",
                                                        float64(bstats.EstimatedSize),
                                                        "GAUGE",
                                                        nil))

            metrics = append(metrics, c.ConvertToMetric("breakers_limit_size_bytes",
                                                        float64(bstats.LimitSize),
                                                        "GAUGE",
                                                        nil))
        }

        // Thread Pool stats
        for pool, pstats := range stats.ThreadPool {
            metrics = append(metrics, c.ConvertToMetric("thread_pool_completed_count",
                                                        float64(pstats.Completed),
                                                        "COUNTER",
                                                        map[string]string{"type": pool}))

            metrics = append(metrics, c.ConvertToMetric("thread_pool_rejected_count",
                                                        float64(pstats.Rejected),
                                                        "COUNTER",
                                                        map[string]string{"type": pool}))

            metrics = append(metrics, c.ConvertToMetric("thread_pool_active_count",
                                                        float64(pstats.Active),
                                                        "GAUGE",
                                                        map[string]string{"type": pool}))

            metrics = append(metrics, c.ConvertToMetric("thread_pool_threads_count",
                                                        float64(pstats.Threads),
                                                        "GAUGE",
                                                        map[string]string{"type": pool}))

            metrics = append(metrics, c.ConvertToMetric("thread_pool_largest_count",
                                                        float64(pstats.Largest),
                                                        "GAUGE",
                                                        map[string]string{"type": pool}))

            metrics = append(metrics, c.ConvertToMetric("thread_pool_queue_count",
                                                        float64(pstats.Queue),
                                                        "GAUGE",
                                                        map[string]string{"type": pool}))
        }

        // JVM Memory Stats
        metrics = append(metrics, c.ConvertToMetric("jvm_memory_committed_bytes",
                                                    float64(stats.JVM.Mem.HeapCommitted),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("jvm_memory_used_bytes",
                                                    float64(stats.JVM.Mem.HeapUsed),
                                                    "GAUGE",
                                                    nil))


        metrics = append(metrics, c.ConvertToMetric("jvm_memory_max_bytes",
                                                    float64(stats.JVM.Mem.HeapMax),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("jvm_memory_committed_bytes",
                                                    float64(stats.JVM.Mem.NonHeapCommitted),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("jvm_memory_used_bytes",
                                                    float64(stats.JVM.Mem.NonHeapUsed),
                                                    "GAUGE",
                                                    nil))

        // Indices Stats)
        metrics = append(metrics, c.ConvertToMetric("indices_fielddata_memory_size_bytes",
                                                    float64(stats.Indices.FieldData.MemorySize),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_fielddata_evictions",
                                                    float64(stats.Indices.FieldData.Evictions),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_filter_cache_memory_size_bytes",
                                                    float64(stats.Indices.FilterCache.MemorySize),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_filter_cache_evictions",
                                                    float64(stats.Indices.FilterCache.Evictions),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_query_cache_memory_size_bytes",
                                                    float64(stats.Indices.QueryCache.MemorySize),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_query_cache_evictions",
                                                    float64(stats.Indices.QueryCache.Evictions),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_request_cache_memory_size_bytes",
                                                    float64(stats.Indices.QueryCache.MemorySize),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_request_cache_evictions",
                                                    float64(stats.Indices.QueryCache.Evictions),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_docs",
                                                    float64(stats.Indices.Docs.Count),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_docs_deleted",
                                                    float64(stats.Indices.Docs.Deleted),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_segments_memory_bytes",
                                                    float64(stats.Indices.Segments.Memory),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_segments_count",
                                                    float64(stats.Indices.Segments.Count),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_store_size_bytes",
                                                    float64(stats.Indices.Store.Size),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_store_throttle_time_ms_total",
                                                    float64(stats.Indices.Store.ThrottleTime),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_flush_total",
                                                    float64(stats.Indices.Flush.Total),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_flush_time_ms_total",
                                                    float64(stats.Indices.Flush.Time),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_indexing_index_time_ms_total",
                                                    float64(stats.Indices.Indexing.IndexTime),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_indexing_index_total",
                                                    float64(stats.Indices.Indexing.IndexTotal),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_merges_total_time_ms_total",
                                                    float64(stats.Indices.Merges.TotalTime),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_merges_total_size_bytes_total",
                                                    float64(stats.Indices.Merges.TotalSize),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_merges_total",
                                                    float64(stats.Indices.Merges.Total),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_refresh_total_time_ms_total",
                                                    float64(stats.Indices.Refresh.TotalTime),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("indices_refresh_total",
                                                    float64(stats.Indices.Refresh.Total),
                                                    "COUNTER",
                                                    nil))

        // Transport Stats)
        metrics = append(metrics, c.ConvertToMetric("transport_rx_packets_total",
                                                    float64(stats.Transport.RxCount),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("transport_rx_size_bytes_total",
                                                    float64(stats.Transport.RxSize),
                                                    "COUNTER",
                                                    nil))


        metrics = append(metrics, c.ConvertToMetric("transport_tx_packets_total",
                                                    float64(stats.Transport.TxCount),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("transport_tx_size_bytes_total",
                                                    float64(stats.Transport.TxSize),
                                                    "COUNTER",
                                                    nil))

        // Process Stats)
        metrics = append(metrics, c.ConvertToMetric("process_cpu_percent",
                                                    float64(stats.Process.CPU.Percent),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("process_mem_resident_size_bytes",
                                                    float64(stats.Process.Memory.Resident),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("process_mem_share_size_bytes",
                                                    float64(stats.Process.Memory.Share),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("process_mem_virtual_size_bytes",
                                                    float64(stats.Process.Memory.TotalVirtual),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("process_open_files_count",
                                                    float64(stats.Process.OpenFD),
                                                    "GAUGE",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("process_cpu_time_seconds_sum",
                                                    float64(stats.Process.CPU.Total / 1000),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("process_cpu_time_seconds_sum",
                                                    float64(stats.Process.CPU.Sys / 1000),
                                                    "COUNTER",
                                                    nil))

        metrics = append(metrics, c.ConvertToMetric("process_cpu_time_seconds_sum",
                                                    float64(stats.Process.CPU.User / 1000),
                                                    "COUNTER",
                                                    nil))

    }

    return metrics
}

func (c *Collector) ConvertToMetric(name string, v float64, t string, l interface{}) *exportertools.Metric {
    processedName := name
    switch l.(type) {
    case map[string]string:
        re := regexp.MustCompile("(.*)_count")
        match := re.FindStringSubmatch(name)[1]
        processedName = fmt.Sprintf("%v_%v_count", match, l.(map[string]string)["type"])
    }

    var desc string
    switch t {
    case "GAUGE":
        desc = gaugeMetrics[name]
    case "COUNTER":
        desc = counterMetrics[name]
    }

    m := exportertools.Metric {
        Name:        processedName,
        Description: desc,
        Type:        exportertools.StringToType(t),
        Value:       v,
        Labels:      c.Labels,
    }

    return &m
}

// hubble/exportertools does not support variable metric, so the workaround above is used
// func (c *Collector) ConvertToMetric(name string, v float64, t string, l interface{}) *exportertools.Metric {
//     labels := map[string]string{}
//     for k,v := range c.Labels {
//         labels[k] = v
//     }
//
//     switch l.(type) {
//     case map[string]string:
//         for k, v := range l.(map[string]string) {
//             labels[k] = v
//         }
//     }
//
//     var desc string
//     switch t {
//     case "GAUGE":
//         desc = gaugeMetrics[name]
//     case "COUNTER":
//         desc = counterMetrics[name]
//     }
//
//     m := exportertools.Metric {
//         Name:        name,
//         Description: desc,
//         Type:        exportertools.StringToType(t),
//         Value:       v,
//         Labels:      labels,
//     }
//
//     return &m
// }
