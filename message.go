package main

import (
	"bytes"
	"fmt"
	"net/http"
	"bufio"
	"os"
	"strings"
	"io"
)
var (
	token     = readConf("config.conf", "TOKEN")
	channelID = readConf("config.conf", "CHANNEL_ID")
	apiURL    = "https://discord.com/api/v10/channels/" + channelID + "/messages"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readConf(fileName string, param string) string {
	f, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		configParam := strings.Split(line, " = ")[0]
		if configParam == "["+param+"]" {
			return strings.Split(line, " = ")[1]
		}
	}
	return ""
}


func postMessage(message string) {
	
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(message)))
	req.Header.Set("Authorization", "Bot "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Message sent! Status:", resp.Status)
}

func readMessages() {
	
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("Authorization", "Bot "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
	    fmt.Println("Error reading response:", err)
	    return
	}
	fmt.Println(string(body))
	fmt.Println("Message sent! Status:", resp.Status)
}


func main() {
	message := `{"content": "Hello, Discord!"}`
	readMessages()
	postMessage(message)
}