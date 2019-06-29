package parser

import (
	"newproject/crawler/engine"
	"regexp"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(
	contents []byte) engine.ParseResult{
		re:=regexp.MustCompile(cityRe)
		matches:=re.FindAllSubmatch(contents,-1)
		result:=engine.ParseResult{}
		for _,m:=range matches{
			name:=string(m[2])
			result.Items=append(
				result.Items,"User "+name)
			result.Requests=append(
				result.Requests,engine.Request{
					Url:	string(m[1]),
					//需要将用户姓名传到ParseProfile，方案是给ParseProfile加一个string参数
					//直接修改ParserFunc的签名不行，ParseCityList和ParseCity不需要这个参数
					//办法是使用闭包，定义一个符合ParserFunc签名的匿名函数，返回ParseProfile

					//另外还有一点要注意，不能将string(m[2])直接传给ParseProfile
					//for-range中，m始终是用同一块内存存储的，值在不停的变化
					//而这里ParseProfile并不马上执行，等ParseProfile执行时m的值已经是最后一个元素了
					//这和在for-range循环中使用groutine是一样的问题
					ParserFunc: func(c []byte) engine.ParseResult {
						return ParseProfile(c,name)
					},
				},
				)
		}
		return result
}

