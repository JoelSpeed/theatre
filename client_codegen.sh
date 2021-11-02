ROOT_PACKAGE="github.com/gocardless/theatre"
CUSTOM_RESOURCE_NAME="workloads"
CUSTOM_RESOURCE_VERSION="v1alpha1"

//go get -u k8s.io/code-generator/...
cd $GOPATH/src/k8s.io/code-generator

./generate-groups.sh all "$ROOT_PACKAGE/pkg/client" "$ROOT_PACKAGE/apis/workloads" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

tree $GOPATH/src/$ROOT_PACKAGE/pkg/client
