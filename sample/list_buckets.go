package sample

import (
	"fmt"
	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
)

func ListBucketsSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	myBuckets := []string{
		"yig-my-bucket-1",
		"yig-my-bucket-11",
		"yig-my-bucket-2",
		"yig-my-bucket-21",
		"yig-my-bucket-22",
		"yig-my-bucket-3",
		"yig-my-bucket-31",
		"yig-my-bucket-32"}
	for _, b := range myBuckets {
		err := sc.MakeBucket(b)
		if err != nil {
			HandleError(err)
		}
	}
	out, err := sc.ListBuckets()
	if err != nil {
		HandleError(err)
	}
	for _, b := range out {
		fmt.Println("bucket:", *b.Name)
	}

	for _, b := range myBuckets {
		err := sc.DeleteBucket(b)
		if err != nil {
			HandleError(err)
		}
	}

	fmt.Printf("ListBucketsSample Run Success!\n\n")
}
