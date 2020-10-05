package main

import "fmt"

type Suku struct{
    Suku string;
    Daerah string;
    PopulasiJuta int;
}

func main(){
    var Data [] Suku

        Data = [] Suku{
            Suku{
                Suku : "Suku Jawa",
                Daerah : "Di Pulau Jawa",
                PopulasiJuta : 94,
            },
            Suku{
                Suku : "Suku Sunda",
                Daerah : "Di Pulau Jawa",
                PopulasiJuta : 36,
            },
            Suku{
                Suku : "Suku Melayu",
                Daerah : "Di Pulau Sumatra",
                PopulasiJuta : 8,
            },
            Suku{
                Suku : "Suku Batak",
                Daerah : " Di Pulau Sumatra",
                PopulasiJuta : 8,
            },
            Suku{
                Suku : "Suku Madura",
                Daerah : "Di Pulau Jawa",
                PopulasiJuta : 7,
            },
            Suku{
                Suku : "Suku Betawi",
                Daerah : "Di Pulau Jawa",
                PopulasiJuta : 6,
            },
            Suku{
                Suku : "Suku Minangkabau",
                Daerah : "Di Pulau Sumatra",
                PopulasiJuta : 6,
            },
            Suku{
                Suku : "Suku Bugis",
                Daerah : "Di Pulau Sulawesi",
                PopulasiJuta : 6,
            },
            Suku{
                Suku : "Suku Banjar",
                Daerah : "Di Pulau Kalimantan",
                PopulasiJuta : 4,
            },
            Suku{
                Suku : "Suku Makassar",
                Daerah : "Di Pulau Sulawesi",
                PopulasiJuta : 2,
            },
        }
        fmt.Println("Data Singkat Suku Terbesar Di Indonesia Berdasarkan Populasi Dalam Juta")
        fmt.Println(Data[0])
        fmt.Println(Data[1])
        fmt.Println(Data[2])
        fmt.Println(Data[3])
        fmt.Println(Data[4])
        fmt.Println(Data[5])
        fmt.Println(Data[6])
        fmt.Println(Data[7])
        fmt.Println(Data[8])
        fmt.Println(Data[9])

}