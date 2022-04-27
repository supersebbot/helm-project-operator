package main

import (
	"os"

	"github.com/rancher/wrangler/pkg/cleanup"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cleanup.Cleanup("./pkg/apis"); err != nil {
		logrus.Fatal(err)
	}
	if err := os.RemoveAll("./pkg/generated"); err != nil {
		logrus.Fatal(err)
	}
	if err := os.RemoveAll("./charts/helm-project-operator-crd/crds"); err != nil {
		logrus.Fatal(err)
	}
	if err := os.RemoveAll("./charts/helm-project-operator-crd/crd-manifest"); err != nil {
		logrus.Fatal(err)
	}
}
