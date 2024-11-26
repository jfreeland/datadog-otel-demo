# datadog otel demo

This uses https://opentelemetry.io/docs/demo/ and datadog-agent `7-ot-beta-rc`
with the built in collector.

```bash
kind create cluster -n dot --kubeconfig ~/.kube/config
helm repo add datadog https://helm.datadoghq.com
helm repo update
kubectl create ns platform
kubectl apply -f rbac.yaml
# update apiKey and appKey in values.yaml
helm upgrade --create-namespace -n platform --install datadog --set-file datadog.otelCollector.config=otel.yaml -f agent-values.yaml datadog/datadog
kubectl apply -f otel-demo.yaml
kind delete cluster -n dot
```

I had to pass `DD_HOSTNAME` to the agent container to make the agent work in
`kind`.
