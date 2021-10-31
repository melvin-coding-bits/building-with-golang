helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install -f k8s/user-service/service-monitor.yaml dezerv-monitoring prometheus-community/kube-prometheus-stack
helm repo add elastic https://helm.elastic.co
kubectl create -f https://download.elastic.co/downloads/eck/1.8.0/crds.yaml
kubectl apply -f https://download.elastic.co/downloads/eck/1.8.0/operator.yaml
kubectl create namespace logging-es
kubectl apply -f k8s/logging/elastic/elastic.yaml -n logging-es
kubectl apply -f k8s/logging/kibana/kibana.yaml -n logging-es
kubectl apply -f k8s/logging/fluentd/config.yaml -n kube-system
kubectl apply -f k8s/logging/fluentd/daemonset.yaml -n kube-system
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm install dezerv-tracing-op jaegertracing/jaeger-operator
helm install -f k8s/jeager/jeager.yaml dezerv-tracing jaegertracing/jaeger