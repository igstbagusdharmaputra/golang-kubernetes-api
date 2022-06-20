## Setup Environment to Access Kubernetes API

export CLUSTER_NAME="docker-desktop"

export APISERVER=$(kubectl config view -o jsonpath="{.clusters[?(@.name==\"$CLUSTER_NAME\")].cluster.server}")

export TOKEN=$(kubectl get secret golang-default-secret -o jsonpath='{.data.token}' | base64 --decode)

curl -X GET $APISERVER/api --header "Authorization: Bearer $TOKEN" --insecure

kubectl config view --raw -o json | jq -r '.clusters[] | select(.name == "'$(kubectl config current-context)'") | .cluster."certificate-authority-data"' | base64 -d > ca.crt

curl -X GET $APISERVER/api --header "Authorization: Bearer $TOKEN" --cacert ca.crt