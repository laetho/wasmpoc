# WASM POC

Or maybe WASM R&D is more precise. This repo is not really a POC of anything coherent, just
experimenting with Web Assembly and how it interacts with browser and javascript and the other
way around.



# Guides and How-To's

## Getting started
- https://www.aaron-powell.com/posts/2019-02-04-golang-wasm-1-introduction/
- https://golangbot.com/webassembly-using-go/
- https://tutorialedge.net/golang/go-webassembly-tutorial/

## Background material

- https://hacks.mozilla.org/2017/02/where-is-webassembly-now-and-whats-next/
- https://hacks.mozilla.org/2017/02/a-crash-course-in-just-in-time-jit-compilers/
- https://hacks.mozilla.org/2018/10/calls-between-javascript-and-webassembly-are-finally-fast-%F0%9F%8E%89/

# Projects and libraries

- DOM
  - https://github.com/gowebapi/webapi
  - https://github.com/dennwc/dom
  - https://github.com/dominikh/go-js-dom
- Canvas
  - https://github.com/markfarnan/go-canvas
  - 
  
# Inspiration

- https://github.com/agnivade/shimmer

# Usage

## Building
    cd wasmpoc
    ./buildall.sh
    
## Running

    cd cmd/server
    ./server
    
http://localhost:9090
