// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
