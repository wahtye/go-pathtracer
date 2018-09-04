package main

import (
	"math"

	"github.com/wahtye/gotracer/geometry"
	"github.com/wahtye/gotracer/material"
	"github.com/wahtye/gotracer/render"
)

func main() {
	scene := buildScene(500.)
	render.NewRenderer(500, 500, scene).Render()
}

func buildScene(size float64) *render.Scene {
	scene := render.NewScene()
	diffuseMaterial := material.NewDiffuseMaterial(1.)
	emissiveMaterial := material.NewEmissiveMaterial(1., 5000)
	floorY := .5*size/2 + size/2

	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, 0, .5*size), geometry.NewVector(0, 0, -1.)))
	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, 0, -5*size), geometry.NewVector(0, 0, 1.)))

	scene.AddObject(emissiveMaterial, geometry.NewPlane(geometry.NewVector(.75*size/2+size/2, 0, 0), geometry.NewVector(-1., 0, 0)))
	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(-.75*size/2+size/2, 0, 0), geometry.NewVector(1., 0, 0)))

	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, -.75*size/2+size/2, 0), geometry.NewVector(0, 1., 0)))
	scene.AddObject(diffuseMaterial, geometry.NewPlane(geometry.NewVector(0, floorY, 0), geometry.NewVector(0, -1., 0)))

	sphereCount := 3
	radius := size / 16.
	for i := 0; i < sphereCount; i++ {
		angle := float64(1.8*math.Pi/float64(sphereCount)) * float64(i)
		spherePosition := geometry.NewVector(math.Sin(angle)*size/6.+size/2., floorY-radius, math.Cos(angle)*size/6.-size/4.)
		sphere := geometry.NewSphere(spherePosition, radius)

		scene.AddObject(diffuseMaterial, sphere)
	}

	return scene
}
