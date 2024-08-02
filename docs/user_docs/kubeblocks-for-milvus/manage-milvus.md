---
title: Manage Milvus with KubeBlocks
description: How to manage Milvus on KubeBlocks
keywords: [milvus, vector database, control plane]
sidebar_position: 1
sidebar_label: Manage Milvus with KubeBlocks
---

# Manage Milvus with KubeBlocks

The popularity of generative AI (Generative AI) has aroused widespread attention and completely ignited the vector database (Vector Database) market.

Milvus is a highly flexible, reliable, and blazing-fast cloud-native, open-source vector database. It powers embedding similarity search and AI applications and strives to make vector databases accessible to every organization. Milvus can store, index, and manage a billion+ embedding vectors generated by deep neural networks and other machine learning (ML) models.

KubeBlocks supports the management of Milvus.

## Before you start

- [Install kbcli](./../installation/install-with-kbcli/install-kbcli.md).
- [Install KubeBlocks](./../installation/install-with-kbcli/install-kubeblocks-with-kbcli.md).
- [Install and enable the milvus addon](./../overview/supported-addons.md#use-addons).

## Create a cluster

***Steps***

1. Execute the following command to create a Milvus cluster. You can change the `cluster-definition` value as any other database supported.

   ```bash
   kbcli cluster create milvus --cluster-definition=milvus-2.3.2
   ```

:::note

View more flags for creating a MySQL cluster to create a cluster with customized specifications.
  
```bash
kbcli cluster create --help
```

:::

2. Check whether the cluster is created.

   ```bash
   kbcli cluster list
   >
   NAME     NAMESPACE   CLUSTER-DEFINITION        VERSION               TERMINATION-POLICY   STATUS           CREATED-TIME
   milvus   default     milvus-2.3.2              milvus-2.3.2          Delete               Running          Jul 05,2024 17:35 UTC+0800   
   ```

3. Check the cluster information.

   ```bash
   kbcli cluster describe milvus
   >
   Name: milvus	 Created Time: Jul 05,2024 17:35 UTC+0800
   NAMESPACE   CLUSTER-DEFINITION   VERSION   STATUS    TERMINATION-POLICY   
   default     milvus-2.3.2                   Running   Delete               

   Endpoints:
   COMPONENT   MODE        INTERNAL                                        EXTERNAL   
   milvus      ReadWrite   milvus-milvus.default.svc.cluster.local:19530   <none>     
   minio       ReadWrite   milvus-minio.default.svc.cluster.local:9000     <none>     
   proxy       ReadWrite   milvus-proxy.default.svc.cluster.local:19530    <none>     
                           milvus-proxy.default.svc.cluster.local:9091                

   Topology:
   COMPONENT   INSTANCE             ROLE     STATUS    AZ       NODE     CREATED-TIME                 
   etcd        milvus-etcd-0        <none>   Running   <none>   <none>   Jul 05,2024 17:35 UTC+0800   
   minio       milvus-minio-0       <none>   Running   <none>   <none>   Jul 05,2024 17:35 UTC+0800   
   milvus      milvus-milvus-0      <none>   Running   <none>   <none>   Jul 05,2024 17:35 UTC+0800   
   indexnode   milvus-indexnode-0   <none>   Running   <none>   <none>   Jul 05,2024 17:35 UTC+0800   
   mixcoord    milvus-mixcoord-0    <none>   Running   <none>   <none>   Jul 05,2024 17:35 UTC+0800   
   querynode   milvus-querynode-0   <none>   Running   <none>   <none>   Jul 05,2024 17:35 UTC+0800   
   datanode    milvus-datanode-0    <none>   Running   <none>   <none>   Jul 05,2024 17:35 UTC+0800   
   proxy       milvus-proxy-0       <none>   Running   <none>   <none>   Jul 05,2024 17:35 UTC+0800   

   Resources Allocation:
   COMPONENT   DEDICATED   CPU(REQUEST/LIMIT)   MEMORY(REQUEST/LIMIT)   STORAGE-SIZE   STORAGE-CLASS     
   milvus      false       1 / 1                1Gi / 1Gi               data:20Gi      csi-hostpath-sc   
   etcd        false       1 / 1                1Gi / 1Gi               data:20Gi      csi-hostpath-sc   
   minio       false       1 / 1                1Gi / 1Gi               data:20Gi      csi-hostpath-sc   
   proxy       false       1 / 1                1Gi / 1Gi               data:20Gi      csi-hostpath-sc   
   mixcoord    false       1 / 1                1Gi / 1Gi               data:20Gi      csi-hostpath-sc   
   datanode    false       1 / 1                1Gi / 1Gi               data:20Gi      csi-hostpath-sc   
   indexnode   false       1 / 1                1Gi / 1Gi               data:20Gi      csi-hostpath-sc   
   querynode   false       1 / 1                1Gi / 1Gi               data:20Gi      csi-hostpath-sc   

   Images:
   COMPONENT   TYPE        IMAGE                                                
   milvus      milvus      milvusdb/milvus:v2.3.2                               
   etcd        etcd        docker.io/milvusdb/etcd:3.5.5-r2                     
   minio       minio       docker.io/minio/minio:RELEASE.2022-03-17T06-34-49Z   
   proxy       proxy       milvusdb/milvus:v2.3.2                               
   mixcoord    mixcoord    milvusdb/milvus:v2.3.2                               
   datanode    datanode    milvusdb/milvus:v2.3.2                               
   indexnode   indexnode   milvusdb/milvus:v2.3.2                               
   querynode   querynode   milvusdb/milvus:v2.3.2                               

   Show cluster events: kbcli cluster list-events -n default milvus
   ```

## Scale

Currently, KubeBlocks supports vertically scaling a Milvus cluster.

Use the following command to perform vertical scaling.

```bash
kbcli cluster vscale milvus --cpu=1 --memory=1Gi --components=milvus 
```

Please wait a few seconds until the scaling process is over.

The `kbcli cluster vscale` command prints a command to help check the progress of scaling operations.

```bash
kbcli cluster describe-ops milvus-verticalscaling-rpw2l -n default
```

To check whether the scaling is done, use the following command.

```bash
kbcli cluster describe milvus
```

## Volume Expansion

***Steps:***

```bash
kbcli cluster volume-expand milvus --storage=40Gi --components=milvus -t data
```

The volume expansion may take a few minutes.

The `kbcli cluster volume-expand` command prints a command to help check the progress of scaling operations.

```bash
kbcli cluster describe-ops milvus-volumeexpansion-5pbd2 -n default
```

To check whether the expanding is done, use the following command.

```bash
kbcli cluster describe milvus
```

## Restart

1. Restart a cluster.

   Configure the values of `components` and `ttlSecondsAfterSucceed` and run the command below to restart a specified cluster.

   ```bash
   kbcli cluster restart milvus --components="milvus" \
   --ttlSecondsAfterSucceed=30
   ```

   - `components` describes the component name that needs to be restarted.
   - `ttlSecondsAfterSucceed` describes the time to live of an OpsRequest job after the restarting succeeds.

2. Validate the restarting.

   Run the command below to check the cluster status to check the restarting status.

   ```bash
   kbcli cluster list milvus
   >
   NAME     NAMESPACE   CLUSTER-DEFINITION     VERSION         TERMINATION-POLICY   STATUS    CREATED-TIME
   milvus   default     milvus-2.3.2           milvus-2.3.2    Delete               Running   Jul 05,2024 18:35 UTC+0800
   ```

   * STATUS=Updating: it means the cluster restart is in progress.
   * STATUS=Running: it means the cluster has been restarted.

## Stop/Start a cluster

You can stop/start a cluster to save computing resources. When a cluster is stopped, the computing resources of this cluster are released, which means the pods of Kubernetes are released, but the storage resources are reserved. You can start this cluster again by snapshots if you want to restore the cluster resources.

### Stop a cluster

1. Configure the name of your cluster and run the command below to stop this cluster.

   ```bash
   kbcli cluster stop milvus
   ```

2. Check the status of the cluster to see whether it is stopped.

    ```bash
    kbcli cluster list
    ```

### Start a cluster

1. Configure the name of your cluster and run the command below to start this cluster.

   ```bash
   kbcli cluster start milvus
   ```

2. Check the status of the cluster to see whether it is running again.

    ```bash
    kbcli cluster list
    ```