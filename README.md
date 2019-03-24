# Indeterminate Music

![Musical compositions](https://user-images.githubusercontent.com/6550035/54872482-b25bee80-4d81-11e9-9378-8b3d1414e649.jpg)



> Indeterminate music is music in which the composer gives up some sort of control over the material she/he uses to create the composition.  It can take three rough shapes:
> - Stochastic Music: which bases the material on mathematical concepts (think Iannis Xenakis)
> - Chance Music: which bases the material on decisions made…well…by chance, like rolling dice or, in John Cage’s working with the I Ching.
> - Aleatory Music: which bases the material on the decisions made in performance (or before the performance) by the performers.
>
> *- [Sound American](http://soundamerican.org/sa_archive/sa10/index.html)*

This project is basically a framework for developing indeterminate music. This project is inspired by [Ólafur Arnalds](https://en.wikipedia.org/wiki/%C3%93lafur_Arnalds#re:member_(2018)), [Dan Tepfer](https://www.npr.org/2017/07/24/538677517/fascinating-algorithm-dan-tepfers-player-piano-is-his-composing-partner), the minimal music style from [Philip Glass](https://en.wikipedia.org/wiki/Philip_Glass#1967%E2%80%931974:_Minimalism:_From_Strung_Out_to_Music_in_12_Parts), and the experimental compositioning by [Christian Wolff](http://www.paristransatlantic.com/magazine/interviews/wolff.html).

Examples (generated based on current commit):

- https://soundcloud.com/schollz/b325dca47b2938c01b0689596047155550b8d527 ([b325dca](https://github.com/schollz/indeterminate-music/commit/b325dca47b2938c01b0689596047155550b8d527))
- https://soundcloud.com/schollz/8b0ba2769854009963e4ade90deaf964e3c6fd14 ([8b0ba27](https://github.com/schollz/indeterminate-music/commit/8b0ba2769854009963e4ade90deaf964e3c6fd14))

## How it works

Music is arranged from many pre-recorded snippets. A snippet is one 4/4 bar of music played with one hand in either the C major scale or the C minor scale. Snippets are later analyzed for their characteristics:

- Intensity (number of notes)
- Handiness (RH or LH)
- Minor / Major

These characteristics are then used to arrange them into a song according to the composition which eschews normal composition forms. The form of composition is TBD. A given chord structure will locate snippets and transpose them to the corresponding chord to be used. Midi files are generated for the RH and LH which can then be combined in the browser or in your favorite program.

## To Do

- [ ] Add getting started section
- [x] Determine/filter on number of notes
- [x] Determine/filter whether its LH or RH
- [ ] Build midi tracks in half-time (for holding notes)
- [ ] Interpolate loudness between pp, mf, and ff from velocity
- [ ] Optional sparsity, to randomly leave out notes of snippet
- [x] Add sustain: use triggerAttack instead of triggerAttackRelease when enabled
- [ ] Add two midi channels to frontend (merge them in JSON)


## Credits

Piano sounds for the front-end from [University of Iowa](http://theremin.music.uiowa.edu/MISpiano.html). The first part of the silence is removed with ffmpeg.

```
for i in *aiff; do ffmpeg -i "$i" -af silenceremove=1:0:-50dB "${i%.*}.mp3"; done
```

## License 

MIT
