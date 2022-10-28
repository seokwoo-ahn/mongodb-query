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
