package metrics

import (
	"os"
	"strconv"

	plugin_api "github.com/k8spacket/plugin-api"
	"github.com/perithompson/k8spacket-plugins/connections/metrics/connections"
	"github.com/perithompson/k8spacket-plugins/connections/metrics/prometheus"
)

func StoreconnectionsMetric(stream plugin_api.ReassembledStream) {
	hideSrcPort, _ := strconv.ParseBool(os.Getenv("K8S_PACKET_HIDE_SRC_PORT"))
	var srcPortMetrics = stream.SrcPort
	if hideSrcPort {
		srcPortMetrics = "dynamic"
	}

	prometheus.K8sPacketBytesSentMetric.WithLabelValues(stream.SrcNamespace, stream.Src, stream.SrcName, srcPortMetrics, stream.Dst, stream.DstName, stream.DstPort, strconv.FormatBool(stream.Closed)).Observe(stream.BytesSent)
	prometheus.K8sPacketBytesReceivedMetric.WithLabelValues(stream.SrcNamespace, stream.Src, stream.SrcName, srcPortMetrics, stream.Dst, stream.DstName, stream.DstPort, strconv.FormatBool(stream.Closed)).Observe(stream.BytesReceived)
	prometheus.K8sPacketDurationSecondsMetric.WithLabelValues(stream.SrcNamespace, stream.Src, stream.SrcName, srcPortMetrics, stream.Dst, stream.DstName, stream.DstPort, strconv.FormatBool(stream.Closed)).Observe(stream.Duration)

	connections.Updateconnections(stream.Src, stream.SrcName, stream.SrcNamespace, stream.Dst, stream.DstName, stream.DstPort, stream.DstNamespace, stream.Closed, stream.BytesSent, stream.BytesReceived, stream.Duration)
}
