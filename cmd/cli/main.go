package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go minimal ../../minimal.bpf.c

import (
	"log"
)

func main() {
	// Load the compiled eBPF ELF and load it into the kernel.
	var objs minimalObjects

	if err := loadMinimalObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()
}
