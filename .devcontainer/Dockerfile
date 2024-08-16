# From official php image.
FROM golang:alpine

RUN mkdir -p /workspace/cli

# Create a user group and account under id 1000.
RUN addgroup -g 1000 -S user && adduser -u 1000 -D user -G user

# Install quality-of-life packages.
RUN apk --update add --no-cache curl git vim fish tmux neovim wget make

RUN \
    # Install persistent / runtime deps
    apk --update add --no-cache clang llvm bcc bcc-tools gcompat libbpf-dev linux-headers

# Make fish the default shell
RUN sed -i 's|/bin/ash$|/usr/bin/fish|' /etc/passwd

# Prevent container from exiting early.
CMD ["sleep", "infinity"]
