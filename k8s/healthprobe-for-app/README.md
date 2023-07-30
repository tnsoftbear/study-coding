# k8s health check pattern

This application is a simple example of how to implement a health check pattern in a kubernetes cluster.  
It has 3 stages: starting up, preparing for readiness and running until termination. Each stage time is configured by CLI arguments.  

```sh
Flags:
  -s, --startup int32   Startup delay for application
  -r, --ready int32     Ready time for application
  -e, --end int32       End time for application (default -1)
  -p, --path string     Marker file path (default "/tmp")
```

For instance, if you want to run the application for 120 seconds, with 30 seconds of startup and 60 seconds of readiness, you can run the following command:

```sh
/app/happ -s 30 -r 60 -e 120 -p /tmp
```

It will create `/tmp/startup-marker` after 30 seconds, `/tmp/ready-marker` after 60 seconds and will terminate after 120 seconds with deleting of marker files.

If you define health check attributes the next way:

```yaml
        startupProbe:
          exec:
            command: ["stat", "/tmp/startup-marker"]
          initialDelaySeconds: 26
          periodSeconds: 3
          failureThreshold: 3
        readinessProbe:
          exec:
            command:
            - stat
            - /tmp/readiness-marker
          initialDelaySeconds: 56
          periodSeconds: 3
          failureThreshold: 3
```

You will see the next output in the pod description:

```sh
Events:
  Type     Reason     Age                From               Message
  ----     ------     ----               ----               -------
  Normal   Scheduled  98s                default-scheduler  Successfully assigned default/deployment-for-healthcheck-5db7f4cccf-4ztl7 to k3d-health-check-cluster-server-0
  Normal   Pulled     98s                kubelet            Container image "health-app:local" already present on machine
  Normal   Created    98s                kubelet            Created container health-app
  Normal   Started    98s                kubelet            Started container health-app
  Warning  Unhealthy  68s (x2 over 71s)  kubelet            Startup probe failed: stat: can't stat '/tmp/startup-marker': No such file or directory
  Warning  Unhealthy  38s (x2 over 41s)  kubelet            Readiness probe failed: stat: can't stat '/tmp/readiness-marker': No such file or directory
```

It says that the startup probe and readiness probe are failed twice, but the 3rd time they are passed and the application is ready to serve requests.

## Links

* [spf13/cobra](https://github.com/spf13/cobra) ~ CLI framework for Go
* [Configure Liveness, Readiness and Startup Probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)
