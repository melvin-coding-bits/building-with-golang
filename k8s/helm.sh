helm install -f k8s/user-service/service-monitor.yaml dezerv-monitoring prometheus-community/kube-prometheus-stack
helm install elasticsearch elastic/elasticsearch
helm install kibana elastic/kibana
helm install dezerv-tracing-op jaegertracing/jaeger-operator
helm install -f jeager.yaml dezerv-tracing jaegertracing/jaeger
