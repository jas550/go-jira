package onpremise

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestStatusCategoryService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/rest/api/2/statuscategory"

	raw, err := os.ReadFile("../testing/mock-data/all_statuscategories.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEdpoint)
		fmt.Fprint(w, string(raw))
	})

	statusCategory, _, err := testClient.StatusCategory.GetList(context.Background())
	if statusCategory == nil {
		t.Error("Expected statusCategory list. StatusCategory list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestStatusCategoryService_Get(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/2/statuscategory/1"

	raw, err := os.ReadFile("../testing/mock-data/status_category.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	statusCategory, _, err := testClient.StatusCategory.Get(context.Background(), "1")

	if err != nil {
		t.Errorf("Error given: %s", err)
	} else if statusCategory == nil {
		t.Error("Expected status category. StatusCategory is nil")
	}
}
