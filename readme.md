`xq` converts xml to json.

## install

```sh
$ go get -v github.com/syui/xq
```

archlinux aur package

```sh
$ sudo pacman -S go-pie
$ yay -S xq
```

if you use go instead of go-pie

`PKGBUILD`

```sh
- makedepends=('go-pie')
+ makedepends=('go')
```

## download

https://github.com/syui/xq/releases

## use

```sh
$ xq /path/to/rss.xml
[{"title":"Creating a new theme","link":"https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/","date":"2018-01-16T00:00:00.000+09:00"}, {"title":"Archive","link":"https://syui.gitlab.io/test-hugo-theme-wave/archive/","date":"2018-01-16T00:00:00.000+09:00"}]

$ xq /path/to/rss.xml|jq .
[
  {
    "title": "Creating a new theme",
    "link": "https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/",
    "date": "2018-01-16T00:00:00.000+09:00"
  }
]
```

## build

```sh
$ git clone https://github.com/syui/xq
$ cd xq
$ go build -o xq
$ ./xq h
```

## example

```sh
$ curl -sLO https://syui.cf/hugo-theme-air/index.xml

$ xq ./index.xml
[{"title":"Creating a new theme","link":"https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/","date":"2018-01-16T00:00:00.000+09:00"}, {"title":"Archive","link":"https://syui.gitlab.io/test-hugo-theme-wave/archive/","date":"2018-01-16T00:00:00.000+09:00"}]

$ xq i index.xml
[{
  "title": "Creating a new theme",
  "link": "https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/",
  "date": "2018-01-16T00:00:00.000+09:00"
}]

# latest update
$ xq u index.xml
2018-01-23T00:00:00.000+09:00

# latest post
$ xq p index.xml
2018-01-16T00:00:00.000+09:00

# latest item link
$ xq latest link index.xml
https://syui.gitlab.io/test-hugo-theme-wave/2016/01/01/creating-a-new-theme/

# latest item title
$ xq l t index.xml
Creating a new theme

# latest item publish
$ xq l p index.xml
2018-01-16T00:00:00.000+09:00

# latest item description
$ xq l d index.xml

# latest item author
$ xq l a index.xml
```

## update

- 0.2.0 : urfave/cli/v2

- 0.2.1 : option/[a]ll

- 0.2.2 : option/[p]ublish

- 0.2.3 : fix option/[i]tem

- 0.2.4 : option/none, $ xq file.xml

- 0.3.0 : option/[l]atest, sub-command/{link, title, published, description}, $ xq l link ./index.xml

- 0.3.1 : fix version, $ xq -v

- 0.3.2 : add alias sub-command, $ xq l d ./index.xml

- 0.3.3 : change option/all, $ xq a ./index.xml

- 0.3.4 : change option/update -> latest, $ xq latest ./index.xml

- 0.3.5 : change option/update -> latest, $ xq latest author ./index.xml
