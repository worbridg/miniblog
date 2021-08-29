<!DOCTYPE html>
<html>

<head>
    <title>{{ .Title }}</title>
</head>

<body>
    <div id="posted-blog">
        <div id="posted-blog-title">
            <h1>{{ .Title }}</h1>
        </div>
        <div id="posted-blog-content">
            {{ .Content }}
        </div>
    </div>
</body>

</html>