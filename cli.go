package main

import (
	"fmt"
	"go-on-jvm/jvm"
	"go/ast"
	"os"
)

type Visitor struct {
}

func (v Visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		println("nil")
	} else {
		fmt.Printf("%#v\n", node)
	}

	return v
}

func main() {
	//fileSite := token.NewFileSet()
	//file, err := parser.ParseFile(fileSite, "test.go", nil, parser.AllErrors)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//ast.Walk(Visitor{}, file)

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

	err = class.Compile(f)
	if err != nil {
		panic(err)
	}
}
