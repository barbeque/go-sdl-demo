package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
  }

  // blow the surface away
  fullScreenRect := sdl.Rect { 0, 0, 800, 600 }
  surface.FillRect(&fullScreenRect, 0x0)

  var running bool
  running = true

  for running {
    sdl.PumpEvents()

    rect := sdl.Rect{0, 0, 200, 200}
  	surface.FillRect(&rect, 0xffff0000)
  	window.UpdateSurface()

    var event sdl.Event
    for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
      switch t := event.(type) {
      case *sdl.QuitEvent:
        running = false
      case *sdl.KeyDownEvent:
        if t.Keysym.Sym == sdl.K_ESCAPE {
          running = false
        }
      }
    }

    sdl.Delay(1)
  }

	//sdl.Delay(1000)
	sdl.Quit()
}
