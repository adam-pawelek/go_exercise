# golang cheat sheet

## Structs

Defining
```
type Vertex struct {
  X int
  Y int
}
func main() {
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X, v.Y)
}
```
Pointers to structs
```
v := &Vertex{1, 2}
v.X = 2
```


## Operations on string
```
fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("PORT"))
```


