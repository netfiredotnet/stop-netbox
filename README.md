```
go get github.com/netfiredotnet/stop-netbox/netbox
```
To generate: 
```
openapi-generator generate -i ./openapi.json -g go -o ./netbox --additional-properties packageName=netbox,isGoSubmodule=true
```
