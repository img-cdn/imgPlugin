package main

// This file is here to test features of the Go compiler.

import (
	"context"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	plugin "github.com/img-cdn/imgPlugin/proto"
)

type greyPlugin struct {
}

func (g *greyPlugin) Modify(context context.Context, request plugin.PluginRequest) (plugin.PluginReply, error) {
	image, err := imgio.Open(request.SrcPath)
	if err != nil {
		return plugin.PluginReply{Message: err.Error()}, err
	}
	image = effect.Grayscale(image)
	err = imgio.Save(request.DestPath, image, imgio.JPEGEncoder(int(request.Quality)))
	if err != nil {
		return plugin.PluginReply{Message: "failed "}, err
	}
	return plugin.PluginReply{Message: "succ"}, nil
}

func main() {
	plugin.RegisterActuator(&greyPlugin{})
}
