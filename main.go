package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/mojotx/profile/pkg/fnord"
)

var startTime time.Time

func init() {
	setTimeZoneToUTC()
	startTime = time.Now()
}

func main() {
	color.Green("Starting CPU profiling")

	/*
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGPROF)
		go func() {
			for {
				<-sig
			}
		}()
	*/
	runtime.SetCPUProfileRate(400)
	fh, err := os.Create("cpu.prof")
	if err != nil {
		panic("cannot create cpu.prof")
	}

	defer fh.Close()

	if serr := pprof.StartCPUProfile(fh); serr != nil {
		panic(serr.Error())
	}
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup

	for i := uint64(0); i < 500; i++ {
		wg.Add(1)
		go doWork(&wg)
		if i%100000 == 0 {
			runtime.GC()
		}
	}
}

func doWork(wg *sync.WaitGroup) {
	defer wg.Done()
	s := fnord.GetWorkingData()

	fnord.BubbleSort(s)
	color.HiBlack("BubbleSort done in %v\n", time.Since(startTime))
}

func setTimeZoneToUTC() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic("Unable to set TZ to UTC")
	}
	time.Local = loc
}
