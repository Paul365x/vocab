package translate

import (
	"internal/config"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var url string = "https://glosbe.com"

type trTuple struct {
	Trans string
	Role string
	OrEx string
	TrEx string
}

func Translate(word string) []trTuple {

	// get dictionary page
	lang := config.GetLanguage()
	res,err := http.Get(url + lang + word)
	if err != nil {
		panic( err)
	}
	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic( err)
	}

	results := []trTuple{}

	pl := doc.Find("ul.pr-1").First()
	pl.Find("li").Each(func(i int, li *goquery.Selection) {
		res := trTuple{}
		la := li.Find("div div div.mt-1").First()
		tr := li.Find("div div div H3").First().Text()
		rl := li.Find("div div div span").First().Text()
		or := la.Find("p").First().Text()
		trex := la.Find("p").Next().Text()
		res.Trans = strings.TrimSpace(strings.Replace(tr,"\n"," ", -1))
		res.Role = strings.TrimSpace(strings.Replace(rl,"\n"," ", -1))
		res.OrEx = strings.TrimSpace(strings.Replace(or,"\n"," ", -1))
		res.TrEx = strings.TrimSpace(strings.Replace(trex,"\n"," ", -1))
		results = append(results,res)	
	})
	return results
}

/*
func main() {

results := translate("quod")
for t := range results {
	fmt.Printf("%s (%s) eg: %s (%s)\n",results[t].trans, results[t].role, results[t].OrEx, results[t].TrEx)	
}

	
}*/