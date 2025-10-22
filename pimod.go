package pimod

import (
	"time"

	"github.com/elgopher/pi/piaudio"
)

func Decode(data []byte) Module {
	mod, err := DecodeOrErr(data)
	if err != nil {
		panic(err)
	}
	return mod
}

func DecodeOrErr(data []byte) (Module, error) {
	return Module{}, nil
}

type Module struct {
}

func (m Module) Play(ch piaudio.Chan, patternNumber int, fade time.Duration) {

}

func (m Module) Stop(fade time.Duration) {

}

func CurrentMod() (Module, bool) {
	return Module{}, false
}
