package main

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// Config 存储 Casdoor 配置信息
type Config struct {
	Endpoint         string
	ClientId         string
	ClientSecret     string
	Certificate      string
	OrganizationName string
	ApplicationName  string
}

// GetDefaultConfig 返回默认配置
func GetDefaultConfig() *Config {
	return &Config{
		Endpoint:         "http://localhost:8000",
		ClientId:         "294b09fbc17f95daf2fe",
		ClientSecret:     "dd8982f7046ccba1bbd7851d5c1ece4e52bf039d",
		OrganizationName: "built-in",
		ApplicationName:  "app-built-in",
		// 证书用于验证 JWT token
		Certificate: `-----BEGIN CERTIFICATE-----
MIIE+TCCAuGgAwIBAgIDAeJAMA0GCSqGSIb3DQEBCwUAMDYxHTAbBgNVBAoTFENh
c2Rvb3IgT3JnYW5pemF0aW9uMRUwEwYDVQQDEwxDYXNkb29yIENlcnQwHhcNMjEw
OTIyMTE1NzE3WhcNNDEwOTIyMTE1NzE3WjA2MR0wGwYDVQQKExRDYXNkb29yIE9y
Z2FuaXphdGlvbjEVMBMGA1UEAxMMQ2FzZG9vciBDZXJ0MIICIjANBgkqhkiG9w0B
AQEFAAOCAg8AMIICCgKCAgEAsInpb5E1/ym0f1RfSDSSE8IR7y+lw+RJjI74e5ej
rq4b8zMYk7HeHCMJXXUmgPLWcP8JyK6Gw5KKs/KN+YN8wYp0x7w8nP/PdnOWMuHN
a7bKkPNzRdkVGK0d8dFPWm2z9iJxX8VdF/OYy+rV0IbJYaLGZPw3RhS5tSd7QkHi
3dBN1f3T8BIKx/8vP8Yr7WsJBp7wQXCaJc/hB7VzNKSdC1v0Uw6LJ3TJLn/xvp0h
X8MKqjSZBSVqMBKIGF8VF9zPBHT8n0D0xn1/t6zYZmZmN8KMQ3vMUEBfJhWJMlVR
oMQ0xEEDy0L5qGa0cqHCqVUhWiGPNL5x9dTEtX1cQGVJCPXB9cpN6yEwmVqNgGpH
HqGnK0eCF5p7iHLNEFVCbPCgP7zKRMm3fQvPFxvCT9vxgTdPPKiPGCWFxLJh0vFO
8G3k8aZQFqJKKMqJjfQBRHvMDQKpQxYeLqWcqV7Z3B3GvGHZOGfZSHGKGHJC6KLJ
NyZxMGLiR7XQV2jNyT3xH5OTL9HV5CdJEGLl1fKCwvPCJ7Y6J4m8cQNvH4DQFGVp
T7QGH6hNrE5qFHrNVvXzrNxNGmLYlN2FJHC2hDHYuGJXEJ7N5eiD5aDN6T9JZvWE
8iIrLvGP1dQXQxhz5vMVnEJ9MJZxbpjHJPTgH/VTQlVV9n5FMlMNLfLHVH+6Mf3f
QbMCAwEAAaMpMCcwDgYDVR0PAQH/BAQDAgWgMBUGA1UdJQQOMAwGCisGAQQBgjcK
AwwwDQYJKoZIhvcNAQELBQADggIBAAlTB5FQGGghUoFKpW9T1CQ3s0I8Y8PQGPC5
qJZGGMVvZLqXCQ7lJFD9nLHH1Dw6L0lHPHJPQpxvf+z6Tpvn8xKKVZ5cJqXZBBJK
8XQvSjLKQB8cCqLGHKLJFr4bHGDYvfQHqFGX4y0rMn/gHqVNTJiJH6k+u3PpQH0J
tHj7vDkHqYq5LrNJ0H5lJcG6QH3MlBvQXL1P2LqpCxKzNYj3P8Lm4xGQHqL+TLHl
GXfFkHhPb3Q0hT6HKPnl+J2bvXLqQJd4YPqJHl6JGqQ2vZvKL7hPH8QLQY5hXJHq
L6Q7JqL6qZJ8qLH6Q7qJqLqJqLqJqLqJqLqJqLqJqLqJqLqJqLqJqLqJqLqJqLq
-----END CERTIFICATE-----`,
	}
}

// InitCasdoorSDK 初始化 Casdoor SDK 客户端
func InitCasdoorSDK(config *Config) {
	casdoorsdk.InitConfig(
		config.Endpoint,
		config.ClientId,
		config.ClientSecret,
		config.Certificate,
		config.OrganizationName,
		config.ApplicationName,
	)
}
