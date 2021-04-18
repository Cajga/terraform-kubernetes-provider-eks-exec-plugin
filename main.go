package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

func main() {
	eksClusterPtr := flag.String("clustername", "", "Name of the cluster")
	flag.Parse()
	if *eksClusterPtr == "" {
		panic("have to provide cluster name with -clustername NAMEOFCLUSTER")
	}

	gen, err := token.NewGenerator(true, false)
	if err != nil {
		panic(err)
	}

	opts := &token.GetTokenOptions{
		ClusterID: aws.StringValue(eksClusterPtr),
	}
	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("{\"kind\": \"ExecCredential\", \"apiVersion\": \"client.authentication.k8s.io/v1alpha1\", \"status\": {\"expirationTimestamp\": \"%v\", \"token\": \"%v\"}}", tok.Expiration.UTC().Format(time.RFC3339), tok.Token)

}
