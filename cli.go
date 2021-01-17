package main

import (
	"fmt"
	"go-on-jvm/compiler"
	"go-on-jvm/compiler/runtime"
	"go-on-jvm/parser"
	os "os"
)

func main() {
	parsed, err := parser.ParseDirectory("./example")
	if err != nil {
		panic(err)
	}
	_ = parsed

	//
	//var w bytes.Buffer
	//_, _ = fmt.Fprintf(&w, "%x", 0xCAFEBABE)
	//println(w.String())

	//s := "cafebabe00000034000307000201000a48656c6c6f576f726c640021000100000000000000000000"
	//bytes, _ := hex.DecodeString(s)
	//f.Write(bytes)

	classes := compiler.Compile(parsed)

	for _, class := range classes {
		write(class)
	}

	write(compiler.CompiledClass{
		Path:  "StandardLibrary",
		Class: runtime.NewStandardLib(),
	})
}

func write(class compiler.CompiledClass) {
	f, err := os.Create(fmt.Sprintf("%s.class", class.Path))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = class.Class.Write(f)
	if err != nil {
		panic(err)
	}
}
