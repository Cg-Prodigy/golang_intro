package main
import(
	"fmt"
)

func main()  {
	// fmt.Println("message")
	// message:="hi international space station"
	// fmt.Println(message)
	// for i:=0;i<len(message);i++{
	// 	char:=message[i] // this is the byte or int8 value of the character
	// 	if char>='a' && char<='z'{ // comparing the int8 value of char against int8 value of a and z
	// 		char-=13
	// 		if char<'a'{
	// 			char+=26
	// 		}
	// 	}
	// 	fmt.Printf("%c",char)
	// }
	// fmt.Println()
	cipher:="uv vagreangvbany fcnpr fgngvba"
	fmt.Println("Decipher message")
	for i:=0;i<len(cipher);i++{
		char:=cipher[i]
		if char>='a' && char<='z'{
			char+=13
			if char>'z'{
				char-=26
			}
		}
		fmt.Printf("%c",char)
	}
	fmt.Println()
}