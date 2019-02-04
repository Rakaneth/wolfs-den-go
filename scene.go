package main

import (
	"fmt"
)

//Scene represents a single game screen
type Scene interface {
	Render()
	Enter()
	Exit()
	Name() string
	HandleInput() Command
}

//DefaultScene is the base scene type
type DefaultScene struct{ name string }

//Render draws the scene
func (s *DefaultScene) Render() {}

//Enter runs on scene entry
func (s *DefaultScene) Enter() { fmt.Printf("Entered %s scene", s.name) }

//Exit runs on scene exit
func (s *DefaultScene) Exit() { fmt.Printf("Exited %s scene.", s.name) }

//Name returns the scene name
func (s *DefaultScene) Name() string { return s.name }

//SceneManager manages game scenes
type SceneManager struct {
	scenes   map[string]Scene
	CurScene Scene
}

//NewSceneManager creates a new SceneManager
func NewSceneManager() *SceneManager {
	return &SceneManager{make(map[string]Scene), nil}
}

//SetScene changes the current scene
func (man *SceneManager) SetScene(sceneName string) {
	if man.CurScene != nil {
		man.CurScene.Exit()
	}
	man.CurScene = man.scenes[sceneName]
	man.CurScene.Enter()
}

//AddScene adds a scene to the Manager
func (man *SceneManager) AddScene(scenes ...Scene) {
	for _, scene := range scenes {
		man.scenes[scene.Name()] = scene
	}
}
