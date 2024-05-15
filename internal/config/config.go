package config

import "github.com/subosito/gotenv"

func Load(path string) error {
	err := gotenv.Load(path)

	if err != nil {
		return err
	}

	return nil
}

type GRPCConfig interface {
	Address() string
}
