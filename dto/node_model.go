package dto

type UpdateLocationReq struct {
}

type NodeDistanceDTO struct {
	NodeNo   string
	Distance float64
}

type NodeCoordDTO struct {
	NodeId string
	X      float64
	Y      float64
	Z      float64
}
