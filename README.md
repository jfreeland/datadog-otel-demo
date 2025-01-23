# datadog otel demo

This uses https://opentelemetry.io/docs/demo/ and datadog-agent with a custom
embedded otel-agent from
https://docs.datadoghq.com/opentelemetry/agent/agent_with_custom_components/.

```bash
docker build . -f Dockerfile.agent-otel -t joeyfreeland/datadog-agent:7-otel-test --no-cache
docker push joeyfreeland/datadog-agent:7-otel-test

kind create cluster -n dot --kubeconfig ~/.kube/config
kc create ns platform
kc apply -f otel-demo.yaml

# add opensearch so we can test the otel-agent elastic plugin
helm repo add opensearch https://opensearch-project.github.io/helm-charts/ && helm repo update
helm upgrade --install opensearch opensearch/opensearch --create-namespace -n opensearch --values opensearch.yaml
helm upgrade --install opensearch-dashboards opensearch/opensearch-dashboards --create-namespace -n opensearch --values opensearch-dashboards.yaml

# add datadog-agent
envsubst < secret.yaml | kc apply -f -
kc apply -f rbac.yaml
# NOTE: I moved the otel.yaml config to be inline in agent-values.yaml
#helm upgrade --create-namespace -n platform --install datadog --set-file datadog.otelCollector.config=otel.yaml -f datadog-agent.yaml datadog/datadog
helm repo add datadog https://helm.datadoghq.com && helm repo update
helm upgrade --create-namespace -n platform --install datadog -f datadog-agent.yaml datadog/datadog

kind delete cluster -n dot
```

I had to pass `DD_HOSTNAME` to the agent container to make the agent work in
`kind`.
