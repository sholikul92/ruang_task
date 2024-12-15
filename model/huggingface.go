package model

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type AIRequest struct {
	Inputs Inputs `json:"inputs"`
}

type TapasResponse struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

type TextGenerationPayload struct {
	Model       string     `json:"model"`
	Messages    []Messages `json:"messages"`
	MaxContents int        `json:"max_contents"`
	Stream      bool       `json:"stream"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type TextGenerationResponse struct {
	Choices []Choices `json:"choices"`
}

type Choices struct {
	Index    int      `json:"index"`
	Messages Messages `json:"message"`
}
