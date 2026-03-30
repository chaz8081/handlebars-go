# Changelog

### hbs 4.0.0 _(March 29, 2026)_

Full Handlebars v4 spec compliance.

**Breaking Changes**
- Minimum Go version bumped to 1.22

**Handlebars v4 Features**
- Subexpressions (`{{foo (bar baz)}}`)
- Raw blocks (`{{{{raw}}}}...{{{{/raw}}}}`)
- Inline partials (`{{#*inline "name"}}...{{/inline}}`)
- Partial blocks (`{{#> partial}}fallback{{/partial}}`) and `@partial-block`
- Decorators (`{{* decorator}}` and `{{#* decorator}}`)
- Block parameters (`{{#each items as |item index|}}`)
- `@data` variables (`@root`, `@first`, `@last`, `@index`, `@key`)
- `helperMissing` / `blockHelperMissing` hooks
- Strict mode (`Template.SetStrict()`)
- Chained else (`{{else if}}`)
- Dynamic partials via subexpressions (`{{> (lookup . "name")}}`)
- Whitespace control (`{{~expression~}}`)
- `includeZero` flag for helpers
- Raw block helpers with `options.RawContent()`

**New Features**
- CLI tools: `handlebars-lint`, `handlebars-vars`, `handlebars-gen`
- VS Code extension with syntax highlighting, autocomplete, and lint integration
- Template variable extraction (`ExtractVariables`)
- Template validation (`Validate`)
- AST access (`Template.AST()`, `NewTemplateFromAST()`)

**Improvements**
- Mustache spec updated to v1.4.3
- CI/CD with GitHub Actions (Go 1.22–1.26)
- golangci-lint v2 with staticcheck, errcheck, govet, unused, gofmt
- Dependabot for Go modules and GitHub Actions
- Comprehensive nil-safety handling

### Handlebars 3.0.0 _(July 14, 2021)_
- Hard Fork from Raymond. Rebrand "handlebars"
- Start with major version number 3 tracking supported version of Handlebars.js


### Raymond 2.0.2 _(March 22, 2018)_

- [IMPROVEMENT] Add `RemoveHelper` and `RemoveAllHelpers` functions
- [IMPROVEMENT] Add the #equal helper (#7)
- [IMPROVEMENT] Add struct tag template variable support (#8)

### Raymond 2.0.1 _(June 01, 2016)_

- [BUGFIX] Removes data races [#3](https://github.com/aymerick/raymond/issues/3) - Thanks [@markbates](https://github.com/markbates)

### Raymond 2.0.0 _(May 01, 2016)_

- [BUGFIX] Fixes passing of context in helper options [#2](https://github.com/aymerick/raymond/issues/2) - Thanks [@GhostRussia](https://github.com/GhostRussia)
- [BREAKING] Renames and unexports constants:

  - `handlebars.DUMP_TPL`
  - `lexer.ESCAPED_ESCAPED_OPEN_MUSTACHE`
  - `lexer.ESCAPED_OPEN_MUSTACHE`
  - `lexer.OPEN_MUSTACHE`
  - `lexer.CLOSE_MUSTACHE`
  - `lexer.CLOSE_STRIP_MUSTACHE`
  - `lexer.CLOSE_UNESCAPED_STRIP_MUSTACHE`
  - `lexer.DUMP_TOKEN_POS`
  - `lexer.DUMP_ALL_TOKENS_VAL`


### Raymond 1.1.0 _(June 15, 2015)_

- Permits templates references with lowercase versions of struct fields.
- Adds `ParseFile()` function.
- Adds `RegisterPartialFile()`, `RegisterPartialFiles()` and `Clone()` methods on `Template`.
- Helpers can now be struct methods.
- Ensures safe concurrent access to helpers and partials.

### Raymond 1.0.0 _(June 09, 2015)_

- This is the first release. Raymond supports almost all handlebars features. See https://github.com/aymerick/raymond#limitations for a list of differences with the javascript implementation.
