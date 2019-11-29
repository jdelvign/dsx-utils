[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

![Version](https://img.shields.io/github/manifest-json/v/jdelvign/dsxutl)

# dsxutl
Simple tool to work with IBM DataStage DSX files

## `dsxutil grep` command
- dsxutl grep search a substring into a dsx file.
- When found, display the job name the substring where found and the line number into the DSX file.

```
Usage: dsxutl grep -substr <SUBSTRING> [-ignoreCase] -dsxfile <DSXFILE>
  -dsxfile string
        The DSX file to search in
  -ignoreCase
        Search the substring in case sensitive (false/default) or not (true)
  -substr string
        The substring to find in the DSX file
```

## `dsxutl ljobs` command
- Display the jobdesigns contained in the DSX file on the standard output

## `dsxutl header` command
- Display the DSX header on the standard output

# How to Build
go build

