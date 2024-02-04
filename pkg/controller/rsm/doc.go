/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/*
Package rsm is a general component aims to hold role-based stateful workloads(such as databases).
RSM stands for Replicated State Machine based on the truth that the workloads are solving state replication related problems.
Treat RSM as an enhanced StatefulSet.

The K8s native StatefulSet can handle stateful workloads well,
but there are more works to do if the workload pods have roles(leader/follower in etcd, primary/secondary in PostgreSQL etc.).

RSM adds an abstract layer above StatefulSet, and provides:
1. role-based update strategy(Serial/Parallel/BestEffortParallel)
2. role-based access modes(ReadWrite/Readonly/None)
3. auto switchover
4. membership reconfiguration
*/
package rsm
