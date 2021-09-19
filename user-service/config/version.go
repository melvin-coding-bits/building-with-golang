package config

//Version of the api endpoints
type Verison string

func (v Verison) String() string {
	return string(v)
}

const (
	//V1 is the first version of the api server
	V1 Verison = "v1"
)
