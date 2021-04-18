# terraform-kubernetes-provider-eks-exec-plugin
A lightweight, single binary, exec plugin for terraform's kubernetes provider for AWS EKS.

Useful for CI pipelines where we do not want to include a full-blown aws cli

# build a static binary
```bash
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .
```

# usage
```yaml
provider "aws" {
  region = "eu-central-1"
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.example.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.example.certificate_authority.0.data)
  exec {
    api_version = "client.authentication.k8s.io/v1alpha1"
    args        = ["-clustername", local.eks_cluster_name]
    command     = "./terraform-kubernetes-provider-eks-exec-plugin"
  }
}

locals {
  eks_cluster_name = "eks-test"
}

data "aws_eks_cluster" "example" {
  name = local.eks_cluster_name
}

data "kubernetes_all_namespaces" "allns" {}

data "kubernetes_config_map" "example" {
  metadata {
    name      = "aws-auth"
    namespace = "kube-system"
  }
}

output "all-ns" {
  value = data.kubernetes_all_namespaces.allns.namespaces
}

output "configmap" {
  value = data.kubernetes_config_map.example
}

```