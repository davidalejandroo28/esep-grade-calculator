package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 52, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func GCBuilder(a, e, s int) *GradeCalculator {
	gc := NewGradeCalculator()
	if a >= 0 {
		gc.AddGrade("A1", a, Assignment)
	}
	if e >= 0 {
		gc.AddGrade("E1", e, Exam)
	}
	if s >= 0 {
		gc.AddGrade("S1", s, Essay)
	}
	return gc
}

func TestClearGrades(t *testing.T) {
	// Clear A
	if g := GCBuilder(100, 100, 100).GetFinalGrade(); g != "A" {
		t.Fatalf("expected A, got %s", g)
	}
	// Clear B
	if g := GCBuilder(85, 82, 80).GetFinalGrade(); g != "B" {
		t.Fatalf("expected B, got %s", g)
	}
	// Clear C
	if g := GCBuilder(75, 70, 70).GetFinalGrade(); g != "C" {
		t.Fatalf("expected C, got %s", g)
	}
	// Clear D
	if g := GCBuilder(65, 60, 60).GetFinalGrade(); g != "D" {
		t.Fatalf("expected D, got %s", g)
	}
}

func TestEmptyCategories_AverageZero(t *testing.T) {
	// 100% on assignments and othes empty automatic F
	gc := NewGradeCalculator()
	gc.AddGrade("A1", 100, Assignment)
	if g := gc.GetFinalGrade(); g != "F" {
		t.Fatalf("expected F with only assignments=100, got %s", g)
	}

	// Only Exams grade and others missing automatic F
	gc = NewGradeCalculator()
	gc.AddGrade("E1", 100, Exam)
	if g := gc.GetFinalGrade(); g != "F" {
		t.Fatalf("expected F with only exams=100, got %s", g)
	}

	// Only essays present just like previous test -> F
	gc = NewGradeCalculator()
	gc.AddGrade("S1", 100, Essay)
	if g := gc.GetFinalGrade(); g != "F" {
		t.Fatalf("expected F with only essays=100, got %s", g)
	}
}

func TestAveragingPerCategoryAndRouting(t *testing.T) {
	// This verifies that AddGrade routes to the correct slice by checking
	// the resulting weighted average when multiple items exist per category.

	gc := NewGradeCalculator()
	gc.AddGrade("A1", 100, Assignment)
	gc.AddGrade("A2", 0, Assignment)
	gc.AddGrade("E1", 80, Exam)
	gc.AddGrade("E2", 100, Exam)
	gc.AddGrade("S1", 60, Essay)
	gc.AddGrade("S2", 60, Essay)
	if g := gc.GetFinalGrade(); g != "D" {
		t.Fatalf("expected D from mixed averages, got %s", g)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String() mismatch: %q", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String() mismatch: %q", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String() mismatch: %q", Essay.String())
	}
}
