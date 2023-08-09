# go-crowdin
Crowdin API in Go - https://crowdin.com/page/api

[![Go Report Card](https://goreportcard.com/badge/github.com/medisafe/go-crowdin)](https://goreportcard.com/report/github.com/medisafe/go-crowdin) [![Documentation](https://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/github.com/medisafe/go-crowdin)  

#### Install

`go get github.com/medisafe/go-crowdin`

- [Initialize](#initialize)
- [API](#api)
- [Debug](#debug)
- [App Engine](#app-engine)

##### Initialize

``` Go
api := crowdin.New("token", "project-name")
```

##### API

:blue_book: Check the doc - [Documentation](https://godoc.org/github.com/medisafe/go-crowdin)

> Examples:
> ``` Go
> // get language status
> files, err := api.GetLanguageStatus("ru")
> 
> // add file
> result, err := api.AddFile(&crowdin.AddFileOptions{
>     Type: "csv",
>     Scheme: "identifier,source_or_translation,context",
>     FirstLineContainsHeader: true,
>     Files: map[string]string{
>         "strings_profile_section.csv" : "local/path/to/strings_profile_section.csv",
>     },
> })
> ```

##### Debug

You can print the internal errors by enabling debug to true

``` Go
api.SetDebug(true, nil)
```

You can also define your own `io.Writer` in case you want to persist the logs somewhere.
For example keeping the errors on file

``` Go
logFile, err := os.Create("crowdin.log")
api.SetDebug(true, logFile)
```

##### App Engine

Initialize app engine client and continue as usual

``` Go
c := appengine.NewContext(r)
client := urlfetch.Client(c)

api := crowdin.New("token", "project-name")
api.SetClient(client)
```

[Documentation](https://godoc.org/github.com/medisafe/go-crowdin)

##### Author

Roman Kushnarenko - [sromku](https://github.com/sromku)

##### License 

Apache License 2.0