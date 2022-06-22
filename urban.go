package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
)

// Response structure
type response struct {
	List []struct {
		Definition string `json:"definition"`
		Example    string `json:"example"`
		Word       string `json:"word"`
	} `json:"list"`
}

// Constants (API URLs)
const (
	API_URL      = "https://api.urbandictionary.com/v0/define?term="
	RAND_API_URL = "https://api.urbandictionary.com/v0/random"
)

func printf(msg string) {
	fmt.Print(msg)
}

// For createing request on API_URL
func req(query string, raw bool) ([]byte, response) {
	var resp response
	body, err := http.Get(API_URL + url.QueryEscape(query))

	buf, err := ioutil.ReadAll(body.Body)
	if raw == false {
		err = json.Unmarshal([]byte(string(buf)), &resp)
	} else {
		return buf, resp
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	defer body.Body.Close()
	return buf, resp
}

// For creating request om RAND_API_URL
func query_random(raw bool) ([]byte, response) {
	var resp response
	body, err := http.Get(RAND_API_URL)
	buf, err := ioutil.ReadAll(body.Body)
	if raw == false {
		err = json.Unmarshal([]byte(string(buf)), &resp)
	} else {
		return buf, resp
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	defer body.Body.Close()
	return buf, resp
}

// Render the structured data with glamour
func render(data string) {
	glam, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	render, err := glam.Render(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	printf(render)
}

// Join all charecters as an arguments
func getArgs(args []string) string {
	return strings.Join(args, " ")
}

// Shows help
func help() {
	printf("urban v0.1\n\nUsage:\nurban [options] [word]\n\nOptions:\n -h, shows help menu\n -d, definition of a word or text\n -e, examples of a word or text\n -j, raw json stdout\n -r, get random definitions or examples (must use as a second param)\n")
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "-d":
		if len(os.Args) < 3 {
			help()
			break
		}

		if os.Args[2] == "-r" {
			_, resp := query_random(false)
			data := ""
			data += "## Definitions of " + strings.Title(resp.List[0].Word) + "\n"
			for i := 0; i < len(resp.List); i++ {
				data += "* " + strings.ReplaceAll(resp.List[i].Definition, "*", "") + "\n"
			}
			render(data)
			break
		}

		_, resp := req(getArgs(os.Args[2:]), false)

		if len(resp.List) == 0 {
			printf("No results found matching your query.\n")
			os.Exit(0)
		}
		data := ""
		data += "## Definitions of " + strings.Title(strings.Join(os.Args[2:], " ")) + "\n"
		for i := 0; i < len(resp.List); i++ {
			data += "* " + strings.ReplaceAll(resp.List[i].Definition, "*", "") + "\n"
		}
		render(data)
		break

	case "-e":
		if len(os.Args) < 3 {
			help()
			break
		}

		if os.Args[2] == "-r" {
			_, resp := query_random(false)
			data := ""
			data += "## Examples of " + strings.Title(strings.Join(os.Args[2:], " ")) + "\n"
			for i := 0; i < len(resp.List); i++ {
				data += "* " + strings.ReplaceAll(resp.List[i].Example, "*", "") + "\n"
			}
			render(data)
			break
		}

		_, resp := req(getArgs(os.Args[2:]), false)

		if len(resp.List) == 0 {
			printf("No results found matching your query.\n")
			break
		}
		data := ""
		data += "## Examples of " + strings.Title(strings.Join(os.Args[2:], " ")) + "\n"
		for i := 0; i < len(resp.List); i++ {
			data += "* " + strings.ReplaceAll(resp.List[i].Example, "*", "") + "\n"
		}
		render(data)
		break

	case "-j":
		if len(os.Args) < 3 {
			help()
			break
		}

		if os.Args[2] == "-r" {
			buf, _ := query_random(true)
			printf(string(buf))
			break
		}
		raw, _ := req(getArgs(os.Args[2:]), true)
		printf(string(raw))
		break

	case "-h":
		help()
		break

	case "-help":
		help()
		break

	case "--help":
		help()
		break

	default:
		if len(os.Args) > 1 {
			printf(os.Args[1] + " isn't a valid option. Try with urban --help for more info.\n")
		}
		break
	}
}
