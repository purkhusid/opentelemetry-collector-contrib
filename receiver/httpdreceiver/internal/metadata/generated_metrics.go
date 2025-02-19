// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "httpdreceiver"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	HttpdCurrentConnections MetricIntf
	HttpdRequests           MetricIntf
	HttpdScoreboard         MetricIntf
	HttpdTraffic            MetricIntf
	HttpdUptime             MetricIntf
	HttpdWorkers            MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"httpd.current_connections",
		"httpd.requests",
		"httpd.scoreboard",
		"httpd.traffic",
		"httpd.uptime",
		"httpd.workers",
	}
}

var metricsByName = map[string]MetricIntf{
	"httpd.current_connections": Metrics.HttpdCurrentConnections,
	"httpd.requests":            Metrics.HttpdRequests,
	"httpd.scoreboard":          Metrics.HttpdScoreboard,
	"httpd.traffic":             Metrics.HttpdTraffic,
	"httpd.uptime":              Metrics.HttpdUptime,
	"httpd.workers":             Metrics.HttpdWorkers,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"httpd.current_connections",
		func(metric pdata.Metric) {
			metric.SetName("httpd.current_connections")
			metric.SetDescription("The number of active connections currently attached to the HTTP server")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"httpd.requests",
		func(metric pdata.Metric) {
			metric.SetName("httpd.requests")
			metric.SetDescription("The number of requests serviced by the HTTP server per second")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"httpd.scoreboard",
		func(metric pdata.Metric) {
			metric.SetName("httpd.scoreboard")
			metric.SetDescription("The number of connections in each state")
			metric.SetUnit("scoreboard")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"httpd.traffic",
		func(metric pdata.Metric) {
			metric.SetName("httpd.traffic")
			metric.SetDescription("Total HTTP server traffic")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"httpd.uptime",
		func(metric pdata.Metric) {
			metric.SetName("httpd.uptime")
			metric.SetDescription("The amount of time that the server has been running in seconds")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"httpd.workers",
		func(metric pdata.Metric) {
			metric.SetName("httpd.workers")
			metric.SetDescription("The number of workers currently attached to the HTTP server")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
	// ScoreboardState (The state of a connection)
	ScoreboardState string
	// ServerName (The name of the Apache HTTP server)
	ServerName string
	// WorkersState (The state of workers)
	WorkersState string
}{
	"state",
	"server_name",
	"state",
}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels

// LabelScoreboardState are the possible values that the label "scoreboard_state" can have.
var LabelScoreboardState = struct {
	Open        string
	Waiting     string
	Starting    string
	Reading     string
	Sending     string
	Keepalive   string
	Dnslookup   string
	Closing     string
	Logging     string
	Finishing   string
	IdleCleanup string
}{
	"open",
	"waiting",
	"starting",
	"reading",
	"sending",
	"keepalive",
	"dnslookup",
	"closing",
	"logging",
	"finishing",
	"idle_cleanup",
}

// LabelWorkersState are the possible values that the label "workers_state" can have.
var LabelWorkersState = struct {
	Busy string
	Idle string
}{
	"busy",
	"idle",
}
