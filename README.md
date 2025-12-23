#README
## Overview 
go pubsub test case

## usage 

go test 
```
go mod tidy
go test -count=1 -run Integration -v
```

check by pulling messages 

```
gcloud pubsub subscriptions pull test-pull-sub2 --auto-ack --limit=10
```

