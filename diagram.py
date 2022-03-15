# -* coding:UTF-8 -*
#!/usr/bin/env python

from diagrams import Cluster, Diagram
from diagrams.aws.compute import EC2, ECS
from diagrams.aws.database import RDS
from diagrams.aws.network import Route53, ELB
from diagrams.onprem.client import Client
from diagrams.onprem.inmemory import Redis
from diagrams.oci.devops import APIService

with Diagram("web_service", show=False):
    dns = Route53("dns")
    client = Client("client")
    load_balancer = ELB("Load Balancer")
    covid_api = APIService("Covid API")

    with Cluster("Webserver Cluster"):
        servers = [EC2("Webserver 1"),
                   EC2("Webserver 2"),
                   EC2("Webserver 3")]

    with Cluster("DB Cluster"):
        db_primary = RDS("primary")
        db_primary - [RDS("replica1"),
                      RDS("replica2")]

    with Cluster("Data pulling Workers"):
        workers = [ECS("worker1"),
                   ECS("worker2"),
                   ECS("worker3")]

    client >> dns >> load_balancer >> servers
    servers >> db_primary

    covid_api << workers >> db_primary
