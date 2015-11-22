package vision

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var output string
var blocks = make(map[string]string)
var assignments = make(map[string]string)

func TemplateFile(tpl_file string) {
	fbuffer, err := ioutil.ReadFile("tpl/hello.tpl")

	if err != nil {
		fmt.Println(err)
		return
	}

	output = string(fbuffer)
	output = parseblocks(output, "")

}


func Assign(name string, value string) {
	assignments[name] = value
}


func Parse(block_name string) {
	output = strings.Replace(output, "<!-- BLOCK: " + block_name + " -->", blocks[block_name] + "<!-- BLOCK: " + block_name + " -->", 1)
	for name, value := range assignments {
		output = strings.Replace(output, "{"+name+"}", value, -1)
	}
}

func Out()  string {
	block_pattern := regexp.MustCompile("<!-- BLOCK: (.*?) -->")
	output = block_pattern.ReplaceAllString(output, "")
	blanklines := regexp.MustCompile("(?ms:(^([[:space:]]+)?[\r\n]|^[\r\n]+))")
	output = blanklines.ReplaceAllString(output, "")
	return output
}

func parseblocks(htmlin string, parent_blockname string) string {
	begin_pattern := regexp.MustCompile("<!-- BEGIN: (.*?) -->")
	raw_block_name := begin_pattern.FindStringSubmatch(htmlin)

	if raw_block_name != nil {
		block_name := raw_block_name[1]

		block_pattern := regexp.MustCompile("<!-- BEGIN: " + block_name + " -->(?ms:(.*?))<!-- END: " + block_name + " -->")
		raw_block_content := block_pattern.FindStringSubmatch(htmlin)

		if raw_block_content != nil {
			if parent_blockname != "" {
				block_name = parent_blockname + "/" + block_name
			}
			block_content := raw_block_content[1]

			blocks[block_name] = parseblocks(block_content, block_name)
			htmlin = strings.Replace(htmlin, raw_block_content[0], "<!-- BLOCK: " + block_name + " -->", -1)

			//fmt.Println("Wrote " + block_name)
		}

		htmlin = parseblocks(htmlin, parent_blockname)
	}


	return htmlin;
}
