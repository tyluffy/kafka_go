### version 3
```
FindCoordinator Request (Version: 3) => key key_type TAG_BUFFER 
  key => COMPACT_STRING
  key_type => INT8
```

```
FindCoordinator Response (Version: 3) => throttle_time_ms error_code error_message node_id host port TAG_BUFFER 
  throttle_time_ms => INT32
  error_code => INT16
  error_message => COMPACT_NULLABLE_STRING
  node_id => INT32
  host => COMPACT_STRING
  port => INT32
```