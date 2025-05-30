package vector

type Document struct {
	ID        string
	Embedding []float32
	Content   string
}

type Result struct {
	ID         string
	Content    string
	Similarity float64
}

type VectorDB struct {
	documents map[string]Document
}

func NewVectorDB() *VectorDB {
	return &VectorDB{documents: make(map[string]Document)}
}

func (vc *VectorDB) AddDocuments(documents []Document) {
	for _, document := range documents {
		document.Embedding, _ = OllamaEnbedding(document.Content)
		vc.documents[document.ID] = document
	}
}

func (vc *VectorDB) Query(text string) *Result {
	queryEnbedding, _ := OllamaEnbedding(text)

	result := &Result{Similarity: -1}
	for _, doc := range vc.documents {
		sim := CosineSimilarity(queryEnbedding, doc.Embedding)
		if sim > result.Similarity {
			result.Similarity = sim
			result.ID = doc.ID
			result.Content = doc.Content
		}
	}
	return result
}
