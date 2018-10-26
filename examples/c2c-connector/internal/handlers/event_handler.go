package handlers

import (
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappcore"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/client/devices"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

type EventHandler struct {
	Api smartappcore.SmartThingsApi
}

func NewEventHandler(api smartappcore.SmartThingsApi) *EventHandler {
	return &EventHandler{
		Api: api,
	}
}

func (h *EventHandler) Handler() smartappcore.Handler {
	handler := smartappcore.NewEventHandler()
	handler.AddEventHandler(&DeviceCommandsHandler{
		Api: h.Api,
	})
	return handler.BuildHandler()
}

type DeviceCommandsHandler struct {
	Api smartappcore.SmartThingsApi
}

func (h *DeviceCommandsHandler) Test(event *smartapp.Event) bool {
	return event != nil && event.EventType == smartapp.EventTypeDEVICECOMMANDSEVENT
}

func (h *DeviceCommandsHandler) Handle(params *smartappcore.SmartAppParams, event *smartapp.Event) error {
	authToken := *params.Request.EventData.AuthToken
	cmd := event.DeviceCommandsEvent.Commands[0]
	err := h.Api.CreateDeviceEvents(&devices.CreateDeviceEventsParams{
		Authorization: authToken,
		Context:       params.Context,
		DeviceID:      event.DeviceCommandsEvent.DeviceID,
		DeviceEventRequest: &models.DeviceEventsRequest{
			DeviceEvents: []*models.DeviceStateEvent{
				{
					Component:  cmd.ComponentID,
					Capability: cmd.Capability,
					Attribute:  "switch",
					Value:      cmd.Command,
				},
			},
		},
	})

	if err != nil {
		return err
	}
	return nil
}
