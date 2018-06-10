## download

```sh
$ go get -v gitlab.com/syui/xq
```

## use

```sh
$ xq /path/to/rss.xml
{
  "title": "Creating a new theme",
  "link": "https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/",
  "date": "2018-01-16T00:00:00.000+09:00"
}
```

## build

```sh
$ git clone https://gitlab.com/syui/xq
$ cd xq
$ go build -o xq
$ ./xq
```

## example

```sh
$ url=https://syui.gitlab.io/test-hugo-theme-wave/index.xml
$ curl -sSL -H "Accept: application/xml" $url -o index.xml
$ cat index.xml

$ xq ./index.xml
{"title":"Creating a new theme","link":"https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/","date":"2018-01-16T00:00:00.000+09:00"}
{"title":"Archive","link":"https://syui.gitlab.io/test-hugo-theme-wave/archive/","date":"2018-01-16T00:00:00.000+09:00"}
```

## ref

https://github.com/urfave/cli

https://github.com/mmcdole/gofeed

http://yatta47.hateblo.jp/entry/2017/05/21/233522
