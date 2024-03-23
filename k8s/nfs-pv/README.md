

```sh
# Install helm
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
# Install nfs common libs
sudo apt-get update
sudo apt-get install nfs-common -y

k create ns nfs
helm repo add nfs-ganesha-server-and-external-provisioner https://kubernetes-sigs.github.io/nfs-ganesha-server-and-external-provisioner/
helm repo update
# helm install my-release nfs-ganesha-server-and-external-provisioner/nfs-server-provisioner
# helm delete my-release
helm pull nfs-ganesha-server-and-external-provisioner/nfs-server-provisioner --untar
cp values.yaml ./nfs-server-provisioner/
# Посмотреть доступное место на диске
df -hT 
```

Изменить в `values.yaml`

```yaml
persistence:
  enabled: true
  storageClass: "-"
  size: 100Mi

storageClass:
  defaultClass: true

nodeSelector:
  kubernetes.io/hostname: node01
```

```sh
k -n nfs apply -f pv.yaml
k -n nfs get pv
helm -n nfs upgrade --install nfs-server-provisioner nfs-ganesha-server-and-external-provisioner/nfs-server-provisioner
k -n nfs get po
k -n nfs get pv
k -n nfs apply -f pvc.yaml
k -n nfs get pvc
# PV создана
k -n nfs get pv
```

```yaml - pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
    - name: nginx
      image: nginx
      volumeMounts:
        - name: test
          mountPath: /test
  volumes:
    - name: test
      persistentVolumeClaim:
        claimName: test-dynamic-volume-claim
```

```sh
k -n nfs apply -f pod.yaml
k -n nfs describe po nginx
# Ошибка на killercoda.com

  Warning  FailedMount  2s (x3 over 3s)  kubelet            MountVolume.SetUp failed for volume "pvc-20e868e3-0dc2-4343-8289-fb21097ef76b" : mount failed: exit status 32
Mounting command: mount
Mounting arguments: -t nfs -o retrans=2,timeo=30,vers=3 10.99.95.198:/export/pvc-20e868e3-0dc2-4343-8289-fb21097ef76b /var/lib/kubelet/pods/a839bc4f-695d-458e-aee0-c3642658eab9/volumes/kubernetes.io~nfs/pvc-20e868e3-0dc2-4343-8289-fb21097ef76b
Output: mount: /var/lib/kubelet/pods/a839bc4f-695d-458e-aee0-c3642658eab9/volumes/kubernetes.io~nfs/pvc-20e868e3-0dc2-4343-8289-fb21097ef76b: bad option; for several filesystems (e.g. nfs, cifs) you might need a /sbin/mount.<type> helper program.
```

