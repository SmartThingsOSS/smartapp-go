package handlers

import (
	"github.com/SmartThingsOSS/smartapp-go/examples/c2c-connector/internal/logging"
	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappcore"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/client/devices"
	st "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
	log "github.com/sirupsen/logrus"
)

type InstallHandler struct {
	Api smartappcore.SmartThingsApi
}

func NewInstallHandler(api smartappcore.SmartThingsApi) *InstallHandler {
	return &InstallHandler{
		Api: api,
	}
}

func (h *InstallHandler) Handler() smartappcore.Handler {
	return func(params *smartappcore.SmartAppParams) (*models.ExecutionResponse, error) {
		ctx := params.Context
		request := params.Request
		installedAppId := request.InstallData.InstalledApp.InstalledAppID.String()
		locationId := request.InstallData.InstalledApp.LocationID
		authToken := request.InstallData.AuthToken
		deviceProfileId := params.GetSetting("DEVICE_PROFILE_ID")
		logger := logging.Logger(ctx)

		device, err := h.Api.InstallDevice(&devices.InstallDeviceParams{
			Authorization: *authToken,
			Context:       ctx,
			InstallationRequest: &st.DeviceInstallRequest{
				Label:      "C2C Virtual Switch",
				LocationID: &locationId,
				App: &st.DeviceInstallRequestApp{
					ProfileID:      &deviceProfileId,
					InstalledAppID: &installedAppId,
					ExternalID:     "someExternalId",
				},
			},
		})

		if err != nil {
			return &models.ExecutionResponse{}, err
		}

		logger.WithFields(log.Fields{
			"deviceId": device.DeviceID,
		}).Debug("Successfully installed device.")

		if err := h.Api.CreateDeviceEvents(&devices.CreateDeviceEventsParams{
			Authorization: *authToken,
			Context:       ctx,
			DeviceID:      device.DeviceID,
			DeviceEventRequest: &st.DeviceEventsRequest{
				DeviceEvents: []*st.DeviceStateEvent{
					{
						Component:  "main",
						Capability: "switch",
						Attribute:  "switch",
						Value:      "off",
					},
				},
			},
		}); err != nil {
			return &models.ExecutionResponse{}, err
		}

		logger.WithFields(log.Fields{
			"deviceId": device.DeviceID,
		}).Debug("Successfully set device state.")

		return &models.ExecutionResponse{StatusCode: 200, InstallData: map[string]string{}}, nil
	}
}
