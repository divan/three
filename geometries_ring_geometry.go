package three

//go:generate go run geometry_method_generator/main.go -geometryType RingGeometry -geometrySlug ring_geometry

import (
	"github.com/gopherjs/gopherjs/js"
	"math"
)

// RingGeometry is the two-dimensional primitive ring geometry class. It is typically
// used for creating a ring of the dimensions provided with the 'innerRadius', 'outerRadius',
// 'thetaSegments', 'phiSegments', 'thetaStart', and 'thetaLength' constructor arguments.
type RingGeometry struct {
	*js.Object

	InnerRadius     float64 `js:"innerRadius"`
	OuterRadius     float64 `js:"outerRadius"`
	ThetaSegments   float64 `js:"thetaSegments"`
	PhiSegments     float64 `js:"phiSegments"`
	ThetaStart      float64 `js:"thetaStart"`
	ThetaLength     float64 `js:"thetaLength"`
}

// RingGeometryParameters .
type RingGeometryParameters struct {
	InnerRadius     float64
	OuterRadius     float64
	ThetaSegments   float64
	PhiSegments     float64
	ThetaStart      float64
	ThetaLength     float64
}

// NewRingGeometry creates a new RingGeometry.
func NewRingGeometry(params *RingGeometryParameters) RingGeometry {
	if params.InnerRadius == 0 {
		params.ThetaLength = 0.5
	}
	if params.OuterRadius == 0 {
		params.OuterRadius = 1
	}
	if params.ThetaSegments < 3 {
		params.ThetaSegments = 3
	}
	if params.PhiSegments < 1 {
		params.PhiSegments = 1
	}
	if params.ThetaLength == 0 {
		params.ThetaLength = 2 * math.Pi
	}
	return RingGeometry{
		Object: three.Get("RingGeometry").New(
			params.InnerRadius,
			params.OuterRadius,
			params.ThetaSegments,
			params.PhiSegments,
			params.ThetaStart,
			params.ThetaLength,
		),
	}
}
