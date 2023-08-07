package server

import "math"

// Тип данных для RPC сервера.
// Может быть любым пользовательским типом.
// Все экспортируемые методы этого типа
// будут доступны для удаленного вызова.
type Server int

// Точка на плоскости.
type Point struct {
	X, Y float64
}

// Аргумент для функции Dist.
type Points struct {
	A, B Point
}

// Метод  доступен для удаленного вызова.
func (s *Server) Dist(args Points, resp *float64) error {

	*resp = math.Sqrt((args.A.X-args.B.X)*(args.A.X-args.B.X) + (args.A.Y-args.B.Y)*(args.A.Y-args.B.Y))
	return nil
}
