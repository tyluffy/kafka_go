### version 0
```
SaslAuthenticate Request (Version: 0) => auth_bytes 
  auth_bytes => BYTES
```
```
SaslAuthenticate Response (Version: 0) => error_code error_message auth_bytes 
  error_code => INT16
  error_message => NULLABLE_STRING
  auth_bytes => BYTES
```
### version 2
```
SaslAuthenticate Request (Version: 2) => auth_bytes TAG_BUFFER 
  auth_bytes => COMPACT_BYTES
```
```
SaslAuthenticate Response (Version: 2) => error_code error_message auth_bytes session_lifetime_ms TAG_BUFFER 
  error_code => INT16
  error_message => COMPACT_NULLABLE_STRING
  auth_bytes => COMPACT_BYTES
  session_lifetime_ms => INT64
```
