| Headers   | Required |
|-----------|:--------:|
| X-Api-Key     | true   |
| Application   | true   |
| Authorization | not all |

| Method | Route  | Ref              |
|--------|--------|------------------|
| GET    | /users      | [Get Users](#GET-users) |
| POST   | /users      | [Post User](#POST-user) |
| PUT    | /users/{id} | [Put User](#PUT-user)   |
| DELETE | /users/{id} | [Delete User](#DELETE-user)|

# GET users

Esta rota precisa do **header de Authorization** como obrigatório

Essa rota possui os seguintes query params:

- ids (opcional)
- page (opcional)
- per_page (opcional)

# POST user

Esta rota deve receber um body com o seguinte formato e todos os campos são obrigatórios:

```json
{
	"name": "John",
	"lastname": "Doe",
	"email": "john.doe@email.com",
	"password": "123456",
	"new_password": "123456"
}
```

**new_password é o equivalente a confirmação de password**

# PUT user

Esta rota precisa do **header de Authorization** como obrigatório

A rota possui o Path param ```/users/{id}``` que corresponde ao ID do usuário

Esta rota deve receber um body com o seguinte formato:

```json
{
	"name": "John",
	"lastname": "Doe",
	"email": "john.doe@email.com",
	"password": "123456",
	"new_password": "123456"
}
```

# DELETE user

Esta rota precisa do **header de Authorization** como obrigatório

A rota possui o Path param ```/users/{id}``` que corresponde ao ID do usuário