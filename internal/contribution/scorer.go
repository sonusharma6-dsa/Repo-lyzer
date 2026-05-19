package contribution

import (
	"strings"

	"github.com/agnivo988/Repo-lyzer/internal/github"
)

// ContributionScore represents the computed contribution friendliness metrics.
type ContributionScore struct {
	Score      float64  `json:"score"`
	Level      string   `json:"level"`
	Strengths  []string `json:"strengths"`
	Weaknesses []string `json:"weaknesses"`
}

var readmeKeywords = []string{
	"installation",
	"setup",
	"getting started",
	"usage",
}

func hasSetupSection(readmeContent string) bool {
	contentLower := strings.ToLower(readmeContent)
	for _, kw := range readmeKeywords {
		if strings.Contains(contentLower, kw) {
			return true
		}
	}
	return false
}

// Calculate computes the contribution score based on various repository metrics.
func Calculate(
	hasContributing bool,
	readmeContent string,
	issues []github.Issue,
	commits []github.Commit,
	contributors []github.Contributor,
) ContributionScore {
	return ContributionScore{}
}
