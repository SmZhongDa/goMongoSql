// hello project main.go
package main

import (
	"fmt"

	"bufio"
	"os"

	"github.com/xwb1989/sqlparser"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Entr struct {
	Key   string
	Value string
}

type Result struct {
	Entr []Entr
}

type Element struct {
	Columns     [10]string
	TableName   string
	Condition   bson.M
	bsonColumns bson.M
}

const (
	URL = "192.168.0.106:27017"

//	URL = "127.0.0.1:27017"
)

type Entry struct {
	Id_   string `bson:"_id"`
	Type  string `bson:"type"`
	Sheng string `bson:"sheng"`
	Name  string `bson:"name"`
}

type Collection struct {
	Entrys []Entry
}

func getMongoData(elemnt *Element) {
	session, err := mgo.Dial(URL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")             //数据库名称
	collection := db.C(elemnt.TableName) //如果该集合已经存在的话，则直接返回

	result := Entry{}
	var entryAll Collection
	iter := collection.Find(elemnt.Condition).Select(bson.M{"sheng": 1, "name": 1, "_id": 1}).Iter()
	i := 0
	for iter.Next(&result) {
		i = i + 1
		fmt.Printf("-----record:%v \n", i)
		fmt.Printf("Result:%v \n", result.Id_)
		fmt.Printf("Result:%v \n", result.Sheng)
		fmt.Printf("Result:%v \n", result.Name)
		entryAll.Entrys = append(entryAll.Entrys, result)
	}
}

func getElement(sql string) *Element {
	tree, err := sqlparser.Parse(sql)
	if err != nil {
		fmt.Println(err)
	}
	element := new(Element)
	//sql
	//	fmt.Println(sqlparser.String(tree))

	//获取limit字段
	//	fmt.Println("获取limit字段:")
	//	buf := sqlparser.NewTrackedBuffer(nil)
	//	tree.(*(sqlparser.Select)).Limit.Rowcount.Format(buf)
	//	fmt.Println(buf)

	//获取group by 字段

	//	tableIden := tree.(*(sqlparser.Select)).SelectExprs[0].(*(sqlparser.AliasedExpr)).Expr.(*(sqlparser.ColName)).Name
	//	fmt.Println(tableIden.String())
	//	element.Columns[0] = tableIden.String()

	//	element.Type = tableIden
	//	fmt.Println(element.Type)

	//	element := new(Element)

	//获取表名
	fmt.Println("获取表名:")
	tableIdent := sqlparser.GetTableName(tree.(*(sqlparser.Select)).From[0].(*(sqlparser.AliasedTableExpr)).Expr)
	element.TableName = tableIdent.String()
	fmt.Println(element.TableName)

	//获取select字段
	//	columnsTemp := make([]bson.M, 0, 0)
	fmt.Println("获取select字段:")
	num := len(tree.(*(sqlparser.Select)).SelectExprs)

	for i := 0; i < num; i++ {
		selec := tree.(*(sqlparser.Select)).SelectExprs[i].(*(sqlparser.AliasedExpr)).Expr.(*(sqlparser.ColName)).Name

		element.Columns[i] = selec.String()
		fmt.Println(selec.String())
	}

	//获取where条件
	fmt.Println("获取where条件:")
	fmt.Println("开始Accept：")
	expr := sqlparser.Accept(tree.(*(sqlparser.Select)).Where.Expr)
	fmt.Println("结束Accept")
	element.Condition = expr
	fmt.Println(element.Condition)

	return element
}

func main() {
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

	//	sql := "select type,sheng,name from tt where _id = 'nanjing' and type = 'food' and item = 'card' and qty = 'xgx'"
	//	sql := "select type,sheng,name from tt where type = 'city' or name = 'nanjing'"
	//	sql := "select type,sheng,name from tt where (type = 'city' and name = 'nanjing') or (type = 'city' and name = 'shanghai')"
	//	sql := "select type,sheng,name from tt where _id >= 1 and _id < 6"
	//	sql := "select type,sheng,name from tt where sheng != 'jiangsu' or _id != 4 and name = 'wuzhen' "

	fmt.Println("please enter sql:")
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	sql := string(data)
	fmt.Println(sql)
	en := getElement(sql)
	fmt.Print("最终的返回值：")
	fmt.Println(en)
	fmt.Println("==========================")
	getMongoData(en)

}
