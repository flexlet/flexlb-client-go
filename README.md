# FlexLB go client

Flexible load balancer go client to control keepalived and haproxy

## Build

### Clone code

```sh
git clone https://github.com/flexlet/flexlb-client-go.git
```

### Build binary

#### For Linux
```sh
cd flexlb-client-go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/flexlb-cli cmd/flexlb-client/main.go
```

## Run

### Generage self-signed certificate

#### Generate CA key and CA certs

See {FLEXLB}/README.md

#### Generate server key and certs

See {FLEXLB}/README.md

#### Generate client key and certs

See {FLEXLB}/README.md

### Run FlexLB API server

See {FLEXLB}/README.md

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
./flexlb-cli -create /tmp/inst1.json
```

#### List instance
```sh
./flexlb-cli -list
./flexlb-cli -list -name inst1
```

#### Modify instance
```sh
# edit /tmp/inst1.json
./flexlb-cli -modify /tmp/inst1.json
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
