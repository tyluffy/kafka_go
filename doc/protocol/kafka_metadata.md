### version 9
```
Metadata Response (Version: 9) => throttle_time_ms [brokers] cluster_id controller_id [topics] cluster_authorized_operations TAG_BUFFER 
  throttle_time_ms => INT32
  brokers => node_id host port rack TAG_BUFFER 
    node_id => INT32
    host => COMPACT_STRING
    port => INT32
    rack => COMPACT_NULLABLE_STRING
  cluster_id => COMPACT_NULLABLE_STRING
  controller_id => INT32
  topics => error_code name is_internal [partitions] topic_authorized_operations TAG_BUFFER 
    error_code => INT16
    name => COMPACT_STRING
    is_internal => BOOLEAN
    partitions => error_code partition_index leader_id leader_epoch [replica_nodes] [isr_nodes] [offline_replicas] TAG_BUFFER 
      error_code => INT16
      partition_index => INT32
      leader_id => INT32
      leader_epoch => INT32
      replica_nodes => INT32
      isr_nodes => INT32
      offline_replicas => INT32
    topic_authorized_operations => INT32
  cluster_authorized_operations => INT32
```
### version 11
```
Metadata Request (Version: 11) => [topics] allow_auto_topic_creation include_topic_authorized_operations TAG_BUFFER 
  topics => topic_id name TAG_BUFFER 
    topic_id => UUID
    name => COMPACT_NULLABLE_STRING
  allow_auto_topic_creation => BOOLEAN
  include_topic_authorized_operations => BOOLEAN
```
```
Metadata Response (Version: 11) => throttle_time_ms [brokers] cluster_id controller_id [topics] TAG_BUFFER 
  throttle_time_ms => INT32
  brokers => node_id host port rack TAG_BUFFER 
    node_id => INT32
    host => COMPACT_STRING
    port => INT32
    rack => COMPACT_NULLABLE_STRING
  cluster_id => COMPACT_NULLABLE_STRING
  controller_id => INT32
  topics => error_code name topic_id is_internal [partitions] topic_authorized_operations TAG_BUFFER 
    error_code => INT16
    name => COMPACT_STRING
    topic_id => UUID
    is_internal => BOOLEAN
    partitions => error_code partition_index leader_id leader_epoch [replica_nodes] [isr_nodes] [offline_replicas] TAG_BUFFER 
      error_code => INT16
      partition_index => INT32
      leader_id => INT32
      leader_epoch => INT32
      replica_nodes => INT32
      isr_nodes => INT32
      offline_replicas => INT32
    topic_authorized_operations => INT32
```