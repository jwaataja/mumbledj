# gumble

## Sub-projects

- gumble [![GoDoc](https://godoc.org/github.com/layeh/gumble/gumble?status.svg)](https://godoc.org/github.com/layeh/gumble/gumble)
    - Client library
- gumbleopenal
    - [OpenAL](http://kcat.strangesoft.net/openal.html) audio system for gumble
- gumbleffmpeg
    - [ffmpeg](https://www.ffmpeg.org/) audio source for gumble
- gumbleutil
    - Extras that can make working with gumble easier

## Example

    package main

    import (
      "github.com/layeh/gumble/gumble"
      "github.com/layeh/gumble/gumbleutil"
    )

    func main() {
      gumbleutil.Main(gumbleutil.Listener{
        UserChange: func(e *gumble.UserChangeEvent) {
          if e.Type.Has(gumble.UserChangeConnected) {
            e.User.Send("Welcome to the server, " + e.User.Name + "!")
          }
        },
      })
    }

## Related projects

- [barnard](https://github.com/layeh/barnard)
    - terminal-based Mumble client
- [piepan](https://github.com/layeh/piepan)
    - an easy to use framework for writing Mumble bots using Lua

## License

MPL 2.0

## Author

Tim Cooper (<tim.cooper@layeh.com>)
