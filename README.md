# WASM POC

Or maybe WASM R&D is more precise. This repo is not really a POC of anything coherent, just
experimenting with Web Assembly and how it interacts with browser and javascript and the other
way around.



# Guides and How-To's

- https://www.aaron-powell.com/posts/2019-02-04-golang-wasm-1-introduction/
- https://golangbot.com/webassembly-using-go/
- https://tutorialedge.net/golang/go-webassembly-tutorial/

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
