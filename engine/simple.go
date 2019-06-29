package engine

import (
	"log"
	"newproject/crawler/fetcher"
)

type SimpleEngine struct {

}


//传入第一层Request，称之为seeds
//变长参数列表，实参需要是[]Request
func (e SimpleEngine)Run(seeds ...Request){

	//创建一个[]Request，将seeds放进去
	var requests []Request
	for _,r := range seeds{
		requests=append(requests,r)
	}

	for len(requests)>0{
		//slice没有pop方法，自己实现很简单
		r:=requests[0]
		requests=requests[1:]

		parseResult,err:=worker(r)
		if err!=nil{
			continue
		}

		//将下一层[]Request放入requests
		requests=append(requests,
			parseResult.Requests...)

		//将本层item打印到日志
		for _,item:=range parseResult.Items{
			log.Printf("Got item %v",
				item)
		}
	}
}

//获取Request，返回ParseResult
func worker(r Request) (ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err:=fetcher.Fetch(r.Url)   //获取url返回的body
	if err != nil {
		log.Printf("Fetcher:error " +
			"fetching url %s: %v",
			r.Url,err)
		//注意返回值的类型，这里返回对应的空值，ParseResult空值为ParseResult{}
		return ParseResult{},err
	}
	//err的空值为nil
	return r.ParserFunc(body),nil   //解析body
}


