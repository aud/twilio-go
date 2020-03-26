package twilio

import (
	"context"
	"net/url"
)

const ServicePathPart = "Services"
const EnvironmentPathPart = "Environments"

type ServicesService struct {
	client *Client
}

type EnvironmentsService struct {
	client *Client
}

type FunctionService struct {
	Services     *ServicesService
	Envrionments *EnvironmentsService
}

type Service struct {
	Sid                string     `json:"sid"`
	AccountSid         string     `json:"account_sid"`
	FriendlyName       string     `json:"friendly_name"`
	UniqueName         string     `json:"unique_name"`
	IncludeCredentials bool       `json:"include_credentials"`
	UiEditable         bool       `json:"ui_editable"`
	DateCreated        TwilioTime `json:"date_created"`
	DateUpdated        TwilioTime `json:"date_updated"`
	URL                string     `json:"url"`
}

// Todo implement rest
// https://www.twilio.com/docs/runtime/functions-assets-api/quickstart?code-sample=code-create-an-environment&code-language=Ruby&code-sdk-version=5.x
type Environment struct {
	Sid string `json:"sid"`
}

func (s *ServicesService) Create(ctx context.Context, data url.Values) (*Service, error) {
	service := new(Service)

	err := s.client.CreateResource(ctx, ServicePathPart, data, service)

	return service, err
}

func (s *EnvrionmentsService) Create(ctx context.Context, serviceSid string, data url.Values) (*Environment, error) {
	environment := new(Environment)

	err := s.client.CreateResource(ctx, ServicePathPart+"/"+serviceSid+EnvironmentPathPart, data, environment)

	return environment, err
}
