package main 

import (

  "fmt"
  "io/ioutil"
  "net/http"
  //"reflect"
  //"strconv"
  "strings"

)

func FindStat(text string, d bool) (string, []string) {

  //fmt.Println("hi")
  //fmt.Println(text)

  dummy := strings.Split(text, "<p>")
  bt, at := dummy[0], dummy[1:]
  dummy = strings.Split(at[0],"</p>")[0:2]
  val, next := dummy[0], at[1]

  bt = strings.Split(bt, "</h4>")[0]

  i := 1

  var stat string;

  //fmt.Println(bt)

  for bt[len(bt) - i] != uint8(62) {
    stat = string(bt[len(bt) - i]) + stat
    i++
  }

  if (d) {

    //fmt.Println(next)

    i = 0

    var val2 string;

    for next[i] != uint8(60) {
      val2 = val2 + string(next[i])
      i++
  }

   

    return stat, []string{val,val2}
 
  }

  return stat, []string{val}
}

func StringMergeDiv(s []string) string {

  i := 0

  var r string;

  for i < len(s) {

    r = r + s[i] + "</div>"
    i++

  }

  return r

}

func PlayerData(url string) map[string][]string {

  resp, _ := http.Get(url)

  defer resp.Body.Close()
  body,_ := ioutil.ReadAll(resp.Body)

  text := string(body[:])

  text = strings.Split(text, "<div class=\"stats_pullout\">\n<div><div>\n<h4>SUMMARY</h4>")[1]

  m := make(map[string][]string)

  double := strings.Contains(strings.Split(text,"</div></div>")[0],"2017")

  text = strings.Split(text, "</div></div>")[1]

  for i := strings.Count(text, "</h4>"); i > 0; i-- {

    q1, q2 := FindStat(text, double)

    m[q1] = q2

    text = StringMergeDiv(strings.Split(text, "</div>")[1:])

    //fmt.Println(text)

  }

  return m  

}



func SortPlayers() {

  resp, _ := http.Get("http://www.baseball-reference.com/players/a/")

  defer resp.Body.Close()
  body,_ := ioutil.ReadAll(resp.Body)

  text := string(body[:])

  text = strings.Split(text, "<div class=\"section_content\" id=\"div_players_\">")[1]

  text = strings.Split(text, "</div>\n</div>")[0]

  players := strings.Split(text, "href=\"")[1:]

  for i := 0; i < len(players); i++ {

    d := strings.Split(players[i],"\">")
    url := "http://www.baseball-reference.com/" + d[0]
    name := strings.Split(d[1], "</a>")[0]

    //fmt.Println(url)

    fmt.Println(name)
    fmt.Println(PlayerData(url))

  }

  

  

}



func main() {

  SortPlayers()
  //PlayerData("http://www.baseball-reference.com/players/r/ruthba01.shtml")
  


  //fmt.Println(body)
  //fmt.Println(text)
  //fmt.Println(reflect.TypeOf(body))

}

