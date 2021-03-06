package ami

// This file was generated by static_resolver_ami_generate.go; DO NOT EDIT.

// StaticImages is a map that holds the list of AMIs to be used by for static resolution
var StaticImages = map[string]map[int]map[string]string{
	"AmazonLinux2": {
		ImageClassGPU: {
			"eu-west-1": "ami-0706dc8a5eed2eed9",
			"us-east-1": "ami-058bfb8c236caae89",
			"us-west-2": "ami-0731694d53ef9604b",
		},
		ImageClassGeneral: {
			"eu-west-1": "ami-0c7a4976cb6fafd3a",
			"us-east-1": "ami-0440e4f6b9713faf6",
			"us-west-2": "ami-0a54c984b9f908c81",
		},
	},
	"Ubuntu1804": {ImageClassGeneral: {
		"eu-west-1": "ami-07036622490f7e97b",
		"us-east-1": "ami-06fd8200ac0eb656d",
		"us-west-2": "ami-6322011b",
	}},
}
