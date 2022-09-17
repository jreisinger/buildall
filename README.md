Build (cross-compile) binaries for all supported platforms from Go source files.
Platform is a combination of operating system and architecture, e.g.
linux/amd64.

```
$ go install 

$ buildall main.go
OK:  go build -o binaries/main-linux-ppc64 main.go
OK:  go build -o binaries/main-android-arm64 main.go
OK:  go build -o binaries/main-aix-ppc64 main.go
OK:  go build -o binaries/main-openbsd-mips64 main.go
OK:  go build -o binaries/main-linux-mips64le main.go
OK:  go build -o binaries/main-freebsd-amd64 main.go
ERR: go build -o binaries/main-android-arm main.go
<...snip...>
```