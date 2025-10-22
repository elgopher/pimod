package pimod

import (
	"fmt"
	"slices"
	"time"

	"github.com/elgopher/pi/piaudio"
	"github.com/elgopher/pimod/internal"
)

func Decode(data []byte) Module {
	mod, err := DecodeOrErr(data)
	if err != nil {
		panic(err)
	}
	return mod
}

func DecodeOrErr(data []byte) (Module, error) {
	title := internal.CStr(data[0:20])

	tag := internal.CStr(data[1080:1085])

	supportedTags := []string{"M.K.", "M!K!", "FLT4", "4CHN", "CHN4", "TDZ4"}

	if !slices.Contains(supportedTags, tag) {
		return Module{}, fmt.Errorf("%s is not a supported tag", title)
	}

	fmt.Println(tag)

	return Module{
		Title: title,
		Tag:   tag,
	}, nil
}

type Module struct {
	Title    string
	Tag      string // aka signature, for example "M.K." for ProTracker, 4 channels
	Samples  []Sample
	Patterns []Pattern
}

type Sample struct {
	Name     string
	Data     []byte
	FineTune uint8
	Volume   uint8
}

type Pattern struct {
}

func (m Module) Play(ch piaudio.Chan, patternNumber int, fade time.Duration) {

}

// TODO Add Pause/unpause methods too?
func (m Module) Stop(fade time.Duration) {

}

func CurrentMod() (Module, bool) {
	return Module{}, false
}
