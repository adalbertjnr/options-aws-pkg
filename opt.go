package apkg

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ebs"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type OptFunc func(*Options)

type (
	Config struct {
		Setup
		Options
	}

	Setup struct {
		Profile string
		Region  string
	}

	Options struct {
		AwsCfg    aws.Config
		R53Client *route53.Client
		S3Client  *s3.Client
		IamClient *iam.Client
		EcrClient *ecr.Client
		SsmClient *ssm.Client
		EbsClient *ebs.Client
	}
)

// Set as opts: WithIam, WithS3, WithEcr, WithR53, WithSSM
func (c *Config) NewConfigOpts(opts ...OptFunc) *Config {
	for _, optFn := range opts {
		optFn(&c.Options)
	}
	return c
}

func MustLoadConfig(profile, region string) *Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Setup: Setup{
			Profile: profile,
			Region:  region,
		},
		Options: Options{
			AwsCfg: cfg,
		},
	}
}

func WithIam(opt *Options) {
	opt.IamClient = iam.NewFromConfig(opt.AwsCfg)
}

func WithS3(opt *Options) {
	opt.S3Client = s3.NewFromConfig(opt.AwsCfg)
}

func WithR53(opt *Options) {
	opt.R53Client = route53.NewFromConfig(opt.AwsCfg)
}

func WithEcr(opt *Options) {
	opt.EcrClient = ecr.NewFromConfig(opt.AwsCfg)
}

func WithSSM(opt *Options) {
	opt.SsmClient = ssm.NewFromConfig(opt.AwsCfg)
}

func WithEBS(opt *Options) {
	opt.EbsClient = ebs.NewFromConfig(opt.AwsCfg)
}
