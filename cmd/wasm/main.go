package main

import (  
    "fmt"
    "math/rand"
    "strconv"
    "syscall/js"
    "time"
)

var Nodes int64 = 1

func main() {
    fmt.Println("Go Web Assembly")
    js.Global().Set("anotherFunc", anotherFunc())
    js.Global().Set("addToDOM", addToDOM())
    js.Global().Set("updateGraph", updateGraph())
    <-make(chan bool)
}

func anotherFunc() js.Func {
    f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return wasmFunction()
    })
    return f
}

func wasmFunction() string {
    return "i'm a return string from a wasm go function called from an exported to js function."
}

func addToDOM() js.Func {
    f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        jsDoc := js.Global().Get("document")
        e := jsDoc.Call("createElement", "p")
        e.Set("innerHTML",wasmFunction())
        jsDoc.Get("body").Call("appendChild", e)
        return nil
    })
    return f
}

func updateGraph() js.Func {
    f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        Nodes++ // bump
        jsDoc := js.Global().Get("document")
        e := jsDoc.Call("getElementById", "data")
        e.Set("innerHTML",genGraph())
        js.Global().Call("draw")
        fmt.Println(Nodes)
        return nil
    })
    return f
}

func genGraph() string {
    start := "digraph {"
    stop := "}"
    graph := start
    for i := int64(1); i <= 100; i++ {
        rand.Seed(time.Now().UnixNano())
        edge := rand.Int63n(1000-0+1)
        graph += strconv.FormatInt(i,10)+" -> "+ strconv.FormatInt(edge, 10)+";\n"
        edge2 := rand.Int63n(1000-0+1)
        graph += strconv.FormatInt(edge,10)+" -> "+ strconv.FormatInt(edge2, 10)+";\n"
    }
    graph += stop
    return graph
}