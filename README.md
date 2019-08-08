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

![Clean Architecture https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html](CleanArchitecture.jpg)

Two main priorities were kept in mind while designing the application: *Maintainability* and *Testability*. These two priorities allow the development of the application to respond quickly to feedback and changing requirements, while maintaining a high level of reliability.

These priorities were met by implementing the SOLID principles and Domain-Driven Design, and following a Clean Architecture structure.


### `main.go`

This file is kept quite short. All that needs to happen here is to initialise the dependencies of the system. No business logic is performed here.

All dependencies are inject "down" through components to the core of the application. This helps with maintainability as the components become more easily reasoned about.

`jsonPresenter` is of the type `presenters.JSONPresenter` which implements the interface `presenters.Presenter` and has the role of presenting some domain-specific input into a JSON formatted output.

`JSONHandler` is of the type `handlers.Handler`, and only becomes a JSON handler when the `Presenter` field is set to an object of type `presenters.JSONPresenter`. 

This aids maintainability and testability by removing the two-way dependency on a presentation layer. 

For example, a future developer might like to re-write the application to return YAML instead. They would simply need to create a new struct which implements the interface `presenters.Presenter` and change only the handler instantiation and nothing else (other than imports).

Similarly, the field `ArticleRepository` can easily be substituted with other structs which implement the interface `entities.ArticleRepository`.

This design allows two key results:
    * The decision of which data storage infrastructure to use can be left as late as possible - the cost of changing the data storage infrastructure is low for the maintainer.
    * The risk of breaking the application due to changes is mitigated and reduced.
    * The `handlers.Handler` struct can be tested end-to-end by injecting fake dependencies, e.g. `FakeJSONPresenter` and `InMemoryArticleRepository`, instead of being forced to test on live/test infrastructure.

### `infrastructure/*.go` (The infrastructure layer)

The file `infrastructure/inmemoryarticlerepository.go` contains a concrete implementation of the interface `entities.ArticleRepository`. This implementation simply stores the articles in memory. It abstracts database operations into the methods `Add`, `Get` and `Find`.

Any concrete implementation which implements `entities.ArticleRepository` can be substituted at a low cost to the developer.

This is especially useful when migrating between database technologies. For example, a migration of data from a MongoDB database to a PostgreSQL database will not require a complete code refactor - the application is **extensible** by writing a new concrete implementation of `entities.ArticleRepository`.

### `handlers/http.go` (The control layer)

The `handlers.Handler` struct acts as a control layer for the application. Here, the HTTP request is handled and processed into relevant domain variables to be passed to a use case. Additionally, CORS is handled to allow remote hosts to access the HTTP REST API.

There are static handlers for the error codes 400, 404, 500. There are 3 handler functions which run use cases: `SubmitArticleHandler`, `GetArticleHandler` and `GetArticlesByTagHandler`. The role of this layer is to convert HTTP request into domain variables.

If a programmer wanted to implement **authentication**, this is the layer where it would occur. It prevents the layers beneath (use cases and entities) from depending on the authentication method.

For example, an authentication struct which implements a new `auth.UserValidator` interface could be injected into the handler, and the handler could validate a Bearer Token, either returning a *403 Forbidden* error code or continuing with execution.

### `usecases/*.go` (The use case layer)

The use case layer contains the specific use cases for the Article resource. A user can submit an article, get an article, and get a list of articles and related tags based on a given tag and date.

These domain variables are injected from the `handlers.Handler` into the relevant use cases.

The concrete implementation of the `entities.ArticleRepository` interface which was injected into the handler from `main.go` must also be injected further down into the use case.

The use case has all its dependencies injected from the control layer. This way, the domain logic of submission, retrieval or otherwise regarding Articles has no concrete dependency on other parts of the application.

This allows the use case to be **testable** by injecting a fake concrete implementation of `entities.ArticleRepository` and other faked variables. This way there is no need to run a web server or a database to test whether the use cases are working correctly, **allowing for automated or parameterized testing and facilitating CI/CD.**

### `entities/*.go` (The entities layer)

At the core of the application is the entities layer. The struct `entities.Article` defines the fields that an Article has, and in future could contain methods which implement domain or business logic concerning a single or multiple instances of `entities.Article`.

Notice that there are no imports, no external dependencies. This is because the concept of Article should not be concerned with how it is stored or represented to a user.

The interface `entities.ArticleRepository` is defined here. This allows dependencies to flow inwards, and outwardly changes cannot affect the inner layers.

## Summary

With this architecture, there is a high level of maintainability and testability of the project, allowing for **automated testing**, **continous integration/continous delivery**, and **agile iteration** - responding to the changing needs of the business or customer.

From a programmer's perspective, the codebase is easy to reason about - conversion from domain variables to JSON (or XML or YAML) happens in the presentation layer. Persistence of data (whether by database, file or API) exists in the infrastructure layer












