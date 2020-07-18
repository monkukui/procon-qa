package lib

import (
	"github.com/monkukui/procon-qa/model"
  "sort"
  "fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func EditDistance(s, t string) int {
  s1 := []rune(s)
  s2 := []rune(t)
  dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i, c1 := range s1 {
		for j, c2 := range s2 {
			dist := 1
			if c1 == c2 {
				dist = 0
			}
			dp[i+1][j+1] = min(dp[i][j+1]+1, min(dp[i+1][j]+1, dp[i][j]+dist))
		}
	}

	return dp[len(s1)][len(s2)]
}

// Distance 順にソート（ここから）
type QuestionWithDistance struct {
  Question model.Question
  Distance int
}

type ByDistance []QuestionWithDistance

func (a ByDistance) Len() int { return len(a) }
func (a ByDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

//（ここまで）

func GetSortedQuestionsByEditDistance(allQuestions model.Questions, queryTitle string) model.Questions {

  var arr ByDistance
  for _, v := range allQuestions {
    arr = append(arr, QuestionWithDistance{
      Question: v,
      Distance: EditDistance(v.Title, queryTitle),
    })
  }

  sort.Sort(ByDistance(arr));

  // 上位 k 件を返す
  var questions = model.Questions{}
  k := 10
  for i := 0; i < k && i < len(arr); i++ {
    fmt.Println(arr[i].Question.Title)
    fmt.Println(arr[i].Distance)
    questions = append(questions, arr[i].Question)
  }

  return questions;
}
