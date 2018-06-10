# NIFCLOUD SDK for Go

nifcloud-sdk-go is the **UNOFFICIAL** NIFCLOUD SDK for the Go programming language.

This repository was forked from [aws-sdk-go](https://github.com/aws/aws-sdk-go) and modified the following items.

* Fix package name and import path.
* Remove AWS services.
* Add JSON models for NIFCLOUD and generate services.
* Add signature version `v2computing` for NIFCLOUD Computing.
* Add protocol `computing` and `script`.
* Fix datetime format.

## Features

* :heavy_check_mark: Full support for NIFCLOUD Computing / RDB / NAS / Script APIs (as of Nov 13, 2017)
* :heavy_check_mark: AWS-SDK-compatible data-driven architecture

## Installing

```
$ go get -u github.com/alice02/nifcloud-sdk-go
```

or if you use dep, within your repo run:

```
$ dep ensure -add github.com/alice02/nifcloud-sdk-go
```

## Usage

Write the following code.

```go
package main

import (
        "fmt"

        "github.com/alice02/nifcloud-sdk-go/nifcloud"
        "github.com/alice02/nifcloud-sdk-go/nifcloud/credentials"
        "github.com/alice02/nifcloud-sdk-go/nifcloud/session"
        "github.com/alice02/nifcloud-sdk-go/service/computing"
)

func main() {
        sess := session.Must(session.NewSession(&nifcloud.Config{
                Region: nifcloud.String("jp-east-1"),
                Credentials: credentials.NewEnvCredentials(),
        }))

        client := computing.New(sess)

        fmt.Println(client.DescribeInstances(nil))
}
```

Set environment variables and execute the program.

```
$ export AWS_ACCESS_KEY_ID=<Your NIFCLOUD Access Key ID>
$ export AWS_SECRET_ACCESS_KEY=<Your NIFCLOUD Secret Access Key>
$ go run main.go
```

## License

This SDK is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.
