package main

import (
	"fmt"
	"github.com/protosam/vision"
)

func main() {

	var tpl vision.New
	tpl.TemplateFile("tpl/hello.tpl")
	
	tpl.Assign("testvar", "Foobar")
	
	tpl.Parse("main")
	tpl.Parse("main/row")
	tpl.Parse("main/row")
	tpl.Parse("main/row")
	
	
	
	tpl.Assign("foovar", "Hello World")
	tpl.Parse("main/vrow")
	tpl.Assign("foovar", "Hello Dog")
	tpl.Parse("main/vrow")
	tpl.Assign("foovar", "Hello Cat")
	tpl.Parse("main/vrow")
	
	
	
	fmt.Println(tpl.Out())
}
