### 准备(linux系统)
#### 1：安装ollama： curl -fsSL https://ollama.com/install.sh | sh
#### 2：ollama pull bge-m3:latest

```go
package main

import (
	"fmt"

	"github.com/button-chen/vector"
)

func main() {
	vecdb := vector.NewVectorDB()

	vecdb.AddDocuments([]vector.Document{
		{
			ID:      "1",
			Content: "The sky is blue because of Rayleigh scattering.",
		},
		{
			ID:      "2",
			Content: "Leaves are green because chlorophyll absorbs red and blue light.",
		},
	})

	res := vecdb.Query("Why is the sky blue?")

	fmt.Printf("ID: %v\nSimilarity: %v\nContent: %v\n", res.ID, res.Similarity, res.Content)
}

output:

ID: 1
Similarity: 0.7347239885667152
Content: The sky is blue because of Rayleigh scattering.

```
