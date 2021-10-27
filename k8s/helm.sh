helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install -f k8s/user-service/service-monitor.yaml dezerv-monitoring prometheus-community/kube-prometheus-stack
helm repo add elastic https://helm.elastic.co
helm install elasticsearch elastic/elasticsearch
helm install kibana elastic/kibana
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm install dezerv-tracing-op jaegertracing/jaeger-operator
helm install -f k8s/jeager/jeager.yaml dezerv-tracing jaegertracing/jaeger