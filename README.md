# datadog otel demo

This uses https://opentelemetry.io/docs/demo/ and datadog-agent `7-ot-beta-rc`
with the built in collector.

```bash
kind create cluster -n dot --kubeconfig ~/.kube/config
helm repo add datadog https://helm.datadoghq.com
helm repo update
# update apiKey and appKey in values.yaml
helm upgrade --create-namespace -n datadog --install datadog -f values.yaml datadog/datadog
# note: --set-file datadog.otelcollector.config=otel.yaml didn't seem to do anything
kubectl apply -f otel-demo.yaml
kind delete cluster -n dot
```

I need to play with the otel collector config a bit and see if it will respect
what I pass in. By default datadog will populate a functional config that will
send data to Datadog.

I had to pass `DD_HOSTNAME` to the agent container to make the agent work.
