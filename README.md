# Go Server for Sorting Array

This repository contains a simple Go server that sorts arrays using sequential and concurrent processing. The server exposes two endpoints, /process-single and /process-concurrent, to demonstrate the efficiency of both methods.

# Table of Contents
* Project Structure
* Usage
* Endpoints
* Performance Measurement
* Dockerization

# Project Structure
The project is organized as follows:
* main.go: The main server implementation with endpoint handlers.
* index.html: HTML file containing a simple user interface to interact with the server.
* Dockerfile: Instructions for building a Docker image of the server.
* .dockerfignore: Files and directories to be ignored during Docker image build.
* go.mod: Go module file where complete go code is written.

# Usage
Clone the repository:
bash code:
git clone https://github.com/kumarkhagendra909/sorting-server.git

cd sorting-server
Build and run the server locally:

bash code:
go run main.go

Access the server UI at http://localhost:8000 to test the endpoints interactively.

# Endpoints
/process-single
Sorts each sub-array sequentially.
Request:

json 
{
  "to_sort": [[3, 2, 1], [6, 5, 4], [9, 8, 7]]
}
Response:

json
{
  "sorted_arrays": [[1, 2, 3], [4, 5, 6], [7, 8, 9]],
  "time_ns": "<time_taken_in_nanoseconds>"
}
/process-concurrent
Sorts each sub-array concurrently using Go's concurrency features.
##Request:

json code
{
  "to_sort": [[3, 2, 1], [6, 5, 4], [9, 8, 7]]
}

Response:

json code
{
  "sorted_arrays": [[1, 2, 3], [4, 5, 6], [7, 8, 9]],
  "time_ns": "<time_taken_in_nanoseconds>"
}

# Performance Measurement
The server measures the time taken to sort all sub-arrays in nanoseconds using Go's time package.

# Dockerization.Docker public view link:
https://hub.docker.com/r/kumarkhagendra909/sortarray
Build the Docker image:
bash code:
docker build -t sortarray .
Run the Docker container:
bash code:
docker run -p 8000:8000 sortarray
Access the server at http://localhost:8000 and test the endpoints.for /process single after entering your array just like the placeholder hint over the text area click on the Sort Sequentially button and for /process concurrent after entering your array just like the placeholder hint over the text area click on the Sort Concurrently button.

# OUTPUT SCREENSHOT:
===============Browser output===============

![Screenshot (809)](https://github.com/kumarkhagendra909/sorting-server/assets/57476268/9cc947a6-a8a8-471d-92a0-4ee97c17cd5c)
![Screenshot (811)](https://github.com/kumarkhagendra909/sorting-server/assets/57476268/68766376-ca89-40d2-a407-b7dc8f54eb45)
![Screenshot (812)](https://github.com/kumarkhagendra909/sorting-server/assets/57476268/293c4a35-0bcb-4a05-9edc-81800c6e8be5)

===============Command Prompt Output===============

![Screenshot (814)](https://github.com/kumarkhagendra909/sorting-server/assets/57476268/bdfcaaf5-5cac-42db-95d9-fa10e120eb83)

