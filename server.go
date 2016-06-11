package main

import (
    "fmt"
    "net"
    "time"
    "encoding/binary"
    "math/rand"
    "os"
)

func printJogada(jogada int64) string{
    if jogada == 0{
        return "Pedra"
    } else if jogada == 1{
        return "Papel"
    } else{
        return "Tesoura"
    }
}

func handleConnection(conn net.Conn) {
    var user int64
    err := binary.Read(conn, binary.LittleEndian, &user)
      if err != nil {
            fmt.Println("Falha ao receber jogada:", err)
        os.Exit(3)
      }
    if user > 2 || user < 0 {
      fmt.Println("Jogada inválida:", user)
      os.Exit(4)
    }
    rand.Seed(time.Now().UTC().UnixNano())
    pc := rand.Int63n(3)
    resultado := "Sua jogada: " + printJogada(user) + "  Jogada do pc: " +
    printJogada(pc)

    difference := (user - pc) % 3
    if difference < 0{
        difference +=3
    }

    switch difference {
    case 0:
            resultado  += "  Empate!\n"
    case 1:
            resultado  +=  "  Você venceu!\n"
    case 2:
            resultado  +=  "  Você Perdeu!\n"
    }
    _, err = fmt.Fprint(conn, resultado)
    if err != nil {
            fmt.Println("Falha enviar resultado:", err)
        os.Exit(5)
      }
}

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Erro ao iniciar servidor:", err);
        os.Exit(1)
    }
    fmt.Println("Servidor disponível");
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Erro ao aceitar conexão:", err);
            os.Exit(2)
        }
        go handleConnection(conn)
    }
}
