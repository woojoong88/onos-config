// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manager

import (
	"fmt"
	"github.com/onosproject/onos-config/pkg/events"
	"github.com/onosproject/onos-config/pkg/store"
	"github.com/onosproject/onos-config/pkg/store/change"
	"log"
	"time"
)

// SetNetworkConfig sets the given value, according to the path on the configuration for the specified targed
func (m *Manager) SetNetworkConfig(target string, configName string, updates map[string]string,
	removes []string) (change.ID, error) {
	if _, ok := m.deviceStore.Store[target]; !ok {
		return nil, fmt.Errorf("Device not present %s", target)
	}
	//checks if config exists, otherwise create new
	deviceConfig, ok := m.ConfigStore.Store[configName];
	if !ok {
		deviceConfig = store.Configuration{
			Name:    configName,
			Device:  target,
			Created: time.Now(),
			Updated: time.Now(),
			//TODO make these usable value
			User:        "foo",
			Description: "test",
			Changes:     make([]change.ID, 0),
		}
	}

	var newChange = make([]*change.Value, 0)
	//updates
	for path, value := range updates {
		changeValue, _ := change.CreateChangeValue(path, value, false)
		newChange = append(newChange, changeValue)
	}

	//removes
	for _, path := range removes {
		changeValue, _ := change.CreateChangeValue(path, "", true)
		newChange = append(newChange, changeValue)
	}

	//TODO gNMI --> no description we need to come up with something
	configChange, err := change.CreateChange(newChange, "bar")
	if err != nil {
		return nil, err
	}
	m.ChangeStore.Store[store.B64(configChange.ID)] = configChange
	log.Println("Added change", store.B64(configChange.ID), "to ChangeStore (in memory)")

	deviceConfig.Changes = append(deviceConfig.Changes, configChange.ID)
	deviceConfig.Updated = time.Now()
	m.ConfigStore.Store[configName] = deviceConfig
	eventValues := make(map[string]string)
	eventValues[events.ChangeID] = store.B64(configChange.ID)
	eventValues[events.Committed] = "true"
	m.changesChannel <- events.CreateEvent(deviceConfig.Device,
		events.EventTypeConfiguration, eventValues)
	return configChange.ID, nil
}
