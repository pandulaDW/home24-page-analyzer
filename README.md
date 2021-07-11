# Web Page Analyzer

This is an application written in Golang and React, which can be used to extract information about a given web page. It provides the below information </br>

    - What HTML version the document is written in?
    - What is the page title?
    - How many headings of what level are in the document?
    - How many internal and external links are in the - - document?
    - How many external links are inaccessible?
    - Does the page contain a login form?

### Deployment

- The application is hosted at heroku and can be accessed using https://home24-page-analyzer.herokuapp.com/

- Docker was used to build both the front-end and back-end using the Dockerfile in this repo.
- Both, React app and the web page analysis API, is served by the Golang server.

### Technical Details

- https://golang.org/pkg/net/http/ package was used to parse html content.
- https://github.com/stretchr/testify package was used as an assertion library.
- Front-end was written in React and Typescript.

### Assumptions

- Inaccessibility is only considered for external links.
- A login form is considered to be a form which contains an element with a certain text wording such as (login).

### Things to improve

- Use a tool like Swagger to properly communicate request, response formats between front-end and
  back-end applications.
- Write more end-to-end testing for the whole application.
- Use a caching mechanism to cache pre-searched urls.
- Use a better logging mechanism in the backend.
