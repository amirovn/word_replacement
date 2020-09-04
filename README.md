# word replacement

replacement of words in files, in the names of directories and files, preserving the style (camel)

замена слов в файлах, в названиях директорий и файлах, с сохранением стиля (camel)



Example
 - test/inIndex.go
 - test/InIndex.go
 - test/in_index.go
main.go

go run main.go inIndex toIndex test/

  - test/toIndex.go
  - test/ToIndex.go
  - test/to_index.go
