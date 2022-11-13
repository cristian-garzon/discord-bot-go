package discord

import "errors"

type Config struct {
	Token  string
	Prefix string
}

func NewConfig(token string, prefix string) (*Config, error) {

	if token == "" {
		return nil, errors.New("error getting token")
	}

	if prefix == "" {
		return nil, errors.New("error getting prefix")
	}

	return &Config{Token: token}, nil
}
