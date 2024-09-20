package bootstrap

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsLock = &sync.Mutex{}
var awsSession *session.Session

// GetAWSSession returns a singleton instance of an AWS session
func GetAWSSession() (*session.Session, error) {
	if awsSession == nil {
		awsLock.Lock()
		defer awsLock.Unlock()
		if awsSession == nil {
			fmt.Println("Creating single AWS session instance now.")

			sess, err := session.NewSession(&aws.Config{
				Region: aws.String("us-east-1"), // Change the region as per your requirement
			})

			if err != nil {
				fmt.Println("Error creating AWS session:", err)
				return nil, err
			}

			awsSession = sess
		} else {
			fmt.Println("AWS session instance already created.")
		}
	} else {
		fmt.Println("AWS session instance already created.")
	}

	return awsSession, nil
}
