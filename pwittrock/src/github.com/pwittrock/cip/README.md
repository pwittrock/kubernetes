# Container Image Pusher

This is not an official Google Project.

*The code in this repository is intended for educational /development
purposes only and has not undergone review or testing.*

**A flood of callbacks will create a lot of requests to the api server.
Do not run this in a cluster serving production traffic until this issue
is resolved.**

This repo provides an example of how to write a controller to update
[Kubernetes](http://github.com/kubernetes/kubernetes)
[Deployments](http://kubernetes.io/docs/user-guide/deployments/)
 by updating the container image in response to
 [Docker Hub](https://hub.docker.com/) webhook callbacks.

## Features

- Listen to [Docker Hub](https://hub.docker.com/) webhook callbacks and
  - Update the Container image on all annotated Deployments to set the image to use the pushed tag
  - Force a rollout on all annotated Deployments without changing a tag to force an image pull
- Require a secret to be provided in the callback to authenticate the origin
- Only update Deployments that are whitelisted with an annotation
- Only perform updates to images matching a specific regex
- Only rollout if the push time is greater than the last push time

## Project Goals

- Demonstrate how to write Kubernetes controller to perform updates in response to external events

## Requirements

- Kubernetes 1.4+

## Example Usage

### 1. Create a Secret to hold your certificate so the requests for Docker Hub are encrypted

```sh

kubectl create secret tls container-tls-certs --cert server.crt --key server.key

```

### 2. Create a Deployment with to run the controller

```sh

echo 'apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  annotations:
    push-image: "true"
  name: container-image-pusher
spec:
  template:
    metadata:
      labels:
        app: container-image-pusher
    spec:
      containers:
      - image: pwittrock/container-image-pusher:v1
        name: container-image-pusher
        command: [
          "./main",
          "--port", "443",
          "--web-hook-secret", "changeme",
          "--tag-whitelist", "v[0-9\\.]*",
          "--cert-path", "/tls/tls.crt",
          "--key-path", "/tls/tls.key"]
        volumeMounts:
          - name: container-tls-certs
            mountPath: /tls/
            readOnly: true
      volumes:
        - name: container-tls-certs
          secret:
            secretName: container-tls-certs
' | kubectl create -f -

```

This will create a new Deployment running container-image-pusher.
 
- Listen for webhook callbacks for image pushes
- Upon getting a callback verify the secret value
- Find deployments with containers using the same repo
- Update the Container in the PodTemplate with the new container image label
- Update the PodTemplate labels with the pushtime of the container (as recorded in the callback json payload)
 
Additionally it:
 
- Listens for secure connections on port 443
  - If your webhook callbacks are already secure - such as through a proxy - then you can bind to an insecure connection by specifying `""` for your tls.
- Uses the certificate files mounted under /tls/
- Expects a query parameter called `secret` to contain the value `changeme`.  If not provided it log a message and ignore the hook.
- Processes webhooks only for image tags matching `v[0-9\.]*`
  - It will only push updates to containers whose images also contain this whitelist
- Updates **only** deployments with the annotation `replicatingperfection.net/push-image` (changed using `--controller-annotation`)
- At update time, creates or updates the label `auto-pushed-image-<repo/name>` with the push time of the image
  - If another update is received with an earlier push time than one recorded, it will be ignore.

### 3. Add the whitelisted annotation to a Deployment you want to update so the controller doesn't ignore it.

```sh

kubectl annotate deployment your-deployment-name replicatingperfection.net/push-image=true

```

### 4. Push a new image to Docker Hub with a whitelisted tag

```sh

docker push your/repo:v1

```

### 5. Watch for the new Pod creation times.


```sh

kubectl get pods

```

## Requeired callback fields

```json
{
  "repository": {
    "repo_name": "your/reponame"
  },
  "push_data": {
    "pushed_at": 77777,
    "tag": "v9"
  }
}

```

## Finer grain image push controll

If you want to push containers with images containing a tag
`fooX` to `fooY` and containers with images contain `barX` to `barY` you
can do this by running multiple instances with separate tag whitelists and
optionally separate annotations.

## Trouble Shooting

First manually send a request to the container-image-pusher to try to update an image:

```sh

curl -i -X POST -H "Content-Type: application/json" -d '{"repository": {"repo_name": "your/reponame"}, "push_data": {"pushed_at": 77777, "tag": "v1"}}'  https://yourdomain.com?secret=changeme

```

Then view the logs to see if the request was received and parsed:

```sh

kubectl logs <podname>

```

or

```sh

kubectl logs $(kubectl get pods --show-labels -l app=container-image-pusher -o name |sed -e s:pod/:: ) 

```

The logs should record if the request was received, the body that was
parsed, and deployments it looked at.

```sh
$ kubectl logs $(kubectl get pods --show-labels -l app=container-image-push -o name |sed -e s:pod/:: ) 
Listening on :443.
Processing request.  Parsed Body: &{CallbackURL: PushData:{Images:[] PushedAt:77777 Pusher: Tag:v9} Repository:{DateCreated:0 Description: IsTrusted:false FullDescription: Name: Namespace: Owner: RepoName:pwittrock/cli-docs RepoURL: Status:}} Values map[secret:[examplesecret]] URI: /?secret=changeme.
Deployment docs found annotation replicatingperfection.net/push-image.
Skipping image: [pwittrock/api-docs] on deployment: [docs] because image does not match pushed repo: [pwittrock/cli-docs].
Deployment kubectldocs found annotation replicatingperfection.net/push-image.
Image pwittrock/cli-docs matches Deployment kubectldocs image pwittrock/cli-docs.
```


## Known issues

- There is no ratelimiting and a `list deployments` request is made to the api server for each callback.
- The response will always be 200 even if the request failed.
