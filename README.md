# Cloud Eye Exporter SDKv3

This document describe how to configure [Flexible Engine](https://cloud.orange-business.com/en/offers/infrastructure-iaas/public-cloud/) Cloud Eye Services exporter for Prometheus.

## Supported Services

Cloud Eye exporter for Prometheus support the following services :

- Elastic Cloud Server
- Elastic LoadBalancer
- Elastic IP
- Elastic Volume Service
- NAT Gateway
- Relational Database Service

## Download

```bash
git clone https://github.com/FlexibleEngineCloud/cloudeye-exporter
```

## (Optional ) Building The Exporter in Golang

```bash
wget https://dl.google.com/go/go1.17.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin # You should put in your .profile or .bashrc
go version # to verify it runs and version #

go get github.com/FlexibleEngineCloud/cloudeye-exporter
cd ~/go/src/github.com/FlexibleEngineCloud/cloudeye-exporter
go build
```

## Configure Exporter

First you need to customize the clouds.yml file to add your credentials :

Modify the "auth_url" value follwing your region (IAM API Endpoints list can be found [here](https://docs.prod-cloud-ocb.orange-business.com/en-us/endpoint/index.html)).

```yaml
global:
  prefix: "flexibleengine"
  port: ":8087"

auth:
  auth_url: "https://iam.eu-west-0.prod-cloud-ocb.orange-business.com/v3"
  project_name: "eu-west-0_myproject"
  access_key: "<your-access-key>"
  secret_key: "<your-secret-key>"
  region: "eu-west-0"
```

or

```yaml
auth:
  auth_url: "https://iam.eu-west-0.prod-cloud-ocb.orange-business.com/v3"
  project_name: "eu-west-0"
  user_name: "username"
  password: "password"
  region: "eu-west-0"
  domain_name: "domain_name"
```

Notes :

- `global.retrieve_offset` (default "0") is an offset (example "-5m") to ask Cloudeye for some older metrics.
- `global.cloudeye_timestamp` (default false) allows Cloudeye Exporter to send metrics with their Cloudeye timestamp
- `global.ignore_empty_datapoints` (default false), when set, will ignore empty datapoints (no warnings)

## Running 

```
./cloudeye-exporter -config=clouds.yml
```

By default service is running on port 8087, default config file location is ./clouds.yml.

Visit metrics in http://localhost:8087/metrics?services=SYS.VPC,SYS.ELB

## Help

```
Usage of ./cloudeye-exporter:
  -config string
        Path to the cloud configuration file (default "./clouds.yml")
  -debug
        If debug the code.
 
```

## Prometheus Configuration

You need to edit Prometheus Configuration file 

```bash
vim /etc/prometheus/prometheus.yml
```

Create a new job `FlexibleEngine`

Pass the The Cloud Eye Exporter address as `targets`

Add `services` to monitor (see supported services)

Example config:

```yaml
global:
  scrape_interval: 1m # Set the scrape interval to every 1 minute seconds. Default is every 1 minute.
  scrape_timeout: 1m
scrape_configs:
  - job_name: 'FlexibleEngine'
    static_configs:
    - targets: ['10.0.0.10:8087']
    params:
      services: ['SYS.VPC,SYS.ELB']
```

Finally you need to restart Prometheus service

```bash
sudo systemctl restart prometheus
```

You can verify that FlexibleEngine targets is now available in Prometheus Targets 

![](/images/prometheus.png)
