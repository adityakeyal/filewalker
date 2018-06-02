# filewalker 
[![Go Report Card](https://goreportcard.com/badge/github.com/adityakeyal/filewalker)](https://goreportcard.com/report/github.com/adityakeyal/filewalker)
Fast recursive file walker library - ideal for multiple recursive folder structure


#import:
`go get "github.com/adityakeyal/filewalker"`

#Usage:
walk.Scan(startFolder string, callBack func(path string, filename string))
 - startFolder : The folder to recursive search
 - callback : Method called when a file is found
> Arguments
>            - path =  The path of the file relative to root
>            - filename = The filename with extension

