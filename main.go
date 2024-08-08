package main

import (
	"log"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const WIDTH, HEIGHT = 800, 600

type Pos struct {
	x, y float64
}

type Bat struct {
	Pos
	w, h   float64
	c      sdl.Color
	xv, yv float64
}

func (b *Bat) Draw(pixels []byte) {
	startX := b.x - b.w/2
	startY := b.y - b.h/2
	endX := b.x + b.w/2
	endY := b.y + b.h/2

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			drawPixel(int(x), int(y), b.c, pixels)
		}
	}
}

func (b *Bat) Update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 && b.y-b.yv-(b.h/2) >= 0 {
		b.y -= b.yv
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 && b.y+b.yv+(b.h/2) <= HEIGHT {
		b.y += b.yv
	}
}

func (b *Bat) AIUpdate(ball *Ball) {
	b.y = ball.y
}

type Ball struct {
	Pos
	r      float64
	c      sdl.Color
	xv, yv float64
}

func (b *Ball) Draw(pixels []byte) {
	for y := -b.r; y < b.r; y++ {
		for x := -b.r; x < b.r; x++ {
			if x*x+y*y < b.r*b.r {
				drawPixel(int(b.x+x), int(b.y+y), b.c, pixels)
			}
		}
	}
}

func (b *Ball) Update(left, right *Bat) {
	b.x += b.xv
	b.y += b.yv

	if b.y < 0+b.r || b.y > HEIGHT-b.r {
		b.yv = -b.yv
	}

	// TODO : Score
	if b.x-b.r < 0 || b.x+b.r > WIDTH {
		b.x = WIDTH / 2
		b.y = HEIGHT / 2
	}

	if b.x < left.x+left.w/2 {
		if b.y > left.y-left.h/2 && b.y < left.y+left.h/2 {
			b.xv = -b.xv
		}
	}

	if b.x > right.x-right.w/2 {
		if b.y > right.y-right.h/2 && b.y < right.y+right.h/2 {
			b.xv = -b.xv
		}
	}
}

func clearPixels(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func drawPixel(x, y int, c sdl.Color, pixels []byte) {
	index := (y*WIDTH + x) * 4
	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.R
		pixels[index+1] = c.G
		pixels[index+2] = c.B
	}
}

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
	keyState := sdl.GetKeyboardState()

	bat1 := Bat{Pos{0 + 40, 140}, 30, 80, sdl.Color{R: 255, G: 255, B: 255, A: 0}, 10, 10}
	bat2 := Bat{Pos{WIDTH - 40, 140}, 30, 80, sdl.Color{R: 255, G: 255, B: 255, A: 0}, 10, 10}
	ball := Ball{Pos{180, 180}, 20, sdl.Color{R: 255, G: 255, B: 255, A: 0}, 10, 10}

	for {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		clearPixels(pixels)

		bat1.Update(keyState)
		bat2.AIUpdate(&ball)
		ball.Update(&bat1, &bat2)

		bat1.Draw(pixels)
		bat2.Draw(pixels)
		ball.Draw(pixels)

		texture.Update(nil, unsafe.Pointer(&pixels[0]), WIDTH*4)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		sdl.Delay(33)
	}
}
