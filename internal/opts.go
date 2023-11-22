package opts

import (
	"encoding/json"
	"log"
)

type Params struct {
	Iface string `           long:"interface"      env:"DB_SERVER"  description:"Server network interface" default:"eno1"`
}

func (o *Params) GetJson() []byte {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return jsonBytes
}
