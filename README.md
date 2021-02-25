# cfg

配置文件采用github.com/spf13/viper



**配置文件使用**

1，普通配置文件


```
viewpath = "/home/www/templates/"

[owner]
name = "Tom Preston-Werner"
organization = "GitHub"
bio = "GitHub Cofounder & CEO\nLikes tater tots and beer."
dob = 1979-05-27T07:32:00Z # 日期时间是一等公民。为什么不呢？

```

获取方法

lib.Instance("proxy").GetString("base.time_location")


2，复杂的map类型配置文件


```
[servers]
  # 你可以依照你的意愿缩进。使用空格或Tab。TOML不会在意。
  [servers.alpha]
  ip = "10.0.0.1"
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"

```
获取方法

lib.Instance("config").GetString("servers.beta.ip")




3，二位数组类型


```
[clients]
data = [["gamma", "delta"],[1, 2]]

```

使用方法：采用gjson方式


```
//a := lib.Instance("config").GetLixinString("clients.data")

//fmt.Println(gjson.Parse(a).Array()[0])  //[gamma delta]

//fmt.Println(gjson.Parse(a).Array()[1].Array()[0])  //1


```

	

4，其它类型（map[string][]string）

```
[supplier_no_brand]
    3 =   [615,757,46596,43172,52,46481,47811,48817]
    7 =   [47778]
    9 =   [47778,4589,12369]

```



使用方法

```
b := lib.Instance("config").GetStringMapStringSlice("supplier_no_brand")
fmt.Println(b["3"])  //[615 757 46596 43172 52 46481 47811 48817]
fmt.Println(b["7"])  //[47778]

```




**初始化**

```
lib.Init("./config/dev/")
lib.Instance("config")//选择需要指定的配置文件

```










