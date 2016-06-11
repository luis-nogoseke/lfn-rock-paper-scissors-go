package main

import (
    "fmt"
    "log"
    "net"
    "encoding/binary"
    "bufio"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Falha na conexão:", err)
        os.Exit(1)
    }
    var user int64
    str1 := "Pedra(0), Papel(1) ou Tesoura(2)?"
    fmt.Println(str1)
    fmt.Scanf("%d", &user)
    if user > 2 || user < 0 {
      fmt.Println("Jogada inválida\nAs jogadas possíveis são: Pedra(0), " +
      "Papel(1) ou Tesoura(2)")
      conn.Close()
      os.Exit(2)
    }

    err = binary.Write(conn, binary.LittleEndian, &user)
      if err != nil {
           fmt.Println("Falha ao enviar jogada:", err)
       conn.Close()
       os.Exit(3)
      }
    var resultado string
    resultado, err = bufio.NewReader(conn).ReadString('\n')
    if err != nil {
           fmt.Println("Falha ao receber resultado:", err)
       conn.Close()
       os.Exit(4)
      }
    fmt.Print(resultado)
    conn.Close()
}
