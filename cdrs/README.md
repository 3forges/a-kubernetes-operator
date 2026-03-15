# The Pesto CRDs

## Shirts

```bash
kubectl get shirts.stable.example.com --field-selector spec.color=blue
```

```Yaml
---
apiVersion: stable.example.com/v1
kind: Shirt
metadata:
  name: blue-large-shirt
  namespace: default
spec:
  color: blue
  size: large
```

## OpenBAO Vaults

```bash
kubectl get shirts.stable.example.com --field-selector spec.color=blue
```


```Yaml
---
apiVersion: stable.pesto.io/v1
kind: Baovaults
metadata:
  name: supervault
  namespace: marvel
spec:
  sharedkeys_number: 23
  sharedkeys_corum: 7
  telegram_api_key_secret_name: telegram_bot_apikey
  organization: avengers
```

Now deploy one:

```bash


cat <<EOF >./avengers.ns.yaml
---
apiVersion: v1
kind: Namespace
metadata:
  name: marvel-mcu
EOF

cat <<EOF >./avengers.vault.yaml
---
apiVersion: stable.pesto.io/v1
kind: Baovaults
metadata:
  name: supervault
  namespace: marvel-mcu
spec:
  sharedkeys_number: 23
  sharedkeys_corum: 7
  telegram_api_key_secret_name: telegram_bot_apikey
  organization: avengers
EOF


```
