package clients

type Service2Client struct {
	BaseURL string
}

func NewService2Client(baseURL string) *Service2Client {
	return &Service2Client{BaseURL: baseURL}
}

func (c *Service2Client) GetOrderByUserId(userId string) ([]any, error) {
	return []any{}, nil
}
