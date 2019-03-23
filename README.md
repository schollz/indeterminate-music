# musical-keyboard
A musical keyboard, make music using MIDI and the browser

Piano sounds from [University of Iowa](http://theremin.music.uiowa.edu/MISpiano.html). The first part of the silence is removed with ffmpeg.

```
for i in *aiff; do ffmpeg -i "$i" -af silenceremove=1:0:-50dB "${i%.*}.mp3"; done
```
