package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Bucket name required as argument")
		os.Exit(1)
	}
	bucket := os.Args[1]

	cfg, _ := config.LoadDefaultConfig(context.TODO())
	// config.WithRegion("us-west-2") is only necessary if you are not using the default region.
	client := s3.NewFromConfig(cfg)

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}

	paginator := s3.NewListObjectsV2Paginator(client, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, object := range page.Contents {
			fmt.Println(*object.Key)
		}
	}
}
