# Contacts App (Go)

A [Go](https://go.dev/) web application made in the spirit of the [official Flask example][htmx-proj] application
that accompanies the book [Hypermedia Systems][htmx-book]. You can use this application as a starting point
to follow along with the book starting at [HTMX Patterns](https://hypermedia.systems/htmx-patterns/).

## Prerequisites

- [Go 1.23](https://go.dev/) or later
- [Air](https://github.com/air-verse/air) for live reloading:
  ```shell
  go install github.com/air-verse/air@latest
  ```

## Development

```shell
# Start the Tailwind CSS build process
make watch-css

# In a separate terminal start the development server (Linux / MacOS)
make dev
```

## Tailwind CSS Development Notes

You can develop and build everything using only the TailwindCSS CLI (installed via `make tailwindcss`) but you likely will
not have Tailwind Intellisense in your IDE. If you need Intellisense installing the NodeJS version should enable it - just 
run `npm install` in the project root and you should be good to go.


## Production

Everything needed to run the application is in the `dist` directory. You can serve the files using any web server or reverse proxy.

```shell
# Build the application
make build
```

## Note

- This project uses the [Library Manager][libman] [CLI][libman-cli] to manage client-side libraries. You do not need it,
  but I find if you have .NET on your system it is really handy for handling non-bundled client-side assets.

  ```shell
  # Install the client-side libraries
  libman restore
  ```

- The data repository is rudimentary as that is not the point of the book, and by using a JSON file
  as the data store, it is easy to follow along without setting up a database and worrying about additional dependencies.
- I will win no awards for design in my lifetime and I am okay with this.

[htmx]: https://htmx.org 'High power tools for HTML'
[htmx-book]: https://hypermedia.systems/ 'Hypermedia Systems Book'
[flask]: https://flask.palletsprojects.com/ 'Flask - A minimal web framework for Python'
[htmx-proj]: https://github.com/bigskysoftware/contact-app 'Contact App - official'
[libman]: https://devblogs.microsoft.com/dotnet/library-manager-client-side-content-manager-for-web-apps/ 'Client-side content manager for web apps'
[libman-cli]: https://learn.microsoft.com/en-us/aspnet/core/client-side/libman/libman-cli
