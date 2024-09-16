package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/asecurityteam/rolling"
	internalMath "github.com/autoscalerhq/autoscaler/lib/math"
	"github.com/kevinconway/loadshed/v2"
	loadshedhttp "github.com/kevinconway/loadshed/v2/stdlib/net/http"
	"github.com/shirou/gopsutil/cpu"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
)

var CPUMetric uint32 = 0

var rollingWindow = rolling.NewPointPolicy(rolling.NewWindow(10))

// CPUCapacity implements the Capacity interface based on CPU usage
type CPUCapacity struct{}

func (c *CPUCapacity) Name(ctx context.Context) string {
	return "CPU Usage"
}

func (c *CPUCapacity) Usage(ctx context.Context) float32 {
	return internalMath.Uint32ToFloat32(atomic.LoadUint32(&CPUMetric))
}

// GetCPUUsage calculates the CPU usage percentage and updates the rolling window and CPUMetric variables.
// It fetches the CPU usage percentages using the cpu.Percent function. If there is an error, it prints the error
// message and sets the returnPercentage to 0.0. Otherwise, it calculates the mean and median of the percentages.
// If mean > median, it returns the higher of the two. If mean < median, it returns the lower of the two.
// If mean == median, it returns the median. The calculated value is then appended to the rolling window.
// Finally, it reduces the rolling window to get the average and assigns it to the CPUMetric variable.
// Decision was based on https://gist.github.com/dim/152e6bf80e1384ea72e17ac717a5000a
func GetCPUUsage() {

	var returnPercentage = 0.0
	percentages, err := cpu.Percent(0, true)

	if err != nil {
		fmt.Println("Error fetching CPU usage:", err)
		returnPercentage = 0.0
	} else {

		mean := internalMath.Mean(percentages)
		median := internalMath.Median(percentages)

		if mean > median {
			// Indicates that there are some high outliers pulling the mean upwards. Return the higher of the two.
			returnPercentage = internalMath.Max(median, mean)
		} else if mean < median {
			// Indicates that there are some low outliers pulling the mean downward. Return the lower of the two.\
			returnPercentage = internalMath.Min(median, mean)
		} else {
			// Default to median as its more accuracy as saying what is the overall usage in the CPU
			returnPercentage = median
		}
	}

	rollingWindow.Append(returnPercentage)

	atomic.StoreUint32(&CPUMetric, uint32(rollingWindow.Reduce(rolling.Avg)))
}

// CPUFailureProbability implements the FailureProbability interface based on CPU usage
type CPUFailureProbability struct {
	*CPUCapacity
}

func (c *CPUFailureProbability) Likelihood(ctx context.Context) float32 {

	println("usage: ", int(CPUMetric))
	if CPUMetric < 80 {
		// No chance of failure due to CPU when under 80% utilization
		return 0.0
	}
	return float32((CPUMetric-80)/20) / 100 // Linear increase from 80% to 100% usage
}

const (
	PriorityHigh   loadshed.Classification = "HIGH"
	PriorityNormal loadshed.Classification = "NORMAL"
	PriorityLow    loadshed.Classification = "LOW"
)

func (*URLClassifier) Classify(ctx context.Context) loadshed.Classification {
	path, ok := ctx.Value(ctxKey).(string)
	println("path: ", path)
	if !ok {
		return PriorityNormal
	}

	// Example logic: Assign priority based on URL path
	if strings.HasPrefix(path, "/health-check") {
		println("high-priority route")
		return PriorityHigh
	}
	if strings.HasPrefix(path, "/api/v1/normal") {
		println("normal-priority route")
		return PriorityNormal
	}

	println("low-priority route")
	return PriorityLow
}

// RejectionHandler is a custom handler for load shed rejections.
type RejectionHandler struct{}

// ServeHTTP implements the http.Handler interface and handles rejected requests.
func (rh *RejectionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var rejectionInfo loadshed.ErrRejection
	println("responding value usage: ", rejectionInfo.Usage*100)
	if err := loadshedhttp.FromHandlerContext(r.Context()); errors.As(err, &rejectionInfo) {
		// Set custom headers with the rejection details
		w.Header().Set("X-Rejection-Name", rejectionInfo.Name)
		w.Header().Set("X-Rejection-Usage", strconv.Itoa(int(rejectionInfo.Usage*100)))
		w.Header().Set("X-Rejection-Rule", rejectionInfo.Rule)
		w.Header().Set("X-Rejection-Classifier", string(rejectionInfo.Classification))
	}

	// Respond with 503 Service Unavailable
	http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
}

func CreateShedder() *loadshed.Shedder {
	cpuCap := &CPUCapacity{}
	cpuFailure := &CPUFailureProbability{CPUCapacity: cpuCap}

	// Create a rejection rate based on the failure probability
	rejectionRate := loadshed.NewRejectionRateCurveIdentity(cpuFailure)

	// Create the classifier
	classifier := &URLClassifier{}

	// Create the Shedder with the rejection rate as a probabilistic policy
	return loadshed.NewShedder(
		loadshed.OptionShedderClassifier(classifier),
		loadshed.OptionShedderRejectionRate(rejectionRate),
	)
}
