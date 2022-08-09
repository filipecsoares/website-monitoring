package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the Website Monitoring")
	for {
		showMenu()
		var option int
		fmt.Scan(&option)
		switch option {
		case 1:
			startMonitoring()
		case 2:
			startLogging()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func showMenu() {
	fmt.Println("Digit an option:")
	fmt.Println("1 - Monitoring")
	fmt.Println("2 - Logging")
	fmt.Println("0 - Exit")
}

func startMonitoring() {
	fmt.Println()
	fmt.Println("Monitoring")
	sites := readWebsiteFile()
	for _, site := range sites {
		resp, err := http.Get(site)
		if err != nil {
			fmt.Println("Error on get site", err)
			return
		}
		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "Successful!")
			saveLog(site, true)
		} else {
			fmt.Println("Something went wrong on site:", site, "Status Code:", resp.StatusCode)
			saveLog(site, false)
		}
	}
}

func startLogging() {
	fmt.Println()
	fmt.Println("Logging")
}

func readWebsiteFile() []string {
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Something went wrong with the file.", err)
		return nil
	}
	reader := bufio.NewReader(file)
	var sites []string
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		if err == io.EOF {
			break
		}
		sites = append(sites, row)
	}
	file.Close()
	return sites
}

func saveLog(site string, online bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online: " + strconv.FormatBool(online) + "\n")
	file.Close()
}
