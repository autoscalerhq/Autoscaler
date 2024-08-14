package dkron

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	Disabled       bool                         `json:"disabled,omitempty"`
	Tags           map[string]string            `json:"tags,omitempty"`
	Metadata       map[string]string            `json:"metadata,omitempty"`
	Retries        int                          `json:"retries,omitempty"`
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
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create or update job, status code: %d", resp.StatusCode)
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
