package utils

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)
func InitAWSSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIAQLHQ3QWM6HKISM6B", "BbBbMPei6IsGSb3dF08s0TtrBUv9HRjwgdOKto6B", ""), //access key, secret key, token
	})

	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	return sess
}
func InitS3Client(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

//Asia Pacific (Singapore) ap-southeast-1