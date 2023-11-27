
**example of usage**

```
func main() {
	var (
		profile = "default"
		region  = "us-east-1"
		ctx     = context.TODO()
	)

	start := MustLoadConfig(profile, region).NewConfigOpts(WithS3)

	otpS3, err := start.S3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range otpS3.Buckets {
		fmt.Printf("bucket: %s\n", *bucket.Name)
	}
}
```