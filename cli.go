package main

import (
	"go-on-jvm/jvm"
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

	f, err := os.Create("Test.class")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//s := "cafebabe00000034000307000201000a48656c6c6f576f726c640021000100000000000000000000"
	//bytes, _ := hex.DecodeString(s)
	//f.Write(bytes)

	class := jvm.NewClass("HelloWorld", jvm.ObjectClass)
	class.WithAccess(jvm.Super, jvm.Public)
	class.AddField(jvm.NewField("field", jvm.Int, jvm.Super, jvm.Public))
	method := jvm.NewMethod("method", jvm.Public)
	method.WithTypeDescriptor(jvm.Int)
	class.AddMethod(method)

	err = class.Write(f)
	if err != nil {
		panic(err)
	}
}
