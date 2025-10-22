# pimod

Play music and sound effects in Pi games, directly from .MOD files (ProTracker modules).

## 🎯 Project Goals

Create package which will allow to load .MOD files and play them directly in Pi games using piaudio package.

### Play MOD tracks directly in piaudio channels

* Track is a column within a pattern that holds information for a specific instrument.

### Support ProTracker M.K. standard

* standard 4-channel, 64-pattern-max MOD.

### Support effects

* [ ] 0xy Arpeggio
* [ ] 1xy Slide up
* [ ] 2xy Slide down
* [ ] 3xy Slide to note
* [ ] 4xy Vibrato
* [ ] 5xy Volume slide + slide to note
* [ ] 6xy Volume slide + vibrato
* [ ] 7xy Tremolo
* [ ] 9xy Set sample offset
* [ ] Axy Volume slide
* [ ] Bxx Jump to offset
* [ ] Cxx Set volume
* [ ] Dxy Jump to row
* [ ] E0x - ??????
* [ ] E1x - Fine slide up
* [ ] E2x - Fine slide down
* [ ] E3x - Screw up 3xx
* [ ] E4x - Set vibrato waveform
* [ ] E5x - Set finetune value for the current sample
* [ ] E6x - loopback
* [ ] E7x - Set tremolo waveform, like E4x
*

## More info

* https://pollak.thebe.de/b/the-mod-format/
* https://wiki.multimedia.cx/index.php/Protracker_Module