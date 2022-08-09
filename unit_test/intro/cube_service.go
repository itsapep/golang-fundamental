package intro

type cubeService struct {
	repo CubeRepository
}

func (c *cubeService) GetVolumeCube(cube Cube) (float64, error) {
	return c.repo.Volume()
}

func (c *cubeService) GetAreaCube(cube Cube) (float64, error) {
	return c.repo.Area()
}

func NewCubeService(repo CubeRepository) cubeService {
	return cubeService{repo: repo}
}
