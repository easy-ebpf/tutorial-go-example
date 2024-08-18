package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go minimal minimal.bpf.c

import (
	"log"
	"math"
	"time"

	"github.com/cilium/ebpf/link"
)

func main() {

	// Load the compiled eBPF ELF and load it into the kernel.
	var objs minimalObjects

	if err := loadMinimalObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	// Attach to the same tracepoint mentioned in minimal.bpf.c
	kp, err := link.Tracepoint("syscalls", "sys_enter_write", objs.minimalPrograms.HandleTp, nil)
	if err != nil {
		log.Fatalf("opening tracepoint: %s", err)
	}
	defer kp.Close()

	// Waiting for events...
	log.Println("Waiting for events..")
	time.Sleep(math.MaxInt64)

}
