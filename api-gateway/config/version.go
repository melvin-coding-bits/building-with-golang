package config

type Verison string

func (v Verison) String() string {
	return string(v)
}

const (
	V1 Verison = "v1"
)
