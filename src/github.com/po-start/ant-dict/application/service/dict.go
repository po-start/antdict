package service

import (
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/po-start/ant-dict/core/db/mysql"

//	"github.com/po-start/ant-dict/application/entity"
)

type DictService struct{}
/*
func ( d *DictService ) Search ( key string ) interface{} {



	var data []map[string]interface{}
	if err := json.Unmarshal(responseData, &data); err != nil {
		fmt.Prentln(err.Error())
	}
	return data
}
*/

func (d *DictService) Search (key string) interface{} {
	data := make([]map[string]interface{}, 0, 1)
//	t := make(map[string]interface{})
//	t["b"] = "ttt"
//	t["ss"] = 123123
//	t["bs"] = mysql.TT()

a := mysql.TT(key).([]map[string]string)

for _, v := range a {
//	fmt.Printf("key: %d, v:%s \n", key, v)
		item := make(map[string]interface{})
        var content map[string]interface{}     
        if err := json.Unmarshal([]byte(v["content"]), &content); err != nil {
            fmt.Println(err.Error())
        }
		id, _ := strconv.Atoi(v["id"])  
		item["id"] = id 
		item["word"] = v["word"]
		item["content"] = content
		data = append(data, item)

	//for k, value := range v {
	//	fmt.Printf("------ key: %s, v: $s\n", k, value)
	//}
}

//data = append(data, t)
//data = append(data, t)
//data = append(data, t)
	return data
}

