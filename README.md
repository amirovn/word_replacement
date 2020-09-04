# word replacement


replacement of words in files, in the names of directories and files

замена слов в файлах, в названиях директорий и файлах



Example

test/
- inIndex.go
- InIndex.go
- in_index.go
main.go

go run main.go inIndex toIndex test/

test/
- toIndex.go
- ToIndex.go
- to_index.go
