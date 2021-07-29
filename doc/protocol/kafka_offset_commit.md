```
OffsetCommit Request (Version: 8) => group_id generation_id member_id group_instance_id [topics] TAG_BUFFER 
  group_id => COMPACT_STRING
  generation_id => INT32
  member_id => COMPACT_STRING
  group_instance_id => COMPACT_NULLABLE_STRING
  topics => name [partitions] TAG_BUFFER 
    name => COMPACT_STRING
    partitions => partition_index committed_offset committed_leader_epoch committed_metadata TAG_BUFFER 
      partition_index => INT32
      committed_offset => INT64
      committed_leader_epoch => INT32
      committed_metadata => COMPACT_NULLABLE_STRING
```

```
OffsetCommit Response (Version: 8) => throttle_time_ms [topics] TAG_BUFFER 
  throttle_time_ms => INT32
  topics => name [partitions] TAG_BUFFER 
    name => COMPACT_STRING
    partitions => partition_index error_code TAG_BUFFER 
      partition_index => INT32
      error_code => INT16
```