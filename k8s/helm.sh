helm install -f service-monitor.yaml dezerv-monitoring prometheus-community/kube-prometheus-stack
helm install elasticsearch elastic/elasticsearch
helm install kibana elastic/kibana
