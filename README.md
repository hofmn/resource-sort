# Go Programm für einheitliche Resource Dateien
Dieses Programm sortiert eine Datei zeilenweise und formartiert sie nach dem Muster: `KEY = VALUE`. 

## Verwendung

### Binary
Die Binary ist die perfomanteste und einfachste Möglichkeit. Lade die Binary herunter und führe sie über ein Terminal deiner Wahl aus. Gebe die zu sortierenden Datein mit ihrem Dateipfad als Option an.

`.\unifrom-resource.exe <filePath1> <filePath2> ...`

### Mithilfe von go run
1. [Installiere go](https://go.dev/doc/install)
2. Lade dir das Repo herunter.
3. `go run cmd\unifrom-resource.go <filePath1> ...`

## Formatierung
Die Formatierung funktioniert wie folgt:

1. Jede Zeile wird anhand des **"="** in Key und Value gesplittet und alle Leerzeichen und Tabs vor und nach den Werten entfernt.
2. Um Key und Value Spaltenweise darzustellen wird der längste Key gesucht. Allen Keys werden nun Leerzeichen bis zur länge des längsten Keys + 1 angefügt.
3. Alle Umlaute (plus ß) werden aus dem Value entfernt und durch unicode ersetzt.
    ```
    replacements := map[string]string{
    	"ä": "\\u00E4",
    	"Ä": "\\u00C4",
    	"ü": "\\u00FC",
    	"Ü": "\\u00DC",
    	"ö": "\\u00F6",
    	"Ö": "\\u00D6",
    	"ß": "\\u00DF",
    	}
    ```

## Bauen der Binary
1. [Installiere go](https://go.dev/doc/install)
2. Lade dir das Repo herunter.
3. `go build cmd\uniform-resource.go`