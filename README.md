# Kubernetes cluster-api-provider-ssh Project

This repository hosts an implementation of a provider using Ironic for the
[cluster-api project](https://sigs.k8s.io/cluster-api).

This project is under heavy development.

# Development notes

## Obtaining the code

```bash
go get github.com/alanmeadows/cluster-api-provider-ironic
cd $GOPATH/src/alanmeadows/cluster-api-provider-ironic
```

## Dependencies/Requirements

- [Kustomize](https://github.com/kubernetes-sigs/kustomize)
- [Go Dep](https://github.com/golang/dep)

## Generating cluster, machine, and provider-components files

```bash
PATH=$PATH:$(go env GOPATH)/bin
dep ensure
make
```

This will build the manager, the CRDs, and finally run the unit tests.

If you would like to `go run` the manager against your `~/.kubeconfig` you can run:

```bash
make run
```

If you want to install the manager and CRDs in your already running cluster (will use `kubectl apply`):

```bash
make install
```

Finally, if you want to build your own image, use the `docker-build` and `docker-push` commands:

```bash
export IMG=quay.io/alanmeadows/cluster-api-provider-ironic:v1.0
make docker-build IMG=${IMG}
make docker-push IMG=${IMG}
```


