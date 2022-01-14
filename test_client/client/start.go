package client

import (
	"fmt"
	"github.com/firehydrant/api-client-go/client/incidents"
	"github.com/firehydrant/api-client-go/fhclient"
	"github.com/firehydrant/api-client-go/models"
)

func Start() (
	*models.IncidentEntity,
	error,
) {
	c := fhclient.NewApiClient(
		fhclient.Config{
			ApiHost: "gmendez.ngrok.io",
			ApiKey:  "<local_bot_token>",
		},
	)
	var name string
	name = "test incident gaby"
	params := incidents.NewPostV1IncidentsParams()
	params.V1Incidents = &models.PostV1Incidents{
		Name: &name,
	}

	incident, err := c.Client.Incidents.PostV1Incidents(
		params,
		c.Auth,
	)
	if err != nil {
		fmt.Println(
			"did this error",
			err,
		)
		return nil, err
	}
	return incident.Payload, nil
}
