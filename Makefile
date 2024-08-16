run:
	go generate
	go build
	./tutorial-minimal-example

clean:
	rm -f *_bpf*.o && rm -f *_bpf*.go && rm -f tutorial-minimal-example
