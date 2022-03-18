# Kazanımlar

1. Command-line flag
    - Herhangi bir kodu run veya build ederken kullanılan flagları eklememizi sağlar.
    - Tüm flag'ler tanımlandıktan sonra flag.Parse() çağırılır.
    - Flag olarak string, int ve bool kullanılabilir.

    ```go
    import "flag"
    csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'") // csv burda flag, problems.csv ise default value, en son değer de command line'da --help yazıldığında gelen bilgi. flag.String sonucunda string pointer'ı dönüyor.
    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")
    flag.Parse()
    ```

2. Dosya okuma (os.Open())
    - Dosya okuma işlemleri os modülü ile yapılır.

    ```go
    import "os"
    file, err := os.Open(*csvFilename)
    if err != nil{
        exit(fmt.Sprintf("Failed to open csv file: %s\n", *csvFilename))
        os.Exit(1)
    }
    ```

3. csv dosyası okuma(csv.NewReader(file))
    - csv parse etme işlemleri "encoding/csv" modülü ile yapılır.

    ```go
    import "encoding/csv"
    r := csv.NewReader(file)
    lines, err := r.ReadAll()
    if err != nil {
        exit("Failed to parse csv file")
    }
    ```
