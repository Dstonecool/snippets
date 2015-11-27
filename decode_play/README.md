# Decoding and playing audio file in Linux

There are two types of snippets here:
* `decode` - snippet gets filename as argument, decodes audio stream from file, and writes decoded samples to stdout;
* `play` - snippet reads decoded samples from stdin and sends them to the soundcard.

Decoded samples are always in the same form:
* linear PCM;
* two channels (front Left, front Right);
* interleaved format (L R L R ...);
* samples are 32-bits floats in little endian (actually CPU should be little-endian too);
* sample rate is 44100 Hz.

### Building

```
$ make          # build all snippets
$ make clean    # cleanup
```

### Example usage

`decode` and `play` snippets may be connected via pipe. Any combinations are allowed, e.g.:

```
$ ./ffmpeg_decode     cool_song.mp3   |  ./alsa_play_tuned
$ ./sox_decode_chain  cool_song.mp3   |  ./ffmpeg_play
$ ./sndfile_decode    cool_song.flac  |  ./sox_play
```

### `decode` snippets

* `ffmpeg_decode` - decode file using [FFmpeg](https://www.ffmpeg.org/) (automatic resampling and channel mapping)
* `sox_decode_simple` - decode file using [SoX](http://sox.sourceforge.net/) (no resampling and channel mapping)
* `sox_decode_chain` - decode file using [SoX](http://sox.sourceforge.net/) (automatic resampling and channel mapping using effects chain)
* `sndfile_decode` - decode file using [libsndfile](http://www.mega-nerd.com/libsndfile/) (no resampling and channel mapping, only limited number of formats supported)

### `play` snippets

* `ffmpeg_play` - play decoded samples using [FFmpeg](https://www.ffmpeg.org/)
* `ffmpeg_play_encoder` - play decoded samples using [FFmpeg](https://www.ffmpeg.org/) (a bit more complex example demonstrating encoder usage)
* `sox_play` - play decoded samples using [SoX](http://sox.sourceforge.net/)
* `alsa_play_simple` - play decoded samples using `libasound` (with default parameters)
* `alsa_play_tuned` - play decoded samples using `libasound` (with customized parameters)
