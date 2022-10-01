# elastic-search
### 创建一个索引
```js
http://127.0.0.1:9200/info/_doc/1
{
    "name":"hello",
    "age":22,
    "hobby":"look up book",
    "id":123
}
```
### 查看健康状态
```js
http://127.0.0.1:9200/_cat/health
```
### 查看索引
```js
http://127.0.0.1:9200/_cat/indices
```
### 查询数据
```js
http://127.0.0.1:9200/user/_search
```

### 分页请求
```js
http://127.0.0.1:9200/haoke/user/_search?size=1&from=1
```
