package token

import "time"


type Maker interface{
	// Create token from username
	CreateToken(username string,duration time.Duration) (string, error)

	// Verify Token
	VerifyToken(token string) (*Payload,error)
}
