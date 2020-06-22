# cloudeye-exporter

Prometheus cloudeye exporter for [Flexible Engine Cloud](https://cloud.orange-business.com/en/offers/infrastructure-iaas/public-cloud/).

## Download
```
$ git clone https://github.com/FlexibleEngineCloud/cloudeye-exporter
```

## (Option) Building The Discovery with Exact steps on clean Ubuntu 16.04 
```
$ wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin # You should put in your .profile or .bashrc
$ go version # to verify it runs and version #

$ go get github.com/FlexibleEngineCloud/cloudeye-exporter
$ cd ~/go/src/github.com/FlexibleEngineCloud/cloudeye-exporter
$ go build
```

## Usage
```
 ./cloudeye-exporter  -config=clouds.yml
```

The default port is 8087, default config file location is ./clouds.yml.

Visit metrics in http://localhost:8087/metrics?services=SYS.VPC,SYS.ELB


## Help
```
Usage of ./cloudeye-exporter:
  -config string
        Path to the cloud configuration file (default "./clouds.yml")
  -debug
        If debug the code.
 
```

## Example of config file(clouds.yml)
The "URL" value can be get from [Identity and Access Management (IAM) endpoint list](https://docs.prod-cloud-ocb.orange-business.com/endpoint/index.html).
```
global:
  prefix: "flexibleengine"
  port: ":8087"
  metric_path: "/metrics"
  retrieve_offset: "0"
  cloudeye_timestamp: false
  ignore_empty_datapoints: false

auth:
  auth_url: "https://iam.eu-west-0.prod-cloud-ocb.orange-business.com/v3"
  project_name: "eu-west-0"
  access_key: "<your-access-key>"
  secret_key: "<your-secret-key>"
  region: "eu-west-0"

```
or

```
auth:
  auth_url: "https://iam.eu-west-0.prod-cloud-ocb.orange-business.com/v3"
  project_name: "eu-west-0"
  user_name: "username"
  password: "password"
  region: "eu-west-0"
  domain_name: "domain_name"

```

Notes :
* `global.retrieve_offset` (default "0") is an offset (example "-5m") to ask Cloudeye for some older metrics.
* `global.cloudeye_timestamp` (default false) allows Cloudeye Exporter to send metrics with their Cloudeye timestamp
* `global.ignore_empty_datapoints` (default false), when set, will ignore empty datapoints (no warnings)

## Prometheus Configuration
The huaweicloud exporter needs to be passed the address as a parameter, this can be done with relabelling.

Example config:

```
global:
  scrape_interval: 1m # Set the scrape interval to every 1 minute seconds. Default is every 1 minute.
  scrape_timeout: 1m
scrape_configs:
  - job_name: 'flexibleengine'
    static_configs:
    - targets: ['10.0.0.10:8087']
    params:
      services: ['SYS.VPC,SYS.ELB']
```
