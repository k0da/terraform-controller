module github.com/rancher/terraform-controller

go 1.13

require (
	github.com/docker/go-units v0.4.0
	github.com/google/go-github/v28 v28.0.0
	github.com/gorilla/mux v1.7.1
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rancher/lasso v0.0.0-20200905045615-7fcb07d6a20b
	github.com/rancher/wrangler v0.7.3-0.20201020003736-e86bc912dfac
	github.com/rancher/wrangler-api v0.6.1-0.20200515193802-dcf70881b087
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	github.com/urfave/cli v1.22.4
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	golang.org/x/tools v0.0.0-20201017001424-6003fad69a88 // indirect
	k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver v0.18.0
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v0.18.8
	k8s.io/gengo v0.0.0-20200114144118-36b2048a9120
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451 // indirect
)

replace github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
