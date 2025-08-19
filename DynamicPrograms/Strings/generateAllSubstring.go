package main


import "fmt"

func main(){
	s := "abc"
	generateAllSubstrings(s)
}

func generateAllSubstrings(s string){
	n := len(s)

	for i:=0;i<=n;i++{
		for j:=i+1;j<=n;j++{
			fmt.Println(s[i:j])
		}
	}
}