# musical-keyboard

A musical keyboard, make music using MIDI and the browser.


## Generative music

Generative music is arranged from snippets. One snippet is one 4/4 bar of music played with one hand in either the C major scale or the C minor scale. Snippets are analyzed for their characteristics:

- Intensity (number of notes)
- Handiness (RH or LH)
- Minor / Major

A given chord structure will locate snippets and transpose them to the corresponding chord to be used. Midi files are generated for the RH and LH which can then be combined in the browser or in your favorite program.

## To Do

- [ ] Interpolate loudness between pp, mf, and ff from velocity
- [ ] Optional sparsity, to randomly leave out notes of snippet
- [ ] Add sustain: use triggerAttack instead of triggerAttackRelease when enabled
- [ ] Add two midi channels to frontend (merge them in JSON)

## Credits

Piano sounds from [University of Iowa](http://theremin.music.uiowa.edu/MISpiano.html). The first part of the silence is removed with ffmpeg.

```
for i in *aiff; do ffmpeg -i "$i" -af silenceremove=1:0:-50dB "${i%.*}.mp3"; done
```

## License 

MIT
