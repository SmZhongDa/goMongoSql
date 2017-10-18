/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sqlparser

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

var fla = "0"

var bsonNow bson.M
var bsonFinal bson.M

func Accept(node Expr) bson.M {

	switch node := node.(type) {
	case *AndExpr:
		fmt.Println("进入 AndExpr")
		AndVisit(node)
		return bsonFinal
	case *OrExpr:
		fmt.Println("进入 OrExpr")
		OrVisit(node)
		return bsonFinal
	case *ComparisonExpr:
		fmt.Println("进入 ComparisonExpr")
		EqualsVisit(node)
		return bsonFinal
	case *NotExpr:
		fmt.Println("进入 NotExpr")
		return bsonFinal
	case *ParenExpr:
		fmt.Println("进入 ParenExpr")
		//		fmt.Println((*ParenExpr).Expr)
		Accept(node.Expr)
		return bsonFinal
	case *RangeCond:
		fmt.Println("进入 RangeCond")
		return bsonFinal
	case *IsExpr:
		fmt.Println("进入 IsExpr")
		return bsonFinal
	case *ExistsExpr:
		fmt.Println("进入 ExistsExpr")
		return bsonFinal
	case *SQLVal:
		fmt.Println("进入 SQLVal")
		return bsonFinal
	case *NullVal:
		fmt.Println("进入 NullVal")
		return bsonFinal
	case BoolVal:
		fmt.Println("进入 BoolVal")
		return bsonFinal
	case *ColName:
		fmt.Println("进入 ColName")
		return bsonFinal
	case ValTuple:
		fmt.Println("进入 ValTuple")
		return bsonFinal
	case *Subquery:
		fmt.Println("进入 Subquery")
		return bsonFinal
	case ListArg:
		fmt.Println("进入 ListArg")
		return bsonFinal
	case *BinaryExpr:
		fmt.Println("进入 BinaryExpr")
		return bsonFinal
	case *UnaryExpr:
		fmt.Println("进入 UnaryExpr")
		return bsonFinal
	case *IntervalExpr:
		fmt.Println("进入 IntervalExpr")
		return bsonFinal
	case *CollateExpr:
		fmt.Println("进入 CollateExpr")
		return bsonFinal
	case *FuncExpr:
		fmt.Println("进入 FuncExpr")
		return bsonFinal
	case *CaseExpr:
		fmt.Println("进入 CaseExpr")
		return bsonFinal
	case *ValuesFuncExpr:
		fmt.Println("进入 ValuesFuncExpr")
		return bsonFinal
	case *ConvertExpr:
		fmt.Println("进入 ConvertExpr")
		return bsonFinal
	case *ConvertUsingExpr:
		fmt.Println("进入 ConvertUsingExpr")
		return bsonFinal
	case *MatchExpr:
		fmt.Println("进入 MatchExpr")
		return bsonFinal
	case *GroupConcatExpr:
		fmt.Println("进入 GroupConcatExpr")
		return bsonFinal
	default:
		fmt.Println("进入 default")
		return bsonFinal
	}
}

func AndVisit(node Expr) {
	query := make([]bson.M, 0, 0)
	if fla == "0" {
		fla = "1"
		Accept(node.(*AndExpr).Left)
		fmt.Print("a:")
		query = append(query, bsonNow)
		fmt.Println(query)

		Accept(node.(*AndExpr).Right)
		fmt.Print("b:")
		query = append(query, bsonNow)
		fmt.Println(query)
		bsonFinal = bson.M{"$and": query}
		fmt.Println(bsonFinal)
	} else {
		Accept(node.(*AndExpr).Left)
		query = append(query, bsonNow)
		fmt.Println(query)

		Accept(node.(*AndExpr).Right)
		query = append(query, bsonNow)
		fmt.Println(query)
		bsonNow = bson.M{"$and": query}
	}
}

func OrVisit(node Expr) {
	query := make([]bson.M, 0, 0)
	if fla == "0" {
		fla = "1"
		Accept(node.(*OrExpr).Left)
		fmt.Print("a:")
		query = append(query, bsonNow)
		fmt.Println(query)

		Accept(node.(*OrExpr).Right)
		fmt.Print("b:")
		query = append(query, bsonNow)
		fmt.Println(query)
		bsonFinal = bson.M{"$or": query}
		fmt.Println(bsonFinal)
	} else {
		Accept(node.(*OrExpr).Left)
		query = append(query, bsonNow)
		fmt.Println(query)

		Accept(node.(*OrExpr).Right)
		query = append(query, bsonNow)
		fmt.Println(query)
		bsonNow = bson.M{"$or": query}
	}
}

func EqualsVisit(node Expr) string {
	bufLeft := NewTrackedBuffer(nil)
	//	bufOperator := NewTrackedBuffer(nil)
	bufRight := NewTrackedBuffer(nil)
	node.(*ComparisonExpr).Left.Format(bufLeft)
	bufOperator := node.(*ComparisonExpr).Operator
	node.(*ComparisonExpr).Right.Format(bufRight)
	if fla == "0" {
		fla = "1"
		switch bufOperator {
		case "=":
			bsonFinal = bson.M{bufLeft.String(): strings.Trim(bufRight.String(), "'")}
			fmt.Println(bsonFinal)
			return "="
		case "!=":
			bsonFinal = bson.M{bufLeft.String(): bson.M{"$ne": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonFinal)
			return "!="
		case ">":
			bsonFinal = bson.M{bufLeft.String(): bson.M{"$gt": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonFinal)
			return ">"
		case ">=":
			bsonFinal = bson.M{bufLeft.String(): bson.M{"$gte": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonFinal)
			return ">="
		case "<":
			bsonFinal = bson.M{bufLeft.String(): bson.M{"$lt": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonFinal)
			return "<"
		case "<=":
			bsonFinal = bson.M{bufLeft.String(): bson.M{"$lte": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonFinal)
			return "<="
		case "in":
			bsonFinal = bson.M{bufLeft.String(): bson.M{"$in": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonFinal)
			return "in"
		default:
			return "default"
		}

	} else {
		switch bufOperator {
		case "=":
			bsonNow = bson.M{bufLeft.String(): strings.Trim(bufRight.String(), "'")}
			fmt.Println(bsonNow)
			return "="
		case "!=":
			bsonNow = bson.M{bufLeft.String(): bson.M{"$ne": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonNow)
			return "!="
		case ">":
			bsonNow = bson.M{bufLeft.String(): bson.M{"$gt": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonNow)
			return ">"
		case ">=":
			bsonNow = bson.M{bufLeft.String(): bson.M{"$gte": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonNow)
			return ">="
		case "<":
			bsonNow = bson.M{bufLeft.String(): bson.M{"$lt": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonNow)
			return "<"
		case "<=":
			bsonNow = bson.M{bufLeft.String(): bson.M{"$lte": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonNow)
			return "<="
		case "in":
			bsonNow = bson.M{bufLeft.String(): bson.M{"$in": strings.Trim(bufRight.String(), "'")}}
			fmt.Println(bsonNow)
			return "in"
		default:
			return "default"
		}
	}

	return "null"

}
