# GoLang Learning

Pour exécuter un fichier `.go` faire : `go run <path>`

> Comments : // One-line or /* Multi-lines */

Le langage Go est organisé en packages, quand on fait un fichier `.go`, on doit s'assurer d'avoir défini un package comme ça :
`package <nom du package>`

On utilisera souvent `package main`.

Quand on fait `import`, on l'écrit comme une fonction, on peut importer un (chaque package est appelé entre ") ou plusieurs packages (entre parenthèse).

`func name(<params>){/* Code */}` Pour faire une fonction.

Fonction principale `main()` à définir à chaque code.

Pour définire des paramètres à notre fonction il faut d'abord le nom du paramètre puis son type (`any` ou `interface{}` pour tout type de valeurs). Après la parenthèse femrante, si on prévoit que notre fonction retourne une valeur, alors on met le type de cette valeur sinon on n'en met pas. Si on ne sait pas de quel type la valeur retourné va être alors utilisé une interface : `interface{}` toujours après la parenthèse fermente.
Pour print des chaînes de caractères dans l'output, il faut importer le package `fmt` de la librarie standard et utiliser `fmt.Println()`. Pour afficher du texte sans saut de ligne : `fmt.Print()`.
Pour déclarer une variable il faut utiliser l'expression `var` puis le nom de la variable et enfin son type. Pour l'affecter, il faut mettre le nom de la variable puis un `=` et enfin sa nouvelle valeur qui doit correspondre à son type. Pour la déclarer et l'affecter d'un coup, il faut mettre le nom de la variable puis l'opérateur `:=` et enfin sa valeur. La variable doit imperativement être utilisé après sa déclaration et sa première affectation, sinon une erreur apparaîtra. Pour contrer cela on peut créer une fonction use() qui quand on l'appelera ne fera rien mais au moins la variable mit dans ses paramètres sera utilisé. Voici le code de cette fonction :
```go
func use(_ any) interface{} {
	return nil
}
```
Et donc lors de la déclaration d'une variable on fera cela :
```go
var variable interface{}; use(variable)
```
À noter : si on change la valeur d'une variable en un valeur qui n'est pas du même type qu'on a selectionner lors de la déclaration de cette variable, alors une erreur se presentera.
Dans les paramètres de fonctions, pour selectionner tout les arguments entrés lors de son appel, il faut choisir le nom de son paramètre et après mettre `...` et ensuite mettre le type des arguments qui seront mit.
Pour récupérer une entrée de la console, on peut utiliser `fmt.Scanln(&variable_to_get_entry)` où `fmt.Scanf("%s", &variable_to_get_entry)` pour spécifier le type. Pour afficher qu'il que chose avant, on peut utiliser `fmt.Print()`.
Pour les boucles, le langage Go ne possède que la boucle `for`, donc pour afficher les nombres de 1 à 5 il faut faire :
```go
for i := 1; i <= 5; i++ {
        print_it(i)
}
```
Pour les boucles while :
```go
i := 0
for {
     fmt.Println(i)
     i++
     if i >= 5 {
        break
     }
}

i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```
Exemple de condition :
```go
if nb == randomNumber {
	print_it("Exact")
} else if nb > randomNumber {
	print_it("Too big")
} else {
	print_it("Too little")
}
```
Alternative aux classes :
```go
// Déclaration d'un type structuré nommé Person
type Person struct {
    Name string
    Age  int
}

// Méthode associée à la struct Person
func (p Person) SayHello() {
    fmt.Printf("Bonjour, je m'appelle %s et j'ai %d ans.\n", p.Name, p.Age)
}

func main() {
	person := Person{name: "John", age: 30}
	person.say("Hello !")
}
```
