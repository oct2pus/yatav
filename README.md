# yatav
*Y*et *A*nother *T*rack *A*rt *V*iewer

I'm currently using CMUS a lot. I wanted something to display the images.
This was supposed to be a simple shell script and it quickly got out of hand, send help.

This has minimal testing, don't expect much.

## Features
- sure handles images
- crashes when a song doesn't have an id3v2 image tag
  - probably crashes even harder if the image doesn't have one at all
- only handles .mp3s
  - .ogg & .flacs also work probably
- under 100 lines of code (and many, *many* gigantic packages imported)
- sometimes the text gets managed by being too long
- doesn't seem to resize right always and i don't know why
- hey, technically it works under wayland!

## Screenshot

![A picture of my desktop. With both Cmus and yatav running.](https://raw.githubusercontent.com/oct2pus/yatav/master/screenshot.png)
