# Proto-filter
Proto-filter is a [protobuf](https://developers.google.com/protocol-buffers/) pre-processor that allows you to filter out items in a proto file based on terms. This can be useful if you have a service that has private methods that you wish to keep hidden from certain clients. Most of what this tool does can also be achieved by just structuring your proto files differently. It is mostly just an exercise for me to better understand the internals of protocol buffers.


## Filtering Rules
The tool will filter items according to the following logic (in this exact order of priority):
* Items without the option present will be included in the output (never filtered)
* Items that have the filtered term marked as `excluded` will be removed
* Items that have the filtered term marked as `included` will be included in the output

This means that an exclude rule will take priority over an include rule in case there is a conflict.

## Example Usage
Consider the following `test.proto` file

```proto
syntax = "proto3";

package com.test;

import "filter/filter.proto"; // Adds filter option

option go_package = "example.protobuf.test";

message Test {
    string jp_string = 1 [(filter.field) = {exclude: ["NA"]}];
    string na_string = 2 [(filter.field) = {include: ["NA"]}];
    Empty nothing = 3;
}
```

Processed by the following command

```bash
proto-filter -i . -term NA test.proto
```

Will result in

```proto
syntax = "proto3";

package com.test;

import "filter/filter.proto"; // Adds filter option

option go_package = "example.protobuf.test";

message Test {
    string na_string = 2 [(filter.field) = {include: ["NA"]}];
    Empty nothing = 3;
}
```