### TODO: escape input? Does redis have same risks as SQL injection?

#### No. The Redis protocol has no concept of string escaping, so injection is impossible under normal circumstances using a normal client library. The protocol uses prefixed-length strings and is completely binary safe.


