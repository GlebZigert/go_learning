package main

import (
	"fmt"
	"net/http"
    "github.com/PuerkitoBio/goquery"
	"os"
//	"strings"
	

)

func check(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func writeFile(data, filename string){
	file,err:=os.Create(filename)
	check(err)
	file.WriteString(data)
}

func main(){

url:="https://techcrunch.com/"

response,err:=http.Get(url)
check(err)

if response.StatusCode>400{
	fmt.Println("Status code:",response.StatusCode)
}

doc,err:=goquery.NewDocumentFromReader(response.Body)
check(err)
/*
river:=doc.Find("div.river").Find("div.post-block").Each(func(index int,item *goquery.Selection){
		h2:=item.Find("h2")
		title:=strings.TrimSpace(h2.Text())
		fmt.Println("title: ",title)	
	})
*/
river:=doc.Find("div.river").Find("div.post-block").Size()

fmt.Println(river)

//writeFile(river,"index.html")

}