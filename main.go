package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sync"
	"syscall"
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGPROF)
	go func() {
		for {
			<-sig
		}
	}()
	runtime.SetCPUProfileRate(400)
	fh, err := os.Create("cpu.prof")
	if err != nil {
		panic("cannot create cpu.prof")
	}

	defer fh.Close()

	pprof.StartCPUProfile(fh)
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup

	for i := uint64(0); i < 5000000; i++ {
		wg.Add(1)
		go doWork(&wg, i)
		if i%100000 == 0 {
			runtime.GC()
		}
	}
}

func doWork(wg *sync.WaitGroup, counter uint64) {
	defer wg.Done()
	s := fnord.GetRandomString(120)
	if fnord.StringContainsName(s) {
		color.Green(s)
		color.Green("HIT: %d\n", counter)
		fmt.Printf("Elapsed time: %v\n", time.Since(startTime))
		os.Exit(0)
	}
	if fnord.StringRegexName(s) {
		color.Green(s)
		color.Green("HIT: %d\n", counter)
		fmt.Printf("Elapsed time: %v\n", time.Since(startTime))
		os.Exit(0)
	}
}

func setTimeZoneToUTC() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic("Unable to set TZ to UTC")
	}
	time.Local = loc
}
