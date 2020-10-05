package main

import "fmt"

type People struct{
    Nama string;
    Suku string;
    Usia int;

}

func main(){
    var ppl People
    fmt.Println("----Anda Suku Apa?----")

    for i :=0; i < 6; i++{

        fmt.Println("Nama Anda :")
        fmt.Scan(&ppl.Nama)
        fmt.Println("Suku Anda :")
        fmt.Scan(&ppl.Suku)
        fmt.Println("Usia Anda :")
        fmt.Scan(&ppl.Usia)
    }
}