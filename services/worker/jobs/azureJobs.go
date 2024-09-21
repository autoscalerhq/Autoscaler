package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/autoscalerhq/autoscaler/internal/bootstrap"
	"gorm.io/gorm"
	"io"
	"net/http"
	"sync"
	"time"
)

// Azure Pricing API and Resource Management API URLs
const azurePricingURL = "https://prices.azure.com/api/retail/prices"
const azureAuthURL = "https://login.microsoftonline.com/%s/oauth2/v2.0/token"
const azureRegionURL = "https://management.azure.com/subscriptions/%s/locations?api-version=2020-01-01"

// Azure credentials - replace with your values
var (
	clientID       = "YOUR_CLIENT_ID"
	clientSecret   = "YOUR_CLIENT_SECRET"
	tenantID       = "YOUR_TENANT_ID"
	subscriptionID = "YOUR_SUBSCRIPTION_ID"
)

// AzureServiceDetails defines the service-specific information needed for fetching pricing
type AzureServiceDetails struct {
	ServiceName string
}

var azureServices = map[string]AzureServiceDetails{
	"VMs":         {ServiceName: "Virtual Machines"},
	"AKS":         {ServiceName: "Azure Kubernetes Service"},
	"AppServices": {ServiceName: "App Services"},
	"SQLDB":       {ServiceName: "Azure SQL Database"}, // New service added
}

// Fetch all Azure regions dynamically using Resource Management API
func getAzureRegions(accessToken string) ([]string, error) {
	url := fmt.Sprintf(azureRegionURL, subscriptionID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println("error closing body")
		}
	}(resp.Body)

	var result struct {
		Value []struct {
			Name string `json:"name"`
		} `json:"value"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	regions := []string{}
	for _, region := range result.Value {
		regions = append(regions, region.Name)
	}

	return regions, nil
}

func fetchAzurePricing() {

	_, _ = bootstrap.GetAzureCredential()

	db, err := bootstrap.GetDBInstance()
	if err != nil {
		fmt.Println("Error getting database session:", err)
	}

	// List of Azure regions (you can define these or fetch them dynamically)
	regions := []string{"eastus", "westus", "westeurope", "southeastasia"}

	// Parallelize fetching for all services across all regions
	var wg sync.WaitGroup
	for _, serviceDetails := range azureServices {
		for _, region := range regions {
			wg.Add(1)
			go func(serviceDetails AzureServiceDetails, region string) {
				defer wg.Done()
				fetchAzurePricingForService(db, serviceDetails, region)
			}(serviceDetails, region)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Finished fetching pricing for all services in all regions")
}

// Function to fetch pricing for any Azure service in a given region
func fetchAzurePricingForService(db *gorm.DB, serviceDetails AzureServiceDetails, region string) {
	// Prepare the API request with filters for service and region
	url := fmt.Sprintf("%s?$filter=serviceName eq '%s' and armRegionName eq '%s'", azurePricingURL, serviceDetails.ServiceName, region)

	// Send the HTTP GET request to the Azure Pricing API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching pricing for %s in region %s: %v\n", serviceDetails.ServiceName, region, err)
		return
	}
	defer resp.Body.Close()

	// Decode the response
	var priceData AzurePricingResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceData); err != nil {
		fmt.Printf("Error decoding Azure pricing data for %s in region %s: %v\n", serviceDetails.ServiceName, region, err)
		return
	}

	// Process each pricing item and insert into the database
	for _, item := range priceData.Items {
		// Parse the pricing item
		unitPrice := item.UnitPrice
		resourceType := item.MeterName
		cores := extractAzureCores(item)
		memoryGB := extractAzureMemory(item)

		// Insert the pricing into the database
		pricingEntry := CloudPricing{
			Provider:      "Azure",
			ServiceName:   serviceDetails.ServiceName,
			ResourceType:  resourceType,
			Region:        region,
			UnitPrice:     unitPrice,
			Unit:          item.UnitOfMeasure,
			Cores:         cores,
			MemoryGB:      memoryGB,
			RetrievalDate: time.Now(),
		}

		err := db.Create(&pricingEntry).Error
		if err != nil {
			fmt.Printf("Error inserting pricing for %s in region %s, resource type %s: %v\n", serviceDetails.ServiceName, region, resourceType, err)
		} else {
			fmt.Printf("Inserted pricing for %s in region %s, resource type %s\n", serviceDetails.ServiceName, region, resourceType)
		}
	}
}

// AzurePricingResponse to unmarshal Azure Pricing API responses
type AzurePricingResponse struct {
	Items []struct {
		MeterName     string  `json:"meterName"`
		UnitPrice     float64 `json:"unitPrice"`
		UnitOfMeasure string  `json:"unitOfMeasure"`
	} `json:"Items"`
}

// Helper functions to extract data from Azure price items
func extractAzureCores(item interface{}) int {
	return 2 // Simplified placeholder
}

func extractAzureMemory(item interface{}) float64 {
	return 4.0 // Simplified placeholder
}
