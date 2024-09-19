package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GCPServiceDetails defines the service-specific information needed for fetching pricing
type GCPServiceDetails struct {
	ServiceName string
}

// Google Cloud Pricing API URL
const gcpPricingURL = "https://cloudpricingcalculator.appspot.com/static/data/pricelist.json"

// Supported GCP services
var gcpServices = map[string]GCPServiceDetails{
	"ComputeEngine": {ServiceName: "Compute Engine"},
	"GKE":           {ServiceName: "Google Kubernetes Engine"},
	"AppEngine":     {ServiceName: "App Engine"},
}

func main() {
	// Connect to the database
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto-migrate the schema to the PostgreSQL database
	db.AutoMigrate(&CloudPricing{})

	// Authenticate with GCP and fetch the list of available regions
	regions, err := getGCPRegions()
	if err != nil {
		fmt.Println("Error fetching GCP regions:", err)
		return
	}

	// Parallelize fetching for all services across all regions
	var wg sync.WaitGroup
	for _, serviceDetails := range gcpServices {
		for _, region := range regions {
			wg.Add(1)
			go func(serviceDetails GCPServiceDetails, region string) {
				defer wg.Done()
				fetchGCPPricingForService(db, serviceDetails, region)
			}(serviceDetails, region)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Finished fetching pricing for all services in all regions")
}

// Fetch all available regions from GCP
func getGCPRegions() ([]string, error) {
	ctx := context.Background()

	// Initialize the Compute Engine client
	client, err := compute.NewService(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		return nil, err
	}

	// List available regions
	regionsService := compute.NewRegionsService(client)
	regionList, err := regionsService.List("your-project-id").Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	regions := []string{}
	for _, region := range regionList.Items {
		regions = append(regions, region.Name)
	}

	return regions, nil
}

// Fetch pricing for any GCP service in a given region
func fetchGCPPricingForService(db *gorm.DB, serviceDetails GCPServiceDetails, region string) {
	// Use the static GCP pricing data (Alternatively, you could use a custom pricing API if available)
	resp, err := http.Get(gcpPricingURL)
	if err != nil {
		fmt.Printf("Error fetching GCP pricing for %s in region %s: %v\n", serviceDetails.ServiceName, region, err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var pricingData map[string]interface{}
	if err := json.Unmarshal(body, &pricingData); err != nil {
		fmt.Printf("Error decoding GCP pricing data for %s in region %s: %v\n", serviceDetails.ServiceName, region, err)
		return
	}

	// Extract relevant pricing data
	// GCP pricing data has a complex structure, so this is a simplified example
	pricingItems := pricingData["gcp_price_list"].(map[string]interface{})["CP-COMPUTEENGINE-VMIMAGE-N1-STANDARD-1"].(map[string]interface{})

	unitPrice := pricingItems["us"].(float64) // Placeholder for example
	resourceType := "n1-standard-1"           // Placeholder for example
	cores := 1                                // Simplified placeholder
	memoryGB := 3.75                          // Simplified placeholder

	// Insert the pricing into the database
	pricingEntry := CloudPricing{
		Provider:      "GCP",
		ServiceName:   serviceDetails.ServiceName,
		ResourceType:  resourceType,
		Region:        region,
		UnitPrice:     unitPrice,
		Unit:          "per hour",
		Cores:         cores,
		MemoryGB:      memoryGB,
		RetrievalDate: time.Now(),
	}

	err = db.Create(&pricingEntry).Error
	if err != nil {
		fmt.Printf("Error inserting pricing for %s in region %s, resource type %s: %v\n", serviceDetails.ServiceName, region, resourceType, err)
	} else {
		fmt.Printf("Inserted pricing for %s in region %s, resource type %s\n", serviceDetails.ServiceName, region, resourceType)
	}
}
