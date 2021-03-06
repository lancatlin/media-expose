package main

import "strings"

func ExampleImportCSV() {
	openTestDB()
	companies := `
"id","name","domain","country","owner","shareholders","invested_by_china","source","note"
1,"中國時報文化事業股份有限公司","https://www.want-want.com","TW","李玉生","蔡合旺事業(47,583)、旺嘉實業(46,172)","true","https://fubon-ebrokerdj.fbs.com.tw/z/zc/zck/zck_2816.djhtm、https://zh.wikipedia.org/wiki/蔡衍明、https://zh.wikipedia.org/wiki/旺旺集團","蔡衍明為旺旺集團創辦人最大股東主張台灣人應親中，不反共，也有因為中共資金被罰款"
2,"壹傳媒有限公司","https://www.nextdigital.com.tw/","HK","黎智英","黎智英(73.49%)、大衛.韋伯(5.02%)、張嘉聲(1.73%)","false","https://zh.wikipedia.org/wiki/%E5%A3%B9%E5%82%B3%E5%AA%92","壹傳媒一直被部分人士砲轟其報導手法譁眾取寵、渲染色情，嚴重破壞傳媒生態"
3,"聯合線上股份有限公司","http://co.udn.com/","TW","劉永平",,"false","http://co.udn.com/",
4,"聯邦企業集團","https://zh.wikipedia.org/wiki/聯邦集團","TW",,,"false",,
`
	media := `
"_id","name","domain","country","company_id","source","note"
1,"中時新聞","www.chinatimes.com","TW",1,"www.chinatimes.com","中時新聞內容包含中國時報、工商時報、旺報、時報周刊、中視"
2,"蘋果即時","tw.appledaily.com","HK",2,"https://tw.appledaily.com/",
3,"聯合線上","co.udn.com","TW",3,"http://co.udn.com/",
4,"自由時報","www.ltn.com.tw","TW",4,"https://www.ltn.com.tw/",
`
	ImportCSV(strings.NewReader(companies), ImportCompanies)
	ImportCSV(strings.NewReader(media), ImportMedia)
	// Output:
	// success import 中國時報文化事業股份有限公司
	// success import 壹傳媒有限公司
	// success import 聯合線上股份有限公司
	// success import 聯邦企業集團
	// success import 中時新聞
	// success import 蘋果即時
	// success import 聯合線上
	// success import 自由時報
}
