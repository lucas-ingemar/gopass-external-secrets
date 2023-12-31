##
# Project Title
#
# @file
# @version 0.1



# end

generate-external-secrets-values:
	helm repo add external-secrets https://charts.external-secrets.io
	helm show values external-secrets/external-secrets > external-secrets-values.yaml

render-external-secrets:
	helm template --name-template external-secrets -f external-secrets-values.yaml -n external-secrets external-secrets/external-secrets > external-secrets-rendered.yaml
