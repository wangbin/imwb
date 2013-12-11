<form action="" method="post">
      {{ with .Form }}
      {{ .NonFieldErrors }}
      <p><label>Name:</label><input type="text" name="username" value="{{ .Name }}" /></p>
      {{ with .ErrorMap }}
      {{ if .username }}
      {{.username }}
      {{ end }}
      {{ end }}
      <p><label>Password:</label><input type="password" name="password" value="{{ .Password }}" /></p>
      {{ with .ErrorMap }}
      {{ if .password }}
      {{.username }}
      {{ end }}
      {{ end }}

      {{ end }}
        <input type="submit" name="submit"/>
</form>