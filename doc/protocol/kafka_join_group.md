```
JoinGroup Request (Version: 7) => group_id session_timeout_ms rebalance_timeout_ms member_id group_instance_id protocol_type [protocols] TAG_BUFFER 
  group_id => COMPACT_STRING
  session_timeout_ms => INT32
  rebalance_timeout_ms => INT32
  member_id => COMPACT_STRING
  group_instance_id => COMPACT_NULLABLE_STRING
  protocol_type => COMPACT_STRING
  protocols => name metadata TAG_BUFFER 
    name => COMPACT_STRING
    metadata => COMPACT_BYTES
```

```
JoinGroup Response (Version: 7) => throttle_time_ms error_code generation_id protocol_type protocol_name leader member_id [members] TAG_BUFFER 
  throttle_time_ms => INT32
  error_code => INT16
  generation_id => INT32
  protocol_type => COMPACT_NULLABLE_STRING
  protocol_name => COMPACT_NULLABLE_STRING
  leader => COMPACT_STRING
  member_id => COMPACT_STRING
  members => member_id group_instance_id metadata TAG_BUFFER 
    member_id => COMPACT_STRING
    group_instance_id => COMPACT_NULLABLE_STRING
    metadata => COMPACT_BYTES
```