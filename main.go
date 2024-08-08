package main

import (
	"log"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const WIDTH, HEIGHT = 800, 800

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, WIDTH, HEIGHT)
	if err != nil {
		log.Fatal(err)
	}
	defer texture.Destroy()

	pixels := make([]byte, WIDTH*HEIGHT*4)

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {

		}
	}

	texture.Update(nil, unsafe.Pointer(&pixels[0]), WIDTH*4)
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}
}
