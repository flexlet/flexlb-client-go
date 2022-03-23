# FlexLB go client

Flexible load balancer go client to control keepalived and haproxy

## Build

### Generate code

```sh
FLEXLB_API=../flexlb-api
swagger generate client -f ${FLEXLB_API}/swagger/flexlb-api-spec.yaml
```

### Build binary

#### For Linux
```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/flexlb-cli cmd/flexlb-client/main.go
```

## Run

### Generage self-signed certificate

#### Generate CA key and CA certs

See {FLEXLB_API}/README.md

#### Generate server key and certs

See {FLEXLB_API}/README.md

#### Generate client key and certs

```sh
cd /etc/flexlb/certs
DNS_NAME="example.com"
openssl genrsa -out client.key 2048
openssl req -new -out client.csr -key client.key -subj "/CN=${DNS_NAME}"
openssl x509 -req -in client.csr -out client.crt -signkey client.key -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650
```

### Run FlexLB API server

See {FLEXLB_API}/README.md

### Run FlexLB Client

#### Show ready status
```sh
./flexlb-cli -status
```

## Test

#### Create instance
```sh
TEMPLATE="test/instance_template.json"
NAME="inst1"
VIP="192.168.2.1"
sed "s/<NAME>/${NAME}/g; s/<VIP>/${VIP}/g" ${TEMPLATE} > /tmp/inst1.json
./flexlb-cli -create -config /tmp/inst1.json
```

#### List instance
```sh
./flexlb-cli -list
./flexlb-cli -list -name inst1
```

#### Modify instance
```sh
# edit /tmp/inst1.json
./flexlb-cli -modify inst1 -config /tmp/inst1.json
```

#### Get instance
```sh
./flexlb-cli -get inst1
```

#### Stop/Start instance
```sh
./flexlb-cli -stop inst1
./flexlb-cli -start inst1
```
#### Delete instance
```sh
./flexlb-cli -delete inst1
```
