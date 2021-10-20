package main
import "fmt"
func main() {
	num := 0
	head := `
           _____  
          / ___ \ 
         | |   | |
         | |   | |
         | |___| |
          \_____/
		`
	fmt.Println("请输入程序员头发数：")
	fmt.Scanln(&num)
	hair := ""
	switch num {
	case 1:
		hair = `
             |
             |`
	case 2:
		hair = `
         \       /
          \     /`
	case 3:
		hair = `
         \   |   /
          \  |  /`
	default:
		fmt.Println("做梦呢，程序员哪有那么多头发")
		return
	}
	fmt.Print(hair)
	fmt.Print(head)
}
