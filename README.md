# rest-api
[![Go Report Card](https://goreportcard.com/badge/github.com/nick1729/rest-api?style=flat-square)](https://goreportcard.com/badge/github.com/nick1729/rest-api)

Test task
<p><code>git clone github.com/nick1729/rest-api.git</code></p>

<br>Just for testing REST-API and Docker</br>

<br>Build and run: <code>docker-compose up -d</code></br>

<br>Start page: <code>http://127.0.0.1:8080/</code></br>
<br>Add new user: <code>http://127.0.0.1:8080/users?firstname=Vasya&lastname=Ivanov&email=vasya_vanov@mail.ru&age=22/</code></br>
<br>Edit user by UUID: <code>http://127.0.0.1:8080/users/?id=08a3a799-cca3-4d67-a3d6-eb116185466a&firstname=Gena&lastname=Petrov&email=gena@mail.gg&age=27</code></br>
<br>Show user data by UUID: <code>http://127.0.0.1:8000/users/a5657a25-b62d-45f8-96f6-41aab04f9ec</code></br>
<br>Show all users: <code>http://127.0.0.1:8080/users/all</code></br>
<br>Delete user by UUID: <code>http://127.0.0.1:8080/users/delete/9c683260-5481-47ca-a399-093d6b3233b5</code></br>