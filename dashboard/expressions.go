/*
Copyright © 2021 Peter Parada <parad.peter@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package dashboard

// TODO: Implement
// TODO: Test if queries really work by increasing requests limit as well as real increasing of memroy - has to be some docker image which handled=s this
// TODO: Advanced Allow filtering by node/s, node selectors
var ClusterMemoryRate = "sum(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes) or ((node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes) / node_memory_MemTotal_bytes)) / sum(kube_node_info)"

var ClusterMemoryUsed = "sum(node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes)"
var ClusterMemoryTotal = "sum(node_memory_MemTotal_bytes)"

var ClusterCPURate = "sum(1 - avg without (mode,cpu) (rate(node_cpu_seconds_total{mode=\"idle\"}[5m]))) / sum(kube_node_info)"

var ClusterCoresUsedCount = ""
var ClusterCoresTotalCount = ""

// var ClusterDiskUsageRate = ""
// var ClusterDiskUsageUsed = ""
// var ClusterDiskUsageTotal = ""

// TODO: Advanced Nginx metrics
// var NginxRequestsPerSecondRate = ""
// var NginxResponseTimeRate = ""
