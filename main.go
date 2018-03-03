package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"log"
	"math"
	"time"
)

func main() {
	log.Println("Gromo starting up")

	sr := beep.SampleRate(48000)
	speaker.Init(sr, sr.N(time.Second/10))

	streamer := SineWave(sr, 256)

	done := make(chan struct{})
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(done)
	})))

	<-done
}

func SineWave(sr beep.SampleRate, freq float64) beep.Streamer {
	t := 0.0
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			y := math.Sin(math.Pi * freq * t)
			samples[i][0] = y
			samples[i][1] = y
			t += sr.D(1).Seconds()
		}
		return len(samples), true
	})
}
