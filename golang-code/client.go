package main

import (
    "fmt"
    "encoding/json"
    "net/http"
   
)
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
      r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
func main() {
    resp, err := http.Get("http://server-clusterip:8080/message")
    if err != nil {
        print(err)
    }
    defer resp.Body.Close()
    var data map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
        print(err)
    }
    
    var v interface{}
     v = data["message"] 
    fmt.Println(Reverse(v.(string)))

}
