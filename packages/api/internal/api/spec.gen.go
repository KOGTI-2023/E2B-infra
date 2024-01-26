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
	"h7fq/21j7Jj4LwAA///cgASSBCQAAA==",
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
