run:
	go generate
	go build -buildvcs=false
	./tutorial-minimal-example

trace:
	bpftool prog trace log

clean:
	rm -f *_bpf*.o && rm -f *_bpf*.go && rm -f tutorial-minimal-example
