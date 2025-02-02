package main

import (
	"html/template"
)

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012,
		length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007,
		length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011,
		length("4m24s")},
}

var trackTable = template.Must(template.New("Track").Parse(`
<h1> Tracks </h1>
<table>
<tr style='text-align: left'>
    <th onclick="submitform('Title')">Title
        <form action="" name="Title" method="post">
            <input type="hidden" name="orderby" value="Title"/>
        </form>
    </th>
    <th>Artist
        <form action="" name="Artist" method="post">
            <input type="hidden" name="orderby" value="Artist"/>
        </form>
    </th>
    <th>Album
        <form action="" name="Album" method="post">
            <input type="hidden" name="orderby" value="Album"/>
        </form>
    </th>
    <th onclick="submitform('Year')">Year
        <form action="" name="Year" method="post">
            <input type="hidden" name="orderby" value="Year"/>
        </form>
    </th>
    <th onclick="submitform('Length')">Length
        <form action="" name="Length" method="post">
            <input type="hidden" name="orderby" value="Length"/>
        </form>
    </th>
</tr>
{{range .Data}}
<tr>
    <td>{{.Title}}</td>
    <td>{{.Artist}}</td>
    <td>{{.Album}}</td>
    <td>{{.Year}}</td>
    <td>{{.Length}}</td>
</tr>
{{end}}
</table>

<script>
function submitform(formname) {
    document[formname].submit();
}
</script>
`))
