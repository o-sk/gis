# gis
`gis` is a command to search for image by Google Image Search

## Installation
Use `go get` to install this package:

```bash
$ go get -u github.com/o-sk/gis/cmd/gis
```

## Usage

### Basic
```bash
$ gis help
NAME:
   gis - Google Image Search

USAGE:
   gis [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     download, d  download image file
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --query value, -q value  Query (default: "apple")
   --help, -h               show help
   --version, -v            print the version
```

### Example

```bash
$ gis -q "猫"
[]gis.Image{
  gis.Image{
    ID:          "JItpTG6FJrAE8M:",
    Cite:        "ja.wikipedia.org",
    Description: "ネコ - Wikipedia",
    Destination: "https://ja.wikipedia.org/wiki/%E3%83%8D%E3%82%B3",
    Extension:   "jpg",
    Source:      "https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/2016-06-14_Orange_and_white_tabby_cat_born_in_2016_%E8%8C%B6%E3%83%88%E3%83%A9%E7%99%BD%E3%81%AD%E3%81%93_DSCF6526%E2%98%86%E5%BD%A1.jpg/200px-2016-06-14_Orange_and_white_tabby_cat_born_in_2016_%E8%8C%B6%E3%83%88%E3%83%A9%E7%99%BD%E3%81%AD%E3%81%93_DSCF6526%E2%98%86%E5%BD%A1.jpg",
    Thumbnail:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSSpBqahonmleOI6IAqQ0ihPSeSz9Dw6lzHdzYS0SRnt3wDkS4uzA",
  },
  gis.Image{
    ID:          "vFiG36xh6DR_0M:",
    Cite:        "psnews.jp",
    Description: "猫は自分の名前を覚える？名前に反応してくれる猫にする３つのコツ！｜猫 ...",
    Destination: "http://psnews.jp/cat/p/24710/",
    Extension:   "jpg",
    Source:      "http://psnews.jp/cat/uploads/2017/02/cat-1646566_1280.jpg",
    Thumbnail:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT7QC31AgY22AtTnX6uN7QTeqRdGwqG3DzXp-a12cTdwxs2RwmsUw",
  },
  gis.Image{
    ID:          "Bj425h-z7-rBgM:",
    Cite:        "grapee.jp",
    Description: "猫を『買う』以外の選択肢 意外と知られていない？ – grape [グレイプ]",
    Destination: "https://grapee.jp/329450",
    Extension:   "jpg",
    Source:      "https://grapee.jp/wp-content/uploads/32187_main2.jpg",
    Thumbnail:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS_b7k_BbYGphNHjs7pGxDjbdCE05wDciCLfl0LCM22jGJkYznOLA",
  },
  gis.Image{
    ID:          "E_zRZ6xSuKjXYM:",
    Cite:        "rocketnews24.com",
    Description: "すげえ】猫を飼っている人の99.7％が「幸せになった」という調査結果が ...",
    Destination: "https://rocketnews24.com/2018/04/24/1051250/",
    Extension:   "jpg",
    Source:      "https://sociorocketnews.files.wordpress.com/2018/04/mokichi004.jpg?w=640&h=480",
    Thumbnail:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRhbQGIr19I8UW6pRUSu5JMSV1YqXEYius6NWRhNLZ8rZCfLbAV",
  },
  gis.Image{
    ID:          "d2BtkouqPQCtWM:",
    Cite:        "diamond.jp",
    Description: "猫は自分の賢さを自覚している？ | 猫はためらわずにノンと言う ...",
    Destination: "https://diamond.jp/articles/-/151866",
    Extension:   "jpg",
    Source:      "https://dol.ismcdn.jp/mwimgs/6/b/670m/img_6b201e1a7a266ad253bf473f846d695e67132.jpg",
    Thumbnail:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR1Jxunfgpgv6oCyk69BKPCyC7dEz1svxGzKYY16NvgLNdC4ilh",
  },

  ...
}
```

### Download command

``` bash
$ gis download -h  
NAME:
   gis download - download image file

USAGE:
   gis download [command options] [arguments...]

OPTIONS:
   --directory value, -d value  download directory
   --filename value, -f value   download filename (ex: <filename>51.jpg) (default: "image")
```

### Example
```bash
$ gis -q "猫" download -d tmp -f cat
```
#### output
![image](https://user-images.githubusercontent.com/44518256/50286810-eafadc00-04a3-11e9-8ae3-39a9d3f755db.png)
