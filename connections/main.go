package main

import (
	plugin_api "github.com/k8spacket/plugin-api"
	"github.com/perithompson/k8spacket-plugins/connections/metrics"
	"github.com/perithompson/k8spacket-plugins/connections/metrics/connections"
)

type stream plugin_api.ReassembledStream

func (s stream) InitPlugin(manager plugin_api.PluginManager) {
	manager.RegisterPlugin(s)
	manager.RegisterHttpHandler("/connections/connections", connections.ConnectionHandler)
	manager.RegisterHttpHandler("/connections/api/health", connections.Health)
}

func (s stream) DistributeReassembledStream(reassembledStream plugin_api.ReassembledStream) {
	metrics.StoreconnectionsMetric(reassembledStream)
}

func (s stream) DistributeTCPPacketPayload(_ plugin_api.TCPPacketPayload) {
	//silent
}

func init() {}

// exported
var StreamPlugin stream
