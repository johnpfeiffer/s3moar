package main

import (
	"context"
	"fmt"
	"os"
	"sync"

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

	maxKeys := int32(1000)

	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := s3.NewFromConfig(cfg)

	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		MaxKeys: maxKeys,
	}

	var wg sync.WaitGroup
	pageCh := make(chan *s3.ListObjectsV2Output)

	paginator := s3.NewListObjectsV2Paginator(client, input)
	for paginator.HasMorePages() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			page, err := paginator.NextPage(context.TODO())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			pageCh <- page
		}()
	}

	go func() {
		wg.Wait()
		close(pageCh)
	}()

	for page := range pageCh {
		for _, object := range page.Contents {
			fmt.Println(*object.Key)
		}
	}
}
