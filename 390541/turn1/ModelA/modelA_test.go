package mathops

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        a, b int
        want  int
        err   bool
    }{
        {1, 2, 3, false},
        {2, 3, 5, false},
        {-1, 5, 0, true}, // this should fail with an error
    }

    for _, tt := range tests {
        got, err := Add(tt.a, tt.b)
        if (err != nil) != tt.err {
            t.Errorf("Add(%d, %d) error = %v, wantErr %v", tt.a, tt.b, err, tt.err)
        }
        if got != tt.want {
            t.Errorf("Add(%d, %d) = %v, want %v", tt.a, tt.b, got, tt.want)
        }
    }
}

func TestMultiply(t *testing.T) {
    if got := Multiply(2, 3); got != 6 {
        t.Errorf("Multiply(2, 3) = %v, want %v", got, 6)
    }
}