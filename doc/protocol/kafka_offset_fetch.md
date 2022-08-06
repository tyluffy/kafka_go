### version 6
```
OffsetFetch Request (Version: 6) => group_id [topics] TAG_BUFFER 
  group_id => COMPACT_STRING
  topics => name [partition_indexes] TAG_BUFFER 
    name => COMPACT_STRING
    partition_indexes => INT32
```
```
OffsetFetch Response (Version: 6) => throttle_time_ms [topics] error_code TAG_BUFFER 
  throttle_time_ms => INT32
  topics => name [partitions] TAG_BUFFER 
    name => COMPACT_STRING
    partitions => partition_index committed_offset committed_leader_epoch metadata error_code TAG_BUFFER 
      partition_index => INT32
      committed_offset => INT64
      committed_leader_epoch => INT32
      metadata => COMPACT_NULLABLE_STRING
      error_code => INT16
  error_code => INT16
```
### version 7
```
OffsetFetch Request (Version: 7) => group_id [topics] require_stable TAG_BUFFER
  group_id => COMPACT_STRING
  topics => name [partition_indexes] TAG_BUFFER
    name => COMPACT_STRING
    partition_indexes => INT32
  require_stable => BOOLEAN
```

```
OffsetFetch Response (Version: 7) => throttle_time_ms [topics] error_code TAG_BUFFER
  throttle_time_ms => INT32
  topics => name [partitions] TAG_BUFFER
    name => COMPACT_STRING
    partitions => partition_index committed_offset committed_leader_epoch metadata error_code TAG_BUFFER
      partition_index => INT32
      committed_offset => INT64
      committed_leader_epoch => INT32
      metadata => COMPACT_NULLABLE_STRING
      error_code => INT16
  error_code => INT16
```
