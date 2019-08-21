# shout
This is a proof-of-concept alternative libshout bindings for golang. Unlike more straight-forward approaches, 
this project provides a WriteCloser adapter to a libshout connection so streaming a file is as easy as `io.Copy(wr, f)`

## Example
```
package main

import (
	"io"
	"os"

	"github.com/viert/shout"
)

func main() {
	f, err := os.Open("song.mp3")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cfg := &shout.Config{
		Host:     "localhost",
		Port:     8000,
		User:     "source",
		Password: "passw0rd",
		Mount:    "/shuffle",
		Proto:    shout.ProtocolHTTP,
		Format:   shout.ShoutFormatMP3,
	}

	w, err := shout.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	io.Copy(w, f)
}

```

### Things to implement
  * non-blocking interface
  * metadata functions
  * tidying up
  * docs
