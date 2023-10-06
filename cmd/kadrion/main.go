package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"sync"
	"time"
)

var max_concurrent_requests int
var number_requests int

type Tconfig struct {
	API []struct {
		Endpoint string                 `yaml:"endpoint"`
		Method   string                 `yaml:"method"`
		JSON     map[string]interface{} `yaml:"JSON"`
	} `yaml:"API"`
}
type TestResult struct {
	ResponseTime time.Duration
}

const Version string = "0.10.0"

var loader bool
var tconfigs Tconfig
var concurrent_response_time time.Duration
var output_table *tablewriter.Table
var endpoint_number int

func main() {

	helpFlag := flag.Bool("help", false, "Tool Usage Information")
	versionFlag := flag.Bool("version", false, "Provide current version rolling")
	args := os.Args[1:]

	flag.Parse()

	if *helpFlag == true {
		toolInfo()
		os.Exit(0)
	}
	if *versionFlag == true {
		fmt.Printf(" kadrion %s\n", Version)
	}
	if len(args) == 2 && args[0] == "apply" && args[1] == "tconfig.yaml" {
		_, err := os.Stat("tconfig.yaml")
		if os.IsNotExist(err) {
			fmt.Println("tconfig.yaml file not found")
			os.Exit(1)
		} else {
			readfile, err := os.ReadFile("tconfig.yaml")
			if err != nil {
				fmt.Println("Error reading tconfig.yaml:", err)
				os.Exit(1)

			} else {
				if err := yaml.Unmarshal(readfile, &tconfigs); err != nil {
					fmt.Println("Unable to read YAML file due to:", err)
					os.Exit(1)
				} else {
					LoadingQueue()
				}
			}
		}
	} else {
		toolInfo()
	}

}

func toolInfo() {
	fmt.Println("\n Kadrion - A Continuous Testing CLI Tool")
	fmt.Println("\n Usage:")
	fmt.Println("\n    kadrion apply tconfig.yaml")
	fmt.Println("\n Available scopes to use are:")
	fmt.Println("\n    api     Performance testing of an API Endpoint")
	fmt.Println("\n    cluster     Validate Kube deployment")
	fmt.Println("\n Documentation for configuration can be found at\n\n https://www.kadriontestops.tech/docs")
	fmt.Println()
	fmt.Println(" Additional Commands --help and --version")
	fmt.Println()
}

func LoadingQueue() {
	frames := []string{"|", "/", "-", "\\"}

	i := 0
	loader = true
	fmt.Println()
	go func() {
		for loader == true {
			fmt.Printf("\r %s Wait running test against your config file", frames[i])
			i = (i + 1) % len(frames)
			time.Sleep(100 * time.Millisecond)
		}

	}()
	endpoint_number = 1
	for _, tconfig := range tconfigs.API {
		time.Sleep(1 * time.Second)
		testEndpoint(tconfig.Endpoint, tconfig.Method, tconfig.JSON)
		time.Sleep(1 * time.Second)
		endpoint_number += 1
	}

}

var result_attempt []int
var result_concurrent_requests []int
var result_response_time []time.Duration

func testEndpoint(endpoint, method string, jsonbody map[string]interface{}) {
	max_concurrent_requests := 2000
	number_requests := 10
	concurrent_request_differential := max_concurrent_requests * 5 / 100
	result_response_time := []time.Duration{}
	result_attempt := []int{}
	result_concurrent_requests := []int{}

	for i := 1; i <= number_requests; i++ {
		concurrent_requests(endpoint, method, jsonbody, concurrent_request_differential)
		result_response_time = append(result_response_time, concurrent_response_time)
		result_attempt = append(result_attempt, i)
		result_concurrent_requests = append(result_concurrent_requests, concurrent_request_differential)

		if i <= number_requests/2 {
			concurrent_request_differential += max_concurrent_requests * 5 / 100
		} else {
			concurrent_request_differential -= max_concurrent_requests * 5 / 100
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Attempt", "Concurrent Requests", "Response Time"})
	for i := 0; i < len(result_attempt); i++ {
		Each_Row_Table := []string{
			fmt.Sprintf("%d", result_attempt[i]),
			fmt.Sprintf("%d", result_concurrent_requests[i]),
			result_response_time[i].String(),
		}
		table.Append(Each_Row_Table)
	}
	loader = false
	time.Sleep(1 * time.Second)
	if endpoint_number == 1 {
		fmt.Println()
	}
	fmt.Println()
	fmt.Printf("Printing results for endpoint %v", endpoint_number)
	fmt.Println()
	table.Render()

}

func concurrent_requests(endpoint, method string, jsonbody map[string]interface{}, concurrent_request_number int) {
	result_channel := make(chan time.Duration, max_concurrent_requests)
	var wg sync.WaitGroup
	for i := 0; i < concurrent_request_number; i++ {
		wg.Add(1)
		go response_time_single(endpoint, method, jsonbody, result_channel, &wg)
	}
	go func() {
		wg.Wait()
		close(result_channel)
	}()
	totalResponseTime := time.Duration(0)
	Num_of_Responses := 0
	for responseTime := range result_channel {
		if responseTime > 0 {
			totalResponseTime += responseTime
			Num_of_Responses++
		}
	}
	if Num_of_Responses > 0 {
		concurrent_response_time = totalResponseTime / time.Duration(Num_of_Responses)
	} else {
		concurrent_response_time = 0
	}
}

func response_time_single(endpoint, method string, jsonbody map[string]interface{}, result_channel chan<- time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	marshaled_data, err := json.Marshal(jsonbody)
	if err != nil {
		fmt.Println("Error processing json data due to:", err)
		result_channel <- 0
		return
	}
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(marshaled_data))
	if err != nil {
		fmt.Println("Error creating API request due to:", err)
		result_channel <- 0
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	startTime := time.Now()
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching API response due to:", err)
		result_channel <- 0
		return
	}
	defer response.Body.Close()
	responseTime := time.Since(startTime)
	result_channel <- responseTime

}
