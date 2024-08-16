setup:
	# Setup Eunomia BPF utils
	wget https://aka.pw/bpf-ecli -O ./ecli && chmod +x ./ecli
	wget https://github.com/eunomia-bpf/eunomia-bpf/releases/latest/download/ecc -O ./ecc && chmod +x ./ecc

run:
	./ecc minimal.bpf.c
	./ecli run package.json

help:
