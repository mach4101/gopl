package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

type Params struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

var params = Params{
	cycles:  5,
	res:     0.001,
	size:    100,
	nframes: 64,
	delay:   8,
}

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0, 0, 0xff, 0xff}, color.RGBA{0xff, 0, 0, 0xff},
}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1
	greenIndex = 2 // next color in palette
	blueIndex  = 3
	redIndex   = 4
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Print(err)
	}

	for k, v := range r.Form {
		if num, err := strconv.Atoi(v[0]); err == nil {
			switch k {
			case "cycles":
				params.cycles = num
			case "res":
				params.res = float64(num)
			case "size":
				params.size = num
			case "nframes":
				params.nframes = num
			case "delay":
				params.delay = num
			}
		}
	}

	lissajous(w)
}

func lissajous(out io.Writer) {

	cycles := params.cycles
	res := params.res
	size := params.size
	nframes := params.nframes
	delay := params.delay

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
