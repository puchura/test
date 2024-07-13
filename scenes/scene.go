package scenes

// Scene interface defines the methods that all scenes must implement
type Scene interface {
	Init(changeScene func(string))
	Update()
	Draw()
	Unload()
}

// SceneManager handles the current scene and scene transitions
type SceneManager struct {
	currentScene Scene
	scenes       map[string]Scene
	changeScene  func(string)
}

// NewSceneManager creates a new SceneManager
func NewSceneManager() *SceneManager {
	sm := &SceneManager{
		scenes: make(map[string]Scene),
	}
	sm.changeScene = sm.SetScene
	return sm
}

// AddScene adds a new scene to the SceneManager
func (sm *SceneManager) AddScene(name string, scene Scene) {
	sm.scenes[name] = scene
}

// SetScene sets the current scene
func (sm *SceneManager) SetScene(name string) {
	if scene, exists := sm.scenes[name]; exists {
		if sm.currentScene != nil {
			sm.currentScene.Unload()
		}
		sm.currentScene = scene
		sm.currentScene.Init(sm.changeScene)
	}
}

// Update calls the Update method of the current scene
func (sm *SceneManager) Update() {
	if sm.currentScene != nil {
		sm.currentScene.Update()
	}
}

// Draw calls the Draw method of the current scene
func (sm *SceneManager) Draw() {
	if sm.currentScene != nil {
		sm.currentScene.Draw()
	}
}
