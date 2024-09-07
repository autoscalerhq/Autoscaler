package jobs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/pricing"
	"gorm.io/gorm"
	"sync"
	"time"
)

// CloudPricing represents the cloud_pricing table schema
type CloudPricing struct {
	ID              uint      `gorm:"primaryKey"`
	Provider        string    `gorm:"size:50;not null"`
	ServiceName     string    `gorm:"size:100;not null"`
	ResourceType    string    `gorm:"size:100"`
	Region          string    `gorm:"size:100"`
	UnitPrice       float64   `gorm:"type:numeric(10,6);not null"`
	Currency        string    `gorm:"size:10;default:'USD'"`
	Unit            string    `gorm:"size:50"`
	Cores           int       `gorm:"type:int"`
	MemoryGB        float64   `gorm:"type:numeric(10,2)"`
	StorageType     string    `gorm:"size:100"`
	PricePerStorage float64   `gorm:"type:numeric(10,6)"`
	RetrievalDate   time.Time `gorm:"autoCreateTime"`
}

// ServiceDetails defines the service-specific information needed for fetching pricing
type ServiceDetails struct {
	ServiceCode string
	ServiceName string
}

// Global variable to hold supported AWS services
var services = map[string]ServiceDetails{
	"EC2":              {ServiceCode: "AmazonEC2", ServiceName: "EC2"},
	"Fargate":          {ServiceCode: "AmazonECS", ServiceName: "Fargate"},
	"RDS":              {ServiceCode: "AmazonRDS", ServiceName: "RDS"},
	"ElasticBeanstalk": {ServiceCode: "ElasticBeanstalk", ServiceName: "Elastic Beanstalk"},
}

func fetchAWSPricing(sess *session.Session, db *gorm.DB) {

	// Get all regions
	regions, err := getAllRegions(sess)
	if err != nil {
		fmt.Println("Error fetching regions:", err)
		return
	}

	// Parallelize fetching for all services across all regions
	var wg sync.WaitGroup
	for _, serviceDetails := range services {
		for _, region := range regions {
			// Get the number of instances in each service
			wg.Add(1)
			go func(serviceDetails ServiceDetails, region string) {
				defer wg.Done()

				fetchPricingForService(sess, db, serviceDetails, region)
			}(serviceDetails, region)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Finished fetching pricing for all regions")
}

// Function to get all available regions from EC2 API, including Wavelength regions
func getAllRegions(sess *session.Session) ([]string, error) {
	svc := ec2.New(sess)
	input := &ec2.DescribeRegionsInput{}
	result, err := svc.DescribeRegions(input)
	if err != nil {
		return nil, err
	}

	regions := []string{}
	for _, region := range result.Regions {
		regions = append(regions, *region.RegionName)
	}

	return regions, nil
}

// Generic function to fetch pricing for any AWS service in a given region
func fetchPricingForService(sess *session.Session, db *gorm.DB, serviceDetails ServiceDetails, region string) {
	svc := pricing.New(sess)

	input := &pricing.GetProductsInput{
		ServiceCode: aws.String(serviceDetails.ServiceCode),
		Filters: []*pricing.Filter{
			{
				Type:  aws.String("TERM_MATCH"),
				Field: aws.String("location"),
				Value: aws.String(region),
			},
		},
	}

	result, err := svc.GetProducts(input)
	if err != nil {
		fmt.Printf("Error fetching pricing for %s in region %s: %v\n", serviceDetails.ServiceName, region, err)
		return
	}

	for _, priceItem := range result.PriceList {
		// Parse priceItem to extract relevant data
		unitPrice := extractUnitPrice(priceItem)
		resourceType := extractInstanceType(priceItem)
		cores := extractCores(priceItem)
		memoryGB := extractMemoryGB(priceItem)

		// Insert pricing data into the database
		pricingEntry := CloudPricing{
			Provider:      "AWS",
			ServiceName:   serviceDetails.ServiceName,
			ResourceType:  resourceType,
			Region:        region,
			UnitPrice:     unitPrice,
			Unit:          "per hour",
			Cores:         cores,
			MemoryGB:      memoryGB,
			RetrievalDate: time.Now(),
		}

		err := db.Create(&pricingEntry).Error
		if err != nil {
			fmt.Printf("Error inserting pricing for %s in region %s, instance type %s: %v\n", serviceDetails.ServiceName, region, resourceType, err)
		} else {
			fmt.Printf("Inserted pricing for %s in region %s, instance type %s\n", serviceDetails.ServiceName, region, resourceType)
		}
	}
}

// You would need to implement these helper functions
func extractUnitPrice(priceItem interface{}) float64 {
	// Parse priceItem and extract the price
	// This is a simplified placeholder
	return 0.01234
}

func extractInstanceType(priceItem interface{}) string {
	// Parse priceItem and extract the instance type
	// This is a simplified placeholder
	return "t3.micro"
}

func extractCores(priceItem interface{}) int {
	// Parse priceItem and extract the number of cores
	// This is a simplified placeholder
	return 2
}

func extractMemoryGB(priceItem interface{}) float64 {
	// Parse priceItem and extract the memory size in GB
	// This is a simplified placeholder
	return 4.0
}
