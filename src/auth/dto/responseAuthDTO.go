package authDTO

import "time"

type ResponseAuthDTO struct {
	AuthToken string      `json:"auth_token"`
	ExpiresIn time.Time   `json:"expires_in"`
	User      interface{} `json:"user"`
}
