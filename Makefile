##
# Project Title
#
# @file
# @version 0.1



# end

install-external-secrets-helm:
	helm repo add external-secrets https://charts.external-secrets.io
	helm show values external-secrets/external-secrets > org_helm_values.yaml

generate-from-helm:
	helm template --validate --name-template external-secrets -f values.yaml -n external-secrets external-secrets/external-secrets > k8s_testfiles/external-secrets-rendered.yaml

start-kind:
	kind delete cluster
	kind create cluster --config k8s_testfiles/kind.yaml

deploy-external-secrets-operator:
	kubectl create namespace external-secrets
	kubectl apply -f k8s_testfiles/external-secrets-rendered.yaml
