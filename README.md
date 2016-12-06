# proto-gen-md

protofile convert markdown.

this is under development project.

## Useage

#### Prepare protofile

```proto
syntax = "proto3";

package hoge;

service Player {
  rpc Info (PlayerInfoRequest) returns (PlayerInfoResponse) {};
  rpc Entry (PlayerEntryRequest) returns (PlayerEntryResponse) {};
  rpc Comment (PlayerCommentRequest) returns (PlayerCommentResponse) {};

}

message PlayerInfoRequest {
    uint64 player_id = 1;
}

message PlayerInfoResponse {
    uint64 id = 1;
    string name = 2;
    uint64 money = 3;
    int64 last_logined_at = 4;
}

message PlayerEntryRequest {
    uint64 player_id = 1;
}

message PlayerEntryResponse {
    uint64 player_id = 1;
    repeated uint32 entry_id = 2;
}


message PlayerComment {
    uint64 id = 1;
    uint32 entry_id = 2;
    string comment = 3;
    int64 created_at = 4;
    int64 updated_at = 5;
}

message PlayerCommentRequest {
    uint64 player_id = 1;
}

message PlayerCommentResponse {
    uint64 id = 1;
    repeated PlayerComment player_comment = 2;
}
```

#### Exec command

```shell
protoc --md_out=./doc ./player.proto
```

#### Generated Markdown Document

```
# Document

## Index

### Player

  - [/api/player/info](#player_info)
  - [/api/player/entry](#player_entry)
  - [/api/player/comment](#player_comment)

## Detail

## Player

### <a name="player_info">/api/player/info</a>

#### Request Method

POST

#### Request Parameter: PlayerInfoRequest

|key|type|
|:--|:--|
|player_id|uint64|


#### Response Parameter: PlayerInfoResponse

|key|type|
|:--|:--|
|id|uint64|
|name|string|
|money|uint64|
|last_logined_at|int64|

### <a name="player_entry">/api/player/entry</a>

#### Request Method

POST

#### Request Parameter: PlayerEntryRequest

|key|type|
|:--|:--|
|player_id|uint64|


#### Response Parameter: PlayerEntryResponse

|key|type|
|:--|:--|
|player_id|uint64|
|entry_id|uint32|

### <a name="player_comment">/api/player/comment</a>

#### Request Method

POST

#### Request Parameter: PlayerCommentRequest

|key|type|
|:--|:--|
|player_id|uint64|


#### Response Parameter: PlayerCommentResponse

|key|type|
|:--|:--|
|id|uint64|
|player_comment|hoge.PlayerComment|

```
