package stats

import (
	"github.com/perithompson/k8spacket-plugins/connections/metrics/connections/model"
	"github.com/perithompson/k8spacket-plugins/connections/metrics/connections/stats/bytes"
	"github.com/perithompson/k8spacket-plugins/connections/metrics/connections/stats/connection"
	"github.com/perithompson/k8spacket-plugins/connections/metrics/connections/stats/duration"
)

func GetConfig(statsType string) model.Config {
	switch statsType {
	case "bytes":
		return bytes.GetConfig()
	case "duration":
		return duration.GetConfig()
	default:
		return connection.GetConfig()
	}
}

func FillNodeStats(statsType string, node *model.Node, connEndpoint model.ConnectionEndpoint) {
	switch statsType {
	case "bytes":
		bytes.FillNodeStats(node, connEndpoint)
	case "duration":
		duration.FillNodeStats(node, connEndpoint)
	default:
		connection.FillNodeStats(node, connEndpoint)
	}
}

func FillEdgeStats(statsType string, edge *model.Edge, connItem model.ConnectionItem) {
	switch statsType {
	case "bytes":
		bytes.FillEdgeStats(edge, connItem)
	case "duration":
		duration.FillEdgeStats(edge, connItem)
	default:
		connection.FillEdgeStats(edge, connItem)
	}
}
