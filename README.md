# Google Cloud Platform Operator

This is a demo operator for the [Google Cloud Platform](https://cloud.google.com) which simplifies requesting google cloud resources in the form on Kubernetes Manifests.

## Annotations

You can set the following Annotations:

| Annotation | Description |
| ---------- | ----------- |
| `cloud.google.com/project-id` | Sets the GCP Project ID if different to that used in operator service account |


## Example Usage

Create a namespace to run the operator in:

```bash
kubectl create namespace gcp-operator
```

Create a secret containing your GCP account credentials:

```bash
kubectl -n gcp-operator create secret \
    generic gcp-operator \
  --from-file=google.json=/path/to/credentials.json
```

Deploy the GCP Operator:

```bash
kubectl -n gcp-operator apply -f deploy/rbac.yaml
kubectl -n gcp-operator apply -f deploy/sa.yaml
kubectl -n gcp-operator apply -f deploy/crd.yaml
kubectl -n gcp-operator apply -f deploy/operator.yaml
```

### Create GCP Instance

Edit `deploy/cr-instance.yaml` replacing the project ID placeholders with your GCP project.

Once the GCP Operator is deployed you can create a GCP instance:

```bash
kubectl -n gcp-operator apply -f deploy/cr-instance.yaml
```

After a few minutes check to see if the new instance exists:

```bash
gcloud compute instances list
NAME                                     ZONE           MACHINE_TYPE               PREEMPTIBLE  INTERNAL_IP  EXTERNAL_IP     STATUS
test                                     us-central1-a  custom (2 vCPU, 4.00 GiB)               10.128.0.2                   RUNNING
```

### Create GCP Address

Edit `deploy/cr-address.yaml` replacing the project ID placeholders with your GCP project.

Once the GCP Operator is deployed you can create a GCP instance:

```bash
kubectl -n gcp-operator apply -f deploy/cr-address.yaml
```

After a few minutes check to see if the new instance exists:

```bash
gcloud compute addresses list
NAME                   REGION       ADDRESS          STATUS
example                us-central1  35.226.61.203    RESERVED

```

Cleanup:

```
kubectl delete -f deploy
```
