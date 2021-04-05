package main

import (
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Open("1.mp3")
	if err != nil {
		log.Fatal(err)
	}
	st, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(st)
	select {}
}
