<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="utf-8">
  <title>PhysPeach掲示板</title>
  <meta name="description" content="簡単なシングルスレッド掲示板">
</head>

<body>
  <header>
    <h1>PhysPeach掲示板へようこそ</h1>
    <p>
      喜んでコメントしろ！
    </p>
  </header>
  <form class="postForm">
    <h1>投稿<h1>
    <p>
      <input type="text" class="inputName" placeholder="名前" />
      <br>
      <input type="text" class="inputComment" placeholder="コメント" />
      <br>
      <input type="submit" value="書き込む">
    </p>
  </form>
  <session class="posts">
  <h1>これまでの書き込み</h1>
  {{range $key, $post := .posts}}
    <div class="post">
      <h2>{{$post.Name}} さんの書き込み({{$post.Created}}):</h2>
      <p>{{$post.Text}}</p>
    </div>
  {{end}}
  </session>
  <footer>
    <p>
      Github:
      <a href="https://{{.Github}}">{{.Github}}</a>
    </p>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
