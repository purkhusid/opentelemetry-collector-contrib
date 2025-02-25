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

package system

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/model/pdata"
	conventions "go.opentelemetry.io/collector/model/semconv/v1.5.0"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

const (
	// TypeStr is type of detector.
	TypeStr = "system"
)

var hostnameSourcesMap = map[string]func(*Detector) (string, error){
	"dns": getFQDN,
	"os":  getHostname,
}

var _ internal.Detector = (*Detector)(nil)

// Detector is a system metadata detector
type Detector struct {
	provider        systemMetadata
	logger          *zap.Logger
	hostnameSources []string
}

// NewDetector creates a new system metadata detector
func NewDetector(p component.ProcessorCreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	if len(cfg.HostnameSources) == 0 {
		cfg.HostnameSources = []string{"dns", "os"}
	}
	return &Detector{provider: &systemMetadataImpl{}, logger: p.Logger, hostnameSources: cfg.HostnameSources}, nil
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(_ context.Context) (resource pdata.Resource, schemaURL string, err error) {
	var hostname string

	res := pdata.NewResource()
	attrs := res.Attributes()

	osType, err := d.provider.OSType()
	if err != nil {
		return res, "", fmt.Errorf("failed getting OS type: %w", err)
	}
	for _, source := range d.hostnameSources {
		getHostFromSource := hostnameSourcesMap[source]
		hostname, err = getHostFromSource(d)
		if err == nil {
			attrs.InsertString(conventions.AttributeHostName, hostname)
			attrs.InsertString(conventions.AttributeOSType, osType)

			return res, conventions.SchemaURL, nil
		}
		d.logger.Debug(err.Error())
	}

	return res, "", errors.New("all hostname sources are failed to get hostname")
}

// getHostname returns OS hostname
func getHostname(d *Detector) (string, error) {
	hostname, err := d.provider.Hostname()
	if err != nil {
		return "", fmt.Errorf("failed getting OS hostname: %w", err)
	}
	return hostname, nil
}

// getFQDN returns FQDN of the host
func getFQDN(d *Detector) (string, error) {
	hostname, err := d.provider.FQDN()
	if err != nil {
		return "", fmt.Errorf("failed getting FQDN: %w", err)
	}
	return hostname, nil
}
