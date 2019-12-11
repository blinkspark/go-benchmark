package config

import "fmt"

// Config *********************************************
type Config struct {
	Name string
	Version
	SHA2
}

// NewConfig creates a v0 config
func NewConfig() *Config {
	return &Config{
		Name:    "V0",
		Version: Version{0, 0, 1},
		SHA2:    SHA2{1000, 1 << 20},
	}
}

// ****************************************************

// Version data ***************************************
type Version struct {
	Major, Minor, Patch int
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// ****************************************************

// SHA2 ***********************************************
type SHA2 struct {
	Round, DataSize int
}

// ****************************************************
