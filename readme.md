## download

```sh
$ go get -v github.com/syui/xq
```

## use

```sh
$ xq i /path/to/rss.xml
{
  "title": "Creating a new theme",
  "link": "https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/",
  "date": "2018-01-16T00:00:00.000+09:00"
}
```

## build

```sh
$ git clone https://github.com/syui/xq
$ cd xq
$ go build -o xq
$ ./xq
```

## example

```sh
$ curl -sLO https://syui.cf/hugo-theme-air/index.xml

$ xq i ./index.xml
{"title":"Creating a new theme","link":"https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/","date":"2018-01-16T00:00:00.000+09:00"}
{"title":"Archive","link":"https://syui.gitlab.io/test-hugo-theme-wave/archive/","date":"2018-01-16T00:00:00.000+09:00"}
```


