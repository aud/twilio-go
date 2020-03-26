package twilio

// Client.Functions.CreateService(
// )

// Client.Functions.CreateEnvironment(
// )

// Client.Functions.CreateFunction(
// )

// Client.Functions.UploadFunction(
// )

const ServicePathPart = "Services"

type ServicesService struct {
	client *Client
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

func (s *ServicesService) Create(ctx context.Context, data url.Values) (*Service, error) {
	service := new(Service)

	err := s.client.UpdateResource(ctx, ServicePathPart, data)

	return service, err
}
