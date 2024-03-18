# PDB

 #k8s #pdb

```sh
$ k apply -f pdb.yaml
$ k get pdb
$ k apply -f deployment.yaml
# observation in terminal 2
$ k get po -w
# observation in terminal 3
$ k get pdb -w
# terminal 1
$ k drain node01 --ignore-daemonsets
node/node01 already cordoned
Warning: ignoring DaemonSet-managed Pods: kube-system/canal-gjxwj, kube-system/kube-proxy-lhxdd
evicting pod kube-system/coredns-86b698fbb6-hqpmj
evicting pod default/nginx-7854ff8877-bwjcw
evicting pod default/nginx-7854ff8877-x995k
evicting pod kube-system/coredns-86b698fbb6-8q542
error when evicting pods/"nginx-7854ff8877-bwjcw" -n "default" (will retry after 5s): Cannot evict pod as it would violate the pod's disruption budget.
pod/nginx-7854ff8877-x995k evicted
evicting pod default/nginx-7854ff8877-bwjcw
pod/nginx-7854ff8877-bwjcw evicted
pod/coredns-86b698fbb6-hqpmj evicted
pod/coredns-86b698fbb6-8q542 evicted
node/node01 drained
```

```sh
$ k apply -f deployment.yaml
# Видим, что все поды на controlplane
$ k get po -w
# Восстановим node01 после drain
$ k uncordon node01
$ k apply -f deployment.yaml
# альтернатива: k rollout restart deploy/nginx
# Видим поды вернулись на node01
$ k get po -w
```

## Links

* [Specifying a Disruption Budget for your Application](https://kubernetes.io/docs/tasks/run-application/configure-pdb/)

---

