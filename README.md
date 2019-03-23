# musical-keyboard

![Musical compositions](https://user-images.githubusercontent.com/6550035/54872482-b25bee80-4d81-11e9-9378-8b3d1414e649.jpg)


## Generative music

Generative music is arranged from snippets. One snippet is one 4/4 bar of music played with one hand in either the C major scale or the C minor scale. Snippets are analyzed for their characteristics:

- Intensity (number of notes)
- Handiness (RH or LH)
- Minor / Major

A given chord structure will locate snippets and transpose them to the corresponding chord to be used. Midi files are generated for the RH and LH which can then be combined in the browser or in your favorite program.

## To Do

- [ ] Determine/filter on number of notes
- [ ] Determine/filter whether its LH or RH
- [ ] Interpolate loudness between pp, mf, and ff from velocity
- [ ] Optional sparsity, to randomly leave out notes of snippet
- [x] Add sustain: use triggerAttack instead of triggerAttackRelease when enabled
- [ ] Add two midi channels to frontend (merge them in JSON)

## Inspiration

This project is inspired out what is possible with simple repetitive elements on the piano. I took a lot of inspriation from [Ólafur Arnalds](https://en.wikipedia.org/wiki/%C3%93lafur_Arnalds#re:member_(2018)), [Dan Tepfer](https://www.npr.org/2017/07/24/538677517/fascinating-algorithm-dan-tepfers-player-piano-is-his-composing-partner), the minimal music style from [Philip Glass](https://en.wikipedia.org/wiki/Philip_Glass#1967%E2%80%931974:_Minimalism:_From_Strung_Out_to_Music_in_12_Parts), and the experimental compositioning by [Christian Wolff](http://www.paristransatlantic.com/magazine/interviews/wolff.html).

## Credits

Piano sounds from [University of Iowa](http://theremin.music.uiowa.edu/MISpiano.html). The first part of the silence is removed with ffmpeg.

```
for i in *aiff; do ffmpeg -i "$i" -af silenceremove=1:0:-50dB "${i%.*}.mp3"; done
```

## License 

MIT