package main
import (
    "fmt"
    "net/http"
	"os"
    "github.com/PuerkitoBio/goquery"
    "log"	
	"strings"
)

func check(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func main() {
	str:="https://www.rusprofile.ru/search?query="+os.Args[1]+"&type=ul"
	fmt.Println(str)
    resp, err := http.Get(str) 
    if err != nil { 
        fmt.Println(err) 
        return
    } 
    defer resp.Body.Close()
/*
    for true {
             
        bs := make([]byte, 1014)
        n, err := resp.Body.Read(bs)
        fmt.Println(string(bs[:n]))
         
        if n == 0 || err != nil{
            break
        }
    }
*/
	doc, err := goquery.NewDocumentFromReader(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    title := doc.Find("title").Text()
    fmt.Println(title)

	name:=doc.Find("div.company-name")

	fmt.Println(name.Size())
	fmt.Println(name.Text())	



    doc.Find("div.company-requisites").Find("div.company-row").Each(func(index int,item *goquery.Selection){


			item.Find("dl.company-col").Each(func(index int,item *goquery.Selection){

				info:=	item.Find("dt.company-info__title")			
//			fmt.Println(info.Text())	

			if info.Text()=="ИНН/КПП"{
				item.Find("dd.company-info__text").Each(func(index int,item *goquery.Selection){
					fmt.Println(index,": ",strings.TrimSpace(item.Text()))
				})
			}

	
		})

		leftcol:=doc.Find("div.leftcol")
		leftcol.Find("div.company-row").Each(func(index int,item *goquery.Selection){

			str:=item.Find("span.company-info__title").Text()
			if str=="Руководитель"{

				res:=item.Find("span.company-info__text").Find("a.link-arrow").Find("span").Text()
				fmt.Println(res)

			}


		})




	})




	check(err)

}