package handlebars

import "testing"

var inlinePartialTests = []Test{
	{
		"basic inline partial definition and use",
		`{{#*inline "myPartial"}}Hello{{/inline}}{{> myPartial}}`,
		nil, nil, nil, nil,
		"Hello",
	},
	{
		"inline partial with context data",
		`{{#*inline "myPartial"}}<span>{{name}}</span>{{/inline}}{{> myPartial}}`,
		map[string]string{"name": "World"},
		nil, nil, nil,
		"<span>World</span>",
	},
	{
		"inline partial overrides registered partial of same name",
		`{{#*inline "myPartial"}}inline version{{/inline}}{{> myPartial}}`,
		nil, nil, nil,
		map[string]string{"myPartial": "registered version"},
		"inline version",
	},
	{
		"multiple inline partials in same template",
		`{{#*inline "p1"}}first{{/inline}}{{#*inline "p2"}}second{{/inline}}{{> p1}} and {{> p2}}`,
		nil, nil, nil, nil,
		"first and second",
	},
	{
		"inline partial inside each block - should still be accessible",
		`{{#*inline "item"}}<li>{{this}}</li>{{/inline}}{{#each items}}{{> item}}{{/each}}`,
		map[string]interface{}{"items": []string{"a", "b", "c"}},
		nil, nil, nil,
		"<li>a</li><li>b</li><li>c</li>",
	},
	{
		"inline partial scoped to template - does not leak",
		`{{#*inline "myPartial"}}inline content{{/inline}}{{> myPartial}}`,
		nil, nil, nil, nil,
		"inline content",
	},
	{
		"inline partial with nested context",
		`{{#*inline "userCard"}}<div>{{user.name}} ({{user.age}})</div>{{/inline}}{{> userCard}}`,
		map[string]interface{}{"user": map[string]interface{}{"name": "Alice", "age": 30}},
		nil, nil, nil,
		"<div>Alice (30)</div>",
	},
	{
		"inline partial used inside if block",
		`{{#*inline "greeting"}}Hello {{name}}!{{/inline}}{{#if show}}{{> greeting}}{{/if}}`,
		map[string]interface{}{"show": true, "name": "Bob"},
		nil, nil, nil,
		"Hello Bob!",
	},
}

func TestInlinePartials(t *testing.T) {
	launchTests(t, inlinePartialTests)
}
