package geofabrik

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetIndex(t *testing.T) {
	viper.Set("log", true)

	tests := []struct {
		name    string
		myURL   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test",
			myURL:   GeofabrikIndexURL,
			wantErr: false,
		},
	}
	for x := range tests {
		mt := tests[x]
		t.Run(mt.name, func(t *testing.T) {
			index, err := GetIndex(mt.myURL)
			if (err != nil) != mt.wantErr {
				t.Errorf("GetIndex() error = %v, wantErr %v", err, mt.wantErr)
			}
			if len(index.Features) < 10 {
				t.Errorf("GetIndex() error I should have more features!!!")
			}
			c, err := Convert(index)
			if c == nil || err != nil {
				t.Errorf("GetIndex() error cant convert !!!\n%v", err)
			}
			if e, err := c.GetElement("france"); err != nil || e == nil {
				t.Errorf("GetIndex() error cant find element !!!\nconfig=%v\nerr=%v", c, err)
			}
		})
	}
}
