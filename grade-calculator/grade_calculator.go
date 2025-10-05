package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numerical := gc.calculateNumericalGrade()

	if numerical >= 90 {
		return "A"
	} else if numerical >= 80 {
		return "B"
	} else if numerical >= 70 {
		return "C"
	} else if numerical >= 60 {
		return "D"
	}
	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignAvg := computeAverageByType(gc.grades, Assignment)
	examAvg := computeAverageByType(gc.grades, Exam)
	essayAvg := computeAverageByType(gc.grades, Essay)

	weighted := float64(assignAvg)*0.50 +
		float64(examAvg)*0.35 +
		float64(essayAvg)*0.15

	return int(weighted)
}

func computeAverageByType(all []Grade, t GradeType) int {
	sum, count := 0, 0
	for _, g := range all {
		if g.Type == t {
			sum += g.Grade
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
