package s3buckets

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetBodyJsonFromBucket(bucketname string, keyname string, region string) (map[string]interface{}, error) {
	// get data from s3 bucket
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svcS3 := s3.New(sess)

	// get object from s3 bucket
	result, err := svcS3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(keyname),
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer result.Body.Close()

	var resultJson map[string]interface{}

	err = json.NewDecoder(result.Body).Decode(&resultJson)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resultJson, nil
}
