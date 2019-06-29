package engine

//架构中是分层的：https://www.cnblogs.com/zhoufeng1989/p/10779120.html
//第一层是城市列表，
//第二层是城市
//第三层是用户

type Request struct{   //本层的request
	Url 		string
	ParserFunc func([]byte) ParseResult   //本层的解析器
}

type ParseResult struct{
	Requests 	[]Request        //下一层的request
	Items 		[]interface{}      //本层的数据
}

func NilParser([]byte) ParseResult{
	return 		ParseResult{}
}

