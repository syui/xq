## build

```sh
$ git clone https://gitlab.com/syui/xq
$ cd xq
$ go build
$ ./xq -h
```

## example

```sh
$ url=https://syui.gitlab.io/test-hugo-theme-wave/index.xml
$ curl -sSL -H "Accept: application/xml" $url -o index.xml
$ xq t ./index.xml
```

## ref

https://github.com/urfave/cli

https://github.com/mmcdole/gofeed

http://yatta47.hateblo.jp/entry/2017/05/21/233522
