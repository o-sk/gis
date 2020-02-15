# gis
[![CircleCI](https://circleci.com/gh/o-sk/gis/tree/master.svg?style=svg)](https://circleci.com/gh/o-sk/gis/tree/master)

~~`gis` is a module to search for image by scraping Google Image Search~~

`gis` is not work correctly

## Usage
```
import "github.com/o-sk/gis"

images, _ := gis.Search("猫")
pp.Print(images)

// Output:
// []gis.Image{
//  gis.Image{
//    ID:          "JItpTG6FJrAE8M:",
//    Cite:        "ja.wikipedia.org",
//    Description: "ネコ - Wikipedia",
//    Destination: "https://ja.wikipedia.org/wiki/%E3%83%8D%E3%82%B3",
//    Extension:   "jpg",
//    Source:      "https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/2016-06-14_Orange_and_white_tabby_cat_born_in_2016_%E8%8C%B6%E3%83%88%E3%83%A9%E7%99%BD%E3%81%AD%E3%81%93_DSCF6526%E2%98%86%E5%BD%A1.jpg/200px-2016-06-14_Orange_and_white_tabby_cat_born_in_2016_%E8%8C%B6%E3%83%88%E3%83%A9%E7%99%BD%E3%81%AD%E3%81%93_DSCF6526%E2%98%86%E5%BD%A1.jpg",
//    Thumbnail:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSSpBqahonmleOI6IAqQ0ihPSeSz9Dw6lzHdzYS0SRnt3wDkS4uzA",
//  },
//  gis.Image{
//    ID:          "vFiG36xh6DR_0M:",
//    Cite:        "psnews.jp",
//    Description: "猫は自分の名前を覚える？名前に反応してくれる猫にする３つのコツ！｜猫 ...",
//    Destination: "http://psnews.jp/cat/p/24710/",
//    Extension:   "jpg",
//    Source:      "http://psnews.jp/cat/uploads/2017/02/cat-1646566_1280.jpg",
//    Thumbnail:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT7QC31AgY22AtTnX6uN7QTeqRdGwqG3DzXp-a12cTdwxs2RwmsUw",
//  },
//  gis.Image{
//    ID:          "Bj425h-z7-rBgM:",
//    Cite:        "grapee.jp",
//    Description: "猫を『買う』以外の選択肢 意外と知られていない？ – grape [グレイプ]",
//    Destination: "https://grapee.jp/329450",
//    Extension:   "jpg",
//    Source:      "https://grapee.jp/wp-content/uploads/32187_main2.jpg",
//    Thumbnail:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS_b7k_BbYGphNHjs7pGxDjbdCE05wDciCLfl0LCM22jGJkYznOLA",
//  },
```