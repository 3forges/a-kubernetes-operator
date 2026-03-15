# Kube Operator

```bash
chmod +x ./utils/cluster/provision.sh
./utils/cluster/provision.sh

```

```bash
export KND_CLUSTER='kind-koperator-play'

kubectl cluster-info --context ${KND_CLUSTER}

kubectl get all
kubectl get ns
kubectl get nodes

```

```bash

chmod +x ./utils/deploy.sh
./utils/deploy.sh


```

## References

* https://github.com/kubernetes/client-go/tree/master/examples
* Especially:
  * https://github.com/kubernetes/client-go/tree/master/examples/workqueue
  * https://github.com/kubernetes/community/blob/master/contributors/devel/sig-api-machinery/controllers.md

* Creating informers:
  * https://blog.dsb.dev/posts/creating-dynamic-informers/
  * https://www.plural.sh/blog/manage-kubernetes-events-informers/
  * https://labex.io/tutorials/kubernetes-how-to-implement-event-driven-workflows-in-kubernetes-392994
* About Golang:
  * https://go.dev/blog/using-go-modules