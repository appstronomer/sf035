package proverb

import (
	"bufio"
	"fmt"
	"io"
	"math/rand/v2"
)

type Proverb struct {
	// Интерфейс структуры позволят только читать из неё - поэтому
	// не использую никакие синхронизирующие сущности
	data []string
}

func (p Proverb) Rnd() string {
	return p.data[rand.IntN(len(p.data))]
}

func NewProverb(fileReader io.ReadCloser) (*Proverb, error) {
	data := []string{}
	defer fileReader.Close()
	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		message := fmt.Sprintf("%v\n", scanner.Text())
		data = append(data, message)
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return &Proverb{data}, nil
}
