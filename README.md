# mongodb-query
## 실행 방법
### prerequisites
* config.toml
```toml
[DataBase]
dataSource = {your mongoDB url}
username = {user name}
password = {password} // optional
db = {mongoDB database name}
blockCollection = {block data collection name}
txCollection = {tx data collection name}
eventCollection = {event data collection name}
```
toml 파일을 실행파일과 같은 경로에 넣는다.
### 실행
main 함수 위치한 곳에서 go build 후 실행파일 생성
toml 파일이 실행파일과 같은 경로가 아니라면
```
./mongodb-query -config {상대경로 or 절대경로}
```
<img width="677" alt="스크린샷 2022-10-31 오후 3 03 49" src="https://user-images.githubusercontent.com/84396735/198942412-b97cdf5d-3ab1-41cf-b13f-213338f8825e.png">

## mongo db collections
### block (blockCollection)
* schema
```json
"properties": {
  "_id": { "bsonType": "objectId" },
  "blockhash": { "bsonType": "string" },
  "blocknumber": { "bsonType": "double" },
  "blocksize": { "bsonType": "double" },
  "gasused": { "bsonType": "double" },
  "parenthash": { "bsonType": "string" },
  "time": { "bsonType": "double" },
  "totaltxs": { "bsonType": "double" }
}
```

### transaction (txCollection)
* schema
```json
"properties": {
  "_id": { "bsonType": "objectId" },
  "amount": { "bsonType": "string" },
  "blocknumber": { "bsonType": "double" },
  "from": { "bsonType": "string" },
  "gasprice": { "bsonType": "double" },
  "gasused": { "bsonType": "double" },
  "nonce": { "bsonType": "double" },
  "status": { "bsonType": "double" },
  "time": { "bsonType": "double" },
  "to": { "bsonType": "string" },
  "totalindex": { "bsonType": "double" },
  "txhash": { "bsonType": "string" },
  "txindex": { "bsonType": "double" },
  "txsize": { "bsonType": "double" }
}
```

### event (eventCollection)
* schema
```json
"properties": {
  "_id": { "bsonType": "objectId" },
  "blocknumber": { "bsonType": "double" },
  "contract": { "bsonType": "string" },
  "contractaddress": { "bsonType": "string" },
  "data": { "bsonType": "document" },
  "event": { "bsonType": "string" },
  "eventfunc": { "bsonType": "string" },
  "logindex": { "bsonType": "double" },
  "removed": { "bsonType": "boolean" },
  "time": { "bsonType": "double" },
  "txhash": { "bsonType": "string" },
  "txindex": { "bsonType": "double" }
}
```

## 지원기능
### 공통 기능
* Index: collection의 index 정보 조회

### tx
* ByHash: 입력한 transaction 해시 값으로 transaction searching
* ByBNGT: 입력한 숫자보다 큰 블록넘버의 블록에 포함된 모든 transaction searching

### block
* ByHash: block 해시 값으로 block searching
* ByBNGT: 입력한 숫자보다 큰 블록넘버의 블록 searching

### event
* ByEventName: 입력한 이벤트 이름으로 모든 이벤트 searching
* ByContractName: 입력한 컨트랙트 이름에서 발생한 모든 이벤트 searching
* ByContractAddress: 입력한 컨트랙트 주소에서 발생한 모든 이벤트 searching
* ByBlockNum: 입력한 블록 넘버의 블록에 포함된 모든 이벤트 searching
* ByTxHash: 입력한 transaction 해시 값의 transaction에 포함된 모든 이벤트 searching
