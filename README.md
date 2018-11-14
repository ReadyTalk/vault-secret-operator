# Vault Secret Operator
This project contains a Kubernetes operator that uses a CRDs create Kubernetes secrets based on Vault values.

Example VaultSecret resource
```
apiVersion: crd.readytalk.com/v1alpha1
kind: VaultSecret
metadata:
  name: static-credentials
spec:
  path: secret/sre/demo
  refreshRate: 10 # default 300 seconds?, minimum 60 seconds
  ttl: 101 # Only applies to dynamic secrets (warning if not); Default 5d
  autoRenew: true # For dynamic secrets, autorenew after TTL; Default true
  autoRefresh: true # For dynamic secrets, get new credentials after maxTTL has been reached; Default true
  revokeOnDelete: true # For dynamic secrets, revoke immediately when the resource is deleted; Default true
```

Built with [Operator Framework](https://github.com/operator-framework/operator-sdk).
