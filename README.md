# go-vue-loader: A Vue-loader written in Go

go-vue-loader is a vue-loader written in Go.
If you use Vue via CDN for small-sized project, but you still want to enjoy the `.vue`, this lib is made for you !

## Features
 - [X] Compiles a `.vue` file, returning a `.js` readable by the client
 - [ ] Returns any the components from a directory and allowing to retrieve several components in one request
 - [ ] Cache system, to maximize performance
 - [ ] On the flight minifier

## Not supported (yet?)
 - Event handling in template -> '@' makes the compiler crash
 - Style -> Ignored by the compiler
 - Any kind of preprocessor (Jade, Babel, ...) -> The `.js` will be return as-is
 - Local component registration -> Obviously, since the client is choosing which component to load

## Examples
See the [examples folder](examples)

