#USE API

change .env for configuration Database
<br>
change docker-compose for configuration
<br>
Installation with "make install" in cmd
<br>
Run API with "make dev" in cmd
<br>
<li>POST /v1/login (for Login User)
<li>POST /v1/signup (for Sign up)
<li>GET /v1/user/:id (for User find By Id)
<li>GET /v1/user (for Find all User)
<li>PUT /v1/user/:id (for Edit User by Id)
<li>DELETE /v1/user/:id (for Delete User)
<li>PUT /v1/user/:id/photo (for Update Photo)
<li>GET /v1/buku (for find Buku by Judul)
<li>GET /v1/buku/allBuku (for find all Buku)
<li>POST /v1/buku/:id (for Create new Buku)
<li>PUT /v1/buku/:id (for Edit Buku by Id)
<li>DELETE /v1/buku/:id (for Delete Buku by Id)
<li>POST /v1/activity/pinjam/:id (for Pinjam Buku)
<li>POST /v1/activity/kembali/:id (for Kembali Buku)
<li>GET /v1/activity/historypinjam/:id (for History Pinjam Buku)
<li>GET /v1/activity/historykembali/:id (for History Buku Kembali)