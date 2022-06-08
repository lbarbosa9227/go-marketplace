package models

import "testing"

func NewReview(stars int, comment string) *CreateReviewCmd {
	return &CreateReviewCmd{
		Stars:   stars,
		Comment: comment,
	}

}

func Test_withCorrectParams(t *testing.T) {

	r := NewReview(4, "the iphone looks good")
	err := r.validate()
	if err != nil {
		t.Error("the validation they dont pass")
		t.Fail()
	}
}

func Test_shuldFailWithWrongNumberOfStars(t *testing.T) {

	r := NewReview(8, "the iphone looks relly good")
	err := r.validate()
	if err != nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}
