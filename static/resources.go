// Code generated by compile_resources. DO NOT EDIT.
//go:generate go run ../pkg/compile_resources

package static

import (
	"github.com/cybozu-go/cke"
)

// Resources is the Kubernetes resource definitions embedded in CKE.
var Resources = []cke.ResourceDefinition{
	{
		Key:        "ServiceAccount/kube-system/cke-cluster-dns",
		Kind:       "ServiceAccount",
		Namespace:  "kube-system",
		Name:       "cke-cluster-dns",
		Revision:   1,
		Image:      "",
		Definition: []byte("apiVersion: v1\nkind: ServiceAccount\nmetadata:\n  name: cke-cluster-dns\n  namespace: kube-system\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n"),
	},
	{
		Key:        "ServiceAccount/kube-system/cke-etcdbackup",
		Kind:       "ServiceAccount",
		Namespace:  "kube-system",
		Name:       "cke-etcdbackup",
		Revision:   1,
		Image:      "",
		Definition: []byte("apiVersion: v1\nkind: ServiceAccount\nmetadata:\n  name: cke-etcdbackup\n  namespace: kube-system\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n"),
	},
	{
		Key:        "ServiceAccount/kube-system/cke-node-dns",
		Kind:       "ServiceAccount",
		Namespace:  "kube-system",
		Name:       "cke-node-dns",
		Revision:   1,
		Image:      "",
		Definition: []byte("apiVersion: v1\nkind: ServiceAccount\nmetadata:\n  name: cke-node-dns\n  namespace: kube-system\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n"),
	},
	{
		Key:        "ClusterRole/system:cluster-dns",
		Kind:       "ClusterRole",
		Namespace:  "",
		Name:       "system:cluster-dns",
		Revision:   1,
		Image:      "",
		Definition: []byte("\nkind: ClusterRole\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: system:cluster-dns\n  labels:\n    kubernetes.io/bootstrapping: rbac-defaults\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n    # turn on auto-reconciliation\n    # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#auto-reconciliation\n    rbac.authorization.kubernetes.io/autoupdate: \"true\"\nrules:\n  - apiGroups: [\"\"]\n    resources:\n      - endpoints\n      - services\n      - pods\n      - namespaces\n    verbs: [\"list\", \"watch\"]\n  - apiGroups: [\"policy\"]\n    resources: [\"podsecuritypolicies\"]\n    verbs: [\"use\"]\n    resourceNames: [\"cke-restricted\"]\n"),
	},
	{
		Key:        "ClusterRole/system:kube-apiserver-to-kubelet",
		Kind:       "ClusterRole",
		Namespace:  "",
		Name:       "system:kube-apiserver-to-kubelet",
		Revision:   1,
		Image:      "",
		Definition: []byte("kind: ClusterRole\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: system:kube-apiserver-to-kubelet\n  labels:\n    kubernetes.io/bootstrapping: rbac-defaults\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n    # turn on auto-reconciliation\n    # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#auto-reconciliation\n    rbac.authorization.kubernetes.io/autoupdate: \"true\"\nrules:\n  - apiGroups: [\"\"]\n    resources:\n      - nodes/proxy\n      - nodes/stats\n      - nodes/log\n      - nodes/spec\n      - nodes/metrics\n    verbs: [\"*\"]\n"),
	},
	{
		Key:        "ClusterRoleBinding/system:cluster-dns",
		Kind:       "ClusterRoleBinding",
		Namespace:  "",
		Name:       "system:cluster-dns",
		Revision:   1,
		Image:      "",
		Definition: []byte("\nkind: ClusterRoleBinding\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: system:cluster-dns\n  labels:\n    kubernetes.io/bootstrapping: rbac-defaults\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n    rbac.authorization.kubernetes.io/autoupdate: \"true\"\nroleRef:\n  apiGroup: rbac.authorization.k8s.io\n  kind: ClusterRole\n  name: system:cluster-dns\nsubjects:\n- kind: ServiceAccount\n  name: cke-cluster-dns\n  namespace: kube-system\n"),
	},
	{
		Key:        "ClusterRoleBinding/system:kube-apiserver",
		Kind:       "ClusterRoleBinding",
		Namespace:  "",
		Name:       "system:kube-apiserver",
		Revision:   1,
		Image:      "",
		Definition: []byte("kind: ClusterRoleBinding\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: system:kube-apiserver\n  labels:\n    kubernetes.io/bootstrapping: rbac-defaults\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n    rbac.authorization.kubernetes.io/autoupdate: \"true\"\nroleRef:\n  apiGroup: rbac.authorization.k8s.io\n  kind: ClusterRole\n  name: system:kube-apiserver-to-kubelet\nsubjects:\n- kind: User\n  name: kubernetes\n"),
	},
	{
		Key:        "PodSecurityPolicy/cke-node-dns",
		Kind:       "PodSecurityPolicy",
		Namespace:  "",
		Name:       "cke-node-dns",
		Revision:   1,
		Image:      "",
		Definition: []byte("apiVersion: policy/v1beta1\nkind: PodSecurityPolicy\nmetadata:\n  name: cke-node-dns\n  annotations:\n    seccomp.security.alpha.kubernetes.io/allowedProfileNames: 'docker/default'\n    seccomp.security.alpha.kubernetes.io/defaultProfileName:  'docker/default'\n    cke.cybozu.com/revision: \"1\"\nspec:\n  privileged: false\n  # Required to prevent escalations to root.\n  allowPrivilegeEscalation: false\n  allowedCapabilities:\n    - NET_BIND_SERVICE\n  # Allow core volume types.\n  volumes:\n    - 'configMap'\n    - 'emptyDir'\n    - 'projected'\n    - 'secret'\n    - 'downwardAPI'\n    # Assume that persistentVolumes set up by the cluster admin are safe to use.\n    - 'persistentVolumeClaim'\n  hostNetwork: true\n  hostIPC: false\n  hostPID: false\n  runAsUser:\n    rule: 'RunAsAny'\n  seLinux:\n    rule: 'RunAsAny'\n  supplementalGroups:\n    rule: 'RunAsAny'\n  fsGroup:\n    rule: 'RunAsAny'\n  readOnlyRootFilesystem: true\n"),
	},
	{
		Key:        "PodSecurityPolicy/cke-restricted",
		Kind:       "PodSecurityPolicy",
		Namespace:  "",
		Name:       "cke-restricted",
		Revision:   1,
		Image:      "",
		Definition: []byte("apiVersion: policy/v1beta1\nkind: PodSecurityPolicy\nmetadata:\n  name: cke-restricted\n  annotations:\n    seccomp.security.alpha.kubernetes.io/allowedProfileNames: 'docker/default'\n    seccomp.security.alpha.kubernetes.io/defaultProfileName:  'docker/default'\n    cke.cybozu.com/revision: \"1\"\nspec:\n  privileged: false\n  # Required to prevent escalations to root.\n  allowPrivilegeEscalation: false\n  # This is redundant with non-root + disallow privilege escalation,\n  # but we can provide it for defense in depth.\n  requiredDropCapabilities:\n    - ALL\n  # Allow core volume types.\n  volumes:\n    - 'configMap'\n    - 'emptyDir'\n    - 'projected'\n    - 'secret'\n    - 'downwardAPI'\n    # Assume that persistentVolumes set up by the cluster admin are safe to use.\n    - 'persistentVolumeClaim'\n  hostNetwork: false\n  hostIPC: false\n  hostPID: false\n  runAsUser:\n    # Require the container to run without root privileges.\n    rule: 'MustRunAsNonRoot'\n  seLinux:\n    # This policy assumes the nodes are using AppArmor rather than SELinux.\n    rule: 'RunAsAny'\n  supplementalGroups:\n    rule: 'MustRunAs'\n    ranges:\n      # Forbid adding the root group.\n      - min: 1\n        max: 65535\n  fsGroup:\n    rule: 'MustRunAs'\n    ranges:\n      # Forbid adding the root group.\n      - min: 1\n        max: 65535\n  readOnlyRootFilesystem: true\n"),
	},
	{
		Key:        "Role/kube-system/system:etcdbackup",
		Kind:       "Role",
		Namespace:  "kube-system",
		Name:       "system:etcdbackup",
		Revision:   1,
		Image:      "",
		Definition: []byte("\nkind: Role\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: system:etcdbackup\n  namespace: kube-system\n  labels:\n    kubernetes.io/bootstrapping: rbac-defaults\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n    # turn on auto-reconciliation\n    # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#auto-reconciliation\n    rbac.authorization.kubernetes.io/autoupdate: \"true\"\nrules:\n  - apiGroups: [\"policy\"]\n    resources: [\"podsecuritypolicies\"]\n    verbs: [\"use\"]\n    resourceNames: [\"cke-restricted\"]\n"),
	},
	{
		Key:        "Role/kube-system/system:node-dns",
		Kind:       "Role",
		Namespace:  "kube-system",
		Name:       "system:node-dns",
		Revision:   1,
		Image:      "",
		Definition: []byte("\nkind: Role\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: system:node-dns\n  namespace: kube-system\n  labels:\n    kubernetes.io/bootstrapping: rbac-defaults\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n    # turn on auto-reconciliation\n    # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#auto-reconciliation\n    rbac.authorization.kubernetes.io/autoupdate: \"true\"\nrules:\n  - apiGroups: [\"policy\"]\n    resources: [\"podsecuritypolicies\"]\n    verbs: [\"use\"]\n    resourceNames: [\"cke-node-dns\"]\n"),
	},
	{
		Key:        "RoleBinding/kube-system/system:etcdbackup",
		Kind:       "RoleBinding",
		Namespace:  "kube-system",
		Name:       "system:etcdbackup",
		Revision:   1,
		Image:      "",
		Definition: []byte("\nkind: RoleBinding\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: system:etcdbackup\n  namespace: kube-system\n  labels:\n    kubernetes.io/bootstrapping: rbac-defaults\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n    rbac.authorization.kubernetes.io/autoupdate: \"true\"\nroleRef:\n  apiGroup: rbac.authorization.k8s.io\n  kind: Role\n  name: system:etcdbackup\nsubjects:\n- kind: ServiceAccount\n  name: cke-etcdbackup\n  namespace: kube-system\n"),
	},
	{
		Key:        "RoleBinding/kube-system/system:node-dns",
		Kind:       "RoleBinding",
		Namespace:  "kube-system",
		Name:       "system:node-dns",
		Revision:   1,
		Image:      "",
		Definition: []byte("\nkind: RoleBinding\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: system:node-dns\n  namespace: kube-system\n  labels:\n    kubernetes.io/bootstrapping: rbac-defaults\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n    rbac.authorization.kubernetes.io/autoupdate: \"true\"\nroleRef:\n  apiGroup: rbac.authorization.k8s.io\n  kind: Role\n  name: system:node-dns\nsubjects:\n- kind: ServiceAccount\n  name: cke-node-dns\n  namespace: kube-system\n"),
	},
	{
		Key:        "DaemonSet/kube-system/node-dns",
		Kind:       "DaemonSet",
		Namespace:  "kube-system",
		Name:       "node-dns",
		Revision:   1,
		Image:      "quay.io/cybozu/unbound:1.11.0.1",
		Definition: []byte("\nkind: DaemonSet\napiVersion: apps/v1\nmetadata:\n  name: node-dns\n  namespace: kube-system\n  annotations:\n    cke.cybozu.com/image: \"quay.io/cybozu/unbound:1.11.0.1\"\n    cke.cybozu.com/revision: \"1\"\nspec:\n  selector:\n    matchLabels:\n      cke.cybozu.com/appname: node-dns\n  updateStrategy:\n    type: RollingUpdate\n    rollingUpdate:\n      maxUnavailable: 1\n  template:\n    metadata:\n      labels:\n        cke.cybozu.com/appname: node-dns\n    spec:\n      priorityClassName: system-node-critical\n      nodeSelector:\n        kubernetes.io/os: linux\n      hostNetwork: true\n      tolerations:\n        - operator: Exists\n      terminationGracePeriodSeconds: 1\n      serviceAccountName: cke-node-dns\n      containers:\n        - name: unbound\n          image: quay.io/cybozu/unbound:1.11.0.1\n          args:\n            - -c\n            - /etc/unbound/unbound.conf\n          securityContext:\n            allowPrivilegeEscalation: false\n            capabilities:\n              add:\n              - NET_BIND_SERVICE\n              drop:\n              - all\n            readOnlyRootFilesystem: true\n          readinessProbe:\n            tcpSocket:\n              port: 53\n              host: localhost\n            periodSeconds: 1\n          livenessProbe:\n            tcpSocket:\n              port: 53\n              host: localhost\n            periodSeconds: 1\n            initialDelaySeconds: 1\n            failureThreshold: 6\n          volumeMounts:\n            - name: config-volume\n              mountPath: /etc/unbound\n        - name: reload\n          image: quay.io/cybozu/unbound:1.11.0.1\n          command:\n          - /usr/local/bin/reload-unbound\n          securityContext:\n            allowPrivilegeEscalation: false\n            capabilities:\n              drop:\n              - all\n            readOnlyRootFilesystem: true\n          volumeMounts:\n            - name: config-volume\n              mountPath: /etc/unbound\n      volumes:\n        - name: config-volume\n          configMap:\n            name: node-dns\n            items:\n            - key: unbound.conf\n              path: unbound.conf\n"),
	},
	{
		Key:        "Deployment/kube-system/cluster-dns",
		Kind:       "Deployment",
		Namespace:  "kube-system",
		Name:       "cluster-dns",
		Revision:   1,
		Image:      "quay.io/cybozu/coredns:1.7.0.1",
		Definition: []byte("\nkind: Deployment\napiVersion: apps/v1\nmetadata:\n  name: cluster-dns\n  namespace: kube-system\n  annotations:\n    cke.cybozu.com/image: \"quay.io/cybozu/coredns:1.7.0.1\"\n    cke.cybozu.com/revision: \"1\"\nspec:\n  replicas: 2\n  strategy:\n    type: RollingUpdate\n    rollingUpdate:\n      maxUnavailable: 1\n  selector:\n    matchLabels:\n      cke.cybozu.com/appname: cluster-dns\n  template:\n    metadata:\n      labels:\n        cke.cybozu.com/appname: cluster-dns\n        k8s-app: coredns # sonobuoy requires\n      annotations:\n        prometheus.io/port: \"9153\"\n    spec:\n      priorityClassName: system-cluster-critical\n      serviceAccountName: cke-cluster-dns\n      tolerations:\n        - key: node-role.kubernetes.io/master\n          effect: NoSchedule\n        - key: \"CriticalAddonsOnly\"\n          operator: \"Exists\"\n        - key: kubernetes.io/e2e-evict-taint-key\n          operator: Exists\n          # for sonobuoy https://github.com/vmware-tanzu/sonobuoy/pull/878\n      containers:\n      - name: coredns\n        image: quay.io/cybozu/coredns:1.7.0.1\n        imagePullPolicy: IfNotPresent\n        resources:\n          limits:\n            memory: 170Mi\n          requests:\n            cpu: 100m\n            memory: 70Mi\n        args: [ \"-conf\", \"/etc/coredns/Corefile\" ]\n        lifecycle:\n          preStop:\n            exec:\n              command: [\"sh\", \"-c\", \"sleep\", \"5\"]\n        volumeMounts:\n        - name: config-volume\n          mountPath: /etc/coredns\n          readOnly: true\n        ports:\n        - containerPort: 1053\n          name: dns\n          protocol: UDP\n        - containerPort: 1053\n          name: dns-tcp\n          protocol: TCP\n        - containerPort: 9153\n          name: metrics\n          protocol: TCP\n        securityContext:\n          allowPrivilegeEscalation: false\n          capabilities:\n            drop:\n            - all\n          readOnlyRootFilesystem: true\n        readinessProbe:\n          httpGet:\n            path: /health\n            port: 8080\n            scheme: HTTP\n          timeoutSeconds: 5\n        livenessProbe:\n          httpGet:\n            path: /health\n            port: 8080\n            scheme: HTTP\n          initialDelaySeconds: 60\n          timeoutSeconds: 5\n          successThreshold: 1\n          failureThreshold: 5\n      dnsPolicy: Default\n      volumes:\n        - name: config-volume\n          configMap:\n            name: cluster-dns\n            items:\n            - key: Corefile\n              path: Corefile\n"),
	},
	{
		Key:        "PodDisruptionBudget/kube-system/cluster-dns-pdb",
		Kind:       "PodDisruptionBudget",
		Namespace:  "kube-system",
		Name:       "cluster-dns-pdb",
		Revision:   1,
		Image:      "",
		Definition: []byte("\napiVersion: policy/v1beta1\nkind: PodDisruptionBudget\nmetadata:\n  name: cluster-dns-pdb\n  namespace: kube-system\n  annotations:\n    cke.cybozu.com/revision: \"1\"\nspec:\n  maxUnavailable: 1\n  selector:\n    matchLabels:\n      cke.cybozu.com/appname: cluster-dns\n"),
	},
	{
		Key:        "Service/kube-system/cluster-dns",
		Kind:       "Service",
		Namespace:  "kube-system",
		Name:       "cluster-dns",
		Revision:   1,
		Image:      "",
		Definition: []byte("\nkind: Service\napiVersion: v1\nmetadata:\n  name: cluster-dns\n  namespace: kube-system\n  annotations:\n    cke.cybozu.com/revision: \"1\"\n  labels:\n    cke.cybozu.com/appname: cluster-dns\nspec:\n  selector:\n    cke.cybozu.com/appname: cluster-dns\n  ports:\n    - name: dns\n      port: 53\n      targetPort: 1053\n      protocol: UDP\n    - name: dns-tcp\n      port: 53\n      targetPort: 1053\n      protocol: TCP\n"),
	},
}
