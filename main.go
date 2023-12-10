/*
Submission By: Khagendra Kumar Mandal
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

// Request structure for input payload
type SortRequest struct {
	ToSort [][]int `json:"to_sort"`
}

// Response structure for JSON response
type SortResponse struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

func main() {
	// Define routes
	// Serve static files (HTML, JS)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/process-single", processSingle)
	http.HandleFunc("/process-concurrent", processConcurrent)
	fmt.Println("Server 8000 port started")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Function to handle /process-single endpoint
func processSingle(w http.ResponseWriter, r *http.Request) {
	var req SortRequest
	decodeRequest(w, r, &req)

	if len(req.ToSort) == 0 {
		http.Error(w, "Input array is empty", http.StatusBadRequest)
		return
	}

	fmt.Println("Received request for /process-single")
	fmt.Printf("to_sort: %v\n", req.ToSort)

	startTime := time.Now()
	sortedArrays := sortSequential(req.ToSort)
	endTime := time.Now()

	fmt.Printf("Start Time: %v\nEndTime: %v\n", startTime, endTime)
	fmt.Printf("Sorting sequential: %v\n", sortedArrays)
	fmt.Printf("Time Taken for Sequential Sorting: %v\n", endTime.Sub(startTime))

	sendResponse(w, sortedArrays, endTime.Sub(startTime).Nanoseconds())
}

// Function to handle /process-concurrent endpoint
func processConcurrent(w http.ResponseWriter, r *http.Request) {
	var req SortRequest
	decodeRequest(w, r, &req)

	if len(req.ToSort) == 0 {
		http.Error(w, "Input array is empty", http.StatusBadRequest)
		return
	}

	fmt.Println("Received request for /process-concurrent")
	fmt.Printf("")

	startTime := time.Now()
	sortedArrays := sortConcurrent(req.ToSort)
	endTime := time.Now()

	fmt.Printf("to_sort: %v\nStart Time: %v\nEndTime: %v\n", req.ToSort, startTime, endTime)
	fmt.Printf("Sorting sequential: %v\n", sortedArrays)
	fmt.Printf("Time Taken for Sequential Sorting: %v\n", endTime.Sub(startTime))

	sendResponse(w, sortedArrays, endTime.Sub(startTime).Nanoseconds())
}

// Function to decode JSON request
func decodeRequest(w http.ResponseWriter, r *http.Request, req *SortRequest) {
	fmt.Println("Decode Request")
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
}

// Function to sort arrays sequentially
func sortSequential(arrays [][]int) [][]int {
	for i := range arrays {
		sort.Ints(arrays[i])
	}
	return arrays
}

// Function to sort arrays concurrently
func sortConcurrent(arrays [][]int) [][]int {
	var wg sync.WaitGroup
	wg.Add(len(arrays))

	for i := range arrays {
		go func(i int) {
			defer wg.Done()
			sort.Ints(arrays[i])
		}(i)
	}

	wg.Wait()
	return arrays
}

// Function to send JSON response
func sendResponse(w http.ResponseWriter, sortedArrays [][]int, timeNs int64) {
	response := SortResponse{
		SortedArrays: sortedArrays,
		TimeNs:       timeNs,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

/*
For output running command line using curl
path\Sort_Array>curl -X POST -H "Content-Type: application/json" -d "{\"to_sort\": [[3, 2, 1], [6, 5, 4], [9, 8, 7]]}" http://localhost:8000/process-single
output is as below
{"sorted_arrays":[[1,2,3],[4,5,6],[7,8,9]],"time_ns":0}
*/
