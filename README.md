# Article Service

## Installation Instructions for MacOS

1.  Run `go get github.com/arizard/nine-publishing-technical-test` at the terminal prompt.
2.  Install the `faker` package for some of the basic tests: `go get github.com/icrowley/fake`
3.  Navigate to the package directory: `cd $GOPATH/src/github.com/arizard/nine-publishing-technical-test/`
4.  Start the application: `go run main.go`
5.  Use your favorite request test tool (e.g. Postman or Paw) to test the endpoints.

## Article Object

The **Article** object models a news article. It has the following fields:
*   `id` `string` 
    *   A unique identitier for the Article. There can be no duplicate `id` values between Articles.
*   `title` `string` 
    *   The descriptive headline for the Article.
*   `date` `string` 
    *   The date the Article was written, formatted as `YYYY-MM-DD`.
*   `body` `string` 
    *   The body content of the Article.
*   `tags` `[]string` 
    *   An array of strings representing the tags of the article. Tags allow articles to be grouped by similar topics.

## Endpoints

All endpoints are of the content type `application/json` and return content of type `application/json`.

### POST `/articles`

Submit a new Article to the service. Returns information as to whether the submission was successful or not.

#### Parameters

*   `id` `string`
*   `title` `string`
*   `date` `string`
*   `body` `string`
*   `tags` `[]string`

#### Example

POST Body:
```json 
{
    "id": "1",
    "title": "latest science shows that potato chips are better for you than sugar",
    "date" : "2016-09-22",
    "body" : "some text, potentially containing simple markup about how potato chips are great",
    "tags" : ["health", "fitness", "science"]
}
```

### GET `/articles/{id}`

Return an Article from the service. The Article must exist and returns a 404 error if it does not. The response contains the Article in full.

#### Parameters

*   `{id}` `string`
    *   The `id` of the Article to be retrieved.

### GET `tag/{tagName}/{date}`

Return a JSON object containing information about all the Articles which have a particular tag and are written on a particular date.

#### Parameters

*   `{tagName}` `string` The tag to match Articles against.
*   `{date}` `string` The date to match Articles against. Formatted as `YYYYMMDD`

#### Example

Response Body:

```json
{
    "tag" : "health",
    "count" : 17,
    "articles" :
        [
            "1",
            "7"
        ],
    "related_tags" :
        [
            "science",
            "fitness"
        ]
}
```

## Assumptions

*   CORS will need to be handled. In this case CORS is permitted for all hosts and all headers are exposed for development purposes and would be configured differently in production.
*   The storage method of the resource does not need to persist between sessions (stored "in memory").
*   The requests to endpoints are usually correctly formatted and response codes do not need to be extremely specific.
*   The development server is accessed via TCP port 8080.

## The Solution

Two main priorities were kept in mind while designing the application: *Maintainability* and *Testability*. These two priorities allow the development of the application to respond quickly to feedback and changing requirements, while maintaining a high level of reliability.

These priorities were met by implementing the SOLID principles and Domain-Driven Design, and following a Clean Architecture structure.

### Maintainability

#### `main.go`



#### `handlers/http.go` (The control layer)



#### `usecases/*.go` (The use case layer)



#### `entities/*.go` (The entities layer)



### Testability











