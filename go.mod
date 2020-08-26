module github.com/rancher/terraform-controller

go 1.13

require (
	github.com/docker/go-units v0.4.0
	github.com/hashicorp/go-getter v1.4.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.7.0 // indirect
	github.com/onsi/gomega v1.4.3 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rancher/wrangler v0.2.0
	github.com/rancher/wrangler-api v0.2.0
	github.com/sirupsen/logrus v1.4.1
	github.com/spf13/pflag v1.0.2 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/urfave/cli v1.20.0
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/gengo v0.0.0-20190327210449-e17681d19d3a
)

replace github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
