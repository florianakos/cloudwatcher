package cloudwatcher

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"

	"log"
	"strconv"
	"time"
)

// global variable needed to generate unique request IDs
// in case multiple metrics are requested
var counter int64

// build data structure to request the metricname from the given instance ID
func buildMetricDataQuery(metricname, instanceID string) *cloudwatch.MetricDataQuery {
	counter++
	return &cloudwatch.MetricDataQuery{
		Id: aws.String("id" + strconv.FormatInt(counter, 10)),
		MetricStat: &cloudwatch.MetricStat{
			Period: aws.Int64(60),
			Stat:   aws.String("Average"),
			Metric: &cloudwatch.Metric{
				MetricName: aws.String(metricname),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("InstanceId"),
						Value: aws.String(instanceID),
					},
				},
				Namespace: aws.String("AWS/EC2"),
			},
		},
	}
}

// GetMetricData is a func to get CPU-Utilization for the last hour for the instance in the specified region
func GetMetricData(region string, instanceID string) *cloudwatch.GetMetricDataOutput {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Println(err)
		return nil
	}
	cw := cloudwatch.New(sess)

	// prepare request input structure
	dataInput := &cloudwatch.GetMetricDataInput{
		StartTime: aws.Time(time.Now().Add(-60 * time.Minute)),
		EndTime:   aws.Time(time.Now()),
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			buildMetricDataQuery("CPUUtilization", instanceID),
			//buildMetricDataQuery("NetworkIn", instanceID),

		},
	}

	// request MetricData in batch
	dataOutput, err := cw.GetMetricData(dataInput)
	if err != nil {
		log.Println("error GetMetricStatistics: ", err)
		return nil
	}

	// return results
	return dataOutput
}
