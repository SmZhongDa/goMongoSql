#goMongoSql
==========
##Notice
-------
goMongoSql 是在开源项目sqlparser（https://github.com/xwb1989/sqlparser ）基础上针对Mongodb sql的二次开发，<br>
sqlparser负责将标准的sql语句解析成语法树，可以从语法树中获取select字段、表名、where条件等sql语句组成部分，并且<br>
可以将where条件解析成最小的元素，比如"where id = 1" 可以解析为Left(id)、Operator(=)、Right(1)。这样就可以将<br>
这种最小的元素拼接成任何数据库的语法<\br>

##Usage
-------
main.go:程序入口<br>
sqlVist.go:mogodb语法拼接，需要将该文件放到sqlparser包中，供main.go调用<\br>

##TEST
-------

		mongodb版本：MongoDB shell version v3.4.9<br>
###测试数据：<br>
```
[2017-10-06 01:06:12](127.0.0.1:27888/test)>db.list.find()<br>
{ "_id" : "1", "type" : "city", "sheng" : "jiangsu", "name" : "nanjing" }<br>
{ "_id" : "2", "type" : "city", "sheng" : "jiangsu", "name" : "xuzhou" }<br>
{ "_id" : "3", "type" : "city", "sheng" : "jiangsu", "name" : "suzhou" }<br>
{ "_id" : "4", "type" : "city", "sheng" : "china", "name" : "beijing" }<br>
{ "_id" : "5", "type" : "city", "sheng" : "china", "name" : "shanghai" }<br>
{ "_id" : "6", "type" : "city", "sheng" : "zhejiang", "name" : "hangzhou" }<br>
{ "_id" : "7", "type" : "city", "sheng" : "zhejiang", "name" : "wuzhen" }<br>
```

###支持的sql类型:
```
sql := "select type,sheng,name from tt where _id = 1 and type = 'city' and sheng = 'jiangsu' and name = 'nanjing'"<br>
sql := "select type,sheng,name from tt where type = 'city' or name = 'nanjing'"<br>
sql := "select type,sheng,name from tt where (type = 'city' and name = 'nanjing') or (type = 'city' and name = 'shanghai')"<br>
sql := "select type,sheng,name from tt where _id >= 1 and _id < 6"<br>
sql := "select type,sheng,name from tt where sheng != 'jiangsu' or _id != 4 and name = 'wuzhen' "<br>
```
  
###流程
1.运行Mongodb容器
```
xgx@xgx-virtual-machine:~$ sudo docker images

REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
<none>              <none>              14be7ede2af7        3 minutes ago       132 MB	
ubunt_v2            latest              831fde02fb68        40 minutes ago      127 MB
mongo               3.4_v1              9cc2afe1ded3        7 hours ago         360.9 MB
ubuntu              16.04               747cb2d60bbe        7 days ago          122 MB
ubuntu              latest              747cb2d60bbe        7 days ago          122 MB
mongo               3.4                 42e262dc0845        8 days ago          360.9 MB

xgx@xgx-virtual-machine:~$ sudo docker  run -p 27017:27017 -v $PWD/db:/data/db -d 9cc2afe1ded3
db50fb249e58e7201f6e5f8f627a35353d52cdee2ed40a4ae465a9b9bfe3f66c  
```
2.运行go解析脚本
```
xgx@xgx-virtual-machine:~$ sudo docker run -ti  831fde02fb68
```

例子1：
```
root@42daacbd6c86:/# /home/hello

please enter sql:

select id,name from tt where name = 'nanjing'

select id, name from tt where name = 'nanjing'tt

获取select字段:

id

name

获取where条件:

进入 ComparisonExpr

&&&&&&&&&&&&

map[name:nanjing]

map[name:nanjing]

最终的返回值：&{[id name        ] tt map[name:nanjing] map[]}

==========================

-----record:1 

Result:1 

Result:jiangsu 

Result:nanjing 
```


例子2：
```
root@42daacbd6c86:/# /home/hello
please enter sql:
select id,name from tt where _id >= 1 and _id < 6
select id, name from tt where _id >= 1 and _id < 6tt
获取select字段:
id
name
获取where条件:
进入 AndExpr
进入 ComparisonExpr
&&&&&&&&&&&&
map[_id:map[$gte:1]]
a:[map[_id:map[$gte:1]]]
进入 ComparisonExpr
&&&&&&&&&&&&
map[_id:map[$lt:6]]
b:[map[_id:map[$gte:1]] map[_id:map[$lt:6]]]
map[$and:[map[_id:map[$gte:1]] map[_id:map[$lt:6]]]]
map[$and:[map[_id:map[$gte:1]] map[_id:map[$lt:6]]]]
最终的返回值：&{[id name        ] tt map[$and:[map[_id:map[$gte:1]] map[_id:map[$lt:6]]]] map[]}
==========================
-----record:1 
Result:1 
Result:jiangsu 
Result:nanjing 
-----record:2 
Result:2 
Result:jiangsu 
Result:xuzhou 
-----record:3 
Result:3 
Result:jiangsu 
Result:suzhou 
-----record:4 
Result:4 
Result:china 
Result:beijing 
-----record:5 
Result:5 
Result:china 
Result:shanghai 
```

例子3：
```
root@42daacbd6c86:/# /home/hello
please enter sql:
select id,name from tt where name = 'nanjing' or _id != 4 and name = 'wuzhen'
select id, name from tt where name = 'nanjing' or _id != 4 and name = 'wuzhen'tt
获取select字段:
id
name
获取where条件:
进入 OrExpr
进入 ComparisonExpr
&&&&&&&&&&&&
map[name:nanjing]
a:[map[name:nanjing]]
进入 AndExpr
进入 ComparisonExpr
&&&&&&&&&&&&
map[_id:map[$ne:4]]
[map[_id:map[$ne:4]]]
进入 ComparisonExpr
&&&&&&&&&&&&
map[name:wuzhen]
[map[_id:map[$ne:4]] map[name:wuzhen]]
b:[map[name:nanjing] map[$and:[map[_id:map[$ne:4]] map[name:wuzhen]]]]
map[$or:[map[name:nanjing] map[$and:[map[_id:map[$ne:4]] map[name:wuzhen]]]]]
map[$or:[map[name:nanjing] map[$and:[map[_id:map[$ne:4]] map[name:wuzhen]]]]]
最终的返回值：&{[id name        ] tt map[$or:[map[name:nanjing] map[$and:[map[_id:map[$ne:4]] map[name:wuzhen]]]]] map[]}
==========================
-----record:1 
Result:1 
Result:jiangsu 
Result:nanjing 
-----record:2 
Result:7 
Result:zhejiang 
Result:wuzhen 
```

例子4：
```
root@42daacbd6c86:/# /home/hello
please enter sql:
select id,name from tt where (type = 'city' and name = 'nanjing') or (type = 'city' and name = 'shanghai')
select id, name from tt where (type = 'city' and name = 'nanjing') or (type = 'city' and name = 'shanghai')tt
获取select字段:
id
name
获取where条件:
进入 OrExpr
进入 ParenExpr
进入 AndExpr
进入 ComparisonExpr
&&&&&&&&&&&&
map[type:city]
[map[type:city]]
进入 ComparisonExpr
&&&&&&&&&&&&
map[name:nanjing]
[map[type:city] map[name:nanjing]]
a:[map[$and:[map[type:city] map[name:nanjing]]]]
进入 ParenExpr
进入 AndExpr
进入 ComparisonExpr
&&&&&&&&&&&&
map[type:city]
[map[type:city]]
进入 ComparisonExpr
&&&&&&&&&&&&
map[name:shanghai]
[map[type:city] map[name:shanghai]]
b:[map[$and:[map[type:city] map[name:nanjing]]] map[$and:[map[type:city] map[name:shanghai]]]]
map[$or:[map[$and:[map[type:city] map[name:nanjing]]] map[$and:[map[type:city] map[name:shanghai]]]]]
map[$or:[map[$and:[map[type:city] map[name:nanjing]]] map[$and:[map[type:city] map[name:shanghai]]]]]
最终的返回值：&{[id name        ] tt map[$or:[map[$and:[map[type:city] map[name:nanjing]]] map[$and:[map[type:city] map[name:shanghai]]]]] map[]}
==========================
-----record:1 
Result:1 
Result:jiangsu 
Result:nanjing 
-----record:2 
Result:5 
Result:china 
Result:shanghai 
```
 
