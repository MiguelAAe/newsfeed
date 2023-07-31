# News feed

News feed is a simple application that pulls RSS feed of news to be later consume as API

## Build the project

```shell
make build
```

## Run the server
```shell
./cmd/newsfeed
```

## Check docs in browser
```
http://localhost:9200/swagger/index.html
```

## Endpoint GET /newsFeed 
```shell
curl http://localhost:9200/newsFeed
```

### Query retrieval with Query Params 
```shell
curl http://localhost:9200/newsFeed?source=bbc&category=technology
```

## Endpoint POST /newsFeed

### Using a complex query as a body request to get a more custom feed list

```shell
curl http://localhost:9200/newsFeed
```
body:
```json
{
    "query":
    {
        "sky": ["entertainment"],
        "bbc": ["technology", "politics"]
    }
}
```

