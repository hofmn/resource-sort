# Go Programm für einheitliche Resource Dateien
Dieses Programm sortiert eine Datei zeilenweise und formartiert sie.
Die Datei muss nach dem Muster: `KEY = VALUE` aufgebaut sein.
Wenn der Key aus mehrern Wörtern besteht, dasnn sollte er mit `_` getrennt sein.

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
2. Wenn ein Key bzw. das erste Wort eines Keys sich ändert. Wird eine Leezeile hinzugefügt.
3. Um Key und Value Spaltenweise darzustellen wird der längste Key gesucht. Allen Keys werden nun Leerzeichen bis zur länge des längsten Keys + 1 angefügt.
4. Alle nicht ASCII Zeichen werden durch ihren Unicode ersetzt.

## Bauen der Binary
1. [Installiere go](https://go.dev/doc/install)
2. Lade dir das Repo herunter.
3. `go build cmd\uniform-resource.go`
