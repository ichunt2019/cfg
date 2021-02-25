package main

import (
	"fmt"
	"testing"

	"github.com/ichunt2019/cfg/lib"
)

func TestConfig(t *testing.T) {


	lib.Init("./config/dev/")


	fmt.Println(lib.Instance("proxy").GetString("base.time_location"))
	fmt.Println(lib.Instance("proxy").GetString("http.addr"))
	fmt.Println(lib.Instance("config").GetString("servers.beta.ip"))
	fmt.Println(lib.Instance("config").GetString("database.default.host"))

	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(a)

	/*
		[clients]
		data = [
			["gamma", "delta"],
			[1, 2]
		]
	*/
	//a := lib.Instance("config").GetLixinString("clients.data")
	//fmt.Println(gjson.Parse(a).Array()[0])  //[gamma delta]
	//fmt.Println(gjson.Parse(a).Array()[1].Array()[0])  //1


	/*
		[supplier_no_brand]
		    3 = [615,757,46596,43172,52,46481,47811,48817]
		    7 =   [47778]
		    9 =   [47778,4589,12369]
	*/
		b := lib.Instance("config").GetStringMapStringSlice("supplier_no_brand")
		fmt.Println(b["3"])  //[615 757 46596 43172 52 46481 47811 48817]
		fmt.Println(b["7"])  //[47778]
}