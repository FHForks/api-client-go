package client

import (
	"fmt"
	"github.com/firehydrant/api-client-go/client/incidents"
	"github.com/firehydrant/api-client-go/fhclient"
	"github.com/firehydrant/api-client-go/models"
	"os"
)

func Start() (
	*models.IncidentEntity,
	error,
) {
	c := fhclient.NewApiClient(
		fhclient.Config{
			ApiHost: os.Getenv("API_HOST"),
			ApiKey:  os.Getenv("API_KEY"),
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
	fmt.Println(
		"incident payload: ",
		incident.Payload,
	)
	return incident.Payload, nil
}
