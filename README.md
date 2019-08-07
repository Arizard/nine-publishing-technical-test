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

Example body:
    
    {
        "id": "1",
        "title": "latest science shows that potato chips are better for you than sugar",
        "date" : "2016-09-22",
        "body" : "some text, potentially containing simple markup about how potato chips are great",
        "tags" : ["health", "fitness", "science"]
    }







