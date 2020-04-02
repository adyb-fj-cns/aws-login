package aws

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// Service struct
type Service struct {
	Profile      string
	MFAARN       string
	MFACode      string
	Region       string
	AccessKey    string
	SecretKey    string
	SessionToken string
}

// NewService returns new instance
func NewService(profile string, mfaARN string, mfaCode string, region string) *Service {
	return &Service{
		Profile: profile,
		MFAARN:  mfaARN,
		MFACode: mfaCode,
		Region:  region,
	}
}

// New returns new instance
func New() *Service {
	return &Service{
		Profile: "default",
	}
}

// Init initializes
func (t *Service) Init() {

}

// GenerateTemporaryCredentialsFromSTS returns from STS with MFA
func (t *Service) GenerateTemporaryCredentialsFromSTS() {

	svc := sts.New(session.New())
	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(900),
		SerialNumber:    aws.String(t.MFAARN),
		TokenCode:       aws.String(t.MFACode),
	}

	result, err := svc.GetSessionToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	t.AccessKey = *result.Credentials.AccessKeyId
	t.SecretKey = *result.Credentials.SecretAccessKey
	t.SessionToken = *result.Credentials.SessionToken

	if runtime.GOOS == "windows" {
		outputWindowsAWSCredentials(t.Profile, t.Region, t.AccessKey, t.SecretKey, t.SessionToken)
	} else {
		outputAWSCredentials(t.Profile, t.Region, t.AccessKey, t.SecretKey, t.SessionToken)
	}

}

func outputWindowsAWSCredentials(profile string, region string, accessKey string, secretKey string, sessionToken string) {
	fmt.Printf("$Env:AWS_ACCESS_KEY_ID = \"%s\"\n", accessKey)
	fmt.Printf("$Env:AWS_SECRET_ACCESS_KEY = \"%s\"\n", secretKey)
	fmt.Printf("$Env:AWS_SESSION_TOKEN = \"%s\"\n", sessionToken)
	fmt.Printf("$Env:AWS_REGION = \"%s\"\n", region)
}

func outputAWSCredentials(profile string, region string, accessKey string, secretKey string, sessionToken string) {
	// Not tested yet
	fmt.Printf("aws configure set aws_access_key_id %s --profile %s;\n", accessKey, profile)
	fmt.Printf("aws configure set aws_secret_access_key %s --profile %s;\n", secretKey, profile)
	fmt.Printf("aws configure set region %s --profile %s;\n", region, profile)
	fmt.Printf("aws configure set format json --profile %s;\n", profile)
}

func setLinuxAWSCredentials(profile string, region string, accessKey string, secretKey string) {
	// Not tested yet
	_, err := exec.Command("/usr/local/bin/aws", "configure", "set", "aws_access_key_id", accessKey, "--profile", profile).Output()
	_, err = exec.Command("/usr/local/bin/aws", "configure", "set", "aws_secret_access_key", secretKey, "--profile", profile).Output()
	_, err = exec.Command("/usr/local/bin/aws", "configure", "set", "region", region, "--profile", profile).Output()
	_, err = exec.Command("/usr/local/bin/aws", "configure", "set", "format", "json", "--profile", profile).Output()

	if err != nil {
		panic(err)
	}
	fmt.Println("Waiting for credentials to be ready...")
	time.Sleep(5 * time.Second)
	fmt.Println("Please open a new terminal window to refresh the updated default environment variables.")
}

func setWindowsAWSCredentials(profile string, region string, accessKey string, secretKey string) {
	// Not tested yet
	//_, err := exec.Command("setx", "AWS_DEFAULT_PROFILE", profile).Output()
	//_, err = exec.Command("setx", "AWS_PROFILE", profile).Output()
	_, err := exec.Command("aws", "configure", "set", "aws_access_key_id", accessKey, "--profile", profile).Output()
	_, err = exec.Command("aws", "configure", "set", "aws_secret_access_key", secretKey, "--profile", profile).Output()
	_, err = exec.Command("aws", "configure", "set", "region", region, "--profile", profile).Output()
	_, err = exec.Command("aws", "configure", "set", "format", "json", "--profile", profile).Output()

	if err != nil {
		panic(err)
	}
	fmt.Println("Waiting for credentials to be ready...")
	time.Sleep(5 * time.Second)
	fmt.Println("Please open a new terminal window to refresh the updated default environment variables.")
}
