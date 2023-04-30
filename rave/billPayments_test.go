package rave

import (
	"fmt"
	"testing"
)

var rv = Rave{
	false,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxx-X",
	"FLWSECK-xxxxxxxxxxxxxxxxxxxxx-X",
	true,
}

var bp = BillPayments{r}

func TestBillpayment_GetBillCategories(t *testing.T) {
	tests := []struct {
		name    string
		filters []BillCategoryFilter
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCategories, err := bp.GetBillCategories(tt.filters...)
			fmt.Printf("categories returned: %+v", gotCategories)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBillCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
