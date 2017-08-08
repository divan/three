package main

// The following directive is necessary to make the package coherent:
// +build ignore

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

const geometryTemplate = `package three
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at:
// {{ .Timestamp }}
//
// using the following cmd:
// geometry_method_generator -geometryType {{ .GeometryType }} -geometrySlug {{ .GeometrySlug }}

import "github.com/gopherjs/gopherjs/js"
	
func (g {{ .GeometryType }}) ApplyMatrix(matrix *Matrix4) {
	g.Object.Call("applyMatrix", matrix)
}

func (g {{ .GeometryType }}) RotateX() {
	g.Object.Call("rotateX")
}

func (g {{ .GeometryType }}) RotateY() {
	g.Object.Call("rotateY")
}

func (g {{ .GeometryType }}) RotateZ() {
	g.Object.Call("rotateZ")
}

func (g {{ .GeometryType }}) Translate() {
	g.Object.Call("translate")
}

func (g {{ .GeometryType }}) Scale() {
	g.Object.Call("scale")
}

func (g {{ .GeometryType }}) LookAt() {
	g.Object.Call("lookAt")
}

func (g {{ .GeometryType }}) FromBufferGeometry(geometry Geometry) {
	g.Object.Call("fromBufferGeometry")
}

func (g {{ .GeometryType }}) Center() {
	g.Object.Call("center")
}

func (g {{ .GeometryType }}) Normalize() {{ .GeometryType }} {
	g.Object.Call("normalize")
	return g
}

func (g {{ .GeometryType }}) ComputeFaceNormals() {
	g.Object.Call("computeFaceNormals")
}

func (g {{ .GeometryType }}) ComputeVertexNormals(areaWeighted bool) {
	g.Object.Call("computeVertexNormals", areaWeighted)
}

func (g {{ .GeometryType }}) ComputeFlatVertexNormals() {
	g.Object.Call("computeFlatVertexNormals")
}

func (g {{ .GeometryType }}) ComputeMorphNormals() {
	g.Object.Call("computeMorphNormals")
}

func (g {{ .GeometryType }}) ComputeLineDistances() {
	g.Object.Call("computeLineDistances")
}

func (g {{ .GeometryType }}) ComputeBoundingBox() {
	g.Object.Call("computeBoundingBox")
}

func (g {{ .GeometryType }}) ComputeBoundingSphere() {
	g.Object.Call("computeBoundingSphere")
}

func (g {{ .GeometryType }}) Merge(geometry Geometry, matrix Matrix4, materialIndexOffset float64) {
	g.Object.Call("merge", geometry, matrix, materialIndexOffset)
}

func (g {{ .GeometryType }}) MergeMesh(mesh Mesh) {
	g.Object.Call("mergeMesh", mesh.getInternalObject())
}

func (g {{ .GeometryType }}) MergeVertices() {
	g.Object.Call("mergeVertices")
}

func (g {{ .GeometryType }}) SortFacesByMaterialIndex() {
	g.Object.Call("sortFacesByMaterialIndex")
}

func (g {{ .GeometryType }}) ToJSON() interface{} {
	return g.Object.Call("toJSON")
}

// func (g {{ .GeometryType }}) Clone() {{ .GeometryType }} {
// 	return g.Object.Call("clone")
// }

func (g {{ .GeometryType }}) Copy(source Object3D, recursive bool) *{{ .GeometryType }} {
	return &{{ .GeometryType }}{Object: g.getInternalObject().Call("copy", source.getInternalObject(), recursive)}
}

func (g {{ .GeometryType}}) Dispose() {
	g.Object.Call("dispose")
}

func (g {{ .GeometryType }}) getInternalObject() *js.Object {
	return g.Object
}

`

var (
	geometryType = flag.String("geometryType", "", "Name of class that extends Geometry e.g. CircleGeometry")
	geometrySlug = flag.String("geometrySlug", "", "Slugified name of class e.g. circle_geometry")
)

func main() {
	flag.Parse()

	if *geometryType == "" {
		log.Fatal("a geometry name argument must be provided (e.g. -geometryType CircleGeometry)")
	}
	if *geometrySlug == "" {
		log.Fatal("a geometry slug argument must be provided (e.g. -geometrySlug circle_geometry)")
	}

	filePath := fmt.Sprintf("./gen_geometry_%s.go", *geometrySlug)

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = template.Must(template.New("").Parse(geometryTemplate)).Execute(f, struct {
		Timestamp    time.Time
		GeometryType string
		GeometrySlug string
	}{
		Timestamp:    time.Now(),
		GeometryType: *geometryType,
		GeometrySlug: *geometrySlug,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Generated file: %s", filePath)
}