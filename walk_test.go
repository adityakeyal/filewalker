package walk

import "testing"

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
