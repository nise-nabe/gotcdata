gotcdata
========

Topcoder Data Feeds for go [link](http://apps.topcoder.com/wiki/display/tc/Algorithm+Data+Feeds)

|DataFeeds                       | Implements |
|--------------------------------|------------|
|Coder List                      |        100%|
|Active Algorithm Coder List     |          0%|
|Algorithm Round List            |          0%|
|Algorithm Round Results         |          0%|
|Algorithm Rating  History       |          0%|
|Algorithm Practice Rooms        |          0%|
|Algorithm Practice Room Details |          0%|

## Usage

```
package main

import (
  "fmt"
  "github.com/nise-nabe/gotcdata"
  "log"
  "os"
)

func main() {
  coderList, err := gotcdata.NewCoderList()
  if err != nil {
    log.Fatalln(err)
    os.Exit(1)
  }

  for _, coder := range coderList {
    fmt.Println(coder.Handle, ":", coder.AlgRating)
  }
}
```

output
```
...
nise_nabe : 964
...
```

## Problem

It's too slow to get the feeds because too large.
