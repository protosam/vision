package vision

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type New struct {
	output string
	blocks map[string]string
	Assignments map[string]string
	GAssignments map[string]string
}

func (tpl *New) TemplateFile(tpl_file string) {

	fbuffer, err := ioutil.ReadFile(tpl_file)

	if err != nil {
		fmt.Println(err)
		return
	}

	tpl.output = string(fbuffer)

	tpl.blocks = make(map[string]string)
	tpl.Assignments = make(map[string]string)
	tpl.GAssignments = make(map[string]string)

	tpl.output = tpl.parseblocks(tpl.output, "")

}


func (tpl *New) Assign(name string, value string) {
	// sanitize the value
	value = strings.Replace(value, "}", "_SANITIZED}", -1)
	tpl.Assignments[name] = value
}
func (tpl *New) GAssign(name string, value string) {
	// sanitize the value
	value = strings.Replace(value, "}", "_SANITIZED}", -1)
	tpl.GAssignments[name] = value
}


func (tpl *New) Parse(block_name string) {
	// removed for problems
	// tpl.output = strings.Replace(tpl.output, "<!-- BLOCK: " + block_name + " -->", tpl.blocks[block_name] + "<!-- BLOCK: " + block_name + " -->", 1)
	
	pos := strings.LastIndex(tpl.output, "<!-- BLOCK: " + block_name + " -->")
	tpl.output = tpl.output[:pos] + tpl.blocks[block_name] + tpl.output[pos:]

	
	for name, value := range tpl.GAssignments {
		tpl.output = strings.Replace(tpl.output, "{"+name+"}", value, -1)
	}
	
	for name, value := range tpl.Assignments {
		tpl.output = strings.Replace(tpl.output, "{"+name+"}", value, -1)
	}
	tpl.Assignments = make(map[string]string)
}

func (tpl *New) Out()  string {
	block_pattern := regexp.MustCompile("<!-- BLOCK: (.*?) -->")
	tpl.output = block_pattern.ReplaceAllString(tpl.output, "")
	//blanklines := regexp.MustCompile("(?ms:(^([[:space:]]+)?[\r\n]|^[\r\n]+))")
	//tpl.output = blanklines.ReplaceAllString(tpl.output, "")
	tpl.output = strings.TrimSpace(tpl.output)
	tpl.output = strings.Replace(tpl.output, "_SANITIZED}", "}", -1)
	return tpl.output
}

func (tpl *New) parseblocks(htmlin string, parent_blockname string) string {
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

			tpl.blocks[block_name] = tpl.parseblocks(block_content, block_name)
			htmlin = strings.Replace(htmlin, raw_block_content[0], "<!-- BLOCK: " + block_name + " -->", -1)

			//fmt.Println("Wrote " + block_name)
		}

		htmlin = tpl.parseblocks(htmlin, parent_blockname)
	}


	return htmlin;
}
