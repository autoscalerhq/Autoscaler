
## Install charts
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo add longhorn https://charts.longhorn.io
helm repo add metallb https://metallb.github.io/metallb
helm repo add istio https://istio-release.storage.googleapis.com/charts


helm repo update

## Set up local prometheus; all other logs and metrics should be sent to grafana cloud
## TODO set up where it could be local or cloud based

FILE="./secrets.yaml"

if [ ! -r "$FILE" ]; then
  echo "Unable to find '$FILE'."
fi

kubectl apply -f secrets.yaml

# 0-monitoring
echo "Installing monitoring charts"
kubectl apply -f ./0-monitoring/0-namespace.yaml

# prometheus monitoring server
helm upgrade --install prometheus prometheus-community/prometheus --values ./0-monitoring/1-prometheus-cloud-write.yaml -n monitoring

# speedtest for ingress and egress network monitoring; https://truecharts.org/charts/stable/speedtest-exporter/
helm upgrade --install speedtest oci://tccr.io/truecharts/speedtest-exporter --values ./0-monitoring/2-speedtest.yaml -n monitoring

# disk monitoring export for disk management; https://truecharts.org/charts/stable/smartctl-exporter/
helm upgrade --install smartctl oci://tccr.io/truecharts/smartctl-exporter --values ./0-monitoring/3-smartctl.yaml -n monitoring

# set up scrutiny to watch disk status ; https://truecharts.org/charts/stable/scrutiny/
helm upgrade --install scrutiny oci://tccr.io/truecharts/scrutiny --values ./0-monitoring/4-scrutiny.yaml -n monitoring

# 1-infra
echo "Installing infrastructure charts"
kubectl apply -f ./1-infra/0-namespace.yaml

# Setup metalLB to be able to give external Ips
helm upgrade --install metallb metallb/metallb --values ./1-infra/1-metallb.values.yaml -n metallb-system --wait

while ! kubectl apply -n metallb-system -f ./1-infra/2-metallb.advertisements.yaml
do sleep 5; done

helm upgrade --install metallb metallb/metallb --values ./1-infra/1-metallb.values.yaml -n metallb-system --wait

# Temp disabled due to conflict with talos CNI and istio CNI when tring to run in ambient mode
## Install Istio with ambient mesh; https://istio.io/latest/docs/ambient/install/helm/
#helm upgrade --install istio-base istio/base -n istio-system --set defaultRevision=default --wait
#
## Kubernetes Gateway API CRDs
#kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.1.0/standard-install.yaml;
#
#helm install istiod istio/istiod -n istio-system --set profile=ambient --wait
#
#helm install istio-cni istio/cni -n istio-system --set profile=ambient --wait
#
#helm install ztunnel istio/ztunnel -n istio-system --wait
#
#helm install istio-ingress istio/gateway -n istio-ingress --create-namespace --wait



