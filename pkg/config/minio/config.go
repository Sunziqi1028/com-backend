package minio

/// will init at the compile time with the CI environment

var (
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
)
