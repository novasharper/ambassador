from kat.harness import Query, kube_server_version, kube_client_version
from kat.manifests import KNATIVE_SERVING

from packaging import version

from abstract_tests import AmbassadorTest, HTTP, ServiceType

KNATIVE_EXAMPLE = """
---
apiVersion: serving.knative.dev/v1alpha1
kind: Service
metadata:
 name: helloworld-go
 namespace: default
spec:
 template:
   spec:
     containers:
     - image: gcr.io/knative-samples/helloworld-go
       env:
       - name: TARGET
         value: "Ambassador is Awesome"
"""


class KnativeTest(AmbassadorTest):
    target: ServiceType

    def init(self) -> None:
        server_version = kube_server_version()
        client_version = kube_client_version()
        if version.parse(server_version) < version.parse('1.11') or version.parse(client_version) < version.parse('1.10'):
            self.skip_node = True

        self.target = HTTP()

    def manifests(self) -> str:
        self.manifest_envs = """
    - name: AMBASSADOR_KNATIVE_SUPPORT
      value: "true"
"""

        return super().manifests() + KNATIVE_SERVING + KNATIVE_EXAMPLE

    def queries(self):
        yield Query(self.url(""), expected=404)
        yield Query(self.url(""), headers={'Host': 'random.host.whatever'}, expected=404)
        yield Query(self.url(""), headers={'Host': 'helloworld-go.default.example.com'}, expected=200)
