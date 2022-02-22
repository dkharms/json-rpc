# 🔥 Json RPC

## 📖 Description

This is simple json-rpc framework for communicating among microservices. It was built on `net/http` library of Go.

## 🎓 Usage

This framework is simple to learn because of it's nature.

### ⚡ Creating server

DI for loggers:

```go
l := log.New(os.Stdin, "server: ", log.Ldate|log.Lshortfile)
s := server.New(l)
```

### ⚡ Creating items

```go
type SumRequest struct {
    A int `json:"a"`
    B int `json:"b"`
}

type SumResponse struct {
    Result int `json:"result"`
}
```

### ⚡ Creating procedures

Adding procedure for remote calling is very simple:

```go
s.AddProcedure(procedure.New("GetSum", "@1",
        func (request *api.JsonRequest, response *api.JsonResponse) error {
            sr := &SumRequest{}
            
            err := request.Get(sr)
            if err != nil {
                return err
            }
            
            res := SumResponse{Result: sr.A + sr.B}
            response.Set(res)
            
            return nil
        }
	)
)
```

### ⚡ Running server

```go
s.Run(":8080")
```

### 🔥 Done
