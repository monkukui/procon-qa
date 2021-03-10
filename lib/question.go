package lib

import (
	"sort"

	"github.com/monkukui/procon-qa/model"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func EditDistance(s1, s2 []rune) int {
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

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

//（ここまで）

func GetSortedQuestionsByEditDistance(allQuestions model.Questions, queryTitle string) model.Questions {

	var arr ByDistance
	for _, v := range allQuestions {
		title := []rune(v.Title)
		query := []rune(queryTitle)
		distance := EditDistance(query, title)

		// title
		for i := 0; i < len(title)-len(query)+1; i++ {
			distance = min(distance, EditDistance(query, title[i:i+len(query)]))
		}

		arr = append(arr, QuestionWithDistance{
			Question: v,
			Distance: distance,
		})
	}

	sort.Sort(ByDistance(arr))

	// 上位 k 件を返す
	var questions = model.Questions{}
	k := 10
	for i := 0; i < k && i < len(arr); i++ {
		questions = append(questions, arr[i].Question)
	}

	return questions
}
