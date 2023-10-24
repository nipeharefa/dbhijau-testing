package demo

import "testing"

func TestExampleNewClient(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name: "Test connection",
			want: "arango",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExampleNewClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExampleNewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExampleNewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
