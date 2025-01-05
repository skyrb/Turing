package main

func fetchData(db *sql.DB) ([]Data, error) {
    defer db.Close()  // Close the database connection after the function returns

    query := "SELECT * FROM data"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    defer rows.Close() // Close the rows object after fetching data
    defer sql.FinishTx(...) // If handling transactions, ensure they are finished

    var data []Data
    for rows.Next() {
        var d Data
        err := rows.Scan(&d)
        if err != nil {
            return nil, err
        }
        data = append(data, d)
    }

    return data, nil
}


func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close() // Close the file after processing

    // Read the file and process it

    return nil
}

func fetchExternalData() ([]byte, error) {
    client := &http.Client{}
    resp, err := client.Get("https://example.com/data")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close() // Close the response body

    defer func() {
        err := client.CloseIdleConnections() // Close idle connections
        if err != nil {
            // Log the error or handle it
        }
    }()

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return data, nil
}

func main(){
	fetchExternalData()
}