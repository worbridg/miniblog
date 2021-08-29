<!doctype html>
<html>

<head>
  <title>Blog Title</title>
</head>

<body>
  <div id="blog-content">
    <a href="/post">create a new blog</a>
    <div id="blog-title">
      <h1>Welcome to our mini blog</h1>
    </div>
    <div id="blog-content">
      Posted blogs
      <ul>
      {{ range $id, $blog := .Blogs }}
        <li><a href="/blog?id={{$id}}">{{ $blog.Title }}</a> <a href="/edit?id={{$id}}">edit</a>|<a href="/delete?id={{$id}}">delete</a></li>
      {{ end }}
      </ul>
    </div>
  </div>
</body>

</html>
