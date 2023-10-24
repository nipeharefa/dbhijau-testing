package demo

import "testing"

func TestExampleNewClient(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		setEnv  func(*testing.T)
		wantErr bool
	}{
		{
			name: "Test connection fail",
			setEnv: func(t *testing.T) {
				t.Setenv("DB_URL", "http://badurl")
			},
			wantErr: true,
		},
		{
			name:   "Test connection success",
			setEnv: func(t *testing.T) {},
			want:   "arango",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnv(t)
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
