/*
short programme to list resources in all aws regions
*/
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {


	regions := []string{"us-east-1",

		"us-west-1",
		"us-east-1",
		"us-west-2",
		"eu-west-1",
		"eu-central-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"sa-east-1",
		//"cn-north-1",
		//"us-gov-west-1",
	}





	for a := range regions {

		//must have your credentials file set up
		svc := ec2.New(session.New(), &aws.Config{Region: aws.String(regions[a])})

		// Call the DescribeInstances Operation
		resp, err := svc.DescribeInstances(nil)
		if err != nil {
			panic(err)
		}

		// resp has all of the response data, pull out instance IDs:
		fmt.Println("> Number of reservation sets: ", regions[a], len(resp.Reservations))
		for idx, res := range resp.Reservations {
			fmt.Println("  > Number of instances: ", len(res.Instances))
			for _, inst := range resp.Reservations[idx].Instances {
				fmt.Println("    - Instance ID: ", *inst.InstanceId)
			}
		}


	}
}
