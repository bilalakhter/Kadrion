package testprocess

import (
	"bytes"
	"fmt"
	"github.com/bilalakhter/kadrion/internal/customtypes"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"sync"
	"time"
)

var loader bool
var concurrent_response_time time.Duration
var output_table *tablewriter.Table
var endpoint_number int
var max_concurrent_requests int
var number_requests int
var tconfigs customtypes.Tconfig
var result_attempt []int
var result_concurrent_requests []int
var result_response_time []time.Duration

func ProcessYaml(readfile []byte) {
	if err := yaml.Unmarshal(readfile, &tconfigs); err != nil {
		fmt.Println("Unable to read YAML file due to:", err)
		os.Exit(1)
	} else {
		TestQueue()
	}
}

func TestQueue() {

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
	for _, tconfig := range tconfigs.API.Endpoints {
		TestEndpoint(tconfig.Endpoint, tconfig.Method, []byte(tconfig.JSON), tconfigs.API.MaxConcurrentRequests, tconfigs.API.NumberOfRequests)
		endpoint_number += 1
	}

}

func TestEndpoint(endpoint, method string, jsonbody []byte, max_concurrent_requests int, number_requests int) {

	concurrent_request_differential := max_concurrent_requests / number_requests
	result_response_time := []time.Duration{}
	result_attempt := []int{}
	result_concurrent_requests := []int{}

	for i := 1; i <= number_requests; i++ {
		Concurrent_requests(endpoint, method, jsonbody, concurrent_request_differential)
		result_response_time = append(result_response_time, concurrent_response_time)
		result_attempt = append(result_attempt, i)
		result_concurrent_requests = append(result_concurrent_requests, concurrent_request_differential)

		concurrent_request_differential += concurrent_request_differential

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

func Concurrent_requests(endpoint, method string, jsonbody []byte, concurrent_request_number int) {
	result_channel := make(chan time.Duration, max_concurrent_requests)
	var wg sync.WaitGroup
	for i := 0; i < concurrent_request_number; i++ {
		wg.Add(1)
		go Response_time_single(endpoint, method, jsonbody, result_channel, &wg)
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

func Response_time_single(endpoint, method string, jsonbody []byte, result_channel chan<- time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonbody))
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
	if response.StatusCode != http.StatusOK {
		fmt.Println("Request did not recieved 200 status:", response.Status)
		os.Exit(0)
	}
	responseTime := time.Since(startTime)
	result_channel <- responseTime

}
