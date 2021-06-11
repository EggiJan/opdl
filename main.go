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

  var chap string
  if chapter < 1000 {
    chap = "0" + strconv.Itoa(chapter)
  } else {
    chap = strconv.Itoa(chapter)
  }

  args := os.Args

  if len(args) < 2 {
    fmt.Println("Argument missing")
    panic("argument missing")
  }

  var wg sync.WaitGroup
  wg.Add(MAXIMAGES)

  fileTypes := []string{"jpg", "png", "tiff"}

  fmt.Println("Chapter:", chap)

  // Create output folder
  os.Mkdir(folder, 0777)

  baseURL := "https://manga-lesen.com/kapitel/"

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

      // Download all filetypes
      for j := 0; j < len(fileTypes); j++ {
        Download(baseURL + chap + "/" + index + "." + fileTypes[j])
      }
      
    } (i)

  }

  wg.Wait()


}
