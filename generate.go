//go:generate go run pkg/codegen/cleanup/main.go
//go:generate go run pkg/codegen/main.go
//go:generate go run ./pkg/codegen crds ./charts/helm-project-operator-crd/crd-manifest ./charts/helm-project-operator-crd/crds

package main
