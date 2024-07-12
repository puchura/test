package scenes

type Scene interface {
	Load()
	Unload()
	Update()
	Draw()
}

type SceneChanger interface {
	ChangeScene(newScene Scene)
}
