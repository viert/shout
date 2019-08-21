# shout
This is a proof-of-concept alternative libshout bindings for golang. Unlike more straight-forward approaches, 
this project provides a WriteCloser adapter to a libshout connection so streaming a file is as easy as `io.Copy(wr, f)`


### Things to implement
  * non-blocking interface
  * metadata functions
  * tidying up
