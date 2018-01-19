## download

```sh
$ go get -u gitlab.com/syui/xq
```

## build

```sh
$ git clone https://gitlab.com/syui/xq
$ cd xq
$ go build -o xq
$ ./xq -h
```

## example

```sh
$ url=https://syui.gitlab.io/test-hugo-theme-wave/index.xml
$ curl -sSL -H "Accept: application/xml" $url -o index.xml
$ cat index.xml

$ xq i ./index.xml
{"title":"Creating a new theme","link":"https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/","date":"2018-01-16T00:00:00.000+09:00"}
{"title":"Archive","link":"https://syui.gitlab.io/test-hugo-theme-wave/archive/","date":"2018-01-16T00:00:00.000+09:00"}
```

```sh
$ date -R
$ date -d "Sat, 20 Jan 2018 00:00:00 +0900" '+%Y-%m-%dT%H:%H:%M'
2018-01-20T00:00:00
```

## version 

### 0.0.3 : latest updated

```sh
$ xq l ./index.xml
2018-01-16T00:00:00.000+09:00
```

## ref

https://github.com/urfave/cli

https://github.com/mmcdole/gofeed

http://yatta47.hateblo.jp/entry/2017/05/21/233522
