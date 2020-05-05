package liuhuo

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	word  string
	vec   []float64
	elems []string
}

// model is the total vec
type Model []*Vector

// vec.String() return word aling with vec value
func (v *Vector) String() string {
	return fmt.Sprintf("%s\t%v", v.word, v.vec)
}

// vec.Word() return word
func (v *Vector) Word() string {
	return v.word
}

func (v *Vector) Vector() []float64 {
	return v.vec
}

// Load word vec From the embedding file
func LoadText(filepath string, dimension int) (Model, error) {
	r, err := os.Open(filepath)
	if err != nil {
		fmt.Errorf("unable to open the embedding file %s", err)
		return nil, err
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)
	result := Model{}
	for scanner.Scan() {
		tokens := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(tokens) < (dimension + 1) { // one word + 100 dimension vec
			continue
		}
		vec := []float64{}
		for i := 1; i < len(tokens); i++ {
			val, err := strconv.ParseFloat(tokens[i], 64)
			if err != nil {
				return nil, err
			}
			vec = append(vec, val)
		}
		result = append(result, &Vector{
			word:  tokens[0],
			vec:   vec,
			elems: []string{tokens[0]},
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// find and return specific vector
func (m Model) Find(word string) *Vector {
	for _, vec := range m {
		if vec.word == word {
			return vec
		}
	}
	return nil
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// compute the similar between two vector
func CosineSim(a *Vector, b *Vector) float64 {
	la := len(a.vec)
	lb := len(b.vec)
	if la != lb {
		return float64(-999) // cannot compare two vec with different dimension
	}
	dot := float64(0.0)
	absa := float64(0.0)
	absb := float64(0.0)
	for i := 0; i < la; i++ {
		dot += a.vec[i] * b.vec[i]
		absa += math.Pow(a.vec[i], 2.0)
		absb += math.Pow(b.vec[i], 2.0)
	}

	cosSim := (dot) / (math.Sqrt(absa) + math.Sqrt(absb))
	return cosSim
}
