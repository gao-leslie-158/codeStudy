package snowflake

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// 生成64位的雪花ID
func GenID() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2023-03-18", 1); err != nil {
//		fmt.Println("snowflake Init() failed, error：", err)
//		return
//	}
//	id := GenID()
//	fmt.Println("id:", id)
//}
