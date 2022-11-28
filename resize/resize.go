//go:build tinygo.wasm

package main

import (
	"bytes"
	"context"
	"github.com/anthonynsimon/bild/transform"
	plugin "github.com/img-cdn/imgPlugin/proto"
	"image"
	"image/jpeg"
	"strconv"
)

type resizePlugin struct {
}

func (g *resizePlugin) Modify(_ context.Context, request plugin.PluginRequest) (plugin.PluginReply, error) {
	imgReader := bytes.NewReader(request.Image)
	img, _, err := image.Decode(imgReader)
	if err != nil {
		return plugin.PluginReply{Status: false, Image: nil}, err
	}
	width, err := strconv.Atoi(request.Parameters["width"])
	if err != nil {
		return plugin.PluginReply{}, err
	}
	high, err := strconv.Atoi(request.Parameters["high"])
	if err != nil {
		return plugin.PluginReply{}, err
	}
	img = transform.Resize(img, width, high, transform.Linear)
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: int(request.Quality)}); err != nil {
		return plugin.PluginReply{Status: false, Image: nil}, err
	}
	return plugin.PluginReply{Status: true, Image: buf.Bytes()}, nil
}

func main() {
	plugin.RegisterActuator(&resizePlugin{})
}
