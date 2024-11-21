package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

// допишите код реализации методов интерфейса
// ...
// String должен уметь сериализовать переменную типа в строку.
func (n *NetAddress) String() string {
	return n.Host + ":" + strconv.Itoa(n.Port)
}

// Set связывает переменную типа со значением флага
// и устанавливает правила парсинга для пользовательского типа.
func (n *NetAddress) Set(flagValue string) error {
	t := strings.Split(flagValue, ":")
	n.Host = t[0]
	n.Port, _ = strconv.Atoi(t[1])
	return nil
}

func main() {
	addr := new(NetAddress)
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)
	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
