package dkron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Job represents a scheduled task to execute.
type Job struct {
	Name           string                       `json:"name"`
	DisplayName    string                       `json:"displayname,omitempty"`
	Schedule       string                       `json:"schedule"`
	Timezone       string                       `json:"timezone,omitempty"`
	Owner          string                       `json:"owner,omitempty"`
	OwnerEmail     string                       `json:"owner_email,omitempty"`
	SuccessCount   int                          `json:"success_count,omitempty"`
	ErrorCount     int                          `json:"error_count,omitempty"`
	LastSuccess    string                       `json:"last_success,omitempty"`
	LastError      string                       `json:"last_error,omitempty"`
	Disabled       bool                         `json:"disabled"`
	Ephemeral      bool                         `json:"ephemeral"`
	Tags           map[string]string            `json:"tags,omitempty"`
	Metadata       map[string]string            `json:"metadata,omitempty"`
	Retries        *int                         `json:"retries,omitempty"`
	ParentJob      string                       `json:"parent_job,omitempty"`
	DependentJobs  []string                     `json:"dependent_jobs,omitempty"`
	Processors     map[string]map[string]string `json:"processors,omitempty"`
	Concurrency    string                       `json:"concurrency,omitempty"`
	Executor       string                       `json:"executor"`
	ExecutorConfig map[string]string            `json:"executor_config"`
	Status         string                       `json:"status,omitempty"`
	Next           string                       `json:"next,omitempty"`
}

// Client is the Dkron API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new Dkron API client.
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// CreateOrUpdateJob creates or updates a job.
func (c *Client) CreateOrUpdateJob(job Job, runOnCreate bool) (*Job, error) {
	url := fmt.Sprintf("%s/jobs?runoncreate=%t", c.BaseURL, runOnCreate)
	jobData, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jobData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println("failed to close body: ", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		}

		// Convert the body to a string
		bodyString := string(body)
		fmt.Println(bodyString)

		return nil, fmt.Errorf("failed to create or update job, status code: %d, %s", resp.StatusCode, bodyString)
	}

	var createdJob Job
	err = json.NewDecoder(resp.Body).Decode(&createdJob)
	if err != nil {
		return nil, err
	}

	return &createdJob, nil
}

// ToggleJob toggles a job.
func (c *Client) ToggleJob(jobName string) (*Job, error) {
	url := fmt.Sprintf("%s/jobs/%s/toggle", c.BaseURL, jobName)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to toggle job, status code: %d", resp.StatusCode)
	}

	var toggledJob Job
	err = json.NewDecoder(resp.Body).Decode(&toggledJob)
	if err != nil {
		return nil, err
	}

	return &toggledJob, nil
}

// DeleteJob deletes a job.
func (c *Client) DeleteJob(jobName string) error {
	url := fmt.Sprintf("%s/jobs/%s", c.BaseURL, jobName)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete job, status code: %d", resp.StatusCode)
	}

	return nil
}

// RunJob runs a job.
func (c *Client) RunJob(jobName string) (*Job, error) {
	url := fmt.Sprintf("%s/jobs/%s", c.BaseURL, jobName)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("failed to run job, status code: %d", resp.StatusCode)
	}

	var runJob Job
	err = json.NewDecoder(resp.Body).Decode(&runJob)
	if err != nil {
		return nil, err
	}

	return &runJob, nil
}

// ShowJobByName shows a job by its name.
func (c *Client) ShowJobByName(jobName string) (*Job, error) {
	url := fmt.Sprintf("%s/jobs/%s", c.BaseURL, jobName)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to show job, status code: %d", resp.StatusCode)
	}

	var job Job
	err = json.NewDecoder(resp.Body).Decode(&job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

// GetStatus retrieves the current status of the Dkron node.
func (c *Client) GetStatus() (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/", c.BaseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println("failed to close body: ", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get status, status code: %d", resp.StatusCode)
	}

	var status map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		return nil, err
	}

	return status, nil
}

// GetLeader retrieves the current leader of the Dkron cluster.
func (c *Client) GetLeader() (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/leader", c.BaseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get leader, status code: %d", resp.StatusCode)
	}

	var leader map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&leader)
	if err != nil {
		return nil, err
	}

	return leader, nil
}

// GetBusy retrieves the currently running executions.
func (c *Client) GetBusy() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/busy", c.BaseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get busy executions, status code: %d", resp.StatusCode)
	}

	var executions []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&executions)
	if err != nil {
		return nil, err
	}

	return executions, nil
}
