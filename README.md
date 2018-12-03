## AWS CloudWatcher Library ##

![GO-Badge](https://goreportcard.com/badge/github.com/florianakos/cloudwatcher)

This is a simple Go library that wraps the "AWS GO SDK" library to provide a way to query AWS Cloudwatch metrics from EC2 instances.

### Example Usage: ###

```go
fmt.Println(cloudwatcher.GetMetricData("aws_region", "your_instance_ID"))
```

### Example output: ###

```go
{
  MetricDataResults: [{
      Id: "id1",
      Label: "CPUUtilization",
      StatusCode: "Complete",
      Timestamps: [
        2018-12-03 14:39:00 +0000 UTC,
        2018-12-03 14:34:00 +0000 UTC,
        2018-12-03 14:29:00 +0000 UTC,
        2018-12-03 14:24:00 +0000 UTC,
        2018-12-03 14:19:00 +0000 UTC,
        2018-12-03 14:14:00 +0000 UTC,
        2018-12-03 14:09:00 +0000 UTC,
        2018-12-03 14:04:00 +0000 UTC,
        2018-12-03 13:59:00 +0000 UTC,
        2018-12-03 13:54:00 +0000 UTC,
        2018-12-03 13:49:00 +0000 UTC
      ],
      Values: [
        1.999435028248588,
        0.0338983050847462,
        0,
        0.032786885245900996,
        0.0338983050847462,
        0,
        0.233333333333334,
        0.0333333333333338,
        0.6595628415300541,
        0.0333333333333332,
        1.240279707326108
      ]
    }]
}

```