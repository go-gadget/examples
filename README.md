# Gadget Examples

This repository contains some sample/experimentation code for the Gadget (golang) frontend framework

## Building

```
$ GOARCH=wasm GOOS=js go build -o lib.wasm github.com/go-gadget/examples/cmd/todo

$ bin/gadget serve
```

(bin/gadget is built from go-gadget/gadget)