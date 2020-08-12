package main

import (
	"fmt"
	"html/template"
	"net/http"
	"math/rand"
	"time"
)
type selectedImg struct{
	Pic1,Pic2,Pic3,Pic4,Pic5,Pic6 string
}
var randmPic selectedImg


func main (){

http.HandleFunc("/",indexHandler)
http.ListenAndServe(":8086",nil)

}
func indexHandler( w http.ResponseWriter, r *http.Request){
	randmPic.Pic1,randmPic.Pic2,randmPic.Pic3,randmPic.Pic4,randmPic.Pic5,randmPic.Pic6 = picSelecter()
	t,error := template.ParseFiles("./template/index.html")
	if error != nil {
		fmt.Println(error)
		fmt.Println("ERROR")
	}else {
		if err := t.Execute(w,randmPic); err != nil {
			fmt.Println(err)
		}

	}
}


func picSelecter()(link1,link2,link3,link4,link5,link6 string){
imgSrc := []string{
		"https://www.rootsofaction.com/wp-content/uploads/2018/03/The-mind-that-opens-to-a-new-idea.png",
		"https://www.marineinsight.com/wp-content/uploads/2018/04/2-1.png",
		"https://www.success.com/wp-content/uploads/2019/08/1.png",
		"https://cdn.lifehack.org/wp-content/uploads/2016/12/08230522/best-motivational-quotes-dont-limit-challenges.png",
		"https://image.marriage.com/quote-images/3311037300.jpg",
		"https://www.rockettes.com/wp-content/uploads/2015/06/11741022_10153239325453551_8060414301275087317_o.jpg",
		"https://edlester.com/wp-content/uploads/2017/01/14980674_1190862611007414_5753585401113410984_n-300x300.jpg",
		"https://img.theculturetrip.com/450x/wp-content/uploads/2017/01/inspirational-quotes-squre-3.png",
		"https://www.ryrob.com/wp-content/uploads/2017/09/hustle_quotes_motivation_I%E2%80%99ve-got-a-dream-that%E2%80%99s-worth-more-than-my-sleep.jpg",
		"https://spotio.com/wp-content/uploads/2017/02/Motivational-Quote-1-300x300.jpg",
		"https://i.pinimg.com/736x/b9/f4/8f/b9f48fc73c7e15abd6598d0c72bda688.jpg",
		"https://spotio.com/wp-content/uploads/2017/02/Motivational-Quote-2.png",
		"https://i.pinimg.com/236x/09/e9/60/09e9602777186e0665edb75366e92aa3--quotes-motivation-motivation-inspiration.jpg",
		"https://theincrementalmama.com/wp-content/uploads/2019/02/5.jpg",
		"https://a2n9b7p7.stackpathcdn.com/wp-content/uploads/happiness_quotes_3.jpg",
		"https://media1.popsugar-assets.com/files/thumbor/lBsNqFIoilJknLXHD0ruYDZJ21w/fit-in/1024x1024/filters:format_auto-!!-:strip_icc-!!-/2018/03/07/083/n/44344577/addurlRs68hY/i/Quotes-About-Sweat.jpg",
		"https://cdn.lifehack.org/wp-content/uploads/2016/12/08232038/best-motivational-quotes-choice.jpg",
		"https://i.pinimg.com/originals/c4/96/d0/c496d0d77e56409c8afa379aa2f5d1a6.jpg",
		"https://i.pinimg.com/originals/cb/c8/4f/cbc84fd11811a4efcd1a1a1449c6f35f.jpg",
		"https://brightdrops.com/wp-content/uploads/2013/10/drseussyouhavebrainsinyourhead.jpg",
		"https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fstatic.onecms.io%2Fwp-content%2Fuploads%2Fsites%2F35%2F2020%2F02%2FSHAPE-motivational-quotes-16-2000.jpg",
		"https://www.success.com/wp-content/uploads/legacy/sites/default/files/14_14.jpg",
		"https://i.pinimg.com/originals/8f/8b/1e/8f8b1e58e335e1adfd65939905e61ade.jpg",
		"https://i.pinimg.com/originals/c3/96/ac/c396ac310b308149385c701f2e7dbdd5.jpg",
		"https://www.quotemaster.org/images/59/599e2c1218aa688a79ed2cb3820e75d2.jpg",
		"https://i.pinimg.com/originals/69/77/d3/6977d38c7c72ffe1602b35174930e722.jpg",
		"https://i.pinimg.com/originals/df/fb/38/dffb38279db58a5ccb79ee9cc50352ca.jpg",
		"https://quotess.net/wp-content/uploads/2017/10/work-quotes-trust-sport-quote-more.jpg",
		"https://i.redd.it/bim1kjhmojh31.jpg",
		"https://cdn.everydaypower.com/wp-content/uploads/2017/03/Inspirational-quote-32.jpg",
		"https://cdn.everydaypower.com/wp-content/uploads/2017/03/Inspirational-quote-7.jpg",
		"https://englishblogmmg.edublogs.org/files/2018/12/dont-judge-me-for-my-past-2eokv47-2k2v9pp.jpg",
		"https://shaneus814.files.wordpress.com/2018/01/its-going-to-take-some-time-motivational-gym-quotes.jpg",
		"https://i.pinimg.com/originals/7d/e1/42/7de142a1595455485fc197d2c54e6b5d.jpg",
		"https://previews.123rf.com/images/kotokomi/kotokomi1903/kotokomi190300002/119628224-trust-yourself-inspirational-quote-modern-calligraphy-motivational-saying.jpg",
	}
	rand.Seed(time.Now().UnixNano())
	
	link1 = imgSrc[rand.Intn(len(imgSrc))]
	link2 = imgSrc[rand.Intn(len(imgSrc))]
	link3 = imgSrc[rand.Intn(len(imgSrc))]
	link4 = imgSrc[rand.Intn(len(imgSrc))]
	link5 = imgSrc[rand.Intn(len(imgSrc))]
	link6 = imgSrc[rand.Intn(len(imgSrc))]
	return
}