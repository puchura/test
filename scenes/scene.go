package scenes

type Scene interface {
	Load()
	Unload()
	Update()
	Draw()
}
