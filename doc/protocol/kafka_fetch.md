### version 11
```
Fetch Request (Version: 11) => replica_id max_wait_ms min_bytes max_bytes isolation_level session_id session_epoch [topics] [forgotten_topics_data] rack_id 
  replica_id => INT32
  max_wait_ms => INT32
  min_bytes => INT32
  max_bytes => INT32
  isolation_level => INT8
  session_id => INT32
  session_epoch => INT32
  topics => topic [partitions] 
    topic => STRING
    partitions => partition current_leader_epoch fetch_offset log_start_offset partition_max_bytes 
      partition => INT32
      current_leader_epoch => INT32
      fetch_offset => INT64
      log_start_offset => INT64
      partition_max_bytes => INT32
  forgotten_topics_data => topic [partitions] 
    topic => STRING
    partitions => INT32
  rack_id => STRING
```
```
Fetch Response (Version: 11) => throttle_time_ms error_code session_id [responses] 
  throttle_time_ms => INT32
  error_code => INT16
  session_id => INT32
  responses => topic [partition_responses] 
    topic => STRING
    partition_responses => partition error_code high_watermark last_stable_offset log_start_offset [aborted_transactions] preferred_read_replica record_set 
      partition => INT32
      error_code => INT16
      high_watermark => INT64
      last_stable_offset => INT64
      log_start_offset => INT64
      aborted_transactions => producer_id first_offset 
        producer_id => INT64
        first_offset => INT64
      preferred_read_replica => INT32
      record_set => RECORDS
```
### version 12
```
Fetch Request (Version: 12) => replica_id max_wait_ms min_bytes max_bytes isolation_level session_id session_epoch [topics] [forgotten_topics_data] rack_id TAG_BUFFER 
  replica_id => INT32
  max_wait_ms => INT32
  min_bytes => INT32
  max_bytes => INT32
  isolation_level => INT8
  session_id => INT32
  session_epoch => INT32
  topics => topic [partitions] TAG_BUFFER 
    topic => COMPACT_STRING
    partitions => partition current_leader_epoch fetch_offset last_fetched_epoch log_start_offset partition_max_bytes TAG_BUFFER 
      partition => INT32
      current_leader_epoch => INT32
      fetch_offset => INT64
      last_fetched_epoch => INT32
      log_start_offset => INT64
      partition_max_bytes => INT32
  forgotten_topics_data => topic [partitions] TAG_BUFFER 
    topic => COMPACT_STRING
    partitions => INT32
  rack_id => COMPACT_STRING
```
```
Fetch Response (Version: 12) => throttle_time_ms error_code session_id [responses] TAG_BUFFER 
  throttle_time_ms => INT32
  error_code => INT16
  session_id => INT32
  responses => topic [partition_responses] TAG_BUFFER 
    topic => COMPACT_STRING
    partition_responses => partition error_code high_watermark last_stable_offset log_start_offset [aborted_transactions] preferred_read_replica record_set TAG_BUFFER 
      partition => INT32
      error_code => INT16
      high_watermark => INT64
      last_stable_offset => INT64
      log_start_offset => INT64
      aborted_transactions => producer_id first_offset TAG_BUFFER 
        producer_id => INT64
        first_offset => INT64
      preferred_read_replica => INT32
      record_set => COMPACT_RECORDS
```
