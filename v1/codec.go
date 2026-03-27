// Package pluginsv1 implements the Kleff plugin gRPC contract.
//
// Wire format: The codec registered here overrides gRPC's default protobuf
// codec with a JSON codec. This lets the package ship without a protoc build
// step while still speaking standard gRPC framing. Both the platform (client)
// and all Go plugins (server) import this package and therefore share the same
// codec automatically via init().
//
// To switch to protobuf wire format: generate pb.go files from the .proto
// files in plugin-sdk/v1/proto/ using protoc, then delete this file and
// replace types.go / *.go stubs with the generated output.
package pluginsv1

import (
	"encoding/json"

	"google.golang.org/grpc/encoding"
)

func init() {
	// Override the default "proto" codec so all gRPC calls in this process
	// use JSON encoding. Both the platform client and plugin servers register
	// this via their shared import of this package.
	encoding.RegisterCodec(jsonCodec{})
}

type jsonCodec struct{}

func (jsonCodec) Name() string { return "proto" }

func (jsonCodec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
