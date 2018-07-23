# simpleprogress

### requirements
- go1.10.3

### sample

```
import "simpleprogress"

func main() {
	message := "waiting"
        ch := simpleprogress.Simpleprogress(message)
	sample()	
        ch <-"end"
}

func sample(){
	...
}

```

sample 実行中に waiting... と表示される
