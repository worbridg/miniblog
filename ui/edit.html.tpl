<!DOCTYPE html>
<html>

<head>
    <title>Edit a blog</title>
</head>

<body>
    <div id="blog-edit-content">
        <div id="blog-edit-title">
            <h1>Post a blog</h1>
        </div>
        <div id="blog-edit-form">
            <form action="/edit?id={{.ID}}" method="post">
                <div id="blog-edit-form-title">
                    <h2>Blog Title</h2>
                    <input type="text" name="blog-title" value="{{ .Blog.Title }}"/>
                </div>
                <div id="blog-edit-form-content">
                    <h2>Blog Content</h2>
                    <textarea name="blog-content">{{ .Blog.Content }}</textarea>
                </div>
                <button type="submit" value="edit">Update</button>
            </form>
        </div>
    </div>
</body>

</html>