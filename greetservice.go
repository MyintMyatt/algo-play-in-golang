package main

type GreetService struct{}

func (g *GreetService) Greet(name string) string {
	return "Hello " + name + "!"
}

func (g *GreetService) Calculate(num1 int, num2 int) int {
	return num1 + num2
}
