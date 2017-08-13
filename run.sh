#!/bin/bash
if go build -ldflags -s ; then
        echo "Here we go..."
        ./sdl-demo
else
        echo "Build failed."
fi
