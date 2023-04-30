package rave

import (
	"fmt"
	"testing"
)

var rv = Rave{
	true,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxx-X",
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxx-X",
	false,
}

var bp = BillPayments{rv}

func TestBillpayment_GetBillCategories(t *testing.T) {
	tests := []struct {
		name    string
		filters []BillCategoryFilter
		wantErr bool
	}{
		{
			name:    "Get Billers test 1: No filter",
			filters: nil,
			wantErr: false,
		},
		{
			name:    "Get Billers test 2: DSTV Only",
			filters: []BillCategoryFilter{Cable},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCategories, err := bp.GetBillCategories(tt.filters...)
			fmt.Printf("categories returned: %+v\n", gotCategories)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBillCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
