package config

import (
	"errors"
	"os"
	"time"
)

type Env struct {
	HttpPort    string `map:"HTTP_PORT"`
	FrontendUrl string `map:"FRONTEND_URL"`

	DbUsername string `map:"DB_USERNAME"`
	DbPassword string `map:"DB_PASSWORD"`
	DbDb       string `map:"DB_DB"`
	DbHost     string `map:"DB_HOST"`
	DbPort     string `map:"DB_PORT"`
	DbSslMode  string `map:"DB_SSL_MODE"`

	AccessTokenPrivateKey string        `map:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey  string        `map:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiredIn  time.Duration `map:"ACCESS_TOKEN_EXPIRED_IN"`

	AwsS3Region        string `map:"AWS_S3_REGION"`
	AwsS3Bucket        string `map:"AWS_S3_BUCKET"`
	AwsAccessKeyID     string `map:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `map:"AWS_SECRET_ACCESS_KEY"`
}

func NewEnv() (*Env, error) {
	env := Env{}

	httpPort, err := getEnv("HTTP_PORT")
	if err != nil {
		return &env, err
	}
	env.HttpPort = httpPort

	frontendUrl, err := getEnv("FRONTEND_URL")
	if err != nil {
		return &env, err
	}
	env.FrontendUrl = frontendUrl

	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		return &env, err
	}
	env.DbHost = dbHost

	dbUsername, err := getEnv("DB_USERNAME")
	if err != nil {
		return &env, err
	}
	env.DbUsername = dbUsername

	dbPassword, err := getEnv("DB_PASSWORD")
	if err != nil {
		return &env, err
	}
	env.DbPassword = dbPassword

	dbDb, err := getEnv("DB_DB")
	if err != nil {
		return &env, err
	}
	env.DbDb = dbDb

	dbPort, err := getEnv("DB_PORT")
	if err != nil {
		return &env, err
	}
	env.DbPort = dbPort

	sslMode, err := getEnv("DB_SSL_MODE")
	if err != nil {
		return &env, err
	}
	env.DbSslMode = sslMode

	accessTokenPrivateKey, err := getEnv("ACCESS_TOKEN_PRIVATE_KEY")
	if err != nil {
		return &env, err
	}
	env.AccessTokenPrivateKey = accessTokenPrivateKey

	accessTokenPublicKey, err := getEnv("ACCESS_TOKEN_PUBLIC_KEY")
	if err != nil {
		return &env, err
	}
	env.AccessTokenPublicKey = accessTokenPublicKey

	duration, err := getEnv("ACCESS_TOKEN_EXPIRED_IN")
	if err != nil {
		return &env, err
	}

	accessTokenExpiredIn, err := time.ParseDuration(duration)
	if err != nil {
		return &env, err
	}
	env.AccessTokenExpiredIn = accessTokenExpiredIn

	awsS3Region, err := getEnv("AWS_S3_REGION")
	if err != nil {
		return &env, err
	}
	env.AwsS3Region = awsS3Region

	awsS3Bucket, err := getEnv("AWS_S3_BUCKET")
	if err != nil {
		return &env, err
	}
	env.AwsS3Bucket = awsS3Bucket

	awsAccessKeyID, err := getEnv("AWS_ACCESS_KEY_ID")
	if err != nil {
		return &env, err
	}
	env.AwsAccessKeyID = awsAccessKeyID

	awsSecretAccessKey, err := getEnv("AWS_SECRET_ACCESS_KEY")
	if err != nil {
		return &env, err
	}
	env.AwsSecretAccessKey = awsSecretAccessKey

	return &env, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New(key + " is not set")
	}
	return value, nil
}
