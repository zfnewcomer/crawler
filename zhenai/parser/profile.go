package parser

import (
	"newproject/crawler/engine"
	"newproject/crawler/model"
	"regexp"
	"strconv"
)

//const ageRe  = `"basicInfo":\[[^]]+,"([\d]+)岁"[^[]+,\]`
//const marriageRe  = `"basicInfo":\[([^,]+).*\]`

var(
	//url1=`<div class="des f-cl" data-v-5b109fc3="">阿坝 | 26岁 | 大学本科 | 未婚 | 156cm | 5001-8000元<a class="online f-fr" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=http%3A%2F%2Falbum.zhenai.com%2Fu%2F1662184411" target="_self" data-v-5b109fc3="">查看最后登录时间</a></div>`
	ageRe=regexp.MustCompile(`<div[^>]+>([\d]+)岁</div>`)
	marriageRe=regexp.MustCompile(`<div[^>]+>[^ ]+ | [^ ]+ | [^ ]+ | ([^ ]+) |[^<]+</div>`)
	heightRe=regexp.MustCompile(`<div[^>]+>([\d]+)cm</div>`)
	weightRe=regexp.MustCompile(`<div[^>]+>([\d]+)kg</div>`)
	incomeRe=regexp.MustCompile(`<div[^>]+>(月收入:[^<]+)</div>`)
	occupationRe=regexp.MustCompile(`<div[^>]+>(工作地:[^<]+)</div>`)
	educationRe=regexp.MustCompile(`<div[^>]+>[^ ]+ | [^ ]+ | ([^ ]+) | [^ ]+ |[^<]+</div> `)
	hukouRe=regexp.MustCompile(`<div[^>]+>([^ ])+ | [\d]+岁 | [^ ]+ | [^ ]+ | [\d]+cm | [^ ]+</div>`)
	//houseRe=regexp.MustCompile()
	//carRe=regexp.MustCompile()
)

func ParseProfile(
	contents []byte,name string) engine.ParseResult{
		profile:=model.Profile{}

		profile.Name=name
		//返回值是否有err，点开strconv.Atoi的定义可知
		//年龄
		age,err:=strconv.Atoi(extractString(contents,ageRe))
		if err ==nil{
			profile.Age=age
		}
		//身高
		height,err:=strconv.Atoi(extractString(contents,heightRe))
		if err ==nil{
			profile.Height=height
		}
		//重量
		weight,err:=strconv.Atoi(extractString(contents,weightRe))
		if err ==nil{
			profile.Weight=weight
		}
		profile.Marriage=extractString(contents,marriageRe)
		profile.Income=extractString(contents,incomeRe)
		profile.Education=extractString(contents,educationRe)
		profile.Occupation=extractString(contents,occupationRe)
		profile.Hukou =extractString(contents,hukouRe)
		//profile.House=extractString(contents,houseRe)
		//profile.Car=extractString(contents,carRe)

		result:= engine.ParseResult{
			Items:[]interface{}{profile},
		}
		return result
}

func extractString(
	contents []byte,re *regexp.Regexp) string{
		match:=re.FindSubmatch(contents)
		if len(match)>=2{
			return string(match[1])
		}else{
			return ""
		}
}


