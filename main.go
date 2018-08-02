package main

import (
        "fmt"
        "os"
        "strconv"
        "flag"
        "sync"
)

var folder string
var chapter int
const MAXIMAGES int = 20

func init()  {
  fmt.Println("init")
  flag.IntVar(&chapter, "chap", 0, "Set chapter to download using -chap=123")
  flag.StringVar(&folder, "output", "./images", "Define output folder like -output='/Users/me/images'")
  flag.Parse()
}

func main() {
  fmt.Println("Downloading file...")

  args := os.Args

  if len(args) < 2 {
    fmt.Println("Argument missing")
    panic("argument missing")
  }

  var wg sync.WaitGroup
  wg.Add(MAXIMAGES)



  fmt.Println("Chapter:", chapter)

  // Create output folder
  os.Mkdir(folder, 0777)

  baseURL := "http://85.14.254.67/manga/kapitel/"

  // Download in a loop
  for i := 1; i <= MAXIMAGES; i++ {

    go func (i int) {
      defer wg.Done()
      var index string
      if i < 10 {
        index = "0" + strconv.Itoa(i)
      } else {
        index = strconv.Itoa(i)
      }

      Download(baseURL + strconv.Itoa(chapter) + "/" + index + ".jpg")
    } (i)

  }

  wg.Wait()


}
