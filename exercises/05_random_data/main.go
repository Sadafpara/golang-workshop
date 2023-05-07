package main

/// needs to be checked
import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Random Data API Docs: https://random-data-api.com/documentation

// TODO: Create constant for the base URL: https://random-data-api.com/api/v2/

 const baseURL = "<https://random-data-api.com/api/v2/>"
// TODO: Create a struct for a resource you'd like to fetch
type BloodType struct {
	Type string `json:"type"`
	Donor string `json:"donor"`
	Recipient string `json:"recipient"`
	ID int `json:"id"`
}
// TODO: Add json tags to the struct fields to unmarshal the response
// For example:
// type someResource struct {
// 	ID int `json:"id"`
// }

// TODO: Uncomment the following function when you're ready to use it.
// fetchRandomData fetches random data of the specified size from the provided resource.
// It takes a context for cancellation and timeout, the resource name, and the desired size.
// The function returns the fetched data as a byte slice and any error encountered.

func fetchRandomData(ctx context.Context, resource string, size int) ([]byte, error) {
	// Create a new HTTP GET request for the resource
	req, err := http.NewRequest("GET", baseURL+resource, nil)
	if err != nil {
		return nil, err
	}

	// Add the 'size' query parameter to the request
	query := req.URL.Query()
	query.Add("size", fmt.Sprintf("%d", size))
	req.URL.RawQuery = query.Encode()

	// Create an HTTP client to execute the request
	client := &http.Client{}
	// Execute the request with the provided context
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	// Ensure the response body is closed after the function returns
	defer resp.Body.Close()

	// Read the response body and return it as a byte slice
	return ioutil.ReadAll(resp.Body)
}

func main() {
    // Create a context with a timeout of 30 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Fetch at least 1000 records from the API using fetchRandomData()
    var bloodTypes []BloodType
    for i := 0; i < 10; i++ {
        data, err := fetchRandomData(ctx, BloodType ,  100)
        if err != nil {
            panic(err)
        }
        // Unmarshal the response data into a slice of your struct
        var tmp []BloodType
        if err := json.Unmarshal(data, &tmp); err != nil {
            panic(err)
        }
        bloodTypes = append(bloodTypes, tmp...)
    }

    // Create a CSV file, where you will write the fetched data
    file, err := os.Create("blood_types.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // First line of the CSV file should be the column names
    writer := csv.NewWriter(file)
    defer writer.Flush()
    if err := writer.Write([]string{"Type", "Donor", "Recipient"}); err != nil {
        panic(err)
    }

    // Put the fetched data into a file
    for _, bloodType := range bloodTypes {
        if err := writer.Write([]string{BloodType.Type, BloodType.Donor, BloodType.Recipient}); err != nil {
            panic(err)
        }
    }

    // Measure the total execution time of the program
    fmt.Println("Execution time:", time.Since(time.Now()))
}

	// General TODO: Always handle errors
	// TODO: Measure the total execution time of the program, hint: time.Now() & time.Since()
	// TODO: Create a CSV file, where you will write the fetched data, don't forget to close it, hint: os.Create()

	// TODO: First line of the CSV file should be the column names, hint: file.WriteString()

	// TODO: Create a context with a reasonable timeout

	// TODO: Fetch at least 1000 records from the API using fetchRandomData(), but notice it can only return 100 records at a time
	// 		you might want to put part of the code in a loop
	// 		also, fetchRandomData() expects a context


	// TODO: Unmarshal the response data into a slice of your struct, hint: json.Unmarshal()

	// TODO: Put the fetched data into a file, hint: fmt.Sprintf()

