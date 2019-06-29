package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
//get an url
//return response.body
func Fetch1(url string) ([]byte,error){
	resp,err:=http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil,
		fmt.Errorf("wrong status code: %d",
			resp.StatusCode )
	}
	return  ioutil.ReadAll(resp.Body)
}

var rateLimiter=time.Tick(10*time.Microsecond)

func Fetch(url string) ([]byte,error){
	//每10ms打一个点，所有Fetch争抢，这样控制每秒只有100条request
	<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil,
			fmt.Errorf("wrong status code: %d",
				resp.StatusCode )
	}
	return  ioutil.ReadAll(resp.Body)
}