package walk

import (
	"os"
	"os/exec"
	"testing"
)

var CB = func(path string, filename string) {

}

func TestScan(t *testing.T) {
	type args struct {
		startFolder string
		callBack    func(string, string)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "printTmp",
			args: args{
				startFolder: "D:\\code\\igv\\nrithai\\BR-GV-FOR-THAGMO-PROD-REL-19-22052018",
				callBack:    CB,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Scan(tt.args.startFolder, tt.args.callBack)

		})
	}
}
func TestCrasher(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		Scan("e:/", CB)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	t.Log(err)
	// if e, ok := err.(*exec.ExitError); ok && !e.Success() {
	// 	return
	// }
	// t.Fatalf("process ran with err %v, want exit status 1", err)
}
