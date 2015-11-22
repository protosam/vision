package main

import (
	"fmt"
	"github.com/protosam/vision"
)

func main() {


	vision.TemplateFile("tpl/hello.tpl")
	
	vision.Assign("testvar", "Foobar")
	
	vision.Parse("main")
	vision.Parse("main/row")
	vision.Parse("main/row")
	vision.Parse("main/row")
	
	
	
	vision.Assign("foovar", "Hello World")
	vision.Parse("main/vrow")
	vision.Assign("foovar", "Hello Dog")
	vision.Parse("main/vrow")
	vision.Assign("foovar", "Hello Cat")
	vision.Parse("main/vrow")
	
	
	
	fmt.Println(vision.Out())
}
