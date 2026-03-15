# Go do that

To deploy a Kubernetes Operator whose container image is pesto-operator:latest, you typically need at least:

A Namespace (optional but recommended)

A ServiceAccount

RBAC (Role/ClusterRole + Binding)

A Deployment running the operator pod

```bash
kubectl apply -f ./kubernetes/
```



