module github.com/Cajga/terraform-kubernetes-provider-eks-exec-plugin

go 1.15

require (
	github.com/aws/aws-sdk-go v1.38.21
	github.com/aws/aws-sdk-go-v2/config v1.1.6
	github.com/aws/aws-sdk-go-v2/service/eks v1.2.2
	github.com/aws/aws-sdk-go-v2/service/s3 v1.5.0 // indirect
	sigs.k8s.io/aws-iam-authenticator v0.5.2
)
