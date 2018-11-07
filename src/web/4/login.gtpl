<!DOCTYPE html>
<html>
  <head>
    <meta name="generator" content="HTML Tidy for HTML5 for Linux version 5.2.0">
    <title></title>
  </head>
  <body>
    <form method="post" action="http://localhost:9090/login?username=liyunteng">
      用户名： <input type="text" name="username">
      密码： <input type="password" name="password">
      <input type="submit" value="登陆">
      <input type="hidden" name="toke" value="{{.}}" />
    </form>
  </body>
</html>
