# Pong 

Pong in Go with SDL2 bindings.

## Installation

You need to have the SDL toolchain for `go-sdl2`. On Mac I do this with: 

```bash
brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
```

There are also instructions [here](https://github.com/veandco/go-sdl2) for other OSs.

Then: 

```bash
git clone https://github.com/alfiehiscox/pong.git
cd pong
go mod tidy
go make
./pong
```

## Improvments: 

- [ ] Score
- [ ] End Game

## Refs:
[Tutorial](https://www.youtube.com/watch?v=4RAwgmLjdCs)

