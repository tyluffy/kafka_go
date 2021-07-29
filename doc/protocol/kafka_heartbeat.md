```
Heartbeat Request (Version: 4) => group_id generation_id member_id group_instance_id TAG_BUFFER 
  group_id => COMPACT_STRING
  generation_id => INT32
  member_id => COMPACT_STRING
  group_instance_id => COMPACT_NULLABLE_STRING
```

```
Heartbeat Response (Version: 4) => throttle_time_ms error_code TAG_BUFFER 
  throttle_time_ms => INT32
  error_code => INT16
```