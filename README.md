## Purpose

This repo consists of SPIKE code, demonstrating the use of a `TelemetryRecorder`
controller and dynamic Kubernetes
[client](https://github.com/kubernetes/client-go/tree/master/dynamic).  The
controller is built using the KubeBuilder
[framework](https://kubebuilder.io/quick-start.html).

## Instructions to try it out

First, you'll need a Kubernetes cluster to run this controller against.

Next, to install the `TelemetryRecorder` CRD into the cluster, run `make install`.

Then, running `kubectl apply -f config/samples/` will do the following:

  * Deploy an instance of the `TelemetryRecorder` CR for the `LogSink` resources
  * Register the `LogSink` CRD
  * Deploy an instance of the `LogSink` CR

Finally, run `make run`, and you should observe the following instrumented data
in the log:

```
Instrumented values are [
    {
        "Host": "example.com",
        "Port": 8080,
        "enable_tls": null,
        "insecure_skip_verify": null,
        "tls": true,
        "type": null
    },
    {
        "Host": "example.com",
        "Port": 8080,
        "enable_tls": null,
        "insecure_skip_verify": null,
        "tls": true,
        "type": null
    },
    {
        "Host": "example.com",
        "Port": 8888,
        "enable_tls": null,
        "insecure_skip_verify": null,
        "tls": true,
        "type": null
    }
]
```
