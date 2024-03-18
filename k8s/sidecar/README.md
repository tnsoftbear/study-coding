# Sidecar Containers

If an init container is created with its `restartPolicy` set to `Always`, it will start and remain running during the entire life of the Pod. This can be helpful for running supporting services separated from the main application containers.

* [Sidecar Containers](https://kubernetes.io/docs/concepts/workloads/pods/sidecar-containers/)


`deployment-sidecar-and-init.yaml` - пример со statup-пробой на сайдкар-контейнере.