# Grafana Monitoring Panel Usage
1. Download Grafana (https://grafana.com/grafana/download)
   ```
   wget https://dl.grafana.com/enterprise/release/grafana-enterprise-9.0.5-1.x86_64.rpm
   sudo yum install grafana-enterprise-9.0.5-1.x86_64.rpm
   service grafana-server start
   ```
2. Grafana access Prometheus data source
   >(1) Log in to Grafana
   >> Visit http://127.0.0.1:3000，and login
   >> ![load.png](pic/login.jpg)
   
   >(2) Configure Prometheus data source
   >> Configuration—》Data source—》Add data source —》Prometheus —》Fill in the Prometheus address —》Save and test
   >> ![config_prometheus.gif](pic/config_prometheus.gif)
3. Configure related cloud service monitoring views

   If you need to use the following template directly, you need to modify the prometheus configuration and add the task of obtaining enterprise project information. The configuration reference is as follows:
   ```
   $ vi prometheus.yml
   global:
     scrape_interval: 1m # Set the scrape interval to every 1 minute seconds. Default is every 1 minute.
     scrape_timeout: 1m
   scrape_configs:
     # If the enterprise project is enabled, configure this task to obtain the enterprise project information for use in the template
     - job_name: 'flexibleengine-eps'
       metrics_path: "/eps-info"
       static_configs:
       - targets: ['192.168.0.xx:8087']
   ```
   ><font size=6>+</font> —》Import —》Input json template file 》load
   >> ![import.png](pic/import.jpg)
   >> ![img.png](pic/load.jpg)
   
   Template file download links: 
   + [Cloud Search Service (CSS)](templates/css(es)_dashboard_template.json)
   + [Direct Connect (DCAAS)](templates/dcaas_dashboard_template.json)
   + [Distributed Cache Service (DCS)](templates/dcs_dashboard_template.json)
   + [Elastic Cloud Server (ECS)](templates/ecs_dashboard_template.json)
   + [Elastic Load Balancer (ELB)](templates/elb_dashboard_template.json)
   + [Relational Database Service (RDS)](templates/rds_dashboard_template.json)
   + [Web Application Firewall (WAF)](templates/waf_dashboard_template.json)
   + [Virtual Private Cloud VPC](templates/vpc_dashboard_template.json)
4. Display of Results：
   >ECS:
   > ![img.png](pic/ecs.jpg)
   
