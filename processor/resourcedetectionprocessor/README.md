# Resource Detection Processor

Supported pipeline types: metrics, traces, logs

The resource detection processor can be used to detect resource information from the host,
in a format that conforms to the [OpenTelemetry resource semantic conventions](https://github.com/open-telemetry/opentelemetry-specification/tree/main/specification/resource/semantic_conventions/), and append or
override the resource value in telemetry data with this information.

Currently supported detectors include:

* Environment Variable: Reads resource information from the `OTEL_RESOURCE_ATTRIBUTES` environment
variable. This is expected to be in the format `<key1>=<value1>,<key2>=<value2>,...`, the
details of which are currently pending confirmation in the OpenTelemetry specification.

* System metadata: Queries the host machine to retrieve the following resource attributes:

    * host.name
    * os.type

By default `host.name` is being set to FQDN if possible, and a hostname provided by OS used as fallback.
This logic can be changed with `hostname_sources` configuration which is set to `["dns", "os"]` by default.

Use the following config to avoid getting FQDN and apply hostname provided by OS only:

    ```yaml
    detectors: ["system"]
    system:
        hostname_sources: ["os"]
    ```

    * all valid options for hostname_sources:
        * "dns"
        * "os"

Use the Docker detector (see below) if running the Collector as a Docker container.

* Docker metadata: Queries the Docker daemon to retrieve the following resource attributes from the host machine:

    * host.name
    * os.type

You need to mount the Docker socket (`/var/run/docker.sock` on Linux) to contact the Docker daemon.
Docker detection does not work on macOS.

* GCE Metadata: Uses the [Google Cloud Client Libraries for Go](https://github.com/googleapis/google-cloud-go)
to read resource information from the [GCE metadata server](https://cloud.google.com/compute/docs/storing-retrieving-metadata) to retrieve the following resource attributes:

    * cloud.provider ("gcp")
    * cloud.platform ("gcp_compute_engine")
    * cloud.account.id
    * cloud.region
    * cloud.availability_zone
    * host.id
    * host.image.id
    * host.type

* GKE: Google Kubernetes Engine

    * cloud.provider ("gcp")
    * cloud.platform ("gcp_gke")
    * k8s.cluster.name (name of the GKE cluster)

* AWS EC2: Uses [AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/api/aws/ec2metadata/) to read resource information from the [EC2 instance metadata API](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) to retrieve the following resource attributes:

    * cloud.provider ("aws")
    * cloud.platform ("aws_ec2")
    * cloud.account.id
    * cloud.region
    * cloud.availability_zone
    * host.id
    * host.image.id
    * host.name
    * host.type

It also can optionally gather tags for the EC2 instance that the collector is running on. 
Note that in order to fetch EC2 tags, the IAM role assigned to the EC2 instance must have a policy that includes the `ec2:DescribeTags` permission.

EC2 custom configuration example:
```yaml
detectors: ["ec2"]
ec2:
    # A list of regex's to match tag keys to add as resource attributes can be specified
    tags:
        - ^tag1$
        - ^tag2$
        - ^label.*$
```

* Amazon ECS: Queries the [Task Metadata Endpoint](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-metadata-endpoint.html) (TMDE) to record information about the current ECS Task. Only TMDE V4 and V3 are supported.

    * cloud.provider ("aws")
    * cloud.platform ("aws_ecs")
    * cloud.account.id
    * cloud.region
    * cloud.availability_zone
    * aws.ecs.cluster.arn
    * aws.ecs.task.arn
    * aws.ecs.task.family
    * aws.ecs.task.revision
    * aws.ecs.launchtype (V4 only)
    * aws.log.group.names (V4 only)
    * aws.log.group.arns (V4 only)
    * aws.log.stream.names (V4 only)
    * aws.log.stream.arns (V4 only)
    
* Amazon Elastic Beanstalk: Reads the AWS X-Ray configuration file available on all Beanstalk instances with [X-Ray Enabled](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-configuration-debugging.html).

    * cloud.provider ("aws")
    * cloud.platform ("aws_elastic_beanstalk")
    * deployment.environment
    * service.instance.id
    * service.version

* Amazon EKS

    * cloud.provider ("aws")
    * cloud.platform ("aws_eks")
    * k8s.cluster.name (name of the EKS cluster)
    
* Azure: Queries the [Azure Instance Metadata Service](https://aka.ms/azureimds) to retrieve the following resource attributes:

    * cloud.provider ("azure")
    * cloud.platform ("azure_vm")
    * cloud.region
    * cloud.account.id (subscription ID)
    * host.id (virtual machine ID)
    * host.name
    * azure.vm.size (virtual machine size)
    * azure.vm.scaleset.name (name of the scale set if any)
    * azure.resourcegroup.name (resource group name)

* Azure AKS

  * cloud.provider ("azure")
  * cloud.platform ("azure_aks")

## Configuration

```yaml
# a list of resource detectors to run, valid options are: "env", "system", "gce", "gke", "ec2", "ecs", "elastic_beanstalk", "eks", "azure"
detectors: [ <string> ]
# determines if existing resource attributes should be overridden or preserved, defaults to true
override: <bool>
```

## Ordering

Note that if multiple detectors are inserting the same attribute name, the first detector to insert wins. For example if you had `detectors: [eks, ec2]` then `cloud.platform` will be `aws_eks` instead of `ec2`. The below ordering is recommended.

### GCP

* gke
* gce

### AWS

* elastic_beanstalk
* eks
* ecs
* ec2

The full list of settings exposed for this extension are documented [here](./config.go)
with detailed sample configurations [here](./testdata/config.yaml).
