<div align="center">
  <picture>
    <source media=   srcset="assets/AutoScaler.png">
    <img src=assets/AutoScaler.png width="500" alt="Auto"/>
  </picture>
</div>
<hr/>
<h1 align="center"> 🚀 Open-Source Autoscaling for Every System 🚀</h1>

## Why Autoscaler? ✨
Autoscaler is a powerful open-source autoscaling solution that reliably scales and right-sizes your applications with dynamic scaling capabilities.
With Autoscaler, you can specify as complex scaling events as you need, ensuring your system's performance is always optimized.
Keep your engineers and product teams aligned by providing clear insights into how your application is scaling by enabling informed decision-making and better collaboration.
Whether you're running a small application or a large-scale PLATFORM, Autoscaler has got you covered! 📈

## Features 🎉

## Planned Features 🔧
🔌 **Ingestion Integrations**: Autoscaler seamlessly integrates with various data sources and monitoring systems, making it easy to manage your system's performance.

⚙️ **Configurable Scaling Algorithms**: Tailor the scaling algorithms to suit your specific application's needs and optimize resource utilization.

🔍 **Monitoring**: Monitor your system's metrics in real-time, ensuring you're always aware of its performance.

📥 **Ingestion**: Ingest data from multiple sources to make informed scaling decisions.

🧩 **Modular Scalers**: Choose from a variety of pre-built scalers or create custom ones that fit your unique requirements.

🗃️ **Multiple Database Support**: Autoscaler supports a wide range of databases, ensuring compatibility with your preferred data storage solution.

💾 **Data Management**: Efficiently handle and manage data, avoiding data-related bottlenecks.

🔗 **Cross-System Correlation**: Autoscaler can intelligently correlate data from different systems, providing a comprehensive view of your entire infrastructure.

🔒 **VPC PrivateLinks**: Keep your data and communication secure with VPC PrivateLinks support.

🤖 **AI Autoscaling**: Utilize the power of AI to make automated, data-driven scaling decisions.

📊 **Analytics & BI**: Gain valuable insights into your system's performance with integrated analytics and Business Intelligence tools.

🏃 **Reliability & HA**: Autoscaler ensures high availability and reliability, guaranteeing a smooth scaling experience.

🔑 **Roles & Permissions**: Manage access to Autoscaler with ease through comprehensive roles and permissions.

🔐 **SSO (Single Sign-On)**: Enhance security and user experience by enabling Single Sign-On authentication.

## Getting Started 🏁
To get started with AutoScaler, follow the instructions below:

1. Install protobuf
- mac: ```brew install protobuf```
- windows: https://stackoverflow.com/questions/13616033/install-protocol-buffers-on-windows

2. Clone the repository
```
git clone https://github.com/autoscalerhq/autoscaler.git
```

3. Install the required dependencies
```
pnpm install
```

4. Navigate to docker directory
```
cd docker
```

5. Run the setup script 
```
chmod +x setup.sh  
./setup.sh
```

## Providers 📡
Autoscaler wants to support a wide range of providers over many different categories.
If you want any of the unsupported items or more, put in an issue for the team to review.

Frameworks
- [ ] Nestjs
- [ ] BullMQ
- [ ] Celery
- [ ] Micronaut
- [ ] Quarkus

Push Based Providers:
- [ ] AWS Cloudwatch
- [ ] Prometheus
- [ ] Datadog
- [ ] New Relic
- [ ] AWS CloudWatch
- [ ] Azure Monitor
- [ ] Google Cloud Monitoring
- [ ] InfluxDB
- [ ] Graphite
- [ ] StatsD
- [ ] SignalFx
- [ ] Splunk
- [ ] Sysdig
- [ ] ElasticSearch
- [ ] AppDynamics
- [ ] Dynatrace

Pull Based Providers:
- [ ] InfluxDB
- [ ] Graphite
- [ ] StatsD
- [ ] SignalFx
- [ ] Splunk
- [ ] ElasticSearch
- [ ] AppDynamics
- [ ] Dynatrace
- [ ] Mixpanel
- [ ] New Relic
- [ ] DataDog
- [ ] MariaDB
- [ ] MySQL
- [ ] MongoDB
- [ ] Cassandra
- [ ] Pulsar
- [ ] Kafka
- [ ] RabbitMQ
- [ ] ActiveMQ
- [ ] Redis
- [ ] Memcached
- [ ] Couchbase
- [ ] CouchDB
- [ ] Neo4j
- [ ] OrientDB
- [ ] ArangoDB
- [ ] Aerospike
- [ ] Hazelcast
- [ ] VoltDB

Analytics
- [ ] Clickhouse
- [ ] Segment
- [ ] Snowflake
- [ ] BigQuery
- [ ] Redshift
- [ ] GCP BigQuery
- [ ] AWS S3

Scalers
- [ ] AWS ECS
- [ ] AWS EC2
- [ ] AWS Beanstalk
- [ ] Azure Containers
- [ ] Azure App Service
- [ ] GCP Containers
- [ ] GCP App Engine
- [ ] Digital Ocean
- [ ] Oracle VM
- [ ] Heroku
- [ ] Fly.io
- [ ] render
- [ ] Capcover
- [ ] microtica
- [ ] Docker

Integrations
- [ ] Kubernetes
- [ ] Nomad
- [ ] Openshift
- [ ] VMWare
- [ ] Proxmox

## Contributing 🤝
We welcome contributions from the community! If you want to contribute to AutoScaler, please follow our Contribution Guidelines to get started.

## License 📜
AutoScaler is licensed under the [FSL License](./LICENSE).

Get ready to supercharge your system's performance with Autoscaler! 🚀 Don't hesitate, start scaling smarter today! 😎💪
