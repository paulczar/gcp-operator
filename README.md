# Google Cloud Platform Operator

This is a demo operator for the [Google Cloud Platform](https://cloud.google.com) which simplifies requesting google cloud resources in the form on Kubernetes Manifests.

## Example:

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

Edit `deploy/cr.yaml` replacing the project ID placeholders with your GCP project.

Once the GCP Operator is deployed you can create a GCP instance:

```bash
kubectl -n gcp-operator apply -f deploy/cr.yaml
```

After a few minutes check to see if the new instance exists:

```bash
gcloud compute instances list
NAME                                     ZONE           MACHINE_TYPE               PREEMPTIBLE  INTERNAL_IP  EXTERNAL_IP     STATUS
test                                     us-central1-a  custom (2 vCPU, 4.00 GiB)               10.128.0.2                   RUNNING
```

Cleanup:

```
kubectl delete -f deploy
```