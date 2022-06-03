# BookStore

## Project Structure
graph <br/>
&nbsp;    -directives<br/>
&nbsp; &nbsp;&nbsp; &nbsp;      **auth.go**<br/>
&nbsp;  -generated<br/>
&nbsp; &nbsp;&nbsp; &nbsp;     **generated.go**<br/>
&nbsp;   -internal<br/>
&nbsp; &nbsp;&nbsp; &nbsp;    --db<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;         **db.clieng.go**<br/>
&nbsp; &nbsp;&nbsp; &nbsp;      --domain<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;          ---books<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;              **book.go**<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;          ---users<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;              **user.go**<br/>
&nbsp; &nbsp;&nbsp; &nbsp;     --helper<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;         **enum.go**<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;        **price.go**<br/>
&nbsp; &nbsp;&nbsp; &nbsp;    --services<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;         **book_svc.go**<br/>
&nbsp;   -model<br/>
&nbsp; &nbsp;&nbsp; &nbsp;        **model_gen.go**<br/>
&nbsp;   -utils<br/>
&nbsp; &nbsp;&nbsp; &nbsp;     -errors<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;         **errors.go**<br/>
&nbsp; &nbsp;&nbsp; &nbsp;     -jwt<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;        **jwt.go**<br/>
&nbsp; &nbsp;&nbsp; &nbsp;     -service<br/>
&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;&nbsp; &nbsp;         **auth.go**<br/>
&nbsp;  **resolver.go**<br/>
&nbsp;   **schema.graphqls**<br/>
&nbsp;   **server.go**<br/>
  
  ## Available query schema
  <pre>
  "data": 
    "__schema": {
      "queryType": {
        "fields": [
          {
            "name": "listBooks"
          },
          {
            "name": "getBook"
          }
        ]
      }
    }
  }
}
</pre>
## Available Mutation Schema
<pre>
{
  "data": {
    "__schema": {
      "mutationType": {
        "fields": [
          {
            "name": "login"
          },
          {
            "name": "createBook"
          },
          {
            "name": "purchaseBook"
          },
          {
            "name": "reSellBook"
          }
        ]
      }
    }
  }
}
</pre>

### And for the login schema. I have one user {username:'suriya",password:"ayirus"}
  