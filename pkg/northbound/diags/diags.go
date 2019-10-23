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

// Package diags implements the diagnostic gRPC service for the configuration subsystem.
package diags

import (
	"fmt"
	"github.com/onosproject/onos-config/pkg/manager"
	"github.com/onosproject/onos-config/pkg/northbound"
	"github.com/onosproject/onos-config/pkg/northbound/admin"
	"github.com/onosproject/onos-config/pkg/store"
	"github.com/onosproject/onos-config/pkg/store/change/network"
	streams "github.com/onosproject/onos-config/pkg/store/stream"
	devicechangetypes "github.com/onosproject/onos-config/pkg/types/change/device"
	networkchangetypes "github.com/onosproject/onos-config/pkg/types/change/network"
	"github.com/onosproject/onos-config/pkg/utils"
	devicetopo "github.com/onosproject/onos-topo/pkg/northbound/device"
	"google.golang.org/grpc"
	log "k8s.io/klog"
)

// Service is a Service implementation for administration.
type Service struct {
	northbound.Service
}

// OpStateDiagsClientFactory : Default OpStateDiagsClient creation.
var OpStateDiagsClientFactory = func(cc *grpc.ClientConn) OpStateDiagsClient {
	return NewOpStateDiagsClient(cc)
}

// CreateOpStateDiagsClient creates and returns a new op state diags client
func CreateOpStateDiagsClient(cc *grpc.ClientConn) OpStateDiagsClient {
	return OpStateDiagsClientFactory(cc)
}

// ConfigDiagsClientFactory : Default OpStateDiagsClient creation.
var ConfigDiagsClientFactory = func(cc *grpc.ClientConn) ConfigDiagsClient {
	return NewConfigDiagsClient(cc)
}

// CreateConfigDiagsClient creates and returns a new op state diags client
func CreateConfigDiagsClient(cc *grpc.ClientConn) ConfigDiagsClient {
	return ConfigDiagsClientFactory(cc)
}

// ChangeServiceClientFactory : Default ChangeServiceClient creation.
var ChangeServiceClientFactory = func(cc *grpc.ClientConn) ChangeServiceClient {
	return NewChangeServiceClient(cc)
}

// CreateChangeServiceClient creates and returns a new change service client
func CreateChangeServiceClient(cc *grpc.ClientConn) ChangeServiceClient {
	return ChangeServiceClientFactory(cc)
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	RegisterConfigDiagsServer(r, Server{})
	RegisterOpStateDiagsServer(r, Server{})
	RegisterChangeServiceServer(r, Server{})
}

// Server implements the gRPC service for diagnostic facilities.
type Server struct {
}

// GetChanges provides a stream of submitted network changes.
// Deprecated GetChanges is an NBI method of getting Changes
// to be replaced by GetDeviceChanges and GetNetworkChanges
func (s Server) GetChanges(r *ChangesRequest, stream ConfigDiags_GetChangesServer) error {
	for _, c := range manager.GetManager().ChangeStore.Store {
		if len(r.ChangeIDs) > 0 && !stringInList(r.ChangeIDs, store.B64(c.ID)) {
			continue
		}

		errInvalid := c.IsValid()
		if errInvalid != nil {
			return errInvalid
		}

		// Build a change message
		msg := &admin.Change{
			Time:         &c.Created,
			Id:           store.B64(c.ID),
			Desc:         c.Description,
			ChangeValues: c.Config,
		}
		err := stream.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetConfigurations provides a stream of submitted network changes.
// Deprecated GetConfigurations is an NBI method of getting Configurations
// to be replaced by GetDeviceChanges and GetNetworkChanges
func (s Server) GetConfigurations(r *ConfigRequest, stream ConfigDiags_GetConfigurationsServer) error {
	for _, c := range manager.GetManager().ConfigStore.Store {
		if len(r.DeviceIDs) > 0 && !stringInList(r.DeviceIDs, c.Device) {
			continue
		}

		changeIDs := make([]string, len(c.Changes))

		for idx, cid := range c.Changes {
			changeIDs[idx] = store.B64(cid)
		}

		msg := &Configuration{
			Name:       string(c.Name),
			DeviceID:   c.Device,
			Version:    c.Version,
			DeviceType: c.Type,
			Updated:    &c.Updated,
			ChangeIDs:  changeIDs,
		}
		err := stream.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetOpState provides a stream of Operational and State data
func (s Server) GetOpState(r *OpStateRequest, stream OpStateDiags_GetOpStateServer) error {
	deviceCache, ok := manager.GetManager().OperationalStateCache[devicetopo.ID(r.DeviceId)]
	if !ok {
		return fmt.Errorf("no Operational State cache available for %s", r.DeviceId)
	}

	for path, value := range deviceCache {
		pathValue := &devicechangetypes.PathValue{
			Path:  path,
			Value: value,
		}

		msg := &OpStateResponse{Type: admin.Type_NONE, Pathvalue: pathValue}
		err := stream.Send(msg)
		if err != nil {
			return err
		}
	}

	if r.Subscribe {
		streamID := fmt.Sprintf("diags-%p", stream)
		listener, err := manager.GetManager().Dispatcher.RegisterOpState(streamID)
		if err != nil {
			log.Warning("Failed setting up a listener for OpState events on ", r.DeviceId)
			return err
		}
		log.Infof("NBI Diags OpState started on %s for %s", streamID, r.DeviceId)
		for {
			select {
			case opStateEvent := <-listener:
				if opStateEvent.Subject() != r.DeviceId {
					// If the event is not for this device then ignore it
					continue
				}
				log.Infof("Event received NBI Diags OpState subscribe channel %s for %s",
					streamID, r.DeviceId)

				pathValue := &devicechangetypes.PathValue{
					Path:  opStateEvent.Path(),
					Value: opStateEvent.Value(),
				}

				msg := &OpStateResponse{Type: admin.Type_ADDED, Pathvalue: pathValue}
				err = stream.SendMsg(msg)
				if err != nil {
					log.Warningf("Error sending message on stream %s. Closing. %v",
						streamID, msg)
					return err
				}
			case <-stream.Context().Done():
				manager.GetManager().Dispatcher.UnregisterOperationalState(streamID)
				log.Infof("NBI Diags OpState subscribe channel %s for %s closed",
					streamID, r.DeviceId)
				return nil
			}
		}
	}

	log.Infof("Closing NBI Diags OpState stream (no subscribe) for %s", r.DeviceId)
	return nil
}

// ListNetworkChanges provides a stream of Network Changes
// If the optional `subscribe` flag is true, then get then return the list of
// changes first, and then hold the connection open and send on
// further updates until the client hangs up
func (s Server) ListNetworkChanges(r *ListNetworkChangeRequest, stream ChangeService_ListNetworkChangesServer) error {
	log.Infof("ListNetworkChanges called with %s. Subscribe %v", r.ChangeID, r.Subscribe)

	// There may be a wildcard given - we only want to reply with changes that match
	matcher := utils.MatchWildcardChNameRegexp(string(r.ChangeID))

	if r.Subscribe {
		eventCh := make(chan streams.Event)
		ctx, err := manager.GetManager().NetworkChangesStore.Watch(eventCh, network.WithReplay())
		if err != nil {
			log.Errorf("Error watching Network Changes %s", err)
			return err
		}
		defer ctx.Close()

		for {
			breakout := false
			select { // Blocks until one of the following are received
			case event, ok := <-eventCh:
				if !ok { // Will happen at the end of stream
					breakout = true
					break
				}

				change := event.Object.(*networkchangetypes.NetworkChange)
				log.Infof("List() channel sent us %v", change)

				if matcher.MatchString(string(change.ID)) {
					msg := &ListNetworkChangeResponse{
						Change: change,
					}
					log.Infof("Sending matching change %v", change.ID)
					err := stream.Send(msg)
					if err != nil {
						log.Errorf("Error sending NetworkChanges %v %v", change.ID, err)
						return err
					}
				}
			case <-stream.Context().Done():
				log.Infof("Remote client closed connection")
				return nil
			}
			if breakout {
				break
			}
		}
	} else {
		changeCh := make(chan *networkchangetypes.NetworkChange)
		ctx, err := manager.GetManager().NetworkChangesStore.List(changeCh)
		if err != nil {
			log.Errorf("Error listing Network Changes %s", err)
			return err
		}
		defer ctx.Close()

		for {
			breakout := false
			select { // Blocks until one of the following are received
			case change, ok := <-changeCh:
				if !ok { // Will happen at the end of stream
					breakout = true
					break
				}

				log.Infof("List() channel sent us %v", change)
				if matcher.MatchString(string(change.ID)) {
					msg := &ListNetworkChangeResponse{
						Change: change,
					}
					log.Infof("Sending matching change %v", change.ID)
					err := stream.Send(msg)
					if err != nil {
						log.Errorf("Error sending NetworkChanges %v %v", change.ID, err)
						return err
					}
				}
			case <-stream.Context().Done():
				log.Infof("Remote client closed connection")
				return nil
			}
			if breakout {
				break
			}
		}
	}
	log.Infof("Closing ListNetworkChanges for %s", r.ChangeID)
	return nil
}

// ListDeviceChanges provides a stream of Device Changes
func (s Server) ListDeviceChanges(r *ListDeviceChangeRequest, stream ChangeService_ListDeviceChangesServer) error {
	log.Infof("ListDeviceChanges called with %s. Subscribe %v", r.ChangeID, r.Subscribe)

	if r.Subscribe {
		eventCh := make(chan streams.Event)
		ctx, err := manager.GetManager().DeviceChangesStore.Watch(devicetopo.ID(r.ChangeID), eventCh)
		if err != nil {
			log.Errorf("Error watching Network Changes %s", err)
			return err
		}
		defer ctx.Close()

		for {
			breakout := false
			select { // Blocks until one of the following are received
			case event, ok := <-eventCh:
				if !ok { // Will happen at the end of stream
					breakout = true
					break
				}

				change := event.Object.(*devicechangetypes.DeviceChange)
				log.Infof("List() channel sent us %v", change)

				msg := &ListDeviceChangeResponse{
					Change: change,
				}
				log.Infof("Sending matching change %v", change.ID)
				err := stream.Send(msg)
				if err != nil {
					log.Errorf("Error sending NetworkChanges %v %v", change.ID, err)
					return err
				}
			case <-stream.Context().Done():
				log.Infof("Remote client closed connection")
				return nil
			}
			if breakout {
				break
			}
		}
	} else {
		changeCh := make(chan *devicechangetypes.DeviceChange)
		ctx, err := manager.GetManager().DeviceChangesStore.List(devicetopo.ID(r.ChangeID), changeCh)
		if err != nil {
			log.Errorf("Error listing Network Changes %s", err)
			return err
		}
		defer ctx.Close()

		for {
			breakout := false
			select { // Blocks until one of the following are received
			case change, ok := <-changeCh:
				if !ok { // Will happen at the end of stream
					breakout = true
					break
				}

				log.Infof("List() channel sent us %v", change)
				msg := &ListDeviceChangeResponse{
					Change: change,
				}
				log.Infof("Sending matching change %v", change.ID)
				err := stream.Send(msg)
				if err != nil {
					log.Errorf("Error sending NetworkChanges %v %v", change.ID, err)
					return err
				}
			case <-stream.Context().Done():
				log.Infof("Remote client closed connection")
				return nil
			}
			if breakout {
				break
			}
		}
	}
	log.Infof("Closing ListDeviceChanges for %s", r.ChangeID)
	return nil
}

func stringInList(haystack []string, needle string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}
	return false
}
