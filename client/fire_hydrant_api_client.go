// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/client/alerts"
	"github.com/firehydrant/api-client-go/client/attachments"
	"github.com/firehydrant/api-client-go/client/change_types"
	"github.com/firehydrant/api-client-go/client/changes"
	"github.com/firehydrant/api-client-go/client/environments"
	"github.com/firehydrant/api-client-go/client/events"
	"github.com/firehydrant/api-client-go/client/functionalities"
	"github.com/firehydrant/api-client-go/client/incident_roles"
	"github.com/firehydrant/api-client-go/client/incident_tags"
	"github.com/firehydrant/api-client-go/client/incident_types"
	"github.com/firehydrant/api-client-go/client/incidents"
	"github.com/firehydrant/api-client-go/client/infrastructures"
	"github.com/firehydrant/api-client-go/client/integrations"
	"github.com/firehydrant/api-client-go/client/metrics"
	"github.com/firehydrant/api-client-go/client/noauth"
	"github.com/firehydrant/api-client-go/client/notes"
	"github.com/firehydrant/api-client-go/client/nunc"
	"github.com/firehydrant/api-client-go/client/nunc_connections"
	"github.com/firehydrant/api-client-go/client/onboarding"
	"github.com/firehydrant/api-client-go/client/organizations"
	"github.com/firehydrant/api-client-go/client/ping"
	"github.com/firehydrant/api-client-go/client/post_mortems"
	"github.com/firehydrant/api-client-go/client/release_notes"
	"github.com/firehydrant/api-client-go/client/reports"
	"github.com/firehydrant/api-client-go/client/runbook_templates"
	"github.com/firehydrant/api-client-go/client/runbooks"
	"github.com/firehydrant/api-client-go/client/saved_searches"
	"github.com/firehydrant/api-client-go/client/scheduled_maintenances"
	"github.com/firehydrant/api-client-go/client/schedules"
	"github.com/firehydrant/api-client-go/client/services"
	"github.com/firehydrant/api-client-go/client/severities"
	"github.com/firehydrant/api-client-go/client/severity_matrix"
	"github.com/firehydrant/api-client-go/client/teams"
	"github.com/firehydrant/api-client-go/client/ticketing"
	"github.com/firehydrant/api-client-go/client/users"
	"github.com/firehydrant/api-client-go/client/webhooks"
)

// Default fire hydrant API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.firehydrant.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new fire hydrant API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *FireHydrantAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new fire hydrant API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *FireHydrantAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new fire hydrant API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *FireHydrantAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(FireHydrantAPI)
	cli.Transport = transport
	cli.Alerts = alerts.New(transport, formats)
	cli.Attachments = attachments.New(transport, formats)
	cli.ChangeTypes = change_types.New(transport, formats)
	cli.Changes = changes.New(transport, formats)
	cli.Environments = environments.New(transport, formats)
	cli.Events = events.New(transport, formats)
	cli.Functionalities = functionalities.New(transport, formats)
	cli.IncidentRoles = incident_roles.New(transport, formats)
	cli.IncidentTags = incident_tags.New(transport, formats)
	cli.IncidentTypes = incident_types.New(transport, formats)
	cli.Incidents = incidents.New(transport, formats)
	cli.Infrastructures = infrastructures.New(transport, formats)
	cli.Integrations = integrations.New(transport, formats)
	cli.Metrics = metrics.New(transport, formats)
	cli.Noauth = noauth.New(transport, formats)
	cli.Notes = notes.New(transport, formats)
	cli.Nunc = nunc.New(transport, formats)
	cli.NuncConnections = nunc_connections.New(transport, formats)
	cli.Onboarding = onboarding.New(transport, formats)
	cli.Organizations = organizations.New(transport, formats)
	cli.Ping = ping.New(transport, formats)
	cli.PostMortems = post_mortems.New(transport, formats)
	cli.ReleaseNotes = release_notes.New(transport, formats)
	cli.Reports = reports.New(transport, formats)
	cli.RunbookTemplates = runbook_templates.New(transport, formats)
	cli.Runbooks = runbooks.New(transport, formats)
	cli.SavedSearches = saved_searches.New(transport, formats)
	cli.ScheduledMaintenances = scheduled_maintenances.New(transport, formats)
	cli.Schedules = schedules.New(transport, formats)
	cli.Services = services.New(transport, formats)
	cli.Severities = severities.New(transport, formats)
	cli.SeverityMatrix = severity_matrix.New(transport, formats)
	cli.Teams = teams.New(transport, formats)
	cli.Ticketing = ticketing.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Webhooks = webhooks.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// FireHydrantAPI is a client for fire hydrant API
type FireHydrantAPI struct {
	Alerts alerts.ClientService

	Attachments attachments.ClientService

	ChangeTypes change_types.ClientService

	Changes changes.ClientService

	Environments environments.ClientService

	Events events.ClientService

	Functionalities functionalities.ClientService

	IncidentRoles incident_roles.ClientService

	IncidentTags incident_tags.ClientService

	IncidentTypes incident_types.ClientService

	Incidents incidents.ClientService

	Infrastructures infrastructures.ClientService

	Integrations integrations.ClientService

	Metrics metrics.ClientService

	Noauth noauth.ClientService

	Notes notes.ClientService

	Nunc nunc.ClientService

	NuncConnections nunc_connections.ClientService

	Onboarding onboarding.ClientService

	Organizations organizations.ClientService

	Ping ping.ClientService

	PostMortems post_mortems.ClientService

	ReleaseNotes release_notes.ClientService

	Reports reports.ClientService

	RunbookTemplates runbook_templates.ClientService

	Runbooks runbooks.ClientService

	SavedSearches saved_searches.ClientService

	ScheduledMaintenances scheduled_maintenances.ClientService

	Schedules schedules.ClientService

	Services services.ClientService

	Severities severities.ClientService

	SeverityMatrix severity_matrix.ClientService

	Teams teams.ClientService

	Ticketing ticketing.ClientService

	Users users.ClientService

	Webhooks webhooks.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *FireHydrantAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Alerts.SetTransport(transport)
	c.Attachments.SetTransport(transport)
	c.ChangeTypes.SetTransport(transport)
	c.Changes.SetTransport(transport)
	c.Environments.SetTransport(transport)
	c.Events.SetTransport(transport)
	c.Functionalities.SetTransport(transport)
	c.IncidentRoles.SetTransport(transport)
	c.IncidentTags.SetTransport(transport)
	c.IncidentTypes.SetTransport(transport)
	c.Incidents.SetTransport(transport)
	c.Infrastructures.SetTransport(transport)
	c.Integrations.SetTransport(transport)
	c.Metrics.SetTransport(transport)
	c.Noauth.SetTransport(transport)
	c.Notes.SetTransport(transport)
	c.Nunc.SetTransport(transport)
	c.NuncConnections.SetTransport(transport)
	c.Onboarding.SetTransport(transport)
	c.Organizations.SetTransport(transport)
	c.Ping.SetTransport(transport)
	c.PostMortems.SetTransport(transport)
	c.ReleaseNotes.SetTransport(transport)
	c.Reports.SetTransport(transport)
	c.RunbookTemplates.SetTransport(transport)
	c.Runbooks.SetTransport(transport)
	c.SavedSearches.SetTransport(transport)
	c.ScheduledMaintenances.SetTransport(transport)
	c.Schedules.SetTransport(transport)
	c.Services.SetTransport(transport)
	c.Severities.SetTransport(transport)
	c.SeverityMatrix.SetTransport(transport)
	c.Teams.SetTransport(transport)
	c.Ticketing.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Webhooks.SetTransport(transport)
}
