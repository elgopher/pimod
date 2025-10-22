package main

import (
	_ "embed"
	"time"

	"github.com/elgopher/pi"
	"github.com/elgopher/pi/piaudio"
	"github.com/elgopher/pi/piebiten"
	"github.com/elgopher/pimod"
)

//go:embed "example.mod"
var exampleMod []byte

func main() {
	pi.Init = func() {
		mod := pimod.Decode(exampleMod)
		mod.Play(piaudio.Chan1|piaudio.Chan2, 0, 100*time.Millisecond)
	}
	piebiten.Run()
}
