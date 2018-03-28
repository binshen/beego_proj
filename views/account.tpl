Hello World
<br>
<br>
{{range $index, $item := .accounts}}
   {{$index}} - {{.Id}} - {{.Name}}
   <br>
{{end}}