# goMongoSql
/*

docker-images:存放docker镜像


hello:入口程序


sqlparse:sql解析程序
*/





/*
		mongodb版本：MongoDB shell version v3.4.9
		
		测试数据：
		
		[2017-10-06 01:06:12](127.0.0.1:27888/test)>db.list.find()
		
		{ "_id" : "1", "type" : "city", "sheng" : "jiangsu", "name" : "nanjing" }
		
		{ "_id" : "2", "type" : "city", "sheng" : "jiangsu", "name" : "xuzhou" }
		
		{ "_id" : "3", "type" : "city", "sheng" : "jiangsu", "name" : "suzhou" }
		
		{ "_id" : "4", "type" : "city", "sheng" : "china", "name" : "beijing" }
		
		{ "_id" : "5", "type" : "city", "sheng" : "china", "name" : "shanghai" }
		
		{ "_id" : "6", "type" : "city", "sheng" : "zhejiang", "name" : "hangzhou" }
		
		{ "_id" : "7", "type" : "city", "sheng" : "zhejiang", "name" : "wuzhen" }
		
	*/
 支持的sql类型:
 
	//	sql := "select type,sheng,name from tt where _id = 'nanjing' and type = 'food' and item = 'card' and qty = 'xgx'"
	
	//	sql := "select type,sheng,name from tt where type = 'city' or name = 'nanjing'"
	
	//	sql := "select type,sheng,name from tt where (type = 'city' and name = 'nanjing') or (type = 'city' and name = 'shanghai')"
	
	//	sql := "select type,sheng,name from tt where _id >= 1 and _id < 6"
	
	//	sql := "select type,sheng,name from tt where sheng != 'jiangsu' or _id != 4 and name = 'wuzhen' "
  
  
