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
$ gis -q "猫" | jq
[
  {
    "Source": "https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/2016-06-14_Orange_and_white_tabby_cat_born_in_2016_%E8%8C%B6%E3%83%88%E3%83%A9%E7%99%BD%E3%81%AD%E3%81%93_DSCF6526%E2%98%86%E5%BD%A1.jpg/200px-2016-06-14_Orange_and_white_tabby_cat_born_in_2016_%E8%8C%B6%E3%83%88%E3%83%A9%E7%99%BD%E3%81%AD%E3%81%93_DSCF6526%E2%98%86%E5%BD%A1.jpg",
    "Description": "ネコ - Wikipedia",
    "Cite": "ja.wikipedia.org"
  },
  {
    "Source": "http://psnews.jp/cat/uploads/2017/02/cat-1646566_1280.jpg",
    "Description": "猫は自分の名前を覚える？名前に反応してくれる猫にする３つのコツ！｜猫 ...",
    "Cite": "psnews.jp"
  },
  {
    "Source": "https://grapee.jp/wp-content/uploads/32187_main2.jpg",
    "Description": "猫を『買う』以外の選択肢 意外と知られていない？ – grape [グレイプ]",
    "Cite": "grapee.jp"
  },
  {
    "Source": "https://upload.wikimedia.org/wikipedia/commons/thumb/3/33/Hannibal_Poenaru_-_Nasty_cat_%21_%28by-sa%29.jpg/270px-Hannibal_Poenaru_-_Nasty_cat_%21_%28by-sa%29.jpg",
    "Description": "ネコ - Wikipedia",
    "Cite": "ja.wikipedia.org"
  },
  {
    "Source": "http://image.itmedia.co.jp/nl/articles/1811/17/kh_1811tio01.jpg",
    "Description": "似てるかニャ？ 『おじさまと猫』のふくまるとそっくりな猫ちゃんが ...",
    "Cite": "nlab.itmedia.co.jp"
  },
  {
    "Source": "https://dol.ismcdn.jp/mwimgs/8/d/670m/img_8db0612c13c0013326bfb1b66431df95645897.jpg",
    "Description": "人が「猫の時間」を生きると、何が起こる？全仏ベストセラーの猫本が ...",
    "Cite": "diamond.jp"
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