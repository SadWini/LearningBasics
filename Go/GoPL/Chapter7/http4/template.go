package main

import "html/template"

// ref: https://stackoverflow.com/questions/21302520/iterating-through-map-in-template
var itemTable = template.Must(template.New("Items").Parse(`
<h1>Items</h1>
<table>
    <tr>
        <th> Item </th>
        <th> Price </th>
    </tr>
    {{ range $k, $v := . }}
        <tr>
            <td>{{ $k }}</td>
            <td>{{ $v }}</td>
        </tr>
    {{end}}
</table>
`))
