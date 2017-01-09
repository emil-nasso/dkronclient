package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//DkronJob - TODO
type DkronJob struct {
	Name          string            `json:"name"`
	Schedule      string            `json:"schedule"`
	Command       string            `json:"command"`
	Shell         bool              `json:"shell,omitempty"`
	Owner         string            `json:"owner,omitempty"`
	OwnerEmail    string            `json:"owner_email,omitempty"`
	SuccessCount  int               `json:"success_count,omitempty"`
	ErrorCount    int               `json:"error_count,omitempty"`
	LastSuccess   string            `json:"last_success,omitempty"`
	LastError     string            `json:"last_error,omitempty"`
	Disabled      bool              `json:"disabled,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
	Retries       int               `json:"parent_job,omitempty"`
	ParentJob     string            `json:"parent_job,omitempty"`
	DependentJobs string            `json:"dependent_jobs,omitempty"`
}

//DkronClient - TODO
type DkronClient struct {
	url          string
	debugEnabled bool
}

//NewClient - TODO
func NewClient() (client *DkronClient) {
	return &DkronClient{url: "http://192.168.99.100:32838/v1/", debugEnabled: true}
}

//CreateJob - TODO
func (client *DkronClient) CreateJob(job DkronJob) {
	client.makeRequest("jobs", "POST", job)
}

func (client *DkronClient) makeRequest(endpoint string, method string, payload interface{}) {
	jsonData, _ := json.Marshal(payload)
	if client.debugEnabled {
		fmt.Println("Request data:", string(jsonData))
	}

	var resp *http.Response
	var err error
	switch method {
	case "POST":
		resp, err = http.Post(client.url+endpoint, "application/json", bytes.NewBuffer(jsonData))
	}
	if err != nil {
		log.Fatalf("Could not post: %v", err)
	}
	defer resp.Body.Close()

	if client.debugEnabled {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
}
